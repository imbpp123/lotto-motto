package main

import (
	"github.com/imbpp123/lotto_motto/internal/lotto_number/application/command"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/data_provider"
	presentation2 "github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/presentation"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	dataProvider := data_provider.NewLotto6Aus49(2)
	rows, err := dataProvider.GetData()
	if err != nil {
		panic(err)
	}

	handler := command.RandomWeightCommandHandler{}

	cmd := command.RandomWeightCommand{
		Parts:               4,
		Min:                 1,
		Max:                 49,
		GenerateNumberCount: 6,
		Data:                rows,
		ResultViewer:        &presentation2.RawPrint{},
	}
	err = handler.Handle(cmd)
	if err != nil {
		panic(err)
	}
}
