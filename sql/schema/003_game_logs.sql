-- +goose Up

CREATE TABLE game_logs (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    game_id INT NOT NULL REFERENCES games(id) ON DELETE CASCADE,
    user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    shelf CHAR NOT NULL DEFAULT('W'),
    UNIQUE(game_id,user_id)
);

-- +goose Down

DROP TABLE game_logs;