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

type objCtorType func(*types.ClassEntry) *types.ZendObject

func SplRegisterClass(name string, objCtor objCtorType, builtinFunctions []types.FunctionEntry) *types.ClassEntry {
	ce := RegisterInternalClass(name, builtinFunctions)
	if objCtor != nil {
		ce.SetCreateObject(objCtor)
	}
	return ce
}
func SplRegisterSubClass(parentCe *types.ClassEntry, name string, objCtor objCtorType, builtinFunctions []types.FunctionEntry) *types.ClassEntry {
	ce := RegisterInternalClassEx(name, builtinFunctions, parentCe)
	if objCtor != nil {
		ce.SetCreateObject(objCtor)
	} else {
		ce.SetCreateObject(parentCe.GetCreateObject())
	}
	return ce
}

func _doRegisterInternalClass(name string, builtinFunctions []types.FunctionEntry, ceFlags uint32) *types.ClassEntry {
	var moduleNumber = EG__().GetCurrentModule().GetModuleNumber()
	var classEntry = types.NewInternalClass(name, moduleNumber)
	classEntry.SetCeFlags(ceFlags | AccConstantsUpdated | AccLinked | AccResolvedParent | AccResolvedInterfaces)
	if len(builtinFunctions) != 0 {
		ZendRegisterFunctions(classEntry, builtinFunctions, classEntry.FunctionTable(), EG__().GetCurrentModule().GetType())
	}
	CG__().ClassTable().Update(classEntry.Name(), classEntry)
	return classEntry
}
