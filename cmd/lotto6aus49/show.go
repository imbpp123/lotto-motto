package lotto6aus49

import (
	"github.com/go-playground/validator"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/application/command"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/presentation"
	"github.com/spf13/cobra"
)

func runShowCommand(cmd *cobra.Command, args []string) error {
	handlerCmd := command.DisplayLastRowCommand{
		RowCount:     10,
		Filename:     "",
		ColumnNumber: nil,
	}

	// validate
	validate := validator.New()
	errs := validate.Struct(cmd)
	if errs != nil {
		return errs
	}

	// create handler
	consoleTablePresentation := presentation.NumberConsoleTablePresentation{}

	handler := command.DisplayLastRowHandler{
		DataLoader:         nil,
		NumberPresentation: consoleTablePresentation,
	}

	// run command
	return handler.Handle(handlerCmd)
}
