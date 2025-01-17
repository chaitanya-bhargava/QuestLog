// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: games.sql

package database

import (
	"context"
	"time"

	"github.com/lib/pq"
)

const createGame = `-- name: CreateGame :one
INSERT INTO games(id,created_at,updated_at,name,genres,image,release_date)
VALUES ($1,$2,$3,$4,$5,$6,$7)
ON CONFLICT(id)
DO UPDATE SET updated_at=CURRENT_TIMESTAMP
RETURNING id, created_at, updated_at, name, genres, image, release_date
`

type CreateGameParams struct {
	ID          int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Genres      []string
	Image       string
	ReleaseDate time.Time
}

func (q *Queries) CreateGame(ctx context.Context, arg CreateGameParams) (Game, error) {
	row := q.db.QueryRowContext(ctx, createGame,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		pq.Array(arg.Genres),
		arg.Image,
		arg.ReleaseDate,
	)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		pq.Array(&i.Genres),
		&i.Image,
		&i.ReleaseDate,
	)
	return i, err
}