# gocotea

#### (COntrol Thread Execution Ansible)

### gocotea is:
Tool that provides Golang API to run Ansible programmatically. *gocotea* is a port of [cotea](https://github.com/ispras/cotea) into Golang. Porting was done using the [gopython](https://github.com/ispras/gopython) tool.

### gocotea allows:
- **To control** Ansible execution by iterating over the Ansible plays and tasks
- **To embed** Ansible into another system
- **To debug** Ansible execution by getting the values of Ansible variables and by retrieving the results of the execution of Ansible tasks/plays

## Installation
It is assumed that Ansible is installed in the current environment. The dependencies of *gocotea* are [cotea](https://github.com/ispras/cotea) and [gopython](https://github.com/ispras/gopython). 

[gopython](https://github.com/ispras/gopython) is a Golang package and it will be installed automatically after *gocotea* installation:
```bash
go get github.com/ispras/gocotea
```

You also have to set PYTHONPATH env variable in the enviroment where you want to use *gocotea*:
```bash
export PYTHONPATH=$GOPATH/src/github.com/ispras/gocotea/src
```

PYTHONPATH variable is used by [gopython](https://github.com/ispras/gopython) for embedding of [cotea](https://github.com/ispras/cotea). This is beacause *gocotea* is based on the embedding of [cotea](https://github.com/ispras/cotea) into Golang. The embedding is done by using the source code of [cotea](https://github.com/ispras/cotea). At the moment there is no support of *pip* installation for [cotea](https://github.com/ispras/cotea), so in the *src/* folder there is a *cotea/* folder with cotea's source code (and the PYTHONPATH variable points just there). 

Include *gocotea* into your code after all these steps:
```Golang
import gocotea "github.com/ispras/gocotea/src"
```

## Quick start
```Golang
package main

import (
	"fmt"

	gocotea "github.com/ispras/gocotea/src"
)

func main() {
	inventory := "/path/to/inventory"
	playbookPath := "/path/to/playbook"

	gocotea.InitPythonInterpretetor()

	var r gocotea.Runner
	var argMaker gocotea.ArgumentMaker

	argMaker.InitArgMaker()
	argMaker.AddArgument("-i", inventory)

	r.InitRunner(&argMaker, playbookPath, "INFO", "/path/to/ansible/log/file")

	for r.HasNextPlay() {
		setupOk := r.SetupPlayForRun()
		fmt.Printf("Play: %s\n", r.GetCurrentPlayName())

		if setupOk {
			for r.HasNextTask() {
				fmt.Printf("\tTask: %s\n", r.GetNextTaskName())
				r.RunNextTask()
			}
		}
	}

	r.FinishAnsibleWork()

	if r.WasError() {
		fmt.Printf("Ansible failed. Error:\n%s\n", r.GetErrorMsg())
	}

	gocotea.FinalizePythonInterpretetor()
}

```
Any argument of the "ansible-playbook" command can be passed by using **ArgumentMaker** object.
The launch and control of the Ansible is carried out using the **Runner** object.

The Ansible output will be forwarded special log file. Path to this file is passed to *InitRunner* method of the Runner structure. 

A detailed overview of all interfaces is provided in [gocotea documentation](https://github.com/ispras/gocotea/blob/main/docs/gocotea_docs.md).