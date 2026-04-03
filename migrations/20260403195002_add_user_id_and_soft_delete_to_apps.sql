-- +goose Up
SELECT 'up SQL query';

ALTER TABLE animals ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP WITH TIME ZONE;

ALTER TABLE applications ADD COLUMN IF NOT EXISTS user_id UUID NOT NULL;

ALTER TABLE applications DROP CONSTRAINT IF EXISTS fk_animal;
ALTER TABLE applications
    ADD CONSTRAINT fk_animal
        FOREIGN KEY (animal_id)
            REFERENCES animals(id)
            ON DELETE SET NULL;

CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_user_animal_app ON applications (user_id, animal_id);

-- +goose Down
SELECT 'down SQL query';
