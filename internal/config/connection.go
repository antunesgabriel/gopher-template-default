package config

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func NewDB(env *Env) (*sql.DB, error) {
	var err error

	db, err := sql.Open("postgres", env.DatabaseURL)

	if err != nil {
		return nil, err
	}

	return db, nil
}
