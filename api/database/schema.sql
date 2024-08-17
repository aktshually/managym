CREATE TABLE IF NOT EXISTS users (
    id text PRIMARY KEY,
    email VARCHAR(120) UNIQUE NOT NULL,
    "password" VARCHAR(120) NOT NULL,
    first_name VARCHAR(120) NOT NULL,
    last_name VARCHAR(120) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    is_active boolean DEFAULT true
);
