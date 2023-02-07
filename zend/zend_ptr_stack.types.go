// <<generate>>

package zend

/**
 * ZendPtrStack
 */
type ZendPtrStack struct {
	top         int
	max         int
	elements    *any
	top_element *any
	persistent  ZendBool
}

// func NewZendPtrStack(top int, max int, elements *any, top_element *any, persistent ZendBool) *ZendPtrStack {
//     return &ZendPtrStack{
//         top:top,
//         max:max,
//         elements:elements,
//         top_element:top_element,
//         persistent:persistent,
//     }
// }
// func MakeZendPtrStack(top int, max int, elements *any, top_element *any, persistent ZendBool) ZendPtrStack {
//     return ZendPtrStack{
//         top:top,
//         max:max,
//         elements:elements,
//         top_element:top_element,
//         persistent:persistent,
//     }
// }
func (this *ZendPtrStack) GetTop() int                  { return this.top }
func (this *ZendPtrStack) SetTop(value int)             { this.top = value }
func (this *ZendPtrStack) GetMax() int                  { return this.max }
func (this *ZendPtrStack) SetMax(value int)             { this.max = value }
func (this *ZendPtrStack) GetElements() *any            { return this.elements }
func (this *ZendPtrStack) SetElements(value *any)       { this.elements = value }
func (this *ZendPtrStack) GetTopElement() *any          { return this.top_element }
func (this *ZendPtrStack) SetTopElement(value *any)     { this.top_element = value }
func (this *ZendPtrStack) GetPersistent() ZendBool      { return this.persistent }
func (this *ZendPtrStack) SetPersistent(value ZendBool) { this.persistent = value }
