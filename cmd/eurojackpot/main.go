package eurojackpot

import (
	"errors"
	"fmt"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/util"
	"math/rand"
	"os"
	"time"
)

type NumPeriod struct {
	Weight int
	Min    int
	Max    int
}

type NumPeriodArray []NumPeriod

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
	defer func() {
		err := os.Remove(dest)
		if err != nil {
			fmt.Println(err)
		}
	}()

	err = l.Unziper.UnzipFile(dest, "filename in zip file", "data.txt")
	if err != nil {
		return nil, err
	}
	defer func() {
		err := os.Remove("filename")
		if err != nil {
			fmt.Println(err)
		}
	}()

	return util.ReadCsv("output/data.txt", '\t')
}

func NewNumPeriodArray(data [][]string) *NumPeriodArray {
	return nil
}

func (npa *NumPeriodArray) RandWithWeight(arr []NumPeriod) (*NumPeriod, error) {
	weightSum := 0
	for _, periodStruct := range arr {
		weightSum += periodStruct.Weight
	}

	randValue := rand.Intn(weightSum)

	for _, periodStruct := range arr {
		if randValue-periodStruct.Weight < 0 {
			return &periodStruct, nil
		}

		weightSum -= periodStruct.Weight
		randValue -= periodStruct.Weight
	}

	return nil, errors.New("can't find proper period")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	httpRows := HttpRowsLoader{
		Downloader: &util.HttpDownloaderNet{},
		Unziper:    &util.UnziperArchive{},
	}

	data, err := httpRows.Load("https://www.lotto-berlin.de/static/gamebroker_7/default/download_files/archiv_eurojackpot.zip")
	if err != nil {
		panic(err)
	}

	nums := NewNumPeriodArray(data)

	for _, num := range nums {
		currentNumber, err := nums.RandWithWeight()
		if err != nil {
			panic(err)
		}

		lottoNumber := rand.Intn(currentNumber.Max-currentNumber.Min) + currentNumber.Min

	}

	fmt.Printf("1 = %s\n", randWithWeight(data))
	fmt.Printf("2 = %s\n", randWithWeight(data))
	fmt.Printf("3 = %s\n", randWithWeight(data))
	fmt.Printf("4 = %s\n", randWithWeight(data))
	fmt.Printf("5 = %s\n", randWithWeight(data))
	fmt.Printf("6 = %s\n", randWithWeight(data))
}
