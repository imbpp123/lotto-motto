package data_provider_test

import (
	"github.com/golang/mock/gomock"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/data_provider"
	"testing"
)

func TestEuroJackpot_GetData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	loader := data_provider.NewEuroJackpot(3)
	data, err := loader.GetData()
	if err != nil {
		t.Error(err)
	}
	if len(data) != 3 {
		t.Error("result array is not equal 3")
	}
	for _, row := range data {
		if len(row) != 7 {
			t.Error("result row is not equal 7: 5 + 2")
		}
	}
}

func TestData6Aus49_GetData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	loader := data_provider.NewLotto6Aus49(3)
	data, err := loader.GetData()
	if err != nil {
		t.Error(err)
	}
	if len(data) != 3 {
		t.Error("result array is not equal 3")
	}
	for _, row := range data {
		if len(row) != 6 {
			t.Error("result row is not equal 6")
		}
	}
}
