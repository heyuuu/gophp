package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func DisplayDisabledClass(classType *types.ClassEntry) *types.Object {
	var intern *types.Object
	intern = types.NewStdObjectSkipPropertiesInit(classType)
	faults.Error(faults.E_WARNING, "%s() has been disabled for security reasons", classType.Name())
	return intern
}
func ZendDisableClass(className string) int {
	ce := CG__().ClassTable().Get(className)
	if ce == nil {
		return types.FAILURE
	}

	disabledClass := types.NewDisabledClass(ce, DisplayDisabledClass)
	CG__().ClassTable().Update(className, disabledClass)
	return types.SUCCESS
}
