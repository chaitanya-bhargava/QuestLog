package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

func transformGameDetails(data *rawg.GameDetailed) map[string]interface{} {
	developers := make([]map[string]interface{}, 0)
	for _, dev := range data.Developers {
		developers = append(developers, map[string]interface{}{
			"id":   dev.ID,
			"name": dev.Name,
			"slug": dev.Slug,
		})
	}

	publishers := make([]map[string]interface{}, 0)
	for _, pub := range data.Publishers {
		publishers = append(publishers, map[string]interface{}{
			"id":   pub.ID,
			"name": pub.Name,
			"slug": pub.Slug,
		})
	}

	genres := make([]map[string]interface{}, 0)
	for _, genre := range data.Genres {
		genres = append(genres, map[string]interface{}{
			"id":   genre.ID,
			"name": genre.Name,
			"slug": genre.Slug,
		})
	}

	tags := make([]map[string]interface{}, 0)
	for _, tag := range data.Tags {
		tags = append(tags, map[string]interface{}{
			"id":   tag.ID,
			"name": tag.Name,
			"slug": tag.Slug,
		})
	}

	return map[string]interface{}{
		"id":              data.ID,
		"name":           data.Name,
		"description":    data.Description,
		"released":       data.Released.Time,
		"background_image": data.ImageBackground,
		"website":        data.Website,
		"metacritic":     data.Metacritic,
		"rating":         data.Rating,
		"ratings_count":  data.RatingsCount,
		"developers":     developers,
		"publishers":     publishers,
		"genres":        genres,
		"tags":          tags,
	}
}

func handlerGetGameByID(w http.ResponseWriter, r *http.Request, rawgCfg rawg.Config) {
	gameID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Error processing search query: ", err))
		return
	}

	rawgClient := rawg.NewClient(http.DefaultClient, &rawgCfg)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*10000))
	defer cancel()
	data, err := rawgClient.GetGame(ctx, gameID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Error fetching game: ", err))
		return
	}

	transformedData := transformGameDetails(data)
	respondWithJSON(w, 200, transformedData)
}