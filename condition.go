package behavior

type ConditionFunc func(Context) bool

// Condition 条件行为
type Condition struct {
	baseBehavior
	cond ConditionFunc
}

func NewCondition(cond ConditionFunc) *Condition {
	var n = &Condition{}
	n.cond = cond
	return n
}

func (this *Condition) Exec(ctx Context) {
	if this.cond(ctx) {
		this.status = Success
	} else {
		this.status = Failure
	}
}
