package behavior

// Loop 循环行为
type Loop struct {
	composite
}

func NewLoop() *Loop {
	var n = &Loop{}
	return n
}
