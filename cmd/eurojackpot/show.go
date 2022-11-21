package eurojackpot

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/application/command"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/presentation"
	"github.com/spf13/cobra"
	"os"
)

func runShowCommand(cmd *cobra.Command, args []string) {
	handlerCmd := command.DisplayLastRowCommand{
		RowCount:     10,
		Filename:     "",
		ColumnNumber: nil,
	}

	// validate
	validate := validator.New()
	errs := validate.Struct(cmd)
	if errs != nil {
		fmt.Fprintln(os.Stderr, errs)
		os.Exit(1)
	}

	// create handler
	consoleTablePresentation := presentation.NumberConsoleTablePresentation{}

	handler := command.DisplayLastRowHandler{
		DataLoader:         nil,
		NumberPresentation: consoleTablePresentation,
	}

	// run command
	if err := handler.Handle(handlerCmd); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
