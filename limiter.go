package behavior

// Limiter 限定指定行为的执行次数
// 限定子行为的执行次数，达到限定次数之后，返回 Failure，否则返回子行为的执行结果
// 子行为返回 Success 或者 Failure，算作一次执行
// count 参数必须大于 0
type Limiter struct {
	Decorator
	limit   int
	current int
}

func NewLimiter(limit int, child Behavior) *Limiter {
	var n = &Limiter{}
	n.SetWorker(n)
	n.limit = limit
	n.child = child
	return n
}

func (this *Limiter) OnExec(ctx Context) Status {
	if this.child == nil || this.limit <= 0 {
		return Error
	}

	if this.current < this.limit {
		var status = this.child.Exec(ctx)
		if status == Success || status == Failure {
			this.current++
		}
		return status
	}

	return Failure
}
