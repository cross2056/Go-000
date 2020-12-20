// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"week04/internal/biz"
	"week04/internal/data"
	"week04/internal/pkg/database"
	"week04/internal/service"

	"github.com/google/wire"
)

var dataRepoSet = wire.NewSet(data.NewMessageRepo, wire.Bind(new(biz.MessageRepo), new(*data.MessageRepo)))

func InitializeHelloServer(dsnURI string) *service.HelloServer {
	wire.Build(service.NewHelloServer, biz.NewMessageBiz, dataRepoSet, database.InitDB)
	return nil
}
