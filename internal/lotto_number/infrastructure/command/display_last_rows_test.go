package command_test

import (
	"os"
	"testing"

	"github.com/imbpp123/lotto_motto/internal/lotto_number/application/command"
	infra_command "github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/command"
)


func TestCreateCommandFromArgs_Success(t *testing.T) {
	os.Args[0] = "4"
	os.Args[1] = "http://example.com/test.zip"

	cmd, err := infra_command.CreateCommandFromArgs()
	
	if (err != nil) {
		t.Error("Error should be nil")
	}
	if cmd.RowCount != 4 {
		t.Error("Wrong value")
	}
	if cmd.Filename != "http://example.com/test.zip" {
		t.Error("Wrong value")
	}
}

func TestValidate_Fail(t *testing.T) {
	cmd := command.DisplayLastRowCommand{
		RowCount: 0,
		Filename: "test",
	}

	err := infra_command.ValidateInput(&cmd)

	if err == nil {
		t.Error("Error can not be nil")
	}
}

func TestValidate_Success(t *testing.T) {
	cmd := command.DisplayLastRowCommand{
		RowCount: 1,
		Filename: "http://example.com/test.zip",
	}

	err := infra_command.ValidateInput(&cmd)

	if err != nil {
		t.Error("Error can not be nil")
	}
}