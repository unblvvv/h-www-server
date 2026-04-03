-- +goose Up
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS animals (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    age VARCHAR(50) NOT NULL,
    sex VARCHAR(20) NOT NULL,
    description TEXT NOT NULL,
    photo_url VARCHAR(500),
    status VARCHAR(50) DEFAULT 'available',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_animals_organization_id ON animals(organization_id);
CREATE INDEX idx_animals_status ON animals(status);

-- +goose Down
SELECT 'down SQL query';
