package zend

import "github.com/heyuuu/gophp/zend/types"

func ZEND_ARRAY_KEY_EXISTS_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var key *types.Zval
	var subject *types.Zval
	var ht *types.Array
	var result uint32
	key = opline.Const1()
	subject = opline.Const2()
	if subject.IsArray() {
	array_key_exists_array:
		ht = subject.Array()
		result = ZendArrayKeyExistsFast(ht, key, opline, executeData)
	} else {
		result = ZendArrayKeyExistsSlow(subject, key, opline, executeData)
	}
	ZEND_VM_SMART_BRANCH(result == types.IS_TRUE, 1)
	opline.Result().SetTypeInfo(result)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var key *types.Zval
	var subject *types.Zval
	var ht *types.Array
	var result uint32
	key = opline.Const1()
	subject = opline.Op2()
	if subject.IsArray() {
	array_key_exists_array:
		ht = subject.Array()
		result = ZendArrayKeyExistsFast(ht, key, opline, executeData)
	} else {
		if subject.IsReference() {
			subject = types.Z_REFVAL_P(subject)
			if subject.IsArray() {
				goto array_key_exists_array
			}
		}
		result = ZendArrayKeyExistsSlow(subject, key, opline, executeData)
	}
	// ZvalPtrDtorNogc(free_op2)
	ZEND_VM_SMART_BRANCH(result == types.IS_TRUE, 1)
	opline.Result().SetTypeInfo(result)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var key *types.Zval
	var subject *types.Zval
	var ht *types.Array
	var result uint32
	key = opline.Const1()
	subject = opline.Op2()
	if subject.IsArray() {
	array_key_exists_array:
		ht = subject.Array()
		result = ZendArrayKeyExistsFast(ht, key, opline, executeData)
	} else {
		if subject.IsReference() {
			subject = types.Z_REFVAL_P(subject)
			if subject.IsArray() {
				goto array_key_exists_array
			}
		}
		result = ZendArrayKeyExistsSlow(subject, key, opline, executeData)
	}
	ZEND_VM_SMART_BRANCH(result == types.IS_TRUE, 1)
	opline.Result().SetTypeInfo(result)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var key *types.Zval
	var subject *types.Zval
	var ht *types.Array
	var result uint32
	key = opline.Op1()
	subject = opline.Const2()
	if subject.IsArray() {
	array_key_exists_array:
		ht = subject.Array()
		result = ZendArrayKeyExistsFast(ht, key, opline, executeData)
	} else {
		result = ZendArrayKeyExistsSlow(subject, key, opline, executeData)
	}
	// ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result == types.IS_TRUE, 1)
	opline.Result().SetTypeInfo(result)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var key *types.Zval
	var subject *types.Zval
	var ht *types.Array
	var result uint32
	key = opline.Op1()
	subject = opline.Op2()
	if subject.IsArray() {
	array_key_exists_array:
		ht = subject.Array()
		result = ZendArrayKeyExistsFast(ht, key, opline, executeData)
	} else {
		if subject.IsReference() {
			subject = types.Z_REFVAL_P(subject)
			if subject.IsArray() {
				goto array_key_exists_array
			}
		}
		result = ZendArrayKeyExistsSlow(subject, key, opline, executeData)
	}
	// ZvalPtrDtorNogc(free_op2)
	// ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result == types.IS_TRUE, 1)
	opline.Result().SetTypeInfo(result)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var key *types.Zval
	var subject *types.Zval
	var ht *types.Array
	var result uint32
	key = opline.Op1()
	subject = opline.Op2()
	if subject.IsArray() {
	array_key_exists_array:
		ht = subject.Array()
		result = ZendArrayKeyExistsFast(ht, key, opline, executeData)
	} else {
		if subject.IsReference() {
			subject = types.Z_REFVAL_P(subject)
			if subject.IsArray() {
				goto array_key_exists_array
			}
		}
		result = ZendArrayKeyExistsSlow(subject, key, opline, executeData)
	}
	// ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result == types.IS_TRUE, 1)
	opline.Result().SetTypeInfo(result)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var key *types.Zval
	var subject *types.Zval
	var ht *types.Array
	var result uint32
	key = opline.Op1()
	subject = opline.Const2()
	if subject.IsArray() {
	array_key_exists_array:
		ht = subject.Array()
		result = ZendArrayKeyExistsFast(ht, key, opline, executeData)
	} else {
		result = ZendArrayKeyExistsSlow(subject, key, opline, executeData)
	}
	ZEND_VM_SMART_BRANCH(result == types.IS_TRUE, 1)
	opline.Result().SetTypeInfo(result)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var key *types.Zval
	var subject *types.Zval
	var ht *types.Array
	var result uint32
	key = opline.Op1()
	subject = opline.Op2()
	if subject.IsArray() {
	array_key_exists_array:
		ht = subject.Array()
		result = ZendArrayKeyExistsFast(ht, key, opline, executeData)
	} else {
		if subject.IsReference() {
			subject = types.Z_REFVAL_P(subject)
			if subject.IsArray() {
				goto array_key_exists_array
			}
		}
		result = ZendArrayKeyExistsSlow(subject, key, opline, executeData)
	}
	// ZvalPtrDtorNogc(free_op2)
	ZEND_VM_SMART_BRANCH(result == types.IS_TRUE, 1)
	opline.Result().SetTypeInfo(result)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var key *types.Zval
	var subject *types.Zval
	var ht *types.Array
	var result uint32
	key = opline.Op1()
	subject = opline.Op2()
	if subject.IsArray() {
	array_key_exists_array:
		ht = subject.Array()
		result = ZendArrayKeyExistsFast(ht, key, opline, executeData)
	} else {
		if subject.IsReference() {
			subject = types.Z_REFVAL_P(subject)
			if subject.IsArray() {
				goto array_key_exists_array
			}
		}
		result = ZendArrayKeyExistsSlow(subject, key, opline, executeData)
	}
	ZEND_VM_SMART_BRANCH(result == types.IS_TRUE, 1)
	opline.Result().SetTypeInfo(result)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
