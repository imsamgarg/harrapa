-- +goose Up
CREATE TYPE USER_ROLE as ENUM('user', 'admin');

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (), role USER_ROLE NOT NULL, name VARCHAR(50) NOT NULL, profile_picture TEXT, email TEXT NOT NULL UNIQUE, password TEXT NOT NULL, created_at TIMESTAMP NOT NULL DEFAULT (now()), updated_at TIMESTAMP NOT NULL DEFAULT (now())
);

-- +goose Down
DROP TABLE users;

DROP TYPE USER_ROLE;