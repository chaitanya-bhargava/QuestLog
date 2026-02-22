-- name: CreateUser :one
INSERT INTO users(id,created_at,updated_at,name,email,avatarURL,provider,nickname,access_token,access_token_secret,refresh_token,expires_at,id_token,username)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)
ON CONFLICT(id)
DO UPDATE SET updated_at=CURRENT_TIMESTAMP, username=$14
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id=$1;

-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE access_token=$1;

-- name: GetUserByUsername :one
SELECT id, name, email, avatarurl, username FROM users WHERE username=$1;