//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"github.com/antunesgabriel/gopher-template-default/internal/application/repository"
	"github.com/antunesgabriel/gopher-template-default/internal/application/usecase"
	"github.com/antunesgabriel/gopher-template-default/internal/delivery/api"
	"github.com/antunesgabriel/gopher-template-default/internal/delivery/api/controller"
	"github.com/antunesgabriel/gopher-template-default/internal/helper"
	"github.com/antunesgabriel/gopher-template-default/internal/infra"
	"github.com/antunesgabriel/gopher-template-default/internal/infra/pgrepository"
	"github.com/google/wire"
)

var RepositorySet = wire.NewSet(
	pgrepository.NewPostgresRepository,
	pgrepository.NewPostgresUserRepository,
	pgrepository.NewPostgresHealthRepository,
)

var UseCaseSet = wire.NewSet(
	HelperSet,
	wire.Bind(new(helper.PasswordHelper), new(*infra.BcryptPasswordHelper)),
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
	wire.Bind(new(api.Router), new(*infra.ChiRouter)),
	api.NewServer,
)

var HelperSet = wire.NewSet(
	infra.NewBcryptPasswordHelper,
)

func InitServer(db *sql.DB) *api.Server {
	wire.Build(
		ServerSet,
		UseCaseSet,
		ControllerSet,
	)

	return &api.Server{}
}
