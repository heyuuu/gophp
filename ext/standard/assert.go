// <<generate>>

package standard

import (
	"sik/core"
	r "sik/runtime"
	g "sik/runtime/grammar"
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

// @type ZendAssertGlobals struct

var AssertGlobals ZendAssertGlobals
var AssertionErrorCe *zend.ZendClassEntry

// #define ASSERTG(v) ZEND_MODULE_GLOBALS_ACCESSOR ( assert , v )

// #define SAFE_STRING(s) ( ( s ) ? ( s ) : "" )

const (
	ASSERT_ACTIVE = 1
	ASSERT_CALLBACK
	ASSERT_BAIL
	ASSERT_WARNING
	ASSERT_QUIET_EVAL
	ASSERT_EXCEPTION
)

func OnChangeCallback(entry *zend.ZendIniEntry, new_value *zend.ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	if zend.EG.current_execute_data != nil {
		if AssertGlobals.callback.u1.v.type_ != 0 {
			zend.ZvalPtrDtor(&(AssertGlobals.GetCallback()))
			&(AssertGlobals.GetCallback()).u1.type_info = 0
		}
		if new_value != nil && (AssertGlobals.callback.u1.v.type_ != 0 || new_value.len_ != 0) {
			var __z *zend.Zval = &(AssertGlobals.GetCallback())
			var __s *zend.ZendString = new_value
			__z.value.str = __s
			if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
				__z.u1.type_info = 6
			} else {
				zend.ZendGcAddref(&__s.gc)
				__z.u1.type_info = 6 | 1<<0<<8
			}
		}
	} else {
		if AssertGlobals.GetCb() != nil {
			g.CondF(true, func() { return zend.Free(AssertGlobals.GetCb()) }, func() { return zend._efree(AssertGlobals.GetCb()) })
		}
		if new_value != nil && new_value.len_ != 0 {
			AssertGlobals.SetCb(zend.__zendMalloc(new_value.len_ + 1))
			memcpy(AssertGlobals.GetCb(), new_value.val, new_value.len_)
			AssertGlobals.GetCb()[new_value.len_] = '0'
		} else {
			AssertGlobals.SetCb(nil)
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
		g.SizeOf("\"1\"") - 1,
		g.SizeOf("\"assert.active\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"assert.bail",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*ZendAssertGlobals)(nil).GetBail())) - (*byte)(nil))),
		any(&AssertGlobals),
		nil,
		"0",
		nil,
		g.SizeOf("\"0\"") - 1,
		g.SizeOf("\"assert.bail\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"assert.warning",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*ZendAssertGlobals)(nil).GetWarning())) - (*byte)(nil))),
		any(&AssertGlobals),
		nil,
		"1",
		nil,
		g.SizeOf("\"1\"") - 1,
		g.SizeOf("\"assert.warning\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{"assert.callback", OnChangeCallback, nil, nil, nil, nil, nil, g.SizeOf("NULL") - 1, g.SizeOf("\"assert.callback\"") - 1, 1<<0 | 1<<1 | 1<<2},
	{
		"assert.quiet_eval",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*ZendAssertGlobals)(nil).GetQuietEval())) - (*byte)(nil))),
		any(&AssertGlobals),
		nil,
		"0",
		nil,
		g.SizeOf("\"0\"") - 1,
		g.SizeOf("\"assert.quiet_eval\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{
		"assert.exception",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*ZendAssertGlobals)(nil).GetException())) - (*byte)(nil))),
		any(&AssertGlobals),
		nil,
		"0",
		nil,
		g.SizeOf("\"0\"") - 1,
		g.SizeOf("\"assert.exception\"") - 1,
		1<<0 | 1<<1 | 1<<2,
	},
	{nil, nil, nil, nil, nil, nil, nil, 0, 0, 0},
}

func PhpAssertInitGlobals(assert_globals_p *ZendAssertGlobals) {
	&assert_globals_p.callback.u1.type_info = 0
	assert_globals_p.SetCb(nil)
}

/* }}} */

