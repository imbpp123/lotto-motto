package contract

import "github.com/imbpp123/lotto_motto/internal/lotto_number/model"

type ResultViewerInterface interface {
	Display(data *model.ResultArray)
}

type DataProviderInterface interface {
	GetData() ([][]int, error)
}
