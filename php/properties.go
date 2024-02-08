package php

import (
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/types"
)

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
