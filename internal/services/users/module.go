package users

import (
	"angmorning.com/internal/services/users/application"
	"angmorning.com/internal/services/users/infrastructure"
	"angmorning.com/internal/services/users/presentation"
	"github.com/google/wire"
)

var UserModule = wire.NewSet(
	infrastructure.New,
	application.New,
	presentation.New,
)
