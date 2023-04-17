package spl

import (
	"fmt"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/zend/types"
)

/**
 * ZendSplGlobals
 */
type ZendSplGlobals struct {
	autoload_extensions *types.String
	autoload_functions  *types.Array
	hashMaskHandle      uint64
	hashMaskHandlers    uint64
	hashMaskInit        bool
	autoloadRunning     int
}

func (this *ZendSplGlobals) Reset() {
	this.autoload_extensions = nil
	this.autoload_functions = nil
	this.hashMaskInit = 0
}

func (this *ZendSplGlobals) Deactivate() {
	if this.autoload_functions != nil {
		this.autoload_functions.Destroy()
	}
	this.Reset()
}

func (this *ZendSplGlobals) SplObjectHash(handle uint32) string {
	if !this.hashMaskInit {
		this.hashMaskHandle = uint64(standard.PhpMtRand() >> 1)
		this.hashMaskHandlers = uint64(standard.PhpMtRand() >> 1)
		this.hashMaskInit = true
	}
	hashHandle := this.hashMaskHandle ^ uint64(handle)
	hashHandles := this.hashMaskHandlers
	return fmt.Sprintf("%016x%016x", hashHandle, hashHandles)
}

func (this *ZendSplGlobals) SetAutoloadExtensions(value *types.String) {
	this.autoload_extensions = value
}

func (this *ZendSplGlobals) SetAutoloadFunctions(value *types.Array) {
	this.autoload_functions = value
}
func (this *ZendSplGlobals) SetAutoloadRunning(value int) { this.autoloadRunning = value }

/**
 * AutoloadFuncInfo
 */
type AutoloadFuncInfo struct {
	func_ptr types.IFunction
	obj      types.Zval
	closure  types.Zval
	ce       *types.ClassEntry
}

func (this *AutoloadFuncInfo) GetFuncPtr() types.IFunction      { return this.func_ptr }
func (this *AutoloadFuncInfo) SetFuncPtr(value types.IFunction) { this.func_ptr = value }
func (this *AutoloadFuncInfo) GetObj() types.Zval               { return this.obj }
func (this *AutoloadFuncInfo) GetClosure() types.Zval           { return this.closure }
func (this *AutoloadFuncInfo) GetCe() *types.ClassEntry         { return this.ce }
func (this *AutoloadFuncInfo) SetCe(value *types.ClassEntry)    { this.ce = value }
