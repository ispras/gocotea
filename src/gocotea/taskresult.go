package gocotea

import (
	"fmt"

	gopython "github.com/ispras/gopython/src/gopython"
)

type TaskResult struct {
	TaskName      string
	TaskStdOut    string
	IsChanged     bool
	IsFailed      bool
	IsSkipped     bool
	IsUnreachable bool
}

func MakeTaskResFromPyObj(taskResPy *gopython.PythonObject) (*TaskResult, error) {
	res := TaskResult{}

	hasTaskName, err := taskResPy.HasAttr("task_name")
	if err != nil {
		return nil, err
	}

	if hasTaskName {
		fmt.Println("P1")
		taskNamePy, err := taskResPy.GetAttr("task_name")
		if err != nil {
			return nil, err
		}

		taskName, err := taskNamePy.ToStandartGoType()
		if err != nil {
			return nil, err
		}

		res.TaskName = taskName.(string)
	}

	hasIsChanged, err := taskResPy.HasAttr("is_changed")
	if err != nil {
		return nil, err
	}

	if hasIsChanged {
		fmt.Println("P2")
		isChangedPy, err := taskResPy.GetAttr("is_changed")
		if err != nil {
			return nil, err
		}

		isChanged, err := isChangedPy.ToStandartGoType()
		if err != nil {
			return nil, err
		}

		res.IsChanged = isChanged.(bool)
	}

	hasIsFailed, err := taskResPy.HasAttr("is_failed")
	if err != nil {
		return nil, err
	}

	if hasIsFailed {
		fmt.Println("P3")
		isFailedPy, err := taskResPy.GetAttr("is_failed")
		if err != nil {
			return nil, err
		}

		isFailed, err := isFailedPy.ToStandartGoType()
		if err != nil {
			return nil, err
		}

		res.IsFailed = isFailed.(bool)
	}

	hasIsSkipped, err := taskResPy.HasAttr("is_skipped")
	if err != nil {
		return nil, err
	}

	if hasIsSkipped {
		fmt.Println("P4")
		isSkippedPy, err := taskResPy.GetAttr("is_skipped")
		if err != nil {
			return nil, err
		}

		isSkipped, err := isSkippedPy.ToStandartGoType()
		if err != nil {
			return nil, err
		}

		res.IsSkipped = isSkipped.(bool)
	}

	hasIsUnreachable, err := taskResPy.HasAttr("is_unreachable")
	if err != nil {
		return nil, err
	}

	if hasIsUnreachable {
		fmt.Println("P5")
		isUnreachablePy, err := taskResPy.GetAttr("is_unreachable")
		if err != nil {
			return nil, err
		}

		isUnreachable, err := isUnreachablePy.ToStandartGoType()
		if err != nil {
			return nil, err
		}

		res.IsUnreachable = isUnreachable.(bool)
	}

	return &res, nil
}
