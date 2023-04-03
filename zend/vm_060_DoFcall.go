package zend

func ZEND_DO_FCALL_SPEC_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var call *ZendExecuteData = executeData.GetCall()
	var fbc types.IFunction = call.GetFunc()
	var ret *types.Zval
	var retval types.Zval
	executeData.GetCall() = call.GetPrevExecuteData()
	if fbc.HasFnFlags(AccAbstract | AccDeprecated) {
		if fbc.IsAbstract() {
			ZendAbstractMethod(fbc)
		fcall_except:
			UNDEF_RESULT()
			ret = &retval
			ret.SetUndef()
			goto fcall_end
		} else {
			ZendDeprecatedFunction(fbc)
			if EG__().GetException() != nil {
				goto fcall_except
			}
		}
	}
	if fbc.GetType() == ZEND_USER_FUNCTION {
		ret = nil

		call.SetPrevExecuteData(executeData)
		executeData = call
		IInitFuncExecuteData(fbc.GetOpArray(), ret, 1, executeData)
		if ZendExecuteEx == ExecuteEx {
			return 1
		} else {
			executeData = executeData.GetPrevExecuteData()
			ZEND_ADD_CALL_FLAG(call, ZEND_CALL_TOP)
			ZendExecuteEx(call)
		}
	} else if fbc.GetType() < ZEND_USER_FUNCTION {
		call.SetPrevExecuteData(executeData)
		EG__().SetCurrentExecuteData(call)
		if fbc.IsHasTypeHints() && ZendVerifyInternalArgTypes(fbc, call) == 0 {
			goto fcall_except
		}
		ret = &retval
		ret.SetNull()
		if ZendExecuteInternal == nil {

			/* saves one function call if zend_execute_internal is not used */

			fbc.GetInternalFunction().GetHandler()(call, ret)

			/* saves one function call if zend_execute_internal is not used */

		} else {
			ZendExecuteInternal(call, ret)
		}
		EG__().SetCurrentExecuteData(executeData)
	fcall_end:
		ZendVmStackFreeArgs(call)
		IZvalPtrDtor(ret)
	} else {
		ret = &retval
		call.SetPrevExecuteData(executeData)
		if ZendDoFcallOverloaded(call, ret) == 0 {
			UNDEF_RESULT()
			return 0
		}
		ZvalPtrDtor(ret)
	}
	if (ZEND_CALL_INFO(call) & ZEND_CALL_RELEASE_THIS) != 0 {
		OBJ_RELEASE(call.GetThis().GetObj())
	}
	ZendVmStackFreeCallFrame(call)
	if EG__().GetException() != nil {
		faults.RethrowException(executeData)
		return 0
	}
	OPLINE = opline + 1
	ZEND_VM_INTERRUPT_CHECK(executeData)
	return 0
}
func ZEND_DO_FCALL_SPEC_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var call *ZendExecuteData = executeData.GetCall()
	var fbc types.IFunction = call.GetFunc()
	var ret *types.Zval
	var retval types.Zval
	executeData.GetCall() = call.GetPrevExecuteData()
	if fbc.HasFnFlags(AccAbstract | AccDeprecated) {
		if fbc.IsAbstract() {
			ZendAbstractMethod(fbc)
		fcall_except:
			UNDEF_RESULT()

			goto fcall_end
		} else {
			ZendDeprecatedFunction(fbc)
			if EG__().GetException() != nil {
				goto fcall_except
			}
		}
	}
	if fbc.GetType() == ZEND_USER_FUNCTION {
		ret = nil
		ret = opline.GetResultZval()
		call.SetPrevExecuteData(executeData)
		executeData = call
		IInitFuncExecuteData(fbc.GetOpArray(), ret, 1, executeData)
		if ZendExecuteEx == ExecuteEx {
			return 1
		} else {
			executeData = executeData.GetPrevExecuteData()
			ZEND_ADD_CALL_FLAG(call, ZEND_CALL_TOP)
			ZendExecuteEx(call)
		}
	} else if fbc.GetType() < ZEND_USER_FUNCTION {
		call.SetPrevExecuteData(executeData)
		EG__().SetCurrentExecuteData(call)
		if fbc.IsHasTypeHints() && ZendVerifyInternalArgTypes(fbc, call) == 0 {
			goto fcall_except
		}
		ret = opline.GetResultZval()
		ret.SetNull()
		if ZendExecuteInternal == nil {

			/* saves one function call if zend_execute_internal is not used */

			fbc.GetInternalFunction().GetHandler()(call, ret)

			/* saves one function call if zend_execute_internal is not used */

		} else {
			ZendExecuteInternal(call, ret)
		}
		EG__().SetCurrentExecuteData(executeData)
	fcall_end:
		ZendVmStackFreeArgs(call)

	} else {
		ret = opline.GetResultZval()
		call.SetPrevExecuteData(executeData)
		if ZendDoFcallOverloaded(call, ret) == 0 {
			UNDEF_RESULT()
			return 0
		}

	}
	if (ZEND_CALL_INFO(call) & ZEND_CALL_RELEASE_THIS) != 0 {
		OBJ_RELEASE(call.GetThis().GetObj())
	}
	ZendVmStackFreeCallFrame(call)
	if EG__().GetException() != nil {
		faults.RethrowException(executeData)
		return 0
	}
	OPLINE = opline + 1
	ZEND_VM_INTERRUPT_CHECK(executeData)
	return 0
}
