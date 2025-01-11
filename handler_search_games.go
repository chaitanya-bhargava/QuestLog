package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dimuska139/rawg-sdk-go/v3"
)

func handlerSearchGames(w http.ResponseWriter, r *http.Request, rawgCfg rawg.Config){

	queryString := r.URL.Query().Get("query")
	pageNo, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		respondWithError(w,400,fmt.Sprint("Error processing search query: ",err))
		return
	}

	type SearchResults struct {
		Data []*rawg.Game
		Total int
	}

	

	rawgClient := rawg.NewClient(http.DefaultClient,&rawgCfg)

	filter := rawg.NewGamesFilter().
        SetSearch(queryString).
        SetPage(pageNo).
        SetPageSize(10).
        ExcludeCollection(1)
    
    ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*5000))
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