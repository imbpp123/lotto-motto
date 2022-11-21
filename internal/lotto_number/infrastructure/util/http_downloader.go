package util

import (
	"io"
	"net/http"
	"os"
)

type HttpDownloaderInterface interface {
	DownloadFile(filenameUrl string, destFilename string) error
}

type HttpDownloaderNet struct {
}

func (h *HttpDownloaderNet) DownloadFile(filenameUrl string, destFilename string) error {
	// Create blank file
	file, err := os.Create(destFilename)
	if err != nil {
		return err
	}
	defer file.Close()

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

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
