package zend

import (
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/zpp"
)

func ZEND_STRLEN_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types2.Zval
	value = opline.Const1()
	if value.IsString() {
		opline.Result().SetLong(value.String().GetLen())
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else {
		var strict types2.ZendBool
		strict = executeData.IsCallUseStrictTypes()
		for {
			if strict == 0 {
				var str *types2.String
				var tmp types2.Zval
				types2.ZVAL_COPY(&tmp, value)
				if zpp.ZendParseArgStrWeak(&tmp, &str) != 0 {
					opline.Result().SetLong(str.GetLen())
					// ZvalPtrDtor(&tmp)
					break
				}
				// ZvalPtrDtor(&tmp)
			}
			if EG__().GetException() == nil {
				faults.InternalTypeError(strict, "strlen() expects parameter 1 to be string, %s given", types2.ZendGetTypeByConst(value.GetType()))
			}
			opline.Result().SetNull()
			break
		}
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_STRLEN_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types2.Zval
	var free_op1 ZendFreeOp
	value = opline.Op1()
	if value.IsString() {
		opline.Result().SetLong(value.String().GetLen())
		// ZvalPtrDtorNogc(free_op1)
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else {
		var strict types2.ZendBool
		if value.IsReference() {
			value = types2.Z_REFVAL_P(value)
			if value.IsString() {
				opline.Result().SetLong(value.String().GetLen())
				// ZvalPtrDtorNogc(free_op1)
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		}
		if value.IsUndef() {
			value = ZVAL_UNDEFINED_OP1(executeData)
		}
		strict = executeData.IsCallUseStrictTypes()
		for {
			if strict == 0 {
				var str *types2.String
				var tmp types2.Zval
				types2.ZVAL_COPY(&tmp, value)
				if zpp.ZendParseArgStrWeak(&tmp, &str) != 0 {
					opline.Result().SetLong(str.GetLen())
					// ZvalPtrDtor(&tmp)
					break
				}
				// ZvalPtrDtor(&tmp)
			}
			if EG__().GetException() == nil {
				faults.InternalTypeError(strict, "strlen() expects parameter 1 to be string, %s given", types2.ZendGetTypeByConst(value.GetType()))
			}
			opline.Result().SetNull()
			break
		}
	}
	// ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_STRLEN_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types2.Zval
	value = opline.Op1()
	if value.IsString() {
		opline.Result().SetLong(value.String().GetLen())
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else {
		var strict types2.ZendBool
		if value.IsReference() {
			value = types2.Z_REFVAL_P(value)
			if value.IsString() {
				opline.Result().SetLong(value.String().GetLen())
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		}
		if value.IsUndef() {
			value = ZVAL_UNDEFINED_OP1(executeData)
		}
		strict = executeData.IsCallUseStrictTypes()
		for {
			if strict == 0 {
				var str *types2.String
				var tmp types2.Zval
				types2.ZVAL_COPY(&tmp, value)
				if zpp.ZendParseArgStrWeak(&tmp, &str) != 0 {
					opline.Result().SetLong(str.GetLen())
					// ZvalPtrDtor(&tmp)
					break
				}
				// ZvalPtrDtor(&tmp)
			}
			if EG__().GetException() == nil {
				faults.InternalTypeError(strict, "strlen() expects parameter 1 to be string, %s given", types2.ZendGetTypeByConst(value.GetType()))
			}
			opline.Result().SetNull()
			break
		}
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
