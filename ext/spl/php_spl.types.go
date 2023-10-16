package spl

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/php/types"
)

/**
 * ZendSplGlobals
 */
type ZendSplGlobals struct {
	autoloadExtensions *string
	autoloadFunctions  *types.Array
	hashMaskHandle     uint64
	hashMaskHandlers   uint64
	hashMaskInit       bool
	autoloadRunning    int
}

func (g *ZendSplGlobals) Ctor() {
	g.autoloadExtensions = nil
	g.autoloadFunctions = nil
	g.autoloadRunning = 0
}
func (g *ZendSplGlobals) Reset() {
	g.autoloadExtensions = nil
	g.autoloadFunctions = nil
	g.hashMaskInit = false
}

func (g *ZendSplGlobals) Deactivate() {
	if g.autoloadFunctions != nil {
		g.autoloadFunctions.Destroy()
	}
	g.Reset()
}

func (g *ZendSplGlobals) SplObjectHash(handle uint) string {
	if !g.hashMaskInit {
		g.hashMaskHandle = uint64(standard.PhpMtRand() >> 1)
		g.hashMaskHandlers = uint64(standard.PhpMtRand() >> 1)
		g.hashMaskInit = true
	}
	hashHandle := g.hashMaskHandle ^ uint64(handle)
	hashHandles := g.hashMaskHandlers
	return fmt.Sprintf("%016x%016x", hashHandle, hashHandles)
}

func (g *ZendSplGlobals) GetAutoloadExtensions() string {
	return b.Option(g.autoloadExtensions, SPL_DEFAULT_FILE_EXTENSIONS)
}
func (g *ZendSplGlobals) SetAutoloadExtensions(value string) {
	g.autoloadExtensions = &value
}
func (g *ZendSplGlobals) GetAutoloadFunctions() *types.Array {
	return g.autoloadFunctions
}
func (g *ZendSplGlobals) SetAutoloadFunctions(value *types.Array) {
	g.autoloadFunctions = value
}
func (g *ZendSplGlobals) SetAutoloadRunning(value int) { g.autoloadRunning = value }

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
func (this *AutoloadFuncInfo) GetObj() *types.Zval              { return &this.obj }
func (this *AutoloadFuncInfo) GetClosure() *types.Zval          { return &this.closure }
func (this *AutoloadFuncInfo) GetCe() *types.ClassEntry         { return this.ce }
func (this *AutoloadFuncInfo) SetCe(value *types.ClassEntry)    { this.ce = value }
