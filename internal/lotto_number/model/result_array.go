package model

import "sort"

type ResultArray struct {
	result []int
}

func (r *ResultArray) Add(num int) {
	r.result = append(r.result, num)
}

func (r *ResultArray) Contains(num int) (contain bool) {
	contain = false
	for i := 0; i < len(r.result); i++ {
		contain = r.result[i] == num
		if contain {
			break
		}
	}
	return
}

func (r *ResultArray) Sort() {
	sort.Ints(r.result)
}

func (r *ResultArray) Result() []int {
	return r.result
}
