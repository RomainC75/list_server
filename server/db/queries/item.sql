-- name: Getitem :one
SELECT * FROM items WHERE id = $1;

-- name: ListItems :many
SELECT * FROM items ORDER BY name;