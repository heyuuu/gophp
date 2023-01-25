// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

// Source: <ext/standard/assert.c>

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
   | Author: Thies C. Arntzen <thies@thieso.net>                          |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_assert.h"

// # include "php_ini.h"

// # include "zend_exceptions.h"

/* }}} */

var AssertGlobals ZendAssertGlobals
var AssertionErrorCe *zend.ZendClassEntry

func ASSERTG(v __auto__) __auto__ { return AssertGlobals.v }
func SAFE_STRING(s *byte) string {
	if s != nil {
		return s
	} else {
		return ""
	}
}

const (
	ASSERT_ACTIVE = 1
	ASSERT_CALLBACK
	ASSERT_BAIL
	ASSERT_WARNING
	ASSERT_QUIET_EVAL
	ASSERT_EXCEPTION
)

func OnChangeCallback(entry *zend.ZendIniEntry, new_value *zend.ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	if zend.ExecutorGlobals.current_execute_data != nil {
		if zend.Z_TYPE(ASSERTG(callback)) != zend.IS_UNDEF {
			zend.ZvalPtrDtor(&ASSERTG(callback))
			zend.ZVAL_UNDEF(&ASSERTG(callback))
		}
		if new_value != nil && (zend.Z_TYPE(ASSERTG(callback)) != zend.IS_UNDEF || zend.ZSTR_LEN(new_value) != 0) {
			zend.ZVAL_STR_COPY(&ASSERTG(callback), new_value)
		}
	} else {
		if ASSERTG(cb) {
			zend.Pefree(ASSERTG(cb), 1)
		}
		if new_value != nil && zend.ZSTR_LEN(new_value) != 0 {
			ASSERTG(cb) = zend.Pemalloc(zend.ZSTR_LEN(new_value)+1, 1)
			memcpy(ASSERTG(cb), zend.ZSTR_VAL(new_value), zend.ZSTR_LEN(new_value))
			ASSERTG(cb)[zend.ZSTR_LEN(new_value)] = '0'
		} else {
			ASSERTG(cb) = nil
		}
	}
	return zend.SUCCESS
}

/* }}} */

var IniEntries []zend.ZendIniEntryDef = []zend.ZendIniEntryDef{
	{
		"assert.active",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*ZendAssertGlobals)(nil).GetActive())) - (*byte)(nil))),
		any(&AssertGlobals),
		nil,
		"1",
		nil,
		b.SizeOf("\"1\"") - 1,
		b.SizeOf("\"assert.active\"") - 1,
		core.PHP_INI_ALL,
	},
	{
		"assert.bail",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*ZendAssertGlobals)(nil).GetBail())) - (*byte)(nil))),
		any(&AssertGlobals),
		nil,
		"0",
		nil,
		b.SizeOf("\"0\"") - 1,
		b.SizeOf("\"assert.bail\"") - 1,
		core.PHP_INI_ALL,
	},
	{
		"assert.warning",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*ZendAssertGlobals)(nil).GetWarning())) - (*byte)(nil))),
		any(&AssertGlobals),
		nil,
		"1",
		nil,
		b.SizeOf("\"1\"") - 1,
		b.SizeOf("\"assert.warning\"") - 1,
		core.PHP_INI_ALL,
	},
	{"assert.callback", OnChangeCallback, nil, nil, nil, nil, nil, b.SizeOf("NULL") - 1, b.SizeOf("\"assert.callback\"") - 1, core.PHP_INI_ALL},
	{
		"assert.quiet_eval",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*ZendAssertGlobals)(nil).GetQuietEval())) - (*byte)(nil))),
		any(&AssertGlobals),
		nil,
		"0",
		nil,
		b.SizeOf("\"0\"") - 1,
		b.SizeOf("\"assert.quiet_eval\"") - 1,
		core.PHP_INI_ALL,
	},
	{
		"assert.exception",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*ZendAssertGlobals)(nil).GetException())) - (*byte)(nil))),
		any(&AssertGlobals),
		nil,
		"0",
		nil,
		b.SizeOf("\"0\"") - 1,
		b.SizeOf("\"assert.exception\"") - 1,
		core.PHP_INI_ALL,
	},
	{nil, nil, nil, nil, nil, nil, nil, 0, 0, 0},
}

