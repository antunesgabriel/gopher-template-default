package controller

import (
	"encoding/json"
	"errors"
	"github.com/antunesgabriel/gopher-template-default/internal/application/usecase"
	m "github.com/antunesgabriel/gopher-template-default/test/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckHealthController_Handle(t *testing.T) {
	t.Run("it should response up true if db is up", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/health", nil)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		usecaseDBUp := newMockUseCase(nil)
		controller := NewCheckHealthController(usecaseDBUp)

		controller.Handle(rr, req)

		response := controllerResponse{}

		err = json.NewDecoder(rr.Body).Decode(&response)

		if err != nil {
			t.Errorf("got %s want %s", err, "no error")
		}

		if rr.Code != http.StatusOK {
			t.Errorf("got %d want %d", rr.Code, http.StatusOK)
		}

		if response.Error != "" {
			t.Errorf("got %s want %s", response.Error, "error should to be empty")
		}

		if response.Data.Up != true {
			t.Errorf("got %t want %t", response.Data.Up, true)
		}
	})

	t.Run("it should response up false if db is down", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/health", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		response := controllerResponse{}

		usecaseDBDown := newMockUseCase(errors.New("db_down"))
		controller := NewCheckHealthController(usecaseDBDown)

		controller.Handle(rr, req)

		err = json.NewDecoder(rr.Body).Decode(&response)

		if err != nil {
			t.Errorf("got %s want %s", err, "no error")
		}

		if rr.Code != http.StatusInternalServerError {
			t.Errorf("got %d want %d", rr.Code, http.StatusInternalServerError)
		}

		if response.Error == "" {
			t.Errorf("got %s want %s", response.Error, "db_down")
		}

		if response.Data.Up != false {
			t.Errorf("got %t want %t", response.Data.Up, false)
		}
	})
}

type upField struct {
	Up bool `json:"up"`
}

type controllerResponse struct {
	Error string  `json:"error,omitempty"`
	Data  upField `json:"data,omitempty"`
}

func newMockUseCase(err error) *usecase.CheckHealthUseCase {
	r := m.MockHealthRepository{
		Return: err,
	}

	u := usecase.NewCheckHealthUseCase(&r)

	return u
}
