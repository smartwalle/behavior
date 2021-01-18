package behavior

type UntilFailure struct {
	Decorator
}

func NewUntilFailure(child Behavior) *UntilFailure {
	var n = &UntilFailure{}
	n.SetWorker(n)
	n.child = child
	return n
}

func (this *UntilFailure) OnExec(ctx Context) Status {
	var status = this.child.Tick(ctx)
	if status == Failure {
		return status
	}
	return Running
}
