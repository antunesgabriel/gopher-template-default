package usecase

import (
	"context"
	"github.com/antunesgabriel/gopher-template-default/internal/application/dto"
	"github.com/antunesgabriel/gopher-template-default/internal/application/repository"
	"github.com/antunesgabriel/gopher-template-default/internal/domain"
	"github.com/antunesgabriel/gopher-template-default/internal/helper"
	"time"
)

type LocalAuthUseCase struct {
	repository     repository.UserRepository
	jwtHelper      helper.JWTHelper
	passwordHelper helper.PasswordHelper
}

func NewLocalAuthUseCase(userRepository repository.UserRepository, jwtHelper helper.JWTHelper, passwordHelper helper.PasswordHelper) *LocalAuthUseCase {
	uc := LocalAuthUseCase{
		repository:     userRepository,
		jwtHelper:      jwtHelper,
		passwordHelper: passwordHelper,
	}

	return &uc
}

func (it *LocalAuthUseCase) Execute(input dto.LocalAuthInput) (string, error) {
	ctx := context.Background()

	user, err := it.repository.FindUserByEmail(ctx, input.Email)

	if err != nil {
		return "", err
	}

	if user == nil {
		return "", domain.UserIsNotRegister
	}

	if user.Password == "" {
		return "", domain.UserIsNotLocal
	}

	err = it.passwordHelper.Compare(input.Password, user.Password)

	if err != nil {
		return "", domain.InvalidPassword
	}

	payload := map[string]interface{}{
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		"iat": time.Now().Unix(),
		"id":  user.ID,
	}

	token, err := it.jwtHelper.Encode(payload)

	return token, err
}
