package gocotea

import (
	gopython "github.com/ispras/gopython/src/gopython"
)

func InitPythonInterpretetor() error {
	gopython.InitPythonInterpretetor()
	return nil // NEED TO FIX
}

func FinalizePythonInterpretetor() error {
	gopython.FinalizePythonInterpretetor()
	return nil // NEED TO FIX
}
