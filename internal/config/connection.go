package config

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func NewDB() (*sql.DB, error) {
	url := os.Getenv("DATABASE_URL")

	var err error

	db, err := sql.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	return db, nil
}
