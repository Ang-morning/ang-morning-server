package reviews

import (
	"angmorning.com/internal/services/reviews/application"
	"angmorning.com/internal/services/reviews/infrastructure"
	"angmorning.com/internal/services/reviews/presentation"
	"github.com/google/wire"
)

var ReviewSet = wire.NewSet(
	presentation.New,
	application.New,
	infrastructure.New,
)
