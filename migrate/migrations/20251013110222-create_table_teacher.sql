-- +migrate Up
-- write your UP migration SQL here
CREATE TABLE IF NOT EXISTS teachers (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    class VARCHAR(50),
    subject VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    subject_id INT REFERENCES subjects(id)
    
);
CREATE INDEX IF NOT EXISTS idx_teachers_email ON teachers (email);

-- +migrate Up
-- write your UP migration SQL here

DROP INDEX IF EXISTS idx_teachers_email;
DROP TABLE IF EXISTS teachers;