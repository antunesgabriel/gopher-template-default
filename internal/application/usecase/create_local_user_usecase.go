package usecase

import (
	"context"

	"github.com/antunesgabriel/gopher-template-default/internal/application/repository"
	"github.com/antunesgabriel/gopher-template-default/internal/domain/entity"
)

type CreateLocalUserUseCase struct {
	repository repository.UserRepository
}

func NewCreateLocalUserUseCase(repository repository.UserRepository) *CreateLocalUserUseCase {
	uc := CreateLocalUserUseCase{
		repository: repository,
	}

	return &uc
}

func (it *CreateLocalUserUseCase) Execute(name, email, password string) error {
	ctx := context.Background()

	u := entity.NewUser(0, name, email, "", password)

	err := u.ValidateNewLocalUser()

	if err != nil {
		return err
	}

	err = it.repository.Create(ctx, u)

	return err
}
