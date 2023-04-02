package zend

import (
	"github.com/heyuuu/gophp/zend/types"
)

func ZendVmGetOpcodeHandler(opcode types.ZendUchar, op *ZendOp) OpcodeHandlerT {
	return vmGetHandler(opcode, op)
}
func ZendVmSetOpcodeHandler(op *ZendOp) {
	//var opcode types.ZendUchar = ZendUserOpcodes[op.GetOpcode()]
	var opcode types.ZendUchar = op.GetOpcode()
	if vmOpcodeIsCommutative(opcode) {
		if op.GetOp1Type() < op.GetOp2Type() {
			ZendSwapOperands(op)
		}
	}
	op.SetHandler(ZendVmGetOpcodeHandler(opcode, op))
}
