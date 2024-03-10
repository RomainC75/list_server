-- name: Getlist :one
SELECT * FROM lists WHERE id = $1;

-- name: Getlists :many
SELECT * FROM lists WHERE user_id = $1 ORDER BY updated_at;

-- name: CreateList :one
INSERT INTO lists (
    name,
    user_id,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4
) RETURNING *;