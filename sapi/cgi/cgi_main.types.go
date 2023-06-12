package cgi

import (
	"github.com/heyuuu/gophp/php/types"
)

/**
 * php_cgi_globals_struct
 */
type php_cgi_globals_struct struct {
	user_config_cache   *types.Array
	redirect_status_env *byte
	rfc2616_headers     bool
	nph                 bool
	check_shebang_line  bool
	fix_pathinfo        bool
	force_redirect      bool
	discard_path        bool
	fcgi_logging        bool
}

func (this *php_cgi_globals_struct) GetUserConfigCache() *types.Array {
	return this.user_config_cache
}
func (this *php_cgi_globals_struct) SetUserConfigCache(value *types.Array) {
	this.user_config_cache = value
}
func (this *php_cgi_globals_struct) GetRedirectStatusEnv() *byte { return this.redirect_status_env }
func (this *php_cgi_globals_struct) SetRedirectStatusEnv(value *byte) {
	this.redirect_status_env = value
}
func (this *php_cgi_globals_struct) GetRfc2616Headers() bool { return this.rfc2616_headers }
func (this *php_cgi_globals_struct) SetRfc2616Headers(value bool) {
	this.rfc2616_headers = value
}
func (this *php_cgi_globals_struct) GetNph() bool      { return this.nph }
func (this *php_cgi_globals_struct) SetNph(value bool) { this.nph = value }
func (this *php_cgi_globals_struct) GetCheckShebangLine() bool {
	return this.check_shebang_line
}
func (this *php_cgi_globals_struct) SetCheckShebangLine(value bool) {
	this.check_shebang_line = value
}
func (this *php_cgi_globals_struct) GetFixPathinfo() bool      { return this.fix_pathinfo }
func (this *php_cgi_globals_struct) SetFixPathinfo(value bool) { this.fix_pathinfo = value }
func (this *php_cgi_globals_struct) GetForceRedirect() bool    { return this.force_redirect }
func (this *php_cgi_globals_struct) SetForceRedirect(value bool) {
	this.force_redirect = value
}
func (this *php_cgi_globals_struct) GetDiscardPath() bool      { return this.discard_path }
func (this *php_cgi_globals_struct) SetDiscardPath(value bool) { this.discard_path = value }
func (this *php_cgi_globals_struct) GetFcgiLogging() bool      { return this.fcgi_logging }
func (this *php_cgi_globals_struct) SetFcgiLogging(value bool) { this.fcgi_logging = value }

/**
 * UserConfigCacheEntry
 */
type UserConfigCacheEntry struct {
	expires     int64
	user_config *types.Array
}

func (this *UserConfigCacheEntry) GetExpires() int64                { return this.expires }
func (this *UserConfigCacheEntry) SetExpires(value int64)           { this.expires = value }
func (this *UserConfigCacheEntry) GetUserConfig() *types.Array      { return this.user_config }
func (this *UserConfigCacheEntry) SetUserConfig(value *types.Array) { this.user_config = value }
