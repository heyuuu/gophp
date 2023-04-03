package zend

func ZEND_GET_CALLED_CLASS_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	if executeData.GetThis().IsObject() {
		opline.GetResultZval().SetStringCopy(types.Z_OBJCE(executeData.GetThis()).GetName())
	} else if executeData.GetThis().GetCe() != nil {
		opline.GetResultZval().SetStringCopy(types.Z_CE(executeData.GetThis()).GetName())
	} else {
		opline.GetResultZval().SetFalse()
		if !(executeData.GetFunc().common.scope) {
			faults.Error(faults.E_WARNING, "get_called_class() called from outside a class")
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
