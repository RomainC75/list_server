-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: Listusers :many
SELECT * FROM users ORDER BY email;

-- name: CreateUser :one
INSERT INTO users (
    email,
    password
    -- created_at,
    -- updated_at
) VALUES (
    $1, $2
) RETURNING *;

