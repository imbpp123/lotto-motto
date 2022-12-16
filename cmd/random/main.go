package main

import (
	"fmt"
	"github.com/imbpp123/lotto_motto/internal/lotto_number/model"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	numNormal := model.NewNumPeriodCollection([]*model.NumPeriod{
		model.NewNumPeriod(1, 10),
		model.NewNumPeriod(11, 20),
		model.NewNumPeriod(21, 30),
		model.NewNumPeriod(21, 40),
		model.NewNumPeriod(41, 50),
	})
	numExtra := model.NewNumPeriodCollection([]*model.NumPeriod{
		model.NewNumPeriod(1, 4),
		model.NewNumPeriod(5, 8),
		model.NewNumPeriod(9, 12),
	})

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
	for idx, num := range resultNormal.Result() {
		fmt.Printf("%d = %d\n", idx+1, num)
	}

	fmt.Println("Extra rows:")
	for idx, num := range resultExtra.Result() {
		fmt.Printf("%d = %d\n", idx+1, num)
	}
}
