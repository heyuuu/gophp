package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
)

// Register Internal class
type objCtorType func(*types.ClassEntry) *types.Object
type objGetIteratorType func(ce *types.ClassEntry, object *types.Zval, by_ref int) *ZendObjectIterator
type ClassDefines struct {
	Name         string
	CreateObject objCtorType
	Functions    []types.FunctionEntry
	Parent       *types.ClassEntry
	Interfaces   []*types.ClassEntry
	GetIterator  objGetIteratorType
	AddCeFlags   uint32
}

func RegisterInternalInterface(name string, builtinFunctions []types.FunctionEntry) *types.ClassEntry {
	return _doRegisterInternalClass(name, builtinFunctions, types.AccInterface)
}

func RegisterClass(name string, objCtor objCtorType, builtinFunctions []types.FunctionEntry) *types.ClassEntry {
	return RegisterClassEx(&ClassDefines{
		Name:         name,
		Functions:    builtinFunctions,
		CreateObject: objCtor,
	})
}
func RegisterSubClass(parentCe *types.ClassEntry, name string, objCtor objCtorType, builtinFunctions []types.FunctionEntry) *types.ClassEntry {
	return RegisterClassEx(&ClassDefines{
		Name:         name,
		Functions:    builtinFunctions,
		Parent:       parentCe,
		CreateObject: objCtor,
	})
}

func RegisterClassEx(def *ClassDefines) *types.ClassEntry {
	b.Assert(def.Name != "")

	ce := _doRegisterInternalClass(def.Name, def.Functions, 0)

	// handle parent
	parent := def.Parent
	if parent != nil {
		ZendDoInheritance(ce, parent)
		ZendBuildPropertiesInfoTable(ce)
	}

	// handle interfaces
	ZendClassImplements(ce, 1, def.Interfaces...)

	// handle objCtor
	if def.CreateObject != nil {
		ce.SetCreateObject(def.CreateObject)
	} else if parent != nil {
		ce.SetCreateObject(parent.GetCreateObject())
	}

	// others
	if def.GetIterator != nil {
		ce.SetGetIterator(def.GetIterator)
	}
	if def.AddCeFlags != 0 {
		ce.AddCeFlags(def.AddCeFlags)
	}

	return ce
}

func _doRegisterInternalClass(name string, builtinFunctions []types.FunctionEntry, ceFlags uint32) *types.ClassEntry {
	var moduleNumber = EG__().GetCurrentModule().GetModuleNumber()
	var classEntry = types.NewInternalClass(name, moduleNumber)
	classEntry.SetCeFlags(ceFlags | types.AccConstantsUpdated | types.AccLinked | types.AccResolvedParent | types.AccResolvedInterfaces)
	if len(builtinFunctions) != 0 {
		ZendRegisterFunctions(classEntry, builtinFunctions, classEntry.FunctionTable(), EG__().GetCurrentModule().GetType())
	}
	CG__().ClassTable().Update(classEntry.Name(), classEntry)
	return classEntry
}
