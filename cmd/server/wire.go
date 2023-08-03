//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/antunesgabriel/gopher-template-default/internal/adapter"
	"github.com/antunesgabriel/gopher-template-default/internal/adapter/repository"
	"github.com/antunesgabriel/gopher-template-default/internal/app"
	"github.com/antunesgabriel/gopher-template-default/internal/app/module/health"
	"github.com/antunesgabriel/gopher-template-default/internal/app/module/user"
	"github.com/google/wire"
)

var RepositorySet = wire.NewSet(
	repository.NewPostgresRespository,
	repository.NewPostgresUserRespository,
	repository.NewPostgresHealthRepository,
)

var ServiceSet = wire.NewSet(
	RepositorySet,
	wire.Bind(new(user.UserRepository), new(*repository.PostgresUserRepository)),
	user.NewUserService,
	wire.Bind(new(health.HealthRepository), new(*repository.PostgresHealthRepository)),
	health.NewHealthService,
)

var ControllerSet = wire.NewSet(
	user.NewUserController,
	health.NewHealthController,
)

var ServerSet = wire.NewSet(
	adapter.NewChiRouter,
	wire.Bind(new(app.Router), new(*adapter.ChiRouter)),
	app.NewServer,
)

func InitServer(db *sql.DB) *app.Server {
	wire.Build(
		ServerSet,
		ServiceSet,
		ControllerSet,
	)

	return &app.Server{}
}
