package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZEND_RETURN_BY_REF_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var retval_ptr *types2.Zval
	for {
		{

			/* Not supposed to happen, but we'll allow it */

			faults.Error(faults.E_NOTICE, "Only variable references should be returned by reference")
			retval_ptr = opline.Const1()
			if !(executeData.GetReturnValue()) {
			} else {
				executeData.GetReturnValue().
					SetNewRef(retval_ptr)
				{
					// retval_ptr.TryAddRefcount()
				}
			}
			break
		}
		retval_ptr = nil
		if executeData.GetReturnValue() {
			if retval_ptr.IsReference() {
				// 				retval_ptr.AddRefcount()
			} else {
				types2.ZVAL_MAKE_REF_EX(retval_ptr, 2)
			}
			executeData.GetReturnValue().
				SetReference(retval_ptr.Reference())
		}
		break
	}
	return zend_leave_helper_SPEC(executeData)
}
func ZEND_RETURN_BY_REF_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var retval_ptr *types2.Zval
	var free_op1 ZendFreeOp
	for {
		{

			/* Not supposed to happen, but we'll allow it */

			faults.Error(faults.E_NOTICE, "Only variable references should be returned by reference")
			retval_ptr = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
			if !(executeData.GetReturnValue()) {
				// ZvalPtrDtorNogc(free_op1)
			} else {
				if retval_ptr.IsReference() {
					types2.ZVAL_COPY_VALUE(executeData.GetReturnValue(), retval_ptr)
					break
				}
				executeData.GetReturnValue().
					SetNewRef(retval_ptr)
			}
			break
		}
		retval_ptr = nil
		{
			b.Assert(retval_ptr != EG__().GetUninitializedZval())
			if opline.GetExtendedValue() == ZEND_RETURNS_FUNCTION && !(retval_ptr.IsReference()) {
				faults.Error(faults.E_NOTICE, "Only variable references should be returned by reference")
				if executeData.GetReturnValue() {
					executeData.GetReturnValue().
						SetNewRef(retval_ptr)
				}
				break
			}
		}
		if executeData.GetReturnValue() {
			if retval_ptr.IsReference() {
				// 				retval_ptr.AddRefcount()
			} else {
				types2.ZVAL_MAKE_REF_EX(retval_ptr, 2)
			}
			executeData.GetReturnValue().
				SetReference(retval_ptr.Reference())
		}
		break
	}
	return zend_leave_helper_SPEC(executeData)
}
func ZEND_RETURN_BY_REF_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var retval_ptr *types2.Zval
	var free_op1 ZendFreeOp
	for {
		{

			/* Not supposed to happen, but we'll allow it */

			faults.Error(faults.E_NOTICE, "Only variable references should be returned by reference")
			retval_ptr = opline.Op1()
			if !(executeData.GetReturnValue()) {
				// ZvalPtrDtorNogc(free_op1)
			} else {
				if retval_ptr.IsReference() {
					types2.ZVAL_COPY_VALUE(executeData.GetReturnValue(), retval_ptr)
					break
				}
				executeData.GetReturnValue().
					SetNewRef(retval_ptr)
			}
			break
		}
		retval_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
		{
			b.Assert(retval_ptr != EG__().GetUninitializedZval())
			if opline.GetExtendedValue() == ZEND_RETURNS_FUNCTION && !(retval_ptr.IsReference()) {
				faults.Error(faults.E_NOTICE, "Only variable references should be returned by reference")
				if executeData.GetReturnValue() {
					executeData.GetReturnValue().
						SetNewRef(retval_ptr)
				} else {
					if free_op1 != nil {
						// ZvalPtrDtorNogc(free_op1)
					}
				}
				break
			}
		}
		if executeData.GetReturnValue() {
			if retval_ptr.IsReference() {
				// 				retval_ptr.AddRefcount()
			} else {
				types2.ZVAL_MAKE_REF_EX(retval_ptr, 2)
			}
			executeData.GetReturnValue().
				SetReference(retval_ptr.Reference())
		}
		if free_op1 != nil {
			// ZvalPtrDtorNogc(free_op1)
		}
		break
	}
	return zend_leave_helper_SPEC(executeData)
}
func ZEND_RETURN_BY_REF_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var retval_ptr *types2.Zval
	for {
		{

			/* Not supposed to happen, but we'll allow it */

			faults.Error(faults.E_NOTICE, "Only variable references should be returned by reference")
			retval_ptr = opline.Cv1OrUndef()
			if !(executeData.GetReturnValue()) {
			} else {
				if retval_ptr.IsReference() {
					types2.ZVAL_COPY_VALUE(executeData.GetReturnValue(), retval_ptr)
					break
				}
				executeData.GetReturnValue().
					SetNewRef(retval_ptr)
			}
			break
		}
		retval_ptr = _get_zval_ptr_cv_BP_VAR_W(opline.GetOp1().GetVar(), executeData)
		{
			b.Assert(retval_ptr != EG__().GetUninitializedZval())
			if opline.GetExtendedValue() == ZEND_RETURNS_FUNCTION && !(retval_ptr.IsReference()) {
				faults.Error(faults.E_NOTICE, "Only variable references should be returned by reference")
				if executeData.GetReturnValue() {
					executeData.GetReturnValue().
						SetNewRef(retval_ptr)
				}
				break
			}
		}
		if executeData.GetReturnValue() {
			if retval_ptr.IsReference() {
				// 				retval_ptr.AddRefcount()
			} else {
				types2.ZVAL_MAKE_REF_EX(retval_ptr, 2)
			}
			executeData.GetReturnValue().
				SetReference(retval_ptr.Reference())
		}
		break
	}
	return zend_leave_helper_SPEC(executeData)
}
