package gocotea

import (
	gopython "github.com/ispras/gopython/src/gopython"
)

func RemovePyModulesFromImported(module_name_like string) error {
	// from cotea.utils import remove_modules_from_imported
	coteaUtilsName := "cotea.utils"

	var utilsModule gopython.PythonModule
	utilsModule.SetModuleName(coteaUtilsName)
	err := utilsModule.MakeImport()
	if err != nil {
		return err
	}

	removeFuncName := "remove_modules_from_imported"

	removeFunc, err := utilsModule.GetObject(removeFuncName)
	if err != nil {
		return err
	}

	var args gopython.PythonMethodArguments
	args.SetArgCount(1)
	args.SetNextArgument("cotea")

	_, err = removeFunc.CallItself(&args)
	if err != nil {
		return err
	}

	return nil
}
