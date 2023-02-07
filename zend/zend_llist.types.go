// <<generate>>

package zend

/**
 * ZendLlistElement
 */
type ZendLlistElement struct {
	next *ZendLlistElement
	prev *ZendLlistElement
	data []byte
}

// func NewZendLlistElement(next *ZendLlistElement, prev *ZendLlistElement, data []byte) *ZendLlistElement {
//     return &ZendLlistElement{
//         next:next,
//         prev:prev,
//         data:data,
//     }
// }
// func MakeZendLlistElement(next *ZendLlistElement, prev *ZendLlistElement, data []byte) ZendLlistElement {
//     return ZendLlistElement{
//         next:next,
//         prev:prev,
//         data:data,
//     }
// }
func (this *ZendLlistElement) GetNext() *ZendLlistElement      { return this.next }
func (this *ZendLlistElement) SetNext(value *ZendLlistElement) { this.next = value }
func (this *ZendLlistElement) GetPrev() *ZendLlistElement      { return this.prev }
func (this *ZendLlistElement) SetPrev(value *ZendLlistElement) { this.prev = value }
func (this *ZendLlistElement) GetData() []byte                 { return this.data }

// func (this *ZendLlistElement) SetData(value []byte) { this.data = value }

/**
 * ZendLlist
 */
type ZendLlist struct {
	head         *ZendLlistElement
	tail         *ZendLlistElement
	count        int
	size         int
	dtor         LlistDtorFuncT
	persistent   uint8
	traverse_ptr *ZendLlistElement
}

// func NewZendLlist(head *ZendLlistElement, tail *ZendLlistElement, count int, size int, dtor LlistDtorFuncT, persistent uint8, traverse_ptr *ZendLlistElement) *ZendLlist {
//     return &ZendLlist{
//         head:head,
//         tail:tail,
//         count:count,
//         size:size,
//         dtor:dtor,
//         persistent:persistent,
//         traverse_ptr:traverse_ptr,
//     }
// }
// func MakeZendLlist(head *ZendLlistElement, tail *ZendLlistElement, count int, size int, dtor LlistDtorFuncT, persistent uint8, traverse_ptr *ZendLlistElement) ZendLlist {
//     return ZendLlist{
//         head:head,
//         tail:tail,
//         count:count,
//         size:size,
//         dtor:dtor,
//         persistent:persistent,
//         traverse_ptr:traverse_ptr,
//     }
// }
func (this *ZendLlist) GetHead() *ZendLlistElement        { return this.head }
func (this *ZendLlist) SetHead(value *ZendLlistElement)   { this.head = value }
func (this *ZendLlist) GetTail() *ZendLlistElement        { return this.tail }
func (this *ZendLlist) SetTail(value *ZendLlistElement)   { this.tail = value }
func (this *ZendLlist) GetCount() int                     { return this.count }
func (this *ZendLlist) SetCount(value int)                { this.count = value }
func (this *ZendLlist) GetSize() int                      { return this.size }
func (this *ZendLlist) SetSize(value int)                 { this.size = value }
func (this *ZendLlist) GetDtor() LlistDtorFuncT           { return this.dtor }
func (this *ZendLlist) SetDtor(value LlistDtorFuncT)      { this.dtor = value }
func (this *ZendLlist) GetPersistent() uint8              { return this.persistent }
func (this *ZendLlist) SetPersistent(value uint8)         { this.persistent = value }
func (this *ZendLlist) GetTraversePtr() *ZendLlistElement { return this.traverse_ptr }

// func (this *ZendLlist) SetTraversePtr(value *ZendLlistElement) { this.traverse_ptr = value }
