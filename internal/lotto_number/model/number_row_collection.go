package model

import "sort"

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