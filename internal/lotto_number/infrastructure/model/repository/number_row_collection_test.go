package repository_test

import (
	"github.com/golang/mock/gomock"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/model/repository"
	"testing"
)

func TestNumberFileRepository_LoadFromFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	numberRepository := repository.NumberFileRepository{}
	collection, err := numberRepository.LoadFromFile("https://www.lotto-berlin.de/static/gamebroker_7/default/download_files/archiv_lotto.zip")
	if err == nil {
		t.Error("Error should be nil")
	}

	if collection.Count() == 0 {
		t.Error("Lines count should be more than 0")
	}
}
