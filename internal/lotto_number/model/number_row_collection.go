package model

type NumberRowCollection struct {
	rows []NumberRow
}

func (c *NumberRowCollection) Add(row NumberRow) {
	c.rows = append(c.rows, row)
}

func (c *NumberRowCollection) Get(index uint) *NumberRow {
	return &c.rows[index]
}

func (c *NumberRowCollection) Count() uint {
	return 0
}

func (c *NumberRowCollection) SortByDate() *NumberRowCollection {
	return c
}

func (c *NumberRowCollection) Slice(startIndex uint, length uint) *NumberRowCollection {
	return c
}