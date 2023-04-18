package standard

import (
	"github.com/heyuuu/gophp/php/types"
)

/**
 * PhpUrl
 */
type PhpUrl struct {
	scheme   *types.String
	user     *types.String
	pass     *types.String
	host     *types.String
	port     uint16
	path     *types.String
	query    *types.String
	fragment *types.String
}

func (this *PhpUrl) GetScheme() *types.String        { return this.scheme }
func (this *PhpUrl) SetScheme(value *types.String)   { this.scheme = value }
func (this *PhpUrl) GetUser() *types.String          { return this.user }
func (this *PhpUrl) SetUser(value *types.String)     { this.user = value }
func (this *PhpUrl) GetPass() *types.String          { return this.pass }
func (this *PhpUrl) SetPass(value *types.String)     { this.pass = value }
func (this *PhpUrl) GetHost() *types.String          { return this.host }
func (this *PhpUrl) SetHost(value *types.String)     { this.host = value }
func (this *PhpUrl) GetPort() uint16                 { return this.port }
func (this *PhpUrl) SetPort(value uint16)            { this.port = value }
func (this *PhpUrl) GetPath() *types.String          { return this.path }
func (this *PhpUrl) SetPath(value *types.String)     { this.path = value }
func (this *PhpUrl) GetQuery() *types.String         { return this.query }
func (this *PhpUrl) SetQuery(value *types.String)    { this.query = value }
func (this *PhpUrl) GetFragment() *types.String      { return this.fragment }
func (this *PhpUrl) SetFragment(value *types.String) { this.fragment = value }
