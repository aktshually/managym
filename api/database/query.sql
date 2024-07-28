-- name: GetUser :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users(
    id,
    email,
    password,
    first_name,
    last_name
) VALUES ($1, $2, $3, $4, $5)
RETURNING id, email, first_name, last_name, created_at, updated_at, is_active;
