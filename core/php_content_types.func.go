package core

import (
	"github.com/heyuuu/gophp/php/types"
)

func PhpDefaultPostReader() {
	if SG__().RequestInfo.request_method == "POST" {
		if nil == SG__().RequestInfo.post_entry {
			/* no post handler registered, so we just swallow the data */
			SapiReadStandardFormData()
		}
	}
}
func PhpStartupSapiContentTypes() int {
	SapiRegisterDefaultPostReader(PhpDefaultPostReader)
	SapiRegisterTreatData(PhpDefaultTreatData)
	SapiRegisterInputFilter(PhpDefaultInputFilter, nil)
	return types.SUCCESS
}
func PhpSetupSapiContentTypes() int {
	SapiRegisterPostEntries(PhpPostEntries)
	return types.SUCCESS
}
