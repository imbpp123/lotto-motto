package model

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
)

type NumPeriodCollection struct {
	periods []*NumPeriod
}

func NewNumPeriodCollection(periods []*NumPeriod) *NumPeriodCollection {
	return &NumPeriodCollection{periods}
}

func NewNumPeriodCollectionByParts(parts int, min int, max int) *NumPeriodCollection {
	step := (max - min + 1) / parts

	collection := NumPeriodCollection{}

	currentMin := min
	for {
		period := NewNumPeriod(currentMin, currentMin+step-1)
		collection.periods = append(collection.periods, period)

		if period.max >= max {
			period.max = max
			break
		}
		currentMin = period.max + 1
	}

	return &collection
}

func (npa *NumPeriodCollection) SortByWeight() {
	sort.Slice(npa.periods, func(i, j int) bool {
		return npa.periods[i].weight > npa.periods[j].weight
	})
}

func (npa *NumPeriodCollection) clearData() {
	for _, num := range npa.periods {
		num.ClearData()
	}
}

func (npa *NumPeriodCollection) incQty(num int) {
	for _, numPeriod := range npa.periods {
		numPeriod.IncQty(num)
	}
}

func (npa *NumPeriodCollection) SetData(data [][]int) {
	fmt.Println("Data:")
	for _, row := range data {
		for _, cell := range row {
			npa.incQty(cell)
		}
	}
}

func (npa *NumPeriodCollection) CalculateWeight() error {
	for _, numPeriod := range npa.periods {
		numPeriod.weight = numPeriod.qty
		fmt.Printf("Period min=%d max=%d weight=%d\n", numPeriod.min, numPeriod.max, numPeriod.weight)
	}

	return nil
}

func (npa *NumPeriodCollection) RandPeriod() (*NumPeriod, error) {
	weightSum := 0
	for _, periodStruct := range npa.periods {
		weightSum += periodStruct.weight
	}

	if weightSum == 0 {
		return nil, errors.New("weight sum can not be 0")
	}

	randValue := rand.Intn(weightSum)
	fmt.Printf("Random number generation: rand=%d total_weight=%d\n", randValue, weightSum)

	for _, periodStruct := range npa.periods {
		if randValue-periodStruct.weight < 0 {
			return periodStruct, nil
		}

		weightSum -= periodStruct.weight
		randValue -= periodStruct.weight
	}

	return nil, errors.New("can't find proper period")
}

func (npa *NumPeriodCollection) GenerateRandomValues(count int) (*ResultArray, error) {
	result := &ResultArray{}

	for i := 0; i < count; i++ {
		period, err := npa.RandPeriod()
		if err != nil {
			return nil, fmt.Errorf("can't generate random value: %w", err)
		}

		for {
			randNumber := period.RandomNumber()
			if !result.Contains(randNumber) {
				result.Add(randNumber)
				break
			}
		}
	}

	return result, nil
}
