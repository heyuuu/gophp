// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
	"sik/zend/argparse"
	"sik/zend/types"
)

func ZifHeader(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var rep types.ZendBool = 1
	var ctr core.SapiHeaderLine = core.MakeSapiHeaderLine(0)
	var len_ int
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			ctr.GetLine(), len_ = fp.ParseString()
			fp.StartOptional()
			rep = fp.ParseBool()
			ctr.GetResponseCode() = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	ctr.SetLineLen(uint32(len_))
	core.SapiHeaderOp(b.Cond(rep != 0, core.SAPI_HEADER_REPLACE, core.SAPI_HEADER_ADD), &ctr)
}
func ZifHeaderRemove(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var ctr core.SapiHeaderLine = core.MakeSapiHeaderLine(0)
	var len_ int = 0
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			fp.StartOptional()
			ctr.GetLine(), len_ = fp.ParseString()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	ctr.SetLineLen(uint32(len_))
	core.SapiHeaderOp(b.Cond(executeData.NumArgs() == 0, core.SAPI_HEADER_DELETE_ALL, core.SAPI_HEADER_DELETE), &ctr)
}
func PhpHeader() int {
	if core.SapiSendHeaders() == types.FAILURE || core.SG__().request_info.headers_only {
		return 0
	} else {
		return 1
	}
}
func PhpSetcookie(
	name *types.ZendString,
	value *types.ZendString,
	expires int64,
	path *types.ZendString,
	domain *types.ZendString,
	secure int,
	httponly int,
	samesite *types.ZendString,
	url_encode int,
) int {
	var dt *types.ZendString
	var ctr core.SapiHeaderLine = core.MakeSapiHeaderLine(0)
	var result int
	var buf zend.SmartStr = zend.MakeSmartStr(0)
	if name.GetLen() == 0 {
		zend.ZendError(zend.E_WARNING, "Cookie names must not be empty")
		return types.FAILURE
	} else if strpbrk(name.GetVal(), "=,; \t\r\n013014") != nil {
		zend.ZendError(zend.E_WARNING, "Cookie names cannot contain any of the following '=,; \\t\\r\\n\\013\\014'")
		return types.FAILURE
	}
	if url_encode == 0 && value != nil && strpbrk(value.GetVal(), ",; \t\r\n013014") != nil {
		zend.ZendError(zend.E_WARNING, "Cookie values cannot contain any of the following ',; \\t\\r\\n\\013\\014'")
		return types.FAILURE
	}
	if path != nil && strpbrk(path.GetVal(), ",; \t\r\n013014") != nil {
		zend.ZendError(zend.E_WARNING, "Cookie paths cannot contain any of the following ',; \\t\\r\\n\\013\\014'")
		return types.FAILURE
	}
	if domain != nil && strpbrk(domain.GetVal(), ",; \t\r\n013014") != nil {
		zend.ZendError(zend.E_WARNING, "Cookie domains cannot contain any of the following ',; \\t\\r\\n\\013\\014'")
		return types.FAILURE
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
		types.ZendStringFree(dt)
	} else {
		buf.AppendString("Set-Cookie: ")
		buf.AppendString(name.GetStr())
		buf.AppendByte('=')
		if url_encode != 0 {
			var encoded_value *types.ZendString = PhpRawUrlEncode(value.GetVal(), value.GetLen())
			buf.AppendString(encoded_value.GetStr())
			types.ZendStringReleaseEx(encoded_value, 0)
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
				types.ZendStringFree(dt)
				buf.Free()
				zend.ZendError(zend.E_WARNING, "Expiry date cannot have a year greater than 9999")
				return types.FAILURE
			}
			buf.AppendString(dt.GetStr())
			types.ZendStringFree(dt)
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
	types.ZendStringRelease(buf.GetS())
	return result
}
func PhpHeadParseCookieOptionsArray(
	options *types.Zval,
	expires *zend.ZendLong,
	path **types.ZendString,
	domain **types.ZendString,
	secure *types.ZendBool,
	httponly *types.ZendBool,
	samesite **types.ZendString,
) {
	var found int = 0
	var key *types.ZendString
	var value *types.Zval
	var __ht *types.HashTable = options.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		key = _p.GetKey()
		value = _z
		if key != nil {
			if types.ZendStringEqualsLiteralCi(key, "expires") {
				*expires = zend.ZvalGetLong(value)
				found++
			} else if types.ZendStringEqualsLiteralCi(key, "path") {
				*path = zend.ZvalGetString(value)
				found++
			} else if types.ZendStringEqualsLiteralCi(key, "domain") {
				*domain = zend.ZvalGetString(value)
				found++
			} else if types.ZendStringEqualsLiteralCi(key, "secure") {
				*secure = zend.ZvalIsTrue(value)
				found++
			} else if types.ZendStringEqualsLiteralCi(key, "httponly") {
				*httponly = zend.ZvalIsTrue(value)
				found++
			} else if types.ZendStringEqualsLiteralCi(key, "samesite") {
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

	if found == 0 && types.Z_ARRVAL_P(options).GetNNumOfElements() > 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "No valid options were found in the given array")
	}

	/* Array is not empty but no valid keys were found */
}
func ZifSetcookie(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var expires_or_options *types.Zval = nil
	var name *types.ZendString
	var value *types.ZendString = nil
	var path *types.ZendString = nil
	var domain *types.ZendString = nil
	var samesite *types.ZendString = nil
	var expires zend.ZendLong = 0
	var secure types.ZendBool = 0
	var httponly types.ZendBool = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 7
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgStr(_arg, &name, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			fp.StartOptional()
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgStr(_arg, &value, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			argparse.Z_PARAM_PROLOGUE(0, 0)
			argparse.ZendParseArgZvalDeref(_arg, &expires_or_options, 0)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgStr(_arg, &path, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgStr(_arg, &domain, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			secure = fp.ParseBool()
			httponly = fp.ParseBool()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if expires_or_options != nil {
		if expires_or_options.IsType(types.IS_ARRAY) {
			if executeData.NumArgs() > 3 {
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
		if PhpSetcookie(name, value, expires, path, domain, secure, httponly, samesite, 1) == types.SUCCESS {
			return_value.SetTrue()
		} else {
			return_value.SetFalse()
		}
	}
	if expires_or_options != nil && expires_or_options.IsType(types.IS_ARRAY) {
		if path != nil {
			types.ZendStringRelease(path)
		}
		if domain != nil {
			types.ZendStringRelease(domain)
		}
		if samesite != nil {
			types.ZendStringRelease(samesite)
		}
	}
}
func ZifSetrawcookie(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var expires_or_options *types.Zval = nil
	var name *types.ZendString
	var value *types.ZendString = nil
	var path *types.ZendString = nil
	var domain *types.ZendString = nil
	var samesite *types.ZendString = nil
	var expires zend.ZendLong = 0
	var secure types.ZendBool = 0
	var httponly types.ZendBool = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 7
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgStr(_arg, &name, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			fp.StartOptional()
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgStr(_arg, &value, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			argparse.Z_PARAM_PROLOGUE(0, 0)
			argparse.ZendParseArgZvalDeref(_arg, &expires_or_options, 0)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgStr(_arg, &path, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgStr(_arg, &domain, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			secure = fp.ParseBool()
			httponly = fp.ParseBool()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if expires_or_options != nil {
		if expires_or_options.IsType(types.IS_ARRAY) {
			if executeData.NumArgs() > 3 {
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
		if PhpSetcookie(name, value, expires, path, domain, secure, httponly, samesite, 0) == types.SUCCESS {
			return_value.SetTrue()
		} else {
			return_value.SetFalse()
		}
	}
	if expires_or_options != nil && expires_or_options.IsType(types.IS_ARRAY) {
		if path != nil {
			types.ZendStringRelease(path)
		}
		if domain != nil {
			types.ZendStringRelease(domain)
		}
		if samesite != nil {
			types.ZendStringRelease(samesite)
		}
	}
}
func ZifHeadersSent(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var arg1 *types.Zval = nil
	var arg2 *types.Zval = nil
	var file *byte = ""
	var line int = 0
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			fp.StartOptional()
			argparse.Z_PARAM_PROLOGUE(0, 0)
			argparse.ZendParseArgZvalDeref(_arg, &arg1, 0)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			argparse.ZendParseArgZvalDeref(_arg, &arg2, 0)
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if core.SG__().headers_sent {
		line = core.PhpOutputGetStartLineno()
		file = core.PhpOutputGetStartFilename()
	}
	switch executeData.NumArgs() {
	case 2:
		zend.ZEND_TRY_ASSIGN_REF_LONG(arg2, line)
		fallthrough
	case 1:
		if file != nil {
			zend.ZEND_TRY_ASSIGN_REF_STRING(arg1, file)
		} else {
			zend.ZEND_TRY_ASSIGN_REF_EMPTY_STRING(arg1)
		}
	}
	if core.SG__().headers_sent {
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
		zend.AddNextIndexString((*types.Zval)(arg), (*byte)(sapi_header.GetHeader()))
	}
}
func ZifHeadersList(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	zend.ArrayInit(return_value)
	core.SG__().sapi_headers.headers.ApplyWithArgument(PhpHeadApplyHeaderListToHash, return_value)
}
func ZifHttpResponseCode(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var response_code zend.ZendLong = 0
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			fp.StartOptional()
			response_code = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if response_code != 0 {
		var old_response_code zend.ZendLong
		old_response_code = core.SG__().sapi_headers.http_response_code
		core.SG__().sapi_headers.http_response_code = int(response_code)
		if old_response_code != 0 {
			return_value.SetLong(old_response_code)
			return
		}
		return_value.SetTrue()
		return
	}
	if !(core.SG__().sapi_headers.http_response_code) {
		return_value.SetFalse()
		return
	}
	return_value.SetLong(core.SG__().sapi_headers.http_response_code)
	return
}
