package behavior

// Sequence 顺序行为
// 顺序执行所有子行为，如果有一个子行为返回 Running 或者 Failure，则修改自身状态为 Running 或者 Failure 并返回
// 类似逻辑与
// 特别注意：如果一个子行为返回 Running 时，需要记录这个行为，下次直接从该行为开始执行
type Sequence struct {
	compositeBehavior
	lastRunningIndex int
}

func NewSequence(children ...IBehavior) *Sequence {
	var n = &Sequence{}
	n.children = children
	return n
}

func (this *Sequence) Reset() {
	this.compositeBehavior.Reset()
	this.lastRunningIndex = 0
}

func (this *Sequence) Exec(ctx Context) {
	if this.status != Running {
		this.lastRunningIndex = 0
	}

	var childStatus = Success
	for i := this.lastRunningIndex; i < len(this.children); i++ {
		var child = this.children[i]
		child.Exec(ctx)

		childStatus = child.Status()

		if childStatus == Success {
			// 如果子行为执行成功，则继续执行
			continue
		}

		if childStatus == Failure {
			break
		}

		if childStatus == Running {
			this.lastRunningIndex = i
			break
		}
	}
	this.status = childStatus
}

func IF(cond ConditionFunc, child IBehavior) IBehavior {
	return NewSequence(NewCondition(cond), child)
}
