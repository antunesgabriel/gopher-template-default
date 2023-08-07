package usecase

import (
	"errors"
	"github.com/antunesgabriel/gopher-template-default/internal/application/dto"
	"github.com/antunesgabriel/gopher-template-default/internal/domain"
	"github.com/antunesgabriel/gopher-template-default/internal/domain/entity"
	"github.com/antunesgabriel/gopher-template-default/test/mock"
	"testing"
)

func TestLocalAuthUseCase_Execute(t *testing.T) {
	t.Run("it should return error UserIsNotRegister user no exist", func(t *testing.T) {
		user := entity.NewUser(1, "Jhon", "jhon@doe.com", "", "secret")

		users := []*entity.User{
			user,
		}

		uc := buildLocalAuthUseCase(users, "", false, 1, "")

		input := dto.LocalAuthInput{
			Email:    "noexist@email.com",
			Password: "secret",
		}

		_, err := uc.Execute(input)

		if !errors.Is(err, domain.UserIsNotRegister) {
			t.Errorf("got %s want %s", err, domain.UserIsNotRegister)
		}
	})

	t.Run("it should return error UserIsNotLocal when user is not local", func(t *testing.T) {
		externalUser := entity.NewUser(1, "Jhon", "jhon@doe.com", "google", "")

		users := []*entity.User{
			externalUser,
		}

		uc := buildLocalAuthUseCase(users, "", false, 1, "")

		input := dto.LocalAuthInput{
			Email:    "jhon@doe.com",
			Password: "secret",
		}

		_, err := uc.Execute(input)

		if !errors.Is(err, domain.UserIsNotLocal) {
			t.Errorf("got %s want %s", err, domain.UserIsNotLocal)
		}
	})

	t.Run("it should return error InvalidPassword when password is invalid", func(t *testing.T) {
		password := "correct"

		externalUser := entity.NewUser(1, "Jhon", "jhon@doe.com", "", password)

		users := []*entity.User{
			externalUser,
		}

		uc := buildLocalAuthUseCase(users, "", false, 1, "")

		input := dto.LocalAuthInput{
			Email:    "jhon@doe.com",
			Password: "wrog",
		}

		_, err := uc.Execute(input)

		if !errors.Is(err, domain.InvalidPassword) {
			t.Errorf("got %s want %s", err, domain.InvalidPassword)
		}

		input = dto.LocalAuthInput{
			Email:    "jhon@doe.com",
			Password: password,
		}

		_, err = uc.Execute(input)

		if errors.Is(err, domain.InvalidPassword) {
			t.Errorf("got %s want %s", "no error", domain.InvalidPassword)
		}
	})

	t.Run("it should return token if user exist, he is local and password is correct", func(t *testing.T) {
		email := "nana@cat.com"
		password := "meuw"

		user := entity.NewUser(1, "Nana", email, "", password)

		users := []*entity.User{
			user,
		}

		token := "exampletoken"

		uc := buildLocalAuthUseCase(users, "", true, 1, token)

		input := dto.LocalAuthInput{
			Email:    email,
			Password: password,
		}

		tkn, err := uc.Execute(input)

		if err != nil {
			t.Errorf("got %s want %s", err, "no error expected")
		}

		if tkn != token {
			t.Errorf("got %s want %s", tkn, token)
		}
	})
}

func buildLocalAuthUseCase(users []*entity.User, fakeHash string, isEqual bool, id int, token string) *LocalAuthUseCase {

	mockRepository := mock.NewMockUserRepository()
	passwordHelper := mock.NewMockPasswordHelper(fakeHash, isEqual)
	jwtHelper := mock.NewMockJWTHelper(id, token)

	mockRepository.Users = users

	uc := LocalAuthUseCase{
		repository:     mockRepository,
		passwordHelper: passwordHelper,
		jwtHelper:      jwtHelper,
	}

	return &uc
}
