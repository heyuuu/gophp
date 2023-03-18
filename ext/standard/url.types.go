// <<generate>>

package standard

import (
	"sik/zend/types"
)

/**
 * PhpUrl
 */
type PhpUrl struct {
	scheme   *types.ZendString
	user     *types.ZendString
	pass     *types.ZendString
	host     *types.ZendString
	port     uint16
	path     *types.ZendString
	query    *types.ZendString
	fragment *types.ZendString
}

//             func MakePhpUrl(
// scheme *zend.ZendString,
// user *zend.ZendString,
// pass *zend.ZendString,
// host *zend.ZendString,
// port uint16,
// path *zend.ZendString,
// query *zend.ZendString,
// fragment *zend.ZendString,
// ) PhpUrl {
//                 return PhpUrl{
//                     scheme:scheme,
//                     user:user,
//                     pass:pass,
//                     host:host,
//                     port:port,
//                     path:path,
//                     query:query,
//                     fragment:fragment,
//                 }
//             }
func (this *PhpUrl) GetScheme() *types.ZendString        { return this.scheme }
func (this *PhpUrl) SetScheme(value *types.ZendString)   { this.scheme = value }
func (this *PhpUrl) GetUser() *types.ZendString          { return this.user }
func (this *PhpUrl) SetUser(value *types.ZendString)     { this.user = value }
func (this *PhpUrl) GetPass() *types.ZendString          { return this.pass }
func (this *PhpUrl) SetPass(value *types.ZendString)     { this.pass = value }
func (this *PhpUrl) GetHost() *types.ZendString          { return this.host }
func (this *PhpUrl) SetHost(value *types.ZendString)     { this.host = value }
func (this *PhpUrl) GetPort() uint16                     { return this.port }
func (this *PhpUrl) SetPort(value uint16)                { this.port = value }
func (this *PhpUrl) GetPath() *types.ZendString          { return this.path }
func (this *PhpUrl) SetPath(value *types.ZendString)     { this.path = value }
func (this *PhpUrl) GetQuery() *types.ZendString         { return this.query }
func (this *PhpUrl) SetQuery(value *types.ZendString)    { this.query = value }
func (this *PhpUrl) GetFragment() *types.ZendString      { return this.fragment }
func (this *PhpUrl) SetFragment(value *types.ZendString) { this.fragment = value }
