package behavior_test

import (
	"github.com/smartwalle/behavior"
	"testing"
)

func TestNewLoop(t *testing.T) {
	// 不重置结果，每次 Tick 执行一次循环
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

	// 执行完一轮循环之后，又开始下一轮循环
	// 由于不重围结果，所以值会继续递增，状态应该为 Running
	status = n.Tick(nil)
	if add.value != 4 && status != behavior.Running {
		t.Fatal("执行结果应该为 2，状态为 Running")
	}
}
