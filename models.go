package main

import (
	"time"

	"github.com/chaitanya-bhargava/QuestLog/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey string `json:"api_key"`
}

func databaseUsertoUser(dbUser database.User) User {
	return User{
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
		APIKey: dbUser.ApiKey,
	}
}

type GameLog struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	GameID int `json:"game_id"`
	UserID uuid.UUID `json:"user_id"`
	Shelf string `json:"shelf"`
}

func databaseGameLogtoGameLog(dbGameLog database.GameLog) GameLog {
	return GameLog{
		ID: dbGameLog.ID,
		CreatedAt: dbGameLog.CreatedAt,
		UpdatedAt: dbGameLog.UpdatedAt,
		GameID: int(dbGameLog.GameID),
		UserID: dbGameLog.UserID,
		Shelf: dbGameLog.Shelf,
	}
}