package php

import (
	"fmt"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
)

const wrongPropertyOffset = 0
const dynamicPropertyOffset = -1

func throwBadPropertyAccess(ctx *Context, propInfo *types.PropertyInfo, ce *types.Class, member string) {
	ThrowError(ctx, nil, fmt.Sprintf("Cannot access %s property %s::$%s", visibilityString(propInfo.Flags()), ce.Name(), member))
}
func throwBadPropertyName(ctx *Context) {
	ThrowError(ctx, nil, "Cannot access property started with '\\0'")
}
func getPropertyOffset(ctx *Context, ce *types.Class, member string, silent bool) (int, *types.PropertyInfo) {
	var propertyInfo *types.PropertyInfo
	var offset uint32

	if propertyInfo = ce.GetProperty(member); propertyInfo == nil {
		if member != "" && member[0] == '\x00' {
			if !silent {
				throwBadPropertyName(ctx)
			}
			return wrongPropertyOffset, nil
		}
		return dynamicPropertyOffset, nil
	}
	if propertyInfo.IsStatic() {
		if !silent {
			Error(ctx, perr.E_NOTICE, fmt.Sprintf("Accessing static property %s::$%s as non static", ce.Name(), member))
		}
		return dynamicPropertyOffset, nil
	}
	offset = propertyInfo.Offset()
	if propertyInfo.Type() == nil {
		propertyInfo = nil
	}
	return int(offset), propertyInfo
}

func StdCastObjectToString(ctx *Context, obj *types.Object) (string, bool) {
	ce := obj.Class()
	if ce.GetTostring() != nil {
		retval := CallMethod(ctx, obj, ce, ce.GetTostring(), nil)
		if retval.IsString() {
			return retval.String(), true
		}
		if !ctx.EG().HasException() {
			ThrowError(ctx, nil, fmt.Sprintf("Method %s::__toString() must return a string value", ce.Name()))
		}
	}
	return "", false
}
