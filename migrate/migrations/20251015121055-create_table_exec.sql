-- +migrate Up
-- write your UP migration SQL here

CREATE TABLE IF NOT EXISTS execs (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    password_changed_at TIMESTAMP NULL,
    password_reset_token VARCHAR(255) NULL,
    password_reset_token_expire TIMESTAMP NULL,
    role VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Indexes for faster lookups
CREATE INDEX idx_execs_email ON execs (email);
CREATE INDEX idx_execs_username ON execs (username);



-- +migrate Down
-- write your DOWN migration SQL here

-- Drop indexes first (important: avoid dependency errors)
DROP INDEX IF EXISTS idx_execs_email;
DROP INDEX IF EXISTS idx_execs_username;

-- Then drop the table
DROP TABLE IF EXISTS execs;