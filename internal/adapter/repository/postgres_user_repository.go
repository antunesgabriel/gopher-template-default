package repository

import (
	"context"
	"database/sql"

	"gihub.com/antunesgabriel/gopher-template-default/internal/adapter/database"
	"gihub.com/antunesgabriel/gopher-template-default/internal/app/module/user"
	"gihub.com/antunesgabriel/gopher-template-default/internal/helper"
)

type PostgresUserRepository struct {
	query *database.Queries
}

func NewPostgresUserRespository(db *sql.DB) *PostgresUserRepository {
	q := database.New(db)

	p := PostgresUserRepository{
		query: q,
	}

	return &p
}

func (p PostgresUserRepository) FindUserByEmail(
	ctx context.Context,
	email string,
) (*user.User, error) {
	u, err := p.query.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	usr := user.New(
		u.ID,
		u.Name,
		u.Email,
		helper.GetStringFromNullString(&u.Provider),
		helper.GetStringFromNullString(&u.Password),
	)

	return usr, nil
}

func (p PostgresUserRepository) FindUserByID(ctx context.Context, id int64) (*user.User, error) {
	u, err := p.query.GetUserByID(ctx, id)

	if err != nil {
		return nil, err
	}

	usr := user.New(
		u.ID,
		u.Name,
		u.Email,
		helper.GetStringFromNullString(&u.Provider),
		helper.GetStringFromNullString(&u.Password),
	)

	return usr, nil
}

func (p PostgresUserRepository) Create(ctx context.Context, u *user.User) error {
	if u.Password != "" {
		_, err := p.query.CreateUserLocal(ctx, database.CreateUserLocalParams{
			Name:  u.Name,
			Email: u.Email,
			Password: sql.NullString{
				String: u.Password,
				Valid:  true,
			},
		})

		return err
	}

	return nil
}

func (p PostgresUserRepository) Update(ctx context.Context, u *user.User) (*user.User, error) {
	return nil, nil
}

func (p PostgresUserRepository) Delete(ctx context.Context, u *user.User) error {
	return nil
}
