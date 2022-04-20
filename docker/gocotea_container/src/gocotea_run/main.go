package main

import (
	"fmt"

	"github.com/ispras/gocotea/src/gocotea"
)

func main() {
	pbPath := "main.yaml"
	inv := "inv"

	gocotea.InitPythonInterpretetor()

	var argMaker gocotea.ArgumentMaker

	argMaker.InitArgMaker()
	argMaker.AddArgument("-i", inv)

	var r gocotea.Runner

	r.InitRunner(&argMaker, pbPath)

	for r.HasNextPlay() {
		for r.HasNextTask() {
			r.RunNextTask()
		}
	}

	r.FinishAnsibleWork()

	gocotea.FinalizePythonInterpretetor()

	fmt.Println("END")
}
