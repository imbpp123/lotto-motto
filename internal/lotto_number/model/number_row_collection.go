package model

import (
	"sort"
	"strconv"
	"time"
)

type NumberRowCollection struct {
	rows []NumberRow
}

func (c *NumberRowCollection) Add(row NumberRow) {
	c.rows = append(c.rows, row)
}

func (c *NumberRowCollection) Get(index uint) *NumberRow {
	return &c.rows[index]
}

func (c *NumberRowCollection) Count() int {
	return len(c.rows)
}

func (c *NumberRowCollection) SortByDate() *NumberRowCollection {
	sort.Slice(c.rows, func(i, j int) bool {
		return c.rows[j].Date.Before(c.rows[i].Date)
	})

	return c
}

func (c *NumberRowCollection) Slice(startIndex uint, length int) *NumberRowCollection {
	if length > len(c.rows) {
		return c
	}

	c.rows = c.rows[0:length]
	return c
}

func createDateFromRow(row []string) (time.Time, error) {
	year, err := strconv.Atoi(row[2])
	if err != nil {
		return time.Time{}, err
	}

	month, err := strconv.Atoi(row[1])
	if err != nil {
		return time.Time{}, err
	}

	day, err := strconv.Atoi(row[0])
	if err != nil {
		return time.Time{}, err
	}

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil
}

func createNumberArray(row []string, numberTypeAmount []int) ([]Number, error) {
	numbers := []Number{}

	for idx, amount := range numberTypeAmount {
		for i := 0; i < amount; i++ {
			parsedNumber, err := strconv.Atoi(row[i])
			if err != nil {
				return nil, err
			}			

			numbers = append(numbers, Number{
				Value: parsedNumber,
				ValueType: idx,
			})
		}
	}

	return numbers, nil
}

func CreateFromStringMap(data [][]string, numberTypeAmount []int) (NumberRowCollection, error) {
	numberRows := []NumberRow{}

	for _, row := range data {
		dateTime, err := createDateFromRow(row)
		if err != nil {
			return NumberRowCollection{}, err
		}

		numbers, err := createNumberArray(row[3:], numberTypeAmount)
		if err != nil {
			return NumberRowCollection{}, err
		}

		numberRows = append(numberRows, NewNumberRow(dateTime, numbers))
	}
	
	collection := NumberRowCollection{
		rows: numberRows,
	}

	return collection, nil
}