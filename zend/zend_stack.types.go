// <<generate>>

package zend

/**
 * ZendStack
 */
type ZendStack struct {
	size     int
	top      int
	max      int
	elements any
}

// func NewZendStack(size int, top int, max int, elements any) *ZendStack {
//     return &ZendStack{
//         size:size,
//         top:top,
//         max:max,
//         elements:elements,
//     }
// }
// func MakeZendStack(size int, top int, max int, elements any) ZendStack {
//     return ZendStack{
//         size:size,
//         top:top,
//         max:max,
//         elements:elements,
//     }
// }
func (this *ZendStack) GetSize() int          { return this.size }
func (this *ZendStack) SetSize(value int)     { this.size = value }
func (this *ZendStack) GetTop() int           { return this.top }
func (this *ZendStack) SetTop(value int)      { this.top = value }
func (this *ZendStack) GetMax() int           { return this.max }
func (this *ZendStack) SetMax(value int)      { this.max = value }
func (this *ZendStack) GetElements() any      { return this.elements }
func (this *ZendStack) SetElements(value any) { this.elements = value }
