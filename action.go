package behavior

// Action 具体行为。
type Action struct {
	base
}

func (this *Action) OnEnter(Context) {
}

func (this *Action) OnOpen(Context) {
}

func (this *Action) OnExec(Context) Status {
	return Error
}

func (this *Action) OnClose(Context) {
}

func (this *Action) OnExit(Context) {
}
