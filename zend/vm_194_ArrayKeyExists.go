package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

func ZEND_ARRAY_KEY_EXISTS_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var key *types.Zval
	var subject *types.Zval
	var ht *types.Array
	var result types.ZvalType
	key = opline.Const1()
	subject = opline.Const2()
	if subject.IsArray() {
		ht = subject.Array()
		result = ZendArrayKeyExistsFast(ht, key, executeData)
	} else {
		result = ZendArrayKeyExistsSlow(subject, key, executeData)
	}
	ZEND_VM_SMART_BRANCH(result == types.IS_TRUE, 1)
	opline.Result().SetType(result)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var key *types.Zval
	var subject *types.Zval
	var ht *types.Array
	var result types.ZvalType
	key = opline.Const1()
	subject = opline.Op2().DeRef()
	if subject.IsArray() {
		ht = subject.Array()
		result = ZendArrayKeyExistsFast(ht, key, executeData)
	} else {
		result = ZendArrayKeyExistsSlow(subject, key, executeData)
	}
	ZEND_VM_SMART_BRANCH(result == types.IS_TRUE, 1)
	opline.Result().SetType(result)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var key *types.Zval
	var subject *types.Zval
	var ht *types.Array
	var result types.ZvalType
	key = opline.Const1()
	subject = opline.Op2().DeRef()
	if subject.IsArray() {
		ht = subject.Array()
		result = ZendArrayKeyExistsFast(ht, key, executeData)
	} else {
		result = ZendArrayKeyExistsSlow(subject, key, executeData)
	}
	ZEND_VM_SMART_BRANCH(result == types.IS_TRUE, 1)
	opline.Result().SetType(result)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var key *types.Zval
	var subject *types.Zval
	var ht *types.Array
	var result types.ZvalType
	key = opline.Op1()
	subject = opline.Const2()
	if subject.IsArray() {
		ht = subject.Array()
		result = ZendArrayKeyExistsFast(ht, key, executeData)
	} else {
		result = ZendArrayKeyExistsSlow(subject, key, executeData)
	}
	ZEND_VM_SMART_BRANCH(result == types.IS_TRUE, 1)
	opline.Result().SetType(result)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var key *types.Zval
	var subject *types.Zval
	var ht *types.Array
	var result types.ZvalType
	key = opline.Op1()
	subject = opline.Op2().DeRef()
	if subject.IsArray() {
		ht = subject.Array()
		result = ZendArrayKeyExistsFast(ht, key, executeData)
	} else {
		result = ZendArrayKeyExistsSlow(subject, key, executeData)
	}
	ZEND_VM_SMART_BRANCH(result == types.IS_TRUE, 1)
	opline.Result().SetType(result)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var key *types.Zval
	var subject *types.Zval
	var ht *types.Array
	var result types.ZvalType
	key = opline.Op1()
	subject = opline.Op2().DeRef()
	if subject.IsArray() {
		ht = subject.Array()
		result = ZendArrayKeyExistsFast(ht, key, executeData)
	} else {
		result = ZendArrayKeyExistsSlow(subject, key, executeData)
	}
	ZEND_VM_SMART_BRANCH(result == types.IS_TRUE, 1)
	opline.Result().SetType(result)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var key *types.Zval
	var subject *types.Zval
	var ht *types.Array
	var result types.ZvalType
	key = opline.Op1()
	subject = opline.Const2()
	if subject.IsArray() {
		ht = subject.Array()
		result = ZendArrayKeyExistsFast(ht, key, executeData)
	} else {
		result = ZendArrayKeyExistsSlow(subject, key, executeData)
	}
	ZEND_VM_SMART_BRANCH(result == types.IS_TRUE, 1)
	opline.Result().SetType(result)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var key *types.Zval
	var subject *types.Zval
	var ht *types.Array
	var result types.ZvalType
	key = opline.Op1()
	subject = opline.Op2().DeRef()
	if subject.IsArray() {
		ht = subject.Array()
		result = ZendArrayKeyExistsFast(ht, key, executeData)
	} else {
		result = ZendArrayKeyExistsSlow(subject, key, executeData)
	}
	ZEND_VM_SMART_BRANCH(result == types.IS_TRUE, 1)
	opline.Result().SetType(result)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var key *types.Zval
	var subject *types.Zval
	var ht *types.Array
	var result types.ZvalType
	key = opline.Op1()
	subject = opline.Op2().DeRef()
	if subject.IsArray() {
		ht = subject.Array()
		result = ZendArrayKeyExistsFast(ht, key, executeData)
	} else {
		result = ZendArrayKeyExistsSlow(subject, key, executeData)
	}
	ZEND_VM_SMART_BRANCH(result == types.IS_TRUE, 1)
	opline.Result().SetType(result)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
