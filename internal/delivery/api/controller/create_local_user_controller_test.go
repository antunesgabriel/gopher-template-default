package controller

import (
	"bytes"
	"encoding/json"
	"github.com/antunesgabriel/gopher-template-default/internal/application/dto"
	"github.com/antunesgabriel/gopher-template-default/internal/application/usecase"
	"github.com/antunesgabriel/gopher-template-default/internal/domain"
	"github.com/antunesgabriel/gopher-template-default/test/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateLocalUserController_Handle(t *testing.T) {
	t.Run("it should correctly answer when the username is invalid", func(t *testing.T) {
		input := dto.CreateUserLocalInput{
			Name:     "Antunes",
			Email:    "this is not a valid email",
			Password: "xpto",
		}

		invalidUser, err := json.Marshal(input)

		if err != nil {
			t.Errorf("got %s want %s", err, "no error expected")
		}

		req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewReader(invalidUser))

		if err != nil {
			t.Errorf("got %s want %s", err, "no error expected")
		}

		rr := httptest.NewRecorder()

		controller := buildController()

		controller.Handle(rr, req)

		response := dto.Response{}

		err = json.NewDecoder(rr.Body).Decode(&response)

		if err != nil {
			t.Errorf("got %s want %s", err, "no error expected")
		}

		if rr.Code != http.StatusBadRequest {
			t.Errorf("got %d want %d", rr.Code, http.StatusBadRequest)
		}

		if response.Error != domain.InvalidEmailError.Error() {
			t.Errorf("got %s want %s", response.Error, domain.InvalidEmailError)
		}
	})

	t.Run("it should answer with correctly status code when user was validated and stored with successfully", func(t *testing.T) {
		input := dto.CreateUserLocalInput{
			Name:     "Antunes",
			Email:    "json@email.io",
			Password: "xpto",
		}

		validUser, err := json.Marshal(input)

		if err != nil {
			t.Errorf("got %s want %s", err, "no error expected")
		}

		req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewReader(validUser))

		if err != nil {
			t.Errorf("got %s want %s", err, "no error expected")
		}

		rr := httptest.NewRecorder()

		controller := buildController()

		controller.Handle(rr, req)

		response := dto.Response{}

		err = json.NewDecoder(rr.Body).Decode(&response)

		if err != nil {
			t.Errorf("got %s want %s", err, "no error expected")
		}

		if rr.Code != http.StatusCreated {
			t.Errorf("got %d want %d", rr.Code, http.StatusCreated)
		}

		if response.Error != "" {
			t.Errorf("got %s want %s", response.Error, "empty error field")
		}
	})
}

func buildController() *CreateLocalUserController {
	r := mock.NewMockUserRepository()
	h := mock.NewMockPasswordHelper("#4asd3", true)
	u := usecase.NewCreateLocalUserUseCase(r, h)

	controller := NewCreateLocalUserController(u)

	return controller
}
