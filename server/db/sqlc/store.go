package db

import (
	"database/sql"
	"fmt"
	"log"

	// TODO: important for sql.Open() !!
	_ "github.com/lib/pq"

	"github.com/RomainC75/todo2/config"
)

type Store interface {
	Querier
}

var DbStore *Store

type SqlStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SqlStore{
		db:      db,
		Queries: New(db),
	}
}

func Connect() {
	cfg := config.Get()

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)

	conn, err := sql.Open("postgres", dsn)
	store := NewStore(conn)

	if err != nil {
		log.Fatal("error trying to connect to the database : ", err)
	}

	DbStore = &store
}

func GetConnection() *Store {
	return DbStore
}
