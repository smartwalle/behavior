package behavior_test

import (
	"github.com/smartwalle/behavior"
	"testing"
)

func TestNewUntilFailure(t *testing.T) {
	var n = behavior.NewUntilFailure(3, NewSuccessAction())

	var status = n.Tick(nil)
	if status != behavior.Running {
		t.Fatal("期望结果为 Running")
	}

	status = n.Tick(nil)
	if status != behavior.Running {
		t.Fatal("期望结果为 Running")
	}

	status = n.Tick(nil)
	if status != behavior.Failure {
		t.Fatal("期望结果为 Failure")
	}
}
