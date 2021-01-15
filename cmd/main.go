package main

import (
	"fmt"
	"github.com/smartwalle/behavior"
)

func main() {
	behavior.IF(func(ctx behavior.Context) bool {
		if ctx.Target() == nil {
			return false
		}
		return true
	}, NewPrintAction()).Exec(behavior.NewContext("hha"))
}

type PrintAction struct {
	behavior.Action
}

func NewPrintAction() *PrintAction {
	var a = &PrintAction{}
	a.SetWorker(a)
	return a
}

func (this *PrintAction) OnExec(ctx behavior.Context) behavior.Status {
	fmt.Println(ctx.Target())
	return behavior.Success
}
