package behavior

import "time"

// Wait 等待行为
type Wait struct {
	Action
	delay time.Duration
	start time.Time
}

func NewWait(delay time.Duration) *Wait {
	var n = &Wait{}
	n.SetWorker(n)
	n.delay = delay
	return n
}

func (this *Wait) OnOpen(ctx Context) {
	this.start = time.Now()
}

func (this *Wait) OnExec(ctx Context) Status {
	if time.Now().Before(this.start.Add(this.delay)) {
		return Running
	}
	return Success
}
