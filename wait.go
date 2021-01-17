package behavior

import "time"

// Wait 等待行为。
// 类似于 Sleep，等待一定的时间之后，返回 Success，否则返回 Running
type Wait struct {
	Action
	duration time.Duration
	nextTime time.Time
}

func NewWait(duration time.Duration) *Wait {
	var n = &Wait{}
	n.SetWorker(n)
	n.duration = duration
	return n
}

func (this *Wait) OnOpen(ctx Context) {
	this.nextTime = time.Now().Add(this.duration)
}

func (this *Wait) OnExec(ctx Context) Status {
	if time.Now().Before(this.nextTime) {
		return Running
	}
	return Success
}
