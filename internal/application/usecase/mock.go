package usecase

import (
	"context"
	"github.com/antunesgabriel/gopher-template-default/internal/domain/entity"
)

type mockUserRepository struct {
	Users []*entity.User
}

func (it *mockUserRepository) FindUserByID(ctx context.Context, id int64) (*entity.User, error) {
	return nil, nil
}

func (it *mockUserRepository) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	return nil, nil
}

func (it *mockUserRepository) Create(ctx context.Context, u *entity.User) error {
	it.Users = append(it.Users, u)
	return nil
}

func (it *mockUserRepository) Update(ctx context.Context, u *entity.User) (*entity.User, error) {
	return nil, nil
}

func (it *mockUserRepository) Delete(ctx context.Context, u *entity.User) error {
	return nil
}

type mockHealthRepository struct {
	Return error
}

func (it *mockHealthRepository) Ping(ctx context.Context) error {
	return it.Return
}