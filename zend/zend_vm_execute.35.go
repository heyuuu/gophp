// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/faults"
	"sik/zend/types"
	"sik/zend/zpp"
)

func ZendVmDtor() {
	if ZendHandlersTable != nil {
		ZendHandlersTable.Destroy()
		Free(ZendHandlersTable)
		ZendHandlersTable = nil
	}
}
func InitOpcodeSerialiser() {
	var i int
	var tmp types.Zval
	ZendHandlersTable = Malloc(b.SizeOf("HashTable"))
	ZendHandlersTable = types.MakeArrayEx(ZendHandlersCount, nil, 1)
	types.ZendHashRealInit(ZendHandlersTable, 0)
	tmp.SetTypeInfo(types.IS_LONG)
	for i = 0; i < ZendHandlersCount; i++ {
		tmp.SetLval(i)
		ZendHandlersTable.IndexAddH(ZendLong(types.ZendUintptrT(ZendOpcodeHandlers[i])), &tmp)
	}
}
func ZendSerializeOpcodeHandler(op *ZendOp) {
	var zv *types.Zval
	if ZendHandlersTable == nil {
		InitOpcodeSerialiser()
	}
	zv = ZendHandlersTable.IndexFindH(ZendLong(types.ZendUintptrT(op.GetHandler())))
	b.Assert(zv != nil)
	op.SetHandler(any(types.ZendUintptrT(zv.GetLval())))
}
func ZendDeserializeOpcodeHandler(op *ZendOp) {
	op.SetHandler(ZendOpcodeHandlers[types.ZendUintptrT(op.GetHandler())])
}
func ZendGetOpcodeHandlerFunc(op *ZendOp) any { return op.GetHandler() }
func ZendGetHaltOp() *ZendOp                  { return nil }
func ZendVmKind() int                         { return ZEND_VM_KIND }
func ZendVmGetOpcodeHandlerEx(spec uint32, op *ZendOp) any {
	var zend_vm_decode []int = []int{_UNUSED_CODE, _CONST_CODE, _TMP_CODE, _UNUSED_CODE, _VAR_CODE, _UNUSED_CODE, _UNUSED_CODE, _UNUSED_CODE, _CV_CODE}
	var offset uint32 = 0
	if (spec & SPEC_RULE_OP1) != 0 {
		offset = offset*5 + zend_vm_decode[op.GetOp1Type()]
	}
	if (spec & SPEC_RULE_OP2) != 0 {
		offset = offset*5 + zend_vm_decode[op.GetOp2Type()]
	}
	if (spec & SPEC_EXTRA_MASK) != 0 {
		if (spec & SPEC_RULE_RETVAL) != 0 {
			offset = offset*2 + (op.GetResultType() != IS_UNUSED)
		} else if (spec & SPEC_RULE_QUICK_ARG) != 0 {
			offset = offset*2 + (op.GetOp2().GetNum() <= MAX_ARG_FLAG_NUM)
		} else if (spec & SPEC_RULE_OP_DATA) != 0 {
			offset = offset*5 + zend_vm_decode[(op+1).GetOp1Type()]
		} else if (spec & SPEC_RULE_ISSET) != 0 {
			offset = offset*2 + (op.GetExtendedValue() & ZEND_ISEMPTY)
		} else if (spec & SPEC_RULE_SMART_BRANCH) != 0 {
			offset = offset * 3
			if (op + 1).GetOpcode() == ZEND_JMPZ {
				offset += 1
			} else if (op + 1).GetOpcode() == ZEND_JMPNZ {
				offset += 2
			}
		}
	}
	return ZendOpcodeHandlers[(spec&SPEC_START_MASK)+offset]
}
func ZendVmGetOpcodeHandler(opcode types.ZendUchar, op *ZendOp) any {
	return ZendVmGetOpcodeHandlerEx(ZendSpecHandlers[opcode], op)
}
func ZendVmSetOpcodeHandler(op *ZendOp) {
	var opcode types.ZendUchar = ZendUserOpcodes[op.GetOpcode()]
	if (ZendSpecHandlers[op.GetOpcode()] & SPEC_RULE_COMMUTATIVE) != 0 {
		if op.GetOp1Type() < op.GetOp2Type() {
			ZendSwapOperands(op)
		}
	}
	op.SetHandler(ZendVmGetOpcodeHandlerEx(ZendSpecHandlers[opcode], op))
}
func ZendVmSetOpcodeHandlerEx(op *ZendOp, op1_info uint32, op2_info uint32, res_info uint32) {
	var opcode types.ZendUchar = ZendUserOpcodes[op.GetOpcode()]
	var spec uint32 = ZendSpecHandlers[opcode]
	switch opcode {
	case ZEND_ADD:
		if res_info == MAY_BE_LONG && op1_info == MAY_BE_LONG && op2_info == MAY_BE_LONG {
			if op.GetOp1Type() == IS_CONST && op.GetOp2Type() == IS_CONST {
				break
			}
			spec = 2312 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_COMMUTATIVE
			if op.GetOp1Type() < op.GetOp2Type() {
				ZendSwapOperands(op)
			}
		} else if op1_info == MAY_BE_LONG && op2_info == MAY_BE_LONG {
			if op.GetOp1Type() == IS_CONST && op.GetOp2Type() == IS_CONST {
				break
			}
			spec = 2337 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_COMMUTATIVE
			if op.GetOp1Type() < op.GetOp2Type() {
				ZendSwapOperands(op)
			}
		} else if op1_info == MAY_BE_DOUBLE && op2_info == MAY_BE_DOUBLE {
			if op.GetOp1Type() == IS_CONST && op.GetOp2Type() == IS_CONST {
				break
			}
			spec = 2362 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_COMMUTATIVE
			if op.GetOp1Type() < op.GetOp2Type() {
				ZendSwapOperands(op)
			}
		}
	case ZEND_SUB:
		if res_info == MAY_BE_LONG && op1_info == MAY_BE_LONG && op2_info == MAY_BE_LONG {
			if op.GetOp1Type() == IS_CONST && op.GetOp2Type() == IS_CONST {
				break
			}
			spec = 2387 | SPEC_RULE_OP1 | SPEC_RULE_OP2
		} else if op1_info == MAY_BE_LONG && op2_info == MAY_BE_LONG {
			if op.GetOp1Type() == IS_CONST && op.GetOp2Type() == IS_CONST {
				break
			}
			spec = 2412 | SPEC_RULE_OP1 | SPEC_RULE_OP2
		} else if op1_info == MAY_BE_DOUBLE && op2_info == MAY_BE_DOUBLE {
			if op.GetOp1Type() == IS_CONST && op.GetOp2Type() == IS_CONST {
				break
			}
			spec = 2437 | SPEC_RULE_OP1 | SPEC_RULE_OP2
		}
	case ZEND_MUL:
		if op.GetOp1Type() < op.GetOp2Type() {
			ZendSwapOperands(op)
		}
		if res_info == MAY_BE_LONG && op1_info == MAY_BE_LONG && op2_info == MAY_BE_LONG {
			if op.GetOp1Type() == IS_CONST && op.GetOp2Type() == IS_CONST {
				break
			}
			spec = 2462 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_COMMUTATIVE
		} else if op1_info == MAY_BE_LONG && op2_info == MAY_BE_LONG {
			if op.GetOp1Type() == IS_CONST && op.GetOp2Type() == IS_CONST {
				break
			}
			spec = 2487 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_COMMUTATIVE
		} else if op1_info == MAY_BE_DOUBLE && op2_info == MAY_BE_DOUBLE {
			if op.GetOp1Type() == IS_CONST && op.GetOp2Type() == IS_CONST {
				break
			}
			spec = 2512 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_COMMUTATIVE
		}
	case ZEND_IS_EQUAL:
		if op.GetOp1Type() < op.GetOp2Type() {
			ZendSwapOperands(op)
		}
		if op1_info == MAY_BE_LONG && op2_info == MAY_BE_LONG {
			if op.GetOp1Type() == IS_CONST && op.GetOp2Type() == IS_CONST {
				break
			}
			spec = 2537 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_SMART_BRANCH | SPEC_RULE_COMMUTATIVE
		} else if op1_info == MAY_BE_DOUBLE && op2_info == MAY_BE_DOUBLE {
			if op.GetOp1Type() == IS_CONST && op.GetOp2Type() == IS_CONST {
				break
			}
			spec = 2612 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_SMART_BRANCH | SPEC_RULE_COMMUTATIVE
		}
	case ZEND_IS_NOT_EQUAL:
		if op.GetOp1Type() < op.GetOp2Type() {
			ZendSwapOperands(op)
		}
		if op1_info == MAY_BE_LONG && op2_info == MAY_BE_LONG {
			if op.GetOp1Type() == IS_CONST && op.GetOp2Type() == IS_CONST {
				break
			}
			spec = 2687 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_SMART_BRANCH | SPEC_RULE_COMMUTATIVE
		} else if op1_info == MAY_BE_DOUBLE && op2_info == MAY_BE_DOUBLE {
			if op.GetOp1Type() == IS_CONST && op.GetOp2Type() == IS_CONST {
				break
			}
			spec = 2762 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_SMART_BRANCH | SPEC_RULE_COMMUTATIVE
		}
	case ZEND_IS_SMALLER:
		if op1_info == MAY_BE_LONG && op2_info == MAY_BE_LONG {
			if op.GetOp1Type() == IS_CONST && op.GetOp2Type() == IS_CONST {
				break
			}
			spec = 2837 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_SMART_BRANCH
		} else if op1_info == MAY_BE_DOUBLE && op2_info == MAY_BE_DOUBLE {
			if op.GetOp1Type() == IS_CONST && op.GetOp2Type() == IS_CONST {
				break
			}
			spec = 2912 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_SMART_BRANCH
		}
	case ZEND_IS_SMALLER_OR_EQUAL:
		if op1_info == MAY_BE_LONG && op2_info == MAY_BE_LONG {
			if op.GetOp1Type() == IS_CONST && op.GetOp2Type() == IS_CONST {
				break
			}
			spec = 2987 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_SMART_BRANCH
		} else if op1_info == MAY_BE_DOUBLE && op2_info == MAY_BE_DOUBLE {
			if op.GetOp1Type() == IS_CONST && op.GetOp2Type() == IS_CONST {
				break
			}
			spec = 3062 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_SMART_BRANCH
		}
	case ZEND_QM_ASSIGN:
		if op1_info == MAY_BE_LONG {
			spec = 3149 | SPEC_RULE_OP1
		} else if op1_info == MAY_BE_DOUBLE {
			spec = 3154 | SPEC_RULE_OP1
		} else if b.CondF1(op.GetOp1Type() == IS_CONST, func() bool { return !(RT_CONSTANT(op, op.GetOp1()).IsRefcounted()) }, !(op1_info&(MAY_BE_ANY|MAY_BE_UNDEF) - (MAY_BE_NULL | MAY_BE_FALSE | MAY_BE_TRUE | MAY_BE_LONG | MAY_BE_DOUBLE))) {
			spec = 3159 | SPEC_RULE_OP1
		}
	case ZEND_PRE_INC:
		if res_info == MAY_BE_LONG && op1_info == MAY_BE_LONG {
			spec = 3137 | SPEC_RULE_RETVAL
		} else if op1_info == MAY_BE_LONG {
			spec = 3139 | SPEC_RULE_RETVAL
		}
	case ZEND_PRE_DEC:
		if res_info == MAY_BE_LONG && op1_info == MAY_BE_LONG {
			spec = 3141 | SPEC_RULE_RETVAL
		} else if op1_info == MAY_BE_LONG {
			spec = 3143 | SPEC_RULE_RETVAL
		}
	case ZEND_POST_INC:
		if res_info == MAY_BE_LONG && op1_info == MAY_BE_LONG {
			spec = 3145
		} else if op1_info == MAY_BE_LONG {
			spec = 3146
		}
	case ZEND_POST_DEC:
		if res_info == MAY_BE_LONG && op1_info == MAY_BE_LONG {
			spec = 3147
		} else if op1_info == MAY_BE_LONG {
			spec = 3148
		}
	case ZEND_JMP:
		if OP_JMP_ADDR(op, op.GetOp1()) > op {
			spec = 2311
		}
	case ZEND_SEND_VAL:
		if op.GetOp1Type() == IS_CONST && !(RT_CONSTANT(op, op.GetOp1()).IsRefcounted()) {
			spec = 3199
		}
	case ZEND_SEND_VAR_EX:
		if op.GetOp2().GetNum() <= MAX_ARG_FLAG_NUM && (op1_info&(MAY_BE_UNDEF|MAY_BE_REF)) == 0 {
			spec = 3194 | SPEC_RULE_OP1
		}
	case ZEND_FE_FETCH_R:
		if op.GetOp2Type() == IS_CV && (op1_info&(MAY_BE_UNDEF|MAY_BE_ANY|MAY_BE_REF)) == MAY_BE_ARRAY {
			spec = 3201 | SPEC_RULE_RETVAL
		}
	case ZEND_FETCH_DIM_R:
		if (op2_info & (MAY_BE_UNDEF | MAY_BE_NULL | MAY_BE_STRING | MAY_BE_ARRAY | MAY_BE_OBJECT | MAY_BE_RESOURCE | MAY_BE_REF)) == 0 {
			if op.GetOp1Type() == IS_CONST && op.GetOp2Type() == IS_CONST {
				break
			}
			spec = 3164 | SPEC_RULE_OP1 | SPEC_RULE_OP2
		}
	case ZEND_SEND_VAL_EX:
		if op.GetOp2().GetNum() <= MAX_ARG_FLAG_NUM && op.GetOp1Type() == IS_CONST && !(RT_CONSTANT(op, op.GetOp1()).IsRefcounted()) {
			spec = 3200
		}
	case ZEND_SEND_VAR:
		if (op1_info & (MAY_BE_UNDEF | MAY_BE_REF)) == 0 {
			spec = 3189 | SPEC_RULE_OP1
		}
	case ZEND_BW_OR:
		fallthrough
	case ZEND_BW_AND:
		fallthrough
	case ZEND_BW_XOR:
		fallthrough
	case ZEND_BOOL_XOR:
		fallthrough
	case ZEND_IS_IDENTICAL:
		fallthrough
	case ZEND_IS_NOT_IDENTICAL:
		if op.GetOp1Type() < op.GetOp2Type() {
			ZendSwapOperands(op)
		}
	case ZEND_USER_OPCODE:
		if (ZendSpecHandlers[op.GetOpcode()] & SPEC_RULE_COMMUTATIVE) != 0 {
			if op.GetOp1Type() < op.GetOp2Type() {
				ZendSwapOperands(op)
			}
		}
	default:

	}
	op.SetHandler(ZendVmGetOpcodeHandlerEx(spec, op))
}
func ZendVmCallOpcodeHandler(ex *ZendExecuteData) int {
	var ret int
	var executeData *ZendExecuteData = ex
	ret = OpcodeHandlerT(OPLINE.handler)(executeData)
	return ret
}
