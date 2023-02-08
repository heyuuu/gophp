// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

func TagDtor(zv *zend.Zval) { zend.Free(zv.GetPtr()) }
func PhpIniOnUpdateTags(
	entry *zend.ZendIniEntry,
	new_value *zend.ZendString,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
	type_ int,
) int {
	var ctx *UrlAdaptStateExT
	var key *byte
	var tmp *byte
	var lasts *byte = nil
	if type_ != 0 {
		ctx = &(BG(url_adapt_session_ex))
	} else {
		ctx = &(BG(url_adapt_output_ex))
	}
	tmp = zend.Estrndup(new_value.GetVal(), new_value.GetLen())
	if ctx.GetTags() != nil {
		ctx.GetTags().Destroy()
	} else {
		ctx.SetTags(zend.Malloc(b.SizeOf("HashTable")))
		if ctx.GetTags() == nil {
			zend.Efree(tmp)
			return zend.FAILURE
		}
	}
	zend.ZendHashInit(ctx.GetTags(), 0, nil, TagDtor, 1)
	for key = core.PhpStrtokR(tmp, ",", &lasts); key != nil; key = core.PhpStrtokR(nil, ",", &lasts) {
		var val *byte
		val = strchr(key, '=')
		if val != nil {
			var q *byte
			var keylen int
			var str *zend.ZendString
			b.PostInc(&(*val)) = '0'
			for q = key; *q; q++ {
				*q = tolower(*q)
			}
			keylen = q - key
			str = zend.ZendStringInit(key, keylen, 1)
			zend.GC_MAKE_PERSISTENT_LOCAL(str)
			zend.ZendHashAddMem(ctx.GetTags(), str, val, strlen(val)+1)
			zend.ZendStringReleaseEx(str, 1)
		}
	}
	zend.Efree(tmp)
	return zend.SUCCESS
}
func OnUpdateSessionTags(
	entry *zend.ZendIniEntry,
	new_value *zend.ZendString,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	return PhpIniOnUpdateTags(entry, new_value, mh_arg1, mh_arg2, mh_arg3, stage, 1)
}
func OnUpdateOutputTags(
	entry *zend.ZendIniEntry,
	new_value *zend.ZendString,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	return PhpIniOnUpdateTags(entry, new_value, mh_arg1, mh_arg2, mh_arg3, stage, 0)
}
func PhpIniOnUpdateHosts(
	entry *zend.ZendIniEntry,
	new_value *zend.ZendString,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
	type_ int,
) int {
	var hosts *zend.HashTable
	var key *byte
	var tmp *byte
	var lasts *byte = nil
	if type_ != 0 {
		hosts = &(BG(url_adapt_session_hosts_ht))
	} else {
		hosts = &(BG(url_adapt_output_hosts_ht))
	}
	hosts.Clean()

	/* Use user supplied host whitelist */

	tmp = zend.Estrndup(new_value.GetVal(), new_value.GetLen())
	for key = core.PhpStrtokR(tmp, ",", &lasts); key != nil; key = core.PhpStrtokR(nil, ",", &lasts) {
		var keylen int
		var tmp_key *zend.ZendString
		var q *byte
		for q = key; *q; q++ {
			*q = tolower(*q)
		}
		keylen = q - key
		if keylen > 0 {
			tmp_key = zend.ZendStringInit(key, keylen, 0)
			zend.ZendHashAddEmptyElement(hosts, tmp_key)
			zend.ZendStringReleaseEx(tmp_key, 0)
		}
	}
	zend.Efree(tmp)
	return zend.SUCCESS
}
func OnUpdateSessionHosts(
	entry *zend.ZendIniEntry,
	new_value *zend.ZendString,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	return PhpIniOnUpdateHosts(entry, new_value, mh_arg1, mh_arg2, mh_arg3, stage, 1)
}
func OnUpdateOutputHosts(
	entry *zend.ZendIniEntry,
	new_value *zend.ZendString,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	return PhpIniOnUpdateHosts(entry, new_value, mh_arg1, mh_arg2, mh_arg3, stage, 0)
}
func AppendModifiedUrl(url *zend.SmartStr, dest *zend.SmartStr, url_app *zend.SmartStr, separator *byte) {
	var url_parts *PhpUrl
	url.ZeroTail()
	url_parts = PhpUrlParseEx(url.GetS().GetVal(), url.GetS().GetLen())

	/* Ignore malformed URLs */

	if url_parts == nil {
		dest.AppendSmartStr(url)
		return
	}

	/* Don't modify URLs of the format "#mark" */

	if url_parts.GetFragment() != nil && '#' == url.GetS().GetVal()[0] {
		dest.AppendSmartStr(url)
		PhpUrlFree(url_parts)
		return
	}

	/* Check protocol. Only http/https is allowed. */

	if url_parts.GetScheme() != nil && !(zend.ZendStringEqualsLiteralCi(url_parts.GetScheme(), "http")) && !(zend.ZendStringEqualsLiteralCi(url_parts.GetScheme(), "https")) {
		dest.AppendSmartStr(url)
		PhpUrlFree(url_parts)
		return
	}

	/* Check host whitelist. If it's not listed, do nothing. */

	if url_parts.GetHost() != nil {
		var tmp *zend.ZendString = zend.ZendStringTolower(url_parts.GetHost())
		if zend.ZendHashExists(&(BG(url_adapt_session_hosts_ht)), tmp) == 0 {
			zend.ZendStringReleaseEx(tmp, 0)
			dest.AppendSmartStr(url)
			PhpUrlFree(url_parts)
			return
		}
		zend.ZendStringReleaseEx(tmp, 0)
	}

	/*
	 * When URL does not have path and query string add "/?".
	 * i.e. If URL is only "?foo=bar", should not add "/?".
	 */

	if url_parts.GetPath() == nil && url_parts.GetQuery() == nil && url_parts.GetFragment() == nil {

		/* URL is http://php.net or like */

		dest.AppendSmartStr(url)
		dest.AppendByte('/')
		dest.AppendByte('?')
		dest.AppendSmartStr(url_app)
		PhpUrlFree(url_parts)
		return
	}
	if url_parts.GetScheme() != nil {
		dest.AppendString(b.CastStrAuto(url_parts.GetScheme().GetVal()))
		dest.AppendString("://")
	} else if (*(url.GetS().GetVal())) == '/' && (*(url.GetS().GetVal() + 1)) == '/' {
		dest.AppendString("//")
	}
	if url_parts.GetUser() != nil {
		dest.AppendString(b.CastStrAuto(url_parts.GetUser().GetVal()))
		if url_parts.GetPass() != nil {
			dest.AppendString(b.CastStrAuto(url_parts.GetPass().GetVal()))
			dest.AppendByte(':')
		}
		dest.AppendByte('@')
	}
	if url_parts.GetHost() != nil {
		dest.AppendString(b.CastStrAuto(url_parts.GetHost().GetVal()))
	}
	if url_parts.GetPort() != 0 {
		dest.AppendByte(':')
		dest.AppendUlong(long(url_parts.GetPort()))
	}
	if url_parts.GetPath() != nil {
		dest.AppendString(b.CastStrAuto(url_parts.GetPath().GetVal()))
	}
	dest.AppendByte('?')
	if url_parts.GetQuery() != nil {
		dest.AppendString(b.CastStrAuto(url_parts.GetQuery().GetVal()))
		dest.AppendString(b.CastStrAuto(separator))
		dest.AppendSmartStr(url_app)
	} else {
		dest.AppendSmartStr(url_app)
	}
	if url_parts.GetFragment() != nil {
		dest.AppendByte('#')
		dest.AppendString(b.CastStrAuto(url_parts.GetFragment().GetVal()))
	}
	PhpUrlFree(url_parts)
}
func TagArg(ctx *UrlAdaptStateExT, quotes byte, type_ byte) {
	var f byte = 0

	/* arg.s is string WITHOUT NUL.
	   To avoid partial match, NUL is added here */

	ctx.GetArg().GetS().GetVal()[ctx.GetArg().GetS().GetLen()] = '0'
	if !(strcasecmp(ctx.GetArg().GetS().GetVal(), ctx.GetLookupData())) {
		f = 1
	}
	if quotes {
		ctx.GetResult().AppendByte(type_)
	}
	if f {
		AppendModifiedUrl(ctx.GetVal(), ctx.GetResult(), ctx.GetUrlApp(), core.PG(arg_separator).output)
	} else {
		ctx.GetResult().AppendSmartStr(ctx.GetVal())
	}
	if quotes {
		ctx.GetResult().AppendByte(type_)
	}
}
func Passthru(ctx *UrlAdaptStateExT, start *byte, YYCURSOR *byte) {
	ctx.GetResult().AppendString(b.CastStr(start, YYCURSOR-start))
}
func CheckHttpHost(target *byte) int {
	var host *zend.Zval
	var tmp *zend.Zval
	var host_tmp *zend.ZendString
	var colon *byte
	if b.Assign(&tmp, zend.EG__().GetSymbolTable().KeyFind(b.CastStr(zend.ZEND_STRL("_SERVER")))) && tmp.IsType(zend.IS_ARRAY) && b.Assign(&host, tmp.GetArr().KeyFind(b.CastStr(zend.ZEND_STRL("HTTP_HOST")))) && host.IsType(zend.IS_STRING) {
		host_tmp = zend.ZendStringInit(zend.Z_STRVAL_P(host), zend.Z_STRLEN_P(host), 0)

		/* HTTP_HOST could be 'localhost:8888' etc. */

		colon = strchr(host_tmp.GetVal(), ':')
		if colon != nil {
			host_tmp.SetLen(colon - host_tmp.GetVal())
			host_tmp.GetVal()[host_tmp.GetLen()] = '0'
		}
		if !(strcasecmp(host_tmp.GetVal(), target)) {
			zend.ZendStringReleaseEx(host_tmp, 0)
			return zend.SUCCESS
		}
		zend.ZendStringReleaseEx(host_tmp, 0)
	}
	return zend.FAILURE
}
func CheckHostWhitelist(ctx *UrlAdaptStateExT) int {
	var url_parts *PhpUrl = nil
	var allowed_hosts *zend.HashTable = b.CondF(ctx.GetType() != 0, func() *__auto__ { return &(BG(url_adapt_session_hosts_ht)) }, func() *__auto__ { return &(BG(url_adapt_output_hosts_ht)) })
	zend.ZEND_ASSERT(ctx.GetTagType() == TAG_FORM)
	if ctx.GetAttrVal().GetS() != nil && ctx.GetAttrVal().GetS().GetLen() != 0 {
		url_parts = PhpUrlParseEx(ctx.GetAttrVal().GetS().GetVal(), ctx.GetAttrVal().GetS().GetLen())
	} else {
		return zend.SUCCESS
	}
	if url_parts == nil {
		return zend.FAILURE
	}
	if url_parts.GetScheme() != nil {

		/* Only http/https should be handled.
		   A bit hacky check this here, but saves a URL parse. */

		if !(zend.ZendStringEqualsLiteralCi(url_parts.GetScheme(), "http")) && !(zend.ZendStringEqualsLiteralCi(url_parts.GetScheme(), "https")) {
			PhpUrlFree(url_parts)
			return zend.FAILURE
		}

		/* Only http/https should be handled.
		   A bit hacky check this here, but saves a URL parse. */

	}
	if url_parts.GetHost() == nil {
		PhpUrlFree(url_parts)
		return zend.SUCCESS
	}
	if !(allowed_hosts.GetNNumOfElements()) && CheckHttpHost(url_parts.GetHost().GetVal()) == zend.SUCCESS {
		PhpUrlFree(url_parts)
		return zend.SUCCESS
	}
	if allowed_hosts.KeyFind(url_parts.GetHost().GetStr()) == nil {
		PhpUrlFree(url_parts)
		return zend.FAILURE
	}
	PhpUrlFree(url_parts)
	return zend.SUCCESS
}
func HandleForm(ctx *UrlAdaptStateExT, start *byte, YYCURSOR *byte) {
	var doit int = 0
	if ctx.GetFormApp().GetS().GetLen() > 0 {
		switch ctx.GetTag().GetS().GetLen() {
		case b.SizeOf("\"form\"") - 1:
			if !(strncasecmp(ctx.GetTag().GetS().GetVal(), "form", ctx.GetTag().GetS().GetLen())) && CheckHostWhitelist(ctx) == zend.SUCCESS {
				doit = 1
			}
			break
		}
	}
	if doit != 0 {
		ctx.GetResult().AppendSmartStr(ctx.GetFormApp())
	}
}
func HandleTag(ctx *UrlAdaptStateExT, start *byte, YYCURSOR *byte) {
	var ok int = 0
	var i uint
	if ctx.GetTag().GetS() != nil {
		ctx.GetTag().GetS().GetLen() = 0
	}
	ctx.GetTag().AppendString(b.CastStr(start, YYCURSOR-start))
	for i = 0; i < ctx.GetTag().GetS().GetLen(); i++ {
		ctx.GetTag().GetS().GetVal()[i] = tolower(int(uint8(ctx.GetTag().GetS().GetVal()[i])))
	}

	/* intentionally using str_find here, in case the hash value is set, but the string val is changed later */

	if b.Assign(&(ctx.GetLookupData()), zend.ZendHashStrFindPtr(ctx.GetTags(), ctx.GetTag().GetS().GetVal(), ctx.GetTag().GetS().GetLen())) != nil {
		ok = 1
		if ctx.GetTag().GetS().GetLen() == b.SizeOf("\"form\"")-1 && !(strncasecmp(ctx.GetTag().GetS().GetVal(), "form", ctx.GetTag().GetS().GetLen())) {
			ctx.SetTagType(TAG_FORM)
		} else {
			ctx.SetTagType(TAG_NORMAL)
		}
	}
	if ok != 0 {
		ctx.SetState(STATE_NEXT_ARG)
	} else {
		ctx.SetState(STATE_PLAIN)
	}
}
func HandleArg(ctx *UrlAdaptStateExT, start *byte, YYCURSOR *byte) {
	if ctx.GetArg().GetS() != nil {
		ctx.GetArg().GetS().GetLen() = 0
	}
	ctx.GetArg().AppendString(b.CastStr(start, YYCURSOR-start))
	if ctx.GetTagType() == TAG_FORM && strncasecmp(ctx.GetArg().GetS().GetVal(), "action", ctx.GetArg().GetS().GetLen()) == 0 {
		ctx.SetAttrType(ATTR_ACTION)
	} else {
		ctx.SetAttrType(ATTR_NORMAL)
	}
}
func HandleVal(ctx *UrlAdaptStateExT, start *byte, YYCURSOR *byte, quotes byte, type_ byte) {
	ctx.GetVal().SetString(b.CastStr(start+quotes, YYCURSOR-start-quotes*2))
	if ctx.GetTagType() == TAG_FORM && ctx.GetAttrType() == ATTR_ACTION {
		ctx.GetAttrVal().SetString(b.CastStr(start+quotes, YYCURSOR-start-quotes*2))
	}
	TagArg(ctx, quotes, type_)
}
func XxMainloop(ctx *UrlAdaptStateExT, newdata *byte, newlen int) {
	var end *byte
	var q *byte
	var xp *byte
	var start *byte
	var rest int
	ctx.GetBuf().AppendString(b.CastStr(newdata, newlen))
	YYCURSOR = ctx.GetBuf().GetS().GetVal()
	YYLIMIT = ctx.GetBuf().GetS().GetVal() + ctx.GetBuf().GetS().GetLen()
	switch ctx.GetState() {
	case STATE_PLAIN:
		goto state_plain
	case STATE_TAG:
		goto state_tag
	case STATE_NEXT_ARG:
		goto state_next_arg
	case STATE_ARG:
		goto state_arg
	case STATE_BEFORE_VAL:
		goto state_before_val
	case STATE_VAL:
		goto state_val
	}
state_plain_begin:
	ctx.SetState(STATE_PLAIN)
state_plain:
	start = YYCURSOR
	var yych uint8
	var yybm []uint8 = []uint8{128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 0, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128}
	if YYLIMIT <= YYCURSOR {
		goto stop
	}
	yych = *YYCURSOR
	if (yybm[0+yych] & 128) != 0 {
		goto yy4
	}
	YYCURSOR++
	Passthru(ctx, start, xp)
	ctx.SetState(STATE_TAG)
	goto state_tag
yy4:
	YYCURSOR++
	if YYLIMIT <= YYCURSOR {
		goto stop
	}
	yych = *YYCURSOR
	if (yybm[0+yych] & 128) != 0 {
		goto yy4
	}
	Passthru(ctx, start, xp)
	goto state_plain
state_tag:
	start = YYCURSOR
	var yych uint8
	var yybm []uint8 = []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 0, 0, 0, 0, 0, 0, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if YYLIMIT-YYCURSOR < 2 {
		goto stop
	}
	yych = *YYCURSOR
	if yych <= '@' {
		if yych != ':' {
			goto yy11
		}
	} else {
		if yych <= 'Z' {
			goto yy9
		}
		if yych <= '`' {
			goto yy11
		}
		if yych >= '{' {
			goto yy11
		}
	}
yy9:
	YYCURSOR++
	yych = *YYCURSOR
	goto yy14
yy10:
	HandleTag(ctx, start, xp)
	Passthru(ctx, start, xp)
	if ctx.GetState() == STATE_PLAIN {
		goto state_plain
	} else {
		goto state_next_arg
	}
yy11:
	YYCURSOR++
	Passthru(ctx, start, xp)
	goto state_plain_begin
yy13:
	YYCURSOR++
	if YYLIMIT <= YYCURSOR {
		goto stop
	}
	yych = *YYCURSOR
yy14:
	if (yybm[0+yych] & 128) != 0 {
		goto yy13
	}
	goto yy10
state_next_arg_begin:
	ctx.SetState(STATE_NEXT_ARG)
state_next_arg:
	start = YYCURSOR
	var yych uint8
	var yybm []uint8 = []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 128, 128, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if YYLIMIT-YYCURSOR < 2 {
		goto stop
	}
	yych = *YYCURSOR
	if yych <= '.' {
		if yych <= 'f' {
			if yych <= 0x8 {
				goto yy25
			}
			if yych <= 'v' {
				goto yy21
			}
			goto yy25
		} else {
			if yych <= '\r' {
				goto yy21
			}
			if yych == ' ' {
				goto yy21
			}
			goto yy25
		}
	} else {
		if yych <= '@' {
			if yych <= '/' {
				goto yy17
			}
			if yych == '>' {
				goto yy19
			}
			goto yy25
		} else {
			if yych <= 'Z' {
				goto yy23
			}
			if yych <= '`' {
				goto yy25
			}
			if yych <= 'z' {
				goto yy23
			}
			goto yy25
		}
	}
yy17:
	YYCURSOR++
	if b.Assign(&yych, *YYCURSOR) == '>' {
		goto yy28
	}
yy18:
	Passthru(ctx, start, xp)
	goto state_plain_begin
yy19:
	YYCURSOR++
yy20:
	Passthru(ctx, start, xp)
	HandleForm(ctx, start, xp)
	goto state_plain_begin
yy21:
	YYCURSOR++
	yych = *YYCURSOR
	goto yy27
yy22:
	Passthru(ctx, start, xp)
	goto state_next_arg
yy23:
	YYCURSOR++
	YYCURSOR--
	ctx.SetState(STATE_ARG)
	goto state_arg
yy25:
	yych = *(b.PreInc(&YYCURSOR))
	goto yy18
yy26:
	YYCURSOR++
	if YYLIMIT <= YYCURSOR {
		goto stop
	}
	yych = *YYCURSOR
yy27:
	if (yybm[0+yych] & 128) != 0 {
		goto yy26
	}
	goto yy22
yy28:
	YYCURSOR++
	yych = *YYCURSOR
	goto yy20
state_arg:
	start = YYCURSOR
	var yych uint8
	var yybm []uint8 = []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 0, 0, 0, 0, 0, 0, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if YYLIMIT-YYCURSOR < 2 {
		goto stop
	}
	yych = *YYCURSOR
	if yych <= '@' {
		goto yy33
	}
	if yych <= 'Z' {
		goto yy31
	}
	if yych <= '`' {
		goto yy33
	}
	if yych >= '{' {
		goto yy33
	}
yy31:
	YYCURSOR++
	yych = *YYCURSOR
	goto yy36
yy32:
	Passthru(ctx, start, xp)
	HandleArg(ctx, start, xp)
	ctx.SetState(STATE_BEFORE_VAL)
	goto state_before_val
yy33:
	YYCURSOR++
	Passthru(ctx, start, xp)
	ctx.SetState(STATE_NEXT_ARG)
	goto state_next_arg
yy35:
	YYCURSOR++
	if YYLIMIT <= YYCURSOR {
		goto stop
	}
	yych = *YYCURSOR
yy36:
	if (yybm[0+yych] & 128) != 0 {
		goto yy35
	}
	goto yy32
state_before_val:
	start = YYCURSOR
	var yych uint8
	var yybm []uint8 = []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if YYLIMIT-YYCURSOR < 2 {
		goto stop
	}
	yych = *YYCURSOR
	if yych == ' ' {
		goto yy39
	}
	if yych == '=' {
		goto yy41
	}
	goto yy43
yy39:
	yych = *(b.Assign(&YYMARKER, b.PreInc(&YYCURSOR)))
	if yych == ' ' {
		goto yy46
	}
	if yych == '=' {
		goto yy44
	}
yy40:
	YYCURSOR--
	goto state_next_arg_begin
yy41:
	YYCURSOR++
	yych = *YYCURSOR
	goto yy45
yy42:
	Passthru(ctx, start, xp)
	ctx.SetState(STATE_VAL)
	goto state_val
yy43:
	yych = *(b.PreInc(&YYCURSOR))
	goto yy40
yy44:
	YYCURSOR++
	if YYLIMIT <= YYCURSOR {
		goto stop
	}
	yych = *YYCURSOR
yy45:
	if (yybm[0+yych] & 128) != 0 {
		goto yy44
	}
	goto yy42
yy46:
	YYCURSOR++
	if YYLIMIT <= YYCURSOR {
		goto stop
	}
	yych = *YYCURSOR
	if yych == ' ' {
		goto yy46
	}
	if yych == '=' {
		goto yy44
	}
	YYCURSOR = YYMARKER
	goto yy40
state_val:
	start = YYCURSOR
	var yych uint8
	var yybm []uint8 = []uint8{224, 224, 224, 224, 224, 224, 224, 224, 224, 192, 192, 224, 224, 192, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 192, 224, 64, 224, 224, 224, 224, 128, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 0, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224}
	if YYLIMIT-YYCURSOR < 2 {
		goto stop
	}
	yych = *YYCURSOR
	if yych <= ' ' {
		if yych <= 'f' {
			if yych <= 0x8 {
				goto yy54
			}
			if yych <= '\n' {
				goto yy56
			}
			goto yy54
		} else {
			if yych <= '\r' {
				goto yy56
			}
			if yych <= 0x1f {
				goto yy54
			}
			goto yy56
		}
	} else {
		if yych <= '&' {
			if yych != '"' {
				goto yy54
			}
		} else {
			if yych <= '\'' {
				goto yy53
			}
			if yych == '>' {
				goto yy56
			}
			goto yy54
		}
	}
	yych = *(b.Assign(&YYMARKER, b.PreInc(&YYCURSOR)))
	if yych != '>' {
		goto yy65
	}
yy52:
	Passthru(ctx, start, xp)
	goto state_next_arg_begin
yy53:
	yych = *(b.Assign(&YYMARKER, b.PreInc(&YYCURSOR)))
	if yych == '>' {
		goto yy52
	}
	goto yy60
yy54:
	YYCURSOR++
	yych = *YYCURSOR
	goto yy58
yy55:
	HandleVal(ctx, start, xp, 0, ' ')
	goto state_next_arg_begin
yy56:
	yych = *(b.PreInc(&YYCURSOR))
	goto yy52
yy57:
	YYCURSOR++
	if YYLIMIT <= YYCURSOR {
		goto stop
	}
	yych = *YYCURSOR
yy58:
	if (yybm[0+yych] & 32) != 0 {
		goto yy57
	}
	goto yy55
yy59:
	YYCURSOR++
	if YYLIMIT <= YYCURSOR {
		goto stop
	}
	yych = *YYCURSOR
yy60:
	if (yybm[0+yych] & 64) != 0 {
		goto yy59
	}
	if yych <= '=' {
		goto yy62
	}
yy61:
	YYCURSOR = YYMARKER
	goto yy52
yy62:
	YYCURSOR++
	HandleVal(ctx, start, xp, 1, '\'')
	goto state_next_arg_begin
yy64:
	YYCURSOR++
	if YYLIMIT <= YYCURSOR {
		goto stop
	}
	yych = *YYCURSOR
yy65:
	if (yybm[0+yych] & 128) != 0 {
		goto yy64
	}
	if yych >= '>' {
		goto yy61
	}
	YYCURSOR++
	HandleVal(ctx, start, xp, 1, '"')
	goto state_next_arg_begin
stop:
	if YYLIMIT < start {

		/* XXX: Crash avoidance. Need to work with reporter to figure out what goes wrong */

		rest = 0

		/* XXX: Crash avoidance. Need to work with reporter to figure out what goes wrong */

	} else {
		rest = YYLIMIT - start
	}
	if rest != 0 {
		memmove(ctx.GetBuf().GetS().GetVal(), start, rest)
	}
	ctx.GetBuf().GetS().GetLen() = rest
}
func PhpUrlScannerAdaptSingleUrl(
	url *byte,
	urllen int,
	name *byte,
	value *byte,
	newlen *int,
	encode int,
) *byte {
	var result *byte
	var surl zend.SmartStr = zend.MakeSmartStr(0)
	var buf zend.SmartStr = zend.MakeSmartStr(0)
	var url_app zend.SmartStr = zend.MakeSmartStr(0)
	var encoded *zend.ZendString
	surl.AppendString(b.CastStr(url, urllen))
	if encode != 0 {
		encoded = PhpRawUrlEncode(name, strlen(name))
		url_app.AppendString(b.CastStr(encoded.GetVal(), encoded.GetLen()))
		zend.ZendStringFree(encoded)
	} else {
		url_app.AppendString(b.CastStrAuto(name))
	}
	url_app.AppendByte('=')
	if encode != 0 {
		encoded = PhpRawUrlEncode(value, strlen(value))
		url_app.AppendString(b.CastStr(encoded.GetVal(), encoded.GetLen()))
		zend.ZendStringFree(encoded)
	} else {
		url_app.AppendString(b.CastStrAuto(value))
	}
	AppendModifiedUrl(&surl, &buf, &url_app, core.PG(arg_separator).output)
	buf.ZeroTail()
	if newlen != nil {
		*newlen = buf.GetS().GetLen()
	}
	result = zend.Estrndup(buf.GetS().GetVal(), buf.GetS().GetLen())
	url_app.Free()
	buf.Free()
	return result
}
func UrlAdaptExt(src *byte, srclen int, newlen *int, do_flush zend.ZendBool, ctx *UrlAdaptStateExT) *byte {
	var retval *byte
	XxMainloop(ctx, src, srclen)
	if ctx.GetResult().GetS() == nil {
		ctx.GetResult().AppendString("")
		*newlen = 0
	} else {
		*newlen = ctx.GetResult().GetS().GetLen()
	}
	ctx.GetResult().ZeroTail()
	if do_flush != 0 {
		ctx.GetResult().AppendString(ctx.GetBuf().GetS().GetStr())
		*newlen += ctx.GetBuf().GetS().GetLen()
		ctx.GetBuf().Free()
		ctx.GetVal().Free()
		ctx.GetAttrVal().Free()
	}
	retval = zend.Estrndup(ctx.GetResult().GetS().GetVal(), ctx.GetResult().GetS().GetLen())
	ctx.GetResult().Free()
	return retval
}
func PhpUrlScannerExActivate(type_ int) int {
	var ctx *UrlAdaptStateExT
	if type_ != 0 {
		ctx = &(BG(url_adapt_session_ex))
	} else {
		ctx = &(BG(url_adapt_output_ex))
	}
	memset(ctx, 0, zend_long((*byte)(&((*UrlAdaptStateExT)(nil).GetTags()))-(*byte)(nil)))
	return zend.SUCCESS
}
func PhpUrlScannerExDeactivate(type_ int) int {
	var ctx *UrlAdaptStateExT
	if type_ != 0 {
		ctx = &(BG(url_adapt_session_ex))
	} else {
		ctx = &(BG(url_adapt_output_ex))
	}
	ctx.GetResult().Free()
	ctx.GetBuf().Free()
	ctx.GetTag().Free()
	ctx.GetArg().Free()
	ctx.GetAttrVal().Free()
	return zend.SUCCESS
}
func PhpUrlScannerSessionHandlerImpl(
	output *byte,
	output_len int,
	handled_output **byte,
	handled_output_len *int,
	mode int,
	type_ int,
) {
	var len_ int
	var url_state *UrlAdaptStateExT
	if type_ != 0 {
		url_state = &(BG(url_adapt_session_ex))
	} else {
		url_state = &(BG(url_adapt_output_ex))
	}
	if url_state.GetUrlApp().GetS().GetLen() != 0 {
		*handled_output = UrlAdaptExt(output, output_len, &len_, zend_bool(b.Cond((mode&(core.PHP_OUTPUT_HANDLER_END|core.PHP_OUTPUT_HANDLER_CONT|core.PHP_OUTPUT_HANDLER_FLUSH|core.PHP_OUTPUT_HANDLER_FINAL)) != 0, 1, 0)), url_state)
		if b.SizeOf("unsigned int") < b.SizeOf("size_t") {
			if len_ > UINT_MAX {
				len_ = UINT_MAX
			}
		}
		*handled_output_len = len_
	} else if url_state.GetUrlApp().GetS().GetLen() == 0 {
		var ctx *UrlAdaptStateExT = url_state
		if ctx.GetBuf().GetS() != nil && ctx.GetBuf().GetS().GetLen() != 0 {
			ctx.GetResult().AppendString(ctx.GetBuf().GetS().GetStr())
			ctx.GetResult().AppendString(b.CastStr(output, output_len))
			*handled_output = zend.Estrndup(ctx.GetResult().GetS().GetVal(), ctx.GetResult().GetS().GetLen())
			*handled_output_len = ctx.GetBuf().GetS().GetLen() + output_len
			ctx.GetBuf().Free()
			ctx.GetResult().Free()
		} else {
			*handled_output = zend.Estrndup(output, b.Assign(&(*handled_output_len), output_len))
		}
	} else {
		*handled_output = nil
	}
}
func PhpUrlScannerSessionHandler(output *byte, output_len int, handled_output **byte, handled_output_len *int, mode int) {
	PhpUrlScannerSessionHandlerImpl(output, output_len, handled_output, handled_output_len, mode, 1)
}
func PhpUrlScannerOutputHandler(output *byte, output_len int, handled_output **byte, handled_output_len *int, mode int) {
	PhpUrlScannerSessionHandlerImpl(output, output_len, handled_output, handled_output_len, mode, 0)
}
func PhpUrlScannerAddVarImpl(
	name *byte,
	name_len int,
	value *byte,
	value_len int,
	encode int,
	type_ int,
) int {
	var sname zend.SmartStr = zend.MakeSmartStr(0)
	var svalue zend.SmartStr = zend.MakeSmartStr(0)
	var hname zend.SmartStr = zend.MakeSmartStr(0)
	var hvalue zend.SmartStr = zend.MakeSmartStr(0)
	var encoded *zend.ZendString
	var url_state *UrlAdaptStateExT
	var handler core.PhpOutputHandlerFuncT
	if type_ != 0 {
		url_state = &(BG(url_adapt_session_ex))
		handler = PhpUrlScannerSessionHandler
	} else {
		url_state = &(BG(url_adapt_output_ex))
		handler = PhpUrlScannerOutputHandler
	}
	if url_state.GetActive() == 0 {
		PhpUrlScannerExActivate(type_)
		core.PhpOutputStartInternal(zend.ZEND_STRL("URL-Rewriter"), handler, 0, core.PHP_OUTPUT_HANDLER_STDFLAGS)
		url_state.SetActive(1)
	}
	if url_state.GetUrlApp().GetS() != nil && url_state.GetUrlApp().GetS().GetLen() != 0 {
		url_state.GetUrlApp().AppendString(b.CastStrAuto(core.PG(arg_separator).output))
	}
	if encode != 0 {
		encoded = PhpRawUrlEncode(name, name_len)
		sname.AppendString(b.CastStr(encoded.GetVal(), encoded.GetLen()))
		zend.ZendStringFree(encoded)
		encoded = PhpRawUrlEncode(value, value_len)
		svalue.AppendString(b.CastStr(encoded.GetVal(), encoded.GetLen()))
		zend.ZendStringFree(encoded)
		encoded = PhpEscapeHtmlEntitiesEx((*uint8)(name), name_len, 0, ENT_QUOTES|ENT_SUBSTITUTE, core.SG(default_charset), 0)
		hname.AppendString(b.CastStr(encoded.GetVal(), encoded.GetLen()))
		zend.ZendStringFree(encoded)
		encoded = PhpEscapeHtmlEntitiesEx((*uint8)(value), value_len, 0, ENT_QUOTES|ENT_SUBSTITUTE, core.SG(default_charset), 0)
		hvalue.AppendString(b.CastStr(encoded.GetVal(), encoded.GetLen()))
		zend.ZendStringFree(encoded)
	} else {
		sname.AppendString(b.CastStr(name, name_len))
		svalue.AppendString(b.CastStr(value, value_len))
		hname.AppendString(b.CastStr(name, name_len))
		hvalue.AppendString(b.CastStr(value, value_len))
	}
	url_state.GetUrlApp().AppendSmartStr(&sname)
	url_state.GetUrlApp().AppendByte('=')
	url_state.GetUrlApp().AppendSmartStr(&svalue)
	url_state.GetFormApp().AppendString("<input type=\"hidden\" name=\"")
	url_state.GetFormApp().AppendSmartStr(&hname)
	url_state.GetFormApp().AppendString("\" value=\"")
	url_state.GetFormApp().AppendSmartStr(&hvalue)
	url_state.GetFormApp().AppendString("\" />")
	sname.Free()
	svalue.Free()
	hname.Free()
	hvalue.Free()
	return zend.SUCCESS
}
func PhpUrlScannerAddSessionVar(name *byte, name_len int, value *byte, value_len int, encode int) int {
	return PhpUrlScannerAddVarImpl(name, name_len, value, value_len, encode, 1)
}
func PhpUrlScannerAddVar(name *byte, name_len int, value *byte, value_len int, encode int) int {
	return PhpUrlScannerAddVarImpl(name, name_len, value, value_len, encode, 0)
}
func PhpUrlScannerResetVarsImpl(type_ int) {
	var url_state *UrlAdaptStateExT
	if type_ != 0 {
		url_state = &(BG(url_adapt_session_ex))
	} else {
		url_state = &(BG(url_adapt_output_ex))
	}
	if url_state.GetFormApp().GetS() != nil {
		url_state.GetFormApp().GetS().GetLen() = 0
	}
	if url_state.GetUrlApp().GetS() != nil {
		url_state.GetUrlApp().GetS().GetLen() = 0
	}
}
func PhpUrlScannerResetSessionVars() int {
	PhpUrlScannerResetVarsImpl(1)
	return zend.SUCCESS
}
func PhpUrlScannerResetVars() int {
	PhpUrlScannerResetVarsImpl(0)
	return zend.SUCCESS
}
func PhpUrlScannerResetVarImpl(name *zend.ZendString, encode int, type_ int) int {
	var start *byte
	var end *byte
	var limit *byte
	var separator_len int
	var sname zend.SmartStr = zend.MakeSmartStr(0)
	var hname zend.SmartStr = zend.MakeSmartStr(0)
	var url_app zend.SmartStr = zend.MakeSmartStr(0)
	var form_app zend.SmartStr = zend.MakeSmartStr(0)
	var encoded *zend.ZendString
	var ret int = zend.SUCCESS
	var sep_removed zend.ZendBool = 0
	var url_state *UrlAdaptStateExT
	if type_ != 0 {
		url_state = &(BG(url_adapt_session_ex))
	} else {
		url_state = &(BG(url_adapt_output_ex))
	}

	/* Short circuit check. Only check url_app. */

	if url_state.GetUrlApp().GetS() == nil || url_state.GetUrlApp().GetS().GetLen() == 0 {
		return zend.SUCCESS
	}
	if encode != 0 {
		encoded = PhpRawUrlEncode(name.GetVal(), name.GetLen())
		sname.AppendString(b.CastStr(encoded.GetVal(), encoded.GetLen()))
		zend.ZendStringFree(encoded)
		encoded = PhpEscapeHtmlEntitiesEx((*uint8)(name.GetVal()), name.GetLen(), 0, ENT_QUOTES|ENT_SUBSTITUTE, core.SG(default_charset), 0)
		hname.AppendString(b.CastStr(encoded.GetVal(), encoded.GetLen()))
		zend.ZendStringFree(encoded)
	} else {
		sname.AppendString(b.CastStr(name.GetVal(), name.GetLen()))
		hname.AppendString(b.CastStr(name.GetVal(), name.GetLen()))
	}
	sname.ZeroTail()
	hname.ZeroTail()
	url_app.AppendSmartStr(&sname)
	url_app.AppendByte('=')
	url_app.ZeroTail()
	form_app.AppendString("<input type=\"hidden\" name=\"")
	form_app.AppendSmartStr(&hname)
	form_app.AppendString("\" value=\"")
	form_app.ZeroTail()

	/* Short circuit check. Only check url_app. */

	start = (*byte)(core.PhpMemnstr(url_state.GetUrlApp().GetS().GetVal(), url_app.GetS().GetVal(), url_app.GetS().GetLen(), url_state.GetUrlApp().GetS().GetVal()+url_state.GetUrlApp().GetS().GetLen()))
	if start == nil {
		ret = zend.FAILURE
		goto finish
	}

	/* Get end of url var */

	limit = url_state.GetUrlApp().GetS().GetVal() + url_state.GetUrlApp().GetS().GetLen()
	end = start + url_app.GetS().GetLen()
	separator_len = strlen(core.PG(arg_separator).output)
	for end < limit {
		if !(memcmp(end, core.PG(arg_separator).output, separator_len)) {
			end += separator_len
			sep_removed = 1
			break
		}
		end++
	}

	/* Remove all when this is the only rewrite var */

	if url_state.GetUrlApp().GetS().GetLen() == end-start {
		PhpUrlScannerResetVarsImpl(type_)
		goto finish
	}

	/* Check preceding separator */

	if sep_removed == 0 && size_t(start-core.PG(arg_separator).output) >= separator_len && !(memcmp(start-separator_len, core.PG(arg_separator).output, separator_len)) {
		start -= separator_len
	}

	/* Remove partially */

	memmove(start, end, url_state.GetUrlApp().GetS().GetLen()-(end-url_state.GetUrlApp().GetS().GetVal()))
	url_state.GetUrlApp().GetS().GetLen() -= end - start
	url_state.GetUrlApp().GetS().GetVal()[url_state.GetUrlApp().GetS().GetLen()] = '0'

	/* Remove form var */

	start = (*byte)(core.PhpMemnstr(url_state.GetFormApp().GetS().GetVal(), form_app.GetS().GetVal(), form_app.GetS().GetLen(), url_state.GetFormApp().GetS().GetVal()+url_state.GetFormApp().GetS().GetLen()))
	if start == nil {

		/* Should not happen */

		ret = zend.FAILURE
		PhpUrlScannerResetVarsImpl(type_)
		goto finish
	}

	/* Get end of form var */

	limit = url_state.GetFormApp().GetS().GetVal() + url_state.GetFormApp().GetS().GetLen()
	end = start + form_app.GetS().GetLen()
	for end < limit {
		if (*end) == '>' {
			end += 1
			break
		}
		end++
	}

	/* Remove partially */

	memmove(start, end, url_state.GetFormApp().GetS().GetLen()-(end-url_state.GetFormApp().GetS().GetVal()))
	url_state.GetFormApp().GetS().GetLen() -= end - start
	url_state.GetFormApp().GetS().GetVal()[url_state.GetFormApp().GetS().GetLen()] = '0'
finish:
	url_app.Free()
	form_app.Free()
	sname.Free()
	hname.Free()
	return ret
}
func PhpUrlScannerResetSessionVar(name *zend.ZendString, encode int) int {
	return PhpUrlScannerResetVarImpl(name, encode, 1)
}
func PhpUrlScannerResetVar(name *zend.ZendString, encode int) int {
	return PhpUrlScannerResetVarImpl(name, encode, 0)
}
func ZmStartupUrlScanner(type_ int, module_number int) int {
	zend.REGISTER_INI_ENTRIES()
	return zend.SUCCESS
}
func ZmShutdownUrlScanner(type_ int, module_number int) int {
	zend.UNREGISTER_INI_ENTRIES()
	return zend.SUCCESS
}
func ZmActivateUrlScanner(type_ int, module_number int) int {
	BG(url_adapt_session_ex).active = 0
	BG(url_adapt_session_ex).tag_type = 0
	BG(url_adapt_session_ex).attr_type = 0
	BG(url_adapt_output_ex).active = 0
	BG(url_adapt_output_ex).tag_type = 0
	BG(url_adapt_output_ex).attr_type = 0
	return zend.SUCCESS
}
func ZmDeactivateUrlScanner(type_ int, module_number int) int {
	if BG(url_adapt_session_ex).active {
		PhpUrlScannerExDeactivate(1)
		BG(url_adapt_session_ex).active = 0
		BG(url_adapt_session_ex).tag_type = 0
		BG(url_adapt_session_ex).attr_type = 0
	}
	BG(url_adapt_session_ex).form_app.Free()
	BG(url_adapt_session_ex).url_app.Free()
	if BG(url_adapt_output_ex).active {
		PhpUrlScannerExDeactivate(0)
		BG(url_adapt_output_ex).active = 0
		BG(url_adapt_output_ex).tag_type = 0
		BG(url_adapt_output_ex).attr_type = 0
	}
	BG(url_adapt_output_ex).form_app.Free()
	BG(url_adapt_output_ex).url_app.Free()
	return zend.SUCCESS
}
