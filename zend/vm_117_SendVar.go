package zend

func ZEND_SEND_VAR_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	return ZEND_SEND_VAR_SPEC_VAR_INLINE_HANDLER(executeData)
}
func ZEND_SEND_VAR_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	return ZEND_SEND_VAR_SPEC_CV_INLINE_HANDLER(executeData)
}
