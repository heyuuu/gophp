package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_NULL_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	faults.ErrorNoreturn(faults.E_ERROR, "Invalid opcode %d/%d/%d.", OPLINE.opcode, OPLINE.op1_type, OPLINE.op2_type)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ExecuteEx(ex *ZendExecuteData) {
	var executeData *ZendExecuteData = ex
	ZEND_VM_LOOP_INTERRUPT_CHECK(executeData)
	for true {
		var ret int
		if b.Assign(&ret, OpcodeHandlerT(OPLINE.handler)(executeData)) != 0 {
			if ret > 0 {
				executeData = CurrEX()
				ZEND_VM_LOOP_INTERRUPT_CHECK(executeData)
			} else {
				return
			}
		}
	}
	faults.ErrorNoreturn(faults.E_CORE_ERROR, "Arrived at end of main loop which shouldn't happen")
}
func ZendExecute(opArray *types.ZendOpArray, returnValue *types.Zval) {
	var executeData *ZendExecuteData
	var objectOrCalledScope any
	var callInfo uint32
	if EG__().GetException() != nil {
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
	executeData.GetPrevExecuteData() = CurrEX()
	IInitCodeExecuteData(executeData, opArray, returnValue)
	ZendExecuteEx(executeData)
	ZendVmStackFreeCallFrame(executeData)
}
