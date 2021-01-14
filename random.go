package behavior

// Random 随机行为
type Random struct {
	composite
}

func NewRandom() *Random {
	var n = &Random{}
	return n
}
