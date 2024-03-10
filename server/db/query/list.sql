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

-- name: GetListForUpdate :one
SELECT * FROM lists
WHERE id = $1 AND user_id = $2 LIMIT 1
FOR NO KEY UPDATE;
-- NO KEY : avoid dead-lock ! 

-- name: UpdateList :one
UPDATE lists 
SET name = $2, updated_at = $3
WHERE id = $1
RETURNING *;

-- name: DeleteList :one
DELETE FROM lists
WHERE id = $1 AND user_id = $2
RETURNING *;