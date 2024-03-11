// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: item.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createItem = `-- name: CreateItem :one
INSERT INTO items (
    name,
    description,
    date,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id, name, description, date, created_at, updated_at
`

type CreateItemParams struct {
	Name        string
	Description sql.NullString
	Date        sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (q *Queries) CreateItem(ctx context.Context, arg CreateItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, createItem,
		arg.Name,
		arg.Description,
		arg.Date,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Date,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getItemsByListName = `-- name: GetItemsByListName :many
SELECT items.id, items.name, items.description, items.date, items.created_at, items.updated_at FROM items 
INNER JOIN list_item ON items.id=list_item.item_id 
INNER JOIN lists ON list_item.list_id = lists.id
WHERE lists.id = $1
`

func (q *Queries) GetItemsByListName(ctx context.Context, id int32) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, getItemsByListName, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Item{}
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Date,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getitem = `-- name: Getitem :one
SELECT id, name, description, date, created_at, updated_at FROM items WHERE id = $1
`

func (q *Queries) Getitem(ctx context.Context, id int32) (Item, error) {
	row := q.db.QueryRowContext(ctx, getitem, id)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Date,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const linkItemToList = `-- name: LinkItemToList :one
INSERT INTO list_item (
    list_id,
    item_id
)VALUES (
    $1, $2
) RETURNING id, list_id, item_id
`

type LinkItemToListParams struct {
	ListID int32
	ItemID int32
}

func (q *Queries) LinkItemToList(ctx context.Context, arg LinkItemToListParams) (ListItem, error) {
	row := q.db.QueryRowContext(ctx, linkItemToList, arg.ListID, arg.ItemID)
	var i ListItem
	err := row.Scan(&i.ID, &i.ListID, &i.ItemID)
	return i, err
}
