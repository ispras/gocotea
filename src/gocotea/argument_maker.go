package gocotea

import (
	gopython "github.com/ispras/gopython/src/gopython"
)

type ArgumentMaker struct {
	classNamePy          string
	argMakerPythonObject *gopython.PythonObject
}

func (argmaker *ArgumentMaker) InitArgMaker() error {
	argumentsMakerModuleName := "cotea.arguments_maker"
	argmaker.classNamePy = "argument_maker"

	var argmakerModule gopython.PythonModule
	argmakerModule.SetModuleName(argumentsMakerModuleName)
	err := argmakerModule.MakeImport()
	if err != nil {
		return &PythonImportError{ModuleName: argumentsMakerModuleName,
			ErrorMsg: err.Error()}
	}

	var argMakerClass *gopython.PythonClass
	argMakerClass, err = argmakerModule.GetClass(argmaker.classNamePy)
	if err != nil {
		return &PythonAttrError{SourceUnitName: argumentsMakerModuleName,
			AttrName: argmaker.classNamePy,
			ErrorMsg: err.Error()}
	}

	var initArgs gopython.PythonMethodArguments
	initArgs.SetArgCount(0) // method doesn't take any arguments

	argmaker.argMakerPythonObject, err = argMakerClass.CreateObject(&initArgs)
	if err != nil {
		return &PythonObjectCreationError{ClassName: argmaker.classNamePy,
			ErrorMsg: err.Error()}
	}

	return nil
}

func (argmaker *ArgumentMaker) AddArgument(args ...string) error {
	var argsToPythonMethod gopython.PythonMethodArguments
	methodName := "add_arg"

	if len(args) == 1 {
		argsToPythonMethod.SetArgCount(1)
		argsToPythonMethod.SetNextArgument(args[0])
	} else {
		argsToPythonMethod.SetArgCount(2)
		argsToPythonMethod.SetNextArgument(args[0])
		argsToPythonMethod.SetNextArgument(args[1])
	}

	_, err := argmaker.argMakerPythonObject.CallMethod(methodName, &argsToPythonMethod)
	if err != nil {
		return &PythonCallMethodError{ClassName: argmaker.classNamePy,
			MethodName: methodName, ErrorMsg: err.Error()}
	}

	return nil
}
