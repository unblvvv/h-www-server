-- +goose Up
SELECT 'up SQL query';

ALTER TABLE users ADD COLUMN role VARCHAR(50) DEFAULT 'user';

-- +goose Down
SELECT 'down SQL query';

ALTER TABLE users DROP COLUMN role;
