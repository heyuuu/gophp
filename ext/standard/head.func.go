// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

func ZifHeader(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var rep zend.ZendBool = 1
	var ctr core.SapiHeaderLine = core.MakeSapiHeaderLine(0)
	var len_ int
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3
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
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
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
			if zend.ZendParseArgString(_arg, ctr.GetLine(), &len_, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &rep, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, ctr.GetResponseCode(), &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
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
	ctr.SetLineLen(uint32(len_))
	core.SapiHeaderOp(b.Cond(rep != 0, core.SAPI_HEADER_REPLACE, core.SAPI_HEADER_ADD), &ctr)
}
func ZifHeaderRemove(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var ctr core.SapiHeaderLine = core.MakeSapiHeaderLine(0)
	var len_ int = 0
	for {
		var _flags int = 0
		var _min_num_args int = 0
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
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
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
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, ctr.GetLine(), &len_, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
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
	ctr.SetLineLen(uint32(len_))
	core.SapiHeaderOp(b.Cond(zend.ZEND_NUM_ARGS() == 0, core.SAPI_HEADER_DELETE_ALL, core.SAPI_HEADER_DELETE), &ctr)
}
func PhpHeader() int {
	if core.SapiSendHeaders() == zend.FAILURE || core.SG(request_info).headers_only {
		return 0
	} else {
		return 1
	}
}
func PhpSetcookie(
	name *zend.ZendString,
	value *zend.ZendString,
	expires int64,
	path *zend.ZendString,
	domain *zend.ZendString,
	secure int,
	httponly int,
	samesite *zend.ZendString,
	url_encode int,
) int {
	var dt *zend.ZendString
	var ctr core.SapiHeaderLine = core.MakeSapiHeaderLine(0)
	var result int
	var buf zend.SmartStr = zend.MakeSmartStr(0)
	if name.GetLen() == 0 {
		zend.ZendError(zend.E_WARNING, "Cookie names must not be empty")
		return zend.FAILURE
	} else if strpbrk(name.GetVal(), "=,; \t\r\n013014") != nil {
		zend.ZendError(zend.E_WARNING, "Cookie names cannot contain any of the following '=,; \\t\\r\\n\\013\\014'")
		return zend.FAILURE
	}
	if url_encode == 0 && value != nil && strpbrk(value.GetVal(), ",; \t\r\n013014") != nil {
		zend.ZendError(zend.E_WARNING, "Cookie values cannot contain any of the following ',; \\t\\r\\n\\013\\014'")
		return zend.FAILURE
	}
	if path != nil && strpbrk(path.GetVal(), ",; \t\r\n013014") != nil {
		zend.ZendError(zend.E_WARNING, "Cookie paths cannot contain any of the following ',; \\t\\r\\n\\013\\014'")
		return zend.FAILURE
	}
	if domain != nil && strpbrk(domain.GetVal(), ",; \t\r\n013014") != nil {
		zend.ZendError(zend.E_WARNING, "Cookie domains cannot contain any of the following ',; \\t\\r\\n\\013\\014'")
		return zend.FAILURE
	}
	if value == nil || value.GetLen() == 0 {

		/*
		 * MSIE doesn't delete a cookie when you set it to a null value
		 * so in order to force cookies to be deleted, even on MSIE, we
		 * pick an expiry date in the past
		 */

		dt = php_format_date("D, d-M-Y H:i:s T", b.SizeOf("\"D, d-M-Y H:i:s T\"")-1, 1, 0)
		buf.AppendString("Set-Cookie: ")
		buf.AppendString(name.GetStr())
		buf.AppendString("=deleted; expires=")
		buf.AppendString(dt.GetStr())
		buf.AppendString("; Max-Age=0")
		zend.ZendStringFree(dt)
	} else {
		buf.AppendString("Set-Cookie: ")
		buf.AppendString(name.GetStr())
		buf.AppendByte('=')
		if url_encode != 0 {
			var encoded_value *zend.ZendString = PhpRawUrlEncode(value.GetVal(), value.GetLen())
			buf.AppendString(encoded_value.GetStr())
			zend.ZendStringReleaseEx(encoded_value, 0)
		} else {
			buf.AppendString(value.GetStr())
		}
		if expires > 0 {
			var p *byte
			var diff float64
			buf.AppendString(b.CastStrAuto(COOKIE_EXPIRES))
			dt = php_format_date("D, d-M-Y H:i:s T", b.SizeOf("\"D, d-M-Y H:i:s T\"")-1, expires, 0)

			/* check to make sure that the year does not exceed 4 digits in length */

			p = zend.ZendMemrchr(dt.GetVal(), '-', dt.GetLen())
			if p == nil || (*(p + 5)) != ' ' {
				zend.ZendStringFree(dt)
				buf.Free()
				zend.ZendError(zend.E_WARNING, "Expiry date cannot have a year greater than 9999")
				return zend.FAILURE
			}
			buf.AppendString(dt.GetStr())
			zend.ZendStringFree(dt)
			diff = difftime(expires, php_time())
			if diff < 0 {
				diff = 0
			}
			buf.AppendString(b.CastStrAuto(COOKIE_MAX_AGE))
			buf.AppendLong(zend.ZendLong(diff))
		}
	}
	if path != nil && path.GetLen() != 0 {
		buf.AppendString(b.CastStrAuto(COOKIE_PATH))
		buf.AppendString(path.GetStr())
	}
	if domain != nil && domain.GetLen() != 0 {
		buf.AppendString(b.CastStrAuto(COOKIE_DOMAIN))
		buf.AppendString(domain.GetStr())
	}
	if secure != 0 {
		buf.AppendString(b.CastStrAuto(COOKIE_SECURE))
	}
	if httponly != 0 {
		buf.AppendString(b.CastStrAuto(COOKIE_HTTPONLY))
	}
	if samesite != nil && samesite.GetLen() != 0 {
		buf.AppendString(b.CastStrAuto(COOKIE_SAMESITE))
		buf.AppendString(samesite.GetStr())
	}
	ctr.SetLine(buf.GetS().GetVal())
	ctr.SetLineLen(uint32(buf.GetS().GetLen()))
	result = core.SapiHeaderOp(core.SAPI_HEADER_ADD, &ctr)
	zend.ZendStringRelease(buf.GetS())
	return result
}
func PhpHeadParseCookieOptionsArray(
	options *zend.Zval,
	expires *zend.ZendLong,
	path **zend.ZendString,
	domain **zend.ZendString,
	secure *zend.ZendBool,
	httponly *zend.ZendBool,
	samesite **zend.ZendString,
) {
	var found int = 0
	var key *zend.ZendString
	var value *zend.Zval
	var __ht *zend.HashTable = options.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *zend.Zval = _p.GetVal()

		key = _p.GetKey()
		value = _z
		if key != nil {
			if zend.ZendStringEqualsLiteralCi(key, "expires") {
				*expires = zend.ZvalGetLong(value)
				found++
			} else if zend.ZendStringEqualsLiteralCi(key, "path") {
				*path = zend.ZvalGetString(value)
				found++
			} else if zend.ZendStringEqualsLiteralCi(key, "domain") {
				*domain = zend.ZvalGetString(value)
				found++
			} else if zend.ZendStringEqualsLiteralCi(key, "secure") {
				*secure = zend.ZvalIsTrue(value)
				found++
			} else if zend.ZendStringEqualsLiteralCi(key, "httponly") {
				*httponly = zend.ZvalIsTrue(value)
				found++
			} else if zend.ZendStringEqualsLiteralCi(key, "samesite") {
				*samesite = zend.ZvalGetString(value)
				found++
			} else {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Unrecognized key '%s' found in the options array", key.GetVal())
			}
		} else {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Numeric key found in the options array")
		}
	}

	/* Array is not empty but no valid keys were found */

	if found == 0 && zend.Z_ARRVAL_P(options).GetNNumOfElements() > 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "No valid options were found in the given array")
	}

	/* Array is not empty but no valid keys were found */
}
func ZifSetcookie(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var expires_or_options *zend.Zval = nil
	var name *zend.ZendString
	var value *zend.ZendString = nil
	var path *zend.ZendString = nil
	var domain *zend.ZendString = nil
	var samesite *zend.ZendString = nil
	var expires zend.ZendLong = 0
	var secure zend.ZendBool = 0
	var httponly zend.ZendBool = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 7
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
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
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
			if zend.ZendParseArgStr(_arg, &name, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgStr(_arg, &value, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &expires_or_options, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgStr(_arg, &path, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgStr(_arg, &domain, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &secure, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &httponly, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
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
	if expires_or_options != nil {
		if expires_or_options.IsType(zend.IS_ARRAY) {
			if zend.ZEND_NUM_ARGS() > 3 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Cannot pass arguments after the options array")
				return_value.SetFalse()
				return
			}
			PhpHeadParseCookieOptionsArray(expires_or_options, &expires, &path, &domain, &secure, &httponly, &samesite)
		} else {
			expires = zend.ZvalGetLong(expires_or_options)
		}
	}
	if zend.EG__().GetException() == nil {
		if PhpSetcookie(name, value, expires, path, domain, secure, httponly, samesite, 1) == zend.SUCCESS {
			return_value.SetTrue()
		} else {
			return_value.SetFalse()
		}
	}
	if expires_or_options != nil && expires_or_options.IsType(zend.IS_ARRAY) {
		if path != nil {
			zend.ZendStringRelease(path)
		}
		if domain != nil {
			zend.ZendStringRelease(domain)
		}
		if samesite != nil {
			zend.ZendStringRelease(samesite)
		}
	}
}
func ZifSetrawcookie(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var expires_or_options *zend.Zval = nil
	var name *zend.ZendString
	var value *zend.ZendString = nil
	var path *zend.ZendString = nil
	var domain *zend.ZendString = nil
	var samesite *zend.ZendString = nil
	var expires zend.ZendLong = 0
	var secure zend.ZendBool = 0
	var httponly zend.ZendBool = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 7
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
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
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
			if zend.ZendParseArgStr(_arg, &name, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgStr(_arg, &value, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &expires_or_options, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgStr(_arg, &path, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgStr(_arg, &domain, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &secure, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &httponly, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
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
	if expires_or_options != nil {
		if expires_or_options.IsType(zend.IS_ARRAY) {
			if zend.ZEND_NUM_ARGS() > 3 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Cannot pass arguments after the options array")
				return_value.SetFalse()
				return
			}
			PhpHeadParseCookieOptionsArray(expires_or_options, &expires, &path, &domain, &secure, &httponly, &samesite)
		} else {
			expires = zend.ZvalGetLong(expires_or_options)
		}
	}
	if zend.EG__().GetException() == nil {
		if PhpSetcookie(name, value, expires, path, domain, secure, httponly, samesite, 0) == zend.SUCCESS {
			return_value.SetTrue()
		} else {
			return_value.SetFalse()
		}
	}
	if expires_or_options != nil && expires_or_options.IsType(zend.IS_ARRAY) {
		if path != nil {
			zend.ZendStringRelease(path)
		}
		if domain != nil {
			zend.ZendStringRelease(domain)
		}
		if samesite != nil {
			zend.ZendStringRelease(samesite)
		}
	}
}
func ZifHeadersSent(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg1 *zend.Zval = nil
	var arg2 *zend.Zval = nil
	var file *byte = ""
	var line int = 0
	for {
		var _flags int = 0
		var _min_num_args int = 0
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
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
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
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &arg1, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &arg2, 0)
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
	if core.SG(headers_sent) {
		line = core.PhpOutputGetStartLineno()
		file = core.PhpOutputGetStartFilename()
	}
	switch zend.ZEND_NUM_ARGS() {
	case 2:
		zend.ZEND_TRY_ASSIGN_REF_LONG(arg2, line)
	case 1:
		if file != nil {
			zend.ZEND_TRY_ASSIGN_REF_STRING(arg1, file)
		} else {
			zend.ZEND_TRY_ASSIGN_REF_EMPTY_STRING(arg1)
		}
		break
	}
	if core.SG(headers_sent) {
		return_value.SetTrue()
		return
	} else {
		return_value.SetFalse()
		return
	}
}
func PhpHeadApplyHeaderListToHash(data any, arg any) {
	var sapi_header *core.SapiHeader = (*core.SapiHeader)(data)
	if arg && sapi_header != nil {
		zend.AddNextIndexString((*zend.Zval)(arg), (*byte)(sapi_header.GetHeader()))
	}
}
func ZifHeadersList(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ArrayInit(return_value)
	zend.ZendLlistApplyWithArgument(core.SG(sapi_headers).headers, PhpHeadApplyHeaderListToHash, return_value)
}
func ZifHttpResponseCode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var response_code zend.ZendLong = 0
	for {
		var _flags int = 0
		var _min_num_args int = 0
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
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
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
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &response_code, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
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
	if response_code != 0 {
		var old_response_code zend.ZendLong
		old_response_code = core.SG(sapi_headers).http_response_code
		core.SG(sapi_headers).http_response_code = int(response_code)
		if old_response_code != 0 {
			return_value.SetLong(old_response_code)
			return
		}
		return_value.SetTrue()
		return
	}
	if !(core.SG(sapi_headers).http_response_code) {
		return_value.SetFalse()
		return
	}
	return_value.SetLong(core.SG(sapi_headers).http_response_code)
	return
}
