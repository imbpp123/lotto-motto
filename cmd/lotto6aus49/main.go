package main

import (
	"context"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/application/command"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/application/query"
	presentation2 "github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/presentation"
	repository2 "github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/repository"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ctx := context.Background()

	repository := repository2.NewGameRepository(repository2.GameRepositoryConfig{
		Url:      "https://www.lotto-berlin.de/static/gamebroker_7/default/download_files/archiv_lotto.zip",
		Filename: "lotto_6aus49_ab_02.12.2000.txt",
		Columns: repository2.GameColumns{
			Year:  2,
			Month: 1,
			Day:   0,
			Levels: []repository2.ColumnLevel{
				{
					Start: 3,
					End:   9,
					Level: 0,
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
		Parts:               4,
		Min:                 1,
		Max:                 49,
		GenerateNumberCount: 6,
		Data:                game.GetRowsByLevel(0),
		ResultViewer:        &presentation2.RawPrint{},
	}
	err = handler.Handle(cmd)
	if err != nil {
		panic(err)
	}
}
