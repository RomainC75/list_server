-- name: Getlist :one
SELECT * FROM lists WHERE id = $1;

-- name: Listlists :many
SELECT * FROM lists ORDER BY name;

-- name: CreateList :one
INSERT INTO lists (
    name,
    user_id,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4
) RETURNING *;