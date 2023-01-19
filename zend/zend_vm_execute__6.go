// <<generate>>

package zend

import g "sik/runtime/grammar"

var ZendHandlersTable *HashTable = nil

func ZendVmDtor() {
	if ZendHandlersTable != nil {
		ZendHashDestroy(ZendHandlersTable)
		Free(ZendHandlersTable)
		ZendHandlersTable = nil
	}
}
func InitOpcodeSerialiser() {
	var i int
	var tmp Zval
	ZendHandlersTable = Malloc(g.SizeOf("HashTable"))
	_zendHashInit(ZendHandlersTable, ZendHandlersCount, nil, 1)
	ZendHashRealInit(ZendHandlersTable, 0)
	tmp.SetTypeInfo(4)
	for i = 0; i < ZendHandlersCount; i++ {
		tmp.GetValue().SetLval(i)
		ZendHashIndexAdd(ZendHandlersTable, ZendLong(ZendUintptrT(ZendOpcodeHandlers[i])), &tmp)
	}
}
func ZendSerializeOpcodeHandler(op *ZendOp) {
	var zv *Zval
	if ZendHandlersTable == nil {
		InitOpcodeSerialiser()
	}
	zv = ZendHashIndexFind(ZendHandlersTable, ZendLong(ZendUintptrT(op.GetHandler())))
	assert(zv != nil)
	op.SetHandler(any(zend_uintptr_t(*zv).value.lval))
}
func ZendDeserializeOpcodeHandler(op *ZendOp) {
	op.SetHandler(ZendOpcodeHandlers[ZendUintptrT(op.GetHandler())])
}
func ZendGetOpcodeHandlerFunc(op *ZendOp) any { return op.GetHandler() }
func ZendGetHaltOp() *ZendOp                  { return nil }
func ZendVmKind() int                         { return 1 }
func ZendVmGetOpcodeHandlerEx(spec uint32, op *ZendOp) any {
	var zend_vm_decode []int = []int{3, 0, 1, 3, 2, 3, 3, 3, 4}
	var offset uint32 = 0
	if (spec & 0x10000) != 0 {
		offset = offset*5 + zend_vm_decode[op.GetOp1Type()]
	}
	if (spec & 0x20000) != 0 {
		offset = offset*5 + zend_vm_decode[op.GetOp2Type()]
	}
	if (spec & 0xfffc0000) != 0 {
		if (spec & 0x80000) != 0 {
			offset = offset*2 + (op.GetResultType() != 0)
		} else if (spec & 0x100000) != 0 {
			offset = offset*2 + (op.GetOp2().GetNum() <= 12)
		} else if (spec & 0x40000) != 0 {
			offset = offset*5 + zend_vm_decode[(op+1).GetOp1Type()]
		} else if (spec & 0x1000000) != 0 {
			offset = offset*2 + (op.GetExtendedValue() & 1 << 0)
		} else if (spec & 0x200000) != 0 {
			offset = offset * 3
			if (op + 1).GetOpcode() == 43 {
				offset += 1
			} else if (op + 1).GetOpcode() == 44 {
				offset += 2
			}
		}
	}
	return ZendOpcodeHandlers[(spec&0xffff)+offset]
}
func ZendVmGetOpcodeHandler(opcode ZendUchar, op *ZendOp) any {
	return ZendVmGetOpcodeHandlerEx(ZendSpecHandlers[opcode], op)
}
func ZendVmSetOpcodeHandler(op *ZendOp) {
	var opcode ZendUchar = ZendUserOpcodes[op.GetOpcode()]
	if (ZendSpecHandlers[op.GetOpcode()] & 0x800000) != 0 {
		if op.GetOp1Type() < op.GetOp2Type() {
			ZendSwapOperands(op)
		}
	}
	op.SetHandler(ZendVmGetOpcodeHandlerEx(ZendSpecHandlers[opcode], op))
}
func ZendVmSetOpcodeHandlerEx(op *ZendOp, op1_info uint32, op2_info uint32, res_info uint32) {
	var opcode ZendUchar = ZendUserOpcodes[op.GetOpcode()]
	var spec uint32 = ZendSpecHandlers[opcode]
	switch opcode {
	case 1:
		if res_info == 1<<4 && op1_info == 1<<4 && op2_info == 1<<4 {
			if op.GetOp1Type() == 1<<0 && op.GetOp2Type() == 1<<0 {
				break
			}
			spec = 2312 | 0x10000 | 0x20000 | 0x800000
			if op.GetOp1Type() < op.GetOp2Type() {
				ZendSwapOperands(op)
			}
		} else if op1_info == 1<<4 && op2_info == 1<<4 {
			if op.GetOp1Type() == 1<<0 && op.GetOp2Type() == 1<<0 {
				break
			}
			spec = 2337 | 0x10000 | 0x20000 | 0x800000
			if op.GetOp1Type() < op.GetOp2Type() {
				ZendSwapOperands(op)
			}
		} else if op1_info == 1<<5 && op2_info == 1<<5 {
			if op.GetOp1Type() == 1<<0 && op.GetOp2Type() == 1<<0 {
				break
			}
			spec = 2362 | 0x10000 | 0x20000 | 0x800000
			if op.GetOp1Type() < op.GetOp2Type() {
				ZendSwapOperands(op)
			}
		}
		break
	case 2:
		if res_info == 1<<4 && op1_info == 1<<4 && op2_info == 1<<4 {
			if op.GetOp1Type() == 1<<0 && op.GetOp2Type() == 1<<0 {
				break
			}
			spec = 2387 | 0x10000 | 0x20000
		} else if op1_info == 1<<4 && op2_info == 1<<4 {
			if op.GetOp1Type() == 1<<0 && op.GetOp2Type() == 1<<0 {
				break
			}
			spec = 2412 | 0x10000 | 0x20000
		} else if op1_info == 1<<5 && op2_info == 1<<5 {
			if op.GetOp1Type() == 1<<0 && op.GetOp2Type() == 1<<0 {
				break
			}
			spec = 2437 | 0x10000 | 0x20000
		}
		break
	case 3:
		if op.GetOp1Type() < op.GetOp2Type() {
			ZendSwapOperands(op)
		}
		if res_info == 1<<4 && op1_info == 1<<4 && op2_info == 1<<4 {
			if op.GetOp1Type() == 1<<0 && op.GetOp2Type() == 1<<0 {
				break
			}
			spec = 2462 | 0x10000 | 0x20000 | 0x800000
		} else if op1_info == 1<<4 && op2_info == 1<<4 {
			if op.GetOp1Type() == 1<<0 && op.GetOp2Type() == 1<<0 {
				break
			}
			spec = 2487 | 0x10000 | 0x20000 | 0x800000
		} else if op1_info == 1<<5 && op2_info == 1<<5 {
			if op.GetOp1Type() == 1<<0 && op.GetOp2Type() == 1<<0 {
				break
			}
			spec = 2512 | 0x10000 | 0x20000 | 0x800000
		}
		break
	case 18:
		if op.GetOp1Type() < op.GetOp2Type() {
			ZendSwapOperands(op)
		}
		if op1_info == 1<<4 && op2_info == 1<<4 {
			if op.GetOp1Type() == 1<<0 && op.GetOp2Type() == 1<<0 {
				break
			}
			spec = 2537 | 0x10000 | 0x20000 | 0x200000 | 0x800000
		} else if op1_info == 1<<5 && op2_info == 1<<5 {
			if op.GetOp1Type() == 1<<0 && op.GetOp2Type() == 1<<0 {
				break
			}
			spec = 2612 | 0x10000 | 0x20000 | 0x200000 | 0x800000
		}
		break
	case 19:
		if op.GetOp1Type() < op.GetOp2Type() {
			ZendSwapOperands(op)
		}
		if op1_info == 1<<4 && op2_info == 1<<4 {
			if op.GetOp1Type() == 1<<0 && op.GetOp2Type() == 1<<0 {
				break
			}
			spec = 2687 | 0x10000 | 0x20000 | 0x200000 | 0x800000
		} else if op1_info == 1<<5 && op2_info == 1<<5 {
			if op.GetOp1Type() == 1<<0 && op.GetOp2Type() == 1<<0 {
				break
			}
			spec = 2762 | 0x10000 | 0x20000 | 0x200000 | 0x800000
		}
		break
	case 20:
		if op1_info == 1<<4 && op2_info == 1<<4 {
			if op.GetOp1Type() == 1<<0 && op.GetOp2Type() == 1<<0 {
				break
			}
			spec = 2837 | 0x10000 | 0x20000 | 0x200000
		} else if op1_info == 1<<5 && op2_info == 1<<5 {
			if op.GetOp1Type() == 1<<0 && op.GetOp2Type() == 1<<0 {
				break
			}
			spec = 2912 | 0x10000 | 0x20000 | 0x200000
		}
		break
	case 21:
		if op1_info == 1<<4 && op2_info == 1<<4 {
			if op.GetOp1Type() == 1<<0 && op.GetOp2Type() == 1<<0 {
				break
			}
			spec = 2987 | 0x10000 | 0x20000 | 0x200000
		} else if op1_info == 1<<5 && op2_info == 1<<5 {
			if op.GetOp1Type() == 1<<0 && op.GetOp2Type() == 1<<0 {
				break
			}
			spec = 3062 | 0x10000 | 0x20000 | 0x200000
		}
		break
	case 31:
		if op1_info == 1<<4 {
			spec = 3149 | 0x10000
		} else if op1_info == 1<<5 {
			spec = 3154 | 0x10000
		} else if g.CondF1(op.GetOp1Type() == 1<<0, func() bool { return !((*Zval)((*byte)(op)+int32(op.GetOp1()).constant).GetTypeFlags() != 0) }, !(op1_info&(1<<1|1<<2|1<<3|1<<4|1<<5|1<<6|1<<7|1<<8|1<<9|1<<0) - (1<<1 | 1<<2 | 1<<3 | 1<<4 | 1<<5))) {
			spec = 3159 | 0x10000
		}
		break
	case 34:
		if res_info == 1<<4 && op1_info == 1<<4 {
			spec = 3137 | 0x80000
		} else if op1_info == 1<<4 {
			spec = 3139 | 0x80000
		}
		break
	case 35:
		if res_info == 1<<4 && op1_info == 1<<4 {
			spec = 3141 | 0x80000
		} else if op1_info == 1<<4 {
			spec = 3143 | 0x80000
		}
		break
	case 36:
		if res_info == 1<<4 && op1_info == 1<<4 {
			spec = 3145
		} else if op1_info == 1<<4 {
			spec = 3146
		}
		break
	case 37:
		if res_info == 1<<4 && op1_info == 1<<4 {
			spec = 3147
		} else if op1_info == 1<<4 {
			spec = 3148
		}
		break
	case 42:
		if (*ZendOp)((*byte)(op)+int(op.GetOp1().GetJmpOffset())) > op {
			spec = 2311
		}
		break
	case 65:
		if op.GetOp1Type() == 1<<0 && (*Zval)((*byte)(op)+int32(op.GetOp1()).constant).GetTypeFlags() == 0 {
			spec = 3199
		}
		break
	case 66:
		if op.GetOp2().GetNum() <= 12 && (op1_info&(1<<0|1<<10)) == 0 {
			spec = 3194 | 0x10000
		}
		break
	case 78:
		if op.GetOp2Type() == 1<<3 && (op1_info&(1<<0|(1<<1|1<<2|1<<3|1<<4|1<<5|1<<6|1<<7|1<<8|1<<9)|1<<10)) == 1<<7 {
			spec = 3201 | 0x80000
		}
		break
	case 81:
		if (op2_info & (1<<0 | 1<<1 | 1<<6 | 1<<7 | 1<<8 | 1<<9 | 1<<10)) == 0 {
			if op.GetOp1Type() == 1<<0 && op.GetOp2Type() == 1<<0 {
				break
			}
			spec = 3164 | 0x10000 | 0x20000
		}
		break
	case 116:
		if op.GetOp2().GetNum() <= 12 && op.GetOp1Type() == 1<<0 && (*Zval)((*byte)(op)+int32(op.GetOp1()).constant).GetTypeFlags() == 0 {
			spec = 3200
		}
		break
	case 117:
		if (op1_info & (1<<0 | 1<<10)) == 0 {
			spec = 3189 | 0x10000
		}
		break
	case 9:

	case 10:

	case 11:

	case 15:

	case 16:

	case 17:
		if op.GetOp1Type() < op.GetOp2Type() {
			ZendSwapOperands(op)
		}
		break
	case 150:
		if (ZendSpecHandlers[op.GetOpcode()] & 0x800000) != 0 {
			if op.GetOp1Type() < op.GetOp2Type() {
				ZendSwapOperands(op)
			}
		}
		break
	default:
		break
	}
	op.SetHandler(ZendVmGetOpcodeHandlerEx(spec, op))
}
func ZendVmCallOpcodeHandler(ex *ZendExecuteData) int {
	var ret int
	var execute_data *ZendExecuteData = ex
	ret = opcode_handler_t(execute_data.GetOpline()).handler(execute_data)
	return ret
}
