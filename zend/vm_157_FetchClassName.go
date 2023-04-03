package zend

func ZEND_FETCH_CLASS_NAME_SPEC_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var fetch_type uint32
	var called_scope *types.ClassEntry
	var scope *types.ClassEntry
	var opline *ZendOp = executeData.GetOpline()
	fetch_type = opline.GetOp1().GetNum()
	scope = executeData.GetFunc().op_array.scope
	if scope == nil {
		faults.ThrowError(nil, "Cannot use \"%s\" when no class scope is active", b.Cond(b.Cond(fetch_type == ZEND_FETCH_CLASS_SELF, "self", fetch_type == ZEND_FETCH_CLASS_PARENT), "parent", "static"))
		opline.GetResultZval().SetUndef()
		return 0
	}
	switch fetch_type {
	case ZEND_FETCH_CLASS_SELF:
		opline.GetResultZval().SetStringCopy(scope.GetName())
	case ZEND_FETCH_CLASS_PARENT:
		if scope.GetParent() == nil {
			faults.ThrowError(nil, "Cannot use \"parent\" when current class scope has no parent")
			opline.GetResultZval().SetUndef()
			return 0
		}
		opline.GetResultZval().SetStringCopy(scope.GetParent().name)
	case ZEND_FETCH_CLASS_STATIC:
		if executeData.GetThis().IsObject() {
			called_scope = types.Z_OBJCE(executeData.GetThis())
		} else {
			called_scope = executeData.GetThis().GetCe()
		}
		opline.GetResultZval().SetStringCopy(called_scope.GetName())
	default:

	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
