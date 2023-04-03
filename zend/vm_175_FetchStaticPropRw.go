package zend

func ZEND_FETCH_STATIC_PROP_RW_SPEC_HANDLER(executeData *ZendExecuteData) int {
	return zend_fetch_static_prop_helper_SPEC(BP_VAR_RW, executeData)
}
