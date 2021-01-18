package behavior_test

import (
	"github.com/smartwalle/behavior"
	"testing"
)

func TestNewUntilSuccess(t *testing.T) {
	var n = behavior.NewUntilSuccess(3, NewFailureAction())

	var status = n.Tick(nil)
	if status != behavior.Running {
		t.Fatal("期望结果为 Running")
	}

	status = n.Tick(nil)
	if status != behavior.Running {
		t.Fatal("期望结果为 Running")
	}

	status = n.Tick(nil)
	if status != behavior.Success {
		t.Fatal("期望结果为 Failure")
	}
}
