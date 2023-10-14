package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
)

func RegisterInternalInterface(name string, builtinFunctions []types.FunctionEntry) *types.ClassEntry {
	return RegisterClassEx(&types.InternalClassDecl{
		Name:        name,
		Functions:   builtinFunctions,
		IsInterface: true,
	})
}

func RegisterClass(name string, objCtor types.ObjCtorType, builtinFunctions []types.FunctionEntry) *types.ClassEntry {
	return RegisterClassEx(&types.InternalClassDecl{
		Name:         name,
		Functions:    builtinFunctions,
		CreateObject: objCtor,
	})
}
func RegisterSubClass(parentCe *types.ClassEntry, name string, objCtor types.ObjCtorType, builtinFunctions []types.FunctionEntry) *types.ClassEntry {
	return RegisterClassEx(&types.InternalClassDecl{
		Name:         name,
		Functions:    builtinFunctions,
		Parent:       parentCe,
		CreateObject: objCtor,
	})
}

func RegisterClassEx(decl *types.InternalClassDecl) *types.ClassEntry {
	b.Assert(decl.Name != "")

	var moduleNumber = EG__().GetCurrentModule().GetModuleNumber()
	var ce = types.NewInternalClassEx(decl, moduleNumber)
	if len(decl.Functions) != 0 {
		ZendRegisterFunctions(ce, decl.Functions, ce.FunctionTable(), EG__().GetCurrentModule().GetType())
	}
	CG__().ClassTable().Update(ce.Name(), ce)

	if !decl.IsInterface {
		// handle parent
		parent := decl.Parent
		if parent != nil {
			ZendDoInheritance(ce, parent)
			ZendBuildPropertiesInfoTable(ce)
		}

		// handle interfaces
		ZendClassImplements(ce, 1, decl.Interfaces...)

		// handle objCtor
		if decl.CreateObject != nil {
			ce.SetCreateObject(decl.CreateObject)
		} else if parent != nil {
			ce.SetCreateObject(parent.GetCreateObject())
		}

		// others
		if decl.GetIterator != nil {
			ce.SetGetIterator(decl.GetIterator)
		}
		if decl.CeFlags != 0 {
			ce.AddCeFlags(decl.CeFlags)
		}
	}

	return ce
}
