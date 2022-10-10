package loader

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type HttpFileLoader struct {
}

func (tl *HttpFileLoader) loadFile(filenameUrl string) (string, error) {
	dest := "output/archive.zip"

	err := downloadFile(filenameUrl, dest)
	if err != nil {
		return "", err
	}

	return unzipFile(dest, "output")
}

// Download file from internet and save it to disk
func downloadFile(filenameUrl string, destFilename string) error {
	// Create blank file
	file, err := os.Create(destFilename)
	if err != nil {
		return err
	}

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	// Put content on file
	resp, err := client.Get(filenameUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	fmt.Printf("Downloaded a file %s with size %d", filenameUrl, size)

	return nil
}

// unzip file
func unzipFile(zipFilename string, dstFolder string) (string, error) {
	archive, err := zip.OpenReader(zipFilename)
	if err != nil {
		return "", err
	}
	defer archive.Close()

	for _, f := range archive.File {
		filePath := filepath.Join(dstFolder, f.Name)
		fmt.Println("unzipping file ", filePath)

		if !strings.HasPrefix(filePath, filepath.Clean(dstFolder)+string(os.PathSeparator)) {
			fmt.Println("invalid file path")
		}
		if f.FileInfo().IsDir() {
			fmt.Println("creating directory...")
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return "", err
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return "", err
		}

		fileInArchive, err := f.Open()
		if err != nil {
			return "", err
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			return "", err
		}

		dstFile.Close()
		fileInArchive.Close()
	}

	return "", nil
}
