-- +goose Up

CREATE TABLE games (
    id INT PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    genres TEXT[] NOT NULL,
    image TEXT NOT NULL,
    release_date TIMESTAMP NOT NULL
);

-- +goose Down

DROP TABLE games;