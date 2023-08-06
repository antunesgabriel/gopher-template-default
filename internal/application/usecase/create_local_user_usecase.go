package usecase

import (
	"context"
	"github.com/antunesgabriel/gopher-template-default/internal/helper"

	"github.com/antunesgabriel/gopher-template-default/internal/application/repository"
	"github.com/antunesgabriel/gopher-template-default/internal/domain/entity"
)

type CreateLocalUserUseCase struct {
	repository     repository.UserRepository
	passwordHelper helper.PasswordHelper
}

func NewCreateLocalUserUseCase(repository repository.UserRepository, passwordHelper helper.PasswordHelper) *CreateLocalUserUseCase {
	uc := CreateLocalUserUseCase{
		repository:     repository,
		passwordHelper: passwordHelper,
	}

	return &uc
}

func (it *CreateLocalUserUseCase) Execute(name, email, password string) error {
	ctx := context.Background()

	// TODO: check if user already exist

	u := entity.NewUser(0, name, email, "", password)

	err := u.ValidateNewLocalUser()

	if err != nil {
		return err
	}

	passBytes, err := it.passwordHelper.Hash(u.Password)

	if err != nil {
		return err
	}

	u.ChangePassword(string(passBytes))

	err = it.repository.Create(ctx, u)

	return err
}
