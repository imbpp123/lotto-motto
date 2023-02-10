package presentation

import (
	"fmt"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/domain/aggregate"

	"github.com/jedib0t/go-pretty/v6/table"
)

type RaffleTable struct {
}

func (p *RaffleTable) Display(game *aggregate.Game) {
	t := table.NewWriter()
	t.SetCaption("Game raffle numbers.\n")

	for i := 0; i < game.RaffleCount(); i++ {
		t.AppendRow(game.RawRaffleData(i))
	}

	fmt.Println(t.Render())
}
