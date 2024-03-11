-- name: Getitem :one
SELECT * FROM items WHERE id = $1;

-- name: GetItemsByListName :many
SELECT items.* FROM items 
INNER JOIN list_item ON items.id=list_item.item_id 
INNER JOIN lists ON list_item.list_id = lists.id
WHERE lists.id = $1;

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
