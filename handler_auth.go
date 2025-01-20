package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/chaitanya-bhargava/QuestLog/internal/database"
	"github.com/markbates/goth/gothic"
)

func handlerBeginAuth(w http.ResponseWriter, r *http.Request) {
	if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
		respondWithJSON(w,200,gothUser)
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}

func (apiCfg *apiConfig) handlerCompleteAuth(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Something went wrong:", err))
		return
	}

	dbUser, err := apiCfg.DB.CreateUser(r.Context(),database.CreateUserParams{
		ID: user.UserID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: user.Name,
		Email: user.Email,
		Avatarurl: user.AvatarURL,
		Provider: user.Provider,
		Nickname: user.NickName,
		AccessToken: user.AccessToken,
		AccessTokenSecret: user.AccessTokenSecret,
		RefreshToken: user.RefreshToken,
		ExpiresAt: user.ExpiresAt,
		IDToken: user.IDToken,
	})

	if err!= nil {
		respondWithError(w, 400, fmt.Sprint("Error creating user:", err))
		return
	}

	session, err := gothic.Store.Get(r, "session-name")
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error retrieving session: %v", err))
        return
    }

    session.Values["user_id"] = dbUser.ID
    err = session.Save(r, w)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error saving session: %v", err))
        return
    }

	frontendURL := os.Getenv("FRONTEND_URL")
    if frontendURL == "" {
        respondWithError(w,http.StatusInternalServerError,"No redirect URL in environment")
    }
    http.Redirect(w, r, frontendURL, http.StatusTemporaryRedirect)
}

func handlerLogout(w http.ResponseWriter, r *http.Request) {
    session, err := gothic.Store.Get(r, "session-name")
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error retrieving session: %v", err))
        return
    }

    session.Values = make(map[interface{}]interface{})
    session.Options.MaxAge = -1 
    err = session.Save(r, w)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error saving session: %v", err))
        return
    }

    respondWithJSON(w, http.StatusOK, map[string]string{"message": "Logged out successfully"})
}

func (apiCfg *apiConfig) handlerGetCurrentUser(w http.ResponseWriter, r *http.Request) {
    session, err := gothic.Store.Get(r, "session-name")
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error retrieving session: %v", err))
        return
    }

    userID, ok := session.Values["user_id"].(string)
    if !ok {
        respondWithError(w, http.StatusUnauthorized, "User not authenticated")
        return
    }

    dbUser, err := apiCfg.DB.GetUserByID(r.Context(), userID)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching user: %v", err))
        return
    }

    respondWithJSON(w, http.StatusOK, databaseUsertoUser(dbUser))
}