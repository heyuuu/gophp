// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
	"sik/zend/argparse"
	"sik/zend/faults"
	"sik/zend/types"
)

func ZifGettype(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var arg *types.Zval
	var type_ *types.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			arg = fp.ParseZval()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	type_ = types.ZendZvalGetType(arg)
	if type_ != nil {
		return_value.SetInternedString(type_)
		return
	} else {
		return_value.SetRawString("unknown type")
		return
	}
}
func ZifSettype(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var var_ *types.Zval
	var type_ *types.ZendString
	var tmp types.Zval
	var ptr *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			var_ = fp.ParseZval()
			type_ = fp.ParseStr()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	b.Assert(var_.IsReference())
	if zend.ZEND_REF_HAS_TYPE_SOURCES(var_.GetRef()) {
		types.ZVAL_COPY(&tmp, types.Z_REFVAL_P(var_))
		ptr = &tmp
	} else {
		ptr = types.Z_REFVAL_P(var_)
	}
	if types.ZendStringEqualsLiteralCi(type_, "integer") {
		zend.ConvertToLong(ptr)
	} else if types.ZendStringEqualsLiteralCi(type_, "int") {
		zend.ConvertToLong(ptr)
	} else if types.ZendStringEqualsLiteralCi(type_, "float") {
		zend.ConvertToDouble(ptr)
	} else if types.ZendStringEqualsLiteralCi(type_, "double") {
		zend.ConvertToDouble(ptr)
	} else if types.ZendStringEqualsLiteralCi(type_, "string") {
		zend.ConvertToString(ptr)
	} else if types.ZendStringEqualsLiteralCi(type_, "array") {
		zend.ConvertToArray(ptr)
	} else if types.ZendStringEqualsLiteralCi(type_, "object") {
		zend.ConvertToObject(ptr)
	} else if types.ZendStringEqualsLiteralCi(type_, "bool") {
		zend.ConvertToBoolean(ptr)
	} else if types.ZendStringEqualsLiteralCi(type_, "boolean") {
		zend.ConvertToBoolean(ptr)
	} else if types.ZendStringEqualsLiteralCi(type_, "null") {
		zend.ConvertToNull(ptr)
	} else {
		if ptr == &tmp {
			zend.ZvalPtrDtor(&tmp)
		}
		if types.ZendStringEqualsLiteralCi(type_, "resource") {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Cannot convert to resource type")
		} else {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid type")
		}
		return_value.SetFalse()
		return
	}
	if ptr == &tmp {
		zend.ZendTryAssignTypedRef(var_.GetRef(), &tmp)
	}
	return_value.SetTrue()
}
func ZifIntval(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var num *types.Zval
	var base zend.ZendLong = 10
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			num = fp.ParseZval()
			fp.StartOptional()
			base = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if num.GetType() != types.IS_STRING || base == 10 {
		return_value.SetLong(zend.ZvalGetLong(num))
		return
	}
	if base == 0 || base == 2 {
		var strval *byte = num.GetStr().GetVal()
		var strlen int = num.GetStr().GetLen()
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
	return_value.SetLong(zend.ZEND_STRTOL(num.GetStr().GetVal(), nil, base))
}
func ZifFloatval(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var num *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			num = fp.ParseZval()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetDouble(zend.ZvalGetDouble(num))
	return
}
func ZifBoolval(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var val *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			val = fp.ParseZval()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	types.ZVAL_BOOL(return_value, zend.ZendIsTrue(val) != 0)
	return
}
func ZifStrval(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var num *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			num = fp.ParseZval()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetString(zend.ZvalGetString(num))
}
func PhpIsType(executeData *zend.ZendExecuteData, return_value *types.Zval, type_ int) {
	var arg *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			arg = fp.ParseZval()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if arg.IsType(type_) {
		if type_ == types.IS_RESOURCE {
			var type_name *byte = zend.ZendRsrcListGetRsrcType(arg.GetRes())
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
func ZifIsNull(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpIsType(executeData, return_value, types.IS_NULL)
}
func ZifIsResource(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpIsType(executeData, return_value, types.IS_RESOURCE)
}
func ZifIsBool(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var arg *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			arg = fp.ParseZval()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	types.ZVAL_BOOL(return_value, arg.IsType(types.IS_FALSE) || arg.IsType(types.IS_TRUE))
	return
}
func ZifIsInt(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpIsType(executeData, return_value, types.IS_LONG)
}
func ZifIsFloat(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpIsType(executeData, return_value, types.IS_DOUBLE)
}
func ZifIsString(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpIsType(executeData, return_value, types.IS_STRING)
}
func ZifIsArray(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpIsType(executeData, return_value, types.IS_ARRAY)
}
func ZifIsObject(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpIsType(executeData, return_value, types.IS_OBJECT)
}
func ZifIsNumeric(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var arg *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			arg = fp.ParseZval()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	switch arg.GetType() {
	case types.IS_LONG:
		fallthrough
	case types.IS_DOUBLE:
		return_value.SetTrue()
		return
	case types.IS_STRING:
		if zend.IsNumericString(arg.GetStr().GetStr(), nil, nil, 0) != 0 {
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
func ZifIsScalar(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var arg *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			arg = fp.ParseZval()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	switch arg.GetType() {
	case types.IS_FALSE:
		fallthrough
	case types.IS_TRUE:
		fallthrough
	case types.IS_DOUBLE:
		fallthrough
	case types.IS_LONG:
		fallthrough
	case types.IS_STRING:
		return_value.SetTrue()
		return
	default:
		return_value.SetFalse()
		return
	}
}
func ZifIsCallable(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var var_ *types.Zval
	var callable_name *types.Zval = nil
	var name *types.ZendString
	var error *byte
	var retval types.ZendBool
	var syntax_only types.ZendBool = 0
	var check_flags int = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			var_ = fp.ParseZval()
			fp.StartOptional()
			syntax_only = fp.ParseBool()
			callable_name = fp.ParseZval()
			if fp.HasError() {
				fp.HandleError()
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
	types.ZVAL_BOOL(return_value, retval != 0)
	return
}
func ZifIsIterable(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var var_ *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			var_ = fp.ParseZval()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	types.ZVAL_BOOL(return_value, zend.ZendIsIterable(var_) != 0)
	return
}
func ZifIsCountable(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var var_ *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			var_ = fp.ParseZval()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	types.ZVAL_BOOL(return_value, zend.ZendIsCountable(var_) != 0)
	return
}
