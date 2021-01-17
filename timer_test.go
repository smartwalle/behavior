package behavior_test

import (
	"github.com/smartwalle/behavior"
	"testing"
	"time"
)

func TestNewTimer(t *testing.T) {
	var nextTime = time.Now().Add(time.Second * 2)
	var n = behavior.NewTimer(nextTime, NewSuccessAction())
	var status = n.Tick(nil)
	if status != behavior.Running {
		t.Fatal("未到时间，定时任务结果应该为 Running")
	}

	time.Sleep(time.Second * 2)
	status = n.Tick(nil)
	if status == behavior.Running {
		t.Fatal("已过时间，定时任务结果不应该为 Running")
	}
}
