package behavior

import "time"

// Delay 延迟执行子行为。
// 等待一定的时间之后执行子行为并返回子行为的执行结果，否则返回 Running。
// 每一次执行完子行为之后，会重新计算下一次执行时间。
type Delay struct {
	Decorator
	duration time.Duration
	nextTime time.Time
}

func NewDelay(duration time.Duration, child Behavior) *Delay {
	var n = &Delay{}
	n.SetWorker(n)
	n.duration = duration
	n.child = child
	return n
}

func (this *Delay) OnOpen(ctx Context) {
	this.nextTime = time.Now().Add(this.duration)
}

func (this *Delay) OnExec(ctx Context) Status {
	if time.Now().Before(this.nextTime) {
		return Running
	}
	return this.child.Tick(ctx)
}
