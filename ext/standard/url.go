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
	port     uint16
	path     *string
	query    *string
	fragment *string
}

func safeString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func (url *PhpUrl) Scheme() string   { return safeString(url.scheme) }
func (url *PhpUrl) User() string     { return safeString(url.user) }
func (url *PhpUrl) Pass() string     { return safeString(url.pass) }
func (url *PhpUrl) Host() string     { return safeString(url.host) }
func (url *PhpUrl) Port() uint16     { return url.port }
func (url *PhpUrl) Path() string     { return safeString(url.path) }
func (url *PhpUrl) Query() string    { return safeString(url.query) }
func (url *PhpUrl) Fragment() string { return safeString(url.fragment) }

func (url *PhpUrl) GetScheme() *types.String   { return (*types.String)(url.scheme) }
func (url *PhpUrl) GetUser() *types.String     { return (*types.String)(url.user) }
func (url *PhpUrl) GetPass() *types.String     { return (*types.String)(url.pass) }
func (url *PhpUrl) GetHost() *types.String     { return (*types.String)(url.host) }
func (url *PhpUrl) GetPort() uint16            { return url.port }
func (url *PhpUrl) GetPath() *types.String     { return (*types.String)(url.path) }
func (url *PhpUrl) GetQuery() *types.String    { return (*types.String)(url.query) }
func (url *PhpUrl) GetFragment() *types.String { return (*types.String)(url.fragment) }

func (url *PhpUrl) SetScheme(value string)   { url.scheme = &value }
func (url *PhpUrl) SetUser(value string)     { url.user = &value }
func (url *PhpUrl) SetPass(value string)     { url.pass = &value }
func (url *PhpUrl) SetHost(value string)     { url.host = &value }
func (url *PhpUrl) SetPort(value uint16)     { url.port = value }
func (url *PhpUrl) SetPath(value string)     { url.path = &value }
func (url *PhpUrl) SetQuery(value string)    { url.query = &value }
func (url *PhpUrl) SetFragment(value string) { url.fragment = &value }
