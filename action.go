package behavior

// Action 具体行为
type Action struct {
	baseBehavior
}

func NewAction() *Action {
	var n = &Action{}
	return n
}
