package repository

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/util"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/model"
)

type NumberFileRepository struct {

}

func (nfr NumberFileRepository) LoadFromFile(filename string) (*model.NumberRowCollection, error) {
	dest := "archive.zip"

	err := util.DownloadFile(filename, dest)
	if err != nil {
		return nil, err
	}

	unzipFilename, err := util.UnzipFile(dest, "output")
	if err != nil {
		return nil, err
	}

	data, err := util.ReadCsv(unzipFilename, '\t')
	if err != nil {
		return nil, err
	}

	

	collection := new(model.NumberRowCollection)
	return collection, nil
}