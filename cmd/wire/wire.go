//go:build wireinject
// +build wireinject

package wire

import (
	"angmorning.com/internal/server"
	"github.com/google/wire"
)

func InitializeServer() (*server.Server, error) {
	wire.Build(server.ProviderSet)
	return &server.Server{}, nil
}
