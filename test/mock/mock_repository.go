package mock

import (
	"context"
	"github.com/antunesgabriel/gopher-template-default/internal/domain/entity"
)

type MockUserRepository struct {
	Users []*entity.User
}

func (it *MockUserRepository) FindUserByID(ctx context.Context, id int64) (*entity.User, error) {
	return nil, nil
}

func (it *MockUserRepository) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	return nil, nil
}

func (it *MockUserRepository) Create(ctx context.Context, u *entity.User) error {
	it.Users = append(it.Users, u)
	return nil
}

func (it *MockUserRepository) Update(ctx context.Context, u *entity.User) (*entity.User, error) {
	return nil, nil
}

func (it *MockUserRepository) Delete(ctx context.Context, u *entity.User) error {
	return nil
}

type MockHealthRepository struct {
	Return error
}

func (it *MockHealthRepository) Ping(ctx context.Context) error {
	return it.Return
}
