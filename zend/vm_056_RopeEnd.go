package zend

import "github.com/heyuuu/gophp/php/types"

func ZEND_ROPE_END_SPEC_TMP_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var rope **types.String
	var var_ *types.Zval
	var ret *types.Zval
	var i uint32
	var len_ int = 0
	var target *byte
	rope = (**types.String)(opline.Op1())
	{
		var_ = opline.Const2()
		rope[opline.GetExtendedValue()] = var_.GetStr()

		// var_.TryAddRefcount()

	}

	for i = 0; i <= opline.GetExtendedValue(); i++ {
		len_ += rope[i].GetLen()
	}
	ret = opline.Result()
	ret.SetString(types.ZendStringAlloc(len_, 0))
	target = ret.GetStr().GetVal()
	for i = 0; i <= opline.GetExtendedValue(); i++ {
		memcpy(target, rope[i].GetVal(), rope[i].GetLen())
		target += rope[i].GetLen()
		// types.ZendStringReleaseEx(rope[i], 0)
	}
	*target = '0'
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_ROPE_END_SPEC_TMP_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var rope **types.String
	var var_ *types.Zval
	var ret *types.Zval
	var i uint32
	var len_ int = 0
	var target *byte
	rope = (**types.String)(opline.Op1())

	{
		var_ = opline.Op2()
		if var_.IsString() {
			{
				rope[opline.GetExtendedValue()] = var_.GetStr().Copy()
			}

		} else {
			if var_.IsUndef() {
				ZVAL_UNDEFINED_OP2(executeData)
			}
			rope[opline.GetExtendedValue()] = ZvalGetString(var_)
			// ZvalPtrDtorNogc(free_op2)
			if EG__().GetException() != nil {
				for i = 0; i <= opline.GetExtendedValue(); i++ {
					// types.ZendStringReleaseEx(rope[i], 0)
				}
				opline.Result().SetUndef()
				return 0
			}
		}
	}
	for i = 0; i <= opline.GetExtendedValue(); i++ {
		len_ += rope[i].GetLen()
	}
	ret = opline.Result()
	ret.SetString(types.ZendStringAlloc(len_, 0))
	target = ret.GetStr().GetVal()
	for i = 0; i <= opline.GetExtendedValue(); i++ {
		memcpy(target, rope[i].GetVal(), rope[i].GetLen())
		target += rope[i].GetLen()
		// types.ZendStringReleaseEx(rope[i], 0)
	}
	*target = '0'
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_ROPE_END_SPEC_TMP_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var rope **types.String
	var var_ *types.Zval
	var ret *types.Zval
	var i uint32
	var len_ int = 0
	var target *byte
	rope = (**types.String)(opline.Op1())

	{
		var_ = opline.Op2()
		if var_.IsString() {
			{
				rope[opline.GetExtendedValue()] = var_.GetStr().Copy()
			}

		} else {
			if var_.IsUndef() {
				ZVAL_UNDEFINED_OP2(executeData)
			}
			rope[opline.GetExtendedValue()] = ZvalGetString(var_)
			if EG__().GetException() != nil {
				for i = 0; i <= opline.GetExtendedValue(); i++ {
					// types.ZendStringReleaseEx(rope[i], 0)
				}
				opline.Result().SetUndef()
				return 0
			}
		}
	}
	for i = 0; i <= opline.GetExtendedValue(); i++ {
		len_ += rope[i].GetLen()
	}
	ret = opline.Result()
	ret.SetString(types.ZendStringAlloc(len_, 0))
	target = ret.GetStr().GetVal()
	for i = 0; i <= opline.GetExtendedValue(); i++ {
		memcpy(target, rope[i].GetVal(), rope[i].GetLen())
		target += rope[i].GetLen()
		// types.ZendStringReleaseEx(rope[i], 0)
	}
	*target = '0'
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
