package behavior

import "time"

// Delay 延迟执行子行为。
// 等待一定的时间之后执行子行为并返回子行为的执行结果，否则返回 Running。
// 子行为返回 Running 的时候，不会重新计算执行时间，下一次 Tick 会继续执行子行为。
// 子行为返回 Success 或者 Failure 的时候，下一次开始执行子行为之前，会重新计算执行时间。
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
