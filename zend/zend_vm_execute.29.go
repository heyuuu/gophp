// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/types"
)

func ZEND_BIND_GLOBAL_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varname *types.String
	var value *types.Zval
	var variable_ptr *types.Zval
	var idx uintPtr
	var ref *types.ZendReference
	for {
		varname = RT_CONSTANT(opline, opline.GetOp2()).GetStr()

		/* We store "hash slot index" + 1 (NULL is a mark of uninitialized cache slot) */

		idx = uintPtr(CACHED_PTR(opline.GetExtendedValue()) - 1)
		if idx < EG__().GetSymbolTable().GetNNumUsed()*b.SizeOf("Bucket") {
			var p *types.Bucket = (*types.Bucket)((*byte)(EG__().GetSymbolTable().GetArData() + idx))
			if p.GetVal().IsNotUndef() && (p.GetKey() == varname || p.GetH() == varname.GetH() && p.GetKey() != nil && types.ZendStringEqualContent(p.GetKey(), varname) != 0) {
				value = (*types.Zval)(p)
				goto check_indirect
			}
		}
		value = EG__().GetSymbolTable().KeyFind(varname.GetStr())
		if value == nil {
			value = EG__().GetSymbolTable().KeyAddNew(varname.GetStr(), EG__().GetUninitializedZval())
			idx = (*byte)(value - (*byte)(EG__().GetSymbolTable().GetArData()))

			/* Store "hash slot index" + 1 (NULL is a mark of uninitialized cache slot) */

			CACHE_PTR(opline.GetExtendedValue(), any(idx+1))

			/* Store "hash slot index" + 1 (NULL is a mark of uninitialized cache slot) */

		} else {
			idx = (*byte)(value - (*byte)(EG__().GetSymbolTable().GetArData()))

			/* Store "hash slot index" + 1 (NULL is a mark of uninitialized cache slot) */

			CACHE_PTR(opline.GetExtendedValue(), any(idx+1))
		check_indirect:

			/* GLOBAL variable may be an INDIRECT pointer to CV */

			if value.IsIndirect() {
				value = value.GetZv()
				if value.IsUndef() {
					value.SetNull()
				}
			}

			/* GLOBAL variable may be an INDIRECT pointer to CV */

		}
		if !(value.IsReference()) {
			types.ZVAL_MAKE_REF_EX(value, 2)
			ref = value.GetRef()
		} else {
			ref = value.GetRef()
			ref.AddRefcount()
		}
		variable_ptr = EX_VAR(opline.GetOp1().GetVar())
		if variable_ptr.IsRefcounted() {
			var ref *types.ZendRefcounted = variable_ptr.GetCounted()
			var refcnt uint32 = ref.DelRefcount()
			if variable_ptr != value {
				if refcnt == 0 {
					RcDtorFunc(ref)
					if EG__().GetException() != nil {
						variable_ptr.SetNull()
						HANDLE_EXCEPTION()
					}
				} else {
					GcCheckPossibleRoot(ref)
				}
			}
		}
		variable_ptr.SetReference(ref)
		if b.PreInc(&opline).opcode != ZEND_BIND_GLOBAL {
			break
		}
	}
	OPLINE = opline
	return 0
	ZEND_VM_NEXT_OPCODE()
}
func ZEND_IN_ARRAY_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var ht *types.Array = RT_CONSTANT(opline, opline.GetOp2()).GetArr()
	var result *types.Zval
	op1 = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	if op1.IsString() {
		result = ht.KeyFind(op1.GetStr().GetStr())
	} else if opline.GetExtendedValue() != 0 {
		if op1.IsLong() {
			result = ht.IndexFind(op1.GetLval())
		} else {
			result = nil
		}
	} else if op1.GetType() <= types.IS_FALSE {
		result = ht.KeyFind(types.ZSTR_EMPTY_ALLOC().GetStr())
	} else {
		var key *types.String
		var key_tmp types.Zval
		var result_tmp types.Zval
		var val *types.Zval
		result = nil
		var __ht *types.Array = ht
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()

			key = _p.GetKey()
			val = _z
			key_tmp.SetString(key)
			CompareFunction(&result_tmp, op1, &key_tmp)
			if result_tmp.GetLval() == 0 {
				result = val
				break
			}
		}
	}
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != nil)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_DIM_R_INDEX_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var dim *types.Zval
	var value *types.Zval
	var offset ZendLong
	var ht *types.Array
	container = EX_VAR(opline.GetOp1().GetVar())
	dim = RT_CONSTANT(opline, opline.GetOp2())
	if container.IsArray() {
	fetch_dim_r_index_array:
		if dim.IsLong() {
			offset = dim.GetLval()
		} else {
			offset = ZvalGetLong(dim)
		}
		ht = container.GetArr()
		value = ht.IndexFind(offset)
		if value == nil {
			goto fetch_dim_r_index_undef
		}
		types.ZVAL_COPY_DEREF(EX_VAR(opline.GetResult().GetVar()), value)
		{
			ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
		}

	} else if container.IsReference() {
		container = types.Z_REFVAL_P(container)
		if container.IsArray() {
			goto fetch_dim_r_index_array
		} else {
			goto fetch_dim_r_index_slow
		}
	} else {
	fetch_dim_r_index_slow:
		if dim.GetU2Extra() == ZEND_EXTRA_VALUE {
			dim++
		}
		zend_fetch_dimension_address_read_R_slow(container, dim, opline, executeData)
		ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
	}
