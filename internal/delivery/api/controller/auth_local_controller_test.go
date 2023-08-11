package controller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/antunesgabriel/gopher-template-default/internal/application/dto"
	"github.com/antunesgabriel/gopher-template-default/internal/application/usecase"
	"github.com/antunesgabriel/gopher-template-default/internal/delivery/api/controller"
	"github.com/antunesgabriel/gopher-template-default/internal/domain"
	"github.com/antunesgabriel/gopher-template-default/internal/domain/entity"
	"github.com/antunesgabriel/gopher-template-default/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthLocalController_Handle(t *testing.T) {
	t.Run("it should returns correct payload if auth is failure", func(t *testing.T) {
		expect := assert.New(t)

		mockPasswordHelper := mocks.NewMockPasswordHelper(t)
		mockJWTHelper := mocks.NewMockJWTHelper(t)
		mockUserRepository := mocks.NewMockUserRepository(t)

		uc := usecase.NewLocalAuthUseCase(mockUserRepository, mockJWTHelper, mockPasswordHelper)

		ctrl := controller.NewAuthLocalController(uc)

		user := entity.NewUser(
			1,
			"Nana",
			"nana@email.com",
			"",
			"hashed_password",
		)

		authInputBody := dto.LocalAuthInput{
			Email:    user.Email,
			Password: "wrong_password",
		}

		bodyBytes, err := json.Marshal(authInputBody)

		expect.Nil(err)

		rr := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodPost, "/auth/local",
			bytes.NewReader(bodyBytes),
		)

		expect.Nil(err)

		ctx := context.Background()

		mockUserRepository.EXPECT().FindUserByEmail(
			ctx,
			authInputBody.Email,
		).Return(user, nil)

		mockPasswordHelper.EXPECT().Compare(
			authInputBody.Password,
			user.Password,
		).Return(errors.New("invalid password"))

		ctrl.Handle(rr, req)

		response := struct {
			Error string `json:"error,omitempty"`
		}{}

		expect.Equal(rr.Code, http.StatusBadRequest)

		err = json.NewDecoder(rr.Body).Decode(&response)

		expect.Nil(err)

		expect.Equal(response.Error, domain.InvalidPasswordError.Error())
	})

	t.Run("it should returns token in response payload when auth is successfully", func(t *testing.T) {
		expect := assert.New(t)

		mockPasswordHelper := mocks.NewMockPasswordHelper(t)
		mockJWTHelper := mocks.NewMockJWTHelper(t)
		mockUserRepository := mocks.NewMockUserRepository(t)

		uc := usecase.NewLocalAuthUseCase(mockUserRepository, mockJWTHelper, mockPasswordHelper)
		ctrl := controller.NewAuthLocalController(uc)

		user := entity.NewUser(
			1,
			"Nana",
			"nana@email.com",
			"",
			"hashed_password",
		)

		authInputBody := dto.LocalAuthInput{
			Email:    user.Email,
			Password: "correct_password",
		}

		bodyBytes, err := json.Marshal(authInputBody)

		expect.Nil(err)

		rr := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodPost, "/auth/local",
			bytes.NewReader(bodyBytes),
		)

		expect.Nil(err)

		ctx := context.Background()

		token := "example_token"

		mockUserRepository.EXPECT().FindUserByEmail(
			ctx,
			authInputBody.Email,
		).Return(user, nil)

		mockPasswordHelper.EXPECT().Compare(authInputBody.Password, user.Password).Return(nil)

		mockJWTHelper.EXPECT().Encode(mock.Anything).Return(token, nil)

		ctrl.Handle(rr, req)

		response := struct {
			Error string              `json:"error"`
			Data  dto.LocalAuthOutput `json:"data"`
		}{}

		err = json.NewDecoder(rr.Body).Decode(&response)

		expect.Nil(err)
		expect.Equal(response.Data.Token, token)
		expect.Empty(response.Error)
	})
}
