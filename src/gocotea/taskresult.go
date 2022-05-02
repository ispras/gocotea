package gocotea

import (
	gopython "github.com/ispras/gopython/src/gopython"
)

type TaskResult struct {
	TaskName      string
	TaskStdOut    string
	TaskStdErr    string
	TaskMsg       string
	IsChanged     bool
	IsFailed      bool
	IsSkipped     bool
	IsUnreachable bool
}

func MakeTaskResFromPyObj(taskResPy *gopython.PythonObject) (*TaskResult, error) {
	res := TaskResult{}

	taskNamePy, err := taskResPy.GetAttr("task_name")
	if err != nil {
		return nil, err
	}

	taskName, err := taskNamePy.ToStandartGoType()
	if err != nil {
		return nil, err
	}

	res.TaskName = taskName.(string)

	taskStdOutPy, err := taskResPy.GetAttr("stdout")
	if err != nil {
		return nil, err
	}

	taskStdOut, err := taskStdOutPy.ToStandartGoType()
	if err != nil {
		return nil, err
	}

	res.TaskStdOut = taskStdOut.(string)

	taskStdErrPy, err := taskResPy.GetAttr("stderr")
	if err != nil {
		return nil, err
	}

	taskStdErr, err := taskStdErrPy.ToStandartGoType()
	if err != nil {
		return nil, err
	}

	res.TaskStdErr = taskStdErr.(string)

	taskMsgPy, err := taskResPy.GetAttr("msg")
	if err != nil {
		return nil, err
	}

	taskMsg, err := taskMsgPy.ToStandartGoType()
	if err != nil {
		return nil, err
	}

	res.TaskMsg = taskMsg.(string)

	isChangedPy, err := taskResPy.GetAttr("is_changed")
	if err != nil {
		return nil, err
	}

	isChanged, err := isChangedPy.ToStandartGoType()
	if err != nil {
		return nil, err
	}

	res.IsChanged = isChanged.(bool)

	isFailedPy, err := taskResPy.GetAttr("is_failed")
	if err != nil {
		return nil, err
	}

	isFailed, err := isFailedPy.ToStandartGoType()
	if err != nil {
		return nil, err
	}

	res.IsFailed = isFailed.(bool)

	isSkippedPy, err := taskResPy.GetAttr("is_skipped")
	if err != nil {
		return nil, err
	}

	isSkipped, err := isSkippedPy.ToStandartGoType()
	if err != nil {
		return nil, err
	}

	res.IsSkipped = isSkipped.(bool)

	isUnreachablePy, err := taskResPy.GetAttr("is_unreachable")
	if err != nil {
		return nil, err
	}

	isUnreachable, err := isUnreachablePy.ToStandartGoType()
	if err != nil {
		return nil, err
	}

	res.IsUnreachable = isUnreachable.(bool)

	return &res, nil
}
