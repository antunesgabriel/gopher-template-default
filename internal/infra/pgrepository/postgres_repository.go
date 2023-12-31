package pgrepository

import (
	"database/sql"

	"github.com/antunesgabriel/gopher-template-default/internal/infra/database"
)

type PostgresRepository struct {
	db    *sql.DB
	query *database.Queries
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	q := database.New(db)

	p := PostgresRepository{
		db:    db,
		query: q,
	}

	return &p
}
