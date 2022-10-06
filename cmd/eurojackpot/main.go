package eurojackpot

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/application/command"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/model/repository"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/presentation"
)

var (
	validate *validator.Validate
	cmd command.DisplayLastRowCommand
)

func createCommandFromArgs() {
	rowCount, err := strconv.ParseUint(os.Args[1], 10, 0)
	if err != nil {
		panic(err)
	}
	filename := os.Args[3]

	cmd = command.DisplayLastRowCommand{
		RowCount: uint(rowCount),
		Filename: filename,
	}
}

func validateInput(cmd command.DisplayLastRowCommand) {
	isFailed := false

	validate := validator.New()
	errs := validate.Var(cmd.RowCount, "required,min=1")
	if errs != nil {
		fmt.Println(errs) 
		isFailed = true
	}

	errs = validate.Var(cmd.Filename, "required,url")
	if errs != nil {
		fmt.Println(errs) 
		isFailed = true
	} 

	if isFailed {
		panic("Validation failed")
	}
}

func main() {
	// get variables from console
	createCommandFromArgs()
	validateInput(cmd)

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

