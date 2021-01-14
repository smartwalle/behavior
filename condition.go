package behavior

type ConditionFunc func(Context) bool

// Condition 条件行为
type Condition struct {
	base
	cond ConditionFunc
}

func NewCondition(cond ConditionFunc) *Condition {
	var n = &Condition{}
	n.cond = cond
	return n
}

func (this *Condition) Reset() {
}

func (this *Condition) Exec(ctx Context) Status {
	if this.cond != nil && this.cond(ctx) {
		return Success
	}
	return Failure
}
