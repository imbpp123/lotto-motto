package aggregate

import (
	"github.com/imbpp123/lotto_motto/internal/lotto_number/domain/value_object"
)

type Game struct {
	raffle []value_object.Raffle
}

func NewGameWithRows(rows []value_object.Raffle) *Game {
	return &Game{
		raffle: rows,
	}
}

func (g *Game) GetRowsByLevel(level int) [][]int {
	var data [][]int

	for _, val := range g.raffle {
		data = append(data, val.NumberByLevel(level))
	}

	return data
}

func (g *Game) RaffleCount() int {
	return len(g.raffle)
}

func (g *Game) RawRaffleData(idx int) []interface{} {
	var data []interface{}

	currentRaffle := g.raffle[idx]

	data = append(data, currentRaffle.Date().Format("2006-02-01"))

	for i := 0; i < currentRaffle.LevelCount(); i++ {
		for _, number := range currentRaffle.NumberByLevel(i) {
			data = append(data, number)
		}
	}

	return data
}
