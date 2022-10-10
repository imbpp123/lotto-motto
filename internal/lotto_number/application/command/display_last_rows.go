package command

import (
	"github.com/imbpp123/lotto_motto/internal/lotto_number/application/presentation"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/model/repository"
)

type DisplayLastRow interface {
	CreateCommandFromArgs() DisplayLastRowCommand
	Validate(cmd *DisplayLastRowCommand)
}

type DisplayLastRowCommand struct {
	RowCount uint `validate:"required,min=1"`
	Filename string `validate:"required"`
}

type DisplayLastRowHandler struct {
	NumberRowCollectionRepository repository.NumberRowCollectionRepository
	NumberPresentation presentation.NumberRowCollectionPresentation
}

// displays table of numbers from lotto
func (h *DisplayLastRowHandler) Handle(cmd DisplayLastRowCommand) error {
	collection, err := h.NumberRowCollectionRepository.LoadFromFile(cmd.Filename)
	if err != nil {
		return err
	}

	collection = collection.SortByDate().Slice(0, int(cmd.RowCount))

	h.NumberPresentation.Display(collection)

	return nil
}
