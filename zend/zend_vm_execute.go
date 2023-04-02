package zend

func vmOffsetBySpec(spec int, op *ZendOp) int {
	var nextOp *ZendOp = op + 1

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
			offset = offset * 2
			if op.GetResultType() != IS_UNUSED {
				offset++
			}
		} else if (spec & SPEC_RULE_QUICK_ARG) != 0 {
			offset = offset * 2
			if op.GetOp2().GetNum() <= MAX_ARG_FLAG_NUM {
				offset++
			}
		} else if (spec & SPEC_RULE_OP_DATA) != 0 {
			offset = offset*5 + zendVmDecodeEx[nextOp.GetOp1Type()]
		} else if (spec & SPEC_RULE_ISSET) != 0 {
			offset = offset * 2
			if op.GetExtendedValue()&ZEND_ISEMPTY != 0 {
				offset++
			}
		} else if (spec & SPEC_RULE_SMART_BRANCH) != 0 {
			offset = offset * 3
			if nextOp.GetOpcode() == ZEND_JMPZ {
				offset += 1
			} else if nextOp.GetOpcode() == ZEND_JMPNZ {
				offset += 2
			}
		}
	}

	return offset
}
