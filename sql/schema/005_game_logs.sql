-- +goose Up

ALTER TABLE game_logs ADD CONSTRAINT fk_game_id FOREIGN KEY (game_id) REFERENCES games(id) ON DELETE CASCADE;

-- +goose Down
ALTER TABLE game_logs DROP CONSTRAINT fk_game_id;
