package presentation

import "github.com/imbpp123/lotto_motto/internal/lotto_number/model"

type NumberRowCollectionPresentation interface {
	Display(collection *model.NumberRowCollection) 
}