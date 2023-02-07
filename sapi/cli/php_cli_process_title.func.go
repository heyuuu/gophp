// <<generate>>

package cli

import (
	"sik/core"
	"sik/zend"
)

func ZifCliSetProcessTitle(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var title *byte = nil
	var title_len int
	var rc int
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "s", &title, &title_len) == zend.FAILURE {
		return
	}
	rc = SetPsTitle(title)
	if rc == PS_TITLE_SUCCESS {
		return_value.SetTrue()
		return
	}
	core.PhpErrorDocref(nil, zend.E_WARNING, "cli_set_process_title had an error: %s", PsTitleErrno(rc))
	return_value.SetFalse()
	return
}
func ZifCliGetProcessTitle(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var length int = 0
	var title *byte = nil
	var rc int
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	rc = GetPsTitle(&length, &title)
	if rc != PS_TITLE_SUCCESS {
		core.PhpErrorDocref(nil, zend.E_WARNING, "cli_get_process_title had an error: %s", PsTitleErrno(rc))
		return_value.SetNull()
		return
	}
	zend.ZVAL_STRINGL(return_value, title, length)
	return
}
