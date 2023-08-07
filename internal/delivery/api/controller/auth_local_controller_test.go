package controller

import (
	"bytes"
	"encoding/json"
	"github.com/antunesgabriel/gopher-template-default/internal/application/dto"
	"github.com/antunesgabriel/gopher-template-default/internal/application/usecase"
	"github.com/antunesgabriel/gopher-template-default/internal/domain"
	"github.com/antunesgabriel/gopher-template-default/internal/domain/entity"
	"github.com/antunesgabriel/gopher-template-default/test/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthLocalController_Handle(t *testing.T) {
	t.Run("it should return correct payload if authentication is failure", func(t *testing.T) {
		user := entity.NewUser(
			1,
			"Nana",
			"nana@email.com",
			"",
			"secret",
		)

		users := []*entity.User{
			user,
		}

		uc := buildLocalAuthUseCase(users, "xtpo", int(user.ID), "faketoken")

		emptyBody, err := json.Marshal(dto.LocalAuthInput{
			Email:    user.Email,
			Password: "wrongpassword",
		})

		if err != nil {
			t.Errorf("got %s want %s", err, "no error expected")
		}

		controller := NewAuthLocalController(uc)

		req, err := http.NewRequest(
			http.MethodPost, "/auth/local",
			bytes.NewReader(emptyBody),
		)

		if err != nil {
			t.Errorf("got %s want %s", err, "no error expected")
		}

		rr := httptest.NewRecorder()

		controller.Handle(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("got %d want %d", rr.Code, http.StatusBadRequest)
		}

		response := struct {
			Error string `json:"error,omitempty"`
		}{}

		err = json.NewDecoder(rr.Body).Decode(&response)

		if err != nil {
			t.Errorf("got %s want %s", err, "no error expected")
		}

		if response.Error != domain.InvalidPasswordError.Error() {
			t.Errorf("got %s want %s", response.Error, domain.InvalidFieldsError.Error())
		}
	})

	t.Run("it should return correct payload if authentication is successfully", func(t *testing.T) {
		user := entity.NewUser(
			1,
			"Nana Two",
			"nanatwo@email.com",
			"",
			"secrettwo",
		)

		users := []*entity.User{
			user,
		}

		token := "testtoken"

		uc := buildLocalAuthUseCase(users, "xtpo", int(user.ID), token)

		emptyBody, err := json.Marshal(dto.LocalAuthInput{
			Email:    user.Email,
			Password: user.Password,
		})

		if err != nil {
			t.Errorf("got %s want %s", err, "no error expected")
		}

		controller := NewAuthLocalController(uc)

		req, err := http.NewRequest(
			http.MethodPost, "/auth/local",
			bytes.NewReader(emptyBody),
		)

		if err != nil {
			t.Errorf("got %s want %s", err, "no error expected")
		}

		rr := httptest.NewRecorder()

		controller.Handle(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("got %d want %d", rr.Code, http.StatusOK)
		}

		response := struct {
			Error string              `json:"error,omitempty"`
			Data  dto.LocalAuthOutput `json:"data,omitempty"`
		}{}

		err = json.NewDecoder(rr.Body).Decode(&response)

		if err != nil {
			t.Errorf("got %s want %s", err, "no error expected")
		}

		if response.Error != "" {
			t.Errorf("got %s want %s", response.Error, "no error expected")
		}

		if response.Data.Token != token {
			t.Errorf("got %s want %s", response.Data.Token, token)
		}
	})
}

func buildLocalAuthUseCase(users []*entity.User, fakeHash string, id int, token string) *usecase.LocalAuthUseCase {

	mockRepository := mock.NewMockUserRepository()
	passwordHelper := mock.NewMockPasswordHelper(fakeHash, false)
	jwtHelper := mock.NewMockJWTHelper(id, token)

	mockRepository.Users = users

	uc := usecase.NewLocalAuthUseCase(mockRepository, jwtHelper, passwordHelper)

	return uc
}
