package zend

func ZEND_FETCH_DIM_FUNC_ARG_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	} else {
		return ZEND_FETCH_DIM_R_SPEC_CONST_CONST_HANDLER(executeData)
	}
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	} else {
		return ZEND_FETCH_DIM_R_SPEC_CONST_TMPVAR_HANDLER(executeData)
	}
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	} else {
		{
			return zend_use_undef_in_read_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	}
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	} else {
		return ZEND_FETCH_DIM_R_SPEC_CONST_CV_HANDLER(executeData)
	}
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_TMP_CONST_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	} else {
		return ZEND_FETCH_DIM_R_SPEC_TMPVAR_CONST_HANDLER(executeData)
	}
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_TMP_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	} else {
		return ZEND_FETCH_DIM_R_SPEC_TMPVAR_TMPVAR_HANDLER(executeData)
	}
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_TMP_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	} else {
		{
			return zend_use_undef_in_read_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	}
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_TMP_CV_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	} else {
		return ZEND_FETCH_DIM_R_SPEC_TMPVAR_CV_HANDLER(executeData)
	}
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_FETCH_DIM_W_SPEC_VAR_CONST_HANDLER(executeData)
	} else {
		return ZEND_FETCH_DIM_R_SPEC_TMPVAR_CONST_HANDLER(executeData)
	}
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_VAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_FETCH_DIM_W_SPEC_VAR_TMPVAR_HANDLER(executeData)
	} else {
		return ZEND_FETCH_DIM_R_SPEC_TMPVAR_TMPVAR_HANDLER(executeData)
	}
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_VAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_FETCH_DIM_W_SPEC_VAR_UNUSED_HANDLER(executeData)
	} else {
		{
			return zend_use_undef_in_read_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	}
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_VAR_CV_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_FETCH_DIM_W_SPEC_VAR_CV_HANDLER(executeData)
	} else {
		return ZEND_FETCH_DIM_R_SPEC_TMPVAR_CV_HANDLER(executeData)
	}
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_FETCH_DIM_W_SPEC_CV_CONST_HANDLER(executeData)
	} else {
		return ZEND_FETCH_DIM_R_SPEC_CV_CONST_HANDLER(executeData)
	}
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_FETCH_DIM_W_SPEC_CV_TMPVAR_HANDLER(executeData)
	} else {
		return ZEND_FETCH_DIM_R_SPEC_CV_TMPVAR_HANDLER(executeData)
	}
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_FETCH_DIM_W_SPEC_CV_UNUSED_HANDLER(executeData)
	} else {
		{
			return zend_use_undef_in_read_context_helper_SPEC(executeData)
		}
		return ZEND_NULL_HANDLER(executeData)
	}
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_SEND_ARG_BY_REF) != 0 {
		{
			return zend_use_tmp_in_write_context_helper_SPEC(executeData)
		}
		return ZEND_FETCH_DIM_W_SPEC_CV_CV_HANDLER(executeData)
	} else {
		return ZEND_FETCH_DIM_R_SPEC_CV_CV_HANDLER(executeData)
	}
}
