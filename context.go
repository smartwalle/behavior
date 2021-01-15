package behavior

type Context interface {
	Target() interface{}
}

type context struct {
	target interface{}
}

func (this *context) Target() interface{} {
	return this.target
}

func NewContext(target interface{}) Context {
	var ctx = &context{}
	ctx.target = target
	return ctx
}
