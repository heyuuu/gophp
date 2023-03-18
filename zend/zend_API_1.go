// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/types"
)

func _ZEND_TRY_ASSIGN_VALUE_EX(zv *types.Zval, other_zv *types.Zval, strict types.ZendBool, is_ref int) {
	for {
		var _zv *types.Zval = zv
		if is_ref != 0 || _zv.IsReference() {
			var ref *types.ZendReference = _zv.GetRef()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefZvalEx(ref, other_zv, strict)
				break
			}
			_zv = ref.GetVal()
		}
		ZvalPtrDtor(_zv)
		types.ZVAL_COPY_VALUE(_zv, other_zv)
		break
	}
}
func ZEND_TRY_ASSIGN_VALUE_EX(zv *types.Zval, other_zv *types.Zval, strict types.ZendBool) {
	_ZEND_TRY_ASSIGN_VALUE_EX(zv, other_zv, strict, 0)
}
func ZEND_TRY_ASSIGN_COPY_EX(zv *types.Zval, other_zv *types.Zval, strict types.ZendBool) {
	other_zv.TryAddRefcount()
	ZEND_TRY_ASSIGN_VALUE_EX(zv, other_zv, strict)
}
func ZendTryArrayInitSize(zv *types.Zval, size uint32) *types.Zval {
	var arr *types.ZendArray = ZendNewArray(size)
	if zv.IsReference() {
		var ref *types.ZendReference = zv.GetRef()
		if ZEND_REF_HAS_TYPE_SOURCES(ref) {
			if ZendTryAssignTypedRefArr(ref, arr) != types.SUCCESS {
				return nil
			}
			return ref.GetVal()
		}
		zv = ref.GetVal()
	}
	ZvalPtrDtor(zv)
	zv.SetArray(arr)
	return zv
}
func ZendTryArrayInit(zv *types.Zval) *types.Zval { return ZendTryArrayInitSize(zv, 0) }
func Z_PARAM_PROLOGUE(deref int, separate int) {
	_i++
	b.Assert(_i <= _min_num_args || _optional == 1)
	b.Assert(_i > _min_num_args || _optional == 0)
	if _optional {
		if _i > _num_args {
			break
		}
	}
	_real_arg++
	_arg = _real_arg
	if deref != 0 {
		if _arg.IsReference() {
			_arg = types.Z_REFVAL_P(_arg)
		}
	}
	if separate != 0 {
		types.SEPARATE_ZVAL_NOREF(_arg)
	}
}
func ZendParseArgBool(arg *types.Zval, dest *types.ZendBool, is_null *types.ZendBool, check_null int) int {
	val, isNull, ok := ParseArgBool(arg, check_null != 0)
	*dest = types.intBool(val)
	if check_null != 0 {
		*is_null = types.intBool(isNull)
	}
	return types.intBool(ok)
}

func ZendParseArgLong00(arg *types.Zval, dest *ZendLong) bool {
	val, _, ok := ParseArgLong(arg, false, false)
	*dest = val
	return ok
}

