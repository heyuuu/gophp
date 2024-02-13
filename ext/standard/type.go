package standard

import (
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/php/zpp"
)

func ZifGettype(var_ types.Zval) string {
	return types.ZvalGetType(var_)
}
func ZifSettype(ctx *php.Context, var_ zpp.RefZval, typ string) bool {
	typ = ascii.StrToLower(typ)
	switch typ {
	case "integer", "int":
		var_.SetVal(types.ZvalLong(php.ZvalGetLong(ctx, var_.Val())))
	case "float", "double":
		var_.SetVal(types.ZvalDouble(php.ZvalGetDouble(ctx, var_.Val())))
	case "string":
		var_.SetVal(types.ZvalString(php.ZvalTryGetStrVal(ctx, var_.Val())))
	case "array":
		var_.SetVal(types.ZvalArray(php.ZvalGetArray(ctx, var_.Val())))
	case "object":
		var_.SetVal(types.ZvalObject(php.ZvalGetObject(ctx, var_.Val())))
	case "bool", "boolean":
		var_.SetVal(types.ZvalBool(php.ZvalIsTrue(ctx, var_.Val())))
	case "null":
		var_.SetVal(types.Null)
	default:
		if typ == "resource" {
			php.ErrorDocRef(ctx, "", perr.E_WARNING, "Cannot convert to resource type")
		} else {
			php.ErrorDocRef(ctx, "", perr.E_WARNING, "Invalid type")
		}
		return false
	}

	return true
}
func ZifIntval(ctx *php.Context, var_ types.Zval, _ zpp.Opt, base_ *int) int {
	var num = var_
	var base = lang.Option(base_, 10)
	if !num.IsString() || base == 10 {
		return php.ZvalGetLong(ctx, num)
	}
	if base == 0 || base == 2 {
		str := num.String()
		for str != "" && ascii.IsSpace(str[0]) {
			str = str[1:]
		}

		/* Length of 3+ covers "0b#" and "-0b" (which results in 0) */
		if len(str) > 2 {
			var offset = 0
			if str[0] == '-' || str[0] == '+' {
				offset = 1
			}
			if str[offset] == '0' && (str[offset+1] == 'b' || str[offset+1] == 'B') {
				str = str[:offset] + str[offset+2:]
				return php.ZendStrToLong(str, 2)
			}
		}
	}
	return php.ZendStrToLong(num.String(), base)
}

// @zif(alias="doubleval")
func ZifFloatval(ctx *php.Context, var_ types.Zval) float64 {
	return php.ZvalGetDouble(ctx, var_)
}
func ZifBoolval(ctx *php.Context, var_ types.Zval) bool {
	return php.ZvalIsTrue(ctx, var_)
}

// @zif(onError=1)
func ZifStrval(ctx *php.Context, var_ types.Zval) string {
	return php.ZvalGetStrVal(ctx, var_)
}
func ZifIsNull(var_ *types.Zval) bool {
	return var_.IsNull()
}
func ZifIsResource(var_ *types.Zval) bool {
	if var_.IsResource() {
		//typeName := php.ZendRsrcListGetRsrcType(var_.Resource())
		//if typeName == nil {
		//	return false
		//}
		return true
	}

	return false
}

// @zif(onError=1)
func ZifIsBool(var_ *types.Zval) bool {
	return var_.IsBool()
}

// @zif(alias="is_integer,is_long")
func ZifIsInt(var_ *types.Zval) bool {
	return var_.IsLong()
}

// @zif(alias="is_double")
func ZifIsFloat(var_ *types.Zval) bool {
	return var_.IsDouble()
}
func ZifIsString(var_ *types.Zval) bool {
	return var_.IsString()
}
func ZifIsArray(var_ *types.Zval) bool {
	return var_.IsArray()
}
func ZifIsObject(var_ *types.Zval) bool {
	return var_.IsObject()
}

func ZifIsNumeric(value *types.Zval) bool {
	switch value.Type() {
	case types.IsLong, types.IsDouble:
		return true
	case types.IsString:
		return php.IsNumericString(value.String())
	default:
		return false
	}
}
func ZifIsScalar(value *types.Zval) bool {
	switch value.Type() {
	case types.IsFalse, types.IsTrue, types.IsLong, types.IsDouble, types.IsString:
		return true
	default:
		return false
	}
}

//func ZifIsCallable(ctx *php.Context, value *types.Zval, _ zpp.Opt, syntaxOnly bool, callableName zpp.RefZval) bool {
//	var checkFlags uint32 = 0
//	if syntaxOnly {
//		checkFlags |= php.IsCallableCheckSyntaxOnly
//	}
//
//	retval := php.IsCallable(ctx, value, nil, checkFlags)
//	if callableName != nil {
//		name := php.GetCallableName(ctx, value, nil)
//		callableName.SetVal(types.ZvalString(name))
//	}
//	return retval
//}
//func ZifIsIterable(value *types.Zval) bool {
//	return php.ZendIsIterable(value)
//}
//func ZifIsCountable(value *types.Zval) bool {
//	return php.ZendIsCountable(value)
//}
