package behavior

// Decorator 修饰行为
type Decorator struct {
	base
	child Behavior
}

func (this *Decorator) OnEnter(Context) {
}

func (this *Decorator) OnOpen(Context) {
}

func (this *Decorator) OnExec(Context) Status {
	return Error
}

func (this *Decorator) OnClose(Context) {
}

func (this *Decorator) OnExit(Context) {
}
