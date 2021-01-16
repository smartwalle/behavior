package behavior

// Repeat 重复执行指定行为。
// 重复执行子行为，达到最大限定次数之后，返回 Success。
// 子行为返回 Success 或者 Failure，算作一次执行。
// count 参数小于 0 的时候，将不作重复次数限定。
type Repeat struct {
	Decorator
	count   int
	current int
}

func NewRepeat(count int, child Behavior) *Repeat {
	var n = &Repeat{}
	n.SetWorker(n)
	n.count = count
	n.child = child
	return n
}

func (this *Repeat) OnOpen(ctx Context) {
	this.current = 0
}

func (this *Repeat) OnExec(ctx Context) Status {
	if this.child == nil {
		return Error
	}
	var status = Success
	for this.count <= 0 || this.current < this.count {
		status = this.child.Exec(ctx)
		if status == Success || status == Failure {
			this.current++
		} else {
			break
		}
	}
	return status
}
