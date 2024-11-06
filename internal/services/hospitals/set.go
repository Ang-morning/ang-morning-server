package hospitals

import (
	"angmorning.com/internal/services/hospitals/application"
	"angmorning.com/internal/services/hospitals/infrastructure"
	"angmorning.com/internal/services/hospitals/presentation"
	"github.com/google/wire"
)

var HospitalSet = wire.NewSet(
	presentation.New,
	application.New,
	infrastructure.New,
)
