package behavior

// Action 具体行为
type Action struct {
	base
}

func NewAction() *Action {
	var n = &Action{}
	return n
}
