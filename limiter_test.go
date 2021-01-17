package behavior_test

import (
	"github.com/smartwalle/behavior"
	"testing"
)

func TestNewLimiter(t *testing.T) {
	var n = behavior.NewLimiter(3, NewSuccessAction())
	var status = n.Tick(nil)
	if status != behavior.Success {
		t.Fatal("期望结果为 Success")
	}

	status = n.Tick(nil)
	status = n.Tick(nil)

	if status != behavior.Success {
		t.Fatal("期望结果为 Success")
	}

	status = n.Tick(nil)

	if status != behavior.Failure {
		t.Fatal("期望结果为 Failure")
	}
}
