// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/core/streams"
	"sik/zend"
)

// Source: <ext/standard/link.c>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Author:                                                              |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_filestat.h"

// # include "php_globals.h"

// # include < stdlib . h >

// # include < unistd . h >

// # include < sys / stat . h >

// # include < string . h >

// # include < pwd . h >

// # include < grp . h >

// # include < errno . h >

// # include < ctype . h >

// # include "php_link.h"

// # include "php_string.h"

const VOLUME_NAME_NT = 0x2
const VOLUME_NAME_DOS = 0x0

/* {{{ proto string readlink(string filename)
   Return the target of a symbolic link */

func ZifReadlink(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var link *byte
	var link_len int
	var buff []byte
	var ret ssize_t
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgPath(_arg, &link, &link_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if core.PhpCheckOpenBasedir(link) != 0 {
		zend.RETVAL_FALSE
		return
	}
	ret = zend.PhpSysReadlink(link, buff, core.MAXPATHLEN-1)
	if ret == -1 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s", strerror(errno))
		zend.RETVAL_FALSE
		return
	}

	/* Append NULL to the end of the string */

	buff[ret] = '0'
	zend.RETVAL_STRINGL(buff, ret)
	return
}

/* }}} */

func ZifLinkinfo(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var link *byte
	var dirname *byte
	var link_len int
	var sb zend.ZendStatT
	var ret int
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgPath(_arg, &link, &link_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	dirname = zend.Estrndup(link, link_len)
	PhpDirname(dirname, link_len)
	if core.PhpCheckOpenBasedir(dirname) != 0 {
		zend.Efree(dirname)
		zend.RETVAL_FALSE
		return
	}
	ret = zend.VCWD_LSTAT(link, &sb)
	if ret == -1 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s", strerror(errno))
		zend.Efree(dirname)
		zend.RETVAL_LONG(zend.Z_L(-1))
		return
	}
	zend.Efree(dirname)
	zend.RETVAL_LONG(zend.ZendLong(sb.st_dev))
	return
}

/* }}} */

func ZifSymlink(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var topath *byte
	var frompath *byte
	var topath_len int
	var frompath_len int
	var ret int
	var source_p []byte
	var dest_p []byte
	var dirname []byte
	var len_ int
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgPath(_arg, &topath, &topath_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgPath(_arg, &frompath, &frompath_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if core.ExpandFilepath(frompath, source_p) == nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "No such file or directory")
		zend.RETVAL_FALSE
		return
	}
	memcpy(dirname, source_p, b.SizeOf("source_p"))
	len_ = PhpDirname(dirname, strlen(dirname))
	if core.ExpandFilepathEx(topath, dest_p, dirname, len_) == nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "No such file or directory")
		zend.RETVAL_FALSE
		return
	}
	if streams.PhpStreamLocateUrlWrapper(source_p, nil, core.STREAM_LOCATE_WRAPPERS_ONLY) != nil || streams.PhpStreamLocateUrlWrapper(dest_p, nil, core.STREAM_LOCATE_WRAPPERS_ONLY) != nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to symlink to a URL")
		zend.RETVAL_FALSE
		return
	}
	if core.PhpCheckOpenBasedir(dest_p) != 0 {
		zend.RETVAL_FALSE
		return
	}
	if core.PhpCheckOpenBasedir(source_p) != 0 {
		zend.RETVAL_FALSE
		return
	}

	/* For the source, an expanded path must be used (in ZTS an other thread could have changed the CWD).
	 * For the target the exact string given by the user must be used, relative or not, existing or not.
	 * The target is relative to the link itself, not to the CWD. */

	ret = zend.PhpSysSymlink(topath, source_p)
	if ret == -1 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s", strerror(errno))
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_TRUE
	return
}

/* }}} */

func ZifLink(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var topath *byte
	var frompath *byte
	var topath_len int
	var frompath_len int
	var ret int
	var source_p []byte
	var dest_p []byte
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgPath(_arg, &topath, &topath_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgPath(_arg, &frompath, &frompath_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if core.ExpandFilepath(frompath, source_p) == nil || core.ExpandFilepath(topath, dest_p) == nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "No such file or directory")
		zend.RETVAL_FALSE
		return
	}
	if streams.PhpStreamLocateUrlWrapper(source_p, nil, core.STREAM_LOCATE_WRAPPERS_ONLY) != nil || streams.PhpStreamLocateUrlWrapper(dest_p, nil, core.STREAM_LOCATE_WRAPPERS_ONLY) != nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to link to a URL")
		zend.RETVAL_FALSE
		return
	}
	if core.PhpCheckOpenBasedir(dest_p) != 0 {
		zend.RETVAL_FALSE
		return
	}
	if core.PhpCheckOpenBasedir(source_p) != 0 {
		zend.RETVAL_FALSE
		return
	}
	ret = zend.PhpSysLink(topath, frompath)
	if ret == -1 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s", strerror(errno))
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_TRUE
	return
}

/* }}} */
