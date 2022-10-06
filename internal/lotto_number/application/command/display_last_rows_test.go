package command_test

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/application/command"
	mock_presentation "github.com/imbpp123/lotto_motto/internal/lotto_number/application/presentation/mock"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/model"
	mock_repository "github.com/imbpp123/lotto_motto/internal/lotto_number/model/repository/mock"
)

func TestHandle_ErrorLoadFromFile(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dummyError := errors.New("dummy error")
	
	cmd := command.DisplayLastRowCommand{
		RowCount: 4,
		Filename: "http://test.com/file.zip",
	}

	mockRepository := mock_repository.NewMockNumberRowCollectionRepository(ctrl)
	mockRepository.EXPECT().LoadFromFile("http://test.com/file.zip").Return(nil, dummyError)

	handler := command.DisplayLastRowHandler{
		NumberRowCollectionRepository: mockRepository,
		NumberPresentation: nil,
	}

	// act
	err := handler.Handle(cmd)

	// assert
	if err == nil {
		t.Error("Error should be returned, but nil instead")
	}
}

func TestHandle_Success(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	cmd := command.DisplayLastRowCommand{
		RowCount: 1,
		Filename: "http://test.com/file.zip",
	}

	collection := model.NumberRowCollection{}
	collection.Add(model.NewNumberRow(time.Now().Add(-1000), model.Number{1, 1}))
	collection.Add(model.NewNumberRow(time.Now().Add(-100), model.Number{2, 2}))
	collection.Add(model.NewNumberRow(time.Now().Add(-10), model.Number{3, 3}))

	mockRepository := mock_repository.NewMockNumberRowCollectionRepository(ctrl)
	mockRepository.EXPECT().LoadFromFile("http://test.com/file.zip").Return(&collection, nil)

	mockPresentation := mock_presentation.NewMockNumberRowCollectionPresentation(ctrl)
	mockPresentation.EXPECT().Display(&collection)

	handler := command.DisplayLastRowHandler{
		NumberRowCollectionRepository: mockRepository,
		NumberPresentation: mockPresentation,
	}

	// act
	err := handler.Handle(cmd)

	// assert
	if err != nil {
		t.Error("Error is returned instead of collection")
	}
	if collection.Count() != 1 {
		t.Error("Collection count should be 1")
	}
	// check date and number for row too!
	// row := collection.Get(0)
}