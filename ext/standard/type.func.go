package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"github.com/heyuuu/gophp/zend/zpp"
)

func ZifGettype(var_ *types.Zval) string {
	return types.ZvalGetType(var_)
}
func ZifSettype(var_ zpp.RefZval, typ string) bool {
	var tmp types.Zval
	var ptr *types.Zval
	b.Assert(var_.IsRef())
	if zend.ZEND_REF_HAS_TYPE_SOURCES(var_.Ref()) {
		types.ZVAL_COPY(&tmp, types.Z_REFVAL_P(var_))
		ptr = &tmp
	} else {
		ptr = types.Z_REFVAL_P(var_)
	}

	typ = ascii.StrToLower(typ)
	switch typ {
	case "integer", "int":
		operators.ConvertToLong(ptr)
	case "float", "double":
		operators.ConvertToDouble(ptr)
	case "string":
		operators.ConvertToString(ptr)
	case "array":
		operators.ConvertToArray(ptr)
	case "object":
		operators.ConvertToObject(ptr)
	case "bool", "boolean":
		operators.ConvertToBoolean(ptr)
	case "null":
		operators.ConvertToNull(ptr)
	default:
		if ptr == &tmp {
			// zend.ZvalPtrDtor(&tmp)
		}
		if typ == "resource" {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Cannot convert to resource type")
		} else {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid type")
		}
		return false
	}

	if ptr == &tmp {
		zend.ZendTryAssignTypedRef(var_.Ref(), &tmp)
	}
	return true
}
func ZifIntval(executeData zpp.Ex, return_value zpp.Ret, var_ *types.Zval, _ zpp.Opt, base *types.Zval) {
	var num *types.Zval
	var base zend.ZendLong = 10
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 2, 0)
			num = fp.ParseZval()
			fp.StartOptional()
			base = fp.ParseLong()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if !num.IsString() || base == 10 {
		return_value.SetLong(operators.ZvalGetLong(num))
		return
	}
	if base == 0 || base == 2 {
		var strval *byte = num.StringEx().GetVal()
		var strlen int = num.StringEx().GetLen()
		for isspace(*strval) && strlen != 0 {
			strval++
			strlen--
		}

		/* Length of 3+ covers "0b#" and "-0b" (which results in 0) */

		if strlen > 2 {
			var offset int = 0
			if strval[0] == '-' || strval[0] == '+' {
				offset = 1
			}
			if strval[offset] == '0' && (strval[offset+1] == 'b' || strval[offset+1] == 'B') {
				var tmpval *byte
				strlen -= 2
				tmpval = zend.Emalloc(strlen + 1)

				/* Place the unary symbol at pos 0 if there was one */

				if offset != 0 {
					tmpval[0] = strval[0]
				}

				/* Copy the data from after "0b" to the end of the buffer */

				memcpy(tmpval+offset, strval+offset+2, strlen-offset)
				tmpval[strlen] = 0
				return_value.SetLong(zend.ZEND_STRTOL(tmpval, nil, 2))
				zend.Efree(tmpval)
				return
			}
		}

		/* Length of 3+ covers "0b#" and "-0b" (which results in 0) */

	}
	return_value.SetLong(zend.ZEND_STRTOL(num.StringEx().GetVal(), nil, base))
}

//@zif -alias doubleval
func ZifFloatval(executeData zpp.Ex, return_value zpp.Ret, var_ *types.Zval) {
	var num *types.Zval
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			num = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	return_value.SetDouble(operators.ZvalGetDouble(num))
	return
}
func ZifBoolval(executeData zpp.Ex, return_value zpp.Ret, var_ *types.Zval) {
	var val *types.Zval
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			val = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	return_value.SetBool(operators.ZvalIsTrue(val))
	return
}
func ZifStrval(executeData zpp.Ex, return_value zpp.Ret, var_ *types.Zval) {
	var num *types.Zval
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			num = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	return_value.SetStringEx(operators.ZvalGetString(num))
}
func PhpIsType(executeData *zend.ZendExecuteData, return_value *types.Zval, type_ int) {
	var arg *types.Zval
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			arg = fp.ParseZval()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if arg.IsType(type_) {
		if type_ == types.IsResource {
			var type_name *byte = zend.ZendRsrcListGetRsrcType(arg.Resource())
			if type_name == nil {
				return_value.SetFalse()
				return
			}
		}
		return_value.SetTrue()
		return
	} else {
		return_value.SetFalse()
		return
	}
}
func ZifIsNull(executeData zpp.Ex, return_value zpp.Ret, var_ *types.Zval) {
	PhpIsType(executeData, return_value, types.IsNull)
}
func ZifIsResource(executeData zpp.Ex, return_value zpp.Ret, var_ *types.Zval) {
	PhpIsType(executeData, return_value, types.IsResource)
}
func ZifIsBool(executeData zpp.Ex, return_value zpp.Ret, var_ *types.Zval) {
	var arg *types.Zval
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			arg = fp.ParseZval()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	return_value.SetBool(arg.IsType(types.IsFalse) || arg.IsType(types.IsTrue))
	return
}