func ZendParseArgLong(arg *types.Zval, dest *ZendLong, is_null *types.ZendBool, check_null int, cap int) bool {
	val, isNull, ok := ParseArgLong(arg, check_null != 0, cap != 0)
	*dest = val
	if is_null != nil {
		*is_null = types.intBool(isNull)
	}
	return ok
}
func ZendParseArgDouble(arg *types.Zval, dest *float64, is_null *types.ZendBool, check_null int) int {
	val, isNull, ok := ParseArgDouble(arg, check_null != 0)
	*dest = val
	if is_null != nil {
		*is_null = types.intBool(isNull)
	}
	return types.intBool(ok)
}
func ZendParseArgStr(arg *types.Zval, dest **types.ZendString, check_null int) int {
	val, _, ok := ParseArgStr(arg, check_null != 0, isArgUseStrictTypes())
	// 为空时 *dest 直接为 nil，不需单独的 is_null 字符安
	*dest = val
	return types.intBool(ok)
}
func ZendParseArgString(arg *types.Zval, dest **byte, dest_len *int, check_null int) int {
	var str *types.ZendString
	if ZendParseArgStr(arg, &str, check_null) == 0 {
		return 0
	}
	if check_null != 0 && str == nil {
		*dest = nil
		*dest_len = 0
	} else {
		*dest = str.GetVal()
		*dest_len = str.GetLen()
	}
	return 1
}
func ZendParseArgPathStr(arg *types.Zval, dest **types.ZendString, check_null int) int {
	if ZendParseArgStr(arg, dest, check_null) == 0 || (*dest) != nil && CHECK_NULL_PATH(dest.GetVal(), dest.GetLen()) {
		return 0
	}
	return 1
}
func ZendParseArgPath(arg *types.Zval, dest **byte, dest_len *int, check_null int) int {
	var str *types.ZendString
	if ZendParseArgPathStr(arg, &str, check_null) == 0 {
		return 0
	}
	if check_null != 0 && str == nil {
		*dest = nil
		*dest_len = 0
	} else {
		*dest = str.GetVal()
		*dest_len = str.GetLen()
	}
	return 1
}
func ZendParseArgArray(arg *types.Zval, dest **types.Zval, check_null int, or_object int) int {
	if arg.IsArray() || or_object != 0 && arg.IsObject() {
		*dest = arg
	} else if check_null != 0 && arg.IsNull() {
		*dest = nil
	} else {
		return 0
	}
	return 1
}
func ZendParseArgArrayHt(arg *types.Zval, dest **types.HashTable, check_null int, or_object int, separate int) int {
	if arg.IsArray() {
		*dest = arg.GetArr()
	} else if or_object != 0 && arg.IsObject() {
		if separate != 0 && types.Z_OBJ_P(arg).GetProperties() != nil && types.Z_OBJ_P(arg).GetProperties().GetRefcount() > 1 {
			if (types.Z_OBJ_P(arg).GetProperties().GetGcFlags() & types.IS_ARRAY_IMMUTABLE) == 0 {
				types.Z_OBJ_P(arg).GetProperties().DelRefcount()
			}
			types.Z_OBJ_P(arg).SetProperties(ZendArrayDup(types.Z_OBJ_P(arg).GetProperties()))
		}
		*dest = types.Z_OBJ_HT_P(arg).GetGetProperties()(arg)
	} else if check_null != 0 && arg.IsNull() {
		*dest = nil
	} else {
		return 0
	}
	return 1
}
func ZendParseArgObject(arg *types.Zval, dest **types.Zval, ce *ZendClassEntry, check_null int) int {
	if arg.IsObject() && (ce == nil || InstanceofFunction(types.Z_OBJCE_P(arg), ce) != 0) {
		*dest = arg
	} else if check_null != 0 && arg.IsNull() {
		*dest = nil
	} else {
		return 0
	}
	return 1
}
func ZendParseArgResource(arg *types.Zval, dest **types.Zval, check_null int) int {
	if arg.IsResource() {
		*dest = arg
	} else if check_null != 0 && arg.IsNull() {
		*dest = nil
	} else {
		return 0
	}
	return 1
}
func ZendParseArgFunc(arg *types.Zval, dest_fci *ZendFcallInfo, dest_fcc *ZendFcallInfoCache, check_null int, error **byte) int {
	if check_null != 0 && arg.IsNull() {
		dest_fci.SetSize(0)
		dest_fcc.SetFunctionHandler(nil)
		*error = nil
	} else if ZendFcallInfoInit(arg, 0, dest_fci, dest_fcc, nil, error) != types.SUCCESS {
		return 0
	}
	return 1
}
func ZendParseArgZvalDeref(arg *types.Zval, dest **types.Zval, check_null int) {
	if check_null != 0 && arg.IsNull() {
		*dest = nil
	} else {
		*dest = arg
	}
}
func _zendGetParametersArrayEx(param_count int, argument_array *types.Zval) int {
	var param_ptr *types.Zval
	var arg_count int
	param_ptr = CurrEX().Arg(1)
	arg_count = CurrEX().NumArgs()
	if param_count > arg_count {
		return types.FAILURE
	}
	for b.PostDec(&param_count) > 0 {
		types.ZVAL_COPY_VALUE(argument_array, param_ptr)
		argument_array++
		param_ptr++
	}
	return types.SUCCESS
}
func ZendCopyParametersArray(param_count int, argument_array *types.Zval) int {
	var param_ptr *types.Zval
	var arg_count int
	param_ptr = CurrEX().Arg(1)
	arg_count = CurrEX().NumArgs()
	if param_count > arg_count {
		return types.FAILURE
	}
	for b.PostDec(&param_count) > 0 {
		param_ptr.TryAddRefcount()
		argument_array.GetArr().NextIndexInsertNew(param_ptr)
		param_ptr++
	}
	return types.SUCCESS
}
func ZendWrongParamCount() {
	ZendInternalArgumentCountError(CurrEX().IsArgUseStrictTypes(), "Wrong parameter count for %s()", GetActiveCalleeName())
}
func ZendGetTypeByConst(type_ types.ZendUchar) string {
	switch type_ {
	case types.IS_FALSE, types.IS_TRUE, types._IS_BOOL:
		return "bool"
	case types.IS_LONG:
		return "int"
	case types.IS_DOUBLE:
		return "float"
	case types.IS_STRING:
		return "string"
	case types.IS_OBJECT:
		return "object"
	case types.IS_RESOURCE:
		return "resource"
	case types.IS_NULL:
		return "null"
	case types.IS_CALLABLE:
		return "callable"
	case types.IS_ITERABLE:
		return "iterable"
	case types.IS_ARRAY:
		return "array"
	case types.IS_VOID:
		return "void"
	case types._IS_NUMBER:
		return "number"
	default:
		return "unknown"
	}
}
func ZendZvalTypeName(arg *types.Zval) string {
	arg = types.ZVAL_DEREF(arg)
	return ZendGetTypeByConst(arg.GetType())
}
func ZendZvalGetType(arg *types.Zval) *types.ZendString {
	switch arg.GetType() {
	case types.IS_NULL:
		return types.ZSTR_KNOWN(types.ZEND_STR_NULL)
	case types.IS_FALSE:

	case types.IS_TRUE:
		return types.ZSTR_KNOWN(types.ZEND_STR_BOOLEAN)
	case types.IS_LONG:
		return types.ZSTR_KNOWN(types.ZEND_STR_INTEGER)
	case types.IS_DOUBLE:
		return types.ZSTR_KNOWN(types.ZEND_STR_DOUBLE)
	case types.IS_STRING:
		return types.ZSTR_KNOWN(types.ZEND_STR_STRING)
	case types.IS_ARRAY:
		return types.ZSTR_KNOWN(types.ZEND_STR_ARRAY)
	case types.IS_OBJECT:
		return types.ZSTR_KNOWN(types.ZEND_STR_OBJECT)
	case types.IS_RESOURCE:
		if ZendRsrcListGetRsrcType(arg.GetRes()) != nil {
			return types.ZSTR_KNOWN(types.ZEND_STR_RESOURCE)
		} else {
			return types.ZSTR_KNOWN(types.ZEND_STR_CLOSED_RESOURCE)
		}
	default:
		return nil
	}
}
