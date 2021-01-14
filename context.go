package behavior

type Context struct {
	target interface{}
}

func (this *Context) Target() interface{} {
	return this.target
}
