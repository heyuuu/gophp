package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

func arrayKeyExistsHelper(executeData *ZendExecuteData, key *types.Zval, subject *types.Zval) {
	var opline *types.ZendOp = executeData.GetOpline()
	var result types.ZvalType
	if subject.IsArray() {
		ht := subject.Array()
		result = ZendArrayKeyExistsFast(ht, key, executeData)
	} else {
		result = ZendArrayKeyExistsSlow(subject, key, executeData)
	}
	ZEND_VM_SMART_BRANCH(result == types.IsTrue, 1)

	//opline.Result().SetType(result)
	switch result {
	case types.IsTrue:
		opline.Result().SetTrue()
	case types.IsFalse:
		opline.Result().SetFalse()
	case types.IsNull:
		opline.Result().SetNull()
	default:
		panic("unreachable")
	}
}

func ZEND_ARRAY_KEY_EXISTS_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var key *types.Zval
	var subject *types.Zval
	key = opline.Const1()
	subject = opline.Const2()
	arrayKeyExistsHelper(executeData, key, subject)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var key *types.Zval
	var subject *types.Zval
	key = opline.Const1()
	subject = opline.Op2().DeRef()
	arrayKeyExistsHelper(executeData, key, subject)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var key *types.Zval
	var subject *types.Zval
	key = opline.Const1()
	subject = opline.Op2().DeRef()
	arrayKeyExistsHelper(executeData, key, subject)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var key *types.Zval
	var subject *types.Zval
	key = opline.Op1()
	subject = opline.Const2()
	arrayKeyExistsHelper(executeData, key, subject)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var key *types.Zval
	var subject *types.Zval
	key = opline.Op1()
	subject = opline.Op2().DeRef()
	arrayKeyExistsHelper(executeData, key, subject)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var key *types.Zval
	var subject *types.Zval
	key = opline.Op1()
	subject = opline.Op2().DeRef()
	arrayKeyExistsHelper(executeData, key, subject)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var key *types.Zval
	var subject *types.Zval
	key = opline.Op1()
	subject = opline.Const2()
	arrayKeyExistsHelper(executeData, key, subject)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var key *types.Zval
	var subject *types.Zval
	key = opline.Op1()
	subject = opline.Op2().DeRef()
	arrayKeyExistsHelper(executeData, key, subject)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var key *types.Zval
	var subject *types.Zval
	key = opline.Op1()
	subject = opline.Op2().DeRef()
	arrayKeyExistsHelper(executeData, key, subject)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
