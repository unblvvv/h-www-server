-- +goose Up
ALTER TABLE animals RENAME COLUMN photo_url TO photo_urls;
ALTER TABLE animals ALTER COLUMN photo_urls TYPE TEXT[] USING CASE WHEN photo_urls IS NOT NULL THEN ARRAY[photo_urls] ELSE NULL END;
-- +goose Down
ALTER TABLE animals ALTER COLUMN photo_urls TYPE VARCHAR(500) USING array_to_string(photo_urls, ',');
ALTER TABLE animals RENAME COLUMN photo_urls TO photo_url;
