package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZEND_THROW_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var value *types.Zval
	value = opline.Const1()
	for {
		{
			faults.ThrowError(nil, "Can only throw objects")
			return 0
		}
		break
	}
	EG__().ExceptionSave()
	{
		// value.TryAddRefcount()
	}
	faults.ThrowExceptionObject(value)
	EG__().ExceptionRestore()
	return 0
}
func ZEND_THROW_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var value *types.Zval
	var free_op1 ZendFreeOp
	value = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	for {
		if !value.IsObject() {
			if value.IsRef() {
				value = types.Z_REFVAL_P(value)
				if value.IsObject() {
					break
				}
			}
			if value.IsUndef() {
				ZVAL_UNDEFINED_OP1(executeData)
				if EG__().HasException() {
					return 0
				}
			}
			faults.ThrowError(nil, "Can only throw objects")
			// ZvalPtrDtorNogc(free_op1)
			return 0
		}
		break
	}
	EG__().ExceptionSave()
	faults.ThrowExceptionObject(value)
	EG__().ExceptionRestore()
	return 0
}
func ZEND_THROW_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var value *types.Zval
	var free_op1 ZendFreeOp
	value = opline.Op1()
	for {
		if !value.IsObject() {
			if value.IsRef() {
				value = types.Z_REFVAL_P(value)
				if value.IsObject() {
					break
				}
			}
			if value.IsUndef() {
				ZVAL_UNDEFINED_OP1(executeData)
				if EG__().HasException() {
					return 0
				}
			}
			faults.ThrowError(nil, "Can only throw objects")
			// ZvalPtrDtorNogc(free_op1)
			return 0
		}
		break
	}
	EG__().ExceptionSave()
	faults.ThrowExceptionObject(value)
	EG__().ExceptionRestore()
	return 0
}
func ZEND_THROW_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var value *types.Zval
	value = opline.Op1()
	for {
		if !value.IsObject() {
			if value.IsRef() {
				value = types.Z_REFVAL_P(value)
				if value.IsObject() {
					break
				}
			}
			if value.IsUndef() {
				ZVAL_UNDEFINED_OP1(executeData)
				if EG__().HasException() {
					return 0
				}
			}
			faults.ThrowError(nil, "Can only throw objects")
			return 0
		}
		break
	}
	EG__().ExceptionSave()
	faults.ThrowExceptionObject(value)
	EG__().ExceptionRestore()
	return 0
}
