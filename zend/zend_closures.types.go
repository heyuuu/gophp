package zend

import (
	types2 "github.com/heyuuu/gophp/php/types"
)

/**
 * ZendClosure
 */
type ZendClosure struct {
	std                   types2.ZendObject
	func_                 types2.IFunction
	this_ptr              types2.Zval
	called_scope          *types2.ClassEntry
	orig_internal_handler ZifHandler
}

func (this *ZendClosure) GetStd() types2.ZendObject               { return this.std }
func (this *ZendClosure) GetFunc() types2.IFunction               { return this.func_ }
func (this *ZendClosure) GetThisPtr() types2.Zval                 { return this.this_ptr }
func (this *ZendClosure) GetCalledScope() *types2.ClassEntry      { return this.called_scope }
func (this *ZendClosure) SetCalledScope(value *types2.ClassEntry) { this.called_scope = value }
func (this *ZendClosure) GetOrigInternalHandler() ZifHandler      { return this.orig_internal_handler }
func (this *ZendClosure) SetOrigInternalHandler(value ZifHandler) { this.orig_internal_handler = value }
