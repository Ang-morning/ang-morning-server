//go:build wireinject
// +build wireinject

package di

import (
	"angmorning.com/internal/libs/db"
	"angmorning.com/internal/server"
	"angmorning.com/internal/services/auth"
	"angmorning.com/internal/services/users"
	"github.com/google/wire"
)

func InitializeServer() (*server.Server, error) {
	wire.Build(db.InitDb, users.UserSet, server.ProviderSet, auth.AuthSet)
	return &server.Server{}, nil
}
