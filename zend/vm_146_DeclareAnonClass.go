package zend

func ZEND_DECLARE_ANON_CLASS_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var zv *types.Zval
	var ce *types.ClassEntry
	var opline *ZendOp = executeData.GetOpline()
	ce = CACHED_PTR(opline.GetExtendedValue())
	if ce == nil {
		var rtd_key *types.String = RT_CONSTANT(opline, opline.GetOp1()).GetStr()
		zv = EG__().GetClassTable().KeyFind(rtd_key.GetStr())
		if zv == nil {
			for {
				b.Assert((executeData.GetFunc().op_array.fn_flags & AccPreloaded) != 0)
				if ZendPreloadAutoload != nil && ZendPreloadAutoload(executeData.GetFunc().op_array.filename) == types.SUCCESS {
					zv = EG__().GetClassTable().KeyFind(rtd_key.GetStr())
					if zv != nil {
						break
					}
				}
				faults.ErrorNoreturn(faults.E_ERROR, "Anonymous class wasn't preloaded")
				break
			}
		}
		b.Assert(zv != nil)
		ce = zv.GetCe()
		if !ce.IsLinked() {
			if ZendDoLinkClass(ce, b.CondF1(opline.GetOp2Type() == IS_CONST, func() *types.String { return RT_CONSTANT(opline, opline.GetOp2()).GetStr() }, nil)) == types.FAILURE {
				return 0
			}
		}
		CACHE_PTR(opline.GetExtendedValue(), ce)
	}
	opline.GetResultZval().SetCe(ce)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
