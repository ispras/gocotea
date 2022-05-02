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
			fmt.Println("Next task name: ", r.GetNextTaskName())

			taskResults := r.RunNextTask()
			if len(taskResults) > 0 {
				fmt.Println("Task IsChanged:", taskResults[0].IsChanged)
			}
		}
	}

	r.FinishAnsibleWork()

	if r.WasError() {
		fmt.Printf("Ansible failed. Error:\n%s\n", r.GetErrorMsg())
	}

	gocotea.FinalizePythonInterpretetor()
}
