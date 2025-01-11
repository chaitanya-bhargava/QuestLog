// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: game_logs.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createGameLog = `-- name: CreateGameLog :one
INSERT INTO game_logs(id,created_at,updated_at,game_id,user_id,shelf)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING id, created_at, updated_at, game_id, user_id, shelf
`

type CreateGameLogParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	GameID    int32
	UserID    uuid.UUID
	Shelf     string
}

func (q *Queries) CreateGameLog(ctx context.Context, arg CreateGameLogParams) (GameLog, error) {
	row := q.db.QueryRowContext(ctx, createGameLog,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.GameID,
		arg.UserID,
		arg.Shelf,
	)
	var i GameLog
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.GameID,
		&i.UserID,
		&i.Shelf,
	)
	return i, err
}
