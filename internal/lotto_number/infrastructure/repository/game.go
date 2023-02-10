package repository

import (
	"context"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/domain/value_object"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/infrastructure/data_provider"
	"time"
)

type ColumnLevel struct {
	Start int
	End   int
	Level int
}

type GameColumns struct {
	Year   int
	Month  int
	Day    int
	Levels []ColumnLevel
}

type GameRepositoryConfig struct {
	Url      string
	Filename string
	Columns  GameColumns
}

type GameRepository struct {
	config GameRepositoryConfig
	reader data_provider.HttpZipReader
}

func NewGameRepository(config GameRepositoryConfig) *GameRepository {
	return &GameRepository{
		config: config,
		reader: data_provider.NewHttpZipReader(
			config.Url,
			config.Filename,
		),
	}
}

func (r *GameRepository) FetchData(ctx context.Context, rows int) ([]value_object.Raffle, error) {
	data, err := r.reader.GetData(rows)
	if err != nil {
		return nil, err
	}

	var result []value_object.Raffle

	// load data to raffle
	for _, raffleVal := range data {
		var numbers []value_object.Number

		for _, level := range r.config.Columns.Levels {
			for _, num := range raffleVal[level.Start:level.End] {
				numbers = append(numbers, *value_object.NewNumber(num, level.Level))
			}
		}

		year := raffleVal[r.config.Columns.Year]
		month := raffleVal[r.config.Columns.Month]
		day := raffleVal[r.config.Columns.Day]
		date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

		result = append(result, *value_object.NewRaffle(numbers, date))
	}

	return result, nil
}
