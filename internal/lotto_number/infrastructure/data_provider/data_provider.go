package data_provider

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type HttpZipReader interface {
	GetData(rows int) ([][]int, error)
}

type httpZipCsvReader struct {
	fileUrl     string
	csvFilename string
}

func NewHttpZipReader(url string, filename string) HttpZipReader {
	return &httpZipCsvReader{
		fileUrl:     url,
		csvFilename: filename,
	}
}

func (p *httpZipCsvReader) GetData(rows int) ([][]int, error) {
	resp, err := http.Get(p.fileUrl)
	if err != nil {
		return nil, fmt.Errorf("can't download file: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can't read file: %w", err)
	}

	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return nil, fmt.Errorf("can't create zip reader: %w", err)
	}

	// Read all the files from zip archive
	var records [][]string

	for _, zipFile := range zipReader.File {
		f, err := zipFile.Open()
		if err != nil {
			continue
		}
		defer f.Close()

		if zipFile.FileInfo().Name() != p.csvFilename {
			continue
		}

		csvReader := csv.NewReader(f)
		csvReader.Comma = '\t'
		records, err = csvReader.ReadAll()
		if err != nil {
			return nil, err
		}
	}

	return p.getProperRecords(records, rows)
}

func (p *httpZipCsvReader) getProperRecords(data [][]string, count int) ([][]int, error) {
	var records [][]int

	length := len(data)
	for i := length - 1; i > length-count-1; i-- {
		dataInt, err := p.sliceAtoi(data[i])
		if err != nil {
			return nil, fmt.Errorf("can't convert string to int: %w", err)
		}
		records = append(records, dataInt)
	}

	return records, nil
}

func (p *httpZipCsvReader) sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		number := strings.TrimSpace(a)
		if len(number) == 0 {
			continue
		}
		i, err := strconv.Atoi(number)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}
