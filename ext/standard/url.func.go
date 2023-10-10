package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/core/streams"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"github.com/heyuuu/gophp/zend/zpp"
	"strings"
	"unicode"
)

func urlReplaceControlChars(s string) string {
	// PhpReplaceControlchars || PhpReplaceControlcharsEx
	return strings.Map(func(r rune) rune {
		if unicode.IsControl(r) {
			return '_'
		}
		return r
	}, s)
}

func BinaryStrcspn(s *byte, e *byte, chars string) *byte {
	for *chars {
		var p *byte = memchr(s, *chars, e-s)
		if p != nil {
			e = p
		}
		chars++
	}
	return e
}

func PhpUrlParse(str *byte) *PhpUrl { return PhpUrlParseString(b.CastStrAuto(str)) }
func PhpUrlParseString(str string) *PhpUrl {
	s := b.CastStrPtr(str)
	return PhpUrlParseEx(s, len(str))
}
func PhpUrlParseEx(str *byte, length int) *PhpUrl {
	var port_buf []byte
	var ret *PhpUrl = &PhpUrl{}
	var s *byte
	var e byte
	var p byte
	var pp byte
	var ue byte
	s = str
	ue = s + length

	/* parse scheme */

	if lang.Assign(&e, memchr(s, ':', length)) && e != s {

		/* validate scheme */

		p = s
		for p < e {

			/* scheme = 1*[ lowalpha | digit | "+" | "-" | "." ] */

			if !(isalpha(*p)) && !(isdigit(*p)) && (*p) != '+' && (*p) != '.' && (*p) != '-' {
				if e+1 < ue && e < BinaryStrcspn(s, ue, "?#") {
					goto parse_port
				} else if s+1 < ue && (*s) == '/' && (*(s + 1)) == '/' {
					s += 2
					e = 0
					goto parse_host
				} else {
					goto just_path
				}
			}
			p++
		}
		if e+1 == ue {
			ret.SetScheme(b.CastStr(s, e-s))
			return ret
		}

		/*
		 * certain schemas like mailto: and zlib: may not have any / after them
		 * this check ensures we support those.
		 */

		if (*(e + 1)) != '/' {

			/* check if the data we get is a port this allows us to
			 * correctly parse things like a.com:80
			 */

			p = e + 1
			for p < ue && isdigit(*p) {
				p++
			}
			if (p == ue || (*p) == '/') && p-e < 7 {
				goto parse_port
			}
			ret.SetScheme(b.CastStr(s, e-s))
			s = e + 1
			goto just_path
		} else {
			ret.SetScheme(b.CastStr(s, e-s))
			if e+2 < ue && (*(e + 2)) == '/' {
				s = e + 3
				if ascii.StrCaseEquals(ret.Scheme(), "file") {
					if e+3 < ue && (*(e + 3)) == '/' {

						/* support windows drive letters as in:
						   file:///c:/somedir/file.txt
						*/

						if e+5 < ue && (*(e + 5)) == ':' {
							s = e + 4
						}
						goto just_path
					}
				}
			} else {
				s = e + 1
				goto just_path
			}
		}

		/*
		 * certain schemas like mailto: and zlib: may not have any / after them
		 * this check ensures we support those.
		 */

	} else if e {
	parse_port:
		p = e + 1
		pp = p
		for pp < ue && pp-p < 6 && isdigit(*pp) {
			pp++
		}
		if pp-p > 0 && pp-p < 6 && (pp == ue || (*pp) == '/') {
			var port zend.ZendLong
			var end *byte
			memcpy(port_buf, p, pp-p)
			port_buf[pp-p] = '0'
			port = zend.ZEND_STRTOL(port_buf, &end, 10)
			if port >= 0 && port <= 65535 && end != port_buf {
				ret.SetPort(uint16(port))
				if s+1 < ue && (*s) == '/' && (*(s + 1)) == '/' {
					s += 2
				}
			} else {
				//PhpUrlFree(ret)
				return nil
			}
		} else if p == pp && pp == ue {
			//PhpUrlFree(ret)
			return nil
		} else if s+1 < ue && (*s) == '/' && (*(s + 1)) == '/' {
			s += 2
		} else {
			goto just_path
		}
	} else if s+1 < ue && (*s) == '/' && (*(s + 1)) == '/' {
		s += 2
	} else {
		goto just_path
	}
parse_host:
	e = BinaryStrcspn(s, ue, "/?#")

	/* check for login and password */

	if lang.Assign(&p, operators.ZendMemrchr(s, '@', e-s)) {
		if lang.Assign(&pp, memchr(s, ':', p-s)) {
			ret.SetUser(b.CastStr(s, pp-s))
			pp++
			ret.SetPass(b.CastStr(pp, p-pp))
		} else {
			ret.SetUser(b.CastStr(s, p-s))
		}
		s = p + 1
	}

	/* check for port */

	if s < ue && (*s) == '[' && (*(e - 1)) == ']' {

		/* Short circuit portscan,
		   we're dealing with an
		   IPv6 embedded address */

		p = nil

		/* Short circuit portscan,
		   we're dealing with an
		   IPv6 embedded address */

	} else {
		p = operators.ZendMemrchr(s, ':', e-s)
	}
	if p {
		if ret.Port() == 0 {
			p++
			if e-p > 5 {
				//PhpUrlFree(ret)
				return nil
			} else if e-p > 0 {
				var port zend.ZendLong
				var end *byte
				memcpy(port_buf, p, e-p)
				port_buf[e-p] = '0'
				port = zend.ZEND_STRTOL(port_buf, &end, 10)
				if port >= 0 && port <= 65535 && end != port_buf {
					ret.SetPort(uint16(port))
				} else {
					//PhpUrlFree(ret)
					return nil
				}
			}
			p--
		}
	} else {
		p = e
	}

	/* check if we have a valid host, if we don't reject the string as url */

	if p-s < 1 {
		//PhpUrlFree(ret)
		return nil
	}
	ret.SetHost(b.CastStr(s, p-s))
	if e == ue {
		return ret
	}
	s = e
just_path:
	e = ue
	p = memchr(s, '#', e-s)
	if p {
		p++
		if p < e {
			ret.SetFragment(b.CastStr(p, e-p))
		}
		e = p - 1
	}
	p = memchr(s, '?', e-s)
	if p {
		p++
		if p < e {
			ret.SetQuery(b.CastStr(p, e-p))
		}
		e = p - 1
	}
	if s < e || s == ue {
		ret.SetPath(b.CastStr(s, e-s))
	}
	return ret
}
func ZifParseUrl(url string, _ zpp.Opt, component *int) *types.Zval {
	var resource *PhpUrl
	var key zend.ZendLong = b.Option(component, -1)
	resource = PhpUrlParseString(url)
	if resource == nil {
		return types.NewZvalFalse()
	}
	if key > -1 {
		switch key {
		case PHP_URL_SCHEME:
			if resource.HasScheme() {
				return types.NewZvalString(resource.Scheme())
			}
		case PHP_URL_HOST:
			if resource.HasHost() {
				return types.NewZvalString(resource.Host())
			}
		case PHP_URL_PORT:
			if resource.HasPort() {
				return types.NewZvalLong(int(resource.Port()))
			}
		case PHP_URL_USER:
			if resource.HasUser() {
				return types.NewZvalString(resource.User())
			}
		case PHP_URL_PASS:
			if resource.HasPass() {
				return types.NewZvalString(resource.Pass())
			}
		case PHP_URL_PATH:
			if resource.HasPath() {
				return types.NewZvalString(resource.Path())
			}
		case PHP_URL_QUERY:
			if resource.HasQuery() {
				return types.NewZvalString(resource.Query())
			}
		case PHP_URL_FRAGMENT:
			if resource.HasFragment() {
				return types.NewZvalString(resource.Fragment())
			}
		default:
			core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid URL component identifier "+zend.ZEND_LONG_FMT, key)
			return types.NewZvalFalse()
		}
	}

	/* add the various elements to the array */
	arr := types.NewArray(0)
	if resource.HasScheme() {
		arr.KeyAddNew("scheme", types.NewZvalString(resource.Scheme()))
	}
	if resource.HasHost() {
		arr.KeyAddNew("host", types.NewZvalString(resource.Host()))
	}
	if resource.HasPort() {
		arr.KeyAddNew("port", types.NewZvalLong(int(resource.Port())))
	}
	if resource.HasUser() {
		arr.KeyAddNew("user", types.NewZvalString(resource.User()))
	}
	if resource.HasPass() {
		arr.KeyAddNew("pass", types.NewZvalString(resource.Pass()))
	}
	if resource.HasPath() {
		arr.KeyAddNew("path", types.NewZvalString(resource.Path()))
	}
	if resource.HasQuery() {
		arr.KeyAddNew("query", types.NewZvalString(resource.Query()))
	}
	if resource.HasFragment() {
		arr.KeyAddNew("fragment", types.NewZvalString(resource.Fragment()))
	}
	return types.NewZvalArray(arr)
}
func _H2I(c byte) byte {
	if '0' <= c && c <= '9' {
		return c - '0'
	} else if 'a' <= c && c <= 'z' {
		return c - 'a' + 10
	} else if 'A' <= c && c <= 'Z' {
		return c - 'A' + 10
	} else {
		return 0
	}
}

