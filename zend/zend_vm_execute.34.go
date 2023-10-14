package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZEND_NULL_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	faults.ErrorNoreturn(faults.E_ERROR, "Invalid opcode %d/%d/%d.", opline.GetOpcode(), opline.GetOp1Type(), opline.GetOp2Type())
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ExecuteEx(ex *ZendExecuteData) {
	var executeData *ZendExecuteData = ex
	ZEND_VM_LOOP_INTERRUPT_CHECK(executeData)
	for {
		ret := ex.GetOpline().GetHandler()(executeData)
		if ret != 0 {
			if ret > 0 {
				executeData = CurrEX()
				ZEND_VM_LOOP_INTERRUPT_CHECK(executeData)
			} else {
				return
			}
		}
	}
}
func ZendExecute(opArray *types.ZendOpArray, returnValue *types.Zval) {
	var executeData *ZendExecuteData
	var objectOrCalledScope any
	var callInfo uint32
	if EG__().HasException() {
		return
	}
	objectOrCalledScope = ZendGetThisObject(CurrEX())
	if objectOrCalledScope == nil {
		objectOrCalledScope = ZendGetCalledScope(CurrEX())
		callInfo = ZEND_CALL_TOP_CODE | ZEND_CALL_HAS_SYMBOL_TABLE
	} else {
		callInfo = ZEND_CALL_TOP_CODE | ZEND_CALL_HAS_SYMBOL_TABLE | ZEND_CALL_HAS_THIS
	}
	executeData = ZendVmStackPushCallFrame(callInfo, (types.IFunction)(opArray), 0, objectOrCalledScope)
	if CurrEX() != nil {
		executeData.SetSymbolTable(ZendRebuildSymbolTable())
	} else {
		executeData.SetSymbolTable(EG__().GetSymbolTable())
	}
	executeData.SetPrevExecuteData(CurrEX())
	IInitCodeExecuteData(executeData, opArray, returnValue)
	ZendExecuteEx(executeData)
	ZendVmStackFreeCallFrame(executeData)
}
