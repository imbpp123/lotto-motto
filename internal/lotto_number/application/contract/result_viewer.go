package contract

import (
	"context"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/domain/aggregate"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/domain/value_object"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/model"
)

type ResultViewerInterface interface {
	Display(data *model.ResultArray)
}

type GameRepository interface {
	FetchData(ctx context.Context, rows int) ([]value_object.Raffle, error)
}

type GameConsoleViewer interface {
	Display(game *aggregate.Game)
}
