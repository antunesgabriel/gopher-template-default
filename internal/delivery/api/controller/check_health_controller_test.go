package controller_test

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/antunesgabriel/gopher-template-default/internal/application/usecase"
	"github.com/antunesgabriel/gopher-template-default/internal/delivery/api/controller"
	"github.com/antunesgabriel/gopher-template-default/test/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckHealthController_Handle(t *testing.T) {
	t.Run("it should response up true if db is up and false if db is down", func(t *testing.T) {
		expect := assert.New(t)

		mockHealthRepository := mocks.NewMockHealthRepository(t)
		uc := usecase.NewCheckHealthUseCase(mockHealthRepository)

		req, err := http.NewRequest(http.MethodGet, "/health", nil)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		ctrl := controller.NewCheckHealthController(uc)

		response := struct {
			Error string `json:"error"`
			Data  struct {
				Up bool `json:"up"`
			} `json:"data"`
		}{}

		ctx := context.Background()

		mockHealthRepository.EXPECT().Ping(ctx).Return(nil)

		ctrl.Handle(rr, req)

		err = json.NewDecoder(rr.Body).Decode(&response)

		expect.Nil(err)

		expect.Equal(rr.Code, http.StatusOK)
		expect.Empty(response.Error)
		expect.True(response.Data.Up)
	})

	t.Run("it should response up false if db is down", func(t *testing.T) {
		expect := assert.New(t)

		mockHealthRepository := mocks.NewMockHealthRepository(t)
		uc := usecase.NewCheckHealthUseCase(mockHealthRepository)

		req, err := http.NewRequest(http.MethodGet, "/health", nil)

		expect.Nil(err)

		rr := httptest.NewRecorder()

		ctrl := controller.NewCheckHealthController(uc)

		response := struct {
			Error string `json:"error"`
			Data  struct {
				Up bool `json:"up"`
			} `json:"data"`
		}{}

		ctx := context.Background()
		mockHealthRepository.EXPECT().Ping(ctx).Return(errors.New("some error"))

		ctrl.Handle(rr, req)

		err = json.NewDecoder(rr.Body).Decode(&response)

		expect.Nil(err)

		expect.Equal(rr.Code, http.StatusInternalServerError)
		expect.Equal(response.Error, "some error")
		expect.False(response.Data.Up)
	})
}
