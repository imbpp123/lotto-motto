package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

type NumPeriod struct {
	weight int
	qty    int
	Min    int
	Max    int
}

type NumPeriodArray struct {
	periods []*NumPeriod
}

type ResultArray []int

func (np *NumPeriod) ClearData() {
	np.weight = 0
	np.qty = 0
}

func (np *NumPeriod) IncQty(num int) {
	if num > np.Min && num < np.Max {
		np.qty++
	}
}

func (np *NumPeriod) RandomNumber() int {
	return rand.Intn(np.Max-np.Min) + np.Min
}

func (r ResultArray) Contains(num int) (contain bool) {
	contain = false
	for i := 0; i < len(r); i++ {
		contain = r[i] == num
		if contain {
			break
		}
	}
	return
}

func (r ResultArray) Sort() {
	sort.Ints(r)
}

func (npa *NumPeriodArray) clearData() {
	for _, num := range npa.periods {
		num.ClearData()
	}
}

func (npa *NumPeriodArray) incQty(num int) {
	for _, numPeriod := range npa.periods {
		numPeriod.IncQty(num)
	}
}

func (npa *NumPeriodArray) SetData(data [][]int) {
	for _, row := range data {
		for _, cell := range row {
			npa.incQty(cell)
		}
	}
}

func (npa *NumPeriodArray) CalculateWeight() error {
	qtySumm := 0

	for _, numPeriod := range npa.periods {
		qtySumm += numPeriod.qty
	}

	if qtySumm == 0 {
		return errors.New("QTY Summ can not be 0")
	}

	for _, numPeriod := range npa.periods {
		numPeriod.weight = int(math.Round(float64(numPeriod.qty * 100 / qtySumm)))
	}

	return nil
}

func (npa *NumPeriodArray) RandPeriod() (*NumPeriod, error) {
	weightSum := 0
	for _, periodStruct := range npa.periods {
		weightSum += periodStruct.weight
	}

	randValue := rand.Intn(weightSum)

	for _, periodStruct := range npa.periods {
		if randValue-periodStruct.weight < 0 {
			return periodStruct, nil
		}

		weightSum -= periodStruct.weight
		randValue -= periodStruct.weight
	}

	return nil, errors.New("can't find proper period")
}

func (npa *NumPeriodArray) GenerateRandomValues(count int) (ResultArray, error) {
	result := ResultArray{}

	for i := 0; i < count; i++ {
		period, err := npa.RandPeriod()
		if err != nil {
			return nil, fmt.Errorf("can't generate random value: %w", err)
		}

		for {
			randNumber := period.RandomNumber()
			if !result.Contains(randNumber) {
				result = append(result, randNumber)
				break
			}
		}
	}

	return result, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	numNormal := NumPeriodArray{
		periods: []*NumPeriod{
			&NumPeriod{
				Min: 1,
				Max: 10,
			},
			&NumPeriod{
				Min: 11,
				Max: 20,
			},
			&NumPeriod{
				Min: 21,
				Max: 30,
			},
			&NumPeriod{
				Min: 31,
				Max: 40,
			},
			&NumPeriod{
				Min: 41,
				Max: 50,
			},
		},
	}
	numExtra := NumPeriodArray{
		periods: []*NumPeriod{
			&NumPeriod{
				Min: 1,
				Max: 4,
			},
			&NumPeriod{
				Min: 5,
				Max: 8,
			},
			&NumPeriod{
				Min: 9,
				Max: 12,
			},
		},
	}

	normalRow := [][]int{
		{3, 13, 33, 36, 47},
		{13, 6, 32, 21, 8},
	}
	numNormal.SetData(normalRow)
	err := numNormal.CalculateWeight()
	if err != nil {
		panic(err)
	}

	extraRow := [][]int{
		{2, 11},
		{3, 6},
	}
	numExtra.SetData(extraRow)
	err = numExtra.CalculateWeight()
	if err != nil {
		panic(err)
	}

	resultNormal, err := numNormal.GenerateRandomValues(5)
	resultNormal.Sort()
	if err != nil {
		panic(err)
	}

	resultExtra, err := numExtra.GenerateRandomValues(2)
	resultExtra.Sort()
	if err != nil {
		panic(err)
	}

	fmt.Println("Normal rows:")
	for idx, num := range resultNormal {
		fmt.Printf("%d = %d\n", idx+1, num)
	}

	fmt.Println("Extra rows:")
	for idx, num := range resultExtra {
		fmt.Printf("%d = %d\n", idx+1, num)
	}
}
