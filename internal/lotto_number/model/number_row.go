package model

import "time"

type NumberRow struct {
	Numbers []Number
	Date time.Time
}

func NewNumberRow(rowDate time.Time, numbers ...Number) NumberRow {
	return NumberRow{
		numbers,
		rowDate,
	}
}
