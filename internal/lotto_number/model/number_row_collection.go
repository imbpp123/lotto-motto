package model

type NumberRowCollection struct {
	numberRow []NumberRow
}

func (c *NumberRowCollection) SortByDate() *NumberRowCollection {
	return c
}

func (c *NumberRowCollection) Slice(startIndex uint, length uint) *NumberRowCollection {
	return c
}