package behavior

// Composite 组合行为。
type Composite struct {
	base
	children []Behavior
}

func (this *Composite) OnEnter(Context) {
}

func (this *Composite) OnOpen(Context) {
}

func (this *Composite) OnExec(Context) Status {
	return Failure
}

func (this *Composite) OnClose(Context) {
}

func (this *Composite) OnExit(Context) {
}
