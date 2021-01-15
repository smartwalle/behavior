package behavior

type ConditionFunc func(Context) bool

// Condition 条件行为
type Condition struct {
	Action
	cond ConditionFunc
}

func NewCondition(cond ConditionFunc) *Condition {
	var n = &Condition{}
	n.SetWorker(n)
	n.cond = cond
	return n
}

func (this *Condition) OnExec(ctx Context) Status {
	if this.cond != nil && this.cond(ctx) {
		return Success
	}
	return Failure
}
