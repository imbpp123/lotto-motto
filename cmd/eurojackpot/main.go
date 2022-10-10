package main

import (
	"fmt"
	"os"

	"github.com/imbpp123/lotto_motto/internal/lotto_number/application/command"
	infra_command "github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/command"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/model/repository"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/presentation"
)

func main() {
	fmt.Println("len", len(os.Args))

	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
	}

	// get variables from console
	cmd, err := infra_command.CreateCommandFromArgs()
	if err != nil {
		panic(err)
	}
	err = infra_command.ValidateInput(cmd)
	if err != nil {
		panic(err)
	}

	// run command
	fileRepository := repository.NumberFileRepository{}
	consoleTablePresentation := presentation.NumberConsoleTablePresentation{}
	handler := command.DisplayLastRowHandler{
		NumberRowCollectionRepository: fileRepository,
		NumberPresentation:            consoleTablePresentation,
	}
	err = handler.Handle(cmd)	
	if err != nil {
		panic(err)
	}
}

