package command_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/application/command"
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

	mockRepository.EXPECT().LoadFromFile().never()

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