package command

import (
	"errors"
	"os"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/application/command"
)

func CreateCommandFromArgs() (*command.DisplayLastRowCommand, error) {
	if (len(os.Args) < 2) {
		return nil, errors.New("there should be 2 parameters: rows and http link")
	}

	rowCount, err := strconv.ParseUint(os.Args[0], 10, 0)
	if err != nil {
		return nil, err
	}
	filename := os.Args[1]

	return &command.DisplayLastRowCommand{
		RowCount: uint(rowCount),
		Filename: filename,
	}, nil
}

func ValidateInput(cmd *command.DisplayLastRowCommand) (error) {
	validate := validator.New()
	return validate.Struct(cmd)
}