// <<generate>>

package spl

import (
	"sik/zend"
	"sik/zend/types"
)

/**
 * ZendSplGlobals
 */
type ZendSplGlobals struct {
	autoload_extensions *types.String
	autoload_functions  *types.HashTable
	hash_mask_handle    intPtr
	hash_mask_handlers  intPtr
	hash_mask_init      int
	autoload_running    int
}

//             func MakeZendSplGlobals(
// autoload_extensions *zend.String,
// autoload_functions *zend.HashTable,
// hash_mask_handle intPtr,
// hash_mask_handlers intPtr,
// hash_mask_init int,
// autoload_running int,
// ) ZendSplGlobals {
//                 return ZendSplGlobals{
//                     autoload_extensions:autoload_extensions,
//                     autoload_functions:autoload_functions,
//                     hash_mask_handle:hash_mask_handle,
//                     hash_mask_handlers:hash_mask_handlers,
//                     hash_mask_init:hash_mask_init,
//                     autoload_running:autoload_running,
//                 }
//             }
// func (this *ZendSplGlobals)  GetAutoloadExtensions() *zend.String      { return this.autoload_extensions }
func (this *ZendSplGlobals) SetAutoloadExtensions(value *types.String) {
	this.autoload_extensions = value
}

// func (this *ZendSplGlobals)  GetAutoloadFunctions() *zend.HashTable      { return this.autoload_functions }
func (this *ZendSplGlobals) SetAutoloadFunctions(value *types.HashTable) {
	this.autoload_functions = value
}

// func (this *ZendSplGlobals)  GetHashMaskHandle() intPtr      { return this.hash_mask_handle }
// func (this *ZendSplGlobals) SetHashMaskHandle(value intPtr) { this.hash_mask_handle = value }
// func (this *ZendSplGlobals)  GetHashMaskHandlers() intPtr      { return this.hash_mask_handlers }
// func (this *ZendSplGlobals) SetHashMaskHandlers(value intPtr) { this.hash_mask_handlers = value }
// func (this *ZendSplGlobals)  GetHashMaskInit() int      { return this.hash_mask_init }
// func (this *ZendSplGlobals) SetHashMaskInit(value int) { this.hash_mask_init = value }
// func (this *ZendSplGlobals)  GetAutoloadRunning() int      { return this.autoload_running }
func (this *ZendSplGlobals) SetAutoloadRunning(value int) { this.autoload_running = value }

/**
 * AutoloadFuncInfo
 */
type AutoloadFuncInfo struct {
	func_ptr *zend.ZendFunction
	obj      types.Zval
	closure  types.Zval
	ce       *types.ClassEntry
}

// func MakeAutoloadFuncInfo(func_ptr *zend.ZendFunction, obj zend.Zval, closure zend.Zval, ce *zend.ClassEntry) AutoloadFuncInfo {
//     return AutoloadFuncInfo{
//         func_ptr:func_ptr,
//         obj:obj,
//         closure:closure,
//         ce:ce,
//     }
// }
func (this *AutoloadFuncInfo) GetFuncPtr() *zend.ZendFunction      { return this.func_ptr }
func (this *AutoloadFuncInfo) SetFuncPtr(value *zend.ZendFunction) { this.func_ptr = value }
func (this *AutoloadFuncInfo) GetObj() types.Zval                  { return this.obj }

// func (this *AutoloadFuncInfo) SetObj(value zend.Zval) { this.obj = value }
func (this *AutoloadFuncInfo) GetClosure() types.Zval { return this.closure }

// func (this *AutoloadFuncInfo) SetClosure(value zend.Zval) { this.closure = value }
func (this *AutoloadFuncInfo) GetCe() *types.ClassEntry      { return this.ce }
func (this *AutoloadFuncInfo) SetCe(value *types.ClassEntry) { this.ce = value }
