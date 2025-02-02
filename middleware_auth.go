package main

import (
	"fmt"
	"net/http"

	"github.com/chaitanya-bhargava/QuestLog/internal/auth"
	"github.com/chaitanya-bhargava/QuestLog/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter,r *http.Request) {
		userID, err := auth.GetUserID(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprint("Auth error: ",err))
			return
		}

		user, err := apiCfg.DB.GetUserByID(r.Context(),userID)
		if err != nil {
			respondWithError(w, 400, fmt.Sprint("Couldn't fetch user: ",err))
			return
		}
		handler(w,r,user)
	}
}