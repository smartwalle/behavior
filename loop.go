package behavior

// Loop 循环行为
type Loop struct {
	compositeBehavior
}

func NewLoop() *Loop {
	var n = &Loop{}
	return n
}
