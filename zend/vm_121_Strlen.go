package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/zpp"
)

func ZEND_STRLEN_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var value *types.Zval
	value = opline.Const1()
	if value.IsString() {
		opline.Result().SetLong(len(value.String()))
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else {
		var strict bool
		strict = executeData.IsCallUseStrictTypes()
		for {
			if strict == 0 {
				var str *types.String
				var tmp types.Zval
				types.ZVAL_COPY(&tmp, value)
				if zpp.ZendParseArgStrWeak(&tmp, &str) != 0 {
					opline.Result().SetLong(str.GetLen())
					// ZvalPtrDtor(&tmp)
					break
				}
				// ZvalPtrDtor(&tmp)
			}
			if EG__().NoException() {
				faults.InternalTypeError(strict, fmt.Sprintf("strlen() expects parameter 1 to be string, %s given", types.ZendGetTypeByConst(value.Type())))
			}
			opline.Result().SetNull()
			break
		}
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_STRLEN_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var value *types.Zval
	var free_op1 ZendFreeOp
	value = opline.Op1()
	if value.IsString() {
		opline.Result().SetLong(len(value.String()))
		// ZvalPtrDtorNogc(free_op1)
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else {
		var strict bool
		if value.IsRef() {
			value = types.Z_REFVAL_P(value)
			if value.IsString() {
				opline.Result().SetLong(len(value.String()))
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
				var str *types.String
				var tmp types.Zval
				types.ZVAL_COPY(&tmp, value)
				if zpp.ZendParseArgStrWeak(&tmp, &str) != 0 {
					opline.Result().SetLong(str.GetLen())
					// ZvalPtrDtor(&tmp)
					break
				}
				// ZvalPtrDtor(&tmp)
			}
			if EG__().NoException() {
				faults.InternalTypeError(strict, fmt.Sprintf("strlen() expects parameter 1 to be string, %s given", types.ZendGetTypeByConst(value.Type())))
			}
			opline.Result().SetNull()
			break
		}
	}
	// ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_STRLEN_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var value *types.Zval
	value = opline.Op1()
	if value.IsString() {
		opline.Result().SetLong(len(value.String()))
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else {
		var strict bool
		if value.IsRef() {
			value = types.Z_REFVAL_P(value)
			if value.IsString() {
				opline.Result().SetLong(len(value.String()))
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		}
		if value.IsUndef() {
			value = ZVAL_UNDEFINED_OP1(executeData)
		}
		strict = executeData.IsCallUseStrictTypes()
		for {
			if strict == 0 {
				var str *types.String
				var tmp types.Zval
				types.ZVAL_COPY(&tmp, value)
				if zpp.ZendParseArgStrWeak(&tmp, &str) != 0 {
					opline.Result().SetLong(str.GetLen())
					// ZvalPtrDtor(&tmp)
					break
				}
				// ZvalPtrDtor(&tmp)
			}
			if EG__().NoException() {
				faults.InternalTypeError(strict, fmt.Sprintf("strlen() expects parameter 1 to be string, %s given", types.ZendGetTypeByConst(value.Type())))
			}
			opline.Result().SetNull()
			break
		}
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
