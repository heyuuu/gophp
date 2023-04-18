package cgi

import (
	types2 "github.com/heyuuu/gophp/php/types"
)

/**
 * php_cgi_globals_struct
 */
type php_cgi_globals_struct struct {
	user_config_cache   *types2.Array
	redirect_status_env *byte
	rfc2616_headers     types2.ZendBool
	nph                 types2.ZendBool
	check_shebang_line  types2.ZendBool
	fix_pathinfo        types2.ZendBool
	force_redirect      types2.ZendBool
	discard_path        types2.ZendBool
	fcgi_logging        types2.ZendBool
}

func (this *php_cgi_globals_struct) GetUserConfigCache() *types2.Array {
	return this.user_config_cache
}
func (this *php_cgi_globals_struct) GetRedirectStatusEnv() *byte { return this.redirect_status_env }
func (this *php_cgi_globals_struct) SetRedirectStatusEnv(value *byte) {
	this.redirect_status_env = value
}
func (this *php_cgi_globals_struct) GetRfc2616Headers() types2.ZendBool { return this.rfc2616_headers }
func (this *php_cgi_globals_struct) SetRfc2616Headers(value types2.ZendBool) {
	this.rfc2616_headers = value
}
func (this *php_cgi_globals_struct) GetNph() types2.ZendBool      { return this.nph }
func (this *php_cgi_globals_struct) SetNph(value types2.ZendBool) { this.nph = value }
func (this *php_cgi_globals_struct) GetCheckShebangLine() types2.ZendBool {
	return this.check_shebang_line
}
func (this *php_cgi_globals_struct) SetCheckShebangLine(value types2.ZendBool) {
	this.check_shebang_line = value
}
func (this *php_cgi_globals_struct) GetFixPathinfo() types2.ZendBool      { return this.fix_pathinfo }
func (this *php_cgi_globals_struct) SetFixPathinfo(value types2.ZendBool) { this.fix_pathinfo = value }
func (this *php_cgi_globals_struct) GetForceRedirect() types2.ZendBool    { return this.force_redirect }
func (this *php_cgi_globals_struct) SetForceRedirect(value types2.ZendBool) {
	this.force_redirect = value
}
func (this *php_cgi_globals_struct) GetDiscardPath() types2.ZendBool      { return this.discard_path }
func (this *php_cgi_globals_struct) SetDiscardPath(value types2.ZendBool) { this.discard_path = value }
func (this *php_cgi_globals_struct) GetFcgiLogging() types2.ZendBool      { return this.fcgi_logging }
func (this *php_cgi_globals_struct) SetFcgiLogging(value types2.ZendBool) { this.fcgi_logging = value }

/**
 * UserConfigCacheEntry
 */
type UserConfigCacheEntry struct {
	expires     int64
	user_config *types2.Array
}

func (this *UserConfigCacheEntry) GetExpires() int64                 { return this.expires }
func (this *UserConfigCacheEntry) SetExpires(value int64)            { this.expires = value }
func (this *UserConfigCacheEntry) GetUserConfig() *types2.Array      { return this.user_config }
func (this *UserConfigCacheEntry) SetUserConfig(value *types2.Array) { this.user_config = value }
