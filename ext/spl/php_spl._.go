package spl

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

const PHP_SPL_VERSION = core.PHP_VERSION
const PhpextSplPtr = &SplModuleEntry

var SplGlobals ZendSplGlobals

const SPL_DEFAULT_FILE_EXTENSIONS = ".inc,.php"

var SplAutoloadFn types.IFunction = nil
var SplAutoloadCallFn types.IFunction = nil

var SplFunctions []types.FunctionEntry = []types.FunctionEntry{
	DefZifSplClasses,
	DefZifSplAutoload,
	DefZifSplAutoloadExtensions,
	DefZifSplAutoloadRegister,
	DefZifSplAutoloadUnregister,
	DefZifSplAutoloadFunctions,
	DefZifSplAutoloadCall,
	DefZifClassParents,
	DefZifClassImplements,
	DefZifClassUses,
	DefZifSplObjectHash,
	DefZifSplObjectId,
	types.MakeZendFunctionEntryEx("iterator_to_array", 0, ZifIteratorToArray, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types.TypeHintClassName("Traversable", false))),
		zend.MakeArgName("use_keys"),
	}),
	types.MakeZendFunctionEntryEx("iterator_count", 0, ZifIteratorCount, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types.TypeHintClassName("Traversable", false))),
	}),
	types.MakeZendFunctionEntryEx("iterator_apply", 0, ZifIteratorApply, []zend.ArgInfo{zend.MakeReturnArgInfo(2),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types.TypeHintClassName("Traversable", false))),
		zend.MakeArgName("function"),
		zend.MakeArgInfo("args", zend.ArgInfoType(types.TypeHintCode(types.IsArray, 1))),
	}),
}

// SplModuleData
type SplModuleData struct {
	globals ZendSplGlobals
}

var _ zend.ModuleData = (*SplModuleData)(nil)

func (s SplModuleData) Name() string                     { return "SPL" }
func (s SplModuleData) Version() string                  { return PHP_SPL_VERSION }
func (s SplModuleData) Functions() []types.FunctionEntry { return SplFunctions }
func (s SplModuleData) ModuleStartup(moduleNumber int) bool {
	ZmStartupSplExceptions()
	ZmStartupSplIterators()
	ZmStartupSplArray()
	ZmStartupSplDirectory()
	ZmStartupSplDllist()
	ZmStartupSplHeap()
	ZmStartupSplFixedarray()
	ZmStartupSplObserver()

	SplAutoloadFn = zend.CG__().FunctionTable().Get("spl_autoload")
	SplAutoloadCallFn = zend.CG__().FunctionTable().Get("spl_autoload_call")
	b.Assert(SplAutoloadFn != nil && SplAutoloadCallFn != nil)

	return true
}
func (s SplModuleData) ModuleShutdown(moduleNumber int) bool {
	return true
}
func (s SplModuleData) RequestStartup(moduleNumber int) bool {
	SplGlobals = ZendSplGlobals{}
	SplGlobals.Reset()
	return true
}
func (s SplModuleData) RequestShutdown(moduleNumber int) bool {
	SplGlobals.Deactivate()
	return true
}

// SplModuleEntry
var SplModuleEntry = zend.MakeZendModuleEntry(SplModuleData{}, ZmInfoSpl)
