package repository

import (
	"github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/util"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/model"
)

type NumberEurojackpotRepository struct {

}

func (nfr NumberEurojackpotRepository) LoadFromFile(filename string) (*model.NumberRowCollection, error) {
	dest := "archive.zip"

	err := util.DownloadFile(filename, dest)
	if err != nil {
		return "", err
	}

	unzipFilename, err := util.UnzipFile(dest, "output")
	if err != nil {
		return "", err
	}

	data, err := util.ReadCsv(unzipFilename, 't')

	collection := new(model.NumberRowCollection)
	
	for (data) {
		collection.Add()
	}

	return util.ReadCsv(unzipFilename, 't')
}
