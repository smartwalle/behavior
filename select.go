package behavior

// Select 选择行为，
// 顺序执行所有子行为，如果一个子行为返回 Success 或者 Running，则返回 Success 或者 Running。 类似逻辑或。
// 如果一个子行为返回 Running 时，会记录这个行为，下次直接从该行为开始执行。
type Select struct {
	Composite
	lastRunningIndex int
}

func NewSelect(children ...IBehavior) *Select {
	var n = &Select{}
	n.SetWorker(n)
	n.children = children
	return n
}

func (this *Select) OnOpen(ctx Context) {
	this.lastRunningIndex = 0
}

func (this *Select) OnExec(ctx Context) Status {
	for i := this.lastRunningIndex; i < len(this.children); i++ {
		var child = this.children[i]
		var status = child.Exec(ctx)

		if status != Failure {
			if status == Running {
				this.lastRunningIndex = i
			}
			return status
		}
	}
	return Failure
}
