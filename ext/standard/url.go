package standard

import (
	"github.com/heyuuu/gophp/php/types"
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

// PhpUrl
type PhpUrl struct {
	scheme   *string
	user     *string
	pass     *string
	host     *string
	port     *uint16
	path     *string
	query    *string
	fragment *string
}

func safeDeref[T any](ptr *T) T {
	var result T
	if ptr != nil {
		return *ptr
	}
	return result
}

func (url *PhpUrl) HasScheme() bool   { return url.scheme != nil }
func (url *PhpUrl) HasUser() bool     { return url.user != nil }
func (url *PhpUrl) HasPass() bool     { return url.pass != nil }
func (url *PhpUrl) HasHost() bool     { return url.host != nil }
func (url *PhpUrl) HasPort() bool     { return url.port != nil }
func (url *PhpUrl) HasPath() bool     { return url.path != nil }
func (url *PhpUrl) HasQuery() bool    { return url.query != nil }
func (url *PhpUrl) HasFragment() bool { return url.fragment != nil }

func (url *PhpUrl) Scheme() string   { return safeDeref(url.scheme) }
func (url *PhpUrl) User() string     { return safeDeref(url.user) }
func (url *PhpUrl) Pass() string     { return safeDeref(url.pass) }
func (url *PhpUrl) Host() string     { return safeDeref(url.host) }
func (url *PhpUrl) Port() uint16     { return safeDeref(url.port) }
func (url *PhpUrl) Path() string     { return safeDeref(url.path) }
func (url *PhpUrl) Query() string    { return safeDeref(url.query) }
func (url *PhpUrl) Fragment() string { return safeDeref(url.fragment) }

func (url *PhpUrl) GetUser() *types.String { return (*types.String)(url.user) }
func (url *PhpUrl) GetPass() *types.String { return (*types.String)(url.pass) }
func (url *PhpUrl) GetHost() *types.String { return (*types.String)(url.host) }
func (url *PhpUrl) GetPath() *types.String { return (*types.String)(url.path) }

func (url *PhpUrl) prepareStr(s string) *string {
	s = urlReplaceControlChars(s)
	return &s
}
func (url *PhpUrl) SetScheme(value string)   { url.scheme = url.prepareStr(value) }
func (url *PhpUrl) SetUser(value string)     { url.user = url.prepareStr(value) }
func (url *PhpUrl) SetPass(value string)     { url.pass = url.prepareStr(value) }
func (url *PhpUrl) SetHost(value string)     { url.host = url.prepareStr(value) }
func (url *PhpUrl) SetPort(value uint16)     { url.port = &value }
func (url *PhpUrl) SetPath(value string)     { url.path = url.prepareStr(value) }
func (url *PhpUrl) SetQuery(value string)    { url.query = url.prepareStr(value) }
func (url *PhpUrl) SetFragment(value string) { url.fragment = url.prepareStr(value) }
