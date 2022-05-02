# gocotea documentation

The [Runner](https://github.com/ispras/gocotea/blob/main/src/runner.go#L10) structure is the main structure in *gocotea*. It wraps *ansible-playbook* command. With the use of *Runner* structure, user runs and controls Ansible programmatically, gets additional information about the execution.

The second point of interaction is the [ArgumentMaker](https://github.com/ispras/gocotea/blob/main/src/argument_maker.go#L7) structure. With the use of it, user can pass any argument for Ansible launch just like he passed arguments for *ansible-playbook* in the command line.

## Python interpretetor
*gocotea* is based on the embedding of [cotea](https://github.com/ispras/cotea) into Golang. The embedding is done using [gopython](https://github.com/ispras/gopython). *gopython* is based on CPython API calls and the main one's are [Py_Initialize](https://docs.python.org/3/c-api/init.html#c.Py_Initialize) and [Py_Finalize](https://docs.python.org/3/c-api/init.html#c.Py_Finalize). This methods have to be called for the correct operation of *gocotea*. At the moment one should call them explicitly: [InitPythonInterpretetor](https://github.com/ispras/gocotea/blob/main/src/python_interpretetor.go#L7) and [FinalizePythonInterpretetor](https://github.com/ispras/gocotea/blob/main/src/python_interpretetor.go#L12). In the future, it is planned to hide calls to these methods from users.

## ArgumentMaker

Firstoffall one should call [InitArgMaker()](https://github.com/ispras/gocotea/blob/main/src/argument_maker.go#L12) method to init ArgumentMaker object properly. The only interface of this object is *AddArgument* method, that set the argument for *ansible-playbook* command just like one do it in the shell.

**AddArgument(args ...string) error**
- *args* can be 1 string or 2 strings. In the case of one string it is a equivalent of non-value argument - for example, -v. 2 strings stand for an argument and its value. For example, *-i, /path/to/inventory* or *--extra-vars, {extravars...}*

Usage examples:
```Golang
# python interpretetor initialization
gocotea.InitPythonInterpretetor()

var argMaker gocotea.ArgumentMaker

// object initialization
argMaker.InitArgMaker()

// without value
argMaker.AddArgument("-vvv")

// with value
inventory = "/path/to/inventory"
argMaker.AddArgument("-i", inventory)

```

After all of the needed actions, *ArgumentMaker* object should be passed to runner structure *InitRunner* method.


## Runner

**InitRunner(argmaker *ArgumentMaker, pbPath) error***
- *argmaker* - object of *argument_maker* class
- *pbPath* - path of the playbook .yaml file
### controlling interfaces

**HasNextPlay() bool**

Checks if there is unexecuted *plays* in current Ansible execution. Returns *true* if there is.

**HasNextTask() bool**

Checks if there is unexecuted *tasks* in currently executing *play*. Returns *true* if there is.

**RunNextTask() []TaskResult**

Runs the next *task* in the currently executing *play*. Returns a slice of [TaskResult](https://github.com/ispras/gocotea/blob/main/docs/gocotea_docs.md#taskresult) structure objects that describe the result of the executed task on each host in the current group of hosts. 

**FinishAnsibleWork() bool**

Starts a bunch of actions that are needed to finish the current Ansible execution.

These four interfaces are the main part of *gocotea*. They let one control the execution of *ansible-playbook* launch. Every usage of gocotea will contain them in the following order:
```python
var r gocotea.Runner
# r = InitRunner(...)

for r.HasNextPlay() {
    for r.HasNextTask() {
        r.RunNextTask()
    }
}

r.FinishAnsibleWork()
```

### debugging interfaces
**WasError() bool**

Returns *true* if Ansible execution ends with an error.

**GetErrorMsg() string**

If Ansible execution ends with an error (*was_error* returns *true*), returns error message.

**GetCurrentPlayName() string**

Returns the current play name.

**GetNextTaskName() string**

Returns the next task name.

**GetPrevTaskName() string**

Returns the previous task name.

### TaskResult
This structure stores the task results in a convenient way. Based on [ansible.executor.task_result.TaskResult](https://github.com/ansible/ansible/blob/devel/lib/ansible/executor/task_result.py#L25).

Fields:
- *TaskName* - the name of the task
- *TaskStdOut* - the stdout of the task
- *TaskStdErr* - the stderr of the task
- *TaskMsg* - Ansible's msg about the task
-  *IsChanged* - True if task is "changed"
-  *IsFailed* - True if task was failed
-  *IsSkipped* - True if task was skipped
-  *IsUnreachable* - True if "unreachable"
