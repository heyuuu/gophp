// <<generate>>

package zend

/**
 * ZendClosure
 */
type ZendClosure struct {
	std                   ZendObject
	func_                 ZendFunction
	this_ptr              Zval
	called_scope          *ZendClassEntry
	orig_internal_handler ZifHandler
}

func (this *ZendClosure) GetStd() ZendObject                      { return this.std }
func (this *ZendClosure) SetStd(value ZendObject)                 { this.std = value }
func (this *ZendClosure) GetFunc() ZendFunction                   { return this.func_ }
func (this *ZendClosure) SetFunc(value ZendFunction)              { this.func_ = value }
func (this *ZendClosure) GetThisPtr() Zval                        { return this.this_ptr }
func (this *ZendClosure) SetThisPtr(value Zval)                   { this.this_ptr = value }
func (this *ZendClosure) GetCalledScope() *ZendClassEntry         { return this.called_scope }
func (this *ZendClosure) SetCalledScope(value *ZendClassEntry)    { this.called_scope = value }
func (this *ZendClosure) GetOrigInternalHandler() ZifHandler      { return this.orig_internal_handler }
func (this *ZendClosure) SetOrigInternalHandler(value ZifHandler) { this.orig_internal_handler = value }
