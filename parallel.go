package behavior

// Parallel 并行行为
type Parallel struct {
	compositeBehavior
}

func NewParalle() *Parallel {
	var n = &Parallel{}
	return n
}
