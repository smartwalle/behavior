package behavior

import "time"

// Timer 定时执行子行为。
// 子节点不会立即执行，而会在指定的时间到达后才开始执行。
// 未到时间执行返回 Running，否则返回子行为的执行结果。
type Timer struct {
	Decorator
	nextTime time.Time
}

func NewTimer(t time.Time, child Behavior) *Timer {
	var n = &Timer{}
	n.SetWorker(n)
	n.nextTime = t
	n.child = child
	return n
}

func (this *Timer) OnExec(ctx Context) Status {
	if time.Now().Before(this.nextTime) {
		return Running
	}
	return this.child.Tick(ctx)
}
