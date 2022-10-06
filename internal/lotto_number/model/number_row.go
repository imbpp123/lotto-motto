package model

import "time"

type NumberRow struct {
	numbers []Number
	date time.Time
}

func (nr *NumberRow) add(number uint64, valueType uint16) {
	nr.numbers = append(nr.numbers, Number{number, valueType})
}