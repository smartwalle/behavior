package behavior

// UntilFailure 直到子行为失败。
// 每一次 Tick 执行一次子行为，直到子行为返回 Failure。
// 如果子行为返回 Failure，则返回 Failure，否则返回 Running。
// 如果 count 参数的值大于 0，则表示重复执行到 count 次数时，一定返回 Failure。
// 比如：count 参数的值为 2 时，本行为执行第二次的时候，一定返回 Failure。
type UntilFailure struct {
	Decorator
	count   int
	current int
}

func NewUntilFailure(count int, child Behavior) *UntilFailure {
	var n = &UntilFailure{}
	n.SetWorker(n)
	n.child = child
	n.count = count
	return n
}

func (this *UntilFailure) OnOpen(ctx Context) {
	this.current = 0
}

func (this *UntilFailure) OnExec(ctx Context) Status {
	var status = this.child.Tick(ctx)
	this.current++

	if this.count > 0 && this.current >= this.count {
		return Failure
	}
	if status == Failure {
		return status
	}
	return Running
}
