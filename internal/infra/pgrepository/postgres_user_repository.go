package pgrepository

import (
	"context"
	"database/sql"

	"github.com/antunesgabriel/gopher-template-default/internal/domain/entity"

	"github.com/antunesgabriel/gopher-template-default/internal/helper"
	"github.com/antunesgabriel/gopher-template-default/internal/infra/database"
)

type PostgresUserRepository struct {
	*PostgresRepository
}

func NewPostgresUserRepository(pr *PostgresRepository) *PostgresUserRepository {

	p := PostgresUserRepository{
		PostgresRepository: pr,
	}

	return &p
}

func (p *PostgresUserRepository) FindUserByEmail(
	ctx context.Context,
	email string,
) (*entity.User, error) {
	u, err := p.query.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	usr := entity.New(
		u.ID,
		u.Name,
		u.Email,
		helper.GetStringFromNullString(&u.Provider),
		helper.GetStringFromNullString(&u.Password),
	)

	return usr, nil
}

func (p *PostgresUserRepository) FindUserByID(ctx context.Context, id int64) (*entity.User, error) {
	u, err := p.query.GetUserByID(ctx, id)

	if err != nil {
		return nil, err
	}

	usr := entity.New(
		u.ID,
		u.Name,
		u.Email,
		helper.GetStringFromNullString(&u.Provider),
		helper.GetStringFromNullString(&u.Password),
	)

	return usr, nil
}

func (p *PostgresUserRepository) Create(ctx context.Context, u *entity.User) error {
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

func (p *PostgresUserRepository) Update(ctx context.Context, u *entity.User) (*entity.User, error) {
	return nil, nil
}

func (p *PostgresUserRepository) Delete(ctx context.Context, u *entity.User) error {
	return nil
}
