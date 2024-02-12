package php

import (
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/types"
)

func DeclClassConst(ce *types.Class, flags uint32, name string, value types.Zval) {
	constant := types.NewClassConstant(name, value, "", flags)
	ce.ConstantTable().Add(ascii.StrToLower(name), constant)
}

func DeclPropertyInfo(ce *types.Class, flags uint32, name string, typ *types.TypeHint, defaultValue types.Zval) *types.PropertyInfo {
	var propName = name
	if flags&types.AccProtected != 0 {
		propName = ManglePropertyName("*", name)
	} else if flags&types.AccPrivate != 0 {
		propName = ManglePropertyName(ce.Name(), name)
	}

	propInfo := types.NewPropertyInfo(flags, propName, typ, defaultValue)

	ce.PropertyTable().Add(ascii.StrToLower(propName), propInfo)
	return propInfo
}

func DeclMethod(ce *types.Class, fn *types.Function) {
	fn.SetScope(ce)
	lcName := ascii.StrToLower(fn.Name())
	ce.FunctionTable().Add(lcName, fn)

	switch lcName {
	case "__constructor":
		ce.SetConstructor(fn)
	case "__destructor":
		ce.SetDestructor(fn)
	case "__clone":
		ce.SetClone(fn)
	case "__get":
		ce.SetGet(fn)
	case "__set":
		ce.SetSet(fn)
	case "__unset":
		ce.SetUnset(fn)
	case "__isset":
		ce.SetIsset(fn)
	case "__call":
		ce.SetCall(fn)
	case "__callstatic":
		ce.SetCallstatic(fn)
	case "__tostring":
		ce.SetTostring(fn)
	case "__debugInfo":
		ce.SetDebugInfo(fn)
	case "__serialize":
		ce.SetSerializeFunc(fn)
	case "__unserialize":
		ce.SetUnserializeFunc(fn)
	}
}
