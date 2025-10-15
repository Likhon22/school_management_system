-- +migrate Up
-- write your UP migration SQL here
CREATE TABLE IF NOT EXISTS students (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    class VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    subject_id INT REFERENCES subjects(id)
  
);
CREATE INDEX IF NOT EXISTS idx_students_email ON students (email);

-- +migrate Down
-- write your DOWN migration SQL here

DROP INDEX IF EXISTS idx_students_email;
DROP TABLE IF EXISTS students;