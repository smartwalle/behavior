package main

import (
	"github.com/smartwalle/behavior"
)

func main() {
	var seq = behavior.NewSequence(behavior.NewSequence())
	seq.Exec(nil)
	seq.Next()

	seq.Exec(nil)
	seq.Next()
}
