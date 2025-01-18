-- name: CreateGameLog :one
INSERT INTO game_logs(id,created_at,updated_at,game_id,user_id,shelf)
VALUES ($1,$2,$3,$4,$5,$6)
ON CONFLICT(game_id,user_id)
DO UPDATE SET updated_at=CURRENT_TIMESTAMP,shelf=$6
RETURNING *;
