package behavior

func IF(cond ConditionFunc, child Behavior) Behavior {
	return NewSequence(NewCondition(cond), child)
}
