package zend

import "github.com/heyuuu/gophp/zend/types"

/**
 * ZendClosure
 */
type ZendClosure struct {
	std                   types.ZendObject
	func_                 types.IFunction
	this_ptr              types.Zval
	called_scope          *types.ClassEntry
	orig_internal_handler ZifHandler
}

func (this *ZendClosure) GetStd() types.ZendObject                { return this.std }
func (this *ZendClosure) GetFunc() types.IFunction                { return this.func_ }
func (this *ZendClosure) GetThisPtr() types.Zval                  { return this.this_ptr }
func (this *ZendClosure) GetCalledScope() *types.ClassEntry       { return this.called_scope }
func (this *ZendClosure) SetCalledScope(value *types.ClassEntry)  { this.called_scope = value }
func (this *ZendClosure) GetOrigInternalHandler() ZifHandler      { return this.orig_internal_handler }
func (this *ZendClosure) SetOrigInternalHandler(value ZifHandler) { this.orig_internal_handler = value }
