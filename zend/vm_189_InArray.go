package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/operators"
)

func ZEND_IN_ARRAY_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var ht *types.Array = opline.Const2().Array()
	var result *types.Zval
	op1 = executeData.GetOp1(opline)
	if op1.IsString() {
		result = ht.KeyFind(op1.String())
	} else if opline.GetExtendedValue() != 0 {
		if op1.IsLong() {
			result = ht.IndexFind(op1.Long())
		} else {
			result = nil
		}
	} else if op1.IsSignFalse() {
		result = ht.KeyFind(types.NewString("").GetStr())
	} else {
		var key *types.String
		var key_tmp types.Zval
		var result_tmp types.Zval
		var val *types.Zval
		result = nil
		var __ht *types.Array = ht
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()

			key = _p.GetKey()
			val = _z
			key_tmp.SetString(key.GetStr())
			operators.CompareFunction(&result_tmp, op1, &key_tmp)
			if result_tmp.Long() == 0 {
				result = val
				break
			}
		}
	}
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != nil)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_IN_ARRAY_SPEC_TMP_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var ht *types.Array = opline.Const2().Array()
	var result *types.Zval
	op1 = executeData.GetVarOp1(opline)
	if op1.IsString() {
		result = ht.KeyFind(op1.String())
	} else if opline.GetExtendedValue() != 0 {
		if op1.IsLong() {
			result = ht.IndexFind(op1.Long())
		} else {
			result = nil
		}
	} else if op1.IsSignFalse() {
		result = ht.KeyFind(types.NewString("").GetStr())
	} else {
		var key *types.String
		var key_tmp types.Zval
		var result_tmp types.Zval
		var val *types.Zval
		result = nil
		var __ht *types.Array = ht
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()

			key = _p.GetKey()
			val = _z
			key_tmp.SetString(key.GetStr())
			operators.CompareFunction(&result_tmp, op1, &key_tmp)
			if result_tmp.Long() == 0 {
				result = val
				break
			}
		}
	}
	// ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != nil)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_IN_ARRAY_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var ht *types.Array = opline.Const2().Array()
	var result *types.Zval
	op1 = executeData.GetVarOp1(opline).DeRef()
	if op1.IsString() {
		result = ht.KeyFind(op1.String())
	} else if opline.GetExtendedValue() != 0 {
		if op1.IsLong() {
			result = ht.IndexFind(op1.Long())
		} else {
			result = nil
		}
	} else if op1.IsSignFalse() {
		result = ht.KeyFind(types.NewString("").GetStr())
	} else {
		var key *types.String
		var key_tmp types.Zval
		var result_tmp types.Zval
		var val *types.Zval
		result = nil
		var __ht *types.Array = ht
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()

			key = _p.GetKey()
			val = _z
			key_tmp.SetString(key.GetStr())
			operators.CompareFunction(&result_tmp, op1, &key_tmp)
			if result_tmp.Long() == 0 {
				result = val
				break
			}
		}
	}
	// ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result, 1)
	opline.Result().SetBool(result != nil)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_IN_ARRAY_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var ht *types.Array = opline.Const2().Array()
	var result *types.Zval
	op1 = executeData.GetCvOp1(opline)
	if op1.IsString() {
		result = ht.KeyFind(op1.String())
	} else if opline.GetExtendedValue() != 0 {
		if op1.IsLong() {
			result = ht.IndexFind(op1.Long())
		} else {
			result = nil
		}
	} else if op1.IsSignFalse() {
		result = ht.KeyFind(types.NewString("").GetStr())
	} else {
		var key *types.String
		var key_tmp types.Zval
		var result_tmp types.Zval
		var val *types.Zval
		result = nil
		var __ht *types.Array = ht
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()

			key = _p.GetKey()
			val = _z
			key_tmp.SetString(key.GetStr())
			operators.CompareFunction(&result_tmp, op1, &key_tmp)
			if result_tmp.Long() == 0 {
				result = val
				break
			}
		}
	}
	ZEND_VM_SMART_BRANCH(result != nil, 1)
	opline.Result().SetBool(result != nil)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
