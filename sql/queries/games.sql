-- name: CreateGame :one
INSERT INTO games(id,created_at,updated_at,name,genres,image,release_date)
VALUES ($1,$2,$3,$4,$5,$6,$7)
ON CONFLICT(id)
DO UPDATE SET updated_at=CURRENT_TIMESTAMP
RETURNING *;

-- name: GetGamesByGameLog :many
SELECT g.id AS id, g.name AS name, g.genres AS genres, g.image AS image, g.release_date AS release_date, g.created_at AS created_at, g.updated_at AS updated_at
FROM games g
JOIN game_logs gl ON g.id = gl.game_id
WHERE gl.user_id = $1 AND gl.shelf = $2;