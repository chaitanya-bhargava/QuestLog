-- name: CreateGame :one
INSERT INTO games(id,created_at,updated_at,name,genres,image,release_date)
VALUES ($1,$2,$3,$4,$5,$6,$7)
ON CONFLICT(id)
DO UPDATE SET updated_at=CURRENT_TIMESTAMP
RETURNING *;