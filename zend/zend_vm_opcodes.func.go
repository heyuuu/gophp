// <<generate>>

package zend

func ZEND_VM_OP1_FLAGS(flags __auto__) int { return flags & 0xff }
func ZEND_VM_OP2_FLAGS(flags __auto__) int { return flags >> 8 & 0xff }
func ZendGetOpcodeName(opcode ZendUchar) *byte {
	if opcode > ZEND_VM_LAST_OPCODE {
		return nil
	}
	return ZendVmOpcodesNames[opcode]
}
func ZendGetOpcodeFlags(opcode ZendUchar) uint32 {
	if opcode > ZEND_VM_LAST_OPCODE {
		opcode = ZEND_NOP
	}
	return ZendVmOpcodesFlags[opcode]
}
