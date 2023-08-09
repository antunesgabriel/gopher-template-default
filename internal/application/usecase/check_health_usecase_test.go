package usecase_test

import (
	"context"
	"errors"
	"github.com/antunesgabriel/gopher-template-default/internal/application/usecase"
	"github.com/antunesgabriel/gopher-template-default/test/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckHealthUseCase(t *testing.T) {
	t.Run("it should called HealthRepository correctly", func(t *testing.T) {
		expect := assert.New(t)

		repoMocked := mocks.NewMockHealthRepository(t)

		uc := usecase.NewCheckHealthUseCase(repoMocked)

		ctx := context.Background()

		repoMocked.EXPECT().Ping(ctx).Return(nil).Times(1)

		err := uc.Execute()

		expect.Nil(err)
	})

	t.Run("it should return error if database have any problems", func(t *testing.T) {
		expect := assert.New(t)

		repoMocked := mocks.NewMockHealthRepository(t)

		uc := usecase.NewCheckHealthUseCase(repoMocked)

		ctx := context.Background()

		errExpected := errors.New("some error")

		repoMocked.EXPECT().Ping(ctx).Return(errExpected).Times(1)

		err := uc.Execute()

		expect.ErrorIs(err, errExpected)
	})
}
