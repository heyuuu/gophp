package argparse

import "sik/zend"

/**
 * FAST_ZPP: PHP7之后新增的参数处理方式
 * 涉及对应以下宏
 * - ZEND_PARSE_PARAMETERS_START
 * - ZEND_PARSE_PARAMETERS_END
 * - ZEND_PARSE_OPTION
 * - ZEND_PARSE_*
 * - ...
 */
func ZppCheckNumArgs(_num_args int, _min_num_args int, _max_num_args int, _flags int) int {
	if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
		if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
			if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
				zend.CheckNumArgsException(_min_num_args, _max_num_args)
			} else {
				zend.CheckNumArgsError(_min_num_args, _max_num_args)
			}
		}
		return ZPP_ERROR_FAILURE
	}

	return ZPP_ERROR_OK
}
