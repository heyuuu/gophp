package zend

func ZEND_USER_OPCODE_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var ret int
	//ret = ZendUserOpcodeHandlers[opline.GetOpcode()](executeData)
	ret = UserOpcodeHandlerT(nil)(executeData)
	opline = executeData.GetOpline()
	switch ret {
	case ZEND_USER_OPCODE_CONTINUE:
		return 0
	case ZEND_USER_OPCODE_RETURN:
		if (EX_CALL_INFO() & ZEND_CALL_GENERATOR) != 0 {
			var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
			ZendGeneratorClose(generator, 1)
			return -1
		} else {
			return zend_leave_helper_SPEC(executeData)
		}
		fallthrough
	case ZEND_USER_OPCODE_ENTER:
		return 1
	case ZEND_USER_OPCODE_LEAVE:
		return 2
	case ZEND_USER_OPCODE_DISPATCH:
		ZEND_VM_DISPATCH(executeData, opline.GetOpcode(), opline)
		fallthrough
	default:
		ZEND_VM_DISPATCH(executeData, zend_uchar(ret&0xff), opline)
	}
}
