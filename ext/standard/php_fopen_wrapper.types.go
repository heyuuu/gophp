package standard

import (
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/zend"
)

/**
 * PhpStreamInputT
 */
type PhpStreamInputT struct {
	body     *core.PhpStream
	position zend.ZendOffT
}

// func MakePhpStreamInputT(body *core.PhpStream, position zend.ZendOffT) PhpStreamInputT {
//     return PhpStreamInputT{
//         body:body,
//         position:position,
//     }
// }
func (this *PhpStreamInputT) GetBody() *core.PhpStream        { return this.body }
func (this *PhpStreamInputT) SetBody(value *core.PhpStream)   { this.body = value }
func (this *PhpStreamInputT) GetPosition() zend.ZendOffT      { return this.position }
func (this *PhpStreamInputT) SetPosition(value zend.ZendOffT) { this.position = value }
