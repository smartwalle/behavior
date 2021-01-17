package behavior

import (
	"time"
)

// TimeAfter 限定子行为在指定时间之后执行。
// 子节点不会立即执行，而会在指定的时间到达后才开始执行，不限定执行次数。
// 未到时间执行返回 Running，否则返回子行为的执行结果。
type TimeAfter struct {
	Decorator
	startTime time.Time
}

func NewTimeAfter(t time.Time, child Behavior) *TimeAfter {
	var n = &TimeAfter{}
	n.SetWorker(n)
	n.startTime = t
	n.child = child
	return n
}

func (this *TimeAfter) OnExec(ctx Context) Status {
	if time.Now().Before(this.startTime) {
		return Running
	}
	return this.child.Tick(ctx)
}

// TimeBefore 限定子行为在指定时间之内执行。
// 在指定时间范围内不限制次数执行子行为，并返回子行为的执行结果，超过时间之后，返回 Failure。
type TimeBefore struct {
	Decorator
	endTime time.Time
}

func NewTimeBefore(t time.Time, child Behavior) *TimeBefore {
	var n = &TimeBefore{}
	n.SetWorker(n)
	n.endTime = t
	n.child = child
	return n
}

func (this *TimeBefore) OnExec(ctx Context) Status {
	if time.Now().Before(this.endTime) {
		return this.child.Tick(ctx)
	}
	return Failure
}

// TimeLimit 限定子行为最长运行时间。
// 指定子行为的最长运行(Running)时间，如果子行为在指定时间到达后还在运行则强制返回 Failure。
// 开始时间从开始执行本行为开始。
type TimeLimit struct {
	Decorator
	duration time.Duration
	endTime  time.Time
}

func NewTimeLimit(duration time.Duration, child Behavior) *TimeLimit {
	var n = &TimeLimit{}
	n.SetWorker(n)
	n.duration = duration
	n.child = child
	return n
}

func (this *TimeLimit) OnOpen(ctx Context) {
	this.endTime = time.Now().Add(this.duration)
}

func (this *TimeLimit) OnExec(ctx Context) Status {
	if time.Now().Before(this.endTime) {
		return this.child.Tick(ctx)
	}
	return Failure
}
