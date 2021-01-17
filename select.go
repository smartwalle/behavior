package behavior

// PrioritySelect 优先选择行为。
// 每次 Tick 从左到右顺序执行所有子行为，如果一个子行为返回 Success 或者 Running，则返回 Success 或者 Running，类似逻辑或。
// 如果一个子行为返回 Running 时，会记录这个行为，下次直接从该行为开始执行。
type PrioritySelect struct {
	Composite
	lastRunningIndex int
}

func NewPrioritySelect(children ...Behavior) *PrioritySelect {
	var n = &PrioritySelect{}
	n.SetWorker(n)
	n.children = children
	return n
}

func (this *PrioritySelect) OnOpen(ctx Context) {
	this.lastRunningIndex = 0
}

func (this *PrioritySelect) OnExec(ctx Context) Status {
	for i := this.lastRunningIndex; i < len(this.children); i++ {
		var child = this.children[i]
		var status = child.Tick(ctx)

		if status != Failure {
			if status == Running {
				this.lastRunningIndex = i
			}
			return status
		}
	}
	return Failure
}

// NonPrioritySelect 非优先选择行为。
// 每次 Tick 从上一次返回结果为 Success 或者 Running 的行为开始执行，如果是第一次 Tick，则从左向右开始执行。
// 如果一个子行为返回 Success 或者 Running 时，会记录这个行为，下次直接从该行为开始执行。
type NonPrioritySelect struct {
	Composite
	lastSelectIndex int
}

func NewNonPrioritySelect(children ...Behavior) *NonPrioritySelect {
	var n = &NonPrioritySelect{}
	n.SetWorker(n)
	n.children = children
	return n
}

func (this *NonPrioritySelect) OnExec(ctx Context) Status {
	var index = this.lastSelectIndex
	var round = false // 标识是否跑完一轮
	for index < len(this.children) {
		var child = this.children[index]
		var status = child.Tick(ctx)

		if status == Failure {
			// 如果已经执行完一轮，并且已经是开始节点的前一个节点，则返回 Failure
			if round && index == this.lastSelectIndex-1 {
				return Failure
			}

			// 执行到节点列表的最后一个节点
			if index == len(this.children)-1 {
				// 如果开始节点是从头开始，表示一轮执行完成
				if this.lastSelectIndex == 0 {
					return Failure
				}

				// 否则一轮未执行完成，从节点列表的头部开始重新执行
				index = 0
				round = true
			} else {
				index++
			}
			continue
		}
		this.lastSelectIndex = index
		return status
	}
	this.lastSelectIndex = 0
	return Failure
}
