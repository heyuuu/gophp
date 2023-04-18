package zend

import "github.com/heyuuu/gophp/php/types"

func ZEND_FAST_CONCAT_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = executeData.GetOp1(opline)
	op2 = executeData.GetOp2(opline)

	var op1_str *types.String = op1.GetStr()
	var op2_str *types.String = op2.GetStr()
	var str *types.String

	str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
	memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
	memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
	opline.Result().SetString(str)

	return ZEND_VM_NEXT_OPCODE(executeData, opline)
	op1_str = op1.GetStr()
	op2_str = op2.GetStr()

	str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
	memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
	memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
	opline.Result().SetString(str)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FAST_CONCAT_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var op1_str *types.String
	var op2_str *types.String
	var str *types.String
	op1 = executeData.GetOp1(opline)
	op2 = opline.Op2()
	if op2.IsString() {
		var op1_str *types.String = op1.GetStr()
		var op2_str *types.String = op2.GetStr()
		var str *types.String

		if op2_str.GetLen() == 0 {
			{
				opline.Result().SetStringCopy(op1_str)
			}

			{
				// types.ZendStringReleaseEx(op2_str, 0)
			}
		} else {
			str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
			memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
			memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
			opline.Result().SetString(str)
			{
				// types.ZendStringReleaseEx(op2_str, 0)
			}
		}
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	{
		op1_str = op1.GetStr()
	}

	if op2.IsString() {
		op2_str = op2.GetStr().Copy()
	} else {
		if op2.IsUndef() {
			ZVAL_UNDEFINED_OP2(executeData)
		}
		op2_str = ZvalGetStringFunc(op2)
	}
	for {
		{
			if op2_str.GetLen() == 0 {
				{
					if op1.IsRefcounted() {
						// 						op1_str.AddRefcount()
					}
				}
				opline.Result().SetString(op1_str)
				// types.ZendStringReleaseEx(op2_str, 0)
				break
			}
		}
		str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
		memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
		memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
		opline.Result().SetString(str)
		{
			// types.ZendStringReleaseEx(op2_str, 0)
		}
		break
	}
	// ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FAST_CONCAT_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var op1_str *types.String
	var op2_str *types.String
	var str *types.String
	op1 = executeData.GetOp1(opline)
	op2 = opline.Op2()
	if op2.IsString() {
		var op1_str *types.String = op1.GetStr()
		var op2_str *types.String = op2.GetStr()
		var str *types.String

		if op2_str.GetLen() == 0 {
			{
				opline.Result().SetStringCopy(op1_str)
			}

			{
				// types.ZendStringReleaseEx(op2_str, 0)
			}
		} else {
			str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
			memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
			memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
			opline.Result().SetString(str)
			{
				// types.ZendStringReleaseEx(op2_str, 0)
			}
		}
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
	{
		op1_str = op1.GetStr()
	}

	if op2.IsString() {
		op2_str = op2.GetStr().Copy()
	} else {
		if op2.IsUndef() {
			ZVAL_UNDEFINED_OP2(executeData)
		}
		op2_str = ZvalGetStringFunc(op2)
	}
	for {
		{
			if op2_str.GetLen() == 0 {
				{
					if op1.IsRefcounted() {
						// 						op1_str.AddRefcount()
					}
				}
				opline.Result().SetString(op1_str)
				// types.ZendStringReleaseEx(op2_str, 0)
				break
			}
		}
		str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
		memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
		memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
		opline.Result().SetString(str)
		{
			// types.ZendStringReleaseEx(op2_str, 0)
		}
		break
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FAST_CONCAT_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var op1_str *types.String
	var op2_str *types.String
	var str *types.String
	op1 = opline.Op1()
	op2 = executeData.GetOp2(opline)
	if op1.IsString() {
		var op1_str *types.String = op1.GetStr()
		var op2_str *types.String = op2.GetStr()
		var str *types.String
		if op1_str.GetLen() == 0 {
			{
				opline.Result().SetStringCopy(op2_str)
			}

			{
				// types.ZendStringReleaseEx(op1_str, 0)
			}
		} else {
			str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
			memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
			memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
			opline.Result().SetString(str)
			{
				// types.ZendStringReleaseEx(op1_str, 0)
			}
		}
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}

	if op1.IsString() {
		op1_str = op1.GetStr().Copy()
	} else {
		if op1.IsUndef() {
			ZVAL_UNDEFINED_OP1(executeData)
		}
		op1_str = ZvalGetStringFunc(op1)
	}
	{
		op2_str = op2.GetStr()
	}

	for {
		{
			if op1_str.GetLen() == 0 {
				{
					if op2.IsRefcounted() {
						// 						op2_str.AddRefcount()
					}
				}
				opline.Result().SetString(op2_str)
				// types.ZendStringReleaseEx(op1_str, 0)
				break
			}
		}
		str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
		memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
		memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
		opline.Result().SetString(str)
		{
			// types.ZendStringReleaseEx(op1_str, 0)
		}
		break
	}
	// ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FAST_CONCAT_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var op1_str *types.String
	var op2_str *types.String
	var str *types.String
	op1 = opline.Op1()
	op2 = opline.Op2()
	if (op1.IsString()) && (op2.IsString()) {
		var op1_str *types.String = op1.GetStr()
		var op2_str *types.String = op2.GetStr()
		var str *types.String
		if op1_str.GetLen() == 0 {
			{
				opline.Result().SetStringCopy(op2_str)
			}

			{
				// types.ZendStringReleaseEx(op1_str, 0)
			}
		} else if op2_str.GetLen() == 0 {
			{
				opline.Result().SetStringCopy(op1_str)
			}

			{
				// types.ZendStringReleaseEx(op2_str, 0)
			}
		} else {
			str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
			memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
			memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
			opline.Result().SetString(str)
			{
				// types.ZendStringReleaseEx(op1_str, 0)
			}
			{
				// types.ZendStringReleaseEx(op2_str, 0)
			}
		}
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}

	if op1.IsString() {
		op1_str = op1.GetStr().Copy()
	} else {
		if op1.IsUndef() {
			ZVAL_UNDEFINED_OP1(executeData)
		}
		op1_str = ZvalGetStringFunc(op1)
	}

	if op2.IsString() {
		op2_str = op2.GetStr().Copy()
	} else {
		if op2.IsUndef() {
			ZVAL_UNDEFINED_OP2(executeData)
		}
		op2_str = ZvalGetStringFunc(op2)
	}
	for {
		{
			if op1_str.GetLen() == 0 {
				opline.Result().SetString(op2_str)
				// types.ZendStringReleaseEx(op1_str, 0)
				break
			}
		}
		{
			if op2_str.GetLen() == 0 {
				opline.Result().SetString(op1_str)
				// types.ZendStringReleaseEx(op2_str, 0)
				break
			}
		}
		str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
		memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
		memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
		opline.Result().SetString(str)
		{
			// types.ZendStringReleaseEx(op1_str, 0)
		}
		{
			// types.ZendStringReleaseEx(op2_str, 0)
		}
		break
	}
	// ZvalPtrDtorNogc(free_op1)
	// ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FAST_CONCAT_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var op1_str *types.String
	var op2_str *types.String
	var str *types.String
	op1 = opline.Op1()
	op2 = opline.Op2()
	if (op1.IsString()) && (op2.IsString()) {
		var op1_str *types.String = op1.GetStr()
		var op2_str *types.String = op2.GetStr()
		var str *types.String
		if op1_str.GetLen() == 0 {
			{
				opline.Result().SetStringCopy(op2_str)
			}

			{
				// types.ZendStringReleaseEx(op1_str, 0)
			}
		} else if op2_str.GetLen() == 0 {
			{
				opline.Result().SetStringCopy(op1_str)
			}

			{
				// types.ZendStringReleaseEx(op2_str, 0)
			}
		} else {
			str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
			memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
			memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
			opline.Result().SetString(str)
			{
				// types.ZendStringReleaseEx(op1_str, 0)
			}
			{
				// types.ZendStringReleaseEx(op2_str, 0)
			}
		}
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}

	if op1.IsString() {
		op1_str = op1.GetStr().Copy()
	} else {
		if op1.IsUndef() {
			ZVAL_UNDEFINED_OP1(executeData)
		}
		op1_str = ZvalGetStringFunc(op1)
	}

	if op2.IsString() {
		op2_str = op2.GetStr().Copy()
	} else {
		if op2.IsUndef() {
			ZVAL_UNDEFINED_OP2(executeData)
		}
		op2_str = ZvalGetStringFunc(op2)
	}
	for {
		{
			if op1_str.GetLen() == 0 {
				opline.Result().SetString(op2_str)
				// types.ZendStringReleaseEx(op1_str, 0)
				break
			}
		}
		{
			if op2_str.GetLen() == 0 {
				opline.Result().SetString(op1_str)
				// types.ZendStringReleaseEx(op2_str, 0)
				break
			}
		}
		str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
		memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
		memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
		opline.Result().SetString(str)
		{
			// types.ZendStringReleaseEx(op1_str, 0)
		}
		{
			// types.ZendStringReleaseEx(op2_str, 0)
		}
		break
	}
	// ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FAST_CONCAT_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var op1_str *types.String
	var op2_str *types.String
	var str *types.String
	op1 = opline.Op1()
	op2 = executeData.GetOp2(opline)
	if op1.IsString() {
		var op1_str *types.String = op1.GetStr()
		var op2_str *types.String = op2.GetStr()
		var str *types.String
		if op1_str.GetLen() == 0 {
			{
				opline.Result().SetStringCopy(op2_str)
			}

			{
				// types.ZendStringReleaseEx(op1_str, 0)
			}
		} else {
			str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
			memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
			memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
			opline.Result().SetString(str)
			{
				// types.ZendStringReleaseEx(op1_str, 0)
			}
		}
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}

	if op1.IsString() {
		op1_str = op1.GetStr().Copy()
	} else {
		if op1.IsUndef() {
			ZVAL_UNDEFINED_OP1(executeData)
		}
		op1_str = ZvalGetStringFunc(op1)
	}
	{
		op2_str = op2.GetStr()
	}

	for {
		{
			if op1_str.GetLen() == 0 {
				{
					if op2.IsRefcounted() {
						// 						op2_str.AddRefcount()
					}
				}
				opline.Result().SetString(op2_str)
				// types.ZendStringReleaseEx(op1_str, 0)
				break
			}
		}
		str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
		memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
		memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
		opline.Result().SetString(str)
		{
			// types.ZendStringReleaseEx(op1_str, 0)
		}
		break
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FAST_CONCAT_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var op1_str *types.String
	var op2_str *types.String
	var str *types.String
	op1 = opline.Op1()
	op2 = opline.Op2()
	if (op1.IsString()) && (op2.IsString()) {
		var op1_str *types.String = op1.GetStr()
		var op2_str *types.String = op2.GetStr()
		var str *types.String
		if op1_str.GetLen() == 0 {
			{
				opline.Result().SetStringCopy(op2_str)
			}

			{
				// types.ZendStringReleaseEx(op1_str, 0)
			}
		} else if op2_str.GetLen() == 0 {
			{
				opline.Result().SetStringCopy(op1_str)
			}

			{
				// types.ZendStringReleaseEx(op2_str, 0)
			}
		} else {
			str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
			memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
			memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
			opline.Result().SetString(str)
			{
				// types.ZendStringReleaseEx(op1_str, 0)
			}
			{
				// types.ZendStringReleaseEx(op2_str, 0)
			}
		}
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}

	if op1.IsString() {
		op1_str = op1.GetStr().Copy()
	} else {
		if op1.IsUndef() {
			ZVAL_UNDEFINED_OP1(executeData)
		}
		op1_str = ZvalGetStringFunc(op1)
	}

	if op2.IsString() {
		op2_str = op2.GetStr().Copy()
	} else {
		if op2.IsUndef() {
			ZVAL_UNDEFINED_OP2(executeData)
		}
		op2_str = ZvalGetStringFunc(op2)
	}
	for {
		{
			if op1_str.GetLen() == 0 {
				opline.Result().SetString(op2_str)
				// types.ZendStringReleaseEx(op1_str, 0)
				break
			}
		}
		{
			if op2_str.GetLen() == 0 {
				opline.Result().SetString(op1_str)
				// types.ZendStringReleaseEx(op2_str, 0)
				break
			}
		}
		str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
		memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
		memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
		opline.Result().SetString(str)
		{
			// types.ZendStringReleaseEx(op1_str, 0)
		}
		{
			// types.ZendStringReleaseEx(op2_str, 0)
		}
		break
	}
	// ZvalPtrDtorNogc(free_op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FAST_CONCAT_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var op1_str *types.String
	var op2_str *types.String
	var str *types.String
	op1 = opline.Op1()
	op2 = opline.Op2()
	if (op1.IsString()) && (op2.IsString()) {
		var op1_str *types.String = op1.GetStr()
		var op2_str *types.String = op2.GetStr()
		var str *types.String
		if op1_str.GetLen() == 0 {
			{
				opline.Result().SetStringCopy(op2_str)
			}

			{
				// types.ZendStringReleaseEx(op1_str, 0)
			}
		} else if op2_str.GetLen() == 0 {
			{
				opline.Result().SetStringCopy(op1_str)
			}

			{
				// types.ZendStringReleaseEx(op2_str, 0)
			}
		} else {
			str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
			memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
			memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
			opline.Result().SetString(str)
			{
				// types.ZendStringReleaseEx(op1_str, 0)
			}
			{
				// types.ZendStringReleaseEx(op2_str, 0)
			}
		}
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}

	if op1.IsString() {
		op1_str = op1.GetStr().Copy()
	} else {
		if op1.IsUndef() {
			ZVAL_UNDEFINED_OP1(executeData)
		}
		op1_str = ZvalGetStringFunc(op1)
	}

	if op2.IsString() {
		op2_str = op2.GetStr().Copy()
	} else {
		if op2.IsUndef() {
			ZVAL_UNDEFINED_OP2(executeData)
		}
		op2_str = ZvalGetStringFunc(op2)
	}
	for {
		{
			if op1_str.GetLen() == 0 {
				opline.Result().SetString(op2_str)
				// types.ZendStringReleaseEx(op1_str, 0)
				break
			}
		}
		{
			if op2_str.GetLen() == 0 {
				opline.Result().SetString(op1_str)
				// types.ZendStringReleaseEx(op2_str, 0)
				break
			}
		}
		str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
		memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
		memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
		opline.Result().SetString(str)
		{
			// types.ZendStringReleaseEx(op1_str, 0)
		}
		{
			// types.ZendStringReleaseEx(op2_str, 0)
		}
		break
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
