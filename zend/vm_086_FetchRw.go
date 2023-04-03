package zend

func ZEND_FETCH_RW_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_CONST_UNUSED(BP_VAR_RW, executeData)
}
func ZEND_FETCH_RW_SPEC_TMPVAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_TMPVAR_UNUSED(BP_VAR_RW, executeData)
}
func ZEND_FETCH_RW_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_CV_UNUSED(BP_VAR_RW, executeData)
}
