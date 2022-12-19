package model_test

import (
	"github.com/golang/mock/gomock"
	model2 "github.com/imbpp123/lotto_motto/internal/lotto_number/model"
	"testing"
)

func TestNumPeriod_IncQty(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	model := model2.NewNumPeriod(1, 10)
	for _, num := range []int{0, 1, 2, 9, 10, 11} {
		model.IncQty(num)
	}
	if model.Qty() != 4 {
		t.Error("Error in QTY calculations")
	}
}
