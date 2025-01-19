-- name: CreateGameLog :one
INSERT INTO game_logs(id,created_at,updated_at,game_id,user_id,shelf)
VALUES ($1,$2,$3,$4,$5,$6)
ON CONFLICT(game_id,user_id)
DO UPDATE SET updated_at=CURRENT_TIMESTAMP,shelf=$6
RETURNING *;

-- name: GetGameLogByGameID :one
SELECT * FROM game_logs WHERE game_id=$1 AND user_id=$2;

-- name: DeleteGameLogByGameID :exec
DELETE FROM game_logs WHERE game_id=$1 AND user_id=$2;