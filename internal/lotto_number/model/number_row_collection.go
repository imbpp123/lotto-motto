package model

import (
	"sort"
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

// [5, 2] [5]
func CreateFromMap(data [][]string, numberTypeAmount []int) (*NumberRowCollection, error) {
	numberRows := []NumberRow{}
	for _, row := range data {
		dateTime, err := time.Parse(row[2] + "-" + row[1] + "-" + row[0])
		if err != nil {
			return nil, err
		}

		numbers := []Number{}
		for idx, numbers := range numberTypeAmount {
			numbers := append(numbers, Number{
				Value: numbers,
				ValueType: idx,
			})
		}

		numberRows = append(numberRows, NewNumberRow(dateTime, row[3:8], row[8:]))
	}
	
	collection := NumberRowCollection{
		rows: numberRows,
	}

	return collection, nil
}