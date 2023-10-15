package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func ZEND_INIT_STATIC_METHOD_CALL_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var ce *types.ClassEntry
	var call_info uint32
	var fbc types.IFunction
	var call *ZendExecuteData
	{

		/* no function found. try a static method in class */

		ce = CACHED_PTR(opline.GetResult().GetNum())
		if ce == nil {
			ce = ZendFetchClassByName(opline.Const1().String(), (opline.Const1() + 1).String(), ZEND_FETCH_CLASS_DEFAULT|ZEND_FETCH_CLASS_EXCEPTION)
			if ce == nil {
				b.Assert(EG__().HasException())
				return 0
			}
		}
	}

	if lang.Assign(&fbc, CACHED_PTR(opline.GetResult().GetNum()+b.SizeOf("void *"))) != nil {
	} else {
		function_name = opline.Const2()
		if ce.GetGetStaticMethod() != nil {
			fbc = ce.GetGetStaticMethod()(ce, function_name.GetStr())
		} else {
			fbc = ZendStdGetStaticMethod(ce, function_name.GetStr(), lang.CondF1(IS_CONST == IS_CONST, func() *types.Zval { return opline.Const2() + 1 }, nil))
		}
		if fbc == nil {
			if EG__().NoException() {
				ZendUndefinedMethod(ce, function_name.GetStr())
			}
			return 0
		}
		if fbc.GetType() <= ZEND_USER_FUNCTION && !fbc.HasFnFlags(types.AccCallViaTrampoline|types.AccNeverCache) {
			CACHE_POLYMORPHIC_PTR(opline.GetResult().GetNum(), ce, fbc)
		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	}

	if !fbc.IsStatic() {
		if executeData.InScope() && operators.InstanceofFunction(executeData.ThisClass(), ce) {
			ce = executeData.ThisClass()
			call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
		} else {
			ZendNonStaticMethodCall(fbc)
			if EG__().HasException() {
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
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var ce *types.ClassEntry
	var call_info uint32
	var fbc types.IFunction
	var call *ZendExecuteData
	{

		/* no function found. try a static method in class */

		ce = CACHED_PTR(opline.GetResult().GetNum())
		if ce == nil {
			ce = ZendFetchClassByName(opline.Const1().String(), (opline.Const1() + 1).String(), ZEND_FETCH_CLASS_DEFAULT|ZEND_FETCH_CLASS_EXCEPTION)
			if ce == nil {
				b.Assert(EG__().HasException())
				// ZvalPtrDtorNogc(opline.Op2())
				return 0
			}
			{
				CACHE_PTR(opline.GetResult().GetNum(), ce)
			}
		}
	}

	{
		var free_op2 ZendFreeOp
		function_name = opline.Op2()
		{
			if !function_name.IsString() {
				for {
					if function_name.IsRef() {
						function_name = types.Z_REFVAL_P(function_name)
						if function_name.IsString() {
							break
						}
					} else if function_name.IsUndef() {
						ZVAL_UNDEFINED_OP2(executeData)
						if EG__().HasException() {
							return 0
						}
					}
					faults.ThrowError(nil, "Function name must be a string")
					// ZvalPtrDtorNogc(free_op2)
					return 0
					break
				}
			}
		}
		if ce.GetGetStaticMethod() != nil {
			fbc = ce.GetGetStaticMethod()(ce, function_name.GetStr())
		} else {
			fbc = ZendStdGetStaticMethod(ce, function_name.GetStr(), lang.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *types.Zval { return opline.Const2() + 1 }, nil))
		}
		if fbc == nil {
			if EG__().NoException() {
				ZendUndefinedMethod(ce, function_name.GetStr())
			}
			// ZvalPtrDtorNogc(free_op2)
			return 0
		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
		{
			// ZvalPtrDtorNogc(free_op2)
		}
	}

	if !fbc.IsStatic() {
		if executeData.InScope() && operators.InstanceofFunction(executeData.ThisClass(), ce) {
			ce = executeData.ThisClass()
			call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
		} else {
			ZendNonStaticMethodCall(fbc)
			if EG__().HasException() {
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
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var ce *types.ClassEntry
	var call_info uint32
	var fbc types.IFunction
	var call *ZendExecuteData
	{

		/* no function found. try a static method in class */

		ce = CACHED_PTR(opline.GetResult().GetNum())
		if ce == nil {
			ce = ZendFetchClassByName(opline.Const1().String(), (opline.Const1() + 1).String(), ZEND_FETCH_CLASS_DEFAULT|ZEND_FETCH_CLASS_EXCEPTION)
			if ce == nil {
				b.Assert(EG__().HasException())
				return 0
			}
			{
				CACHE_PTR(opline.GetResult().GetNum(), ce)
			}
		}
	}

	{
		if ce.GetConstructor() == nil {
			faults.ThrowError(nil, "Cannot call constructor")
			return 0
		}
		if executeData.InScope() && executeData.ThisClass() != ce.GetConstructor().GetScope() && ce.GetConstructor().IsPrivate() {
			faults.ThrowError(nil, "Cannot call private %s::__construct()", ce.Name())
			return 0
		}
		fbc = ce.GetConstructor()
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	}
	if !fbc.IsStatic() {
		if executeData.InScope() && operators.InstanceofFunction(executeData.ThisClass(), ce) {
			ce = executeData.ThisClass()
			call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
		} else {
			ZendNonStaticMethodCall(fbc)
			if EG__().HasException() {
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
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var ce *types.ClassEntry
	var call_info uint32
	var fbc types.IFunction
	var call *ZendExecuteData
	{

		/* no function found. try a static method in class */

		ce = CACHED_PTR(opline.GetResult().GetNum())
		if ce == nil {
			ce = ZendFetchClassByName(opline.Const1().String(), (opline.Const1() + 1).String(), ZEND_FETCH_CLASS_DEFAULT|ZEND_FETCH_CLASS_EXCEPTION)
			if ce == nil {
				b.Assert(EG__().HasException())
				return 0
			}
			{
				CACHE_PTR(opline.GetResult().GetNum(), ce)
			}
		}
	}

	{
		function_name = opline.Op2()
		{
			if !function_name.IsString() {
				for {
					if function_name.IsRef() {
						function_name = types.Z_REFVAL_P(function_name)
						if function_name.IsString() {
							break
						}
					} else if function_name.IsUndef() {
						ZVAL_UNDEFINED_OP2(executeData)
						if EG__().HasException() {
							return 0
						}
					}
					faults.ThrowError(nil, "Function name must be a string")
					return 0
					break
				}
			}
		}
		if ce.GetGetStaticMethod() != nil {
			fbc = ce.GetGetStaticMethod()(ce, function_name.GetStr())
		} else {
			fbc = ZendStdGetStaticMethod(ce, function_name.GetStr(), lang.CondF1(IS_CV == IS_CONST, func() *types.Zval { return opline.Const2() + 1 }, nil))
		}
		if fbc == nil {
			if EG__().NoException() {
				ZendUndefinedMethod(ce, function_name.GetStr())
			}
			return 0
		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	}

	if !fbc.IsStatic() {
		if executeData.InScope() && operators.InstanceofFunction(executeData.ThisClass(), ce) {
			ce = executeData.ThisClass()
			call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
		} else {
			ZendNonStaticMethodCall(fbc)
			if EG__().HasException() {
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
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var ce *types.ClassEntry
	var call_info uint32
	var fbc types.IFunction
	var call *ZendExecuteData

	/* no function found. try a static method in class */

	{
		ce = opline.Op1().Class()
	}

	if CACHED_PTR(opline.GetResult().GetNum()) == ce {
		fbc = CACHED_PTR(opline.GetResult().GetNum() + b.SizeOf("void *"))
	} else {
		function_name = opline.Const2()
		if ce.GetGetStaticMethod() != nil {
			fbc = ce.GetGetStaticMethod()(ce, function_name.GetStr())
		} else {
			fbc = ZendStdGetStaticMethod(ce, function_name.GetStr(), lang.CondF1(IS_CONST == IS_CONST, func() *types.Zval { return opline.Const2() + 1 }, nil))
		}
		if fbc == nil {
			if EG__().NoException() {
				ZendUndefinedMethod(ce, function_name.GetStr())
			}
			return 0
		}
		if fbc.GetType() <= ZEND_USER_FUNCTION && !fbc.HasFnFlags(types.AccCallViaTrampoline|types.AccNeverCache) {
			CACHE_POLYMORPHIC_PTR(opline.GetResult().GetNum(), ce, fbc)
		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	}

	if !fbc.IsStatic() {
		if executeData.InScope() && operators.InstanceofFunction(executeData.ThisClass(), ce) {
			ce = executeData.ThisClass()
			call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
		} else {
			ZendNonStaticMethodCall(fbc)
			if EG__().HasException() {
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
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_VAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var ce *types.ClassEntry
	var call_info uint32
	var fbc types.IFunction
	var call *ZendExecuteData

	/* no function found. try a static method in class */

	{
		ce = opline.Op1().Class()
	}

	{
		var free_op2 ZendFreeOp
		function_name = opline.Op2()
		{
			if !function_name.IsString() {
				for {
					if function_name.IsRef() {
						function_name = types.Z_REFVAL_P(function_name)
						if function_name.IsString() {
							break
						}
					} else if function_name.IsUndef() {
						ZVAL_UNDEFINED_OP2(executeData)
						if EG__().HasException() {
							return 0
						}
					}
					faults.ThrowError(nil, "Function name must be a string")
					// ZvalPtrDtorNogc(free_op2)
					return 0
					break
				}
			}
		}
		if ce.GetGetStaticMethod() != nil {
			fbc = ce.GetGetStaticMethod()(ce, function_name.GetStr())
		} else {
			fbc = ZendStdGetStaticMethod(ce, function_name.GetStr(), lang.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *types.Zval { return opline.Const2() + 1 }, nil))
		}
		if fbc == nil {
			if EG__().NoException() {
				ZendUndefinedMethod(ce, function_name.GetStr())
			}
			// ZvalPtrDtorNogc(free_op2)
			return 0
		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
		{
			// ZvalPtrDtorNogc(free_op2)
		}
	}

	if !fbc.IsStatic() {
		if executeData.InScope() && operators.InstanceofFunction(executeData.ThisClass(), ce) {
			ce = executeData.ThisClass()
			call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
		} else {
			ZendNonStaticMethodCall(fbc)
			if EG__().HasException() {
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
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_VAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var ce *types.ClassEntry
	var call_info uint32
	var fbc types.IFunction
	var call *ZendExecuteData

	/* no function found. try a static method in class */

	{
		ce = opline.Op1().Class()
	}

	{
		if ce.GetConstructor() == nil {
			faults.ThrowError(nil, "Cannot call constructor")
			return 0
		}
		if executeData.InScope() && executeData.ThisClass() != ce.GetConstructor().GetScope() && ce.GetConstructor().IsPrivate() {
			faults.ThrowError(nil, "Cannot call private %s::__construct()", ce.Name())
			return 0
		}
		fbc = ce.GetConstructor()
		if fbc.IsUserFunction() && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	}
	if !fbc.IsStatic() {
		if executeData.InScope() && operators.InstanceofFunction(executeData.ThisClass(), ce) {
			ce = executeData.ThisClass()
			call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
		} else {
			ZendNonStaticMethodCall(fbc)
			if EG__().HasException() {
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
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_VAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var ce *types.ClassEntry
	var call_info uint32
	var fbc types.IFunction
	var call *ZendExecuteData

	/* no function found. try a static method in class */

	{
		ce = opline.Op1().Class()
	}

	{
		function_name = opline.Op2()
		{
			if !function_name.IsString() {
				for {
					if function_name.IsRef() {
						function_name = types.Z_REFVAL_P(function_name)
						if function_name.IsString() {
							break
						}
					} else if function_name.IsUndef() {
						ZVAL_UNDEFINED_OP2(executeData)
						if EG__().HasException() {
							return 0
						}
					}
					faults.ThrowError(nil, "Function name must be a string")
					return 0
					break
				}
			}
		}
		if ce.GetGetStaticMethod() != nil {
			fbc = ce.GetGetStaticMethod()(ce, function_name.GetStr())
		} else {
			fbc = ZendStdGetStaticMethod(ce, function_name.GetStr(), lang.CondF1(IS_CV == IS_CONST, func() *types.Zval { return opline.Const2() + 1 }, nil))
		}
		if fbc == nil {
			if EG__().NoException() {
				ZendUndefinedMethod(ce, function_name.GetStr())
			}
			return 0
		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	}

	if !fbc.IsStatic() {
		if executeData.InScope() && operators.InstanceofFunction(executeData.ThisClass(), ce) {
			ce = executeData.ThisClass()
			call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
		} else {
			ZendNonStaticMethodCall(fbc)
			if EG__().HasException() {
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
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var ce *types.ClassEntry
	var call_info uint32
	var fbc types.IFunction
	var call *ZendExecuteData

	/* no function found. try a static method in class */

	{
		ce = ZendFetchClass(nil, opline.GetOp1().GetNum())
		if ce == nil {
			b.Assert(EG__().HasException())
			return 0
		}
	}

	if CACHED_PTR(opline.GetResult().GetNum()) == ce {
		fbc = CACHED_PTR(opline.GetResult().GetNum() + b.SizeOf("void *"))
	} else {
		function_name = opline.Const2()
		if ce.GetGetStaticMethod() != nil {
			fbc = ce.GetGetStaticMethod()(ce, function_name.GetStr())
		} else {
			fbc = ZendStdGetStaticMethod(ce, function_name.GetStr(), lang.CondF1(IS_CONST == IS_CONST, func() *types.Zval { return opline.Const2() + 1 }, nil))
		}
		if fbc == nil {
			if EG__().NoException() {
				ZendUndefinedMethod(ce, function_name.GetStr())
			}
			return 0
		}
		if fbc.GetType() <= ZEND_USER_FUNCTION && !fbc.HasFnFlags(types.AccCallViaTrampoline|types.AccNeverCache) {
			CACHE_POLYMORPHIC_PTR(opline.GetResult().GetNum(), ce, fbc)
		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	}

	if !fbc.IsStatic() {
		if executeData.InScope() && operators.InstanceofFunction(executeData.ThisClass(), ce) {
			ce = executeData.ThisClass()
			call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
		} else {
			ZendNonStaticMethodCall(fbc)
			if EG__().HasException() {
				return 0
			}
			goto check_parent_and_self
		}
	} else {
	check_parent_and_self:

		/* previous opcode is ZEND_FETCH_CLASS */

		if (opline.GetOp1().GetNum()&ZEND_FETCH_CLASS_MASK) == ZEND_FETCH_CLASS_PARENT || (opline.GetOp1().GetNum()&ZEND_FETCH_CLASS_MASK) == ZEND_FETCH_CLASS_SELF {
			ce = executeData.ThisClass()
		}
		call_info = ZEND_CALL_NESTED_FUNCTION
	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), ce)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_UNUSED_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var ce *types.ClassEntry
	var call_info uint32
	var fbc types.IFunction
	var call *ZendExecuteData

	/* no function found. try a static method in class */

	{
		ce = ZendFetchClass(nil, opline.GetOp1().GetNum())
		if ce == nil {
			b.Assert(EG__().HasException())
			// ZvalPtrDtorNogc(opline.Op2())
			return 0
		}
	}

	{
		var free_op2 ZendFreeOp
		function_name = opline.Op2()
		{
			if !function_name.IsString() {
				for {
					if function_name.IsRef() {
						function_name = types.Z_REFVAL_P(function_name)
						if function_name.IsString() {
							break
						}
					} else if function_name.IsUndef() {
						ZVAL_UNDEFINED_OP2(executeData)
						if EG__().HasException() {
							return 0
						}
					}
					faults.ThrowError(nil, "Function name must be a string")
					// ZvalPtrDtorNogc(free_op2)
					return 0
				}
			}
		}
		if ce.GetGetStaticMethod() != nil {
			fbc = ce.GetGetStaticMethod()(ce, function_name.GetStr())
		} else {
			fbc = ZendStdGetStaticMethod(ce, function_name.GetStr(), lang.CondF1((IS_TMP_VAR|IS_VAR) == IS_CONST, func() *types.Zval { return opline.Const2() + 1 }, nil))
		}
		if fbc == nil {
			if EG__().NoException() {
				ZendUndefinedMethod(ce, function_name.GetStr())
			}
			// ZvalPtrDtorNogc(free_op2)
			return 0
		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
		{
			// ZvalPtrDtorNogc(free_op2)
		}
	}

	if !fbc.IsStatic() {
		if executeData.InScope() && operators.InstanceofFunction(executeData.ThisClass(), ce) {
			ce = executeData.ThisClass()
			call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
		} else {
			ZendNonStaticMethodCall(fbc)
			if EG__().HasException() {
				return 0
			}
			goto check_parent_and_self
		}
	} else {
	check_parent_and_self:

		/* previous opcode is ZEND_FETCH_CLASS */

		if (opline.GetOp1().GetNum()&ZEND_FETCH_CLASS_MASK) == ZEND_FETCH_CLASS_PARENT || (opline.GetOp1().GetNum()&ZEND_FETCH_CLASS_MASK) == ZEND_FETCH_CLASS_SELF {
			ce = executeData.ThisClass()
		}
		call_info = ZEND_CALL_NESTED_FUNCTION
	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), ce)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var ce *types.ClassEntry
	var call_info uint32
	var fbc types.IFunction
	var call *ZendExecuteData

	/* no function found. try a static method in class */

	{
		ce = ZendFetchClass(nil, opline.GetOp1().GetNum())
		if ce == nil {
			b.Assert(EG__().HasException())
			return 0
		}
	}

	{
		if ce.GetConstructor() == nil {
			faults.ThrowError(nil, "Cannot call constructor")
			return 0
		}
		if executeData.InScope() && executeData.ThisClass() != ce.GetConstructor().GetScope() && ce.GetConstructor().IsPrivate() {
			faults.ThrowError(nil, "Cannot call private %s::__construct()", ce.Name())
			return 0
		}
		fbc = ce.GetConstructor()
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	}
	if !fbc.IsStatic() {
		if executeData.InScope() && operators.InstanceofFunction(executeData.ThisClass(), ce) {
			ce = executeData.ThisClass()
			call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
		} else {
			ZendNonStaticMethodCall(fbc)
			if EG__().HasException() {
				return 0
			}
			goto check_parent_and_self
		}
	} else {
	check_parent_and_self:

		/* previous opcode is ZEND_FETCH_CLASS */
		if (opline.GetOp1().GetNum()&ZEND_FETCH_CLASS_MASK) == ZEND_FETCH_CLASS_PARENT || (opline.GetOp1().GetNum()&ZEND_FETCH_CLASS_MASK) == ZEND_FETCH_CLASS_SELF {
			ce = executeData.ThisClass()
		}
		call_info = ZEND_CALL_NESTED_FUNCTION
	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), ce)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var ce *types.ClassEntry
	var call_info uint32
	var fbc types.IFunction
	var call *ZendExecuteData

	/* no function found. try a static method in class */

	{
		ce = ZendFetchClass(nil, opline.GetOp1().GetNum())
		if ce == nil {
			b.Assert(EG__().HasException())
			return 0
		}
	}

	{
		function_name = opline.Op2()
		{
			if !function_name.IsString() {
				for {
					if function_name.IsRef() {
						function_name = types.Z_REFVAL_P(function_name)
						if function_name.IsString() {
							break
						}
					} else if function_name.IsUndef() {
						ZVAL_UNDEFINED_OP2(executeData)
						if EG__().HasException() {
							return 0
						}
					}
					faults.ThrowError(nil, "Function name must be a string")
					return 0
					break
				}
			}
		}
		if ce.GetGetStaticMethod() != nil {
			fbc = ce.GetGetStaticMethod()(ce, function_name.GetStr())
		} else {
			fbc = ZendStdGetStaticMethod(ce, function_name.GetStr(), lang.CondF1(IS_CV == IS_CONST, func() *types.Zval { return opline.Const2() + 1 }, nil))
		}
		if fbc == nil {
			if EG__().NoException() {
				ZendUndefinedMethod(ce, function_name.GetStr())
			}
			return 0
		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	}

	if !fbc.IsStatic() {
		if executeData.InScope() && operators.InstanceofFunction(executeData.ThisClass(), ce) {
			ce = executeData.ThisClass()
			call_info = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_HAS_THIS
		} else {
			ZendNonStaticMethodCall(fbc)
			if EG__().HasException() {
				return 0
			}
			goto check_parent_and_self
		}
	} else {
	check_parent_and_self:

		/* previous opcode is ZEND_FETCH_CLASS */

		if (opline.GetOp1().GetNum()&ZEND_FETCH_CLASS_MASK) == ZEND_FETCH_CLASS_PARENT || (opline.GetOp1().GetNum()&ZEND_FETCH_CLASS_MASK) == ZEND_FETCH_CLASS_SELF {
			ce = executeData.ThisClass()
		}
		call_info = ZEND_CALL_NESTED_FUNCTION
	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), ce)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
