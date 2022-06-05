// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"module4/internal/biz"
	"module4/internal/conf"
	"module4/internal/data"
	"module4/internal/server"
	"module4/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
