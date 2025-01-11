-- +goose Up

CREATE TABLE game_logs (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    game_id INT NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    shelf CHAR NOT NULL DEFAULT('W'),
    UNIQUE(game_id,user_id,shelf)
);

-- +goose Down

DROP TABLE game_logs;