package behavior

// Random 随机行为
type Random struct {
	Composite
}

func NewRandom() *Random {
	var n = &Random{}
	return n
}
