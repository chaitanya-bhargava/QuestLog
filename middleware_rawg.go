package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dimuska139/rawg-sdk-go/v3"
)

type rawgHandler func(http.ResponseWriter, *http.Request, rawg.Config)

func middlewareRawg(handler rawgHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter,r *http.Request) {
		rawgAPIKey := os.Getenv("RAWG_API_KEY")
		if rawgAPIKey == "" {
			log.Fatal("RAWG_API_KEY is not found in the environment")
			respondWithError(w,400,"API_KEY NOT SET")
			return
		}

		rawgCfg := rawg.Config{
			ApiKey: rawgAPIKey,
			Language: "en",
			Rps: 5,
		}
		handler(w,r,rawgCfg)
	}
}