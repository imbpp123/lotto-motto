package value_object

type Number struct {
	value int
	level int
}

func NewNumber(value int, level int) *Number {
	return &Number{
		value: value,
		level: level,
	}
}

func (n *Number) Value() int {
	return n.value
}

func (n *Number) Level() int {
	return n.level
}
