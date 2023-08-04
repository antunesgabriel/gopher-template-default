//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/antunesgabriel/gopher-template-default/internal/application/repository"
	"github.com/antunesgabriel/gopher-template-default/internal/application/usecase"
	"github.com/antunesgabriel/gopher-template-default/internal/infra"
	"github.com/antunesgabriel/gopher-template-default/internal/infra/pgrepository"
	"github.com/antunesgabriel/gopher-template-default/internal/presentation"
	"github.com/antunesgabriel/gopher-template-default/internal/presentation/controller"
	"github.com/google/wire"
)

var RepositorySet = wire.NewSet(
	pgrepository.NewPostgresRepository,
	pgrepository.NewPostgresUserRepository,
	pgrepository.NewPostgresHealthRepository,
)

var UseCaseSet = wire.NewSet(
	RepositorySet,
	wire.Bind(new(repository.UserRepository), new(*pgrepository.PostgresUserRepository)),
	usecase.NewCreateLocalUserUseCase,
	wire.Bind(new(repository.HealthRepository), new(*pgrepository.PostgresHealthRepository)),
	usecase.NewCheckHealthUseCase,
)

var ControllerSet = wire.NewSet(
	controller.NewCreateLocalUserController,
	controller.NewCheckHealthController,
)

var ServerSet = wire.NewSet(
	infra.NewChiRouter,
	wire.Bind(new(presentation.Router), new(*infra.ChiRouter)),
	presentation.NewServer,
)

func InitServer(db *sql.DB) *presentation.Server {
	wire.Build(
		ServerSet,
		UseCaseSet,
		ControllerSet,
	)

	return &presentation.Server{}
}
