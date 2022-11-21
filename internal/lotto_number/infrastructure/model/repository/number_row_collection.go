package repository

import (
	"github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/util"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/model"
	"os"
)

type NumberFileRepository struct {
	Downloader       util.HttpDownloaderInterface
	Unziper          util.UnziperInterface
	NumberTypeAmount []int
	DataFilename     string
}

func (n *NumberFileRepository) LoadFromFile(filename string) (*model.NumberRowCollection, error) {
	dest := "archive.zip"

	err := n.Downloader.DownloadFile(filename, dest)
	if err != nil {
		return nil, err
	}
	defer os.Remove(dest)

	err = n.Unziper.UnzipFile(dest, "filename in zip file", "data.txt")
	if err != nil {
		return nil, err
	}

	data, err := util.ReadCsv("output/"+n.DataFilename, '\t')
	if err != nil {
		return nil, err
	}

	collection, err := model.CreateFromStringMap(data, n.NumberTypeAmount)

	return &collection, err
}
