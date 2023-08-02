package repository

import "context"

type PostgresHealthRepository struct {
	*PostgresRepository
}

func NewPostgresHealthRepository(pr *PostgresRepository) *PostgresHealthRepository {
	r := PostgresHealthRepository{
		PostgresRepository: pr,
	}

	return &r
}

func (r *PostgresHealthRepository) Ping(ctx context.Context) error {
	return r.db.Ping()
}
