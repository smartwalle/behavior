package main

import (
	"fmt"
	"github.com/smartwalle/behavior"
	"time"
)

func main() {

	var node = behavior.IF(func(ctx behavior.Context) bool {
		if ctx.Target() == nil {
			return false
		}
		return true
	}, behavior.NewSequence(
		behavior.NewWait(time.Second*10),
		NewPrintAction(),
	))

	for {
		fmt.Println(node.Exec(behavior.NewContext("hha")))
		time.Sleep(time.Second)
	}

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
