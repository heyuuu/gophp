// <<generate>>

package zend

import "sik/zend/types"

/**
 * _zendVmStack
 */
type _zendVmStack struct {
	top  *types.Zval
	end  *types.Zval
	prev ZendVmStack
}

// func Make_zendVmStack(top *Zval, end *Zval, prev ZendVmStack) _zendVmStack {
//     return _zendVmStack{
//         top:top,
//         end:end,
//         prev:prev,
//     }
// }
func (this *_zendVmStack) GetTop() *types.Zval       { return this.top }
func (this *_zendVmStack) SetTop(value *types.Zval)  { this.top = value }
func (this *_zendVmStack) GetEnd() *types.Zval       { return this.end }
func (this *_zendVmStack) SetEnd(value *types.Zval)  { this.end = value }
func (this *_zendVmStack) GetPrev() ZendVmStack      { return this.prev }
func (this *_zendVmStack) SetPrev(value ZendVmStack) { this.prev = value }
