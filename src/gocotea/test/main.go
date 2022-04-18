package main

import (
	"fmt"
	"time"

	"github.com/ispras/gocotea/src/gocotea"
)

func main() {
	start := time.Now()
	pbPath := "main_sqlserver.yaml"
	inv := "inv_root"

	gocotea.InitPythonInterpretetor()

	var argMaker gocotea.ArgumentMaker

	argMaker.InitArgMaker()
	argMaker.AddArgument("-i", inv)

	var r gocotea.Runner

	r.InitRunner(&argMaker, pbPath, "DEBUG", "log.log")

	//fmt.Println(r)

	for r.HasNextPlay() {
		for r.HasNextTask() {
			r.RunNextTask()
		}
	}

	r.FinishAnsibleWork()

	// if r.WasError() {
	// 	fmt.Println("THERE WAS AN MATHAFAKA ERROR")
	// }

	gocotea.FinalizePythonInterpretetor()

	duration := time.Since(start)

	fmt.Println("TIME = ", duration)

	fmt.Println("END")
}
