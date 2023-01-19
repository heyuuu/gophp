// <<generate>>

package zend

/**
 * _zendVmStack
 */
type _zendVmStack struct {
	top  *Zval
	end  *Zval
	prev ZendVmStack
}

func (this _zendVmStack) GetTop() *Zval              { return this.top }
func (this *_zendVmStack) SetTop(value *Zval)        { this.top = value }
func (this _zendVmStack) GetEnd() *Zval              { return this.end }
func (this *_zendVmStack) SetEnd(value *Zval)        { this.end = value }
func (this _zendVmStack) GetPrev() ZendVmStack       { return this.prev }
func (this *_zendVmStack) SetPrev(value ZendVmStack) { this.prev = value }
