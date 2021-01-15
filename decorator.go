package behavior

// Decorator 修饰行为
type Decorator struct {
	Composite
}

func NewDecorator() *Decorator {
	var n = &Decorator{}
	return n
}
