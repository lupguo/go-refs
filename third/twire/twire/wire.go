//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package twire

import (
	"github.com/google/wire"
)

func InitApp(dsn string) (*App, error) {
	wire.Build(
		NewBussInfra,
		NewApp,
		NewService,
		wire.Bind(new(IRepos), new(*BussInfra)),
	)
	return &App{}, nil
}
