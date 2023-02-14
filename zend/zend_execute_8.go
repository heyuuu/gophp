// <<generate>>

package zend

import (
	b "sik/builtin"
)

func ZendVmStackExtendCallFrame(call **ZendExecuteData, passed_args uint32, additional_args uint32) {
	if uint32(EG__().GetVmStackEnd()-EG__().GetVmStackTop()) > additional_args {
		EG__().SetVmStackTop(EG__().GetVmStackTop() + additional_args)
	} else {
		*call = ZendVmStackCopyCallFrame(*call, passed_args, additional_args)
	}
}
func ZendGetRunningGenerator(EXECUTE_DATA_D) *ZendGenerator {
	/* The generator object is stored in EX(return_value) */

	var generator *ZendGenerator = (*ZendGenerator)(EX(return_value))

	/* However control may currently be delegated to another generator.
	 * That's the one we're interested in. */

	return generator

	/* However control may currently be delegated to another generator.
	 * That's the one we're interested in. */
}
func CleanupUnfinishedCalls(execute_data *ZendExecuteData, op_num uint32) {
	if EX(call) {
		var call *ZendExecuteData = EX(call)
		var opline *ZendOp = EX(func_).op_array.opcodes + op_num
		var level int
		var do_exit int
		if opline.GetOpcode() == ZEND_INIT_FCALL || opline.GetOpcode() == ZEND_INIT_FCALL_BY_NAME || opline.GetOpcode() == ZEND_INIT_NS_FCALL_BY_NAME || opline.GetOpcode() == ZEND_INIT_DYNAMIC_CALL || opline.GetOpcode() == ZEND_INIT_USER_CALL || opline.GetOpcode() == ZEND_INIT_METHOD_CALL || opline.GetOpcode() == ZEND_INIT_STATIC_METHOD_CALL || opline.GetOpcode() == ZEND_NEW {
			ZEND_ASSERT(op_num != 0)
			opline--
		}
		for {

			/* If the exception was thrown during a function call there might be
			 * arguments pushed to the stack that have to be dtor'ed. */

			level = 0
			do_exit = 0
			for {
				switch opline.GetOpcode() {
				case ZEND_DO_FCALL:
					fallthrough
				case ZEND_DO_ICALL:
					fallthrough
				case ZEND_DO_UCALL:
					fallthrough
				case ZEND_DO_FCALL_BY_NAME:
					level++
				case ZEND_INIT_FCALL:
					fallthrough
				case ZEND_INIT_FCALL_BY_NAME:
					fallthrough
				case ZEND_INIT_NS_FCALL_BY_NAME:
					fallthrough
				case ZEND_INIT_DYNAMIC_CALL:
					fallthrough
				case ZEND_INIT_USER_CALL:
					fallthrough
				case ZEND_INIT_METHOD_CALL:
					fallthrough
				case ZEND_INIT_STATIC_METHOD_CALL:
					fallthrough
				case ZEND_NEW:
					if level == 0 {
						ZEND_CALL_NUM_ARGS(call) = 0
						do_exit = 1
					}
					level--
				case ZEND_SEND_VAL:
					fallthrough
				case ZEND_SEND_VAL_EX:
					fallthrough
				case ZEND_SEND_VAR:
					fallthrough
				case ZEND_SEND_VAR_EX:
					fallthrough
				case ZEND_SEND_FUNC_ARG:
					fallthrough
				case ZEND_SEND_REF:
					fallthrough
				case ZEND_SEND_VAR_NO_REF:
					fallthrough
				case ZEND_SEND_VAR_NO_REF_EX:
					fallthrough
				case ZEND_SEND_USER:
					if level == 0 {
						ZEND_CALL_NUM_ARGS(call) = opline.GetOp2().GetNum()
						do_exit = 1
					}
				case ZEND_SEND_ARRAY:
					fallthrough
				case ZEND_SEND_UNPACK:
					if level == 0 {
						do_exit = 1
					}
				}
				if do_exit == 0 {
					opline--
				}
				if do_exit != 0 {
					break
				}
			}
			if call.GetPrevExecuteData() != nil {

				/* skip current call region */

				level = 0
				do_exit = 0
				for {
					switch opline.GetOpcode() {
					case ZEND_DO_FCALL:
						fallthrough
					case ZEND_DO_ICALL:
						fallthrough
					case ZEND_DO_UCALL:
						fallthrough
					case ZEND_DO_FCALL_BY_NAME:
						level++
					case ZEND_INIT_FCALL:
						fallthrough
					case ZEND_INIT_FCALL_BY_NAME:
						fallthrough
					case ZEND_INIT_NS_FCALL_BY_NAME:
						fallthrough
					case ZEND_INIT_DYNAMIC_CALL:
						fallthrough
					case ZEND_INIT_USER_CALL:
						fallthrough
					case ZEND_INIT_METHOD_CALL:
						fallthrough
					case ZEND_INIT_STATIC_METHOD_CALL:
						fallthrough
					case ZEND_NEW:
						if level == 0 {
							do_exit = 1
						}
						level--
					}
					opline--
					if do_exit != 0 {
						break
					}
				}
			}
			ZendVmStackFreeArgs(EX(call))
			if (ZEND_CALL_INFO(call) & ZEND_CALL_RELEASE_THIS) != 0 {
				OBJ_RELEASE(call.GetThis().GetObj())
			}
			if call.GetFunc().IsClosure() {
				ZendObjectRelease(ZEND_CLOSURE_OBJECT(call.GetFunc()))
			} else if call.GetFunc().IsCallViaTrampoline() {
				ZendStringReleaseEx(call.GetFunc().GetFunctionName(), 0)
				ZendFreeTrampoline(call.GetFunc())
			}
			EX(call) = call.GetPrevExecuteData()
			ZendVmStackFreeCallFrame(call)
			call = EX(call)
			if call == nil {
				break
			}
		}
	}
}
func FindLiveRange(op_array *ZendOpArray, op_num uint32, var_num uint32) *ZendLiveRange {
	var i int
	for i = 0; i < op_array.GetLastLiveRange(); i++ {
		var range_ *ZendLiveRange = op_array.GetLiveRange()[i]
		if op_num >= range_.GetStart() && op_num < range_.GetEnd() && var_num == (range_.GetVar() & ^ZEND_LIVE_MASK) {
			return range_
		}
	}
	return nil
}
func CleanupLiveVars(execute_data *ZendExecuteData, op_num uint32, catch_op_num uint32) {
	var i int
	for i = 0; i < EX(func_).op_array.last_live_range; i++ {
		var range_ *ZendLiveRange = EX(func_).op_array.live_range[i]
		if range_.GetStart() > op_num {

			/* further blocks will not be relevant... */

			break

			/* further blocks will not be relevant... */

		} else if op_num < range_.GetEnd() {
			if catch_op_num == 0 || catch_op_num >= range_.GetEnd() {
				var kind uint32 = range_.GetVar() & ZEND_LIVE_MASK
				var var_num uint32 = range_.GetVar() & ^ZEND_LIVE_MASK
				var var_ *Zval = EX_VAR(var_num)
				if kind == ZEND_LIVE_TMPVAR {
					ZvalPtrDtorNogc(var_)
				} else if kind == ZEND_LIVE_NEW {
					var obj *ZendObject
					ZEND_ASSERT(var_.IsObject())
					obj = var_.GetObj()
					ZendObjectStoreCtorFailed(obj)
					OBJ_RELEASE(obj)
				} else if kind == ZEND_LIVE_LOOP {
					if var_.GetType() != IS_ARRAY && var_.GetFeIterIdx() != uint32-1 {
						ZendHashIteratorDel(var_.GetFeIterIdx())
					}
					ZvalPtrDtorNogc(var_)
				} else if kind == ZEND_LIVE_ROPE {
					var rope **ZendString = (**ZendString)(var_)
					var last *ZendOp = EX(func_).op_array.opcodes + op_num
					for last.GetOpcode() != ZEND_ROPE_ADD && last.GetOpcode() != ZEND_ROPE_INIT || last.GetResult().GetVar() != var_num {
						ZEND_ASSERT(last >= EX(func_).op_array.opcodes)
						last--
					}
					if last.GetOpcode() == ZEND_ROPE_INIT {
						ZendStringReleaseEx(*rope, 0)
					} else {
						var j int = last.GetExtendedValue()
						for {
							ZendStringReleaseEx(rope[j], 0)
							if !(b.PostDec(&j)) {
								break
							}
						}
					}
				} else if kind == ZEND_LIVE_SILENCE {

					/* restore previous error_reporting value */

					if EG__().GetErrorReporting() == 0 && var_.GetLval() != 0 {
						EG__().SetErrorReporting(var_.GetLval())
					}

					/* restore previous error_reporting value */

				}
			}
		}
	}
}
func ZendCleanupUnfinishedExecution(execute_data *ZendExecuteData, op_num uint32, catch_op_num uint32) {
	CleanupUnfinishedCalls(execute_data, op_num)
	CleanupLiveVars(execute_data, op_num, catch_op_num)
}
func ZendSwapOperands(op *ZendOp) {
	var tmp ZnodeOp
	var tmp_type ZendUchar
	tmp = op.GetOp1()
	tmp_type = op.GetOp1Type()
	op.SetOp1(op.GetOp2())
	op.SetOp1Type(op.GetOp2Type())
	op.SetOp2(tmp)
	op.SetOp2Type(tmp_type)
}
func ZendInitDynamicCallString(function *ZendString, num_args uint32) *ZendExecuteData {
	var fbc *ZendFunction
	var func_ *Zval
	var called_scope *ZendClassEntry
	var lcname *ZendString
	var colon *byte
	if b.Assign(&colon, ZendMemrchr(function.GetVal(), ':', function.GetLen())) != nil && colon > function.GetVal() && (*(colon - 1)) == ':' {
		var mname *ZendString
		var cname_length int = colon - function.GetVal() - 1
		var mname_length int = function.GetLen() - cname_length - (b.SizeOf("\"::\"") - 1)
		lcname = ZendStringInit(function.GetVal(), cname_length, 0)
		called_scope = ZendFetchClassByName(lcname, nil, ZEND_FETCH_CLASS_DEFAULT|ZEND_FETCH_CLASS_EXCEPTION)
		if called_scope == nil {
			ZendStringReleaseEx(lcname, 0)
			return nil
		}
		mname = ZendStringInit(function.GetVal()+(cname_length+b.SizeOf("\"::\"")-1), mname_length, 0)
		if called_scope.GetGetStaticMethod() != nil {
			fbc = called_scope.GetGetStaticMethod()(called_scope, mname)
		} else {
			fbc = ZendStdGetStaticMethod(called_scope, mname, nil)
		}
		if fbc == nil {
			if EG__().GetException() == nil {
				ZendUndefinedMethod(called_scope, mname)
			}
			ZendStringReleaseEx(lcname, 0)
			ZendStringReleaseEx(mname, 0)
			return nil
		}
		ZendStringReleaseEx(lcname, 0)
		ZendStringReleaseEx(mname, 0)
		if !fbc.IsStatic() {
			ZendNonStaticMethodCall(fbc)
			if EG__().GetException() != nil {
				return nil
			}
		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	} else {
		if function.GetVal()[0] == '\\' {
			lcname = ZendStringAlloc(function.GetLen()-1, 0)
			ZendStrTolowerCopy(lcname.GetVal(), function.GetVal()+1, function.GetLen()-1)
		} else {
			lcname = ZendStringTolower(function)
		}
		if b.Assign(&func_, EG__().GetFunctionTable().KeyFind(lcname.GetStr())) == nil {
			ZendThrowError(nil, "Call to undefined function %s()", function.GetVal())
			ZendStringReleaseEx(lcname, 0)
			return nil
		}
		ZendStringReleaseEx(lcname, 0)
		fbc = func_.GetFunc()
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
		called_scope = nil
	}
	return ZendVmStackPushCallFrame(ZEND_CALL_NESTED_FUNCTION|ZEND_CALL_DYNAMIC, fbc, num_args, called_scope)
}
func ZendInitDynamicCallObject(function *Zval, num_args uint32) *ZendExecuteData {
	var fbc *ZendFunction
	var object_or_called_scope any
	var called_scope *ZendClassEntry
	var object *ZendObject
	var call_info uint32 = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_DYNAMIC
	if Z_OBJ_HT(*function).GetGetClosure() != nil && Z_OBJ_HT(*function).GetGetClosure()(function, &called_scope, &fbc, &object) == SUCCESS {
		object_or_called_scope = called_scope
		if fbc.IsClosure() {

			/* Delay closure destruction until its invocation */

			ZEND_CLOSURE_OBJECT(fbc).AddRefcount()
			call_info |= ZEND_CALL_CLOSURE
			if fbc.IsFakeClosure() {
				call_info |= ZEND_CALL_FAKE_CLOSURE
			}
			if object != nil {
				call_info |= ZEND_CALL_HAS_THIS
				object_or_called_scope = object
			}
		} else if object != nil {
			call_info |= ZEND_CALL_RELEASE_THIS | ZEND_CALL_HAS_THIS
			object.AddRefcount()
			object_or_called_scope = object
		}
	} else {
		ZendThrowError(nil, "Function name must be a string")
		return nil
	}
	if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
		InitFuncRunTimeCache(fbc.GetOpArray())
	}
	return ZendVmStackPushCallFrame(call_info, fbc, num_args, object_or_called_scope)
}
func ZendInitDynamicCallArray(function *ZendArray, num_args uint32) *ZendExecuteData {
	var fbc *ZendFunction
	var object_or_called_scope any
	var call_info uint32 = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_DYNAMIC
	if function.GetNNumOfElements() == 2 {
		var obj *Zval
		var method *Zval
		obj = function.IndexFindH(0)
		method = function.IndexFindH(1)
		if obj == nil || method == nil {
			ZendThrowError(nil, "Array callback has to contain indices 0 and 1")
			return nil
		}
		ZVAL_DEREF(obj)
		if obj.GetType() != IS_STRING && obj.GetType() != IS_OBJECT {
			ZendThrowError(nil, "First array member is not a valid class name or object")
			return nil
		}
		ZVAL_DEREF(method)
		if method.GetType() != IS_STRING {
			ZendThrowError(nil, "Second array member is not a valid method")
			return nil
		}
		if obj.IsString() {
			var called_scope *ZendClassEntry = ZendFetchClassByName(obj.GetStr(), nil, ZEND_FETCH_CLASS_DEFAULT|ZEND_FETCH_CLASS_EXCEPTION)
			if called_scope == nil {
				return nil
			}
			if called_scope.GetGetStaticMethod() != nil {
				fbc = called_scope.GetGetStaticMethod()(called_scope, method.GetStr())
			} else {
				fbc = ZendStdGetStaticMethod(called_scope, method.GetStr(), nil)
			}
			if fbc == nil {
				if EG__().GetException() == nil {
					ZendUndefinedMethod(called_scope, method.GetStr())
				}
				return nil
			}
			if !fbc.IsStatic() {
				ZendNonStaticMethodCall(fbc)
				if EG__().GetException() != nil {
					return nil
				}
			}
			object_or_called_scope = called_scope
		} else {
			var object *ZendObject = obj.GetObj()
			fbc = Z_OBJ_HT_P(obj).GetGetMethod()(&object, method.GetStr(), nil)
			if fbc == nil {
				if EG__().GetException() == nil {
					ZendUndefinedMethod(object.GetCe(), method.GetStr())
				}
				return nil
			}
			if fbc.IsStatic() {
				object_or_called_scope = object.GetCe()
			} else {
				call_info |= ZEND_CALL_RELEASE_THIS | ZEND_CALL_HAS_THIS
				object.AddRefcount()
				object_or_called_scope = object
			}
		}
	} else {
		ZendThrowError(nil, "Function name must be a string")
		return nil
	}
	if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
		InitFuncRunTimeCache(fbc.GetOpArray())
	}
	return ZendVmStackPushCallFrame(call_info, fbc, num_args, object_or_called_scope)
}
