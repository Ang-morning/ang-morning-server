package users

import (
	"angmorning.com/internal/libs/oauth"
	"angmorning.com/internal/services/users/application"
	"angmorning.com/internal/services/users/infrastructure"
	"angmorning.com/internal/services/users/presentation"
	"github.com/google/wire"
)

var UserSet = wire.NewSet(
	oauth.NewFactory,
	infrastructure.New,
	application.New,
	presentation.New,
)
