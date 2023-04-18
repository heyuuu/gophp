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
	rfc2616_headers     types.ZendBool
	nph                 types.ZendBool
	check_shebang_line  types.ZendBool
	fix_pathinfo        types.ZendBool
	force_redirect      types.ZendBool
	discard_path        types.ZendBool
	fcgi_logging        types.ZendBool
}

func (this *php_cgi_globals_struct) GetUserConfigCache() *types.Array {
	return this.user_config_cache
}
func (this *php_cgi_globals_struct) GetRedirectStatusEnv() *byte { return this.redirect_status_env }
func (this *php_cgi_globals_struct) SetRedirectStatusEnv(value *byte) {
	this.redirect_status_env = value
}
func (this *php_cgi_globals_struct) GetRfc2616Headers() types.ZendBool { return this.rfc2616_headers }
func (this *php_cgi_globals_struct) SetRfc2616Headers(value types.ZendBool) {
	this.rfc2616_headers = value
}
func (this *php_cgi_globals_struct) GetNph() types.ZendBool      { return this.nph }
func (this *php_cgi_globals_struct) SetNph(value types.ZendBool) { this.nph = value }
func (this *php_cgi_globals_struct) GetCheckShebangLine() types.ZendBool {
	return this.check_shebang_line
}
func (this *php_cgi_globals_struct) SetCheckShebangLine(value types.ZendBool) {
	this.check_shebang_line = value
}
func (this *php_cgi_globals_struct) GetFixPathinfo() types.ZendBool      { return this.fix_pathinfo }
func (this *php_cgi_globals_struct) SetFixPathinfo(value types.ZendBool) { this.fix_pathinfo = value }
func (this *php_cgi_globals_struct) GetForceRedirect() types.ZendBool    { return this.force_redirect }
func (this *php_cgi_globals_struct) SetForceRedirect(value types.ZendBool) {
	this.force_redirect = value
}
func (this *php_cgi_globals_struct) GetDiscardPath() types.ZendBool      { return this.discard_path }
func (this *php_cgi_globals_struct) SetDiscardPath(value types.ZendBool) { this.discard_path = value }
func (this *php_cgi_globals_struct) GetFcgiLogging() types.ZendBool      { return this.fcgi_logging }
func (this *php_cgi_globals_struct) SetFcgiLogging(value types.ZendBool) { this.fcgi_logging = value }

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
