// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/faults"
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
	faults.ZendInternalArgumentCountError(CurrEX().IsArgUseStrictTypes(), "Wrong parameter count for %s()", GetActiveCalleeName())
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
		return types.ZSTR_NULL
	case types.IS_FALSE:

	case types.IS_TRUE:
		return types.ZSTR_BOOLEAN
	case types.IS_LONG:
		return types.ZSTR_INTEGER
	case types.IS_DOUBLE:
		return types.ZSTR_DOUBLE
	case types.IS_STRING:
		return types.ZSTR_STRING
	case types.IS_ARRAY:
		return types.ZSTR_ARRAY
	case types.IS_OBJECT:
		return types.ZSTR_OBJECT
	case types.IS_RESOURCE:
		if ZendRsrcListGetRsrcType(arg.GetRes()) != nil {
			return types.ZSTR_RESOURCE
		} else {
			return types.ZSTR_CLOSED_RESOURCE
		}
	default:
		return nil
	}
}
