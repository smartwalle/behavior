package behavior

type FailurePolicy int8

const (
	FailurePolicyOnOne FailurePolicy = 1
	FailurePolicyOnAll FailurePolicy = 2
)

type SuccessPolicy int8

const (
	SuccessPolicyOnOne SuccessPolicy = 1
	SuccessPolicyOnAll SuccessPolicy = 2
)

// Parallel 并行行为。
type Parallel struct {
	Composite
	failurePolicy FailurePolicy
	successPolicy SuccessPolicy

	completed  []bool // 记录已经完成状态
	success    bool   // 有成功
	failure    bool   // 有失败
	running    bool   // 有运行中
	allSuccess bool   // 全部成功
	allFailure bool   // 全部失败
}

func NewParallel(sPolicy SuccessPolicy, fPolicy FailurePolicy, children ...Behavior) *Parallel {
	var n = &Parallel{}
	n.SetWorker(n)
	n.successPolicy = sPolicy
	n.failurePolicy = fPolicy
	n.children = children
	return n
}

func (this *Parallel) OnOpen(ctx Context) {
	this.completed = make([]bool, len(this.children))

	this.success = false
	this.failure = false
	this.running = false
	this.allSuccess = true
	this.allFailure = true
}

func (this *Parallel) OnExec(ctx Context) Status {
	for i := 0; i < len(this.children); i++ {
		if this.completed[i] {
			continue
		}

		var child = this.children[i]
		var status = child.Tick(ctx)
		switch status {
		case Running:
			this.running = true
			this.allFailure = false
			this.allSuccess = false
		case Failure:
			this.failure = true
			this.allSuccess = false
			this.completed[i] = true
		case Success:
			this.success = true
			this.allFailure = false
			this.completed[i] = true
		}
	}

	if this.running {
		return Running
	}
	if (this.failurePolicy == FailurePolicyOnAll && this.allFailure) || (this.failurePolicy == FailurePolicyOnOne && this.failure) {
		return Failure
	}
	if (this.successPolicy == SuccessPolicyOnAll && this.allSuccess) || (this.successPolicy == SuccessPolicyOnOne && this.success) {
		return Success
	}
	return Failure
}
