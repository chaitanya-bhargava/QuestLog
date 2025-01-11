-- name: CreateGameLog :one
INSERT INTO game_logs(id,created_at,updated_at,game_id,user_id,shelf)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING *;
