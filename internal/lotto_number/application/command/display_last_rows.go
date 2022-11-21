package command

import (
	"github.com/imbpp123/lotto_motto/internal/lotto_number/application/presentation"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/model"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/model/rows_loader"
)

type DisplayLastRow interface {
	CreateCommandFromArgs() DisplayLastRowCommand
	Validate(cmd *DisplayLastRowCommand)
}

type DisplayLastRowCommand struct {
	RowCount     int    `validate:"required,min=1"`
	Filename     string `validate:"required"`
	ColumnNumber []int
}

type DisplayLastRowHandler struct {
	DataLoader         rows_loader.RowDataLoader
	NumberPresentation presentation.NumberRowCollectionPresentation
}

func (h *DisplayLastRowHandler) Handle(cmd DisplayLastRowCommand) error {
	data, err := h.DataLoader.Load(cmd.Filename)

	collection, err := model.CreateFromStringMap(data, cmd.ColumnNumber)
	if err != nil {
		return err
	}

	collection.SortByDate().Slice(0, cmd.RowCount)

	h.NumberPresentation.Display(collection)

	return nil
}
