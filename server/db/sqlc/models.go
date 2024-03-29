// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql"
	"time"
)

type Item struct {
	ID          int32
	Name        string
	Description sql.NullString
	Date        sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type List struct {
	ID        int32
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int32
}

type ListItem struct {
	ID     int32
	ListID int32
	ItemID int32
}

type User struct {
	ID        int32
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
