package behavior

// Sequence 顺序行为，
// 顺序执行所有子行为，如果有一个子行为返回 Running 或者 Failure，则返回 Running 或者 Failure， 类似逻辑与。
// 如果一个子行为返回 Running 时，会记录这个行为，下次直接从该行为开始执行。
type Sequence struct {
	composite
	lastRunningIndex int
}

func NewSequence(children ...IBehavior) *Sequence {
	var n = &Sequence{}
	n.children = children
	return n
}

func (this *Sequence) Reset() {
	this.lastRunningIndex = 0
}

func (this *Sequence) OnExec(ctx *Context) Status {
	for i := this.lastRunningIndex; i < len(this.children); i++ {
		var child = this.children[i]
		var status = child.OnExec(ctx)

		if status != Success {
			if status == Running {
				this.lastRunningIndex = i
			}
			return status
		}
	}
	return Success
}

func IF(cond ConditionFunc, child IBehavior) IBehavior {
	return NewSequence(NewCondition(cond), child)
}
