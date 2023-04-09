package zend

func ZEND_BEGIN_SILENCE_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	opline.Result().SetLong(EG__().GetErrorReporting())
	if EG__().GetErrorReporting() != 0 {
		for {
			EG__().SetErrorReporting(0)
			if EG__().GetErrorReportingIniEntry() == nil {
				var zv *types.Zval = EG__().GetIniDirectives().KeyFind(types.STR_ERROR_REPORTING)
				if zv != nil {
					EG__().SetErrorReportingIniEntry((*ZendIniEntry)(zv.GetPtr()))
				} else {
					break
				}
			}
			if EG__().GetErrorReportingIniEntry().GetModified() == 0 {
				if EG__().GetModifiedIniDirectives() == nil {
					ALLOC_HASHTABLE(EG__().GetModifiedIniDirectives())
					EG__().GetModifiedIniDirectives() = types.MakeArrayEx(8, nil, 0)
				}
				if types.ZendHashAddPtr(EG__().GetModifiedIniDirectives(), types.STR_ERROR_REPORTING, EG__().GetErrorReportingIniEntry()) != nil {
					EG__().GetErrorReportingIniEntry().SetOrigValue(EG__().GetErrorReportingIniEntry().GetValue())
					EG__().GetErrorReportingIniEntry().SetOrigModifiable(EG__().GetErrorReportingIniEntry().GetModifiable())
					EG__().GetErrorReportingIniEntry().SetModified(1)
				}
			}
			break
		}
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