func PhpHtoi(s *byte) int {
	var value int
	var c int
	c = (*uint8)(s)[0]
	if isupper(c) {
		c = tolower(c)
	}
	value = lang.Cond(c >= '0' && c <= '9', c-'0', c-'a'+10) * 16
	c = (*uint8)(s)[1]
	if isupper(c) {
		c = tolower(c)
	}
	if c >= '0' && c <= '9' {
		value += c - '0'
	} else {
		value += c - 'a' + 10
	}
	return value
}
func PhpUrlEncodeEx(s string) string {
	var buf strings.Builder
	for _, c := range []byte(s) {
		if c == ' ' {
			buf.WriteByte('+')
		} else if !ascii.IsAlphaNum(c) && c != '-' && c != '.' && c != '_' {
			buf.WriteByte('%')
			buf.WriteByte(c >> 4)
			buf.WriteByte(c & 0x0F)
		} else {
			buf.WriteByte(c)
		}
	}
	return buf.String()
}
func PhpUrlEncode(str string) string {
	var buf strings.Builder
	for _, c := range []byte(str) {
		if c == ' ' {
			buf.WriteByte('+')
		} else if c < '0' && c != '-' && c != '.' || c < 'A' && c > '9' || c > 'Z' && c < 'a' && c != '_' || c > 'z' {
			buf.WriteByte('%')
			buf.WriteByte(Hexchars[c>>4])
			buf.WriteByte(Hexchars[c&15])
		} else {
			buf.WriteByte(c)
		}
	}
	return buf.String()
}