func PhpAssertInitGlobals(assert_globals_p *ZendAssertGlobals) {
	zend.ZVAL_UNDEF(&assert_globals_p.callback)
	assert_globals_p.SetCb(nil)
}

/* }}} */

func ZmStartupAssert(type_ int, module_number int) int {
	var ce zend.ZendClassEntry
	PhpAssertInitGlobals(&AssertGlobals)
	zend.REGISTER_INI_ENTRIES()
	zend.REGISTER_LONG_CONSTANT("ASSERT_ACTIVE", ASSERT_ACTIVE, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("ASSERT_CALLBACK", ASSERT_CALLBACK, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("ASSERT_BAIL", ASSERT_BAIL, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("ASSERT_WARNING", ASSERT_WARNING, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("ASSERT_QUIET_EVAL", ASSERT_QUIET_EVAL, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("ASSERT_EXCEPTION", ASSERT_EXCEPTION, zend.CONST_CS|zend.CONST_PERSISTENT)
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.name = zend.ZendStringInitInterned("AssertionError", b.SizeOf("\"AssertionError\"")-1, 1)
	ce.info.internal.builtin_functions = nil
	AssertionErrorCe = zend.ZendRegisterInternalClassEx(&ce, zend.ZendCeError)
	return zend.SUCCESS
}

/* }}} */

func ZmShutdownAssert(type_ int, module_number int) int {
	if ASSERTG(cb) {
		zend.Pefree(ASSERTG(cb), 1)
		ASSERTG(cb) = nil
	}
	return zend.SUCCESS
}

/* }}} */

func ZmDeactivateAssert(type_ int, module_number int) int {
	if zend.Z_TYPE(ASSERTG(callback)) != zend.IS_UNDEF {
		zend.ZvalPtrDtor(&ASSERTG(callback))
		zend.ZVAL_UNDEF(&ASSERTG(callback))
	}
	return zend.SUCCESS
}

/* }}} */

func ZmInfoAssert(ZEND_MODULE_INFO_FUNC_ARGS) { zend.DISPLAY_INI_ENTRIES() }

/* }}} */

func ZifAssert(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var assertion *zend.Zval
	var description *zend.Zval = nil
	var val int
	var myeval *byte = nil
	var compiled_string_description *byte
	if !(ASSERTG(active)) {
		zend.RETVAL_TRUE
		return
	}
	for {
		var _flags int = 0
		var _min_num_args int = 1
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
			zend.ZendParseArgZvalDeref(_arg, &assertion, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &description, 0)
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
	if zend.Z_TYPE_P(assertion) == zend.IS_STRING {
		var retval zend.Zval
		var old_error_reporting int = 0
		if zend.ZendForbidDynamicCall("assert() with string argument") == zend.FAILURE {
			zend.RETVAL_FALSE
			return
		}
		core.PhpErrorDocref(nil, zend.E_DEPRECATED, "Calling assert() with a string argument is deprecated")
		myeval = zend.Z_STRVAL_P(assertion)
		if ASSERTG(quiet_eval) {
			old_error_reporting = zend.ExecutorGlobals.error_reporting
			zend.ExecutorGlobals.error_reporting = 0
		}
		compiled_string_description = zend.ZendMakeCompiledStringDescription("assert code")
		if zend.ZendEvalStringl(myeval, zend.Z_STRLEN_P(assertion), &retval, compiled_string_description) == zend.FAILURE {
			zend.Efree(compiled_string_description)
			if description == nil {
				zend.ZendThrowError(nil, "Failure evaluating code: %s%s", core.PHP_EOL, myeval)
			} else {
				var str *zend.ZendString = zend.ZvalGetString(description)
				zend.ZendThrowError(nil, "Failure evaluating code: %s%s:\"%s\"", core.PHP_EOL, zend.ZSTR_VAL(str), myeval)
				zend.ZendStringReleaseEx(str, 0)
			}
			if ASSERTG(bail) {
				zend.ZendBailout()
			}
			zend.RETVAL_FALSE
			return
		}
		zend.Efree(compiled_string_description)
		if ASSERTG(quiet_eval) {
			zend.ExecutorGlobals.error_reporting = old_error_reporting
		}
		zend.ConvertToBoolean(&retval)
		val = zend.Z_TYPE(retval) == zend.IS_TRUE
	} else {
		val = zend.ZendIsTrue(assertion)
	}
	if val != 0 {
		zend.RETVAL_TRUE
		return
	}
	if zend.Z_TYPE(ASSERTG(callback)) == zend.IS_UNDEF && ASSERTG(cb) {
		zend.ZVAL_STRING(&ASSERTG(callback), ASSERTG(cb))
	}
	if zend.Z_TYPE(ASSERTG(callback)) != zend.IS_UNDEF {
		var args []zend.Zval
		var retval zend.Zval
		var lineno uint32 = zend.ZendGetExecutedLineno()
		var filename *byte = zend.ZendGetExecutedFilename()
		zend.ZVAL_STRING(&args[0], SAFE_STRING(filename))
		zend.ZVAL_LONG(&args[1], lineno)
		zend.ZVAL_STRING(&args[2], SAFE_STRING(myeval))
		zend.ZVAL_FALSE(&retval)

		/* XXX do we want to check for error here? */

		if description == nil {
			zend.CallUserFunction(nil, nil, &ASSERTG(callback), &retval, 3, args)
			zend.ZvalPtrDtor(&args[2])
			zend.ZvalPtrDtor(&args[0])
		} else {
			zend.ZVAL_STR(&args[3], zend.ZvalGetString(description))
			zend.CallUserFunction(nil, nil, &ASSERTG(callback), &retval, 4, args)
			zend.ZvalPtrDtor(&args[3])
			zend.ZvalPtrDtor(&args[2])
			zend.ZvalPtrDtor(&args[0])
		}
		zend.ZvalPtrDtor(&retval)
	}
	if ASSERTG(exception) {
		if description == nil {
			zend.ZendThrowException(AssertionErrorCe, nil, zend.E_ERROR)
		} else if zend.Z_TYPE_P(description) == zend.IS_OBJECT && zend.InstanceofFunction(zend.Z_OBJCE_P(description), zend.ZendCeThrowable) != 0 {
			zend.Z_ADDREF_P(description)
			zend.ZendThrowExceptionObject(description)
		} else {
			var str *zend.ZendString = zend.ZvalGetString(description)
			zend.ZendThrowException(AssertionErrorCe, zend.ZSTR_VAL(str), zend.E_ERROR)
			zend.ZendStringReleaseEx(str, 0)
		}
	} else if ASSERTG(warning) {
		if description == nil {
			if myeval != nil {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Assertion \"%s\" failed", myeval)
			} else {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Assertion failed")
			}
		} else {
			var str *zend.ZendString = zend.ZvalGetString(description)
			if myeval != nil {
				core.PhpErrorDocref(nil, zend.E_WARNING, "%s: \"%s\" failed", zend.ZSTR_VAL(str), myeval)
			} else {
				core.PhpErrorDocref(nil, zend.E_WARNING, "%s failed", zend.ZSTR_VAL(str))
			}
			zend.ZendStringReleaseEx(str, 0)
		}
	}
	if ASSERTG(bail) {
		zend.ZendBailout()
	}
	zend.RETVAL_FALSE
	return
}

/* }}} */

func ZifAssertOptions(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var value *zend.Zval = nil
	var what zend.ZendLong
	var oldint zend.ZendBool
	var ac int = zend.ZEND_NUM_ARGS()
	var key *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
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
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &what, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &value, 0)
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
	switch what {
	case ASSERT_ACTIVE:
		oldint = ASSERTG(active)
		if ac == 2 {
			var value_str *zend.ZendString = zend.ZvalTryGetString(value)
			if zend.UNEXPECTED(value_str == nil) {
				return
			}
			key = zend.ZendStringInit("assert.active", b.SizeOf("\"assert.active\"")-1, 0)
			zend.ZendAlterIniEntryEx(key, value_str, core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME, 0)
			zend.ZendStringReleaseEx(key, 0)
			zend.ZendStringReleaseEx(value_str, 0)
		}
		zend.RETVAL_LONG(oldint)
		return
		break
	case ASSERT_BAIL:
		oldint = ASSERTG(bail)
		if ac == 2 {
			var value_str *zend.ZendString = zend.ZvalTryGetString(value)
			if zend.UNEXPECTED(value_str == nil) {
				return
			}
			key = zend.ZendStringInit("assert.bail", b.SizeOf("\"assert.bail\"")-1, 0)
			zend.ZendAlterIniEntryEx(key, value_str, core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME, 0)
			zend.ZendStringReleaseEx(key, 0)
			zend.ZendStringReleaseEx(value_str, 0)
		}
		zend.RETVAL_LONG(oldint)
		return
		break
	case ASSERT_QUIET_EVAL:
		oldint = ASSERTG(quiet_eval)
		if ac == 2 {
			var value_str *zend.ZendString = zend.ZvalTryGetString(value)
			if zend.UNEXPECTED(value_str == nil) {
				return
			}
			key = zend.ZendStringInit("assert.quiet_eval", b.SizeOf("\"assert.quiet_eval\"")-1, 0)
			zend.ZendAlterIniEntryEx(key, value_str, core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME, 0)
			zend.ZendStringReleaseEx(key, 0)
			zend.ZendStringReleaseEx(value_str, 0)
		}
		zend.RETVAL_LONG(oldint)
		return
		break
	case ASSERT_WARNING:
		oldint = ASSERTG(warning)
		if ac == 2 {
			var value_str *zend.ZendString = zend.ZvalTryGetString(value)
			if zend.UNEXPECTED(value_str == nil) {
				return
			}
			key = zend.ZendStringInit("assert.warning", b.SizeOf("\"assert.warning\"")-1, 0)
			zend.ZendAlterIniEntryEx(key, value_str, core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME, 0)
			zend.ZendStringReleaseEx(key, 0)
			zend.ZendStringReleaseEx(value_str, 0)
		}
		zend.RETVAL_LONG(oldint)
		return
		break
	case ASSERT_CALLBACK:
		if zend.Z_TYPE(ASSERTG(callback)) != zend.IS_UNDEF {
			zend.ZVAL_COPY(return_value, &ASSERTG(callback))
		} else if ASSERTG(cb) {
			zend.RETVAL_STRING(ASSERTG(cb))
		} else {
			zend.RETVAL_NULL()
		}
		if ac == 2 {
			zend.ZvalPtrDtor(&ASSERTG(callback))
			zend.ZVAL_COPY(&ASSERTG(callback), value)
		}
		return
	case ASSERT_EXCEPTION:
		oldint = ASSERTG(exception)
		if ac == 2 {
			var val *zend.ZendString = zend.ZvalTryGetString(value)
			if zend.UNEXPECTED(val == nil) {
				return
			}
			key = zend.ZendStringInit("assert.exception", b.SizeOf("\"assert.exception\"")-1, 0)
			zend.ZendAlterIniEntryEx(key, val, core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME, 0)
			zend.ZendStringReleaseEx(val, 0)
			zend.ZendStringReleaseEx(key, 0)
		}
		zend.RETVAL_LONG(oldint)
		return
		break
	default:
		core.PhpErrorDocref(nil, zend.E_WARNING, "Unknown value "+zend.ZEND_LONG_FMT, what)
		break
	}
	zend.RETVAL_FALSE
	return
}

/* }}} */
