package zend

import "github.com/heyuuu/gophp/php/types"

func ZEND_RECV_SPEC_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var arg_num uint32 = opline.GetOp1().GetNum()
	if arg_num > executeData.NumArgs() {
		ZendMissingArgError(executeData)
		return 0
	} else {
		var param *types.Zval = opline.Result()
		if ZendVerifyRecvArgType(executeData.GetFunc(), arg_num, param, nil, CACHE_ADDR(opline.GetOp2().GetNum())) == 0 {
			return 0
		}
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
