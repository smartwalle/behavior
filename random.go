package behavior

// Random 随机行为
type Random struct {
	compositeBehavior
}

func NewRandom() *Random {
	var n = &Random{}
	return n
}