func ZifUrlencode(str string) string {
	return PhpUrlEncodeEx(str)
}
func ZifUrldecode(str string) string {
	return PhpUrlDecodeEx(str)
}

func PhpUrlDecodeEx(str string) string {
	var buf strings.Builder
	for i := 0; i < len(str); i++ {
		c := str[i]
		if c == '+' {
			buf.WriteByte(' ')
		} else if c == '%' && i+2 < len(str) && ascii.IsXDigit(str[i+1]) && ascii.IsXDigit(str[i+2]) {
			tmp := _H2I(str[i+1])*18 + _H2I(str[i+2])
			buf.WriteByte(tmp)
			i += 2
		} else {
			buf.WriteByte(c)
		}
	}
	return buf.String()
}
func PhpUrlDecode(str *byte, len_ int) int {
	var dest *byte = str
	var data *byte = str
	for lang.PostDec(&len_) {
		if (*data) == '+' {
			*dest = ' '
		} else if (*data) == '%' && len_ >= 2 && isxdigit(int(*(data + 1))) && isxdigit(int(*(data + 2))) {
			*dest = byte(PhpHtoi(data + 1))
			data += 2
			len_ -= 2
		} else {
			*dest = *data
		}
		data++
		dest++
	}
	*dest = '0'
	return dest - str
}

