package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZEND_DO_FCALL_BY_NAME_SPEC_RETVAL_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var call *ZendExecuteData = executeData.GetCall()
	var fbc types.IFunction = call.GetFunc()
	var ret *types.Zval
	executeData.GetCall() = call.GetPrevExecuteData()
	if fbc.GetType() == ZEND_USER_FUNCTION {
		ret = nil

		call.SetPrevExecuteData(executeData)
		executeData = call
		IInitFuncExecuteData(fbc.GetOpArray(), ret, 0, executeData)
		return 1
	} else {
		var retval types.Zval
		b.Assert(fbc.GetType() == ZEND_INTERNAL_FUNCTION)
		if fbc.IsDeprecated() {
			ZendDeprecatedFunction(fbc)
			if EG__().HasException() {
				UNDEF_RESULT()
				ret = &retval
				ret.SetUndef()
				goto fcall_by_name_end
			}
		}
		call.SetPrevExecuteData(executeData)
		EG__().SetCurrentExecuteData(call)
		if fbc.IsHasTypeHints() && ZendVerifyInternalArgTypes(fbc, call) == 0 {
			UNDEF_RESULT()
			ret = &retval
			ret.SetUndef()
			goto fcall_by_name_end
		}
		ret = &retval
		ret.SetNull()
		fbc.GetInternalFunction().GetHandler()(call, ret)
		EG__().SetCurrentExecuteData(executeData)
	fcall_by_name_end:
		ZendVmStackFreeArgs(call)
		ZendVmStackFreeCallFrame(call)
		// IZvalPtrDtor(ret)
	}
	if EG__().HasException() {
		faults.RethrowException(executeData)
		return 0
	}
	OPLINE = opline + 1
	ZEND_VM_INTERRUPT_CHECK(executeData)
	return 0
}
func ZEND_DO_FCALL_BY_NAME_SPEC_RETVAL_USED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var call *ZendExecuteData = executeData.GetCall()
	var fbc types.IFunction = call.GetFunc()
	var ret *types.Zval
	executeData.GetCall() = call.GetPrevExecuteData()
	if fbc.GetType() == ZEND_USER_FUNCTION {
		ret = nil
		ret = opline.Result()
		call.SetPrevExecuteData(executeData)
		executeData = call
		IInitFuncExecuteData(fbc.GetOpArray(), ret, 0, executeData)
		return 1
	} else {
		var retval types.Zval
		b.Assert(fbc.GetType() == ZEND_INTERNAL_FUNCTION)
		if fbc.IsDeprecated() {
			ZendDeprecatedFunction(fbc)
			if EG__().HasException() {
				UNDEF_RESULT()

				goto fcall_by_name_end
			}
		}
		call.SetPrevExecuteData(executeData)
		EG__().SetCurrentExecuteData(call)
		if fbc.IsHasTypeHints() && ZendVerifyInternalArgTypes(fbc, call) == 0 {
			UNDEF_RESULT()

			goto fcall_by_name_end
		}
		ret = opline.Result()
		ret.SetNull()
		fbc.GetInternalFunction().GetHandler()(call, ret)
		EG__().SetCurrentExecuteData(executeData)
	fcall_by_name_end:
		ZendVmStackFreeArgs(call)
		ZendVmStackFreeCallFrame(call)

	}
	if EG__().HasException() {
		faults.RethrowException(executeData)
		return 0
	}
	OPLINE = opline + 1
	ZEND_VM_INTERRUPT_CHECK(executeData)
	return 0
}
