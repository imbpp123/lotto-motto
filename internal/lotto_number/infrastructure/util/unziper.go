package util

import (
	"archive/zip"
	"io"
	"os"
)

type UnziperInterface interface {
	UnzipFile(zipFilename string, sourceFile string, destFile string) error
}

type UnziperArchive struct {
}

func (u *UnziperArchive) UnzipFile(zipFilename string, sourceFile string, destFile string) error {
	archive, err := zip.OpenReader(zipFilename)
	if err != nil {
		return err
	}
	defer archive.Close()

	for _, f := range archive.File {
		if f.FileInfo().Name() != sourceFile {
			continue
		}

		dstFile, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		fileInArchive, err := f.Open()
		if err != nil {
			return err
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			return err
		}

		dstFile.Close()
		fileInArchive.Close()
	}

	return nil
}
