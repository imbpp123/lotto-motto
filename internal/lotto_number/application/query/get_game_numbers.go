package query

import (
	"context"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/application/contract"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/domain/aggregate"
)

type GetGameNumbersCommand struct {
	Rows int
}

type GetGameNumbersQuery struct {
	repository contract.GameRepository
}

func NewGetGameNumbersQuery(repository contract.GameRepository) *GetGameNumbersQuery {
	return &GetGameNumbersQuery{repository: repository}
}

func (q *GetGameNumbersQuery) FetchData(ctx context.Context, command GetGameNumbersCommand) (*aggregate.Game, error) {
	rows, err := q.repository.FetchData(ctx, command.Rows)
	if err != nil {
		return nil, err
	}
	return aggregate.NewGameWithRows(rows), nil
}
