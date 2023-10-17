package zend

import (
	"fmt"
	b "github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func _ZEND_TRY_ASSIGN_VALUE_EX(zv *types.Zval, other_zv *types.Zval, strict bool, is_ref int) {
	for {
		var _zv *types.Zval = zv
		if is_ref != 0 || _zv.IsRef() {
			var ref *types.Reference = _zv.Ref()
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
func ZEND_TRY_ASSIGN_VALUE_EX(zv *types.Zval, other_zv *types.Zval, strict bool) {
	_ZEND_TRY_ASSIGN_VALUE_EX(zv, other_zv, strict, 0)
}
func ZEND_TRY_ASSIGN_COPY_EX(zv *types.Zval, other_zv *types.Zval, strict bool) {
	// other_zv.TryAddRefcount()
	ZEND_TRY_ASSIGN_VALUE_EX(zv, other_zv, strict)
}
func ZendTryArrayInitSize(zv *types.Zval, size uint32) *types.Zval {
	var arr *types.Array = types.NewArrayCap(size)
	if zv.IsRef() {
		var ref *types.Reference = zv.Ref()
		if ZEND_REF_HAS_TYPE_SOURCES(ref) {
			if ZendTryAssignTypedRefArr(ref, arr) != types.SUCCESS {
				return nil
			}
			return ref.GetVal()
		}
		zv = ref.GetVal()
	}
	zv.SetArray(arr)
	return zv
}
func ZendTryArrayInit(zv *types.Zval) *types.Zval { return ZendTryArrayInitSize(zv, 0) }
func ZendCopyParametersArray(param_count int, argument_array *types.Zval) int {
	var param_ptr *types.Zval
	var arg_count int
	param_ptr = CurrEX().Arg(1)
	arg_count = CurrEX().NumArgs()
	if param_count > arg_count {
		return types.FAILURE
	}
	for b.PostDec(&param_count) > 0 {
		// param_ptr.TryAddRefcount()
		argument_array.Array().AppendNew(param_ptr)
		param_ptr++
	}
	return types.SUCCESS
}
func ZendWrongParamCount() {
	faults.InternalArgumentCountError(CurrEX().IsArgUseStrictTypes(), fmt.Sprintf("Wrong parameter count for %s()", CurrEX().CalleeName()))
}
