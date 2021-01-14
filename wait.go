package behavior

// Wait 等待行为
type Wait struct {
	composite
}

func NewWait() *Wait {
	var n = &Wait{}
	return n
}
