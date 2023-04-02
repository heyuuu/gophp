package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_TYPE_CHECK_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var result int = 0
	value = RT_CONSTANT(opline, opline.GetOp1())
	if (opline.GetExtendedValue() >> uint32(value.GetType()) & 1) != 0 {
	type_check_resource:
		if value.GetType() != types.IS_RESOURCE || nil != ZendRsrcListGetRsrcType(value.GetRes()) {
			result = 1
		}
	} else {
	}

	{
		ZEND_VM_SMART_BRANCH(result, 0)
		types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
}
func ZEND_DEFINED_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var c *ZendConstant
	c = CACHED_PTR(opline.GetExtendedValue())
	if c != nil {
		if IS_SPECIAL_CACHE_VAL(c) == 0 {
		defined_true:
			ZEND_VM_SMART_BRANCH_TRUE()
			EX_VAR(opline.GetResult().GetVar()).SetTrue()
			return ZEND_VM_NEXT_OPCODE(executeData, opline)
		} else if EG__().GetZendConstants().Len() == DECODE_SPECIAL_CACHE_NUM(c) {
		defined_false:
			ZEND_VM_SMART_BRANCH_FALSE()
			EX_VAR(opline.GetResult().GetVar()).SetFalse()
			return ZEND_VM_NEXT_OPCODE(executeData, opline)
		}
	}
	if ZendQuickCheckConstant(RT_CONSTANT(opline, opline.GetOp1()), opline, executeData) != types.SUCCESS {
		CACHE_PTR(opline.GetExtendedValue(), ENCODE_SPECIAL_CACHE_NUM(EG__().GetZendConstants().Len()))
		goto defined_false
	} else {
		goto defined_true
	}
}
func ZEND_ADD_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	var d1 float64
	var d2 float64
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	return zend_add_helper_SPEC(op1, op2, executeData)
}
func ZEND_SUB_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	var d1 float64
	var d2 float64
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	return zend_sub_helper_SPEC(op1, op2, executeData)
}
func ZEND_MUL_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	var d1 float64
	var d2 float64
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	return zend_mul_helper_SPEC(op1, op2, executeData)
}
func ZEND_DIV_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	FastDivFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_MOD_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	/* Prevent overflow error/crash if op1==ZEND_LONG_MIN */

	/* Prevent overflow error/crash if op1==ZEND_LONG_MIN */

	return zend_mod_helper_SPEC(op1, op2, executeData)
}
func ZEND_SL_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	/* Perform shift on unsigned numbers to get well-defined wrap behavior. */

	return zend_shift_left_helper_SPEC(op1, op2, executeData)
}
func ZEND_SR_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	return zend_shift_right_helper_SPEC(op1, op2, executeData)
}
func ZEND_POW_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	PowFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_IS_IDENTICAL_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = FastIsIdenticalFunction(op1, op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_IS_NOT_IDENTICAL_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var result types.ZendBool
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	result = FastIsNotIdenticalFunction(op1, op2)
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_IS_EQUAL_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	return zend_is_equal_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_NOT_EQUAL_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	return zend_is_not_equal_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_SMALLER_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	return zend_is_smaller_helper_SPEC(op1, op2, executeData)
}
func ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	var d1 float64
	var d2 float64
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	return zend_is_smaller_or_equal_helper_SPEC(op1, op2, executeData)
}
func ZEND_SPACESHIP_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	CompareFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_BW_OR_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	return zend_bw_or_helper_SPEC(op1, op2, executeData)
}
func ZEND_BW_AND_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	return zend_bw_and_helper_SPEC(op1, op2, executeData)
}
func ZEND_BW_XOR_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	return zend_bw_xor_helper_SPEC(op1, op2, executeData)
}
func ZEND_BOOL_XOR_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = RT_CONSTANT(opline, opline.GetOp2())
	BooleanXorFunction(EX_VAR(opline.GetResult().GetVar()), op1, op2)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_R_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var dim *types.Zval
	var value *types.Zval
	container = RT_CONSTANT(opline, opline.GetOp1())
	dim = RT_CONSTANT(opline, opline.GetOp2())

	{
		zend_fetch_dimension_address_read_R(container, dim, IS_CONST, opline, executeData)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_IS_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	container = RT_CONSTANT(opline, opline.GetOp1())
	zend_fetch_dimension_address_read_IS(container, RT_CONSTANT(opline, opline.GetOp2()), IS_CONST, opline, executeData)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	} else {
		return ZEND_FETCH_DIM_R_SPEC_CONST_CONST_HANDLER(executeData)
	}
}
func ZEND_FETCH_OBJ_R_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var offset *types.Zval
	offset = RT_CONSTANT(opline, opline.GetOp2())
	ZendWrongPropertyRead(offset)
	EX_VAR(opline.GetResult().GetVar()).SetNull()
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_IS_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	EX_VAR(opline.GetResult().GetVar()).SetNull()
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		/* Behave like FETCH_OBJ_W */
		return zend_use_tmp_in_write_context_helper_SPEC(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_CONST_CONST_HANDLER(executeData)
	}
}
func ZEND_FETCH_LIST_R_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	container = RT_CONSTANT(opline, opline.GetOp1())
	zend_fetch_dimension_address_LIST_r(container, RT_CONSTANT(opline, opline.GetOp2()), IS_CONST, opline, executeData)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FAST_CONCAT_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var op1 *types.Zval
	var op2 *types.Zval
	op1 = RT_CONSTANT(opline, opline.GetOp1())
	op2 = RT_CONSTANT(opline, opline.GetOp2())

	var op1_str *types.String = op1.GetStr()
	var op2_str *types.String = op2.GetStr()
	var str *types.String

	str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
	memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
	memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
	EX_VAR(opline.GetResult().GetVar()).SetString(str)

	return ZEND_VM_NEXT_OPCODE(executeData, opline)
	op1_str = op1.GetStr()
	op2_str = op2.GetStr()

	str = types.ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
	memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
	memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
	EX_VAR(opline.GetResult().GetVar()).SetString(str)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_INIT_METHOD_CALL_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var free_op1 ZendFreeOp
	var object *types.Zval
	var fbc types.IFunction
	var called_scope *types.ClassEntry
	var obj *types.ZendObject
	var call *ZendExecuteData
	var call_info uint32
	object = RT_CONSTANT(opline, opline.GetOp1())
	function_name = RT_CONSTANT(opline, opline.GetOp2())
	ZendInvalidMethodCall(object, function_name)
	return 0
	obj = object.GetObj()
	called_scope = obj.GetCe()
	if CACHED_PTR(opline.GetResult().GetNum()) == called_scope {
		fbc = CACHED_PTR(opline.GetResult().GetNum() + b.SizeOf("void *"))
	} else {
		var orig_obj *types.ZendObject = obj
		{
			function_name = RT_CONSTANT(opline, opline.GetOp2())
		}

		/* First, locate the function. */

		fbc = obj.GetHandlers().GetGetMethod()(&obj, function_name.GetStr(), b.CondF1(IS_CONST == IS_CONST, func() *types.Zval { return RT_CONSTANT(opline, opline.GetOp2()) + 1 }, nil))
		if fbc == nil {
			if EG__().GetException() == nil {
				ZendUndefinedMethod(obj.GetCe(), function_name.GetStr())
			}
			return 0
		}
		if fbc.GetType() <= ZEND_USER_FUNCTION && !fbc.HasFnFlags(AccCallViaTrampoline|AccNeverCache) && obj == orig_obj {
			CACHE_POLYMORPHIC_PTR(opline.GetResult().GetNum(), called_scope, fbc)
		}
		/* Reset "object" to trigger reference counting */

		/* Reset "object" to trigger reference counting */

		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	}
	call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
	if fbc.IsStatic() {
		/* call static method */

		obj = (*types.ZendObject)(called_scope)
		call_info = ZEND_CALL_NESTED_FUNCTION
	}

	/* CV may be changed indirectly (e.g. when it's a reference) */

	/* CV may be changed indirectly (e.g. when it's a reference) */

	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), obj)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var ce *types.ClassEntry
	var call_info uint32
	var fbc types.IFunction
	var call *ZendExecuteData
	{

		/* no function found. try a static method in class */

		ce = CACHED_PTR(opline.GetResult().GetNum())
		if ce == nil {
			ce = ZendFetchClassByName(RT_CONSTANT(opline, opline.GetOp1()).GetStr(), (RT_CONSTANT(opline, opline.GetOp1()) + 1).GetStr(), ZEND_FETCH_CLASS_DEFAULT|ZEND_FETCH_CLASS_EXCEPTION)
			if ce == nil {
				b.Assert(EG__().GetException() != nil)
				return 0
			}
		}
	}

	if b.Assign(&fbc, CACHED_PTR(opline.GetResult().GetNum()+b.SizeOf("void *"))) != nil {
	} else {
		function_name = RT_CONSTANT(opline, opline.GetOp2())
		if ce.GetGetStaticMethod() != nil {
			fbc = ce.GetGetStaticMethod()(ce, function_name.GetStr())
		} else {
			fbc = ZendStdGetStaticMethod(ce, function_name.GetStr(), b.CondF1(IS_CONST == IS_CONST, func() *types.Zval { return RT_CONSTANT(opline, opline.GetOp2()) + 1 }, nil))
		}
		if fbc == nil {
			if EG__().GetException() == nil {
				ZendUndefinedMethod(ce, function_name.GetStr())
			}
			return 0
		}
		if fbc.GetType() <= ZEND_USER_FUNCTION && !fbc.HasFnFlags(AccCallViaTrampoline|AccNeverCache) {
			CACHE_POLYMORPHIC_PTR(opline.GetResult().GetNum(), ce, fbc)
		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	}

	if !fbc.IsStatic() {
		if executeData.GetThis().IsObject() && InstanceofFunction(types.Z_OBJCE(executeData.GetThis()), ce) != 0 {
			ce = (*types.ClassEntry)(executeData.GetThis().GetObj())
			call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
		} else {
			ZendNonStaticMethodCall(fbc)
			if EG__().GetException() != nil {
				return 0
			}
			goto check_parent_and_self
		}
	} else {
	check_parent_and_self:
		/* previous opcode is ZEND_FETCH_CLASS */

		call_info = ZEND_CALL_NESTED_FUNCTION
	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), ce)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_INIT_USER_CALL_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var fcc types.ZendFcallInfoCache
	var error *byte = nil
	var func_ types.IFunction
	var object_or_called_scope any
	var call *ZendExecuteData
	var call_info uint32 = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_DYNAMIC
	function_name = RT_CONSTANT(opline, opline.GetOp2())
	if ZendIsCallableEx(function_name, nil, 0, nil, &fcc, &error) != 0 {
		func_ = fcc.GetFunctionHandler()
		if error != nil {
			Efree(error)

			/* This is the only soft error is_callable() can generate */

			ZendNonStaticMethodCall(func_)
			if EG__().GetException() != nil {
				return 0
			}
		}
		object_or_called_scope = fcc.GetCalledScope()
		if func_.IsClosure() {

			/* Delay closure destruction until its invocation */

			ZEND_CLOSURE_OBJECT(func_).AddRefcount()
			call_info |= ZEND_CALL_CLOSURE
			if func_.IsFakeClosure() {
				call_info |= ZEND_CALL_FAKE_CLOSURE
			}
			if fcc.GetObject() != nil {
				object_or_called_scope = fcc.GetObject()
				call_info |= ZEND_CALL_HAS_THIS
			}
		} else if fcc.GetObject() != nil {
			fcc.GetObject().AddRefcount()
			object_or_called_scope = fcc.GetObject()
			call_info |= ZEND_CALL_RELEASE_THIS | ZEND_CALL_HAS_THIS
		}
		if func_.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(func_.GetOpArray())) {
			InitFuncRunTimeCache(func_.GetOpArray())
		}
	} else {
		faults.InternalTypeError(executeData.IsCallUseStrictTypes(), "%s() expects parameter 1 to be a valid callback, %s", RT_CONSTANT(opline, opline.GetOp1()).GetStr().GetVal(), error)
		Efree(error)
		if EG__().GetException() != nil {
			return 0
		}
		func_ = (types.IFunction)(&ZendPassFunction)
		object_or_called_scope = nil
	}
	call = ZendVmStackPushCallFrame(call_info, func_, opline.GetExtendedValue(), object_or_called_scope)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_FETCH_CLASS_CONSTANT_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var ce *types.ClassEntry
	var scope *types.ClassEntry
	var c *ZendClassConstant
	var value *types.Zval
	var zv *types.Zval
	var opline *ZendOp = executeData.GetOpline()
	for {
		{
			if CACHED_PTR(opline.GetExtendedValue() + b.SizeOf("void *")) {
				value = CACHED_PTR(opline.GetExtendedValue() + b.SizeOf("void *"))
				break
			} else if CACHED_PTR(opline.GetExtendedValue()) {
				ce = CACHED_PTR(opline.GetExtendedValue())
			} else {
				ce = ZendFetchClassByName(RT_CONSTANT(opline, opline.GetOp1()).GetStr(), (RT_CONSTANT(opline, opline.GetOp1()) + 1).GetStr(), ZEND_FETCH_CLASS_DEFAULT|ZEND_FETCH_CLASS_EXCEPTION)
				if ce == nil {
					b.Assert(EG__().GetException() != nil)
					EX_VAR(opline.GetResult().GetVar()).SetUndef()
					return 0
				}
			}
		}

		zv = ce.GetConstantsTable().KeyFind(RT_CONSTANT(opline, opline.GetOp2()).GetStr().GetStr())
		if zv != nil {
			c = zv.GetPtr()
			scope = executeData.GetFunc().op_array.scope
			if ZendVerifyConstAccess(c, scope) == 0 {
				faults.ThrowError(nil, "Cannot access %s const %s::%s", ZendVisibilityString(c.GetValue().GetAccessFlags()), ce.GetName().GetVal(), RT_CONSTANT(opline, opline.GetOp2()).GetStr().GetVal())
				EX_VAR(opline.GetResult().GetVar()).SetUndef()
				return 0
			}
			value = c.GetValue()
			if value.IsConstant() {
				ZvalUpdateConstantEx(value, c.GetCe())
				if EG__().GetException() != nil {
					EX_VAR(opline.GetResult().GetVar()).SetUndef()
					return 0
				}
			}
			CACHE_POLYMORPHIC_PTR(opline.GetExtendedValue(), ce, value)
		} else {
			faults.ThrowError(nil, "Undefined class constant '%s'", RT_CONSTANT(opline, opline.GetOp2()).GetStr().GetVal())
			EX_VAR(opline.GetResult().GetVar()).SetUndef()
			return 0
		}
		break
	}
	types.ZVAL_COPY_OR_DUP(EX_VAR(opline.GetResult().GetVar()), value)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var expr_ptr *types.Zval
	var new_expr types.Zval

	expr_ptr = RT_CONSTANT(opline, opline.GetOp1())
	expr_ptr.TryAddRefcount()
	{
		var offset *types.Zval = RT_CONSTANT(opline, opline.GetOp2())
		var str *types.String
		var hval ZendUlong
	add_again:
		if offset.IsString() {
			str = offset.GetStr()
		str_index:
			EX_VAR(opline.GetResult().GetVar()).GetArr().KeyUpdate(str.GetStr(), expr_ptr)
		} else if offset.IsLong() {
			hval = offset.GetLval()
		num_index:
			EX_VAR(opline.GetResult().GetVar()).GetArr().IndexUpdate(hval, expr_ptr)
		} else if offset.IsNull() {
			str = types.ZSTR_EMPTY_ALLOC()
			goto str_index
		} else if offset.IsDouble() {
			hval = DvalToLval(offset.GetDval())
			goto num_index
		} else if offset.IsFalse() {
			hval = 0
			goto num_index
		} else if offset.IsTrue() {
			hval = 1
			goto num_index
		} else if offset.IsResource() {
			ZendUseResourceAsOffset(offset)
			hval = types.Z_RES_HANDLE_P(offset)
			goto num_index
		} else {
			ZendIllegalOffset()
			ZvalPtrDtorNogc(expr_ptr)
		}
	}

	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_INIT_ARRAY_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = EX_VAR(opline.GetResult().GetVar())
	{
		size = opline.GetExtendedValue() >> ZEND_ARRAY_SIZE_SHIFT
		array.SetArray(types.NewArray(size))

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & ZEND_ARRAY_NOT_PACKED) != 0 {
			types.ZendHashRealInitMixed(array.GetArr())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_CONST_HANDLER(executeData)
	}

}
func ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var result int
	var hval ZendUlong
	var offset *types.Zval
	container = RT_CONSTANT(opline, opline.GetOp1())
	offset = RT_CONSTANT(opline, opline.GetOp2())
	if container.IsArray() {
		var ht *types.Array
		var value *types.Zval
		var str *types.String
	isset_dim_obj_array:
		ht = container.GetArr()
	isset_again:
		if offset.IsString() {
			str = offset.GetStr()
			value = types.ZendHashFindInd(ht, str.GetStr())
		} else if offset.IsLong() {
			hval = offset.GetLval()
		num_index_prop:
			value = ht.IndexFind(hval)
		} else {
			value = ZendFindArrayDimSlow(ht, offset, executeData)
			if EG__().GetException() != nil {
				result = 0
				goto isset_dim_obj_exit
			}
		}
		if (opline.GetExtendedValue() & ZEND_ISEMPTY) == 0 {

			/* > IS_NULL means not IS_UNDEF and not IS_NULL */

			result = value != nil && value.GetType() > types.IS_NULL && (!(value.IsReference()) || types.Z_REFVAL_P(value).GetType() != types.IS_NULL)
			{

				/* avoid exception check */

				ZEND_VM_SMART_BRANCH(result, 0)
				types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else {
			result = value == nil || IZendIsTrue(value) == 0
		}
		goto isset_dim_obj_exit
	} else {
	}

	if offset.GetU2Extra() == ZEND_EXTRA_VALUE {
		offset++
	}
	if (opline.GetExtendedValue() & ZEND_ISEMPTY) == 0 {
		result = ZendIssetDimSlow(container, offset, executeData)
	} else {
		result = ZendIsemptyDimSlow(container, offset, executeData)
	}
isset_dim_obj_exit:
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var result int
	var offset *types.Zval
	container = RT_CONSTANT(opline, opline.GetOp1())
	offset = RT_CONSTANT(opline, opline.GetOp2())
	{

		{
			result = opline.GetExtendedValue() & ZEND_ISEMPTY
			goto isset_object_finish
		}
	}
	result = opline.GetExtendedValue()&ZEND_ISEMPTY ^ types.Z_OBJ_HT_P(container).GetHasProperty()(container, offset, opline.GetExtendedValue()&ZEND_ISEMPTY, b.CondF1(IS_CONST == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_ISEMPTY) }, nil))
isset_object_finish:
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(EX_VAR(opline.GetResult().GetVar()), result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
