-- +goose Up
ALTER TABLE users ADD COLUMN username TEXT NOT NULL DEFAULT '';
UPDATE users SET username = split_part(email, '@', 1);
ALTER TABLE users ADD CONSTRAINT users_username_unique UNIQUE (username);

-- +goose Down
ALTER TABLE users DROP CONSTRAINT users_username_unique;
ALTER TABLE users DROP COLUMN username;
