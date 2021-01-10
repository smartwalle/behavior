package behavior

// Select 选择行为
// 顺序执行所有子行为，如果都返回 Failure ，则修改自身状态为 Failure 并返回
// 只要有一个子行为返回 Success 或者 Running，则修改自身状态为 Success 或者 Running 并返回
// 类似逻辑或
// 特别注意：如果一个子行为返回 Running 时，需要记录这个行为，下次直接从该行为开始执行
type Select struct {
	compositeBehavior
	lastRunningIndex int
}

func NewSelect(children ...IBehavior) *Select {
	var n = &Select{}
	n.children = children
	return n
}

func (this *Select) Reset() {
	this.compositeBehavior.Reset()
	this.lastRunningIndex = 0
}

func (this *Select) Exec(ctx Context) {
	if this.status != Running {
		this.lastRunningIndex = 0
	}

	var childStatus = Failure
	for i := this.lastRunningIndex; i < len(this.children); i++ {
		var child = this.children[i]
		child.Exec(ctx)

		childStatus = child.Status()

		if childStatus == Success {
			break
		}

		if childStatus == Failure {
			// 如果执行失败，则继续执行
			continue
		}

		if childStatus == Running {
			this.lastRunningIndex = i
			break
		}
	}
	this.status = childStatus
}
