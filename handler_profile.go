package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/chaitanya-bhargava/QuestLog/internal/database"
	"github.com/go-chi/chi"
)

func (apiCfg *apiConfig) handlerGetPublicProfile(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	if username == "" {
		respondWithError(w, 400, "Username is required")
		return
	}

	dbUser, err := apiCfg.DB.GetUserByUsername(r.Context(), username)
	if err != nil {
		if err == sql.ErrNoRows {
			respondWithError(w, 404, "Profile not found")
			return
		}
		respondWithError(w, 500, fmt.Sprint("Error fetching profile: ", err))
		return
	}

	respondWithJSON(w, 200, databaseUserToPublicProfile(dbUser))
}

func (apiCfg *apiConfig) handlerGetPublicGamesByShelf(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	if username == "" {
		respondWithError(w, 400, "Username is required")
		return
	}

	shelf := r.URL.Query().Get("shelf")
	if shelf == "" {
		respondWithError(w, 400, "Shelf query parameter is required")
		return
	}

	dbUser, err := apiCfg.DB.GetUserByUsername(r.Context(), username)
	if err != nil {
		if err == sql.ErrNoRows {
			respondWithError(w, 404, "Profile not found")
			return
		}
		respondWithError(w, 500, fmt.Sprint("Error fetching user: ", err))
		return
	}

	games, err := apiCfg.DB.GetGamesByGameLog(r.Context(), database.GetGamesByGameLogParams{
		UserID: dbUser.ID,
		Shelf:  shelf,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprint("Error fetching games: ", err))
		return
	}

	respondWithJSON(w, 200, databaseGamestoGames(games))
}
