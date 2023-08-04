package usecase

import (
	"context"

	"github.com/antunesgabriel/gopher-template-default/internal/application/repository"
)

type CheckHealthUseCase struct {
	repository repository.HealthRepository
}

func NewCheckHealthUseCase(repository repository.HealthRepository) *CheckHealthUseCase {
	uc := CheckHealthUseCase{
		repository: repository,
	}

	return &uc
}

func (it *CheckHealthUseCase) Execute() error {
	ctx := context.Background()

	if err := it.repository.Ping(ctx); err != nil {
		return err
	}

	return nil
}
