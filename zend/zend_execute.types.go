package zend

import "github.com/heyuuu/gophp/zend/types"

/**
 * ZendVmStack
 */
type ZendVmStack = *_zendVmStack
type _zendVmStack struct {
	top  *types.Zval
	end  *types.Zval
	prev *_zendVmStack
	// 隐藏内存
	elements any
}

func (this *_zendVmStack) GetTop() *types.Zval       { return this.top }
func (this *_zendVmStack) SetTop(value *types.Zval)  { this.top = value }
func (this *_zendVmStack) GetEnd() *types.Zval       { return this.end }
func (this *_zendVmStack) SetEnd(value *types.Zval)  { this.end = value }
func (this *_zendVmStack) GetPrev() ZendVmStack      { return this.prev }
func (this *_zendVmStack) SetPrev(value ZendVmStack) { this.prev = value }

// 隐藏内存
func (this *_zendVmStack) Elements() *types.Zval          { return this.elements.(*types.Zval) }
func (this *_zendVmStack) ElementsAsEx() *ZendExecuteData { return this.elements.(*ZendExecuteData) }
