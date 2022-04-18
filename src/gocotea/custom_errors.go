package gocotea

import (
	"fmt"
)

type PythonImportError struct {
	ModuleName string
	ErrorMsg   string
}

func (p *PythonImportError) Error() string {
	format := "Can't import module with name %s."
	format += "Error message:\n%s.\n"
	format += "Probably you don't set PYTHONPATH env var properly. "
	format += "Make:\n\texport PYTHONPATH=$GOPATH/src/github.com/ispras/gocotea/src"
	return fmt.Sprintf(format, p.ModuleName, p.ErrorMsg)
}

type PythonAttrError struct {
	SourceUnitName string // it can be module or object
	AttrName       string // it can be any oject
	ErrorMsg       string
}

func (p *PythonAttrError) Error() string {
	format := "Module or object %s has no attr %s\n"
	format += "Error message:\n%s\n"
	return fmt.Sprintf(format, p.SourceUnitName, p.AttrName,
		p.ErrorMsg)
}

type PythonObjectCreationError struct {
	ClassName string
	ErrorMsg  string
}

func (p *PythonObjectCreationError) Error() string {
	format := "Can't create object of %s class.\n"
	format += "Error message:\n%s\n"
	return fmt.Sprintf(format, p.ClassName, p.ErrorMsg)
}

type PythonCallMethodError struct {
	MethodName string
	ClassName  string
	ErrorMsg   string
}

func (p *PythonCallMethodError) Error() string {
	format := "Can't call %s.%s method.\n"
	format += "Error message:\n%s\n"
	return fmt.Sprintf(format, p.ClassName, p.MethodName, p.ErrorMsg)
}
