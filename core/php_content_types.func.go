// <<generate>>

package core

import (
	"sik/zend"
)

func PhpDefaultPostReader() {
	if !(strcmp(SG(request_info).request_method, "POST")) {
		if nil == SG(request_info).post_entry {

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
	return zend.SUCCESS
}
func PhpSetupSapiContentTypes() int {
	SapiRegisterPostEntries(PhpPostEntries)
	return zend.SUCCESS
}
