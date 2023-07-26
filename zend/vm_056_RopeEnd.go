package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/operators"
)

func ZEND_ROPE_END_SPEC_TMP_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var rope []*types.String
	var var_ *types.Zval
	var i uint32
	rope = ([]*types.String)(opline.Op1())
	{
		var_ = opline.Const2()
		rope[opline.GetExtendedValue()] = var_.String()
	}

	var retStr string
	for i = 0; i <= opline.GetExtendedValue(); i++ {
		retStr += rope[i].GetStr()
	}
	opline.Result().SetString(retStr)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_ROPE_END_SPEC_TMP_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var rope []*types.String
	var var_ *types.Zval
	var i uint32
	rope = ([]*types.String)(opline.Op1())
	{
		var_ = opline.Op2()
		if var_.IsString() {
			rope[opline.GetExtendedValue()] = var_.String().Copy()
		} else {
			if var_.IsUndef() {
				ZVAL_UNDEFINED_OP2(executeData)
			}
			rope[opline.GetExtendedValue()] = operators.ZvalGetString(var_)
			if EG__().GetException() != nil {
				opline.Result().SetUndef()
				return 0
			}
		}
	}

	var retStr string
	for i = 0; i <= opline.GetExtendedValue(); i++ {
		retStr += rope[i].GetStr()
	}
	opline.Result().SetString(retStr)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_ROPE_END_SPEC_TMP_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var rope []*types.String
	var var_ *types.Zval
	var i uint32
	rope = ([]*types.String)(opline.Op1())

	{
		var_ = opline.Op2()
		if var_.IsString() {
			rope[opline.GetExtendedValue()] = var_.String().Copy()
		} else {
			if var_.IsUndef() {
				ZVAL_UNDEFINED_OP2(executeData)
			}
			rope[opline.GetExtendedValue()] = operators.ZvalGetString(var_)
			if EG__().GetException() != nil {
				opline.Result().SetUndef()
				return 0
			}
		}
	}

	var retStr string
	for i = 0; i <= opline.GetExtendedValue(); i++ {
		retStr += rope[i].GetStr()
	}
	opline.Result().SetString(retStr)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
