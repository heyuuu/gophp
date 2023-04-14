package zend

func ZEND_BIND_LEXICAL_SPEC_TMP_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var closure *types.Zval
	var var_ *types.Zval
	closure = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	if (opline.GetExtendedValue() & ZEND_BIND_REF) != 0 {

		/* By-ref binding */

		var_ = _get_zval_ptr_cv_BP_VAR_W(opline.GetOp2().GetVar(), executeData)
		if var_.IsReference() {
			var_.AddRefcount()
		} else {
			types.ZVAL_MAKE_REF_EX(var_, 2)
		}
	} else {
		var_ = opline.Op2()
		if var_.IsUndef() && (opline.GetExtendedValue()&ZEND_BIND_IMPLICIT) == 0 {
			var_ = ZVAL_UNDEFINED_OP2(executeData)
			if EG__().GetException() != nil {
				return 0
			}
		}
		var_ = types.ZVAL_DEREF(var_)
		var_.TryAddRefcount()
	}
	ZendClosureBindVarEx(closure, opline.GetExtendedValue() & ^(ZEND_BIND_REF|ZEND_BIND_IMPLICIT), var_)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
