package model

import "math/rand"

type NumPeriod struct {
	weight int
	qty    int
	min    int
	max    int
}

func (np *NumPeriod) ClearData() {
	np.weight = 0
	np.qty = 0
}

func (np *NumPeriod) IncQty(num int) {
	if num > np.min && num < np.max {
		np.qty++
	}
}

func (np *NumPeriod) RandomNumber() int {
	return rand.Intn(np.max-np.min) + np.min
}

func NewNumPeriod(min int, max int) *NumPeriod {
	return &NumPeriod{min: min, max: max}
}
