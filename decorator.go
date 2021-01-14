package behavior

// Decorator 修饰行为
type Decorator struct {
	composite
}

func NewDecorator() *Decorator {
	var n = &Decorator{}
	return n
}
