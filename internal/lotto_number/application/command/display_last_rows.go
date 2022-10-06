package command

import "github.com/imbpp123/lotto_motto/internal/lotto_number/application/loader"

type DisplayLastRowCommand struct {
	RowCount uint16
	Filename string
}

type DisplayLastRowHandler struct {
	NumberLoader       loader.NumberLoader
	NumberPresentation presentation.NumberPresentation
}

// displays table of numbers from lotto
func (h *DisplayLastRowHandler) Handle(cmd DisplayLastRowCommand) error {
	collection, err := h.NumberLoader.load(cmd.Filename)
	if err != nil {
		return err
	}

	collection = collection.SortByDate().Slice(0, cmd.RowCount)

	h.NumberPresentation.Display(collection)
}
