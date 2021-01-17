package behavior_test

import (
	"github.com/smartwalle/behavior"
	"testing"
)

func TestNewInverter(t *testing.T) {
	var n = behavior.NewInverter(NewSuccessAction())
	var status = n.Tick(nil)
	if status != behavior.Failure {
		t.Fatal("返回状态期望 Failure, 实际为", status)
	}

	n = behavior.NewInverter(NewFailureAction())
	status = n.Tick(nil)
	if status != behavior.Success {
		t.Fatal("返回状态期望 Success, 实际为", status)
	}

	n = behavior.NewInverter(NewRunningAction())
	status = n.Tick(nil)
	if status != behavior.Running {
		t.Fatal("返回状态期望 Running, 实际为", status)
	}
}
