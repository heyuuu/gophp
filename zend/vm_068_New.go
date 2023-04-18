package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
)

func ZEND_NEW_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var result *types.Zval
	var constructor types.IFunction
	var ce *types.ClassEntry
	var call *ZendExecuteData
	{
		ce = CACHED_PTR(opline.GetOp2().GetNum())
		if ce == nil {
			ce = ZendFetchClassByName(opline.Const1().String(), (opline.Const1() + 1).GetStr(), ZEND_FETCH_CLASS_DEFAULT|ZEND_FETCH_CLASS_EXCEPTION)
			if ce == nil {
				b.Assert(EG__().GetException() != nil)
				opline.Result().SetUndef()
				return 0
			}
			CACHE_PTR(opline.GetOp2().GetNum(), ce)
		}
	}

	result = opline.Result()
	if ObjectInitEx(result, ce) != types.SUCCESS {
		result.SetUndef()
		return 0
	}
	constructor = types.Z_OBJ_HT_P(result).GetGetConstructor()(result.Object())
	if constructor == nil {
		if EG__().GetException() != nil {
			return 0
		}

		/* If there are no arguments, skip over the DO_FCALL opcode. We check if the next
		 * opcode is DO_FCALL in case EXT instructions are used. */

		if opline.GetExtendedValue() == 0 && (opline+1).GetOpcode() == ZEND_DO_FCALL {
			OPLINE = executeData.GetOpline() + 2
			return 0
		}

		/* Perform a dummy function call */

		call = ZendVmStackPushCallFrame(ZEND_CALL_FUNCTION, (types.IFunction)(&ZendPassFunction), opline.GetExtendedValue(), nil)

		/* Perform a dummy function call */

	} else {
		if constructor.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(constructor.GetOpArray())) {
			InitFuncRunTimeCache(constructor.GetOpArray())
		}

		/* We are not handling overloaded classes right now */

		call = ZendVmStackPushCallFrame(ZEND_CALL_FUNCTION|ZEND_CALL_RELEASE_THIS|ZEND_CALL_HAS_THIS, constructor, opline.GetExtendedValue(), result.Object())
		// 		result.AddRefcount()
	}
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_NEW_SPEC_VAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var result *types.Zval
	var constructor types.IFunction
	var ce *types.ClassEntry
	var call *ZendExecuteData

	{
		ce = opline.Op1().Class()
	}
	result = opline.Result()
	if ObjectInitEx(result, ce) != types.SUCCESS {
		result.SetUndef()
		return 0
	}
	constructor = types.Z_OBJ_HT_P(result).GetGetConstructor()(result.Object())
	if constructor == nil {
		if EG__().GetException() != nil {
			return 0
		}

		/* If there are no arguments, skip over the DO_FCALL opcode. We check if the next
		 * opcode is DO_FCALL in case EXT instructions are used. */

		if opline.GetExtendedValue() == 0 && (opline+1).GetOpcode() == ZEND_DO_FCALL {
			OPLINE = executeData.GetOpline() + 2
			return 0
		}

		/* Perform a dummy function call */

		call = ZendVmStackPushCallFrame(ZEND_CALL_FUNCTION, (types.IFunction)(&ZendPassFunction), opline.GetExtendedValue(), nil)

		/* Perform a dummy function call */

	} else {
		if constructor.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(constructor.GetOpArray())) {
			InitFuncRunTimeCache(constructor.GetOpArray())
		}

		/* We are not handling overloaded classes right now */

		call = ZendVmStackPushCallFrame(ZEND_CALL_FUNCTION|ZEND_CALL_RELEASE_THIS|ZEND_CALL_HAS_THIS, constructor, opline.GetExtendedValue(), result.Object())
		// 		result.AddRefcount()
	}
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_NEW_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var result *types.Zval
	var constructor types.IFunction
	var ce *types.ClassEntry
	var call *ZendExecuteData

	{
		ce = ZendFetchClass(nil, opline.GetOp1().GetNum())
		if ce == nil {
			b.Assert(EG__().GetException() != nil)
			opline.Result().SetUndef()
			return 0
		}
	}

	result = opline.Result()
	if ObjectInitEx(result, ce) != types.SUCCESS {
		result.SetUndef()
		return 0
	}
	constructor = types.Z_OBJ_HT_P(result).GetGetConstructor()(result.Object())
	if constructor == nil {
		if EG__().GetException() != nil {
			return 0
		}

		/* If there are no arguments, skip over the DO_FCALL opcode. We check if the next
		 * opcode is DO_FCALL in case EXT instructions are used. */

		if opline.GetExtendedValue() == 0 && (opline+1).GetOpcode() == ZEND_DO_FCALL {
			OPLINE = executeData.GetOpline() + 2
			return 0
		}

		/* Perform a dummy function call */

		call = ZendVmStackPushCallFrame(ZEND_CALL_FUNCTION, (types.IFunction)(&ZendPassFunction), opline.GetExtendedValue(), nil)

		/* Perform a dummy function call */

	} else {
		if constructor.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(constructor.GetOpArray())) {
			InitFuncRunTimeCache(constructor.GetOpArray())
		}

		/* We are not handling overloaded classes right now */

		call = ZendVmStackPushCallFrame(ZEND_CALL_FUNCTION|ZEND_CALL_RELEASE_THIS|ZEND_CALL_HAS_THIS, constructor, opline.GetExtendedValue(), result.Object())
		// 		result.AddRefcount()
	}
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