func PhpRawUrlEncode(str string) string {
	var buf strings.Builder
	for _, c := range []byte(str) {
		if c < '0' && c != '-' && c != '.' || c < 'A' && c > '9' || c > 'Z' && c < 'a' && c != '_' || c > 'z' && c != '~' {
			buf.WriteByte('%')
			buf.WriteByte(Hexchars[c>>4])
			buf.WriteByte(Hexchars[c&15])
		} else {
			buf.WriteByte(c)
		}
	}
	return buf.String()
}
func ZifRawurlencode(executeData zpp.Ex, return_value zpp.Ret, str *types.Zval) {
	var in_str *types.String
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			in_str = fp.ParseStr()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	return_value.SetString(PhpRawUrlEncode(in_str.GetStr()))
	return
}
func ZifRawurldecode(str string) string {
	return PhpRawUrlDecodeEx(str)
}
func PhpRawUrlDecodeEx(str string) string {
	var buf strings.Builder
	for i := 0; i < len(str); i++ {
		c := str[i]
		if c == '%' && i+2 < len(str) && ascii.IsXDigit(str[i+1]) && ascii.IsXDigit(str[i+2]) {
			tmp := _H2I(str[i+1])*18 + _H2I(str[i+2])
			buf.WriteByte(tmp)
			i += 2
		} else {
			buf.WriteByte(c)
		}
	}
	return buf.String()
}
func PhpRawUrlDecode(str *byte, len_ int) int {
	var dest *byte = str
	var data *byte = str
	for lang.PostDec(&len_) {
		if (*data) == '%' && len_ >= 2 && isxdigit(int(*(data + 1))) && isxdigit(int(*(data + 2))) {
			*dest = byte(PhpHtoi(data + 1))
			data += 2
			len_ -= 2
		} else {
			*dest = *data
		}
		data++
		dest++
	}
	*dest = '0'
	return dest - str
}
func ZifGetHeaders(executeData zpp.Ex, return_value zpp.Ret, url *types.Zval, _ zpp.Opt, format *types.Zval, context *types.Zval) {
	var url *byte
	var url_len int
	var stream *core.PhpStream
	var prev_val *types.Zval
	var hdr *types.Zval = nil
	var format zend.ZendLong = 0
	var zcontext *types.Zval = nil
	var context *core.PhpStreamContext
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 3, 0)
			url, url_len = fp.ParsePath()
			fp.StartOptional()
			format = fp.ParseLong()
			zcontext = fp.ParseResourceEx(true, false)
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	context = streams.PhpStreamContextFromZval(zcontext, 0)
	if !(lang.Assign(&stream, core.PhpStreamOpenWrapperEx(url, "r", core.REPORT_ERRORS|core.STREAM_USE_URL|core.STREAM_ONLY_GET_HEADERS, nil, context))) {
		return_value.SetFalse()
		return
	}
	if stream.GetWrapperdata().Type() != types.IsArray {
		core.PhpStreamClose(stream)
		return_value.SetFalse()
		return
	}
	zend.ArrayInit(return_value)
	var __ht *types.Array = stream.GetWrapperdata().Array()
	for _, _p := range __ht.ForeachData() {
		var _z *types.Zval = _p.GetVal()

		hdr = _z
		if !hdr.IsString() {
			continue
		}
		if format == 0 {
		no_name_header:
			zend.AddNextIndexStr(return_value, hdr.StringEx().Copy())
		} else {
			var c byte
			var s *byte
			var p *byte
			if lang.Assign(&p, strchr(hdr.StringEx().GetVal(), ':')) {
				c = *p
				*p = '0'
				s = p + 1
				for isspace(int(*((*uint8)(s)))) {
					s++
				}
				if lang.Assign(&prev_val, return_value.Array().KeyFind(b.CastStr(hdr.StringEx().GetVal(), p-hdr.StringEx().GetVal()))) == nil {
					zend.AddAssocStringlEx(return_value, b.CastStr(hdr.StringEx().GetVal(), p-hdr.StringEx().GetVal()), b.CastStr(s, hdr.StringEx().GetLen()-(s-hdr.StringEx().GetVal())))
				} else {
					operators.ConvertToArray(prev_val)
					zend.AddNextIndexStringl(prev_val, s, hdr.StringEx().GetLen()-(s-hdr.StringEx().GetVal()))
				}
				*p = c
			} else {
				goto no_name_header
			}
		}
	}
	core.PhpStreamClose(stream)
}
