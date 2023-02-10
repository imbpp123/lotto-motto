package command

import (
	"fmt"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/presentation"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/model"
)

type ResultViewerInterface interface {
	Display(data *model.ResultArray)
}

type RandomWeightCommand struct {
	Parts               int
	Min                 int
	Max                 int
	GenerateNumberCount int
	Data                [][]int
	ResultViewer        ResultViewerInterface
}

type RandomWeightCommandHandler struct {
}

func (h *RandomWeightCommandHandler) Handle(cmd RandomWeightCommand) error {
	collection := model.NewNumPeriodCollectionByParts(cmd.Parts, cmd.Min, cmd.Max)
	collection.SetData(cmd.Data)

	if err := collection.CalculateWeight(); err != nil {
		return fmt.Errorf("can't calculate weight: %w", err)
	}

	resultNormal, err := collection.GenerateRandomValues(cmd.GenerateNumberCount)
	if err != nil {
		return fmt.Errorf("can't generate random values: %w", err)
	}
	resultNormal.Sort()

	present := presentation.RawPrint{}
	present.Display(resultNormal)

	return nil
}
