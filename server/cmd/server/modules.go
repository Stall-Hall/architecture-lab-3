//go:build wireinject
// +build wireinject

package main

import (
	"architecture-lab-3/server/virtual-machines"

	"github.com/google/wire"
)

func ComposeApiServer(port HttpPortNumber) (*APIServer, error) {
	wire.Build(
		NewDbConnection,
		virtualmachines.Providers,
		wire.Struct(new(APIServer), "Port", "vmHandler"),
	)
	return nil, nil
}
