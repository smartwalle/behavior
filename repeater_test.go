package behavior_test

import (
	"github.com/smartwalle/behavior"
	"testing"
)

func TestNewRepeater(t *testing.T) {
	// Tick 两次，每次重复两次，不重置结果
	var add = NewAddAction(false, behavior.Success)
	var n = behavior.NewRepeater(2, add)

	var status = n.Tick(nil)
	if add.value != 2 && status != behavior.Success {
		t.Fatal("执行结果应该为 2，状态应该为 Success")
	}

	status = n.Tick(nil)
	if add.value != 4 && status != behavior.Success {
		t.Fatal("执行结果应该为 4，状态应该为 Success")
	}

	// Tick 两次，每次重复两次，重置子行为的结果
	// AddAction 的 OnExec 每次都返回 Success，所以其 OnOpen 在每一次 Tick 之前都会被调用
	// 由于其 reset 为 true，所以其 value 值每次执行之前都会被重置为 0
	add = NewAddAction(true, behavior.Success)
	n = behavior.NewRepeater(2, add)

	status = n.Tick(nil)
	if add.value != 1 && status != behavior.Success {
		t.Fatal("执行结果应该为 1，状态应该为 Success")
	}

	status = n.Tick(nil)
	if add.value != 1 && status != behavior.Success {
		t.Fatal("执行结果应该为 1，状态应该为 Success")
	}

	// Tick 两次，由于子行为返回的是 Running，所以每次 Tick 只会重复一次
	// 虽然设置了子行为要重置结果，但是由于是在子行为的 OnOpen 回调中并且子行为状态为 Running, 所以不会重置结果
	add = NewAddAction(true, behavior.Running)
	n = behavior.NewRepeater(2, add)
	status = n.Tick(nil)
	if add.value != 1 && status != behavior.Running {
		t.Fatal("执行结果应该为 1，状态应该为 Running")
	}

	status = n.Tick(nil)
	if add.value != 2 && status != behavior.Running {
		t.Fatal("执行结果应该为 1，状态应该为 Running")
	}
}
