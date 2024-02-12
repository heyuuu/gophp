package standard

import (
	"fmt"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/php/zpp"
	"strconv"
	"strings"
	"unicode"
)

const PHP_URL_SCHEME = 0
const PHP_URL_HOST = 1
const PHP_URL_PORT = 2
const PHP_URL_USER = 3
const PHP_URL_PASS = 4
const PHP_URL_PATH = 5
const PHP_URL_QUERY = 6
const PHP_URL_FRAGMENT = 7
const PHP_QUERY_RFC1738 = 1
const PHP_QUERY_RFC3986 = 2

const lcHexChars = "0123456789abcdef"

// PhpUrl
type PhpUrl struct {
	scheme   string
	host     string
	user     string
	pass     string
	path     string
	query    string
	fragment string
	port     uint16
	flags    uint16
}

func (url *PhpUrl) mark(typ int) {
	if typ < PHP_URL_SCHEME || typ > PHP_URL_FRAGMENT {
		panic(fmt.Sprintf("unexpected url part type: %d", typ))
	}
	url.flags |= 1 << typ
}
func (url *PhpUrl) marked(typ int) bool {
	if typ < PHP_URL_SCHEME || typ > PHP_URL_FRAGMENT {
		panic(fmt.Sprintf("unexpected url part type: %d", typ))
	}
	return url.flags&(1<<typ) != 0
}

func (url *PhpUrl) HasScheme() bool   { return url.marked(PHP_URL_SCHEME) }
func (url *PhpUrl) HasUser() bool     { return url.marked(PHP_URL_USER) }
func (url *PhpUrl) HasPass() bool     { return url.marked(PHP_URL_PASS) }
func (url *PhpUrl) HasHost() bool     { return url.marked(PHP_URL_HOST) }
func (url *PhpUrl) HasPort() bool     { return url.marked(PHP_URL_PORT) }
func (url *PhpUrl) HasPath() bool     { return url.marked(PHP_URL_PATH) }
func (url *PhpUrl) HasQuery() bool    { return url.marked(PHP_URL_QUERY) }
func (url *PhpUrl) HasFragment() bool { return url.marked(PHP_URL_FRAGMENT) }

func (url *PhpUrl) Scheme() string   { return url.scheme }
func (url *PhpUrl) User() string     { return url.user }
func (url *PhpUrl) Pass() string     { return url.pass }
func (url *PhpUrl) Host() string     { return url.host }
func (url *PhpUrl) Port() uint16     { return url.port }
func (url *PhpUrl) Path() string     { return url.path }
func (url *PhpUrl) Query() string    { return url.query }
func (url *PhpUrl) Fragment() string { return url.fragment }

func (url *PhpUrl) prepareStr(s string) *string {
	s = urlReplaceControlChars(s)
	return &s
}
func (url *PhpUrl) SetScheme(value string)   { url.mark(PHP_URL_SCHEME); url.scheme = value }
func (url *PhpUrl) SetUser(value string)     { url.mark(PHP_URL_USER); url.user = value }
func (url *PhpUrl) SetPass(value string)     { url.mark(PHP_URL_PASS); url.pass = value }
func (url *PhpUrl) SetHost(value string)     { url.mark(PHP_URL_HOST); url.host = value }
func (url *PhpUrl) SetPort(value uint16)     { url.mark(PHP_URL_PORT); url.port = value }
func (url *PhpUrl) SetPath(value string)     { url.mark(PHP_URL_PATH); url.path = value }
func (url *PhpUrl) SetQuery(value string)    { url.mark(PHP_URL_QUERY); url.query = value }
func (url *PhpUrl) SetFragment(value string) { url.mark(PHP_URL_FRAGMENT); url.fragment = value }

// --- functions
func urlReplaceControlChars(s string) string {
	// PhpReplaceControlchars || PhpReplaceControlcharsEx
	return strings.Map(func(r rune) rune {
		if unicode.IsControl(r) {
			return '_'
		}
		return r
	}, s)
}

func phpUrlMaybePort(s string) int {
	l := 0
	for l < len(s) && ascii.IsDigit(s[l]) {
		l++
	}
	if (l == len(s) || s[l] == '/') && l < 6 {
		return l
	}
	return 0
}
func phpUrlGetScheme(s string) (string, bool) {
	/* parse scheme */
	if idx := strings.IndexByte(s, ':'); idx > 0 {
		for i := 0; i < idx; i++ {
			c := s[i]
			/* scheme = 1*[ lowalpha | digit | "+" | "-" | "." ] */
			if !ascii.IsAlpha(c) && !ascii.IsDigit(c) && c != '+' && c != '.' && c != '-' {
				return "", false
			}
		}
		if phpUrlMaybePort(s[idx+1:]) > 0 {
			return "", false
		}

		return s[:idx], true
	}
	return "", false
}
func phpUrlMaybeHost(s string, hasScheme bool) (host string, path string, ok bool) {
	if strings.HasPrefix(s, "//") {
		s = s[2:]
	} else if hasScheme {
		return "", s, false
	} else if idx := strings.IndexAny(s, ":"); idx >= 0 {
		for i := 0; i < idx; i++ {

		}
	} else {
		return "", s, false
	}

	if idx := strings.IndexAny(s, "/?#"); idx >= 0 {
		host, path = s[:idx], s[idx:]
	} else {
		host, path = s, ""
	}
	return host, path, host != ""
}

