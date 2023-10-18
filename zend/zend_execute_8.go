package zend

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func ZendVmStackExtendCallFrame(call **ZendExecuteData, passed_args uint32, additional_args uint32) {
	(*call).Extend(passed_args, additional_args)
}
func ZendGetRunningGenerator(executeData *ZendExecuteData) *ZendGenerator {
	/* The generator object is stored in EX(return_value) */

	/* However control may currently be delegated to another generator.
	 * That's the one we're interested in. */
	var generator *ZendGenerator = (*ZendGenerator)(executeData.GetReturnValue())

	return generator
}
func CleanupUnfinishedCalls(executeData *ZendExecuteData, op_num uint32) {
	if executeData.GetCall() {
		var call *ZendExecuteData = executeData.GetCall()
		var opline *types.ZendOp = executeData.GetFunc().GetOpArray().opcodes + op_num
		var level int
		var do_exit int
		if opline.GetOpcode() == ZEND_INIT_FCALL || opline.GetOpcode() == ZEND_INIT_FCALL_BY_NAME || opline.GetOpcode() == ZEND_INIT_NS_FCALL_BY_NAME || opline.GetOpcode() == ZEND_INIT_DYNAMIC_CALL || opline.GetOpcode() == ZEND_INIT_USER_CALL || opline.GetOpcode() == ZEND_INIT_METHOD_CALL || opline.GetOpcode() == ZEND_INIT_STATIC_METHOD_CALL || opline.GetOpcode() == ZEND_NEW {
			b.Assert(op_num != 0)
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
						call.NumArgs() = 0
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
						call.NumArgs() = opline.GetOp2().GetNum()
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
			ZendVmStackFreeArgs(executeData.GetCall())
			if (ZEND_CALL_INFO(call) & ZEND_CALL_RELEASE_THIS) != 0 {
				// OBJ_RELEASE(call.GetThis().Object())
			}
			if call.GetFunc().IsClosure() {
				//ZendObjectRelease(ZEND_CLOSURE_OBJECT(call.GetFunc()))
			} else if call.GetFunc().IsCallViaTrampoline() {
				// types.ZendStringReleaseEx(call.GetFunc().GetFunctionName(), 0)
				ZendFreeTrampoline(call.GetFunc())
			}
			executeData.GetCall() = call.GetPrevExecuteData()
			ZendVmStackFreeCallFrame(call)
			call = executeData.GetCall()
			if call == nil {
				break
			}
		}
	}
}
func FindLiveRange(op_array *types.ZendOpArray, op_num uint32, var_num uint32) *ZendLiveRange {
	var i int
	for i = 0; i < op_array.GetLastLiveRange(); i++ {
		var range_ *ZendLiveRange = op_array.GetLiveRange()[i]
		if op_num >= range_.GetStart() && op_num < range_.GetEnd() && var_num == (range_.GetVar() & ^ZEND_LIVE_MASK) {
			return range_
		}
	}
	return nil
}
func CleanupLiveVars(executeData *ZendExecuteData, op_num uint32, catch_op_num uint32) {
	var i int
	for i = 0; i < executeData.GetFunc().GetOpArray().last_live_range; i++ {
		var range_ *ZendLiveRange = executeData.GetFunc().GetOpArray().live_range[i]
		if range_.GetStart() > op_num {

			/* further blocks will not be relevant... */

			break

			/* further blocks will not be relevant... */

		} else if op_num < range_.GetEnd() {
			if catch_op_num == 0 || catch_op_num >= range_.GetEnd() {
				var kind uint32 = range_.GetVar() & ZEND_LIVE_MASK
				var var_num uint32 = range_.GetVar() & ^ZEND_LIVE_MASK
				var var_ *types.Zval = EX_VAR(executeData, var_num)
				if kind == ZEND_LIVE_TMPVAR {
					// ZvalPtrDtorNogc(var_)
				} else if kind == ZEND_LIVE_NEW {
					var obj *types.Object
					b.Assert(var_.IsObject())
					obj = var_.Object()
					ZendObjectStoreCtorFailed(obj)
					// OBJ_RELEASE(obj)
				} else if kind == ZEND_LIVE_LOOP {
					if !var_.IsArray() && var_.GetFeIterIdx() != uint32-1 {
						EG__().DelArrayIterator(var_.GetFeIterIdx())
					}
					// ZvalPtrDtorNogc(var_)
				} else if kind == ZEND_LIVE_ROPE {
					var rope **types.String = (**types.String)(var_)
					var last *types.ZendOp = executeData.GetFunc().GetOpArray().opcodes + op_num
					for last.GetOpcode() != ZEND_ROPE_ADD && last.GetOpcode() != ZEND_ROPE_INIT || last.GetResult().GetVar() != var_num {
						b.Assert(last >= executeData.GetFunc().GetOpArray().opcodes)
						last--
					}
					if last.GetOpcode() == ZEND_ROPE_INIT {
						// types.ZendStringReleaseEx(*rope, 0)
					} else {
						var j int = last.GetExtendedValue()
						for {
							// types.ZendStringReleaseEx(rope[j], 0)
							if !(lang.PostDec(&j)) {
								break
							}
						}
					}
				} else if kind == ZEND_LIVE_SILENCE {

					/* restore previous error_reporting value */

					if EG__().GetErrorReporting() == 0 && var_.Long() != 0 {
						EG__().SetErrorReporting(var_.Long())
					}

					/* restore previous error_reporting value */

				}
			}
		}
	}
}
func ZendCleanupUnfinishedExecution(executeData *ZendExecuteData, op_num uint32, catch_op_num uint32) {
	CleanupUnfinishedCalls(executeData, op_num)
	CleanupLiveVars(executeData, op_num, catch_op_num)
}
func ZendSwapOperands(op *types.ZendOp) {
	var tmp types.ZnodeOp
	var tmp_type uint8
	tmp = op.GetOp1()
	tmp_type = op.GetOp1Type()
	op.SetOp1(op.GetOp2())
	op.SetOp1Type(op.GetOp2Type())
	op.SetOp2(tmp)
	op.SetOp2Type(tmp_type)
}
func ZendInitDynamicCallString(function *types.String, num_args uint32) *ZendExecuteData {
	var fbc types.IFunction
	var called_scope *types.ClassEntry
	var lcname string
	var colon *byte
	if lang.Assign(&colon, operators.ZendMemrchr(function.GetStr(), ':', function.GetLen())) != nil && colon > function.GetStr() && (*(colon - 1)) == ':' {
		var mname *types.String
		var cname_length int = colon - function.GetStr() - 1
		var mname_length int = function.GetLen() - cname_length - (b.SizeOf(`"::"`) - 1)
		lcname = function.GetStr()[:cname_length]
		called_scope = ZendFetchClassByName(lcname, "", ZEND_FETCH_CLASS_DEFAULT|ZEND_FETCH_CLASS_EXCEPTION)
		if called_scope == nil {
			// types.ZendStringReleaseEx(lcname, 0)
			return nil
		}
		mname = types.NewString(b.CastStr(function.GetStr()+(cname_length+b.SizeOf(`"::"`)-1), mname_length))
		if called_scope.GetGetStaticMethod() != nil {
			fbc = called_scope.GetGetStaticMethod()(called_scope, mname.GetStr())
		} else {
			fbc = ZendStdGetStaticMethod(called_scope, mname.GetStr(), nil)
		}
		if fbc == nil {
			if EG__().NoException() {
				ZendUndefinedMethod(called_scope, mname.GetStr())
			}
			// types.ZendStringReleaseEx(lcname, 0)
			// types.ZendStringReleaseEx(mname, 0)
			return nil
		}
		// types.ZendStringReleaseEx(lcname, 0)
		// types.ZendStringReleaseEx(mname, 0)
		if !fbc.IsStatic() {
			ZendNonStaticMethodCall(fbc)
			if EG__().HasException() {
				return nil
			}
		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	} else {
		if function.GetStr()[0] == '\\' {
			lcname = ascii.StrToLower(function.GetStr()[1:])
		} else {
			lcname = ascii.StrToLower(function.GetStr())
		}

		fbc = EG__().FunctionTable().Get(lcname)
		if fbc == nil {
			faults.ThrowError(nil, fmt.Sprintf("Call to undefined function %s()", function.GetStr()))
			return nil
		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
		called_scope = nil
	}
	return ZendVmStackPushCallFrame(ZEND_CALL_NESTED_FUNCTION|ZEND_CALL_DYNAMIC, fbc, num_args, called_scope)
}
func ZendInitDynamicCallObject(function *types.Zval, num_args uint32) *ZendExecuteData {
	var fbc types.IFunction
	var object_or_called_scope any
	var called_scope *types.ClassEntry
	var object *types.Object
	var call_info uint32 = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_DYNAMIC
	if function.Object().CanGetClosure() && function.Object().GetClosure(function, &called_scope, &fbc, &object) == types.SUCCESS {
		object_or_called_scope = called_scope
		if fbc.IsClosure() {

			/* Delay closure destruction until its invocation */

			//ZEND_CLOSURE_OBJECT(fbc).AddRefcount()
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
			// 			object.AddRefcount()
			object_or_called_scope = object
		}
	} else {
		faults.ThrowError(nil, "Function name must be a string")
		return nil
	}
	if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
		InitFuncRunTimeCache(fbc.GetOpArray())
	}
	return ZendVmStackPushCallFrame(call_info, fbc, num_args, object_or_called_scope)
}
func ZendInitDynamicCallArray(function *types.Array, num_args uint32) *ZendExecuteData {
	var fbc types.IFunction
	var object_or_called_scope any
	var call_info uint32 = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_DYNAMIC
	if function.Len() == 2 {
		var obj *types.Zval
		var method *types.Zval
		obj = function.IndexFind(0)
		method = function.IndexFind(1)
		if obj == nil || method == nil {
			faults.ThrowError(nil, "Array callback has to contain indices 0 and 1")
			return nil
		}
		obj = types.ZVAL_DEREF(obj)
		if !obj.IsString() && !obj.IsObject() {
			faults.ThrowError(nil, "First array member is not a valid class name or object")
			return nil
		}
		method = types.ZVAL_DEREF(method)
		if !method.IsString() {
			faults.ThrowError(nil, "Second array member is not a valid method")
			return nil
		}
		if obj.IsString() {
			var called_scope *types.ClassEntry = ZendFetchClassByName(obj.String(), nil, ZEND_FETCH_CLASS_DEFAULT|ZEND_FETCH_CLASS_EXCEPTION)
			if called_scope == nil {
				return nil
			}
			if called_scope.GetGetStaticMethod() != nil {
				fbc = called_scope.GetGetStaticMethod()(called_scope, method.String())
			} else {
				fbc = ZendStdGetStaticMethod(called_scope, method.String(), nil)
			}
			if fbc == nil {
				if EG__().NoException() {
					ZendUndefinedMethod(called_scope, method.String())
				}
				return nil
			}
			if !fbc.IsStatic() {
				ZendNonStaticMethodCall(fbc)
				if EG__().HasException() {
					return nil
				}
			}
			object_or_called_scope = called_scope
		} else {
			fbc = obj.Object().GetMethod(method.String(), nil)
			if fbc == nil {
				if EG__().NoException() {
					ZendUndefinedMethod(obj.Object().GetCe(), method.String())
				}
				return nil
			}
			if fbc.IsStatic() {
				object_or_called_scope = obj.Object().GetCe()
			} else {
				call_info |= ZEND_CALL_RELEASE_THIS | ZEND_CALL_HAS_THIS
				object_or_called_scope = obj.Object()
			}
		}
	} else {
		faults.ThrowError(nil, "Function name must be a string")
		return nil
	}
	if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
		InitFuncRunTimeCache(fbc.GetOpArray())
	}
	return ZendVmStackPushCallFrame(call_info, fbc, num_args, object_or_called_scope)
}
