package main

import (
	"net/http"
	"os"

	"github.com/chaitanya-bhargava/QuestLog/internal/database"
	"github.com/dimuska139/rawg-sdk-go/v3"
)

func (apiCfg *apiConfig) middlewareAuthRawg(
	handler func(http.ResponseWriter, *http.Request, database.User, rawg.Config),
) func(http.ResponseWriter, *http.Request) {
	return apiCfg.middlewareAuth(func(w http.ResponseWriter, r *http.Request, user database.User) {
		rawgAPIKey := os.Getenv("RAWG_API_KEY")
		if rawgAPIKey == "" {
			respondWithError(w, 400, "API_KEY NOT SET")
			return
		}

		rawgCfg := rawg.Config{
			ApiKey:   rawgAPIKey,
			Language: "en",
			Rps:      5,
		}

		handler(w, r, user, rawgCfg)
	})
}