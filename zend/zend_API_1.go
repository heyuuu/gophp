// <<generate>>

package zend

import (
	b "sik/builtin"
)

func _ZEND_TRY_ASSIGN_VALUE_EX(zv *Zval, other_zv *Zval, strict ZendBool, is_ref int) {
	for {
		var _zv *Zval = zv
		if is_ref != 0 || _zv.IsReference() {
			var ref *ZendReference = _zv.GetRef()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefZvalEx(ref, other_zv, strict)
				break
			}
			_zv = ref.GetVal()
		}
		ZvalPtrDtor(_zv)
		ZVAL_COPY_VALUE(_zv, other_zv)
		break
	}
}
func ZEND_TRY_ASSIGN_VALUE_EX(zv *Zval, other_zv *Zval, strict ZendBool) {
	_ZEND_TRY_ASSIGN_VALUE_EX(zv, other_zv, strict, 0)
}
func ZEND_TRY_ASSIGN_COPY_EX(zv *Zval, other_zv *Zval, strict ZendBool) {
	other_zv.TryAddRefcount()
	ZEND_TRY_ASSIGN_VALUE_EX(zv, other_zv, strict)
}
func ZendTryArrayInitSize(zv *Zval, size uint32) *Zval {
	var arr *ZendArray = ZendNewArray(size)
	if zv.IsReference() {
		var ref *ZendReference = zv.GetRef()
		if ZEND_REF_HAS_TYPE_SOURCES(ref) {
			if ZendTryAssignTypedRefArr(ref, arr) != SUCCESS {
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
func ZendTryArrayInit(zv *Zval) *Zval { return ZendTryArrayInitSize(zv, 0) }
func Z_PARAM_PROLOGUE(deref int, separate int) {
	_i++
	ZEND_ASSERT(_i <= _min_num_args || _optional == 1)
	ZEND_ASSERT(_i > _min_num_args || _optional == 0)
	if _optional {
		if _i > _num_args {
			break
		}
	}
	_real_arg++
	_arg = _real_arg
	if deref != 0 {
		if _arg.IsReference() {
			_arg = Z_REFVAL_P(_arg)
		}
	}
	if separate != 0 {
		SEPARATE_ZVAL_NOREF(_arg)
	}
}
func ZendParseArgBool(arg *Zval, dest *ZendBool, is_null *ZendBool, check_null int) int {
	if check_null != 0 {
		*is_null = 0
	}
	if arg.IsTrue() {
		*dest = 1
	} else if arg.IsFalse() {
		*dest = 0
	} else if check_null != 0 && arg.IsNull() {
		*is_null = 1
		*dest = 0
	} else {
		return ZendParseArgBoolSlow(arg, dest)
	}
	return 1
}
func ZendParseArgLong(arg *Zval, dest *ZendLong, is_null *ZendBool, check_null int, cap int) int {
	if check_null != 0 {
		*is_null = 0
	}
	if arg.IsLong() {
		*dest = arg.GetLval()
	} else if check_null != 0 && arg.IsNull() {
		*is_null = 1
		*dest = 0
	} else if cap != 0 {
		return ZendParseArgLongCapSlow(arg, dest)
	} else {
		return ZendParseArgLongSlow(arg, dest)
	}
	return 1
}
func ZendParseArgDouble(arg *Zval, dest *float64, is_null *ZendBool, check_null int) int {
	if check_null != 0 {
		*is_null = 0
	}
	if arg.IsDouble() {
		*dest = arg.GetDval()
	} else if check_null != 0 && arg.IsNull() {
		*is_null = 1
		*dest = 0.0
	} else {
		return ZendParseArgDoubleSlow(arg, dest)
	}
	return 1
}
func ZendParseArgStr(arg *Zval, dest **ZendString, check_null int) int {
	if arg.IsString() {
		*dest = arg.GetStr()
	} else if check_null != 0 && arg.IsNull() {
		*dest = nil
	} else {
		return ZendParseArgStrSlow(arg, dest)
	}
	return 1
}
func ZendParseArgString(arg *Zval, dest **byte, dest_len *int, check_null int) int {
	var str *ZendString
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
func ZendParseArgPathStr(arg *Zval, dest **ZendString, check_null int) int {
	if ZendParseArgStr(arg, dest, check_null) == 0 || (*dest) != nil && CHECK_NULL_PATH(dest.GetVal(), dest.GetLen()) {
		return 0
	}
	return 1
}
func ZendParseArgPath(arg *Zval, dest **byte, dest_len *int, check_null int) int {
	var str *ZendString
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
func ZendParseArgArray(arg *Zval, dest **Zval, check_null int, or_object int) int {
	if arg.IsArray() || or_object != 0 && arg.IsObject() {
		*dest = arg
	} else if check_null != 0 && arg.IsNull() {
		*dest = nil
	} else {
		return 0
	}
	return 1
}
func ZendParseArgArrayHt(arg *Zval, dest **HashTable, check_null int, or_object int, separate int) int {
	if arg.IsArray() {
		*dest = arg.GetArr()
	} else if or_object != 0 && arg.IsObject() {
		if separate != 0 && Z_OBJ_P(arg).GetProperties() != nil && Z_OBJ_P(arg).GetProperties().GetRefcount() > 1 {
			if (Z_OBJ_P(arg).GetProperties().GetGcFlags() & IS_ARRAY_IMMUTABLE) == 0 {
				Z_OBJ_P(arg).GetProperties().DelRefcount()
			}
			Z_OBJ_P(arg).SetProperties(ZendArrayDup(Z_OBJ_P(arg).GetProperties()))
		}
		*dest = Z_OBJ_HT_P(arg).GetGetProperties()(arg)
	} else if check_null != 0 && arg.IsNull() {
		*dest = nil
	} else {
		return 0
	}
	return 1
}
func ZendParseArgObject(arg *Zval, dest **Zval, ce *ZendClassEntry, check_null int) int {
	if arg.IsObject() && (ce == nil || InstanceofFunction(Z_OBJCE_P(arg), ce) != 0) {
		*dest = arg
	} else if check_null != 0 && arg.IsNull() {
		*dest = nil
	} else {
		return 0
	}
	return 1
}
func ZendParseArgResource(arg *Zval, dest **Zval, check_null int) int {
	if arg.IsResource() {
		*dest = arg
	} else if check_null != 0 && arg.IsNull() {
		*dest = nil
	} else {
		return 0
	}
	return 1
}
func ZendParseArgFunc(arg *Zval, dest_fci *ZendFcallInfo, dest_fcc *ZendFcallInfoCache, check_null int, error **byte) int {
	if check_null != 0 && arg.IsNull() {
		dest_fci.SetSize(0)
		dest_fcc.SetFunctionHandler(nil)
		*error = nil
	} else if ZendFcallInfoInit(arg, 0, dest_fci, dest_fcc, nil, error) != SUCCESS {
		return 0
	}
	return 1
}
func ZendParseArgZvalDeref(arg *Zval, dest **Zval, check_null int) {
	if check_null != 0 && arg.IsNull() {
		*dest = nil
	} else {
		*dest = arg
	}
}
func _zendGetParametersArrayEx(param_count int, argument_array *Zval) int {
	var param_ptr *Zval
	var arg_count int
	param_ptr = ZEND_CALL_ARG(CurrEX(), 1)
	arg_count = ZEND_CALL_NUM_ARGS(CurrEX())
	if param_count > arg_count {
		return FAILURE
	}
	for b.PostDec(&param_count) > 0 {
		ZVAL_COPY_VALUE(argument_array, param_ptr)
		argument_array++
		param_ptr++
	}
	return SUCCESS
}
func ZendCopyParametersArray(param_count int, argument_array *Zval) int {
	var param_ptr *Zval
	var arg_count int
	param_ptr = ZEND_CALL_ARG(CurrEX(), 1)
	arg_count = ZEND_CALL_NUM_ARGS(CurrEX())
	if param_count > arg_count {
		return FAILURE
	}
	for b.PostDec(&param_count) > 0 {
		param_ptr.TryAddRefcount()
		argument_array.GetArr().NextIndexInsertNew(param_ptr)
		param_ptr++
	}
	return SUCCESS
}
func ZendWrongParamCount() {
	ZendInternalArgumentCountError(ZEND_ARG_USES_STRICT_TYPES(), "Wrong parameter count for %s()", GetActiveCalleeName())
}
func ZendGetTypeByConst(type_ ZendUchar) string {
	switch type_ {
	case IS_FALSE, IS_TRUE, _IS_BOOL:
		return "bool"
	case IS_LONG:
		return "int"
	case IS_DOUBLE:
		return "float"
	case IS_STRING:
		return "string"
	case IS_OBJECT:
		return "object"
	case IS_RESOURCE:
		return "resource"
	case IS_NULL:
		return "null"
	case IS_CALLABLE:
		return "callable"
	case IS_ITERABLE:
		return "iterable"
	case IS_ARRAY:
		return "array"
	case IS_VOID:
		return "void"
	case _IS_NUMBER:
		return "number"
	default:
		return "unknown"
	}
}
func ZendZvalTypeName(arg *Zval) string {
	ZVAL_DEREF(arg)
	return ZendGetTypeByConst(arg.GetType())
}
func ZendZvalGetType(arg *Zval) *ZendString {
	switch arg.GetType() {
	case IS_NULL:
		return ZSTR_KNOWN(ZEND_STR_NULL)
	case IS_FALSE:

	case IS_TRUE:
		return ZSTR_KNOWN(ZEND_STR_BOOLEAN)
	case IS_LONG:
		return ZSTR_KNOWN(ZEND_STR_INTEGER)
	case IS_DOUBLE:
		return ZSTR_KNOWN(ZEND_STR_DOUBLE)
	case IS_STRING:
		return ZSTR_KNOWN(ZEND_STR_STRING)
	case IS_ARRAY:
		return ZSTR_KNOWN(ZEND_STR_ARRAY)
	case IS_OBJECT:
		return ZSTR_KNOWN(ZEND_STR_OBJECT)
	case IS_RESOURCE:
		if ZendRsrcListGetRsrcType(arg.GetRes()) != nil {
			return ZSTR_KNOWN(ZEND_STR_RESOURCE)
		} else {
			return ZSTR_KNOWN(ZEND_STR_CLOSED_RESOURCE)
		}
	default:
		return nil
	}
}
