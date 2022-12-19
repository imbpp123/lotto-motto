package presentation

import (
	"fmt"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/model"
)

type RawPrint struct {
}

func (p *RawPrint) Display(data *model.ResultArray) {
	fmt.Println("Rows:")
	for idx, num := range data.Result() {
		fmt.Printf("%d = %d\n", idx+1, num)
	}
}
