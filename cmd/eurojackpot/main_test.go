package main_test

import (
	"testing"

	"github.com/imbpp123/lotto_motto/internal/lotto_number/application/command"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/model/repository"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/presentation"
)

func TestMain_Success(t *testing.T) {
	cmd := command.DisplayLastRowCommand{
		RowCount: 5,
		Filename: "https://www.lotto-berlin.de/static/gamebroker_7/default/download_files/archiv_eurojackpot.zip",
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