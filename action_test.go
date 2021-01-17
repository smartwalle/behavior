package behavior_test

import (
	"fmt"
	"github.com/smartwalle/behavior"
)

// PrintAction
type PrintAction struct {
	behavior.Action
	m string
}

func NewPrintAction(m string) *PrintAction {
	var n = &PrintAction{}
	n.SetWorker(n)
	n.m = m
	return n
}

func (this *PrintAction) OnExec(ctx behavior.Context) behavior.Status {
	fmt.Println(this.m)
	return behavior.Success
}

// SuccessAction
type SuccessAction struct {
	behavior.Action
}

func NewSuccessAction() *SuccessAction {
	var n = &SuccessAction{}
	n.SetWorker(n)
	return n
}

func (this *SuccessAction) OnExec(ctx behavior.Context) behavior.Status {
	return behavior.Success
}

// FailureAction
type FailureAction struct {
	behavior.Action
}

func NewFailureAction() *FailureAction {
	var n = &FailureAction{}
	n.SetWorker(n)
	return n
}

func (this *FailureAction) OnTick(ctx behavior.Context) behavior.Status {
	return behavior.Failure
}

// AddAction
type AddAction struct {
	behavior.Action
	value  int
	reset  bool
	status behavior.Status
}

func NewAddAction(reset bool, status behavior.Status) *AddAction {
	var n = &AddAction{}
	n.SetWorker(n)
	n.value = 0
	n.reset = reset
	n.status = status
	return n
}

func (this *AddAction) OnOpen(ctx behavior.Context) {
	if this.reset {
		this.value = 0
	}
}

func (this *AddAction) OnExec(ctx behavior.Context) behavior.Status {
	this.value += 1
	return this.status
}
