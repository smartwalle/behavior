package behavior

// Repeater 重复执行子行为。
// 每次 Tick 都会重复执行子行为固定次数，当子行为返回 Running 的时候，会中止本次重复行为，等待下次 Tick 从上次中止的次数开始继续执行。
// 达到最大限定次数之后，返回 Success。
// 子行为返回 Success 或者 Failure，算作一次执行。
// count 参数小于 0 的时候，将不作重复次数限定。
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
	var status = Success
	for this.count <= 0 || this.current < this.count {
		status = this.child.Tick(ctx)
		if status == Success || status == Failure {
			this.current++
		} else {
			break
		}
	}
	return status
}
