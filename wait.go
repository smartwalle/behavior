package behavior

// Wait 等待行为
type Wait struct {
	compositeBehavior
}

func NewWait() *Wait {
	var n = &Wait{}
	return n
}
