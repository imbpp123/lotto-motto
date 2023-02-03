package main

import (
	"github.com/imbpp123/lotto_motto/internal/lotto_number/application/command"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/data_provider"
	presentation2 "github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/presentation"
	"hash/maphash"
	"math/rand"
)

func main() {
	rand.Seed(int64(new(maphash.Hash).Sum64()))

	dataProvider := data_provider.NewEuroJackpot(3)
	rows, err := dataProvider.GetData()
	if err != nil {
		panic(err)
	}

	normalRows := [][]int{}
	extraRows := [][]int{}
	for _, row := range rows {
		normalRows = append(normalRows, row[0:5])
		extraRows = append(extraRows, row[5:7])
	}

	handler := command.RandomWeightCommandHandler{}

	cmd := command.RandomWeightCommand{
		Parts:               5,
		Min:                 1,
		Max:                 50,
		GenerateNumberCount: 5,
		Data:                normalRows,
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
		Data:                extraRows,
		ResultViewer:        &presentation2.RawPrint{},
	}
	err = handler.Handle(cmd)
	if err != nil {
		panic(err)
	}
}
