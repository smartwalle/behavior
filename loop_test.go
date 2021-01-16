package behavior_test

import (
	"github.com/smartwalle/behavior"
	"testing"
)

func TestNewLoop(t *testing.T) {
	var add = NewAddAction(false, behavior.Success)
	var n = behavior.NewLoop(3, add)
	var status = n.Tick(nil)
	if add.value != 1 && status != behavior.Running {
		t.Fatal("执行结果应该为 1，状态为 Running")
	}

	status = n.Tick(nil)
	if add.value != 2 && status != behavior.Running {
		t.Fatal("执行结果应该为 2，状态为 Running")
	}

	status = n.Tick(nil)
	if add.value != 3 && status != behavior.Success {
		t.Fatal("执行结果应该为 2，状态为 Success")
	}
}
