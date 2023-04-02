package zend

import (
	"github.com/heyuuu/gophp/zend/types"
)

func ZendVmGetOpcodeHandler(opcode types.ZendUchar, op *ZendOp) OpcodeHandlerT {
	return ZendVmGetOpcodeHandlerEx(opcode, op, op+1)
}

func ZendVmGetOpcodeHandlerEx(opcode types.ZendUchar, op *ZendOp, nextOp *ZendOp) OpcodeHandlerT {
	var spec = ZendSpecHandlers[opcode]
	var zendVmDecodeEx = map[uint8]int{
		IS_UNUSED:  _UNUSED_CODE, // 0    : 3
		IS_CONST:   _CONST_CODE,  // 1<<0 : 0
		IS_TMP_VAR: _TMP_CODE,    // 1<<1 : 1
		IS_VAR:     _VAR_CODE,    // 1<<2 : 2
		IS_CV:      _CV_CODE,     // 1<<3 : 4
	}

	var offset = 0
	if (spec & SPEC_RULE_OP1) != 0 {
		offset = offset*5 + zendVmDecodeEx[op.GetOp1Type()]
	}
	if (spec & SPEC_RULE_OP2) != 0 {
		offset = offset*5 + zendVmDecodeEx[op.GetOp2Type()]
	}
	if (spec & SPEC_EXTRA_MASK) != 0 {
		if (spec & SPEC_RULE_RETVAL) != 0 {
			offset = offset*2 + (op.GetResultType() != IS_UNUSED)
		} else if (spec & SPEC_RULE_QUICK_ARG) != 0 {
			offset = offset*2 + (op.GetOp2().GetNum() <= MAX_ARG_FLAG_NUM)
		} else if (spec & SPEC_RULE_OP_DATA) != 0 {
			offset = offset*5 + zendVmDecodeEx[nextOp.GetOp1Type()]
		} else if (spec & SPEC_RULE_ISSET) != 0 {
			offset = offset*2 + (op.GetExtendedValue() & ZEND_ISEMPTY)
		} else if (spec & SPEC_RULE_SMART_BRANCH) != 0 {
			offset = offset * 3
			if nextOp.GetOpcode() == ZEND_JMPZ {
				offset += 1
			} else if nextOp.GetOpcode() == ZEND_JMPNZ {
				offset += 2
			}
		}
	}
	return ZendOpcodeHandlers[(spec&SPEC_START_MASK)+offset]
}
func ZendVmSetOpcodeHandler(op *ZendOp) {
	//var opcode types.ZendUchar = ZendUserOpcodes[op.GetOpcode()]
	var opcode types.ZendUchar = op.GetOpcode()
	if (ZendSpecHandlers[op.GetOpcode()] & SPEC_RULE_COMMUTATIVE) != 0 {
		if op.GetOp1Type() < op.GetOp2Type() {
			ZendSwapOperands(op)
		}
	}
	op.SetHandler(ZendVmGetOpcodeHandler(opcode, op))
}
