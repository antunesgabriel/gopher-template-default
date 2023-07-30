package config

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func NewDB() error {
	url := os.Getenv("POSTGRES_DATABASE_CONNECTION")

	var err error

	DB, err = sql.Open("postgres", url)

	if err != nil {
		return err
	}

	return nil
}
