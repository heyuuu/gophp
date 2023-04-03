package zend

func ZEND_FETCH_STATIC_PROP_R_SPEC_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_static_prop_helper_SPEC(BP_VAR_R, executeData)
}
