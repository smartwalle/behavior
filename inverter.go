package behavior

// Inverter 倒置子行为结果。
// 对子节点的返回结果取“非”，即子节点返回 Success 则该节点返回 Failure，子节点返回 Failure 则该节点返回 Success。
type Inverter struct {
	Decorator
}

func NewInverter(child Behavior) *Inverter {
	var n = &Inverter{}
	n.SetWorker(n)
	n.child = child
	return n
}

func (this *Inverter) OnExec(ctx Context) Status {
	var status = this.child.Tick(ctx)
	if status == Success {
		return Failure
	} else if status == Failure {
		return Success
	}
	return status
}
