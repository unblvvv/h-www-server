-- +goose Up
SELECT 'up SQL query';

ALTER TABLE applications ADD COLUMN user_id UUID NOT NULL;

-- +goose Down
SELECT 'down SQL query';
