package main

import (
	"flag"

	"github.com/go-playground/validator"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/application/command"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/model/repository"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/presentation"
)

func main() {
	cmd := command.DisplayLastRowCommand{
		RowCount: *flag.Uint("rows", 10, "rows to show in table"),
		Filename: *flag.String("file", "http://example.com", "zip archive with history in CSV format"),
	}

	validate := validator.New()
	errs := validate.Struct(cmd)
	if errs != nil {
		panic(errs)
	}

	// run command
	fileRepository := repository.NumberFileRepository{}
	consoleTablePresentation := presentation.NumberConsoleTablePresentation{}
	handler := command.DisplayLastRowHandler{
		NumberRowCollectionRepository: fileRepository,
		NumberPresentation:            consoleTablePresentation,
	}
	err := handler.Handle(cmd)	
	if err != nil {
		panic(err)
	}
}

