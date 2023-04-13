package zend

import "github.com/heyuuu/gophp/zend/types"

func ZEND_BEGIN_SILENCE_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	opline.Result().SetLong(EG__().GetErrorReporting())
	if EG__().GetErrorReporting() != 0 {
		for {
			EG__().SetErrorReporting(0)
			if EG__().GetErrorReportingIniEntry() == nil {
				var iniEntry = EG__().IniDirectives().Get(types.STR_ERROR_REPORTING)
				if iniEntry != nil {
					EG__().SetErrorReportingIniEntry(iniEntry)
				} else {
					break
				}
			}
			if EG__().GetErrorReportingIniEntry().GetModified() == 0 {
				if EG__().ModifiedIniDirectives() == nil {
					EG__().ModifiedIniDirectives()
				}
				if EG__().ModifiedIniDirectives().Add(types.STR_ERROR_REPORTING, EG__().GetErrorReportingIniEntry()) {
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
