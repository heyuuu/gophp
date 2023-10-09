package php

import (
	"fmt"
	"github.com/heyuuu/gophp/php/types"
	"strconv"
)

func formatDouble(v float64) string {
	return strconv.FormatFloat(v, 'G', getPrec(), 64)
}

func ZvalIsTrue(zv *types.Zval) bool {
	zv = zv.DeRef()
	switch zv.Type() {
	case types.IsTrue:
		return true
	case types.IsLong:
		return zv.Long() != 0
	case types.IsDouble:
		return zv.Double() != 0
	case types.IsString:
		str := zv.String()
		return str != "" && str != "0"
	case types.IsArray:
		return zv.Array().Len() != 0
	case types.IsObject:
		return ObjectIsTrue(zv.Object())
	case types.IsResource:
		return zv.Resource().Handle() != 0
	}
	return false
}

func ObjectIsTrue(obj *types.Object) bool {
	// todo try cast
	// todo try get
	return true
}

func ZvalGetLong(zv *types.Zval) int {
	return _zvalGetLongEx(zv, true)
}

func _zvalGetLongEx(zv *types.Zval, silent bool) int {
	zv = zv.DeRef()
	switch zv.Type() {
	case types.IsUndef, types.IsNull, types.IsFalse:
		return 0
	case types.IsTrue:
		return 1
	case types.IsLong:
		return zv.Long()
	case types.IsDouble:
		return DvalToLval(zv.Double())
	case types.IsString:
		// todo parse
		return 0
	case types.IsResource:
		return zv.Resource().Handle()
	case types.IsArray:
		if zv.Array().Len() != 0 {
			return 1
		} else {
			return 0
		}
	case types.IsObject:
		// todo try object cast
		// todo try object get
		return 1
	default:
		return 0
	}
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
