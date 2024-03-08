-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: Listusers :many
SELECT * FROM users ORDER BY email;