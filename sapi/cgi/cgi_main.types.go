// <<generate>>

package cgi

import (
	"sik/zend"
)

/**
 * php_cgi_globals_struct
 */
type php_cgi_globals_struct struct {
	user_config_cache   zend.HashTable
	redirect_status_env *byte
	rfc2616_headers     zend.ZendBool
	nph                 zend.ZendBool
	check_shebang_line  zend.ZendBool
	fix_pathinfo        zend.ZendBool
	force_redirect      zend.ZendBool
	discard_path        zend.ZendBool
	fcgi_logging        zend.ZendBool
}

func (this *php_cgi_globals_struct) GetUserConfigCache() zend.HashTable {
	return this.user_config_cache
}
func (this *php_cgi_globals_struct) SetUserConfigCache(value zend.HashTable) {
	this.user_config_cache = value
}
func (this *php_cgi_globals_struct) GetRedirectStatusEnv() *byte { return this.redirect_status_env }
func (this *php_cgi_globals_struct) SetRedirectStatusEnv(value *byte) {
	this.redirect_status_env = value
}
func (this *php_cgi_globals_struct) GetRfc2616Headers() zend.ZendBool { return this.rfc2616_headers }
func (this *php_cgi_globals_struct) SetRfc2616Headers(value zend.ZendBool) {
	this.rfc2616_headers = value
}
func (this *php_cgi_globals_struct) GetNph() zend.ZendBool      { return this.nph }
func (this *php_cgi_globals_struct) SetNph(value zend.ZendBool) { this.nph = value }
func (this *php_cgi_globals_struct) GetCheckShebangLine() zend.ZendBool {
	return this.check_shebang_line
}
func (this *php_cgi_globals_struct) SetCheckShebangLine(value zend.ZendBool) {
	this.check_shebang_line = value
}
func (this *php_cgi_globals_struct) GetFixPathinfo() zend.ZendBool      { return this.fix_pathinfo }
func (this *php_cgi_globals_struct) SetFixPathinfo(value zend.ZendBool) { this.fix_pathinfo = value }
func (this *php_cgi_globals_struct) GetForceRedirect() zend.ZendBool    { return this.force_redirect }
func (this *php_cgi_globals_struct) SetForceRedirect(value zend.ZendBool) {
	this.force_redirect = value
}
func (this *php_cgi_globals_struct) GetDiscardPath() zend.ZendBool      { return this.discard_path }
func (this *php_cgi_globals_struct) SetDiscardPath(value zend.ZendBool) { this.discard_path = value }
func (this *php_cgi_globals_struct) GetFcgiLogging() zend.ZendBool      { return this.fcgi_logging }
func (this *php_cgi_globals_struct) SetFcgiLogging(value zend.ZendBool) { this.fcgi_logging = value }

/**
 * UserConfigCacheEntry
 */
type UserConfigCacheEntry struct {
	expires     int64
	user_config *zend.HashTable
}

func (this *UserConfigCacheEntry) GetExpires() int64                   { return this.expires }
func (this *UserConfigCacheEntry) SetExpires(value int64)              { this.expires = value }
func (this *UserConfigCacheEntry) GetUserConfig() *zend.HashTable      { return this.user_config }
func (this *UserConfigCacheEntry) SetUserConfig(value *zend.HashTable) { this.user_config = value }
