package presentation

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/model"
)

type NumberConsoleTablePresentation struct {

}

func (p NumberConsoleTablePresentation) Display(collection *model.NumberRowCollection) {
	t := table.NewWriter()
    t.SetOutputMirror(os.Stdout)

    t.AppendHeader(table.Row{"Date", "Normal", "Extra"})
    
    t.AppendRows([]table.Row{
        {"Arya", "Stark", 3000},
        {"Jon", "Snow", 2000},
    })
    
    t.AppendSeparator()
    t.Render()
}