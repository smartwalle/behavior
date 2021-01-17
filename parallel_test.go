package behavior_test

import (
	"github.com/smartwalle/behavior"
	"testing"
)

func TestNewParallel(t *testing.T) {
	// 运行状态
	var n = behavior.NewParallel(
		behavior.SuccessPolicyOnAll,
		behavior.FailurePolicyOnOne,
		NewSuccessAction(),
		NewFailureAction(),
		NewRunningAction(),
	)
	var status = n.Tick(nil)
	if status != behavior.Running {
		t.Fatal("期望状态为 Running， 实际状态为", status)
	}

	// 失败状态
	n = behavior.NewParallel(
		behavior.SuccessPolicyOnAll,
		behavior.FailurePolicyOnOne,
		NewSuccessAction(),
		NewFailureAction(),
		NewSuccessAction(),
	)
	status = n.Tick(nil)
	if status != behavior.Failure {
		t.Fatal("期望状态为 Failure， 实际状态为", status)
	}

	// 成功状态
	n = behavior.NewParallel(
		behavior.SuccessPolicyOnAll,
		behavior.FailurePolicyOnOne,
		NewSuccessAction(),
		NewSuccessAction(),
		NewSuccessAction(),
	)
	status = n.Tick(nil)
	if status != behavior.Success {
		t.Fatal("期望状态为 Failure， 实际状态为", status)
	}
}
