package behavior

import (
	"math/rand"
	"time"
)

// PrioritySelector 优先选择行为。
// 每次 Tick 从左到右顺序执行所有子行为，如果一个子行为返回 Success 或者 Running，则返回 Success 或者 Running，类似逻辑或。
// 如果一个子行为返回 Running 时，会记录这个行为，下次直接从该行为开始执行。
type PrioritySelector struct {
	Composite
	lastRunningIndex int
}

func NewPrioritySelector(children ...Behavior) *PrioritySelector {
	var n = &PrioritySelector{}
	n.SetWorker(n)
	n.children = children
	return n
}

func (this *PrioritySelector) OnOpen(ctx Context) {
	this.lastRunningIndex = 0
}

func (this *PrioritySelector) OnExec(ctx Context) Status {
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

// NonPrioritySelector 非优先选择行为。
// 每次 Tick 从上一次返回结果为 Success 或者 Running 的行为开始执行，如果是第一次 Tick，则从左向右开始执行。
// 如果一个子行为返回 Success 或者 Running 时，会记录这个行为，下次直接从该行为开始执行。
type NonPrioritySelector struct {
	Composite
	lastSelectIndex int
}

func NewNonPrioritySelector(children ...Behavior) *NonPrioritySelector {
	var n = &NonPrioritySelector{}
	n.SetWorker(n)
	n.children = children
	return n
}

func (this *NonPrioritySelector) OnExec(ctx Context) Status {
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

// RandomSelector 随机选择行为。
// 选择之前，会对子行为进行随机排序，其它逻辑和 PrioritySelector 一致。
// 如果执行结果为 Success 或者 Failure，下一次执行之前会对子行为进行重新排序。
type RandomSelector struct {
	PrioritySelector
	r *rand.Rand
}

func NewRandomSelector(children ...Behavior) *RandomSelector {
	var n = &RandomSelector{}
	n.SetWorker(n)
	n.r = rand.New(rand.NewSource(time.Now().Unix()))
	n.children = children
	return n
}

func (this *RandomSelector) OnOpen(ctx Context) {
	this.PrioritySelector.OnOpen(ctx)
	this.rand()
}

func (this *RandomSelector) rand() {
	for i := len(this.children) - 1; i > 0; i-- {
		var rIdx = this.r.Intn(i)
		this.children[i], this.children[rIdx] = this.children[rIdx], this.children[i]
	}
}

// WeightBehavior 权重行为。
type WeightBehavior interface {
	Behavior

	// Weight 返回当前行为所占权重值
	Weight() int
}

// WeightSelector 权重选择行为。
// 每一个子行为都有一个权重值，每次 Tick 会根据权重值随机选择一个节点进行执行，并返回该节点的执行结果。
// 如果选中的子行为返回 Running 时，会记录这个行为，下次直接执行该行为。
//
// 随机选择的算法为：
// 1、计算出各子行为的总权重值，然后随机一个最大值不超过总权重值的数，这个数即为我们选中的值。
// 2、每次 Tick 从左到右顺序获取子行为的权重值，对子行为的权重值进行累加，当累加的结果大于等于随机选择的值，则当前的子行为则为选中的行为。
// 3、执行该子行为并返回其结果，如果结果为 Running，则会记录这个行为，下次 Tick 继续执行该行为。
//
// 与其它选择行为不一样的地方在于，本选择行为每次只会选择一个子行为执行，无论该子行为返回何种状态。
type WeightSelector struct {
	Composite
	lastRunningIndex int
	children         []WeightBehavior
	totalSum         int
	r                *rand.Rand
}

func NewWeightSelector(children ...WeightBehavior) *WeightSelector {
	var n = &WeightSelector{}
	n.SetWorker(n)
	n.children = children
	n.lastRunningIndex = -1
	for _, child := range children {
		var w = child.Weight()
		n.totalSum += w
	}
	n.r = rand.New(rand.NewSource(time.Now().Unix()))
	return n
}

func (this *WeightSelector) OnOpen(ctx Context) {
	this.lastRunningIndex = -1
}

func (this *WeightSelector) OnExec(ctx Context) Status {
	var chosen = this.r.Intn(this.totalSum)

	if this.lastRunningIndex > -1 {
		var child = this.children[this.lastRunningIndex]
		var status = child.Tick(ctx)
		return status
	}

	var sum = 0
	for i := 0; i < len(this.children); i++ {
		var child = this.children[i]
		var w = child.Weight()
		sum += w

		if w > 0 && sum >= chosen {
			var status = child.Tick(ctx)

			if status == Running {
				this.lastRunningIndex = i
			}
			return status
		}
	}
	return Failure
}
