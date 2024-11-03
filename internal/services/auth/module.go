package auth

import (
	"angmorning.com/internal/services/auth/application"
	"angmorning.com/internal/services/auth/infrastructure"
	"github.com/google/wire"
)

var AuthSet = wire.NewSet(
	application.New,
	infrastructure.New,
)
