// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/faults"
	"sik/zend/types"
	"sik/zend/zpp"
)

func ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = op1.GetLval() <= op2.GetLval()
	ZEND_VM_SMART_BRANCH_JMPNZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = op1.GetDval() <= op2.GetDval()
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = op1.GetDval() <= op2.GetDval()
	ZEND_VM_SMART_BRANCH_JMPZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result int
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = EX_VAR(opline.GetOp2().GetVar())
	result = op1.GetDval() <= op2.GetDval()
	ZEND_VM_SMART_BRANCH_JMPNZ(result, 0)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_FETCH_LIST_R_SPEC_TMPVARCV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var container *types.Zval
	container = EX_VAR(opline.GetOp1().GetVar())
	zend_fetch_dimension_address_LIST_r(container, _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData), IS_TMP_VAR|IS_VAR, opline, executeData)
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_LIST_R_SPEC_TMPVARCV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	container = EX_VAR(opline.GetOp1().GetVar())
	zend_fetch_dimension_address_LIST_r(container, EX_VAR(opline.GetOp2().GetVar()), IS_CV, opline, executeData)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_BW_NOT_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if op1.IsLong() {
		EX_VAR(opline.GetResult().GetVar()).SetLong(^(op1.GetLval()))
		ZEND_VM_NEXT_OPCODE()
	}
	if op1.IsUndef() {
		op1 = ZVAL_UNDEFINED_OP1()
	}
	BitwiseNotFunction(EX_VAR(opline.GetResult().GetVar()), op1)
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_BOOL_NOT_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	var free_op1 ZendFreeOp
	val = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if val.IsTrue() {
		EX_VAR(opline.GetResult().GetVar()).SetFalse()
	} else if val.GetTypeInfo() <= types.IS_TRUE {

		/* The result and op1 can be the same cv zval */

		var orig_val_type uint32 = val.GetTypeInfo()
		EX_VAR(opline.GetResult().GetVar()).SetTrue()
		if orig_val_type == types.IS_UNDEF {
			ZVAL_UNDEFINED_OP1()
			ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
		}
	} else {
		types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), IZendIsTrue(val) == 0)
		ZvalPtrDtorNogc(free_op1)
		ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_ECHO_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var z *types.Zval
	z = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if z.IsString() {
		var str *types.String = z.GetStr()
		if str.GetLen() != 0 {
			ZendWrite(str.GetStr())
		}
	} else {
		var str *types.String = ZvalGetStringFunc(z)
		if str.GetLen() != 0 {
			ZendWrite(str.GetStr())
		} else if z.IsUndef() {
			ZVAL_UNDEFINED_OP1()
		}
		types.ZendStringReleaseEx(str, 0)
	}
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_JMPZ_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var val *types.Zval
	val = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if val.IsTrue() {
		ZEND_VM_NEXT_OPCODE()
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		if val.IsUndef() {
			ZVAL_UNDEFINED_OP1()
			if EG__().GetException() != nil {
				HANDLE_EXCEPTION()
			}
		}
		ZEND_VM_JMP_EX(OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	if IZendIsTrue(val) != 0 {
		opline++
	} else {
		opline = OP_JMP_ADDR(opline, opline.GetOp2())
	}
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_JMP(opline)
}
func ZEND_JMPNZ_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var val *types.Zval
	val = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if val.IsTrue() {
		ZEND_VM_JMP_EX(OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		if val.IsUndef() {
			ZVAL_UNDEFINED_OP1()
			if EG__().GetException() != nil {
				HANDLE_EXCEPTION()
			}
		}
		ZEND_VM_NEXT_OPCODE()
	}
	if IZendIsTrue(val) != 0 {
		opline = OP_JMP_ADDR(opline, opline.GetOp2())
	} else {
		opline++
	}
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_JMP(opline)
}
func ZEND_JMPZNZ_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var val *types.Zval
	val = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if val.IsTrue() {
		ZEND_VM_SET_RELATIVE_OPCODE(opline, opline.GetExtendedValue())
		return 0
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		if val.IsUndef() {
			ZVAL_UNDEFINED_OP1()
			if EG__().GetException() != nil {
				HANDLE_EXCEPTION()
			}
		}
		ZEND_VM_JMP_EX(OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	if IZendIsTrue(val) != 0 {
		opline = ZEND_OFFSET_TO_OPLINE(opline, opline.GetExtendedValue())
	} else {
		opline = OP_JMP_ADDR(opline, opline.GetOp2())
	}
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_JMP(opline)
}
func ZEND_JMPZ_EX_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var val *types.Zval
	var ret int
	val = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if val.IsTrue() {
		EX_VAR(opline.GetResult().GetVar()).SetTrue()
		ZEND_VM_NEXT_OPCODE()
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		EX_VAR(opline.GetResult().GetVar()).SetFalse()
		if val.IsUndef() {
			ZVAL_UNDEFINED_OP1()
			if EG__().GetException() != nil {
				HANDLE_EXCEPTION()
			}
		}
		ZEND_VM_JMP_EX(OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	ret = IZendIsTrue(val)
	ZvalPtrDtorNogc(free_op1)
	if ret != 0 {
		EX_VAR(opline.GetResult().GetVar()).SetTrue()
		opline++
	} else {
		EX_VAR(opline.GetResult().GetVar()).SetFalse()
		opline = OP_JMP_ADDR(opline, opline.GetOp2())
	}
	ZEND_VM_JMP(opline)
}
func ZEND_JMPNZ_EX_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var val *types.Zval
	var ret int
	val = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if val.IsTrue() {
		EX_VAR(opline.GetResult().GetVar()).SetTrue()
		ZEND_VM_JMP_EX(OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	} else if val.GetTypeInfo() <= types.IS_TRUE {
		EX_VAR(opline.GetResult().GetVar()).SetFalse()
		if val.IsUndef() {
			ZVAL_UNDEFINED_OP1()
			ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
		} else {
			ZEND_VM_NEXT_OPCODE()
		}
	}
	ret = IZendIsTrue(val)
	ZvalPtrDtorNogc(free_op1)
	if ret != 0 {
		EX_VAR(opline.GetResult().GetVar()).SetTrue()
		opline = OP_JMP_ADDR(opline, opline.GetOp2())
	} else {
		EX_VAR(opline.GetResult().GetVar()).SetFalse()
		opline++
	}
	ZEND_VM_JMP(opline)
}
func ZEND_FREE_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	ZvalPtrDtorNogc(EX_VAR(opline.GetOp1().GetVar()))
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FE_FREE_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var var_ *types.Zval
	var opline *ZendOp = executeData.GetOpline()
	var_ = EX_VAR(opline.GetOp1().GetVar())
	if var_.GetType() != types.IS_ARRAY && var_.GetFeIterIdx() != uint32-1 {
		types.ZendHashIteratorDel(var_.GetFeIterIdx())
	}
	ZvalPtrDtorNogc(var_)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_SEND_VAL_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var arg *types.Zval
	var free_op1 ZendFreeOp
	value = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	types.ZVAL_COPY_VALUE(arg, value)
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_BOOL_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var val *types.Zval
	var free_op1 ZendFreeOp
	val = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if val.IsTrue() {
		EX_VAR(opline.GetResult().GetVar()).SetTrue()
	} else if val.GetTypeInfo() <= types.IS_TRUE {

		/* The result and op1 can be the same cv zval */

		var orig_val_type uint32 = val.GetTypeInfo()
		EX_VAR(opline.GetResult().GetVar()).SetFalse()
		if orig_val_type == types.IS_UNDEF {
			ZVAL_UNDEFINED_OP1()
			ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
		}
	} else {
		types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), IZendIsTrue(val) != 0)
		ZvalPtrDtorNogc(free_op1)
		ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_CLONE_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var obj *types.Zval
	var ce *types.ClassEntry
	var scope *types.ClassEntry
	var clone *ZendFunction
	var clone_call ZendObjectCloneObjT
	obj = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	for {
		if obj.GetType() != types.IS_OBJECT {
			if obj.IsReference() {
				obj = types.Z_REFVAL_P(obj)
				if obj.IsObject() {
					break
				}
			}
			EX_VAR(opline.GetResult().GetVar()).SetUndef()
			if obj.IsUndef() {
				ZVAL_UNDEFINED_OP1()
				if EG__().GetException() != nil {
					HANDLE_EXCEPTION()
				}
			}
			faults.ThrowError(nil, "__clone method called on non-object")
			ZvalPtrDtorNogc(free_op1)
			HANDLE_EXCEPTION()
		}
		break
	}
	ce = types.Z_OBJCE_P(obj)
	clone = ce.GetClone()
	clone_call = types.Z_OBJ_HT_P(obj).GetCloneObj()
	if clone_call == nil {
		faults.ThrowError(nil, "Trying to clone an uncloneable object of class %s", ce.GetName().GetVal())
		ZvalPtrDtorNogc(free_op1)
		EX_VAR(opline.GetResult().GetVar()).SetUndef()
		HANDLE_EXCEPTION()
	}
	if clone != nil && !clone.IsPublic() {
		scope = executeData.GetFunc().op_array.scope
		if clone.GetScope() != scope {
			if clone.IsPrivate() || ZendCheckProtected(ZendGetFunctionRootClass(clone), scope) == 0 {
				ZendWrongCloneCall(clone, scope)
				ZvalPtrDtorNogc(free_op1)
				EX_VAR(opline.GetResult().GetVar()).SetUndef()
				HANDLE_EXCEPTION()
			}
		}
	}
	EX_VAR(opline.GetResult().GetVar()).SetObject(clone_call(obj))
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_INCLUDE_OR_EVAL_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var new_op_array *ZendOpArray
	var free_op1 ZendFreeOp
	var inc_filename *types.Zval
	inc_filename = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	new_op_array = ZendIncludeOrEval(inc_filename, opline.GetExtendedValue())
	ZvalPtrDtorNogc(free_op1)
	if EG__().GetException() != nil {
		if new_op_array != ZEND_FAKE_OP_ARRAY && new_op_array != nil {
			DestroyOpArray(new_op_array)
			EfreeSize(new_op_array, b.SizeOf("zend_op_array"))
		}
		UNDEF_RESULT()
		HANDLE_EXCEPTION()
	} else if new_op_array == ZEND_FAKE_OP_ARRAY {
		if RETURN_VALUE_USED(opline) {
			EX_VAR(opline.GetResult().GetVar()).SetTrue()
		}
	} else if new_op_array != nil {
		var return_value *types.Zval = nil
		var call *ZendExecuteData
		if RETURN_VALUE_USED(opline) {
			return_value = EX_VAR(opline.GetResult().GetVar())
		}
		new_op_array.SetScope(executeData.GetFunc().op_array.scope)
		call = ZendVmStackPushCallFrame(executeData.GetThis().GetTypeInfo()&ZEND_CALL_HAS_THIS|ZEND_CALL_NESTED_CODE|ZEND_CALL_HAS_SYMBOL_TABLE, (*ZendFunction)(new_op_array), 0, executeData.GetThis().GetPtr())
		if (EX_CALL_INFO() & ZEND_CALL_HAS_SYMBOL_TABLE) != 0 {
			call.SetSymbolTable(executeData.GetSymbolTable())
		} else {
			call.SetSymbolTable(ZendRebuildSymbolTable())
		}
		call.SetPrevExecuteData(executeData)
		IInitCodeExecuteData(call, new_op_array, return_value)
		if ZendExecuteEx == ExecuteEx {
			return 1
		} else {
			ZEND_ADD_CALL_FLAG(call, ZEND_CALL_TOP)
			ZendExecuteEx(call)
			ZendVmStackFreeCallFrame(call)
		}
		DestroyOpArray(new_op_array)
		EfreeSize(new_op_array, b.SizeOf("zend_op_array"))
		if EG__().GetException() != nil {
			faults.RethrowException(executeData)
			UNDEF_RESULT()
			HANDLE_EXCEPTION()
		}
	} else if RETURN_VALUE_USED(opline) {
		EX_VAR(opline.GetResult().GetVar()).SetFalse()
	}
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_STRLEN_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var free_op1 ZendFreeOp
	value = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if value.IsString() {
		EX_VAR(opline.GetResult().GetVar()).SetLong(value.GetStr().GetLen())
		ZvalPtrDtorNogc(free_op1)
		ZEND_VM_NEXT_OPCODE()
	} else {
		var strict types.ZendBool
		if value.IsReference() {
			value = types.Z_REFVAL_P(value)
			if value.IsString() {
				EX_VAR(opline.GetResult().GetVar()).SetLong(value.GetStr().GetLen())
				ZvalPtrDtorNogc(free_op1)
				ZEND_VM_NEXT_OPCODE()
			}
		}
		if value.IsUndef() {
			value = ZVAL_UNDEFINED_OP1()
		}
		strict = executeData.IsCallUseStrictTypes()
		for {
			if strict == 0 {
				var str *types.String
				var tmp types.Zval
				types.ZVAL_COPY(&tmp, value)
				if zpp.ZendParseArgStrWeak(&tmp, &str) != 0 {
					EX_VAR(opline.GetResult().GetVar()).SetLong(str.GetLen())
					ZvalPtrDtor(&tmp)
					break
				}
				ZvalPtrDtor(&tmp)
			}
			if EG__().GetException() == nil {
				faults.InternalTypeError(strict, "strlen() expects parameter 1 to be string, %s given", types.ZendGetTypeByConst(value.GetType()))
			}
			EX_VAR(opline.GetResult().GetVar()).SetNull()
			break
		}
	}
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_TYPE_CHECK_SPEC_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var result int = 0
	var free_op1 ZendFreeOp
	value = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	if (opline.GetExtendedValue() >> uint32(value.GetType()) & 1) != 0 {
	type_check_resource:
		if value.GetType() != types.IS_RESOURCE || nil != ZendRsrcListGetRsrcType(value.GetRes()) {
			result = 1
		}
	} else if value.IsReference() {
		value = types.Z_REFVAL_P(value)
		if (opline.GetExtendedValue() >> uint32(value.GetType()) & 1) != 0 {
			goto type_check_resource
		}
	} else if value.IsUndef() {
		result = (1 << types.IS_NULL & opline.GetExtendedValue()) != 0
		ZVAL_UNDEFINED_OP1()
		if EG__().GetException() != nil {
			EX_VAR(opline.GetResult().GetVar()).SetUndef()
			HANDLE_EXCEPTION()
		}
	}
	{
		ZvalPtrDtorNogc(free_op1)
		ZEND_VM_SMART_BRANCH(result, 1)
		types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
		ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
	}

}
func ZEND_DIV_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	FastDivFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_POW_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	PowFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_CONCAT_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	if op1.IsString() {
		var op1_str *types.String = op1.GetStr()
		var op2_str *types.String = op2.GetStr()
		var str *types.String
		if op1_str.GetLen() == 0 {
			{
				EX_VAR(opline.GetResult().GetVar()).SetStringCopy(op2_str)
			}

			{
				types.ZendStringReleaseEx(op1_str, 0)
			}
		} else {
			str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
			memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
			memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
			EX_VAR(opline.GetResult().GetVar()).SetString(str)
			{
				types.ZendStringReleaseEx(op1_str, 0)
			}
		}
		ZEND_VM_NEXT_OPCODE()
	} else {
		if op1.IsUndef() {
			op1 = ZVAL_UNDEFINED_OP1()
		}
		ConcatFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
		ZvalPtrDtorNogc(free_op1)
		ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
	}
}
func ZEND_IS_EQUAL_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() == op2.GetLval() {
			is_equal_true:
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				ZEND_VM_NEXT_OPCODE()
			} else {
			is_equal_false:
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				ZEND_VM_NEXT_OPCODE()
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto is_equal_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		is_equal_double:
			if d1 == d2 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto is_equal_double
		}
	} else if op1.IsString() {
		if op2.IsString() {
			var result int = ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
			{

			}
			if result != 0 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		}
	}
	return zend_is_equal_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_EQUAL_SPEC_TMPVAR_CONST_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() == op2.GetLval() {
			is_equal_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPZ()
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				ZEND_VM_NEXT_OPCODE()
			} else {
			is_equal_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPZ()
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				ZEND_VM_NEXT_OPCODE()
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto is_equal_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		is_equal_double:
			if d1 == d2 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto is_equal_double
		}
	} else if op1.IsString() {
		if op2.IsString() {
			var result int = ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
			{

			}
			if result != 0 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		}
	}
	return zend_is_equal_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_EQUAL_SPEC_TMPVAR_CONST_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() == op2.GetLval() {
			is_equal_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPNZ()
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				ZEND_VM_NEXT_OPCODE()
			} else {
			is_equal_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPNZ()
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				ZEND_VM_NEXT_OPCODE()
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto is_equal_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		is_equal_double:
			if d1 == d2 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto is_equal_double
		}
	} else if op1.IsString() {
		if op2.IsString() {
			var result int = ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
			{

			}
			if result != 0 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		}
	}
	return zend_is_equal_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() != op2.GetLval() {
			is_not_equal_true:
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				ZEND_VM_NEXT_OPCODE()
			} else {
			is_not_equal_false:
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				ZEND_VM_NEXT_OPCODE()
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto is_not_equal_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		is_not_equal_double:
			if d1 != d2 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto is_not_equal_double
		}
	} else if op1.IsString() {
		if op2.IsString() {
			var result int = ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
			{

			}
			if result == 0 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		}
	}
	return zend_is_not_equal_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_CONST_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() != op2.GetLval() {
			is_not_equal_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPZ()
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				ZEND_VM_NEXT_OPCODE()
			} else {
			is_not_equal_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPZ()
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				ZEND_VM_NEXT_OPCODE()
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto is_not_equal_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		is_not_equal_double:
			if d1 != d2 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto is_not_equal_double
		}
	} else if op1.IsString() {
		if op2.IsString() {
			var result int = ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
			{

			}
			if result == 0 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		}
	}
	return zend_is_not_equal_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_CONST_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	if op1.IsLong() {
		if op2.IsLong() {
			if op1.GetLval() != op2.GetLval() {
			is_not_equal_true:
				ZEND_VM_SMART_BRANCH_TRUE_JMPNZ()
				EX_VAR(opline.GetResult().GetVar()).SetTrue()
				ZEND_VM_NEXT_OPCODE()
			} else {
			is_not_equal_false:
				ZEND_VM_SMART_BRANCH_FALSE_JMPNZ()
				EX_VAR(opline.GetResult().GetVar()).SetFalse()
				ZEND_VM_NEXT_OPCODE()
			}
		} else if op2.IsDouble() {
			d1 = float64(op1.GetLval())
			d2 = op2.GetDval()
			goto is_not_equal_double
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			d1 = op1.GetDval()
			d2 = op2.GetDval()
		is_not_equal_double:
			if d1 != d2 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		} else if op2.IsLong() {
			d1 = op1.GetDval()
			d2 = float64(op2.GetLval())
			goto is_not_equal_double
		}
	} else if op1.IsString() {
		if op2.IsString() {
			var result int = ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
			{

			}
			if result == 0 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		}
	}
	return zend_is_not_equal_helper_SPEC(op1, op2, executeData)
}
func ZEND_SPACESHIP_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	CompareFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_BOOL_XOR_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	BooleanXorFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_DIM_R_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var dim *types.Zval
	var value *types.Zval
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	dim = RT_CONSTANT(opline, opline.GetOp2())
	{
		if container.IsArray() {
		fetch_dim_r_array:
			value = ZendFetchDimensionAddressInner(container.GetArr(), dim, IS_CONST, BP_VAR_R, executeData)
			types.ZVAL_COPY_DEREF(EX_VAR(opline.GetResult().GetVar()), value)
		} else if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if container.IsArray() {
				goto fetch_dim_r_array
			} else {
				goto fetch_dim_r_slow
			}
		} else {
		fetch_dim_r_slow:
			if dim.GetU2Extra() == ZEND_EXTRA_VALUE {
				dim++
			}
			zend_fetch_dimension_address_read_R_slow(container, dim, opline, executeData)
		}
	}

	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_DIM_IS_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	zend_fetch_dimension_address_read_IS(container, RT_CONSTANT(opline, opline.GetOp2()), IS_CONST, opline, executeData)
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_OBJ_R_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var container *types.Zval
	var offset *types.Zval
	var cache_slot *any = nil
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	offset = RT_CONSTANT(opline, opline.GetOp2())
	if container.GetType() != types.IS_OBJECT {
		for {
			if container.IsReference() {
				container = types.Z_REFVAL_P(container)
				if container.IsObject() {
					break
				}
			}
			if container.IsUndef() {
				ZVAL_UNDEFINED_OP1()
			}
			ZendWrongPropertyRead(offset)
			EX_VAR(opline.GetResult().GetVar()).SetNull()
			goto fetch_obj_r_finish
			break
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj *types.ZendObject = container.GetObj()
	var retval *types.Zval
	{
		cache_slot = CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_FETCH_REF)
		if zobj.GetCe() == CACHED_PTR_EX(cache_slot) {
			var prop_offset uintPtr = uintPtr(CACHED_PTR_EX(cache_slot + 1))
			if IS_VALID_PROPERTY_OFFSET(prop_offset) {
				retval = OBJ_PROP(zobj, prop_offset)
				if !retval.IsUndef() {
					{
						goto fetch_obj_r_copy
					}

				}
			} else if zobj.GetProperties() != nil {
				if !(IS_UNKNOWN_DYNAMIC_PROPERTY_OFFSET(prop_offset)) {
					var idx uintPtr = ZEND_DECODE_DYN_PROP_OFFSET(prop_offset)
					if idx < zobj.GetProperties().GetNNumUsed()*b.SizeOf("Bucket") {
						var p *types.Bucket = (*types.Bucket)((*byte)(zobj.GetProperties().GetArData() + idx))
						if p.GetVal().IsNotUndef() && (p.GetKey() == offset.GetStr() || (p.IsStrKey() && p.StrKey() == offset.GetStrVal())) {
							retval = p.GetVal()
							goto fetch_obj_r_copy
						}
					}
					CACHE_PTR_EX(cache_slot+1, any(ZEND_DYNAMIC_PROPERTY_OFFSET))
				}
				retval = zobj.GetProperties().KeyFind(offset.GetStr().GetStr())
				if retval != nil {
					var idx uintPtr = (*byte)(retval - (*byte)(zobj.GetProperties().GetArData()))
					CACHE_PTR_EX(cache_slot+1, any(ZEND_ENCODE_DYN_PROP_OFFSET(idx)))
					{
						goto fetch_obj_r_copy
					}

				}
			}
		}
	}

	retval = zobj.GetHandlers().GetReadProperty()(container, offset, BP_VAR_R, cache_slot, EX_VAR(opline.GetResult().GetVar()))
	if retval != EX_VAR(opline.GetResult().GetVar()) {
	fetch_obj_r_copy:
		types.ZVAL_COPY_DEREF(EX_VAR(opline.GetResult().GetVar()), retval)
	} else if retval.IsReference() {
		ZendUnwrapReference(retval)
	}
fetch_obj_r_finish:
	ZvalPtrDtorNogc(free_op1)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
