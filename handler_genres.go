package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dimuska139/rawg-sdk-go/v3"
)

func handlerGetGamesByGenre(w http.ResponseWriter, r *http.Request, rawgCfg rawg.Config){

	genre := r.URL.Query().Get("genre")
	pageNo, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		respondWithError(w,400,fmt.Sprint("Error processing search query: ",err))
		return
	}

	type SearchResults struct {
		Data []*rawg.Game `json:"data"`
		Total int `json:"total"`
	}

	rawgClient := rawg.NewClient(http.DefaultClient,&rawgCfg)

	filter := rawg.NewGamesFilter().
		SetGenres(genre).
        SetPage(pageNo).
        SetPageSize(12).
        ExcludeCollection(1)
    
    ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*10000))
	defer cancel()
    data, total, err := rawgClient.GetGames(ctx, filter)
	if err != nil {
		respondWithError(w,400,fmt.Sprint("Error fetching games: ",err))
		return
	}

	respondWithJSON(w,200,SearchResults{
		Data: data,
		Total: total,
	})
}

func handlerGetGenres(w http.ResponseWriter, r *http.Request, rawgCfg rawg.Config){

	type SearchResults struct {
		Data []*rawg.Genre `json:"data"`
		Total int `json:"total"`
	}

	rawgClient := rawg.NewClient(http.DefaultClient,&rawgCfg)

    ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*10000))
	defer cancel()
    data, total, err := rawgClient.GetGenres(ctx, 1, 50,"games_count")
	if err != nil {
		respondWithError(w,400,fmt.Sprint("Error fetching genres: ",err))
		return
	}

	respondWithJSON(w,200,SearchResults{
		Data: data,
		Total: total,
	})
}