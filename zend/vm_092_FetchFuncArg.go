package zend

import (
	b "github.com/heyuuu/gophp/php/lang"
)

func ZEND_FETCH_FUNC_ARG_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var fetch_type int = b.Cond((ZEND_CALL_INFO(executeData.GetCall())&ZEND_CALL_SEND_ARG_BY_REF) != 0, BP_VAR_W, BP_VAR_R)
	return zend_fetch_var_address_helper_SPEC_CONST_UNUSED(fetch_type, executeData)
}
func ZEND_FETCH_FUNC_ARG_SPEC_TMPVAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var fetch_type int = b.Cond((ZEND_CALL_INFO(executeData.GetCall())&ZEND_CALL_SEND_ARG_BY_REF) != 0, BP_VAR_W, BP_VAR_R)
	return zend_fetch_var_address_helper_SPEC_TMPVAR_UNUSED(fetch_type, executeData)
}
func ZEND_FETCH_FUNC_ARG_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var fetch_type int = b.Cond((ZEND_CALL_INFO(executeData.GetCall())&ZEND_CALL_SEND_ARG_BY_REF) != 0, BP_VAR_W, BP_VAR_R)
	return zend_fetch_var_address_helper_SPEC_CV_UNUSED(fetch_type, executeData)
}