fetch_dim_r_index_undef:
	EX_VAR(opline.GetResult().GetVar()).SetNull()
	ZendUndefinedOffset(offset)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_DIM_R_INDEX_SPEC_CV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var dim *types.Zval
	var value *types.Zval
	var offset ZendLong
	var ht *types.Array
	container = EX_VAR(opline.GetOp1().GetVar())
	dim = EX_VAR(opline.GetOp2().GetVar())
	if container.IsArray() {
	fetch_dim_r_index_array:
		if dim.IsLong() {
			offset = dim.GetLval()
		} else {
			offset = ZvalGetLong(dim)
		}
		ht = container.GetArr()
		value = ht.IndexFind(offset)
		if value == nil {
			goto fetch_dim_r_index_undef
		}
		types.ZVAL_COPY_DEREF(EX_VAR(opline.GetResult().GetVar()), value)
		{
			ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
		}

	} else if container.IsReference() {
		container = types.Z_REFVAL_P(container)
		if container.IsArray() {
			goto fetch_dim_r_index_array
		} else {
			goto fetch_dim_r_index_slow
		}
	} else {
	fetch_dim_r_index_slow:
		zend_fetch_dimension_address_read_R_slow(container, dim, opline, executeData)
		ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
	}
fetch_dim_r_index_undef:
	EX_VAR(opline.GetResult().GetVar()).SetNull()
	ZendUndefinedOffset(offset)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_DIV_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	FastDivFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_POW_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	PowFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_CONCAT_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	if (op1.IsString()) && (op2.IsString()) {
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
		} else if op2_str.GetLen() == 0 {
			{
				EX_VAR(opline.GetResult().GetVar()).SetStringCopy(op1_str)
			}

			{
				types.ZendStringReleaseEx(op2_str, 0)
			}
		} else {
			str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
			memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
			memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
			EX_VAR(opline.GetResult().GetVar()).SetString(str)
			{
				types.ZendStringReleaseEx(op1_str, 0)
			}
			{
				types.ZendStringReleaseEx(op2_str, 0)
			}
		}
		ZEND_VM_NEXT_OPCODE()
	} else {
		if op1.IsUndef() {
			op1 = ZVAL_UNDEFINED_OP1()
		}
		if op2.IsUndef() {
			op2 = ZVAL_UNDEFINED_OP2()
		}
		ConcatFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
		ZvalPtrDtorNogc(free_op2)
		ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
	}
}
func ZEND_IS_EQUAL_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)

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
func ZEND_IS_EQUAL_SPEC_CV_TMPVAR_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)

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
func ZEND_IS_EQUAL_SPEC_CV_TMPVAR_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)

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
func ZEND_IS_NOT_EQUAL_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)

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
func ZEND_IS_NOT_EQUAL_SPEC_CV_TMPVAR_JMPZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)

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
func ZEND_IS_NOT_EQUAL_SPEC_CV_TMPVAR_JMPNZ_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = EX_VAR(opline.GetOp1().GetVar())
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)

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
func ZEND_SPACESHIP_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	CompareFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_BOOL_XOR_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	BooleanXorFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_ASSIGN_OBJ_OP_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var zptr *types.Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = EX_VAR(opline.GetOp1().GetVar())
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	for {
		value = GetOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data)
		if object.GetType() != types.IS_OBJECT {
			if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
				object = types.Z_REFVAL_P(object)
				goto assign_op_object
			}
			if object.IsUndef() {
				ZVAL_UNDEFINED_OP1()
			}
			object = MakeRealObject(object, property, opline, executeData)
			if object == nil {
				break
			}
		}
	assign_op_object:

		/* here we are sure we are dealing with an object */

		{
			cache_slot = nil
		}
		if b.Assign(&zptr, types.Z_OBJ_HT_P(object).GetGetPropertyPtrPtr()(object, property, BP_VAR_RW, cache_slot)) != nil {
			if zptr.IsError() {
				if RETURN_VALUE_USED(opline) {
					EX_VAR(opline.GetResult().GetVar()).SetNull()
				}
			} else {
				var orig_zptr *types.Zval = zptr
				var ref *types.ZendReference
				for {
					if zptr.IsReference() {
						ref = zptr.GetRef()
						zptr = types.Z_REFVAL_P(zptr)
						if ZEND_REF_HAS_TYPE_SOURCES(ref) {
							ZendBinaryAssignOpTypedRef(ref, value, opline, executeData)
							break
						}
					}

					{
						prop_info = ZendObjectFetchPropertyTypeInfo(object.GetObj(), orig_zptr)
					}
					if prop_info != nil {

						/* special case for typed properties */

						ZendBinaryAssignOpTypedProp(prop_info, zptr, value, opline, executeData)

						/* special case for typed properties */

					} else {
						ZendBinaryOp(zptr, zptr, value, opline)
					}
					break
				}
				if RETURN_VALUE_USED(opline) {
					types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), zptr)
				}
			}
		} else {
			ZendAssignOpOverloadedProperty(object, property, cache_slot, value, opline, executeData)
		}
		break
	}
	FREE_OP(free_op_data)
	ZvalPtrDtorNogc(free_op2)

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_DIM_OP_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var free_op_data1 ZendFreeOp
	var var_ptr *types.Zval
	var value *types.Zval
	var container *types.Zval
	var dim *types.Zval
	container = EX_VAR(opline.GetOp1().GetVar())
	if container.IsArray() {
	assign_dim_op_array:
		types.SEPARATE_ARRAY(container)
	assign_dim_op_new_array:
		dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)

		{

			{
				var_ptr = zend_fetch_dimension_address_inner_RW(container.GetArr(), dim, executeData)
			}
			if var_ptr == nil {
				goto assign_dim_op_ret_null
			}
		}
		value = GetOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data1)
		for {
			if var_ptr.IsReference() {
				var ref *types.ZendReference = var_ptr.GetRef()
				var_ptr = types.Z_REFVAL_P(var_ptr)
				if ZEND_REF_HAS_TYPE_SOURCES(ref) {
					ZendBinaryAssignOpTypedRef(ref, value, opline, executeData)
					break
				}
			}
			ZendBinaryOp(var_ptr, var_ptr, value, opline)
			break
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), var_ptr)
		}
		FREE_OP(free_op_data1)
	} else {
		if container.IsReference() {
			container = types.Z_REFVAL_P(container)
			if container.IsArray() {
				goto assign_dim_op_array
			}
		}
		dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
		if container.IsObject() {
			ZendBinaryAssignOpObjDim(container, dim, opline, executeData)
		} else if container.GetType() <= types.IS_FALSE {
			if container.IsUndef() {
				ZVAL_UNDEFINED_OP1()
			}
			container.SetArray(types.NewArray(8))
			goto assign_dim_op_new_array
		} else {
			ZendBinaryAssignOpDimSlow(container, dim, opline, executeData)
		assign_dim_op_ret_null:
			FREE_UNFETCHED_OP((opline + 1).GetOp1Type(), (opline + 1).GetOp1().GetVar())
			if RETURN_VALUE_USED(opline) {
				EX_VAR(opline.GetResult().GetVar()).SetNull()
			}
		}
	}
	ZvalPtrDtorNogc(free_op2)
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OP_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var var_ptr *types.Zval
	var value *types.Zval
	value = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	var_ptr = _get_zval_ptr_cv_BP_VAR_RW(opline.GetOp1().GetVar(), executeData)
	if var_ptr.IsError() {
		if RETURN_VALUE_USED(opline) {
			EX_VAR(opline.GetResult().GetVar()).SetNull()
		}
	} else {
		for {
			if var_ptr.IsReference() {
				var ref *types.ZendReference = var_ptr.GetRef()
				var_ptr = types.Z_REFVAL_P(var_ptr)
				if ZEND_REF_HAS_TYPE_SOURCES(ref) {
					ZendBinaryAssignOpTypedRef(ref, value, opline, executeData)
					break
				}
			}
			ZendBinaryOp(var_ptr, var_ptr, value, opline)
			break
		}
		if RETURN_VALUE_USED(opline) {
			types.ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), var_ptr)
		}
	}
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_PRE_INC_OBJ_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var zptr *types.Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = EX_VAR(opline.GetOp1().GetVar())
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	for {
		if object.GetType() != types.IS_OBJECT {
			if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
				object = types.Z_REFVAL_P(object)
				goto pre_incdec_object
			}
			if object.IsUndef() {
				ZVAL_UNDEFINED_OP1()
			}
			object = MakeRealObject(object, property, opline, executeData)
			if object == nil {
				break
			}
		}
	pre_incdec_object:

		/* here we are sure we are dealing with an object */

		{
			cache_slot = nil
		}
		if b.Assign(&zptr, types.Z_OBJ_HT_P(object).GetGetPropertyPtrPtr()(object, property, BP_VAR_RW, cache_slot)) != nil {
			if zptr.IsError() {
				if RETURN_VALUE_USED(opline) {
					EX_VAR(opline.GetResult().GetVar()).SetNull()
				}
			} else {

				{
					prop_info = ZendObjectFetchPropertyTypeInfo(object.GetObj(), zptr)
				}
				ZendPreIncdecPropertyZval(zptr, prop_info, opline, executeData)
			}
		} else {
			ZendPreIncdecOverloadedProperty(object, property, cache_slot, opline, executeData)
		}
		break
	}
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_POST_INC_OBJ_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var zptr *types.Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = EX_VAR(opline.GetOp1().GetVar())
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	for {
		if object.GetType() != types.IS_OBJECT {
			if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
				object = types.Z_REFVAL_P(object)
				goto post_incdec_object
			}
			if object.IsUndef() {
				ZVAL_UNDEFINED_OP1()
			}
			object = MakeRealObject(object, property, opline, executeData)
			if object == nil {
				break
			}
		}
	post_incdec_object:

		/* here we are sure we are dealing with an object */

		{
			cache_slot = nil
		}
		if b.Assign(&zptr, types.Z_OBJ_HT_P(object).GetGetPropertyPtrPtr()(object, property, BP_VAR_RW, cache_slot)) != nil {
			if zptr.IsError() {
				EX_VAR(opline.GetResult().GetVar()).SetNull()
			} else {

				{
					prop_info = ZendObjectFetchPropertyTypeInfo(object.GetObj(), zptr)
				}
				ZendPostIncdecPropertyZval(zptr, prop_info, opline, executeData)
			}
		} else {
			ZendPostIncdecOverloadedProperty(object, property, cache_slot, opline, executeData)
		}
		break
	}
	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
func ZEND_FETCH_DIM_R_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op2 ZendFreeOp
	var container *types.Zval
	var dim *types.Zval
	var value *types.Zval
	container = EX_VAR(opline.GetOp1().GetVar())
	dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, executeData)
	{
		if container.IsArray() {
		fetch_dim_r_array:
			value = ZendFetchDimensionAddressInner(container.GetArr(), dim, IS_TMP_VAR|IS_VAR, BP_VAR_R, executeData)
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
			zend_fetch_dimension_address_read_R_slow(container, dim, opline, executeData)
		}
	}

	ZvalPtrDtorNogc(free_op2)
	ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION()
}
