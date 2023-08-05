package usecase

import (
	"errors"
	m "github.com/antunesgabriel/gopher-template-default/test/mock"
	"testing"
)

func TestCheckHealthUseCase(t *testing.T) {
	t.Run("it should called HealthRepository correctly", func(t *testing.T) {
		mock := m.MockHealthRepository{
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
