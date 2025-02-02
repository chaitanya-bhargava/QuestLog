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

const getGamesByGameLog = `-- name: GetGamesByGameLog :many
SELECT g.id AS id, g.name AS name, g.genres AS genres, g.image AS image, g.release_date AS release_date, g.created_at AS created_at, g.updated_at AS updated_at
FROM games g
JOIN game_logs gl ON g.id = gl.game_id
WHERE gl.user_id = $1 AND gl.shelf = $2
`

type GetGamesByGameLogParams struct {
	UserID string
	Shelf  string
}

type GetGamesByGameLogRow struct {
	ID          int32
	Name        string
	Genres      []string
	Image       string
	ReleaseDate time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (q *Queries) GetGamesByGameLog(ctx context.Context, arg GetGamesByGameLogParams) ([]GetGamesByGameLogRow, error) {
	rows, err := q.db.QueryContext(ctx, getGamesByGameLog, arg.UserID, arg.Shelf)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetGamesByGameLogRow
	for rows.Next() {
		var i GetGamesByGameLogRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			pq.Array(&i.Genres),
			&i.Image,
			&i.ReleaseDate,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
