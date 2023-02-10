package main

import (
	"context"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/application/command"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/application/query"
	presentation2 "github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/presentation"
	repository2 "github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/repository"
	"hash/maphash"
	"math/rand"
)

func main() {
	rand.Seed(int64(new(maphash.Hash).Sum64()))
	ctx := context.Background()

	repository := repository2.NewGameRepository(repository2.GameRepositoryConfig{
		Url:      "https://www.lotto-berlin.de/static/gamebroker_7/default/download_files/archiv_eurojackpot.zip",
		Filename: "eurojackpot.txt",
		Columns: repository2.GameColumns{
			Year:  2,
			Month: 1,
			Day:   0,
			Levels: []repository2.ColumnLevel{
				{
					Start: 3,
					End:   8,
					Level: 0,
				},
				{
					Start: 8,
					End:   10,
					Level: 1,
				},
			},
		},
	})

	getGameQuery := query.NewGetGameNumbersQuery(repository)
	game, err := getGameQuery.FetchData(ctx, query.GetGameNumbersCommand{
		Rows: 5,
	})
	if err != nil {
		panic(err)
	}

	raffleRender := presentation2.RaffleTable{}
	raffleRender.Display(game)

	handler := command.RandomWeightCommandHandler{}

	cmd := command.RandomWeightCommand{
		Parts:               5,
		Min:                 1,
		Max:                 50,
		GenerateNumberCount: 5,
		Data:                game.GetRowsByLevel(0),
		ResultViewer:        &presentation2.RawPrint{},
	}
	err = handler.Handle(cmd)
	if err != nil {
		panic(err)
	}

	cmd = command.RandomWeightCommand{
		Parts:               3,
		Min:                 1,
		Max:                 12,
		GenerateNumberCount: 2,
		Data:                game.GetRowsByLevel(1),
		ResultViewer:        &presentation2.RawPrint{},
	}
	err = handler.Handle(cmd)
	if err != nil {
		panic(err)
	}
}
