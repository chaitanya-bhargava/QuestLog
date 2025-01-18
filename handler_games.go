package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/chaitanya-bhargava/QuestLog/internal/database"
	"github.com/dimuska139/rawg-sdk-go/v3"
)

func GenreParser(genres []*rawg.Genre) []string {
	genreNames := make([]string,0)
	for _,genre := range genres{
		genreNames=append(genreNames,genre.Name)
	}
	return genreNames
}

func (apiCfg *apiConfig) handlerCreateGame(w http.ResponseWriter, r *http.Request,user database.User, rawgCfg rawg.Config){
	type parameters struct {
		ID int `json:"id"`
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
    data,  err := rawgClient.GetGame(ctx, params.ID)
	if err != nil {
		respondWithError(w,400,fmt.Sprint("Error fetching games: ",err))
		return
	}



	game, err := apiCfg.DB.CreateGame(r.Context(),database.CreateGameParams{
		ID: int32(params.ID),
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

	respondWithJSON(w,201,databaseGametoGame(game))
}

func (apiCfg *apiConfig) handlerGetGamesByGameLog(w http.ResponseWriter, r *http.Request,user database.User) {
	shelf := r.URL.Query().Get("shelf")
	
	games,err := apiCfg.DB.GetGamesByGameLog(r.Context(),database.GetGamesByGameLogParams{
		UserID: user.ID,
		Shelf: shelf,
	})

	if err!= nil {
		respondWithError(w, 400, fmt.Sprint("Error fetching games:", err))
		return
	}

	respondWithJSON(w,200,databaseGamestoGames(games))
}