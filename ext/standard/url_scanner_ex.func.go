package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"strings"
)

func PhpIniOnUpdateTags(
	entry *zend.ZendIniEntry,
	new_value *types.String,
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
		ctx = &(BG__().url_adapt_session_ex)
	} else {
		ctx = &(BG__().url_adapt_output_ex)
	}
	tmp = zend.Estrndup(new_value.GetVal(), new_value.GetLen())
	if ctx.GetTags() != nil {
		ctx.GetTags().Destroy()
	}
	ctx.SetTags(types.NewArray())
	for key = core.PhpStrtokR(tmp, ",", &lasts); key != nil; key = core.PhpStrtokR(nil, ",", &lasts) {
		var val *byte
		val = strchr(key, '=')
		if val != nil {
			var q *byte
			var keylen int
			var str *types.String
			lang.PostInc(&(*val)) = '0'
			for q = key; *q; q++ {
				*q = tolower(*q)
			}
			keylen = q - key
			str = types.NewString(b.CastStr(key, keylen))
			//types.GC_MAKE_PERSISTENT_LOCAL(str)
			types.ZendHashAddMem(ctx.GetTags(), str.GetStr(), val, strlen(val)+1)
			// types.ZendStringReleaseEx(str, 1)
		}
	}
	zend.Efree(tmp)
	return types.SUCCESS
}
func OnUpdateSessionTags(
	entry *zend.ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	return PhpIniOnUpdateTags(entry, new_value, mh_arg1, mh_arg2, mh_arg3, stage, 1)
}
func OnUpdateOutputTags(
	entry *zend.ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	return PhpIniOnUpdateTags(entry, new_value, mh_arg1, mh_arg2, mh_arg3, stage, 0)
}
func PhpIniOnUpdateHosts(
	entry *zend.ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
	type_ int,
) int {
	var hosts *types.Array
	var key *byte
	var tmp *byte
	var lasts *byte = nil
	if type_ != 0 {
		hosts = BG__().url_adapt_session_hosts_ht
	} else {
		hosts = BG__().url_adapt_output_hosts_ht
	}
	hosts.Clean()

	/* Use user supplied host whitelist */

	tmp = zend.Estrndup(new_value.GetVal(), new_value.GetLen())
	for key = core.PhpStrtokR(tmp, ",", &lasts); key != nil; key = core.PhpStrtokR(nil, ",", &lasts) {
		var keylen int
		var tmp_key *types.String
		var q *byte
		for q = key; *q; q++ {
			*q = tolower(*q)
		}
		keylen = q - key
		if keylen > 0 {
			tmp_key = types.NewString(b.CastStr(key, keylen))
			types.ZendHashAddEmptyElement(hosts, tmp_key.GetStr())
			// types.ZendStringReleaseEx(tmp_key, 0)
		}
	}
	zend.Efree(tmp)
	return types.SUCCESS
}
func OnUpdateSessionHosts(
	entry *zend.ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	return PhpIniOnUpdateHosts(entry, new_value, mh_arg1, mh_arg2, mh_arg3, stage, 1)
}
func OnUpdateOutputHosts(
	entry *zend.ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	return PhpIniOnUpdateHosts(entry, new_value, mh_arg1, mh_arg2, mh_arg3, stage, 0)
}
func AppendModifiedUrl(url *zend.SmartStr, dest *zend.SmartStr, url_app *zend.SmartStr, separator *byte) {
	var url_parts *PhpUrl
	//url.ZeroTail()
	url_parts = PhpUrlParseEx(url.GetS().GetVal(), url.GetS().GetLen())

	/* Ignore malformed URLs */

	if url_parts == nil {
		dest.AppendSmartStr(url)
		return
	}

	/* Don't modify URLs of the format "#mark" */
	if fragment := url_parts.Fragment(); len(fragment) > 0 && fragment[0] == '#' {
		dest.AppendSmartStr(url)
		return
	}

	/* Check protocol. Only http/https is allowed. */
	if lcScheme := ascii.StrToLower(url_parts.Scheme()); lcScheme == "http" || lcScheme == "https" {
		dest.AppendSmartStr(url)
		return
	}

	/* Check host whitelist. If it's not listed, do nothing. */
	if url_parts.HasHost() {
		var tmp = ascii.StrToLower(url_parts.Host())
		if !BG__().url_adapt_session_hosts_ht.KeyExists(tmp) {
			dest.AppendSmartStr(url)
			return
		}
	}

	/*
	 * When URL does not have path and query string add "/?".
	 * i.e. If URL is only "?foo=bar", should not add "/?".
	 */
	if !url_parts.HasPath() && !url_parts.HasQuery() && !url_parts.HasFragment() {
		/* URL is http://php.net or like */
		dest.AppendSmartStr(url)
		dest.WriteByte('/')
		dest.WriteByte('?')
		dest.AppendSmartStr(url_app)
		return
	}
	if url_parts.HasScheme() {
		dest.WriteString(url_parts.Scheme())
		dest.WriteString("://")
	} else if (*(url.GetS().GetVal())) == '/' && (*(url.GetS().GetVal() + 1)) == '/' {
		dest.WriteString("//")
	}
	if url_parts.HasUser() {
		dest.WriteString(url_parts.User())
		if url_parts.HasPass() {
			dest.WriteString(url_parts.Pass())
			dest.WriteByte(':')
		}
		dest.WriteByte('@')
	}
	if url_parts.HasHost() {
		dest.WriteString(url_parts.Host())
	}
	if url_parts.Port() != 0 {
		dest.WriteByte(':')
		dest.WriteUlong(uint(url_parts.Port()))
	}
	if url_parts.HasPath() {
		dest.WriteString(url_parts.Path())
	}
	dest.WriteByte('?')
	if url_parts.HasQuery() {
		dest.WriteString(url_parts.Query())
		dest.WriteString(b.CastStrAuto(separator))
		dest.AppendSmartStr(url_app)
	} else {
		dest.AppendSmartStr(url_app)
	}
	if url_parts.HasFragment() {
		dest.WriteByte('#')
		dest.WriteString(url_parts.Fragment())
	}
}
func TagArg(ctx *UrlAdaptStateExT, quotes byte, type_ byte) {
	var f byte = 0

	/* arg.s is string WITHOUT NUL.
	   To avoid partial match, NUL is added here */

	ctx.GetArg().GetS().GetStr()[ctx.GetArg().GetS().GetLen()] = '0'
	if !(strcasecmp(ctx.GetArg().GetS().GetVal(), ctx.GetLookupData())) {
		f = 1
	}
	if quotes {
		ctx.GetResult().WriteByte(type_)
	}
	if f {
		AppendModifiedUrl(ctx.GetVal(), ctx.GetResult(), ctx.GetUrlApp(), core.PG__().arg_separator.output)
	} else {
		ctx.GetResult().AppendSmartStr(ctx.GetVal())
	}
	if quotes {
		ctx.GetResult().WriteByte(type_)
	}
}
func Passthru(ctx *UrlAdaptStateExT, start *byte, YYCURSOR *byte) {
	ctx.GetResult().WriteString(b.CastStr(start, YYCURSOR-start))
}
func CheckHttpHost(target string) bool {
	tmp := zend.EG__().GetSymbolTable().KeyFind("_SERVER")
	if tmp != nil && tmp.IsArray() {
		host := tmp.Array().KeyFind("HTTP_HOST")
		if host != nil && host.IsString() {
			hostStr := host.String()

			/* HTTP_HOST could be 'localhost:8888' etc. */
			if pos := strings.IndexByte(hostStr, ':'); pos >= 0 {
				hostStr = hostStr[:pos]
			}
			if ascii.StrCaseEquals(hostStr, target) {
				return true
			}
		}
	}

	return false
}
func CheckHostWhitelist(ctx *UrlAdaptStateExT) int {
	var url_parts *PhpUrl = nil
	var allowed_hosts *types.Array = lang.CondF(ctx.GetType() != 0, func() *types.Array { return BG__().url_adapt_session_hosts_ht }, func() *types.Array { return BG__().url_adapt_output_hosts_ht })
	b.Assert(ctx.GetTagType() == TAG_FORM)
	if ctx.GetAttrVal().GetS() != nil && ctx.GetAttrVal().GetS().GetLen() != 0 {
		url_parts = PhpUrlParseEx(ctx.GetAttrVal().GetS().GetVal(), ctx.GetAttrVal().GetS().GetLen())
	} else {
		return types.SUCCESS
	}
	if url_parts == nil {
		return types.FAILURE
	}
	if url_parts.HasScheme() {
		/* Only http/https should be handled.
		   A bit hacky check this here, but saves a URL parse. */
		if lcScheme := ascii.StrToLower(url_parts.Scheme()); lcScheme == "http" || lcScheme == "https" {
			return types.FAILURE
		}
	}
	if url_parts.HasHost() {
		return types.SUCCESS
	}
	if allowed_hosts.Len() == 0 && CheckHttpHost(url_parts.Host()) {
		return types.SUCCESS
	}
	if allowed_hosts.KeyFind(url_parts.Host()) == nil {
		return types.FAILURE
	}
	return types.SUCCESS
}
func HandleForm(ctx *UrlAdaptStateExT, start *byte, YYCURSOR *byte) {
	var doit int = 0
	if ctx.GetFormApp().GetS().GetLen() > 0 {
		switch ctx.GetTag().GetS().GetLen() {
		case b.SizeOf("\"form\"") - 1:
			if !(strncasecmp(ctx.GetTag().GetS().GetVal(), "form", ctx.GetTag().GetS().GetLen())) && CheckHostWhitelist(ctx) == types.SUCCESS {
				doit = 1
			}
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
	ctx.GetTag().WriteString(b.CastStr(start, YYCURSOR-start))
	for i = 0; i < ctx.GetTag().GetS().GetLen(); i++ {
		ctx.GetTag().GetS().GetStr()[i] = tolower(int(uint8(ctx.GetTag().GetS().GetStr()[i])))
	}

	/* intentionally using str_find here, in case the hash value is set, but the string val is changed later */

	if lang.Assign(&(ctx.GetLookupData()), types.ZendHashStrFindPtr(ctx.GetTags(), ctx.GetTag().GetS().GetStr())) != nil {
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
	ctx.GetArg().WriteString(b.CastStr(start, YYCURSOR-start))
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
	ctx.GetBuf().WriteString(b.CastStr(newdata, newlen))
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
	if lang.Assign(&yych, *YYCURSOR) == '>' {
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
	yych = *(lang.PreInc(&YYCURSOR))
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
	yych = *(lang.Assign(&YYMARKER, lang.PreInc(&YYCURSOR)))
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
	yych = *(lang.PreInc(&YYCURSOR))
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
	yych = *(lang.Assign(&YYMARKER, lang.PreInc(&YYCURSOR)))
	if yych != '>' {
		goto yy65
	}
yy52:
	Passthru(ctx, start, xp)
	goto state_next_arg_begin
yy53:
	yych = *(lang.Assign(&YYMARKER, lang.PreInc(&YYCURSOR)))
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
	yych = *(lang.PreInc(&YYCURSOR))
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
	var surl zend.SmartStr
	var buf zend.SmartStr
	var url_app zend.SmartStr
	var encoded string
	surl.WriteString(b.CastStr(url, urllen))
	if encode != 0 {
		encoded = PhpRawUrlEncode(name)
		url_app.WriteString(encoded)
	} else {
		url_app.WriteString(b.CastStrAuto(name))
	}
	url_app.WriteByte('=')
	if encode != 0 {
		encoded = PhpRawUrlEncode(value)
		url_app.WriteString(encoded)
	} else {
		url_app.WriteString(b.CastStrAuto(value))
	}
	AppendModifiedUrl(&surl, &buf, &url_app, core.PG__().arg_separator.output)
	//buf.ZeroTail()
	if newlen != nil {
		*newlen = buf.GetS().GetLen()
	}
	result = zend.Estrndup(buf.GetS().GetVal(), buf.GetS().GetLen())
	url_app.Free()
	buf.Free()
	return result
}
func UrlAdaptExt(src *byte, srclen int, newlen *int, do_flush bool, ctx *UrlAdaptStateExT) *byte {
	var retval *byte
	XxMainloop(ctx, src, srclen)
	if ctx.GetResult().GetS() == nil {
		ctx.GetResult().WriteString("")
		*newlen = 0
	} else {
		*newlen = ctx.GetResult().GetS().GetLen()
	}
	//ctx.GetResult().ZeroTail()
	if do_flush != 0 {
		ctx.GetResult().WriteString(ctx.GetBuf().GetS().GetStr())
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
		ctx = &(BG__().url_adapt_session_ex)
	} else {
		ctx = &(BG__().url_adapt_output_ex)
	}
	memset(ctx, 0, zend_long((*byte)(&((*UrlAdaptStateExT)(nil).GetTags()))-(*byte)(nil)))
	return types.SUCCESS
}
func PhpUrlScannerExDeactivate(type_ int) int {
	var ctx *UrlAdaptStateExT
	if type_ != 0 {
		ctx = &(BG__().url_adapt_session_ex)
	} else {
		ctx = &(BG__().url_adapt_output_ex)
	}
	ctx.GetResult().Free()
	ctx.GetBuf().Free()
	ctx.GetTag().Free()
	ctx.GetArg().Free()
	ctx.GetAttrVal().Free()
	return types.SUCCESS
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
		url_state = &(BG__().url_adapt_session_ex)
	} else {
		url_state = &(BG__().url_adapt_output_ex)
	}
	if url_state.GetUrlApp().GetS().GetLen() != 0 {
		*handled_output = UrlAdaptExt(output, output_len, &len_, zend_bool(lang.Cond((mode&(core.PHP_OUTPUT_HANDLER_END|core.PHP_OUTPUT_HANDLER_CONT|core.PHP_OUTPUT_HANDLER_FLUSH|core.PHP_OUTPUT_HANDLER_FINAL)) != 0, 1, 0)), url_state)
		if b.SizeOf("unsigned int") < b.SizeOf("size_t") {
			if len_ > UINT_MAX {
				len_ = UINT_MAX
			}
		}
		*handled_output_len = len_
	} else if url_state.GetUrlApp().GetS().GetLen() == 0 {
		var ctx *UrlAdaptStateExT = url_state
		if ctx.GetBuf().GetS() != nil && ctx.GetBuf().GetS().GetLen() != 0 {
			ctx.GetResult().WriteString(ctx.GetBuf().GetS().GetStr())
			ctx.GetResult().WriteString(b.CastStr(output, output_len))
			*handled_output = zend.Estrndup(ctx.GetResult().GetS().GetVal(), ctx.GetResult().GetS().GetLen())
			*handled_output_len = ctx.GetBuf().GetS().GetLen() + output_len
			ctx.GetBuf().Free()
			ctx.GetResult().Free()
		} else {
			*handled_output = zend.Estrndup(output, lang.Assign(&(*handled_output_len), output_len))
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
	var sname zend.SmartStr
	var svalue zend.SmartStr
	var hname zend.SmartStr
	var hvalue zend.SmartStr
	var url_state *UrlAdaptStateExT
	var handler core.PhpOutputHandlerFuncT
	if type_ != 0 {
		url_state = &(BG__().url_adapt_session_ex)
		handler = PhpUrlScannerSessionHandler
	} else {
		url_state = &(BG__().url_adapt_output_ex)
		handler = PhpUrlScannerOutputHandler
	}
	if url_state.GetActive() == 0 {
		PhpUrlScannerExActivate(type_)
		core.PhpOutputStartInternal("URL-Rewriter", handler, 0, core.PHP_OUTPUT_HANDLER_STDFLAGS)
		url_state.SetActive(1)
	}
	if url_state.GetUrlApp().GetS() != nil && url_state.GetUrlApp().GetS().GetLen() != 0 {
		url_state.GetUrlApp().WriteString(b.CastStrAuto(core.PG__().arg_separator.output))
	}
	if encode != 0 {
		encoded := PhpRawUrlEncode(b.CastStr(name, name_len))
		sname.WriteString(encoded)
		//types.ZendStringFree(encoded)
		encoded = PhpRawUrlEncode(b.CastStr(value, value_len))
		svalue.WriteString(encoded)
		//types.ZendStringFree(encoded)
		encoded = PhpEscapeHtmlEntitiesEx((*uint8)(name), name_len, 0, ENT_QUOTES|ENT_SUBSTITUTE, core.SG__().default_charset, 0).GetStr()
		hname.WriteString(encoded)
		//types.ZendStringFree(encoded)
		encoded = PhpEscapeHtmlEntitiesEx((*uint8)(value), value_len, 0, ENT_QUOTES|ENT_SUBSTITUTE, core.SG__().default_charset, 0).GetStr()
		hvalue.WriteString(encoded)
		//types.ZendStringFree(encoded)
	} else {
		sname.WriteString(b.CastStr(name, name_len))
		svalue.WriteString(b.CastStr(value, value_len))
		hname.WriteString(b.CastStr(name, name_len))
		hvalue.WriteString(b.CastStr(value, value_len))
	}
	url_state.GetUrlApp().AppendSmartStr(&sname)
	url_state.GetUrlApp().WriteByte('=')
	url_state.GetUrlApp().AppendSmartStr(&svalue)
	url_state.GetFormApp().WriteString("<input type=\"hidden\" name=\"")
	url_state.GetFormApp().AppendSmartStr(&hname)
	url_state.GetFormApp().WriteString("\" value=\"")
	url_state.GetFormApp().AppendSmartStr(&hvalue)
	url_state.GetFormApp().WriteString("\" />")
	sname.Free()
	svalue.Free()
	hname.Free()
	hvalue.Free()
	return types.SUCCESS
}
func PhpUrlScannerAddVar(name string, value string, encode int) bool {
	return PhpUrlScannerAddVarImpl(b.CastStrPtr(name), len(name), b.CastStrPtr(value), len(value), encode, 0) == types.SUCCESS
}
func PhpUrlScannerResetVarsImpl(type_ int) {
	var url_state *UrlAdaptStateExT
	if type_ != 0 {
		url_state = &(BG__().url_adapt_session_ex)
	} else {
		url_state = &(BG__().url_adapt_output_ex)
	}
	if url_state.GetFormApp().GetS() != nil {
		url_state.GetFormApp().GetS().GetLen() = 0
	}
	if url_state.GetUrlApp().GetS() != nil {
		url_state.GetUrlApp().GetS().GetLen() = 0
	}
}
func PhpUrlScannerResetVars() {
	PhpUrlScannerResetVarsImpl(0)
}
func PhpUrlScannerResetVarImpl(name *types.String, encode int, type_ int) int {
	var start *byte
	var end *byte
	var limit *byte
	var separator_len int
	var sname zend.SmartStr
	var hname zend.SmartStr
	var url_app zend.SmartStr
	var form_app zend.SmartStr
	var ret int = types.SUCCESS
	var sep_removed bool = 0
	var url_state *UrlAdaptStateExT
	if type_ != 0 {
		url_state = &(BG__().url_adapt_session_ex)
	} else {
		url_state = &(BG__().url_adapt_output_ex)
	}

	/* Short circuit check. Only check url_app. */

	if url_state.GetUrlApp().GetS() == nil || url_state.GetUrlApp().GetS().GetLen() == 0 {
		return types.SUCCESS
	}
	if encode != 0 {
		encoded := PhpRawUrlEncode(name.GetStr())
		sname.WriteString(encoded)
		encoded = PhpEscapeHtmlEntitiesEx((*uint8)(name.GetVal()), name.GetLen(), 0, ENT_QUOTES|ENT_SUBSTITUTE, core.SG__().default_charset, 0).GetStr()
		hname.WriteString(encoded)
	} else {
		sname.WriteString(name.GetStr())
		hname.WriteString(name.GetStr())
	}
	//sname.ZeroTail()
	//hname.ZeroTail()
	url_app.AppendSmartStr(&sname)
	url_app.WriteByte('=')
	//url_app.ZeroTail()
	form_app.WriteString("<input type=\"hidden\" name=\"")
	form_app.AppendSmartStr(&hname)
	form_app.WriteString("\" value=\"")
	//form_app.ZeroTail()

	/* Short circuit check. Only check url_app. */

	start = (*byte)(core.PhpMemnstr(url_state.GetUrlApp().GetS().GetVal(), url_app.GetS().GetVal(), url_app.GetS().GetLen(), url_state.GetUrlApp().GetS().GetVal()+url_state.GetUrlApp().GetS().GetLen()))
	if start == nil {
		ret = types.FAILURE
		goto finish
	}

	/* Get end of url var */

	limit = url_state.GetUrlApp().GetS().GetVal() + url_state.GetUrlApp().GetS().GetLen()
	end = start + url_app.GetS().GetLen()
	separator_len = strlen(core.PG__().arg_separator.output)
	for end < limit {
		if !(memcmp(end, core.PG__().arg_separator.output, separator_len)) {
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

	if sep_removed == 0 && size_t(start-core.PG__().arg_separator.output) >= separator_len && !(memcmp(start-separator_len, core.PG__().arg_separator.output, separator_len)) {
		start -= separator_len
	}

	/* Remove partially */

	memmove(start, end, url_state.GetUrlApp().GetS().GetLen()-(end-url_state.GetUrlApp().GetS().GetVal()))
	url_state.GetUrlApp().GetS().GetLen() -= end - start
	url_state.GetUrlApp().GetS().GetStr()[url_state.GetUrlApp().GetS().GetLen()] = '0'

	/* Remove form var */

	start = (*byte)(core.PhpMemnstr(url_state.GetFormApp().GetS().GetVal(), form_app.GetS().GetVal(), form_app.GetS().GetLen(), url_state.GetFormApp().GetS().GetVal()+url_state.GetFormApp().GetS().GetLen()))
	if start == nil {

		/* Should not happen */

		ret = types.FAILURE
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
	url_state.GetFormApp().GetS().GetStr()[url_state.GetFormApp().GetS().GetLen()] = '0'
finish:
	url_app.Free()
	form_app.Free()
	sname.Free()
	hname.Free()
	return ret
}
func PhpUrlScannerResetSessionVar(name *types.String, encode int) int {
	return PhpUrlScannerResetVarImpl(name, encode, 1)
}
func PhpUrlScannerResetVar(name *types.String, encode int) int {
	return PhpUrlScannerResetVarImpl(name, encode, 0)
}
func ZmStartupUrlScanner(type_ int, module_number int) int {
	zend.REGISTER_INI_ENTRIES(module_number)
	return types.SUCCESS
}
func ZmShutdownUrlScanner(type_ int, module_number int) int {
	zend.UNREGISTER_INI_ENTRIES(module_number)
	return types.SUCCESS
}
func ZmActivateUrlScanner(type_ int, module_number int) int {
	BG__().url_adapt_session_ex.active = 0
	BG__().url_adapt_session_ex.tag_type = 0
	BG__().url_adapt_session_ex.attr_type = 0
	BG__().url_adapt_output_ex.active = 0
	BG__().url_adapt_output_ex.tag_type = 0
	BG__().url_adapt_output_ex.attr_type = 0
	return types.SUCCESS
}
func ZmDeactivateUrlScanner(type_ int, module_number int) int {
	if BG__().url_adapt_session_ex.active {
		PhpUrlScannerExDeactivate(1)
		BG__().url_adapt_session_ex.active = 0
		BG__().url_adapt_session_ex.tag_type = 0
		BG__().url_adapt_session_ex.attr_type = 0
	}
	BG__().url_adapt_session_ex.form_app.Free()
	BG__().url_adapt_session_ex.url_app.Free()
	if BG__().url_adapt_output_ex.active {
		PhpUrlScannerExDeactivate(0)
		BG__().url_adapt_output_ex.active = 0
		BG__().url_adapt_output_ex.tag_type = 0
		BG__().url_adapt_output_ex.attr_type = 0
	}
	BG__().url_adapt_output_ex.form_app.Free()
	BG__().url_adapt_output_ex.url_app.Free()
	return types.SUCCESS
}
