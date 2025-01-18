-- name: CreateUser :one
INSERT INTO users(id,created_at,updated_at,name,email,avatarURL,provider,nickname,access_token,access_token_secret,refresh_token,expires_at,id_token)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)
ON CONFLICT(id)
DO UPDATE SET updated_at=CURRENT_TIMESTAMP
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id=$1;

-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE access_token=$1;