// notice: php 的 urlParse 和 go 标准库的 url.Parse 差异很大，兼容了很多情况
func phpUrlParse(s string) *PhpUrl {
	var ret = &PhpUrl{}
	if s == "" {
		ret.SetPath("")
		return ret
	}

	/* parse scheme */
	if scheme, ok := phpUrlGetScheme(s); ok {
		ret.SetScheme(scheme)
		s = s[len(scheme)+1:]
	}

	/* parse host */
	host, s, ok := phpUrlMaybeHost(s, ret.Scheme() != "")
	if ok {
		/* check for login and password */
		if idx := strings.LastIndexByte(host, '@'); idx >= 0 {
			if eIdx := strings.IndexByte(host[:idx], ':'); eIdx >= 0 {
				ret.SetUser(host[:eIdx])
				ret.SetPass(host[eIdx+1 : idx])
			} else {
				ret.SetUser(host[:idx])
			}
			host = host[idx+1:]
		}

		/* check for port */
		if host != "" && host[0] == '[' && host[len(host)-1] == ']' {
			/* Short circuit portscan, we're dealing with an IPv6 embedded address */
			ret.SetHost(host)
		} else if pIdx := strings.LastIndexByte(host, ':'); pIdx >= 0 {
			ret.SetHost(host[:pIdx])
			portStr := host[pIdx+1:]
			if len(portStr) > 5 {
				return nil
			} else if len(portStr) > 0 {
				port, err := strconv.Atoi(portStr)
				if err == nil && 0 <= port && port <= 65535 {
					ret.SetPort(uint16(port))
				} else {
					return nil
				}
			}
		} else {
			ret.SetHost(host)
		}
	}

	/* parse fragment */
	if idx := strings.IndexByte(s, '#'); idx >= 0 {
		if idx+1 < len(s) {
			ret.SetFragment(s[idx+1:])
		}
		s = s[:idx]
	}

	/* parse query */
	if idx := strings.IndexByte(s, '?'); idx >= 0 {
		if idx+1 < len(s) {
			ret.SetQuery(s[idx+1:])
		}
		s = s[:idx]
	}

	/* parse path */
	if s != "" {
		ret.SetPath(s)
	}

	return ret
}

func ZifParseUrl(ctx *php.Context, url string, _ zpp.Opt, component *int) types.Zval {
	var resource *PhpUrl
	var key = lang.Option(component, -1)
	resource = phpUrlParse(url)
	if resource == nil {
		return types.ZvalFalse()
	}
	if key > -1 {
		switch key {
		case PHP_URL_SCHEME:
			if resource.HasScheme() {
				return types.ZvalString(resource.Scheme())
			}
		case PHP_URL_HOST:
			if resource.HasHost() {
				return types.ZvalString(resource.Host())
			}
		case PHP_URL_PORT:
			if resource.HasPort() {
				return types.ZvalLong(int(resource.Port()))
			}
		case PHP_URL_USER:
			if resource.HasUser() {
				return types.ZvalString(resource.User())
			}
		case PHP_URL_PASS:
			if resource.HasPass() {
				return types.ZvalString(resource.Pass())
			}
		case PHP_URL_PATH:
			if resource.HasPath() {
				return types.ZvalString(resource.Path())
			}
		case PHP_URL_QUERY:
			if resource.HasQuery() {
				return types.ZvalString(resource.Query())
			}
		case PHP_URL_FRAGMENT:
			if resource.HasFragment() {
				return types.ZvalString(resource.Fragment())
			}
		default:
			php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Invalid URL component identifier %d", key))
			return types.ZvalFalse()
		}
		return php.UninitializedZval()
	}

	/* add the various elements to the array */
	arr := types.NewArray()
	if resource.HasScheme() {
		arr.KeyAdd("scheme", types.ZvalString(resource.Scheme()))
	}
	if resource.HasHost() {
		arr.KeyAdd("host", types.ZvalString(resource.Host()))
	}
	if resource.HasPort() {
		arr.KeyAdd("port", types.ZvalLong(int(resource.Port())))
	}
	if resource.HasUser() {
		arr.KeyAdd("user", types.ZvalString(resource.User()))
	}
	if resource.HasPass() {
		arr.KeyAdd("pass", types.ZvalString(resource.Pass()))
	}
	if resource.HasPath() {
		arr.KeyAdd("path", types.ZvalString(resource.Path()))
	}
	if resource.HasQuery() {
		arr.KeyAdd("query", types.ZvalString(resource.Query()))
	}
	if resource.HasFragment() {
		arr.KeyAdd("fragment", types.ZvalString(resource.Fragment()))
	}
	return types.ZvalArray(arr)
}

func PhpUrlEncode(str string) string {
	var buf strings.Builder
	for _, c := range []byte(str) {
		if c == ' ' {
			buf.WriteByte('+')
		} else if c < '0' && c != '-' && c != '.' || c < 'A' && c > '9' || c > 'Z' && c < 'a' && c != '_' || c > 'z' {
			buf.WriteByte('%')
			buf.WriteByte(lcHexChars[c>>4])
			buf.WriteByte(lcHexChars[c&15])
		} else {
			buf.WriteByte(c)
		}
	}
	return buf.String()
}
func PhpRawUrlEncode(str string) string {
	var buf strings.Builder
	for _, c := range []byte(str) {
		if c < '0' && c != '-' && c != '.' || c < 'A' && c > '9' || c > 'Z' && c < 'a' && c != '_' || c > 'z' && c != '~' {
			buf.WriteByte('%')
			buf.WriteByte(lcHexChars[c>>4])
			buf.WriteByte(lcHexChars[c&15])
		} else {
			buf.WriteByte(c)
		}
	}
	return buf.String()
}
