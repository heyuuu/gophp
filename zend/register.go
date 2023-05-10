package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

// Register Internal class

func RegisterInternalClass(name string, builtinFunctions []types.FunctionEntry) *types.ClassEntry {
	return _doRegisterInternalClass(name, builtinFunctions, 0)
}
func RegisterInternalInterface(name string, builtinFunctions []types.FunctionEntry) *types.ClassEntry {
	return _doRegisterInternalClass(name, builtinFunctions, AccInterface)
}
func RegisterInternalClassEx(name string, builtinFunctions []types.FunctionEntry, parentCe *types.ClassEntry) *types.ClassEntry {
	registerClass := RegisterInternalClass(name, builtinFunctions)
	if parentCe != nil {
		ZendDoInheritance(registerClass, parentCe)
		ZendBuildPropertiesInfoTable(registerClass)
	}
	return registerClass
}

func _doRegisterInternalClass(name string, builtinFunctions []types.FunctionEntry, ceFlags uint32) *types.ClassEntry {
	var classEntry = types.NewInternalClass(name)
	classEntry.SetCeFlags(ceFlags | AccConstantsUpdated | AccLinked | AccResolvedParent | AccResolvedInterfaces)
	classEntry.SetModule(EG__().GetCurrentModule())
	if len(builtinFunctions) != 0 {
		ZendRegisterFunctions(classEntry, builtinFunctions, classEntry.FunctionTable(), EG__().GetCurrentModule().GetType())
	}
	CG__().ClassTable().Update(classEntry.Name(), classEntry)
	return classEntry
}
