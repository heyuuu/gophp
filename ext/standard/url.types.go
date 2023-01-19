// <<generate>>

package standard

import (
	"sik/zend"
)

/**
 * PhpUrl
 */
type PhpUrl struct {
	scheme   *zend.ZendString
	user     *zend.ZendString
	pass     *zend.ZendString
	host     *zend.ZendString
	port     uint16
	path     *zend.ZendString
	query    *zend.ZendString
	fragment *zend.ZendString
}

func (this PhpUrl) GetScheme() *zend.ZendString         { return this.scheme }
func (this *PhpUrl) SetScheme(value *zend.ZendString)   { this.scheme = value }
func (this PhpUrl) GetUser() *zend.ZendString           { return this.user }
func (this *PhpUrl) SetUser(value *zend.ZendString)     { this.user = value }
func (this PhpUrl) GetPass() *zend.ZendString           { return this.pass }
func (this *PhpUrl) SetPass(value *zend.ZendString)     { this.pass = value }
func (this PhpUrl) GetHost() *zend.ZendString           { return this.host }
func (this *PhpUrl) SetHost(value *zend.ZendString)     { this.host = value }
func (this PhpUrl) GetPort() uint16                     { return this.port }
func (this *PhpUrl) SetPort(value uint16)               { this.port = value }
func (this PhpUrl) GetPath() *zend.ZendString           { return this.path }
func (this *PhpUrl) SetPath(value *zend.ZendString)     { this.path = value }
func (this PhpUrl) GetQuery() *zend.ZendString          { return this.query }
func (this *PhpUrl) SetQuery(value *zend.ZendString)    { this.query = value }
func (this PhpUrl) GetFragment() *zend.ZendString       { return this.fragment }
func (this *PhpUrl) SetFragment(value *zend.ZendString) { this.fragment = value }
