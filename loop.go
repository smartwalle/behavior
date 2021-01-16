package behavior

import "fmt"

// Loop 循环执行指定行为。
// 每一次 Tick 只会执行一次子行为。
// 当子行为返回 Failure 或者 Error 的时候，返回 Failure 或者 Error，算作一次循环。
// 当子行为返回 Success 算作一次循环，如果不是最后一次循环，会返回 Running，以表示当前行为正在执行中，如果是最后一次循环，则返回 Success。
// 当子行为返回 Running 的时候，下次 Tick 会从上次中止的次数开始继续执行。
// 达到最大限定次数之后，返回子行为的状态，并重置已执行次数。
// count 参数小于 0 的时候，将不作重复次数限定。
type Loop struct {
	Decorator
	count   int
	current int
}

func NewLoop(count int, child Behavior) *Loop {
	var n = &Loop{}
	n.SetWorker(n)
	n.count = count
	n.child = child
	return n
}

func (this *Loop) OnExec(ctx Context) Status {
	if this.child == nil {
		return Error
	}

	var status = Failure
	if this.count <= 0 || this.current < this.count {
		fmt.Println(this.current)
		status = this.child.Tick(ctx)

		if status == Running {
			return status
		}

		this.current++

		if status == Failure || status == Error {
			return status
		}

		if status == Success && (this.count == 0 || this.current < this.count) {
			status = Running
			return status
		}
	}
	this.current = 0
	return status
}
