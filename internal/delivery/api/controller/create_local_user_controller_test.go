package controller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/antunesgabriel/gopher-template-default/internal/application/dto"
	"github.com/antunesgabriel/gopher-template-default/internal/application/usecase"
	"github.com/antunesgabriel/gopher-template-default/internal/delivery/api/controller"
	"github.com/antunesgabriel/gopher-template-default/internal/domain"
	"github.com/antunesgabriel/gopher-template-default/internal/domain/entity"
	"github.com/antunesgabriel/gopher-template-default/test/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateLocalUserController_Handle(t *testing.T) {
	t.Run("it should return the correct response payload if you get an error when trying to register a user", func(t *testing.T) {
		expect := assert.New(t)

		mockUserRepository := mocks.NewMockUserRepository(t)
		mockPasswordHelper := mocks.NewMockPasswordHelper(t)
		uc := usecase.NewCreateLocalUserUseCase(mockUserRepository, mockPasswordHelper)

		ctrl := controller.NewCreateLocalUserController(uc)

		fakeUserStored := entity.NewUser(1, "Jhon", "jhon@email.com", "", "hashed")

		createNewUserInput := dto.CreateUserLocalInput{
			Name:     "Another Jhon",
			Email:    fakeUserStored.Email,
			Password: "xpto",
		}

		bodyBytes, err := json.Marshal(createNewUserInput)

		expect.Nil(err)

		req, err := http.NewRequest(http.MethodPost, "/signup/local", bytes.NewReader(bodyBytes))

		expect.Nil(err)

		rr := httptest.NewRecorder()

		ctx := context.Background()

		mockUserRepository.EXPECT().FindUserByEmail(ctx, createNewUserInput.Email).Return(fakeUserStored, nil)

		ctrl.Handle(rr, req)

		response := dto.Response{}

		err = json.NewDecoder(rr.Body).Decode(&response)

		expect.Nil(err)
		expect.Equal(response.Error, domain.UserAlreadyExistError.Error())
		expect.Equal(rr.Code, http.StatusBadRequest)
	})

	t.Run("it should create user and returns correct code if new user input is valid", func(t *testing.T) {
		expect := assert.New(t)

		mockUserRepository := mocks.NewMockUserRepository(t)
		mockPasswordHelper := mocks.NewMockPasswordHelper(t)
		uc := usecase.NewCreateLocalUserUseCase(mockUserRepository, mockPasswordHelper)

		ctrl := controller.NewCreateLocalUserController(uc)

		createNewUserInput := dto.CreateUserLocalInput{
			Name:     "Nana Cat",
			Email:    "nana@catspace.io",
			Password: "secret",
		}

		bodyBytes, err := json.Marshal(createNewUserInput)

		expect.Nil(err)

		req, err := http.NewRequest(http.MethodPost, "/signup/local", bytes.NewReader(bodyBytes))

		expect.Nil(err)

		rr := httptest.NewRecorder()

		ctx := context.Background()

		mockUserRepository.EXPECT().FindUserByEmail(
			ctx,
			createNewUserInput.Email,
		).Return(nil, nil)

		hashedPassword := []byte("encrypted password")

		mockPasswordHelper.EXPECT().Hash(createNewUserInput.Password).Return(
			hashedPassword,
			nil,
		)

		newUserExpectedParams := entity.NewUser(
			0,
			createNewUserInput.Name,
			createNewUserInput.Email,
			"",
			string(hashedPassword),
		)

		mockUserRepository.EXPECT().Create(ctx, newUserExpectedParams).Return(nil)

		ctrl.Handle(rr, req)

		response := dto.Response{}

		err = json.NewDecoder(rr.Body).Decode(&response)

		expect.Nil(err)
		expect.Equal(rr.Code, http.StatusCreated)
		expect.Empty(response.Error)
	})
}
