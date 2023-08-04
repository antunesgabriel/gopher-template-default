package repository

import (
	"context"

	"github.com/antunesgabriel/gopher-template-default/internal/domain/entity"
)

type UserRepository interface {
	FindUserByID(ctx context.Context, id int64) (*entity.User, error)
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
	Create(ctx context.Context, u *entity.User) error
	Update(ctx context.Context, u *entity.User) (*entity.User, error)
	Delete(ctx context.Context, u *entity.User) error
}
