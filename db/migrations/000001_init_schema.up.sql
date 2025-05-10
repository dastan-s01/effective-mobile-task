CREATE TABLE IF NOT EXISTS people (
    id UUID PRIMARY KEY,
    full_name TEXT NOT NULL,
    age INT,
    gender TEXT,
    nationality TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
    );