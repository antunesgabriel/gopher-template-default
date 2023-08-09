package usecase_test

import (
	"context"
	"errors"
	"github.com/antunesgabriel/gopher-template-default/internal/application/dto"
	"github.com/antunesgabriel/gopher-template-default/internal/application/usecase"
	"github.com/antunesgabriel/gopher-template-default/internal/domain"
	"github.com/antunesgabriel/gopher-template-default/internal/domain/entity"
	"github.com/antunesgabriel/gopher-template-default/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestLocalAuthUseCase_Execute(t *testing.T) {
	t.Run("it should return error UserIsNotRegisterError user no exist", func(t *testing.T) {
		expect := assert.New(t)

		mockUserRepository := mocks.NewMockUserRepository(t)
		mockPasswordHelper := mocks.NewMockPasswordHelper(t)
		mockJWTHelper := mocks.NewMockJWTHelper(t)

		uc := usecase.NewLocalAuthUseCase(mockUserRepository, mockJWTHelper, mockPasswordHelper)

		input := dto.LocalAuthInput{
			Email:    "jhon@doe.com",
			Password: "secret",
		}

		mockUserRepository.EXPECT().FindUserByEmail(context.Background(), mock.Anything).Return(nil, nil).Times(1)

		token, err := uc.Execute(&input)

		expect.ErrorIs(err, domain.UserIsNotRegisterError)

		expect.Empty(token)
	})

	t.Run("it should return error UserIsNotLocalError when user is not local", func(t *testing.T) {
		expect := assert.New(t)

		externalUser := entity.NewUser(1, "Jhon", "jhon@doe.com", "google", "")

		input := dto.LocalAuthInput{
			Email:    externalUser.Email,
			Password: "secret",
		}

		mockUserRepository := mocks.NewMockUserRepository(t)
		mockPasswordHelper := mocks.NewMockPasswordHelper(t)
		mockJWTHelper := mocks.NewMockJWTHelper(t)

		uc := usecase.NewLocalAuthUseCase(mockUserRepository, mockJWTHelper, mockPasswordHelper)

		mockUserRepository.EXPECT().FindUserByEmail(context.Background(), input.Email).Return(externalUser, nil)

		token, err := uc.Execute(&input)

		expect.ErrorIs(err, domain.UserIsNotLocalError)
		expect.Empty(token)
	})

	t.Run("it should return error InvalidPasswordError when PasswordHelper.Compare returns error", func(t *testing.T) {
		expect := assert.New(t)

		userStored := entity.NewUser(1, "Jhon", "jhon@doe.com", "", "hashed_password")

		input := dto.LocalAuthInput{
			Email:    userStored.Email,
			Password: "wrong",
		}

		mockUserRepository := mocks.NewMockUserRepository(t)
		mockPasswordHelper := mocks.NewMockPasswordHelper(t)
		mockJWTHelper := mocks.NewMockJWTHelper(t)

		uc := usecase.NewLocalAuthUseCase(mockUserRepository, mockJWTHelper, mockPasswordHelper)

		mockUserRepository.EXPECT().FindUserByEmail(context.Background(), input.Email).Return(userStored, nil)
		mockPasswordHelper.EXPECT().Compare(input.Password, userStored.Password).Return(errors.New("it is different")).Times(1)

		token, err := uc.Execute(&input)

		expect.ErrorIs(err, domain.InvalidPasswordError)
		expect.Empty(token)
	})

	t.Run("it should calls JWTToken Encode if auth input is valid and returns auth token", func(t *testing.T) {
		expect := assert.New(t)

		userStored := entity.NewUser(1, "Jhon", "jhon@doe.com", "", "hashed_password")

		input := dto.LocalAuthInput{
			Email:    userStored.Email,
			Password: "secret",
		}

		mockUserRepository := mocks.NewMockUserRepository(t)
		mockPasswordHelper := mocks.NewMockPasswordHelper(t)
		mockJWTHelper := mocks.NewMockJWTHelper(t)

		uc := usecase.NewLocalAuthUseCase(mockUserRepository, mockJWTHelper, mockPasswordHelper)

		mockUserRepository.EXPECT().FindUserByEmail(context.Background(), input.Email).Return(userStored, nil)
		mockPasswordHelper.EXPECT().Compare(input.Password, userStored.Password).Return(nil)

		tokenExpected := "fake_token"

		expectedClaimsMatcher := mock.MatchedBy(func(claims map[string]interface{}) bool {
			exp, expOk := claims["exp"].(int64)
			iat, iatOk := claims["iat"].(int64)
			id, idOk := claims["id"].(int64)
			return expOk && iatOk && idOk && exp > 0 && iat > 0 && id == userStored.ID
		})

		mockJWTHelper.EXPECT().Encode(expectedClaimsMatcher).Return(tokenExpected, nil).Times(1)

		token, err := uc.Execute(&input)

		expect.Nil(err)
		expect.Equal(token, tokenExpected)
	})
}
