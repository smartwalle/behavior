package behavior_test

import (
	"github.com/smartwalle/behavior"
	"testing"
	"time"
)

func TestNewTimeAfter(t *testing.T) {
	var nextTime = time.Now().Add(time.Second * 2)
	var n = behavior.NewTimeAfter(nextTime, NewSuccessAction())
	var status = n.Tick(nil)
	if status != behavior.Running {
		t.Fatal("未到时间，任务结果应该为 Running")
	}

	time.Sleep(time.Second * 2)
	status = n.Tick(nil)
	if status == behavior.Running {
		t.Fatal("已过时间，任务结果不应该为 Running")
	}
}

func TestNewTimeBefore(t *testing.T) {
	var endTime = time.Now().Add(time.Second * 2)
	var n = behavior.NewTimeBefore(endTime, NewSuccessAction())

	var status = n.Tick(nil)
	if status != behavior.Success {
		t.Fatal("时间范围内，任务结果应该为 Success")
	}

	time.Sleep(time.Second * 2)
	status = n.Tick(nil)
	if status != behavior.Failure {
		t.Fatal("已过时间，任务结果应该为 Failure")
	}
}

func TestNewTimeLimit(t *testing.T) {
	// 限定时间为 2 秒，子行为执行时间为 5 秒
	var n = behavior.NewTimeLimit(time.Second*2, behavior.NewWait(time.Second*5))

	var status = n.Tick(nil)
	if status != behavior.Running {
		t.Fatal("时间限定内，任务结果应该为 Running")
	}

	time.Sleep(time.Second * 3)
	// 超过了限定时间 2 秒，子行为还在执行，所以返回的状态应该 Failure
	status = n.Tick(nil)
	if status != behavior.Failure {
		t.Fatal("时间限定外，任务结果应该为 Failure")
	}
}
