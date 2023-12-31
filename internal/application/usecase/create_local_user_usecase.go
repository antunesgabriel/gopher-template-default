package usecase

import (
	"context"
	"github.com/antunesgabriel/gopher-template-default/internal/application/dto"
	"github.com/antunesgabriel/gopher-template-default/internal/domain"
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

func (it *CreateLocalUserUseCase) Execute(input *dto.CreateUserLocalInput) error {
	ctx := context.Background()

	u := entity.NewUser(0, input.Name, input.Email, "", input.Password)

	err := u.ValidateNewLocalUser()

	if err != nil {
		return err
	}

	exist, err := it.repository.FindUserByEmail(ctx, u.Email)

	if exist != nil {
		return domain.UserAlreadyExistError
	}

	passBytes, err := it.passwordHelper.Hash(u.Password)

	if err != nil {
		return err
	}

	u.ChangePassword(string(passBytes))

	err = it.repository.Create(ctx, u)

	return err
}