//@zif -alias is_integer,is_long
func ZifIsInt(executeData zpp.Ex, return_value zpp.Ret, var_ *types.Zval) {
	PhpIsType(executeData, return_value, types.IsLong)
}

//@zif -alias is_double
func ZifIsFloat(executeData zpp.Ex, return_value zpp.Ret, var_ *types.Zval) {
	PhpIsType(executeData, return_value, types.IsDouble)
}
func ZifIsString(executeData zpp.Ex, return_value zpp.Ret, var_ *types.Zval) {
	PhpIsType(executeData, return_value, types.IsString)
}
func ZifIsArray(executeData zpp.Ex, return_value zpp.Ret, var_ *types.Zval) {
	PhpIsType(executeData, return_value, types.IsArray)
}
func ZifIsObject(executeData zpp.Ex, return_value zpp.Ret, var_ *types.Zval) {
	PhpIsType(executeData, return_value, types.IsObject)
}
func ZifIsNumeric(executeData zpp.Ex, return_value zpp.Ret, value *types.Zval) {
	var arg *types.Zval
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			arg = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	switch arg.Type() {
	case types.IsLong:
		fallthrough
	case types.IsDouble:
		return_value.SetTrue()
		return
	case types.IsString:
		if operators.IsNumericString(arg.StringEx().GetStr(), nil, nil, 0) != 0 {
			return_value.SetTrue()
			return
		} else {
			return_value.SetFalse()
			return
		}
	default:
		return_value.SetFalse()
		return
	}
}
func ZifIsScalar(executeData zpp.Ex, return_value zpp.Ret, value *types.Zval) {
	var arg *types.Zval
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			arg = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	switch arg.Type() {
	case types.IsFalse:
		fallthrough
	case types.IsTrue:
		fallthrough
	case types.IsDouble:
		fallthrough
	case types.IsLong:
		fallthrough
	case types.IsString:
		return_value.SetTrue()
		return
	default:
		return_value.SetFalse()
		return
	}
}
func ZifIsCallable(executeData zpp.Ex, return_value zpp.Ret, var_ *types.Zval, _ zpp.Opt, syntaxOnly *types.Zval, callableName zpp.RefZval) {
	var var_ *types.Zval
	var callable_name *types.Zval = nil
	var name *types.String
	var error *byte
	var retval bool
	var syntax_only bool = 0
	var check_flags int = 0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 3, 0)
			var_ = fp.ParseZval()
			fp.StartOptional()
			syntax_only = fp.ParseBool()
			callable_name = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if syntax_only != 0 {
		check_flags |= zend.IS_CALLABLE_CHECK_SYNTAX_ONLY
	}
	if executeData.NumArgs() > 2 {
		retval = zend.ZendIsCallableEx(var_, nil, check_flags, &name, nil, &error)
		zend.ZEND_TRY_ASSIGN_REF_STR(callable_name, name)
	} else {
		retval = zend.ZendIsCallableEx(var_, nil, check_flags, nil, nil, &error)
	}
	if error != nil {

		/* ignore errors */

		zend.Efree(error)

		/* ignore errors */

	}
	return_value.SetBool(retval != 0)
	return
}
func ZifIsIterable(executeData zpp.Ex, return_value zpp.Ret, var_ *types.Zval) {
	var var_ *types.Zval
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			var_ = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	return_value.SetBool(zend.ZendIsIterable(var_) != 0)
	return
}
func ZifIsCountable(executeData zpp.Ex, return_value zpp.Ret, var_ *types.Zval) {
	var var_ *types.Zval
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			var_ = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	return_value.SetBool(zend.ZendIsCountable(var_) != 0)
	return
}
