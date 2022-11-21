package rows_loader

import (
	"github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/util"
	"os"
)

type HttpRowsLoader struct {
	Downloader util.HttpDownloaderInterface
	Unziper    util.UnziperInterface
}

func (l *HttpRowsLoader) Load(filename string) ([][]string, error) {
	dest := "archive.zip"

	err := l.Downloader.DownloadFile(filename, dest)
	if err != nil {
		return nil, err
	}
	defer os.Remove(dest)

	err = l.Unziper.UnzipFile(dest, "filename in zip file", "data.txt")
	if err != nil {
		return nil, err
	}
	defer os.Remove("filename")

	return util.ReadCsv("output/"+l.DataFilename, '\t')
}
