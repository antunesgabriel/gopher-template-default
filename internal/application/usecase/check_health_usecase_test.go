package usecase

import (
	"context"
	"errors"
	"testing"
)

type mockHealthRepository struct {
	Return error
}

func (it *mockHealthRepository) Ping(ctx context.Context) error {
	return it.Return
}

func TestCheckHealthUseCase(t *testing.T) {
	t.Run("it should called HealthRepository correctly", func(t *testing.T) {
		mock := mockHealthRepository{
			nil,
		}

		usecase := NewCheckHealthUseCase(&mock)

		if err := usecase.Execute(); err != nil {
			t.Errorf("got %s want %s", err, "nil")
		}

		errTst := errors.New("vasco")

		mock.Return = errTst

		if err := usecase.Execute(); err == nil || !errors.Is(err, errTst) {
			t.Errorf("got %s want %s", err, errTst)
		}
	})
}
