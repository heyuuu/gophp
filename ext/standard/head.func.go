package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"github.com/heyuuu/gophp/zend/zpp"
)

func ZifHeader(executeData zpp.Ex, return_value zpp.Ret, header *types.Zval, _ zpp.Opt, replace *types.Zval, httpResponseCode *types.Zval) {
	var rep bool = 1
	var ctr core.SapiHeaderLine = core.MakeSapiHeaderLine(0)
	var len_ int
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 3, 0)
			ctr.GetLine(), len_ = fp.ParseString()
			fp.StartOptional()
			rep = fp.ParseBool()
			ctr.GetResponseCode() = fp.ParseLong()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	ctr.SetLineLen(uint32(len_))
	core.SapiHeaderOp(lang.Cond(rep != 0, core.SAPI_HEADER_REPLACE, core.SAPI_HEADER_ADD), &ctr)
}
func ZifHeaderRemove(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, name *types.Zval) {
	var ctr core.SapiHeaderLine = core.MakeSapiHeaderLine(0)
	var len_ int = 0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 0, 1, 0)
			fp.StartOptional()
			ctr.GetLine(), len_ = fp.ParseString()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	ctr.SetLineLen(uint32(len_))
	core.SapiHeaderOp(lang.Cond(executeData.NumArgs() == 0, core.SAPI_HEADER_DELETE_ALL, core.SAPI_HEADER_DELETE), &ctr)
}
func PhpHeader() int {
	if core.SapiSendHeaders() == types.FAILURE || core.SG__().RequestInfo.headers_only {
		return 0
	} else {
		return 1
	}
}
func PhpSetcookie(
	name *types.String,
	value *types.String,
	expires int64,
	path *types.String,
	domain *types.String,
	secure int,
	httponly int,
	samesite *types.String,
	url_encode int,
) int {
	var dt *types.String
	var ctr core.SapiHeaderLine = core.MakeSapiHeaderLine(0)
	var result int
	var buf zend.SmartStr = zend.MakeSmartStr(0)
	if name.GetLen() == 0 {
		faults.Error(faults.E_WARNING, "Cookie names must not be empty")
		return types.FAILURE
	} else if strpbrk(name.GetVal(), "=,; \t\r\n013014") != nil {
		faults.Error(faults.E_WARNING, "Cookie names cannot contain any of the following '=,; \\t\\r\\n\\013\\014'")
		return types.FAILURE
	}
	if url_encode == 0 && value != nil && strpbrk(value.GetVal(), ",; \t\r\n013014") != nil {
		faults.Error(faults.E_WARNING, "Cookie values cannot contain any of the following ',; \\t\\r\\n\\013\\014'")
		return types.FAILURE
	}
	if path != nil && strpbrk(path.GetVal(), ",; \t\r\n013014") != nil {
		faults.Error(faults.E_WARNING, "Cookie paths cannot contain any of the following ',; \\t\\r\\n\\013\\014'")
		return types.FAILURE
	}
	if domain != nil && strpbrk(domain.GetVal(), ",; \t\r\n013014") != nil {
		faults.Error(faults.E_WARNING, "Cookie domains cannot contain any of the following ',; \\t\\r\\n\\013\\014'")
		return types.FAILURE
	}
	if value == nil || value.GetLen() == 0 {

		/*
		 * MSIE doesn't delete a cookie when you set it to a null value
		 * so in order to force cookies to be deleted, even on MSIE, we
		 * pick an expiry date in the past
		 */

		dt = php_format_date("D, d-M-Y H:i:s T", b.SizeOf("\"D, d-M-Y H:i:s T\"")-1, 1, 0)
		buf.WriteString("Set-Cookie: ")
		buf.WriteString(name.GetStr())
		buf.WriteString("=deleted; expires=")
		buf.WriteString(dt.GetStr())
		buf.WriteString("; Max-Age=0")
		//types.ZendStringFree(dt)
	} else {
		buf.WriteString("Set-Cookie: ")
		buf.WriteString(name.GetStr())
		buf.WriteByte('=')
		if url_encode != 0 {
			var encodedValue = PhpRawUrlEncode(value.GetStr())
			buf.WriteString(encodedValue)
		} else {
			buf.WriteString(value.GetStr())
		}
		if expires > 0 {
			var p *byte
			var diff float64
			buf.WriteString(b.CastStrAuto(COOKIE_EXPIRES))
			dt = php_format_date("D, d-M-Y H:i:s T", b.SizeOf("\"D, d-M-Y H:i:s T\"")-1, expires, 0)

			/* check to make sure that the year does not exceed 4 digits in length */

			p = operators.ZendMemrchr(dt.GetVal(), '-', dt.GetLen())
			if p == nil || (*(p + 5)) != ' ' {
				//types.ZendStringFree(dt)
				buf.Free()
				faults.Error(faults.E_WARNING, "Expiry date cannot have a year greater than 9999")
				return types.FAILURE
			}
			buf.WriteString(dt.GetStr())
			//types.ZendStringFree(dt)
			diff = difftime(expires, php_time())
			if diff < 0 {
				diff = 0
			}
			buf.WriteString(b.CastStrAuto(COOKIE_MAX_AGE))
			buf.WriteLong(zend.ZendLong(diff))
		}
	}
	if path != nil && path.GetLen() != 0 {
		buf.WriteString(b.CastStrAuto(COOKIE_PATH))
		buf.WriteString(path.GetStr())
	}
	if domain != nil && domain.GetLen() != 0 {
		buf.WriteString(b.CastStrAuto(COOKIE_DOMAIN))
		buf.WriteString(domain.GetStr())
	}
	if secure != 0 {
		buf.WriteString(b.CastStrAuto(COOKIE_SECURE))
	}
	if httponly != 0 {
		buf.WriteString(b.CastStrAuto(COOKIE_HTTPONLY))
	}
	if samesite != nil && samesite.GetLen() != 0 {
		buf.WriteString(b.CastStrAuto(COOKIE_SAMESITE))
		buf.WriteString(samesite.GetStr())
	}
	ctr.SetLine(buf.GetS().GetVal())
	ctr.SetLineLen(uint32(buf.GetS().GetLen()))
	result = core.SapiHeaderOp(core.SAPI_HEADER_ADD, &ctr)
	// types.ZendStringRelease(buf.GetS())
	return result
}
func PhpHeadParseCookieOptionsArray(
	options *types.Zval,
	expires *zend.ZendLong,
	path **types.String,
	domain **types.String,
	secure *bool,
	httponly *bool,
	samesite **types.String,
) {
	var found int = 0
	var key *types.String
	var value *types.Zval
	options.Array().Foreach(func(key types.ArrayKey, value *types.Zval) {
		if key.IsStrKey() {
			strKey := key.StrKey()
			if ascii.StrCaseEquals(strKey, "expires") {
				*expires = operators.ZvalGetLong(value)
				found++
			} else if ascii.StrCaseEquals(strKey, "path") {
				*path = operators.ZvalGetString(value)
				found++
			} else if ascii.StrCaseEquals(strKey, "domain") {
				*domain = operators.ZvalGetString(value)
				found++
			} else if ascii.StrCaseEquals(strKey, "secure") {
				*secure = operators.IZendIsTrue(value)
				found++
			} else if ascii.StrCaseEquals(strKey, "httponly") {
				*httponly = operators.IZendIsTrue(value)
				found++
			} else if ascii.StrCaseEquals(strKey, "samesite") {
				*samesite = operators.ZvalGetString(value)
				found++
			} else {
				core.PhpErrorDocref("", faults.E_WARNING, "Unrecognized key '%s' found in the options array", strKey)
			}
		} else {
			core.PhpErrorDocref("", faults.E_WARNING, "Numeric key found in the options array")
		}
	})

	/* Array is not empty but no valid keys were found */

	if found == 0 && options.Array().Len() > 0 {
		core.PhpErrorDocref("", faults.E_WARNING, "No valid options were found in the given array")
	}

	/* Array is not empty but no valid keys were found */
}
func ZifSetcookie(executeData zpp.Ex, return_value zpp.Ret, name *types.Zval, _ zpp.Opt, value *types.Zval, expiresOrOptions *types.Zval, path *types.Zval, domain *types.Zval, secure *types.Zval, httponly *types.Zval) {
	var expires_or_options *types.Zval = nil
	var name *types.String
	var value *types.String = nil
	var path *types.String = nil
	var domain *types.String = nil
	var samesite *types.String = nil
	var expires zend.ZendLong = 0
	var secure bool = 0
	var httponly bool = 0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 7, 0)
			name = fp.ParseStr()
			fp.StartOptional()
			value = fp.ParseStr()
			expires_or_options = fp.ParseZval()
			path = fp.ParseStr()
			domain = fp.ParseStr()
			secure = fp.ParseBool()
			httponly = fp.ParseBool()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if expires_or_options != nil {
		if expires_or_options.IsType(types.IsArray) {
			if executeData.NumArgs() > 3 {
				core.PhpErrorDocref("", faults.E_WARNING, "Cannot pass arguments after the options array")
				return_value.SetFalse()
				return
			}
			PhpHeadParseCookieOptionsArray(expires_or_options, &expires, &path, &domain, &secure, &httponly, &samesite)
		} else {
			expires = operators.ZvalGetLong(expires_or_options)
		}
	}
	if zend.EG__().NoException() {
		if PhpSetcookie(name, value, expires, path, domain, secure, httponly, samesite, 1) == types.SUCCESS {
			return_value.SetTrue()
		} else {
			return_value.SetFalse()
		}
	}
	if expires_or_options != nil && expires_or_options.IsType(types.IsArray) {
		if path != nil {
			// types.ZendStringRelease(path)
		}
		if domain != nil {
			// types.ZendStringRelease(domain)
		}
		if samesite != nil {
			// types.ZendStringRelease(samesite)
		}
	}
}
func ZifSetrawcookie(executeData zpp.Ex, return_value zpp.Ret, name *types.Zval, _ zpp.Opt, value *types.Zval, expiresOrOptions *types.Zval, path *types.Zval, domain *types.Zval, secure *types.Zval, httponly *types.Zval) {
	var expires_or_options *types.Zval = nil
	var name *types.String
	var value *types.String = nil
	var path *types.String = nil
	var domain *types.String = nil
	var samesite *types.String = nil
	var expires zend.ZendLong = 0
	var secure bool = 0
	var httponly bool = 0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 7, 0)
			name = fp.ParseStr()
			fp.StartOptional()
			value = fp.ParseStr()
			expires_or_options = fp.ParseZval()
			path = fp.ParseStr()
			domain = fp.ParseStr()
			secure = fp.ParseBool()
			httponly = fp.ParseBool()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if expires_or_options != nil {
		if expires_or_options.IsType(types.IsArray) {
			if executeData.NumArgs() > 3 {
				core.PhpErrorDocref("", faults.E_WARNING, "Cannot pass arguments after the options array")
				return_value.SetFalse()
				return
			}
			PhpHeadParseCookieOptionsArray(expires_or_options, &expires, &path, &domain, &secure, &httponly, &samesite)
		} else {
			expires = operators.ZvalGetLong(expires_or_options)
		}
	}
	if zend.EG__().NoException() {
		if PhpSetcookie(name, value, expires, path, domain, secure, httponly, samesite, 0) == types.SUCCESS {
			return_value.SetTrue()
		} else {
			return_value.SetFalse()
		}
	}
	if expires_or_options != nil && expires_or_options.IsType(types.IsArray) {
		if path != nil {
			// types.ZendStringRelease(path)
		}
		if domain != nil {
			// types.ZendStringRelease(domain)
		}
		if samesite != nil {
			// types.ZendStringRelease(samesite)
		}
	}
}
func ZifHeadersSent(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, file_ zpp.RefZval, line_ zpp.RefZval) {
	var arg1 *types.Zval = nil
	var arg2 *types.Zval = nil
	var file string = ""
	var line int = 0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 0, 2, 0)
			fp.StartOptional()
			arg1 = fp.ParseZval()
			arg2 = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if core.SG__().headers_sent {
		line = core.OG__().StartLineno()
		file = core.OG__().StartFilename()

	}
	switch executeData.NumArgs() {
	case 2:
		zend.ZEND_TRY_ASSIGN_REF_LONG(arg2, line)
		fallthrough
	case 1:
		if file != "" {
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
func ZifHeadersList(executeData zpp.Ex, return_value zpp.Ret) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	zend.ArrayInit(return_value)
	core.SG__().SapiHeaders().GetHeaders().Each(func(h *core.SapiHeader) {
		PhpHeadApplyHeaderListToHash(h, return_value)
	})
}
func ZifHttpResponseCode(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, responseCode *types.Zval) {
	var response_code zend.ZendLong = 0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 0, 1, 0)
			fp.StartOptional()
			response_code = fp.ParseLong()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if response_code != 0 {
		var old_response_code zend.ZendLong
		old_response_code = core.SG__().SapiHeaders().HttpResponseCode()
		core.SG__().SapiHeaders().SetHttpResponseCode(response_code)
		if old_response_code != 0 {
			return_value.SetLong(old_response_code)
			return
		}
		return_value.SetTrue()
		return
	}
	if !(core.SG__().SapiHeaders().HttpResponseCode()) {
		return_value.SetFalse()
		return
	}
	return_value.SetLong(core.SG__().SapiHeaders().HttpResponseCode())
	return
}
