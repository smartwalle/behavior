package behavior

// Repeater 重复执行指定行为
// 重复执行子行为，达到最大限定次数之后，返回 Success
// 子行为返回 Success 或者 Failure，算作一次执行
// count 参数小于 0 的时候，表示不作重复次数限定
type Repeater struct {
	Decorator
	count   int
	current int
}

func NewRepeater(count int, child Behavior) *Repeater {
	var n = &Repeater{}
	n.SetWorker(n)
	n.count = count
	n.child = child
	return n
}

func (this *Repeater) OnOpen(ctx Context) {
	this.current = 0
}

func (this *Repeater) OnExec(ctx Context) Status {
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
