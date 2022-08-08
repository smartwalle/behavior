package behavior

// UntilSuccess 直到子行为成功。
// 每一次 Tick 执行一次子行为，直到子行为返回 Success。
// 如果子行为返回 Success，则返回 Success，否则返回 Running。
// 如果 count 参数的值大于 0，则表示重复执行到 count 次数时，一定返回 Success。
// 比如：count 参数的值为 2 时，本行为执行第二次的时候，一定返回 Success。
type UntilSuccess struct {
	Decorator
	count   int
	current int
}

func NewUntilSuccess(count int, child Behavior) *UntilSuccess {
	var n = &UntilSuccess{}
	n.SetWorker(n)
	n.child = child
	n.count = count
	return n
}

func (this *UntilSuccess) OnOpen(ctx Context) {
	this.current = 0
}

func (this *UntilSuccess) OnExec(ctx Context) Status {
	var status = this.child.Tick(ctx)
	this.current++

	if this.count > 0 && this.current >= this.count {
		return Success
	}
	if status == Success {
		return status
	}
	return Running
}
