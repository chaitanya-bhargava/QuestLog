package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/chaitanya-bhargava/QuestLog/internal/database"
	"github.com/dimuska139/rawg-sdk-go/v3"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateGameLog(w http.ResponseWriter, r *http.Request,user database.User,rawgCfg rawg.Config){
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

	rawgClient := rawg.NewClient(http.DefaultClient,&rawgCfg)
    
    ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*10000))
	defer cancel()
    data,  err := rawgClient.GetGame(ctx, params.GameID)
	if err != nil {
		respondWithError(w,400,fmt.Sprint("Error fetching games: ",err))
		return
	}

	_, err = apiCfg.DB.CreateGame(r.Context(),database.CreateGameParams{
		ID: int32(params.GameID),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: data.Name,
		Genres: GenreParser(data.Genres),
		Image: data.ImageBackground,
		ReleaseDate: data.Released.Time,
	})

	if err!= nil {
		respondWithError(w, 400, fmt.Sprint("Error creating game:", err))
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
		respondWithError(w, 400, fmt.Sprint("Error creating game_log:", err))
		return
	}

	respondWithJSON(w,201,databaseGameLogtoGameLog(gameLog))
}