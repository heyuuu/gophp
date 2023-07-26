package zend

import (
	b "github.com/heyuuu/gophp/php/lang"
)

func ZEND_FETCH_STATIC_PROP_FUNC_ARG_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var fetch_type int = b.Cond((ZEND_CALL_INFO(executeData.GetCall())&ZEND_CALL_SEND_ARG_BY_REF) != 0, BP_VAR_W, BP_VAR_R)
	return zend_fetch_static_prop_helper_SPEC(fetch_type, executeData)
}
