package main

import "github.com/imbpp123/lotto_motto/internal/lotto_number/application/command"

var (
	handler command.DisplayLastRowHandler
)

func main() {
	var cmd command.DisplayLastRowCommand
	cmd.RowCount = 4
	cmd.Filename = "https://www.lotto-berlin.de/static/gamebroker_7/default/download_files/archiv_eurojackpot.zip"}
	
	err := handler.Handle(cmd)	
	if err != nil {
		panic(err)
	}
}