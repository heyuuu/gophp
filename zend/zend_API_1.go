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
	var arr *types.ZendArray = types.ZendNewArray(size)
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
	faults.InternalArgumentCountError(CurrEX().IsArgUseStrictTypes(), "Wrong parameter count for %s()", GetActiveCalleeName())
}