func ZmStartupAssert(type_ int, module_number int) int {
	var ce zend.ZendClassEntry
	PhpAssertInitGlobals(&AssertGlobals)
	zend.ZendRegisterIniEntries(IniEntries, module_number)
	zend.ZendRegisterLongConstant("ASSERT_ACTIVE", g.SizeOf("\"ASSERT_ACTIVE\"")-1, ASSERT_ACTIVE, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("ASSERT_CALLBACK", g.SizeOf("\"ASSERT_CALLBACK\"")-1, ASSERT_CALLBACK, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("ASSERT_BAIL", g.SizeOf("\"ASSERT_BAIL\"")-1, ASSERT_BAIL, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("ASSERT_WARNING", g.SizeOf("\"ASSERT_WARNING\"")-1, ASSERT_WARNING, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("ASSERT_QUIET_EVAL", g.SizeOf("\"ASSERT_QUIET_EVAL\"")-1, ASSERT_QUIET_EVAL, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("ASSERT_EXCEPTION", g.SizeOf("\"ASSERT_EXCEPTION\"")-1, ASSERT_EXCEPTION, 1<<0|1<<1, module_number)
	memset(&ce, 0, g.SizeOf("zend_class_entry"))
	ce.name = zend.ZendStringInitInterned("AssertionError", g.SizeOf("\"AssertionError\"")-1, 1)
	ce.info.internal.builtin_functions = nil
	AssertionErrorCe = zend.ZendRegisterInternalClassEx(&ce, zend.ZendCeError)
	return zend.SUCCESS
}

/* }}} */

func ZmShutdownAssert(type_ int, module_number int) int {
	if AssertGlobals.GetCb() != nil {
		g.CondF(true, func() { return zend.Free(AssertGlobals.GetCb()) }, func() { return zend._efree(AssertGlobals.GetCb()) })
		AssertGlobals.SetCb(nil)
	}
	return zend.SUCCESS
}

/* }}} */

func ZmDeactivateAssert(type_ int, module_number int) int {
	if AssertGlobals.callback.u1.v.type_ != 0 {
		zend.ZvalPtrDtor(&(AssertGlobals.GetCallback()))
		&(AssertGlobals.GetCallback()).u1.type_info = 0
	}
	return zend.SUCCESS
}

/* }}} */

func ZmInfoAssert(zend_module *zend.ZendModuleEntry) { core.DisplayIniEntries(zend_module) }

/* }}} */

func ZifAssert(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var assertion *zend.Zval
	var description *zend.Zval = nil
	var val int
	var myeval *byte = nil
	var compiled_string_description *byte
	if AssertGlobals.GetActive() == 0 {
		return_value.u1.type_info = 3
		return
	}
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &assertion, 0)
			_optional = 1
			_i++
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &description, 0)
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	if assertion.u1.v.type_ == 6 {
		var retval zend.Zval
		var old_error_reporting int = 0
		if zend.ZendForbidDynamicCall("assert() with string argument") == zend.FAILURE {
			return_value.u1.type_info = 2
			return
		}
		core.PhpErrorDocref(nil, 1<<13, "Calling assert() with a string argument is deprecated")
		myeval = assertion.value.str.val
		if AssertGlobals.GetQuietEval() != 0 {
			old_error_reporting = zend.EG.error_reporting
			zend.EG.error_reporting = 0
		}
		compiled_string_description = zend.ZendMakeCompiledStringDescription("assert code")
		if zend.ZendEvalStringl(myeval, assertion.value.str.len_, &retval, compiled_string_description) == zend.FAILURE {
			zend._efree(compiled_string_description)
			if description == nil {
				zend.ZendThrowError(nil, "Failure evaluating code: %s%s", "\n", myeval)
			} else {
				var str *zend.ZendString = zend.ZvalGetString(description)
				zend.ZendThrowError(nil, "Failure evaluating code: %s%s:\"%s\"", "\n", str.val, myeval)
				zend.ZendStringReleaseEx(str, 0)
			}
			if AssertGlobals.GetBail() != 0 {
				zend._zendBailout(__FILE__, __LINE__)
			}
			return_value.u1.type_info = 2
			return
		}
		zend._efree(compiled_string_description)
		if AssertGlobals.GetQuietEval() != 0 {
			zend.EG.error_reporting = old_error_reporting
		}
		zend.ConvertToBoolean(&retval)
		val = retval.u1.v.type_ == 3
	} else {
		val = zend.ZendIsTrue(assertion)
	}
	if val != 0 {
		return_value.u1.type_info = 3
		return
	}
	if AssertGlobals.callback.u1.v.type_ == 0 && AssertGlobals.GetCb() != nil {
		var _s *byte = AssertGlobals.GetCb()
		var __z *zend.Zval = &(AssertGlobals.GetCallback())
		var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
	}
	if AssertGlobals.callback.u1.v.type_ != 0 {
		var args []zend.Zval
		var retval zend.Zval
		var lineno uint32 = zend.ZendGetExecutedLineno()
		var filename *byte = zend.ZendGetExecutedFilename()
		var _s *byte = g.Cond(filename != nil, filename, "")
		var __z *zend.Zval = &args[0]
		var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		var __z *zend.Zval = &args[1]
		__z.value.lval = lineno
		__z.u1.type_info = 4
		var _s *byte = g.Cond(myeval != nil, myeval, "")
		var __z *zend.Zval = &args[2]
		var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		&retval.u1.type_info = 2

		/* XXX do we want to check for error here? */

		if description == nil {
			zend._callUserFunctionEx(nil, &(AssertGlobals.GetCallback()), &retval, 3, args, 1)
			zend.ZvalPtrDtor(&args[2])
			zend.ZvalPtrDtor(&args[0])
		} else {
			var __z *zend.Zval = &args[3]
			var __s *zend.ZendString = zend.ZvalGetString(description)
			__z.value.str = __s
			if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
				__z.u1.type_info = 6
			} else {
				__z.u1.type_info = 6 | 1<<0<<8
			}
			zend._callUserFunctionEx(nil, &(AssertGlobals.GetCallback()), &retval, 4, args, 1)
			zend.ZvalPtrDtor(&args[3])
			zend.ZvalPtrDtor(&args[2])
			zend.ZvalPtrDtor(&args[0])
		}
		zend.ZvalPtrDtor(&retval)
	}
	if AssertGlobals.GetException() != 0 {
		if description == nil {
			zend.ZendThrowException(AssertionErrorCe, nil, 1<<0)
		} else if description.u1.v.type_ == 8 && zend.InstanceofFunction(description.value.obj.ce, zend.ZendCeThrowable) != 0 {
			zend.ZvalAddrefP(description)
			zend.ZendThrowExceptionObject(description)
		} else {
			var str *zend.ZendString = zend.ZvalGetString(description)
			zend.ZendThrowException(AssertionErrorCe, str.val, 1<<0)
			zend.ZendStringReleaseEx(str, 0)
		}
	} else if AssertGlobals.GetWarning() != 0 {
		if description == nil {
			if myeval != nil {
				core.PhpErrorDocref(nil, 1<<1, "Assertion \"%s\" failed", myeval)
			} else {
				core.PhpErrorDocref(nil, 1<<1, "Assertion failed")
			}
		} else {
			var str *zend.ZendString = zend.ZvalGetString(description)
			if myeval != nil {
				core.PhpErrorDocref(nil, 1<<1, "%s: \"%s\" failed", str.val, myeval)
			} else {
				core.PhpErrorDocref(nil, 1<<1, "%s failed", str.val)
			}
			zend.ZendStringReleaseEx(str, 0)
		}
	}
	if AssertGlobals.GetBail() != 0 {
		zend._zendBailout(__FILE__, __LINE__)
	}
	return_value.u1.type_info = 2
	return
}

/* }}} */

func ZifAssertOptions(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var value *zend.Zval = nil
	var what zend.ZendLong
	var oldint zend.ZendBool
	var ac int = execute_data.This.u2.num_args
	var key *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &what, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &value, 0)
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
		oldint = AssertGlobals.GetActive()
		if ac == 2 {
			var value_str *zend.ZendString = zend.ZvalTryGetString(value)
			if value_str == nil {
				return
			}
			key = zend.ZendStringInit("assert.active", g.SizeOf("\"assert.active\"")-1, 0)
			zend.ZendAlterIniEntryEx(key, value_str, 1<<0, 1<<4, 0)
			zend.ZendStringReleaseEx(key, 0)
			zend.ZendStringReleaseEx(value_str, 0)
		}
		var __z *zend.Zval = return_value
		__z.value.lval = oldint
		__z.u1.type_info = 4
		return
		break
	case ASSERT_BAIL:
		oldint = AssertGlobals.GetBail()
		if ac == 2 {
			var value_str *zend.ZendString = zend.ZvalTryGetString(value)
			if value_str == nil {
				return
			}
			key = zend.ZendStringInit("assert.bail", g.SizeOf("\"assert.bail\"")-1, 0)
			zend.ZendAlterIniEntryEx(key, value_str, 1<<0, 1<<4, 0)
			zend.ZendStringReleaseEx(key, 0)
			zend.ZendStringReleaseEx(value_str, 0)
		}
		var __z *zend.Zval = return_value
		__z.value.lval = oldint
		__z.u1.type_info = 4
		return
		break
	case ASSERT_QUIET_EVAL:
		oldint = AssertGlobals.GetQuietEval()
		if ac == 2 {
			var value_str *zend.ZendString = zend.ZvalTryGetString(value)
			if value_str == nil {
				return
			}
			key = zend.ZendStringInit("assert.quiet_eval", g.SizeOf("\"assert.quiet_eval\"")-1, 0)
			zend.ZendAlterIniEntryEx(key, value_str, 1<<0, 1<<4, 0)
			zend.ZendStringReleaseEx(key, 0)
			zend.ZendStringReleaseEx(value_str, 0)
		}
		var __z *zend.Zval = return_value
		__z.value.lval = oldint
		__z.u1.type_info = 4
		return
		break
	case ASSERT_WARNING:
		oldint = AssertGlobals.GetWarning()
		if ac == 2 {
			var value_str *zend.ZendString = zend.ZvalTryGetString(value)
			if value_str == nil {
				return
			}
			key = zend.ZendStringInit("assert.warning", g.SizeOf("\"assert.warning\"")-1, 0)
			zend.ZendAlterIniEntryEx(key, value_str, 1<<0, 1<<4, 0)
			zend.ZendStringReleaseEx(key, 0)
			zend.ZendStringReleaseEx(value_str, 0)
		}
		var __z *zend.Zval = return_value
		__z.value.lval = oldint
		__z.u1.type_info = 4
		return
		break
	case ASSERT_CALLBACK:
		if AssertGlobals.callback.u1.v.type_ != 0 {
			var _z1 *zend.Zval = return_value
			var _z2 *zend.Zval = &(AssertGlobals.GetCallback())
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			if (_t & 0xff00) != 0 {
				zend.ZendGcAddref(&_gc.gc)
			}
		} else if AssertGlobals.GetCb() != nil {
			var _s *byte = AssertGlobals.GetCb()
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
		} else {
			return_value.u1.type_info = 1
		}
		if ac == 2 {
			zend.ZvalPtrDtor(&(AssertGlobals.GetCallback()))
			var _z1 *zend.Zval = &(AssertGlobals.GetCallback())
			var _z2 *zend.Zval = value
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			if (_t & 0xff00) != 0 {
				zend.ZendGcAddref(&_gc.gc)
			}
		}
		return
	case ASSERT_EXCEPTION:
		oldint = AssertGlobals.GetException()
		if ac == 2 {
			var val *zend.ZendString = zend.ZvalTryGetString(value)
			if val == nil {
				return
			}
			key = zend.ZendStringInit("assert.exception", g.SizeOf("\"assert.exception\"")-1, 0)
			zend.ZendAlterIniEntryEx(key, val, 1<<0, 1<<4, 0)
			zend.ZendStringReleaseEx(val, 0)
			zend.ZendStringReleaseEx(key, 0)
		}
		var __z *zend.Zval = return_value
		__z.value.lval = oldint
		__z.u1.type_info = 4
		return
		break
	default:
		core.PhpErrorDocref(nil, 1<<1, "Unknown value "+"%"+"lld", what)
		break
	}
	return_value.u1.type_info = 2
	return
}

/* }}} */
