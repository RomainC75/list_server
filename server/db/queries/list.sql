-- name: Getlist :one
SELECT * FROM lists WHERE id = $1;

-- name: Listlists :many
SELECT * FROM lists ORDER BY name;