# gocotea

#### (cotea - COntrol Thread Execution Ansible)

### gocotea is:
Tool that provides Golang API to run Ansible programmatically. *gocotea* is a port of [cotea](https://github.com/ispras/cotea) into Golang. Porting was done using the [gopython](https://github.com/ispras/gopython) tool.

### gocotea allows:
- **To control** Ansible execution by iterating over the Ansible plays and tasks
- **To embed** Ansible into another system
- **To debug** Ansible execution by getting the values of Ansible variables and by retrieving the results of the execution of Ansible tasks/plays

## Installation
Tested on ubuntu 20.04 with golang 1.18 and python 3.8.10. cotea 1.3.3 and gopython 0.2.1 are required.

1. Install ansible:
```bash
pip install ansible==2.9.4
```

2. Install cotea:
```bash
pip install -i https://test.pypi.org/simple/ cotea==1.3.3
```

3. Create go module:
```bash
go mod init PREFERED_NAME
```

4. Include gocotea to your code (located in the created module) with this import:
```Golang
import "github.com/ispras/gocotea/src/gocotea"
```

5. Make go mod tidy. This command will download required golang packages (including gocotea and gopython)
```bash
go mod tidy
```

Creating a go module is necessary for correct installation of gocotea.

## Quick start
```Golang
package main

import (
	"fmt"

	"github.com/ispras/gocotea/src/gocotea"
)

func main() {
	pbPath := "/path/to/playbook"
	inv := "/path/to/inventory"

	gocotea.InitPythonInterpretetor()

	var argMaker gocotea.ArgumentMaker

	argMaker.InitArgMaker()
	argMaker.AddArgument("-i", inv)

	var r gocotea.Runner

	r.InitRunner(&argMaker, pbPath)

    for r.HasNextPlay() {
        for r.HasNextTask() {
		    fmt.Println("Next task name: ", r.GetNextTaskName())
		    
            r.RunNextTask()
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

A detailed overview of all interfaces is provided in [gocotea documentation](https://github.com/ispras/gocotea/blob/main/docs/gocotea_docs.md).
