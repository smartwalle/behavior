package behavior

// Invert 倒置行为。
// 对子节点的返回结果取“非”，即子节点返回 Success 则该节点返回 Failure，子节点返回 Failure 则该节点返回 Success。
type Invert struct {
	Decorator
}

func NewInvert(child Behavior) *Invert {
	var n = &Invert{}
	n.SetWorker(n)
	n.child = child
	return n
}

func (this *Invert) OnExec(ctx Context) Status {
	var status = this.child.Tick(ctx)
	if status == Success {
		return Failure
	} else if status == Failure {
		return Success
	}
	return status
}
