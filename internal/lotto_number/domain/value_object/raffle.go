package value_object

import (
	"sort"
	"time"
)

type Raffle struct {
	numbers []Number
	date    time.Time
}

func NewRaffle(numbers []Number, date time.Time) *Raffle {
	return &Raffle{
		numbers: numbers,
		date:    date,
	}
}

func (r *Raffle) Date() time.Time {
	return r.date
}

func (r *Raffle) NumberByLevel(level int) []int {
	var data []int

	for _, val := range r.numbers {
		if val.Level() == level {
			data = append(data, val.Value())
		}
	}

	sort.Ints(data)

	return data
}

func (r *Raffle) LevelCount() int {
	if len(r.numbers) == 0 {
		return 0
	}

	count := 1
	currentLevel := 0

	for _, number := range r.numbers {
		if number.level != currentLevel {
			count++
		}
	}

	return count
}
