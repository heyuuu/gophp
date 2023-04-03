package zend

func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		/* Behave like FETCH_OBJ_W */
		return zend_use_tmp_in_write_context_helper_SPEC(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_CONST_CONST_HANDLER(executeData)
	}
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {

		/* Behave like FETCH_OBJ_W */

		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_CONST_TMPVAR_HANDLER(executeData)
	}
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {

		/* Behave like FETCH_OBJ_W */

		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_CONST_CV_HANDLER(executeData)
	}
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_TMP_CONST_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {

		/* Behave like FETCH_OBJ_W */

		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_TMPVAR_CONST_HANDLER(executeData)
	}
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_TMP_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {

		/* Behave like FETCH_OBJ_W */

		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_TMPVAR_TMPVAR_HANDLER(executeData)
	}
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_TMP_CV_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {

		/* Behave like FETCH_OBJ_W */

		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_TMPVAR_CV_HANDLER(executeData)
	}
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {

		/* Behave like FETCH_OBJ_W */

		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_FETCH_OBJ_W_SPEC_VAR_CONST_HANDLER(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_TMPVAR_CONST_HANDLER(executeData)
	}
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_VAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {

		/* Behave like FETCH_OBJ_W */

		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_FETCH_OBJ_W_SPEC_VAR_TMPVAR_HANDLER(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_TMPVAR_TMPVAR_HANDLER(executeData)
	}
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_VAR_CV_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {

		/* Behave like FETCH_OBJ_W */

		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_FETCH_OBJ_W_SPEC_VAR_CV_HANDLER(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_TMPVAR_CV_HANDLER(executeData)
	}
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		/* Behave like FETCH_OBJ_W */

		return ZEND_FETCH_OBJ_W_SPEC_UNUSED_CONST_HANDLER(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_UNUSED_CONST_HANDLER(executeData)
	}
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_UNUSED_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		/* Behave like FETCH_OBJ_W */

		return ZEND_FETCH_OBJ_W_SPEC_UNUSED_TMPVAR_HANDLER(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_UNUSED_TMPVAR_HANDLER(executeData)
	}
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		/* Behave like FETCH_OBJ_W */

		return ZEND_FETCH_OBJ_W_SPEC_UNUSED_CV_HANDLER(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_UNUSED_CV_HANDLER(executeData)
	}
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {

		/* Behave like FETCH_OBJ_W */

		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_FETCH_OBJ_W_SPEC_CV_CONST_HANDLER(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_CV_CONST_HANDLER(executeData)
	}
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {

		/* Behave like FETCH_OBJ_W */

		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_FETCH_OBJ_W_SPEC_CV_TMPVAR_HANDLER(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_CV_TMPVAR_HANDLER(executeData)
	}
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {

		/* Behave like FETCH_OBJ_W */

		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_FETCH_OBJ_W_SPEC_CV_CV_HANDLER(executeData)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_CV_CV_HANDLER(executeData)
	}
}
