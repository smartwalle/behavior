package behavior

import (
	"time"
)

// Delay 延迟执行子行为。
// 等待一定的时间之后执行子行为并返回子行为的执行结果，否则返回 Running。
type Delay struct {
	Decorator
	duration time.Duration
	start    time.Time
}

func NewDelay(duration time.Duration, child Behavior) *Delay {
	var n = &Delay{}
	n.SetWorker(n)
	n.duration = duration
	n.child = child
	return n
}

func (this *Delay) OnOpen(ctx Context) {
	this.start = time.Now()
}

func (this *Delay) OnExec(ctx Context) Status {
	if time.Now().Before(this.start.Add(this.duration)) {
		return Running
	}
	return this.child.Tick(ctx)
}