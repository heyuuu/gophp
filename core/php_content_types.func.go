package core

import (
	"github.com/heyuuu/gophp/zend/types"
)

func PhpDefaultPostReader() {
	if !(strcmp(SG__().request_info.request_method, "POST")) {
		if nil == SG__().request_info.post_entry {

			/* no post handler registered, so we just swallow the data */

			SapiReadStandardFormData()

			/* no post handler registered, so we just swallow the data */

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
