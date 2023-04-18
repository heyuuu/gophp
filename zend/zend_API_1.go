package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func _ZEND_TRY_ASSIGN_VALUE_EX(zv *types2.Zval, other_zv *types2.Zval, strict types2.ZendBool, is_ref int) {
	for {
		var _zv *types2.Zval = zv
		if is_ref != 0 || _zv.IsReference() {
			var ref *types2.ZendReference = _zv.Reference()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefZvalEx(ref, other_zv, strict)
				break
			}
			_zv = ref.GetVal()
		}
		// ZvalPtrDtor(_zv)
		_zv.CopyValueFrom(other_zv)
		break
	}
}
func ZEND_TRY_ASSIGN_VALUE_EX(zv *types2.Zval, other_zv *types2.Zval, strict types2.ZendBool) {
	_ZEND_TRY_ASSIGN_VALUE_EX(zv, other_zv, strict, 0)
}
func ZEND_TRY_ASSIGN_COPY_EX(zv *types2.Zval, other_zv *types2.Zval, strict types2.ZendBool) {
	// other_zv.TryAddRefcount()
	ZEND_TRY_ASSIGN_VALUE_EX(zv, other_zv, strict)
}
func ZendTryArrayInitSize(zv *types2.Zval, size uint32) *types2.Zval {
	var arr *types2.Array = types2.NewArray(size)
	if zv.IsReference() {
		var ref *types2.ZendReference = zv.Reference()
		if ZEND_REF_HAS_TYPE_SOURCES(ref) {
			if ZendTryAssignTypedRefArr(ref, arr) != types2.SUCCESS {
				return nil
			}
			return ref.GetVal()
		}
		zv = ref.GetVal()
	}
	// ZvalPtrDtor(zv)
	zv.SetArray(arr)
	return zv
}
func ZendTryArrayInit(zv *types2.Zval) *types2.Zval { return ZendTryArrayInitSize(zv, 0) }
func _zendGetParametersArrayEx(param_count int, argument_array *types2.Zval) int {
	var param_ptr *types2.Zval
	var arg_count int
	param_ptr = CurrEX().Arg(1)
	arg_count = CurrEX().NumArgs()
	if param_count > arg_count {
		return types2.FAILURE
	}
	for b.PostDec(&param_count) > 0 {
		argument_array.CopyValueFrom(param_ptr)
		argument_array++
		param_ptr++
	}
	return types2.SUCCESS
}
func ZendCopyParametersArray(param_count int, argument_array *types2.Zval) int {
	var param_ptr *types2.Zval
	var arg_count int
	param_ptr = CurrEX().Arg(1)
	arg_count = CurrEX().NumArgs()
	if param_count > arg_count {
		return types2.FAILURE
	}
	for b.PostDec(&param_count) > 0 {
		// param_ptr.TryAddRefcount()
		argument_array.Array().NextIndexInsertNew(param_ptr)
		param_ptr++
	}
	return types2.SUCCESS
}
func ZendWrongParamCount() {
	faults.InternalArgumentCountError(CurrEX().IsArgUseStrictTypes(), "Wrong parameter count for %s()", GetActiveCalleeName())
}
