package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/chaitanya-bhargava/QuestLog/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateGameLog(w http.ResponseWriter, r *http.Request,user database.User){
	type parameters struct {
		GameID int `json:"game_id"`
		Shelf string `json:"shelf"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Error parsing JSON:", err))
		return
	}

	gameLog, err := apiCfg.DB.CreateGameLog(r.Context(),database.CreateGameLogParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		GameID: int32(params.GameID),
		UserID: user.ID,
		Shelf: params.Shelf,
	})

	if err!= nil {
		respondWithError(w, 400, fmt.Sprint("Error creating user:", err))
		return
	}

	respondWithJSON(w,201,databaseGameLogtoGameLog(gameLog))
}