package behavior_test

import (
	"fmt"
	"github.com/smartwalle/behavior"
	"testing"
)

func TestNewPrioritySelector(t *testing.T) {
	// 执行第一个子行为
	var ctx = &SelectContext{}
	ctx.Value = "2"
	var n = behavior.NewPrioritySelector(
		NewSelectA(1, "1"),
		NewSelectA(2, "2"),
	)
	n.Tick(ctx)
	if ctx.ExecId != 2 {
		t.Fatal("期望值为 2， 实际值为", ctx.ExecId)
	}
	n.Tick(ctx)
	if ctx.ExecId != 2 {
		t.Fatal("期望值为 2， 实际值为", ctx.ExecId)
	}
}

type SelectContext struct {
	ExecId int
	Value  string
}

type SelectA struct {
	behavior.Action
	id    int
	value string
}

func NewSelectA(id int, value string) *SelectA {
	var n = &SelectA{}
	n.SetWorker(n)
	n.id = id
	n.value = value
	return n
}

func (this *SelectA) OnExec(ctx behavior.Context) behavior.Status {
	var nCtx = ctx.(*SelectContext)
	if nCtx.Value == this.value {
		nCtx.ExecId = this.id
		return behavior.Success
	}
	return behavior.Failure
}

func TestNewNonPrioritySelector(t *testing.T) {
	var ctx = &SelectContext{}
	ctx.Value = "4"
	var n = behavior.NewNonPrioritySelector(
		NewSelectA(1, "1"),
		NewSelectA(2, "2"),
		NewSelectA(3, "3"),
		NewSelectA(4, "4"),
		NewSelectA(5, "5"),
	)
	n.Tick(ctx)
	if ctx.ExecId != 4 {
		t.Fatal("期望值为 4， 实际值为", ctx.ExecId)
	}
}

func TestNewWeightSelector(t *testing.T) {
	var ctx = &SelectContext{}
	var n = behavior.NewWeightSelector(
		NewWeightBehavior(10, "1"),
		NewWeightBehavior(20, "2"),
		NewWeightBehavior(30, "3"),
		NewWeightBehavior(40, "4"),
		NewWeightBehavior(50, "5"),
	)
	n.Tick(ctx)
}

type WeightBehavior struct {
	behavior.Action
	weight int
	value  string
}

func NewWeightBehavior(w int, value string) *WeightBehavior {
	var n = &WeightBehavior{}
	n.SetWorker(n)
	n.weight = w
	n.value = value
	return n
}

func (this *WeightBehavior) Weight() int {
	return this.weight
}

func (this *WeightBehavior) OnExec(ctx behavior.Context) behavior.Status {
	fmt.Println(this.value)
	return behavior.Success
}
