package main

import (
	"time"

	"github.com/chaitanya-bhargava/QuestLog/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID                string    `json:"id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Name              string    `json:"name"`
	Email             string    `json:"email"`
	Avatarurl         string    `json:"avatar_url"`
	Provider          string    `json:"provider"`
	Nickname          string    `json:"nickname"`
	AccessToken       string    `json:"access_token"`
	AccessTokenSecret string    `json:"access_token_secret"`
	RefreshToken      string    `json:"refresh_token"`
	ExpiresAt         time.Time `json:"expires_at"`
	IDToken           string    `json:"id_token"`
	Username          string    `json:"username"`
}

func databaseUsertoUser(dbUser database.User) User {
	return User{
		ID:                dbUser.ID,
		CreatedAt:         dbUser.CreatedAt,
		UpdatedAt:         dbUser.UpdatedAt,
		Name:              dbUser.Name,
		Email:             dbUser.Email,
		Avatarurl:         dbUser.Avatarurl,
		Provider:          dbUser.Provider,
		Nickname:          dbUser.Nickname,
		AccessToken:       dbUser.AccessToken,
		AccessTokenSecret: dbUser.AccessTokenSecret,
		RefreshToken:      dbUser.RefreshToken,
		ExpiresAt:         dbUser.ExpiresAt,
		IDToken:           dbUser.IDToken,
		Username:          dbUser.Username,
	}
}

type PublicProfile struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Avatarurl string `json:"avatar_url"`
	Username  string `json:"username"`
}

func databaseUserToPublicProfile(dbUser database.GetUserByUsernameRow) PublicProfile {
	return PublicProfile{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		Avatarurl: dbUser.Avatarurl,
		Username:  dbUser.Username,
	}
}

type GameLog struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	GameID int `json:"game_id"`
	UserID string `json:"user_id"`
	Shelf string `json:"shelf"`
	Rating int `json:"rating"`
}

func databaseGameLogtoGameLog(dbGameLog database.GameLog) GameLog {
	return GameLog{
		ID: dbGameLog.ID,
		CreatedAt: dbGameLog.CreatedAt,
		UpdatedAt: dbGameLog.UpdatedAt,
		GameID: int(dbGameLog.GameID),
		UserID: dbGameLog.UserID,
		Shelf: dbGameLog.Shelf,
		Rating: int(dbGameLog.Rating),
	}
}

type Game struct {
	ID          int      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string   `json:"name"`
	Genres      []Genre  `json:"genres"`
	Image       string   `json:"image"`
	ReleaseDate time.Time `json:"release_date"`
}

func databaseGametoGame(dbGame database.Game) Game {
	genres := make([]Genre, 0, len(dbGame.Genres))
	for _, genre := range dbGame.Genres {
		genres = append(genres, Genre{Name: genre})
	}

	return Game{
		ID:          int(dbGame.ID),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Name:        dbGame.Name,
		Genres:      genres,
		Image:       dbGame.Image,
		ReleaseDate: dbGame.ReleaseDate,
	}
}

type Genre struct {
	Name string `json:"name"`
}

func databaseGamestoGames(dbGames []database.GetGamesByGameLogRow) []Game {
	temp := make([]Game, 0, len(dbGames))
	for _, dbGame := range dbGames {
		genres := make([]Genre, 0, len(dbGame.Genres))
		for _, genre := range dbGame.Genres {
			genres = append(genres, Genre{Name: genre})
		}

		temp = append(temp, Game{
			ID:          int(dbGame.ID),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Name:        dbGame.Name,
			Genres:      genres,
			Image:       dbGame.Image,
			ReleaseDate: dbGame.ReleaseDate,
		})
	}
	return temp
}