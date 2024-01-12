package php

import (
	"fmt"
	"github.com/heyuuu/gophp/php/types"
	"strconv"
)

func formatDouble(v float64) string {
	return strconv.FormatFloat(v, 'G', getPrec(), 64)
}

func ZvalGetString(zv *types.Zval) string {
	str, _ := _zvalGetStringEx(zv, false)
	return str
}

func _zvalGetStringEx(zv *types.Zval, try bool) (string, bool) {
	zv = zv.DeRef()
	switch zv.Type() {
	case types.IsString:
		return zv.String(), true
	case types.IsUndef, types.IsNull, types.IsFalse:
		return "", true
	case types.IsTrue:
		return "1", true
	case types.IsResource:
		return fmt.Sprintf("Resource id #%d", zv.Resource().Handle()), true
	case types.IsLong:
		return strconv.Itoa(zv.Long()), true
	case types.IsDouble:
		return formatDouble(zv.Double()), true
	case types.IsArray:
		// todo error

		return "Array", true
	case types.IsObject:
		// todo try object cast
		// todo try object get
		if try {
			return "", false
		} else {
			return "", true
		}
	default:
		return "", false
	}
}
