package spl

import (
	"github.com/heyuuu/gophp/zend/types"
)

/**
 * ZendSplGlobals
 */
type ZendSplGlobals struct {
	autoload_extensions *types.String
	autoload_functions  *types.Array
	hash_mask_handle    int
	hash_mask_handlers  int
	hash_mask_init      int
	autoload_running    int
}

func (this *ZendSplGlobals) SetAutoloadExtensions(value *types.String) {
	this.autoload_extensions = value
}

func (this *ZendSplGlobals) SetAutoloadFunctions(value *types.Array) {
	this.autoload_functions = value
}
func (this *ZendSplGlobals) SetAutoloadRunning(value int) { this.autoload_running = value }

/**
 * AutoloadFuncInfo
 */
type AutoloadFuncInfo struct {
	func_ptr *types.ZendFunction
	obj      types.Zval
	closure  types.Zval
	ce       *types.ClassEntry
}

func (this *AutoloadFuncInfo) GetFuncPtr() *types.ZendFunction      { return this.func_ptr }
func (this *AutoloadFuncInfo) SetFuncPtr(value *types.ZendFunction) { this.func_ptr = value }
func (this *AutoloadFuncInfo) GetObj() types.Zval                   { return this.obj }
func (this *AutoloadFuncInfo) GetClosure() types.Zval               { return this.closure }
func (this *AutoloadFuncInfo) GetCe() *types.ClassEntry             { return this.ce }
func (this *AutoloadFuncInfo) SetCe(value *types.ClassEntry)        { this.ce = value }
