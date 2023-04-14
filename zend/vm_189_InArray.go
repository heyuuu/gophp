package zend

import "github.com/heyuuu/gophp/zend/types"

func ZEND_IN_ARRAY_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var ht *types.Array = opline.Const2().Array()
	var result *types.Zval
	op1 = opline.Const1()
	if op1.IsString() {
		result = ht.KeyFind(op1.GetStrVal())
	} else if opline.GetExtendedValue() != 0 {
		if op1.IsLong() {
			result = ht.IndexFind(op1.Long())
		} else {
			result = nil
		}
	} else if op1.GetType() <= types.IS_FALSE {
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
			key_tmp.SetString(key)
			CompareFunction(&result_tmp, op1, &key_tmp)
			if result_tmp.Long() == 0 {
				result = val
				break
			}
		}
	}
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(opline.Result(), result != nil)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_IN_ARRAY_SPEC_TMP_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var ht *types.Array = opline.Const2().Array()
	var result *types.Zval
	op1 = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	if op1.IsString() {
		result = ht.KeyFind(op1.GetStrVal())
	} else if opline.GetExtendedValue() != 0 {
		if op1.IsLong() {
			result = ht.IndexFind(op1.Long()())
		} else {
			result = nil
		}
	} else if op1.GetType() <= types.IS_FALSE {
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
			key_tmp.SetString(key)
			CompareFunction(&result_tmp, op1, &key_tmp)
			if result_tmp.Long()() == 0 {
				result = val
				break
			}
		}
	}
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(opline.Result(), result != nil)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_IN_ARRAY_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var ht *types.Array = opline.Const2().Array()
	var result *types.Zval
	op1 = _getZvalPtrVarDeref(opline.GetOp1().GetVar(), &free_op1, executeData)
	if op1.IsString() {
		result = ht.KeyFind(op1.GetStrVal())
	} else if opline.GetExtendedValue() != 0 {
		if op1.IsLong() {
			result = ht.IndexFind(op1.Long()())
		} else {
			result = nil
		}
	} else if op1.GetType() <= types.IS_FALSE {
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
			key_tmp.SetString(key)
			CompareFunction(&result_tmp, op1, &key_tmp)
			if result_tmp.Long()() == 0 {
				result = val
				break
			}
		}
	}
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(opline.Result(), result != nil)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_IN_ARRAY_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var ht *types.Array = opline.Const2().Array()
	var result *types.Zval
	op1 = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	if op1.IsString() {
		result = ht.KeyFind(op1.GetStrVal())
	} else if opline.GetExtendedValue() != 0 {
		if op1.IsLong() {
			result = ht.IndexFind(op1.Long()())
		} else {
			result = nil
		}
	} else if op1.GetType() <= types.IS_FALSE {
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
			key_tmp.SetString(key)
			CompareFunction(&result_tmp, op1, &key_tmp)
			if result_tmp.Long()() == 0 {
				result = val
				break
			}
		}
	}
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(opline.Result(), result != nil)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
