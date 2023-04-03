package zend

func ZEND_FETCH_W_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_CONST_UNUSED(BP_VAR_W, executeData)
}
func ZEND_FETCH_W_SPEC_TMPVAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_TMPVAR_UNUSED(BP_VAR_W, executeData)
}
func ZEND_FETCH_W_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_CV_UNUSED(BP_VAR_W, executeData)
}
