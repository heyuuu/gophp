package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
)

func RegisterInterface(decl *types.InternalClassDecl) *types.ClassEntry {
	decl.IsInterface = true
	return RegisterClass(decl)
}

func RegisterClass(decl *types.InternalClassDecl) *types.ClassEntry {
	b.Assert(decl.Name != "")

	var moduleNumber = EG__().GetCurrentModule().ModuleNumber()
	var ce = types.NewInternalClassEx(decl, moduleNumber)
	if len(decl.Functions) != 0 {
		ZendRegisterFunctions(ce, decl.Functions, ce.FunctionTable())
	}
	CG__().ClassTable().Update(ce.Name(), ce)

	// handle interfaces
	for _, iface := range decl.Interfaces {
		ZendDoImplementInterface(ce, iface)
	}

	if decl.IsInterface {
		if decl.InterfaceGetsImplemented != nil {
			ce.SetInterfaceGetsImplemented(decl.InterfaceGetsImplemented)
		}
	} else {
		// handle parent
		parent := decl.Parent
		if parent != nil {
			ZendDoInheritance(ce, parent)
			ZendBuildPropertiesInfoTable(ce)
		}

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
