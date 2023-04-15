package zend

func ZEND_CALL_TRAMPOLINE_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var args *types.Array = nil
	var fbc types.IFunction = executeData.GetFunc()
	var ret *types.Zval = executeData.GetReturnValue()
	var call_info uint32 = EX_CALL_INFO() & (ZEND_CALL_NESTED | ZEND_CALL_TOP | ZEND_CALL_RELEASE_THIS)
	var num_args uint32 = executeData.NumArgs()
	var call *ZendExecuteData
	if num_args != 0 {
		var p *types.Zval = executeData.Arg(1)
		var end *types.Zval = p + num_args
		args = types.NewArray(num_args)
		types.ZendHashRealInitPacked(args)
		for {
			fillScope := types.PackedFillStart(args)
			for {
				fillScope.FillSet(p)
				fillScope.FillNext()
				p++
				if p == end {
					break
				}
			}
			fillScope.FillEnd()
			break
		}
	}
	call = executeData
	EG__().SetCurrentExecuteData(executeData.GetPrevExecuteData())
	executeData = CurrEX()
	if fbc.GetOpArray().IsStatic() {
		call.SetFunc(fbc.GetOpArray().GetScope().GetCallstatic())
	} else {
		call.SetFunc(fbc.GetOpArray().GetScope().GetCall())
	}
	b.Assert(ZendVmCalcUsedStack(2, call.GetFunc()) <= size_t((*byte)(EG__().GetVmStackEnd())-(*byte)(call)))
	call.NumArgs() = 2
	call.Arg(1).SetString(fbc.GetFunctionName())
	if args != nil {
		call.Arg(2).SetArray(args)
	} else {
		call.Arg(2).SetEmptyArray()
	}
	ZendFreeTrampoline(fbc)
	fbc = call.GetFunc()
	if fbc.GetType() == ZEND_USER_FUNCTION {
		if !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
		executeData = call
		IInitFuncExecuteData(fbc.GetOpArray(), ret, 0, executeData)
		if ZendExecuteEx == ExecuteEx {
			return 1
		} else {
			executeData = executeData.GetPrevExecuteData()
			ZEND_ADD_CALL_FLAG(call, ZEND_CALL_TOP)
			ZendExecuteEx(call)
		}
	} else {
		var retval types.Zval
		b.Assert(fbc.GetType() == ZEND_INTERNAL_FUNCTION)
		EG__().SetCurrentExecuteData(call)
		if fbc.IsHasTypeHints() && ZendVerifyInternalArgTypes(fbc, call) == 0 {
			ZendVmStackFreeCallFrame(call)
			if ret != nil {
				ret.SetUndef()
			}
			goto call_trampoline_end
		}
		if ret == nil {
			ret = &retval
		}
		ret.SetNull()
		if ZendExecuteInternal == nil {

			/* saves one function call if zend_execute_internal is not used */

			fbc.GetInternalFunction().GetHandler()(call, ret)

			/* saves one function call if zend_execute_internal is not used */

		} else {
			ZendExecuteInternal(call, ret)
		}
		EG__().SetCurrentExecuteData(call.GetPrevExecuteData())
	call_trampoline_end:
		ZendVmStackFreeArgs(call)
		if ret == &retval {
			// ZvalPtrDtor(ret)
		}
	}
	executeData = CurrEX()
	if !(executeData.GetFunc()) || !(ZEND_USER_CODE(executeData.GetFunc().type_)) || (call_info&ZEND_CALL_TOP) != 0 {
		return -1
	}
	if (call_info & ZEND_CALL_RELEASE_THIS) != 0 {
		var object *types.ZendObject = call.GetThis().Object()
		// OBJ_RELEASE(object)
	}
	ZendVmStackFreeCallFrame(call)
	if EG__().GetException() != nil {
		faults.RethrowException(executeData)
		return 2
	}
	ZEND_VM_INC_OPCODE(executeData)
	return 2
}
