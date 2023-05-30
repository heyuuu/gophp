package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

func ZEND_RECV_VARIADIC_SPEC_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var arg_num uint32 = opline.GetOp1().GetNum()
	var arg_count uint32 = uint32(executeData.NumArgs())
	var params *types.Zval
	params = opline.Result()
	if arg_num <= arg_count {
		ArrayInitSize(params, arg_count-arg_num+1)
		paramIdx := executeData.GetFunc().GetOpArray().GetLastVar() + int(executeData.GetFunc().GetOpArray().T)
		if (executeData.GetFunc().GetOpArray().GetFnFlags() & types.AccHasTypeHints) != 0 {
			ZEND_ADD_CALL_FLAG(executeData, ZEND_CALL_FREE_EXTRA_ARGS)

			for i := 0; i < int(arg_count-arg_num); i++ {
				param := executeData.VarNum(paramIdx + i)

				ZendVerifyVariadicArgType(executeData.GetFunc(), arg_num, param, nil, CACHE_ADDR(opline.GetOp2().GetNum()))

				params.Array().Append(param)
			}
		} else {
			for i := 0; i < int(arg_count-arg_num); i++ {
				param := executeData.VarNum(paramIdx + i)
				params.Array().Append(param)
			}
		}
	} else {
		params.SetEmptyArray()
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
