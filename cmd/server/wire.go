//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"gihub.com/antunesgabriel/gopher-template-default/internal/adapter"
	"gihub.com/antunesgabriel/gopher-template-default/internal/adapter/repository"
	"gihub.com/antunesgabriel/gopher-template-default/internal/app"
	"gihub.com/antunesgabriel/gopher-template-default/internal/app/module/user"
	"github.com/google/wire"
)

var RepositorySet = wire.NewSet(
	repository.NewPostgresRespository,
	repository.NewPostgresUserRespository,
)

var ServiceSet = wire.NewSet(
	wire.Bind(new(user.UserRepository, new(*repository.PostgresUserRepository))),
	user.NewUserService,
)

var ControllerSet = wire.NewSet(
	user.NewUserController,
)

var ServerSet = wire.NewSet(
	wire.Bind(new(app.Router), new(*adapter.ChiRouter)),
	app.NewServer,
)

func InitServer(db *sql.DB) (*app.Server, error) {
	err := wire.Build(
		RepositorySet,
		ServiceSet,
		ControllerSet,
		ServerSet,
	)

	return &app.NewServer{}, err
}
