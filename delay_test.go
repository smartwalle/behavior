package behavior_test

import (
	"github.com/smartwalle/behavior"
	"testing"
	"time"
)

func TestNewDelay(t *testing.T) {
	var n = behavior.NewDelay(time.Second, NewSuccessAction())
	var status = n.Tick(nil)
	if status != behavior.Running {
		t.Fatal("未到时间，延时任务结果应该为 Running")
	}
	time.Sleep(time.Second * 1)
	status = n.Tick(nil)
	if status == behavior.Running {
		t.Fatal("已过时间，延时任务结果不应该为 Running")
	}

	// 第二轮延时开始，此次 Tick 会重新计算延时时间
	status = n.Tick(nil)
	if status != behavior.Running {
		t.Fatal("未到时间，延时任务结果应该为 Running")
	}
	time.Sleep(time.Second * 1)
	status = n.Tick(nil)
	if status == behavior.Running {
		t.Fatal("已过时间，延时任务结果不应该为 Running")
	}

}
