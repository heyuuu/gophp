package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

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
	object = opline.Const1()
	function_name = opline.Const2()
	ZendInvalidMethodCall(object, function_name)
	return 0
	obj = object.GetObj()
	called_scope = obj.GetCe()
	if CACHED_PTR(opline.GetResult().GetNum()) == called_scope {
		fbc = CACHED_PTR(opline.GetResult().GetNum() + b.SizeOf("void *"))
	} else {
		var orig_obj *types.ZendObject = obj
		{
			function_name = opline.Const2()
		}

		/* First, locate the function. */

		fbc = obj.GetHandlers().GetGetMethod()(&obj, function_name.GetStr(), b.CondF1(IS_CONST == IS_CONST, func() *types.Zval { return opline.Const2() + 1 }, nil))
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
func ZEND_INIT_METHOD_CALL_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var object *types.Zval
	var fbc types.IFunction
	var called_scope *types.ClassEntry
	var obj *types.ZendObject
	var call *ZendExecuteData
	var call_info uint32
	object = opline.Const1()
	{
		function_name = opline.Op2()
	}
	if function_name.GetType() != types.IS_STRING {
		for {
			if function_name.IsReference() {
				function_name = types.Z_REFVAL_P(function_name)
				if function_name.IsString() {
					break
				}
			} else if function_name.IsUndef() {
				ZVAL_UNDEFINED_OP2(executeData)
				if EG__().GetException() != nil {
					return 0
				}
			}
			faults.ThrowError(nil, "Method name must be a string")
			// ZvalPtrDtorNogc(free_op2)
			return 0
			break
		}
	}
	{
		for {
			{
				ZendInvalidMethodCall(object, function_name)
				// ZvalPtrDtorNogc(free_op2)
				return 0
			}
			break
		}
	}
	obj = object.GetObj()
	called_scope = obj.GetCe()

	{
		var orig_obj *types.ZendObject = obj
		/* First, locate the function. */

		fbc = obj.GetHandlers().GetGetMethod()(&obj, function_name.GetStr(), b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *types.Zval { return opline.Const2() + 1 }, nil))
		if fbc == nil {
			if EG__().GetException() == nil {
				ZendUndefinedMethod(obj.GetCe(), function_name.GetStr())
			}
			// ZvalPtrDtorNogc(free_op2)
			return 0
		}
		/* Reset "object" to trigger reference counting */

		/* Reset "object" to trigger reference counting */

		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	}
	{
		// ZvalPtrDtorNogc(free_op2)
	}
	call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
	if fbc.IsStatic() {
		/* call static method */

		obj = (*types.ZendObject)(called_scope)
		call_info = ZEND_CALL_NESTED_FUNCTION
	} else {
	}

	/* CV may be changed indirectly (e.g. when it's a reference) */

	/* CV may be changed indirectly (e.g. when it's a reference) */

	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), obj)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_INIT_METHOD_CALL_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var free_op1 ZendFreeOp
	var object *types.Zval
	var fbc types.IFunction
	var called_scope *types.ClassEntry
	var obj *types.ZendObject
	var call *ZendExecuteData
	var call_info uint32
	object = opline.Const1()
	{
		function_name = opline.Op2()
	}
	if function_name.GetType() != types.IS_STRING {
		for {
			if function_name.IsReference() {
				function_name = types.Z_REFVAL_P(function_name)
				if function_name.IsString() {
					break
				}
			} else if function_name.IsUndef() {
				ZVAL_UNDEFINED_OP2(executeData)
				if EG__().GetException() != nil {
					return 0
				}
			}
			faults.ThrowError(nil, "Method name must be a string")
			return 0
			break
		}
	}
	{
		for {
			{
				ZendInvalidMethodCall(object, function_name)
				return 0
			}
			break
		}
	}
	obj = object.GetObj()
	called_scope = obj.GetCe()

	{
		var orig_obj *types.ZendObject = obj
		/* First, locate the function. */

		fbc = obj.GetHandlers().GetGetMethod()(&obj, function_name.GetStr(), b.CondF1(IS_CV == IS_CONST, func() *types.Zval { return opline.Const2() + 1 }, nil))
		if fbc == nil {
			if EG__().GetException() == nil {
				ZendUndefinedMethod(obj.GetCe(), function_name.GetStr())
			}
			return 0
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
	} else {
	}

	/* CV may be changed indirectly (e.g. when it's a reference) */

	/* CV may be changed indirectly (e.g. when it's a reference) */

	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), obj)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_INIT_METHOD_CALL_SPEC_TMPVAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var free_op1 ZendFreeOp
	var object *types.Zval
	var fbc types.IFunction
	var called_scope *types.ClassEntry
	var obj *types.ZendObject
	var call *ZendExecuteData
	var call_info uint32
	object = opline.Op1()
	{
		for {
			if object.GetType() != types.IS_OBJECT {
				if object.IsReference() {
					object = types.Z_REFVAL_P(object)
					if object.IsObject() {
						break
					}
				}
				if object.IsUndef() {
					object = ZVAL_UNDEFINED_OP1(executeData)
					if EG__().GetException() != nil {
						return 0
					}
				}
				{
					function_name = opline.Const2()
				}
				ZendInvalidMethodCall(object, function_name)
				// ZvalPtrDtorNogc(free_op1)
				return 0
			}
			break
		}
	}
	obj = object.GetObj()
	called_scope = obj.GetCe()
	if CACHED_PTR(opline.GetResult().GetNum()) == called_scope {
		fbc = CACHED_PTR(opline.GetResult().GetNum() + b.SizeOf("void *"))
	} else {
		var orig_obj *types.ZendObject = obj
		{
			function_name = opline.Const2()
		}

		/* First, locate the function. */

		fbc = obj.GetHandlers().GetGetMethod()(&obj, function_name.GetStr(), b.CondF1(IS_CONST == IS_CONST, func() *types.Zval { return opline.Const2() + 1 }, nil))
		if fbc == nil {
			if EG__().GetException() == nil {
				ZendUndefinedMethod(obj.GetCe(), function_name.GetStr())
			}
			// ZvalPtrDtorNogc(free_op1)
			return 0
		}
		if fbc.GetType() <= ZEND_USER_FUNCTION && !fbc.HasFnFlags(AccCallViaTrampoline|AccNeverCache) && obj == orig_obj {
			CACHE_POLYMORPHIC_PTR(opline.GetResult().GetNum(), called_scope, fbc)
		}
		if obj != orig_obj {

			/* Reset "object" to trigger reference counting */

			object = nil

			/* Reset "object" to trigger reference counting */

		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	}
	call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
	if fbc.IsStatic() {
		// ZvalPtrDtorNogc(free_op1)
		if EG__().GetException() != nil {
			return 0
		}

		/* call static method */

		obj = (*types.ZendObject)(called_scope)
		call_info = ZEND_CALL_NESTED_FUNCTION
	} else {
		{
			// 			obj.AddRefcount()
		}

		/* CV may be changed indirectly (e.g. when it's a reference) */

		call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS | ZEND_CALL_RELEASE_THIS

		/* CV may be changed indirectly (e.g. when it's a reference) */

	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), obj)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_INIT_METHOD_CALL_SPEC_TMPVAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var object *types.Zval
	var fbc types.IFunction
	var called_scope *types.ClassEntry
	var obj *types.ZendObject
	var call *ZendExecuteData
	var call_info uint32
	object = opline.Op1()
	{
		function_name = opline.Op2()
	}
	if function_name.GetType() != types.IS_STRING {
		for {
			if function_name.IsReference() {
				function_name = types.Z_REFVAL_P(function_name)
				if function_name.IsString() {
					break
				}
			} else if function_name.IsUndef() {
				ZVAL_UNDEFINED_OP2(executeData)
				if EG__().GetException() != nil {
					// ZvalPtrDtorNogc(free_op1)
					return 0
				}
			}
			faults.ThrowError(nil, "Method name must be a string")
			// ZvalPtrDtorNogc(free_op2)
			// ZvalPtrDtorNogc(free_op1)
			return 0
			break
		}
	}
	{
		for {
			if object.GetType() != types.IS_OBJECT {
				if object.IsReference() {
					object = types.Z_REFVAL_P(object)
					if object.IsObject() {
						break
					}
				}
				if object.IsUndef() {
					object = ZVAL_UNDEFINED_OP1(executeData)
					if EG__().GetException() != nil {
						{
							// ZvalPtrDtorNogc(free_op2)
						}
						return 0
					}
				}
				ZendInvalidMethodCall(object, function_name)
				// ZvalPtrDtorNogc(free_op2)
				// ZvalPtrDtorNogc(free_op1)
				return 0
			}
			break
		}
	}
	obj = object.GetObj()
	called_scope = obj.GetCe()

	{
		var orig_obj *types.ZendObject = obj
		/* First, locate the function. */

		fbc = obj.GetHandlers().GetGetMethod()(&obj, function_name.GetStr(), b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *types.Zval { return opline.Const2() + 1 }, nil))
		if fbc == nil {
			if EG__().GetException() == nil {
				ZendUndefinedMethod(obj.GetCe(), function_name.GetStr())
			}
			// ZvalPtrDtorNogc(free_op2)
			// ZvalPtrDtorNogc(free_op1)
			return 0
		}
		if obj != orig_obj {

			/* Reset "object" to trigger reference counting */

			object = nil

			/* Reset "object" to trigger reference counting */

		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	}
	{
		// ZvalPtrDtorNogc(free_op2)
	}
	call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
	if fbc.IsStatic() {
		// ZvalPtrDtorNogc(free_op1)
		if EG__().GetException() != nil {
			return 0
		}

		/* call static method */

		obj = (*types.ZendObject)(called_scope)
		call_info = ZEND_CALL_NESTED_FUNCTION
	} else {
		{
			// 			obj.AddRefcount()
		}

		/* CV may be changed indirectly (e.g. when it's a reference) */

		call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS | ZEND_CALL_RELEASE_THIS

		/* CV may be changed indirectly (e.g. when it's a reference) */

	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), obj)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_INIT_METHOD_CALL_SPEC_TMPVAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var free_op1 ZendFreeOp
	var object *types.Zval
	var fbc types.IFunction
	var called_scope *types.ClassEntry
	var obj *types.ZendObject
	var call *ZendExecuteData
	var call_info uint32
	object = opline.Op1()
	{
		function_name = opline.Op2()
	}
	if function_name.GetType() != types.IS_STRING {
		for {
			if function_name.IsReference() {
				function_name = types.Z_REFVAL_P(function_name)
				if function_name.IsString() {
					break
				}
			} else if function_name.IsUndef() {
				ZVAL_UNDEFINED_OP2(executeData)
				if EG__().GetException() != nil {
					// ZvalPtrDtorNogc(free_op1)
					return 0
				}
			}
			faults.ThrowError(nil, "Method name must be a string")
			// ZvalPtrDtorNogc(free_op1)
			return 0
			break
		}
	}
	{
		for {
			if object.GetType() != types.IS_OBJECT {
				if object.IsReference() {
					object = types.Z_REFVAL_P(object)
					if object.IsObject() {
						break
					}
				}
				if object.IsUndef() {
					object = ZVAL_UNDEFINED_OP1(executeData)
					if EG__().GetException() != nil {
						return 0
					}
				}
				ZendInvalidMethodCall(object, function_name)
				// ZvalPtrDtorNogc(free_op1)
				return 0
			}
			break
		}
	}
	obj = object.GetObj()
	called_scope = obj.GetCe()

	{
		var orig_obj *types.ZendObject = obj
		/* First, locate the function. */

		fbc = obj.GetHandlers().GetGetMethod()(&obj, function_name.GetStr(), b.CondF1(IS_CV == IS_CONST, func() *types.Zval { return opline.Const2() + 1 }, nil))
		if fbc == nil {
			if EG__().GetException() == nil {
				ZendUndefinedMethod(obj.GetCe(), function_name.GetStr())
			}
			// ZvalPtrDtorNogc(free_op1)
			return 0
		}
		if obj != orig_obj {

			/* Reset "object" to trigger reference counting */

			object = nil

			/* Reset "object" to trigger reference counting */

		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	}
	call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
	if fbc.IsStatic() {
		// ZvalPtrDtorNogc(free_op1)
		if EG__().GetException() != nil {
			return 0
		}

		/* call static method */

		obj = (*types.ZendObject)(called_scope)
		call_info = ZEND_CALL_NESTED_FUNCTION
	} else {
		{
			// 			obj.AddRefcount()
		}

		/* CV may be changed indirectly (e.g. when it's a reference) */

		call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS | ZEND_CALL_RELEASE_THIS

		/* CV may be changed indirectly (e.g. when it's a reference) */

	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), obj)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_INIT_METHOD_CALL_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var free_op1 ZendFreeOp
	var object *types.Zval
	var fbc types.IFunction
	var called_scope *types.ClassEntry
	var obj *types.ZendObject
	var call *ZendExecuteData
	var call_info uint32
	object = &(executeData.GetThis())
	if object.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	obj = object.GetObj()
	called_scope = obj.GetCe()
	if CACHED_PTR(opline.GetResult().GetNum()) == called_scope {
		fbc = CACHED_PTR(opline.GetResult().GetNum() + b.SizeOf("void *"))
	} else {
		var orig_obj *types.ZendObject = obj
		{
			function_name = opline.Const2()
		}

		/* First, locate the function. */

		fbc = obj.GetHandlers().GetGetMethod()(&obj, function_name.GetStr(), b.CondF1(IS_CONST == IS_CONST, func() *types.Zval { return opline.Const2() + 1 }, nil))
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
	} else {
	}

	/* CV may be changed indirectly (e.g. when it's a reference) */

	/* CV may be changed indirectly (e.g. when it's a reference) */

	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), obj)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_INIT_METHOD_CALL_SPEC_UNUSED_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var object *types.Zval
	var fbc types.IFunction
	var called_scope *types.ClassEntry
	var obj *types.ZendObject
	var call *ZendExecuteData
	var call_info uint32
	object = &(executeData.GetThis())
	if object.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	{
		function_name = opline.Op2()
	}
	if function_name.GetType() != types.IS_STRING {
		for {
			if function_name.IsReference() {
				function_name = types.Z_REFVAL_P(function_name)
				if function_name.IsString() {
					break
				}
			} else if function_name.IsUndef() {
				ZVAL_UNDEFINED_OP2(executeData)
				if EG__().GetException() != nil {
					return 0
				}
			}
			faults.ThrowError(nil, "Method name must be a string")
			// ZvalPtrDtorNogc(free_op2)
			return 0
			break
		}
	}
	obj = object.GetObj()
	called_scope = obj.GetCe()

	{
		var orig_obj *types.ZendObject = obj
		/* First, locate the function. */

		fbc = obj.GetHandlers().GetGetMethod()(&obj, function_name.GetStr(), b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *types.Zval { return opline.Const2() + 1 }, nil))
		if fbc == nil {
			if EG__().GetException() == nil {
				ZendUndefinedMethod(obj.GetCe(), function_name.GetStr())
			}
			// ZvalPtrDtorNogc(free_op2)
			return 0
		}
		/* Reset "object" to trigger reference counting */

		/* Reset "object" to trigger reference counting */

		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	}
	{
		// ZvalPtrDtorNogc(free_op2)
	}
	call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
	if fbc.IsStatic() {
		/* call static method */

		obj = (*types.ZendObject)(called_scope)
		call_info = ZEND_CALL_NESTED_FUNCTION
	} else {
	}

	/* CV may be changed indirectly (e.g. when it's a reference) */

	/* CV may be changed indirectly (e.g. when it's a reference) */

	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), obj)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_INIT_METHOD_CALL_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var free_op1 ZendFreeOp
	var object *types.Zval
	var fbc types.IFunction
	var called_scope *types.ClassEntry
	var obj *types.ZendObject
	var call *ZendExecuteData
	var call_info uint32
	object = &(executeData.GetThis())
	if object.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	{
		function_name = opline.Op2()
	}
	if function_name.GetType() != types.IS_STRING {
		for {
			if function_name.IsReference() {
				function_name = types.Z_REFVAL_P(function_name)
				if function_name.IsString() {
					break
				}
			} else if function_name.IsUndef() {
				ZVAL_UNDEFINED_OP2(executeData)
				if EG__().GetException() != nil {
					return 0
				}
			}
			faults.ThrowError(nil, "Method name must be a string")
			return 0
			break
		}
	}
	obj = object.GetObj()
	called_scope = obj.GetCe()

	{
		var orig_obj *types.ZendObject = obj
		/* First, locate the function. */

		fbc = obj.GetHandlers().GetGetMethod()(&obj, function_name.GetStr(), b.CondF1(IS_CV == IS_CONST, func() *types.Zval { return opline.Const2() + 1 }, nil))
		if fbc == nil {
			if EG__().GetException() == nil {
				ZendUndefinedMethod(obj.GetCe(), function_name.GetStr())
			}
			return 0
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
	} else {
	}

	/* CV may be changed indirectly (e.g. when it's a reference) */

	/* CV may be changed indirectly (e.g. when it's a reference) */

	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), obj)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_INIT_METHOD_CALL_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var free_op1 ZendFreeOp
	var object *types.Zval
	var fbc types.IFunction
	var called_scope *types.ClassEntry
	var obj *types.ZendObject
	var call *ZendExecuteData
	var call_info uint32
	object = opline.Op1()
	{
		for {
			if object.GetType() != types.IS_OBJECT {
				if object.IsReference() {
					object = types.Z_REFVAL_P(object)
					if object.IsObject() {
						break
					}
				}
				if object.IsUndef() {
					object = ZVAL_UNDEFINED_OP1(executeData)
					if EG__().GetException() != nil {
						return 0
					}
				}
				{
					function_name = opline.Const2()
				}
				ZendInvalidMethodCall(object, function_name)
				return 0
			}
			break
		}
	}
	obj = object.GetObj()
	called_scope = obj.GetCe()
	if CACHED_PTR(opline.GetResult().GetNum()) == called_scope {
		fbc = CACHED_PTR(opline.GetResult().GetNum() + b.SizeOf("void *"))
	} else {
		var orig_obj *types.ZendObject = obj
		{
			function_name = opline.Const2()
		}

		/* First, locate the function. */

		fbc = obj.GetHandlers().GetGetMethod()(&obj, function_name.GetStr(), b.CondF1(IS_CONST == IS_CONST, func() *types.Zval { return opline.Const2() + 1 }, nil))
		if fbc == nil {
			if EG__().GetException() == nil {
				ZendUndefinedMethod(obj.GetCe(), function_name.GetStr())
			}
			return 0
		}
		if fbc.GetType() <= ZEND_USER_FUNCTION && !fbc.HasFnFlags(AccCallViaTrampoline|AccNeverCache) && obj == orig_obj {
			CACHE_POLYMORPHIC_PTR(opline.GetResult().GetNum(), called_scope, fbc)
		}
		if obj != orig_obj {

			/* Reset "object" to trigger reference counting */

			object = nil

			/* Reset "object" to trigger reference counting */

		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	}
	call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
	if fbc.IsStatic() {
		if EG__().GetException() != nil {
			return 0
		}

		/* call static method */

		obj = (*types.ZendObject)(called_scope)
		call_info = ZEND_CALL_NESTED_FUNCTION
	} else {
		{
			// 			obj.AddRefcount()
		}

		/* CV may be changed indirectly (e.g. when it's a reference) */

		call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS | ZEND_CALL_RELEASE_THIS

		/* CV may be changed indirectly (e.g. when it's a reference) */

	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), obj)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_INIT_METHOD_CALL_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var object *types.Zval
	var fbc types.IFunction
	var called_scope *types.ClassEntry
	var obj *types.ZendObject
	var call *ZendExecuteData
	var call_info uint32
	object = opline.Op1()
	{
		function_name = opline.Op2()
	}
	if function_name.GetType() != types.IS_STRING {
		for {
			if function_name.IsReference() {
				function_name = types.Z_REFVAL_P(function_name)
				if function_name.IsString() {
					break
				}
			} else if function_name.IsUndef() {
				ZVAL_UNDEFINED_OP2(executeData)
				if EG__().GetException() != nil {
					return 0
				}
			}
			faults.ThrowError(nil, "Method name must be a string")
			// ZvalPtrDtorNogc(free_op2)
			return 0
			break
		}
	}
	{
		for {
			if object.GetType() != types.IS_OBJECT {
				if object.IsReference() {
					object = types.Z_REFVAL_P(object)
					if object.IsObject() {
						break
					}
				}
				if object.IsUndef() {
					object = ZVAL_UNDEFINED_OP1(executeData)
					if EG__().GetException() != nil {
						{
							// ZvalPtrDtorNogc(free_op2)
						}
						return 0
					}
				}
				ZendInvalidMethodCall(object, function_name)
				// ZvalPtrDtorNogc(free_op2)
				return 0
			}
			break
		}
	}
	obj = object.GetObj()
	called_scope = obj.GetCe()

	{
		var orig_obj *types.ZendObject = obj
		/* First, locate the function. */

		fbc = obj.GetHandlers().GetGetMethod()(&obj, function_name.GetStr(), b.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *types.Zval { return opline.Const2() + 1 }, nil))
		if fbc == nil {
			if EG__().GetException() == nil {
				ZendUndefinedMethod(obj.GetCe(), function_name.GetStr())
			}
			// ZvalPtrDtorNogc(free_op2)
			return 0
		}
		if obj != orig_obj {

			/* Reset "object" to trigger reference counting */

			object = nil

			/* Reset "object" to trigger reference counting */

		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	}
	{
		// ZvalPtrDtorNogc(free_op2)
	}
	call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
	if fbc.IsStatic() {
		if EG__().GetException() != nil {
			return 0
		}

		/* call static method */

		obj = (*types.ZendObject)(called_scope)
		call_info = ZEND_CALL_NESTED_FUNCTION
	} else {
		{
			// 			obj.AddRefcount()
		}

		/* CV may be changed indirectly (e.g. when it's a reference) */

		call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS | ZEND_CALL_RELEASE_THIS

		/* CV may be changed indirectly (e.g. when it's a reference) */

	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), obj)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_INIT_METHOD_CALL_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var free_op1 ZendFreeOp
	var object *types.Zval
	var fbc types.IFunction
	var called_scope *types.ClassEntry
	var obj *types.ZendObject
	var call *ZendExecuteData
	var call_info uint32
	object = opline.Op1()
	{
		function_name = opline.Op2()
	}
	if function_name.GetType() != types.IS_STRING {
		for {
			if function_name.IsReference() {
				function_name = types.Z_REFVAL_P(function_name)
				if function_name.IsString() {
					break
				}
			} else if function_name.IsUndef() {
				ZVAL_UNDEFINED_OP2(executeData)
				if EG__().GetException() != nil {
					return 0
				}
			}
			faults.ThrowError(nil, "Method name must be a string")
			return 0
			break
		}
	}
	{
		for {
			if object.GetType() != types.IS_OBJECT {
				if object.IsReference() {
					object = types.Z_REFVAL_P(object)
					if object.IsObject() {
						break
					}
				}
				if object.IsUndef() {
					object = ZVAL_UNDEFINED_OP1(executeData)
					if EG__().GetException() != nil {
						return 0
					}
				}
				ZendInvalidMethodCall(object, function_name)
				return 0
			}
			break
		}
	}
	obj = object.GetObj()
	called_scope = obj.GetCe()

	{
		var orig_obj *types.ZendObject = obj
		/* First, locate the function. */

		fbc = obj.GetHandlers().GetGetMethod()(&obj, function_name.GetStr(), b.CondF1(IS_CV == IS_CONST, func() *types.Zval { return opline.Const2() + 1 }, nil))
		if fbc == nil {
			if EG__().GetException() == nil {
				ZendUndefinedMethod(obj.GetCe(), function_name.GetStr())
			}
			return 0
		}
		if obj != orig_obj {

			/* Reset "object" to trigger reference counting */

			object = nil

			/* Reset "object" to trigger reference counting */

		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	}
	call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
	if fbc.IsStatic() {
		if EG__().GetException() != nil {
			return 0
		}

		/* call static method */

		obj = (*types.ZendObject)(called_scope)
		call_info = ZEND_CALL_NESTED_FUNCTION
	} else {
		{
			// 			obj.AddRefcount()
		}

		/* CV may be changed indirectly (e.g. when it's a reference) */

		call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS | ZEND_CALL_RELEASE_THIS

		/* CV may be changed indirectly (e.g. when it's a reference) */

	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), obj)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
