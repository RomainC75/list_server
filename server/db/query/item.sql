-- name: Getitem :one
SELECT * FROM items WHERE id = $1;

-- name: ListItems :many
SELECT * FROM items ORDER BY name;

-- name: CreateItem :one
INSERT INTO items (
    name,
    description,
    date,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: LinkItemToList :one
INSERT INTO list_item (
    list_id,
    item_id
)VALUES (
    $1, $2
) RETURNING *;