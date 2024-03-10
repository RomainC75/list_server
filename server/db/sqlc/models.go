// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql"
)

type Item struct {
	ID          int32
	Name        string
	Description sql.NullString
	Date        sql.NullTime
	UserID      sql.NullInt32
}

type List struct {
	ID     int32
	Name   string
	UserID sql.NullInt32
}

type User struct {
	ID       int32
	Email    string
	Password string
}
