package repository

import "github.com/imbpp123/lotto_motto/internal/lotto_number/model"

type NumberRowCollectionRepository interface {
	LoadFromFile(filename string) (*model.NumberRowCollection, error)
}
