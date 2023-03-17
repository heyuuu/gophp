// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

func ASSERTG(v __auto__) __auto__ { return AssertGlobals.v }
func SAFE_STRING(s *byte) string {
	if s != nil {
		return s
	} else {
		return ""
	}
}
func OnChangeCallback(
	entry *zend.ZendIniEntry,
	new_value *zend.ZendString,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	if zend.CurrEX() != nil {
		if ASSERTG(callback).u1.v.type_ != zend.IS_UNDEF {
			zend.ZvalPtrDtor(&(ASSERTG(callback)))
			ASSERTG(callback).SetUndef()
		}
		if new_value != nil && (ASSERTG(callback).u1.v.type_ != zend.IS_UNDEF || new_value.GetLen() != 0) {
			ASSERTG(callback).SetStringCopy(new_value)
		}
	} else {
		if ASSERTG(cb) {
			zend.Pefree(ASSERTG(cb), 1)
		}
		if new_value != nil && new_value.GetLen() != 0 {
			ASSERTG(cb) = zend.Pemalloc(new_value.GetLen()+1, 1)
			memcpy(ASSERTG(cb), new_value.GetVal(), new_value.GetLen())
			ASSERTG(cb)[new_value.GetLen()] = '0'
		} else {
			ASSERTG(cb) = nil
		}
	}
	return zend.SUCCESS
}
func PhpAssertInitGlobals(assert_globals_p *ZendAssertGlobals) {
	assert_globals_p.GetCallback().SetUndef()
	assert_globals_p.SetCb(nil)
}
func ZmStartupAssert(type_ int, module_number int) int {
	var ce zend.ZendClassEntry
	PhpAssertInitGlobals(&AssertGlobals)
	zend.REGISTER_INI_ENTRIES(module_number)
	zend.REGISTER_LONG_CONSTANT("ASSERT_ACTIVE", ASSERT_ACTIVE, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("ASSERT_CALLBACK", ASSERT_CALLBACK, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("ASSERT_BAIL", ASSERT_BAIL, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("ASSERT_WARNING", ASSERT_WARNING, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("ASSERT_QUIET_EVAL", ASSERT_QUIET_EVAL, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("ASSERT_EXCEPTION", ASSERT_EXCEPTION, zend.CONST_CS|zend.CONST_PERSISTENT)
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(zend.ZendStringInitInterned("AssertionError", b.SizeOf("\"AssertionError\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	AssertionErrorCe = zend.ZendRegisterInternalClassEx(&ce, zend.ZendCeError)
	return zend.SUCCESS
}
func ZmShutdownAssert(type_ int, module_number int) int {
	if ASSERTG(cb) {
		zend.Pefree(ASSERTG(cb), 1)
		ASSERTG(cb) = nil
	}
	return zend.SUCCESS
}
func ZmDeactivateAssert(type_ int, module_number int) int {
	if ASSERTG(callback).u1.v.type_ != zend.IS_UNDEF {
		zend.ZvalPtrDtor(&(ASSERTG(callback)))
		ASSERTG(callback).SetUndef()
	}
	return zend.SUCCESS
}
func ZmInfoAssert(zend_module *zend.ZendModuleEntry) { zend.DISPLAY_INI_ENTRIES() }
func ZifAssert(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var assertion *zend.Zval
	var description *zend.Zval = nil
	var val int
	var myeval *byte = nil
	var compiled_string_description *byte
	if !(ASSERTG(active)) {
		return_value.SetTrue()
		return
	}
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
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
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.CheckNumArgsException(_min_num_args, _max_num_args)
					} else {
						zend.CheckNumArgsError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = executeData.Arg(0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &assertion, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &description, 0)
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
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
	if assertion.IsType(zend.IS_STRING) {
		var retval zend.Zval
		var old_error_reporting int = 0
		if zend.ZendForbidDynamicCall("assert() with string argument") == zend.FAILURE {
			return_value.SetFalse()
			return
		}
		core.PhpErrorDocref(nil, zend.E_DEPRECATED, "Calling assert() with a string argument is deprecated")
		myeval = assertion.GetStr().GetVal()
		if ASSERTG(quiet_eval) {
			old_error_reporting = zend.EG__().GetErrorReporting()
			zend.EG__().SetErrorReporting(0)
		}
		compiled_string_description = zend.ZendMakeCompiledStringDescription("assert code")
		if zend.ZendEvalStringl(myeval, assertion.GetStr().GetLen(), &retval, compiled_string_description) == zend.FAILURE {
			zend.Efree(compiled_string_description)
			if description == nil {
				zend.ZendThrowError(nil, "Failure evaluating code: %s%s", core.PHP_EOL, myeval)
			} else {
				var str *zend.ZendString = zend.ZvalGetString(description)
				zend.ZendThrowError(nil, "Failure evaluating code: %s%s:\"%s\"", core.PHP_EOL, str.GetVal(), myeval)
				zend.ZendStringReleaseEx(str, 0)
			}
			if ASSERTG(bail) {
				zend.ZendBailout()
			}
			return_value.SetFalse()
			return
		}
		zend.Efree(compiled_string_description)
		if ASSERTG(quiet_eval) {
			zend.EG__().SetErrorReporting(old_error_reporting)
		}
		zend.ConvertToBoolean(&retval)
		val = retval.IsType(zend.IS_TRUE)
	} else {
		val = zend.ZendIsTrue(assertion)
	}
	if val != 0 {
		return_value.SetTrue()
		return
	}
	if ASSERTG(callback).u1.v.type_ == zend.IS_UNDEF && ASSERTG(cb) {
		(ASSERTG(callback)).SetRawString(b.CastStrAuto(ASSERTG(cb)))
	}
	if ASSERTG(callback).u1.v.type_ != zend.IS_UNDEF {
		var args []zend.Zval
		var retval zend.Zval
		var lineno uint32 = zend.ZendGetExecutedLineno()
		var filename *byte = zend.ZendGetExecutedFilename()
		args[0].SetRawString(b.CastStrAuto(SAFE_STRING(filename)))
		args[1].SetLong(lineno)
		args[2].SetRawString(b.CastStrAuto(SAFE_STRING(myeval)))
		retval.SetFalse()

		/* XXX do we want to check for error here? */

		if description == nil {
			zend.CallUserFunction(nil, &(ASSERTG(callback)), &retval, 3, args)
			zend.ZvalPtrDtor(&args[2])
			zend.ZvalPtrDtor(&args[0])
		} else {
			args[3].SetString(zend.ZvalGetString(description))
			zend.CallUserFunction(nil, &(ASSERTG(callback)), &retval, 4, args)
			zend.ZvalPtrDtor(&args[3])
			zend.ZvalPtrDtor(&args[2])
			zend.ZvalPtrDtor(&args[0])
		}
		zend.ZvalPtrDtor(&retval)
	}
	if ASSERTG(exception) {
		if description == nil {
			zend.ZendThrowException(AssertionErrorCe, nil, zend.E_ERROR)
		} else if description.IsType(zend.IS_OBJECT) && zend.InstanceofFunction(zend.Z_OBJCE_P(description), zend.ZendCeThrowable) != 0 {
			description.AddRefcount()
			zend.ZendThrowExceptionObject(description)
		} else {
			var str *zend.ZendString = zend.ZvalGetString(description)
			zend.ZendThrowException(AssertionErrorCe, str.GetVal(), zend.E_ERROR)
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
				core.PhpErrorDocref(nil, zend.E_WARNING, "%s: \"%s\" failed", str.GetVal(), myeval)
			} else {
				core.PhpErrorDocref(nil, zend.E_WARNING, "%s failed", str.GetVal())
			}
			zend.ZendStringReleaseEx(str, 0)
		}
	}
	if ASSERTG(bail) {
		zend.ZendBailout()
	}
	return_value.SetFalse()
	return
}
func ZifAssertOptions(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var value *zend.Zval = nil
	var what zend.ZendLong
	var oldint zend.ZendBool
	var ac int = executeData.NumArgs()
	var key *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
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
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.CheckNumArgsException(_min_num_args, _max_num_args)
					} else {
						zend.CheckNumArgsError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = executeData.Arg(0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &what, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &value, 0)
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
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
			if value_str == nil {
				return
			}
			key = zend.ZendStringInit("assert.active", b.SizeOf("\"assert.active\"")-1, 0)
			zend.ZendAlterIniEntryEx(key, value_str, core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME, 0)
			zend.ZendStringReleaseEx(key, 0)
			zend.ZendStringReleaseEx(value_str, 0)
		}
		return_value.SetLong(oldint)
		return
	case ASSERT_BAIL:
		oldint = ASSERTG(bail)
		if ac == 2 {
			var value_str *zend.ZendString = zend.ZvalTryGetString(value)
			if value_str == nil {
				return
			}
			key = zend.ZendStringInit("assert.bail", b.SizeOf("\"assert.bail\"")-1, 0)
			zend.ZendAlterIniEntryEx(key, value_str, core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME, 0)
			zend.ZendStringReleaseEx(key, 0)
			zend.ZendStringReleaseEx(value_str, 0)
		}
		return_value.SetLong(oldint)
		return
	case ASSERT_QUIET_EVAL:
		oldint = ASSERTG(quiet_eval)
		if ac == 2 {
			var value_str *zend.ZendString = zend.ZvalTryGetString(value)
			if value_str == nil {
				return
			}
			key = zend.ZendStringInit("assert.quiet_eval", b.SizeOf("\"assert.quiet_eval\"")-1, 0)
			zend.ZendAlterIniEntryEx(key, value_str, core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME, 0)
			zend.ZendStringReleaseEx(key, 0)
			zend.ZendStringReleaseEx(value_str, 0)
		}
		return_value.SetLong(oldint)
		return
	case ASSERT_WARNING:
		oldint = ASSERTG(warning)
		if ac == 2 {
			var value_str *zend.ZendString = zend.ZvalTryGetString(value)
			if value_str == nil {
				return
			}
			key = zend.ZendStringInit("assert.warning", b.SizeOf("\"assert.warning\"")-1, 0)
			zend.ZendAlterIniEntryEx(key, value_str, core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME, 0)
			zend.ZendStringReleaseEx(key, 0)
			zend.ZendStringReleaseEx(value_str, 0)
		}
		return_value.SetLong(oldint)
		return
	case ASSERT_CALLBACK:
		if ASSERTG(callback).u1.v.type_ != zend.IS_UNDEF {
			zend.ZVAL_COPY(return_value, &(ASSERTG(callback)))
		} else if ASSERTG(cb) {
			return_value.SetRawString(b.CastStrAuto(ASSERTG(cb)))
		} else {
			return_value.SetNull()
		}
		if ac == 2 {
			zend.ZvalPtrDtor(&(ASSERTG(callback)))
			zend.ZVAL_COPY(&(ASSERTG(callback)), value)
		}
		return
	case ASSERT_EXCEPTION:
		oldint = ASSERTG(exception)
		if ac == 2 {
			var val *zend.ZendString = zend.ZvalTryGetString(value)
			if val == nil {
				return
			}
			key = zend.ZendStringInit("assert.exception", b.SizeOf("\"assert.exception\"")-1, 0)
			zend.ZendAlterIniEntryEx(key, val, core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME, 0)
			zend.ZendStringReleaseEx(val, 0)
			zend.ZendStringReleaseEx(key, 0)
		}
		return_value.SetLong(oldint)
		return
	default:
		core.PhpErrorDocref(nil, zend.E_WARNING, "Unknown value "+zend.ZEND_LONG_FMT, what)
	}
	return_value.SetFalse()
	return
}
