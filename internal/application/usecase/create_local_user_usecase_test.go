package usecase_test

import (
	"context"
	"github.com/antunesgabriel/gopher-template-default/internal/application/dto"
	"github.com/antunesgabriel/gopher-template-default/internal/application/usecase"
	"github.com/antunesgabriel/gopher-template-default/internal/domain"
	"github.com/antunesgabriel/gopher-template-default/internal/domain/entity"
	"github.com/antunesgabriel/gopher-template-default/test/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateLocalUserUseCase(t *testing.T) {
	t.Run("it should return error if user is invalid", func(t *testing.T) {
		expect := assert.New(t)

		invalidUser := dto.CreateUserLocalInput{
			Name:     "Invalid User",
			Email:    "its_not_a_email",
			Password: "password",
		}

		mockRepository := mocks.NewMockUserRepository(t)

		mockPasswordHelper := mocks.NewMockPasswordHelper(t)

		uc := usecase.NewCreateLocalUserUseCase(mockRepository, mockPasswordHelper)

		err := uc.Execute(&invalidUser)

		expect.ErrorIs(err, domain.InvalidEmailError)
	})

	t.Run("it should return error UserAlreadyExistError if user already exist", func(t *testing.T) {
		expect := assert.New(t)

		mockRepository := mocks.NewMockUserRepository(t)
		mockPasswordHelper := mocks.NewMockPasswordHelper(t)

		email := "mouse@email.com"

		userInput := dto.CreateUserLocalInput{
			Name:     "Another Mouse",
			Email:    email,
			Password: "password",
		}

		ctx := context.Background()

		returns := entity.NewUser(0, "Mouse", email, "", "secret")

		mockRepository.EXPECT().FindUserByEmail(ctx, email).Return(returns, nil)

		uc := usecase.NewCreateLocalUserUseCase(mockRepository, mockPasswordHelper)

		err := uc.Execute(&userInput)

		expect.ErrorIs(err, domain.UserAlreadyExistError)
	})

	t.Run("it should calls UserRepository to create new user if he's is valid and hash with his password encrypted", func(t *testing.T) {
		expect := assert.New(t)

		userInput := dto.CreateUserLocalInput{
			Name:     "Archimedes",
			Email:    "archimedes@gopher.io",
			Password: "eureka",
		}

		mockRepository := mocks.NewMockUserRepository(t)
		mockPasswordHelper := mocks.NewMockPasswordHelper(t)

		passwordEncrypted := "xpto"

		ctx := context.Background()

		userExpected := entity.NewUser(0, userInput.Name, userInput.Email, "", passwordEncrypted)

		mockPasswordHelper.EXPECT().Hash(userInput.Password).Return([]byte(passwordEncrypted), nil)
		mockRepository.EXPECT().FindUserByEmail(ctx, userInput.Email).Return(nil, nil)

		mockRepository.EXPECT().Create(ctx, userExpected).Return(nil)

		uc := usecase.NewCreateLocalUserUseCase(mockRepository, mockPasswordHelper)

		err := uc.Execute(&userInput)

		expect.Nil(err)
	})

}
