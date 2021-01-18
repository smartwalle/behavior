package behavior

type UntilSuccess struct {
	Decorator
}

func NewUntilSuccess(child Behavior) *UntilSuccess {
	var n = &UntilSuccess{}
	n.SetWorker(n)
	n.child = child
	return n
}

func (this *UntilSuccess) OnExec(ctx Context) Status {
	var status = this.child.Tick(ctx)
	if status == Success {
		return status
	}
	return Running
}
