package spl

import (
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"unsafe"
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

/* {{{ spl_module_entry
 */

var SplModuleEntry zend.ModuleEntry = zend.MakeZendModuleEntry(
	"SPL",
	SplFunctions,
	ZmStartupSpl,
	nil,
	ZmActivateSpl,
	ZmDeactivateSpl,
	ZmInfoSpl,
	PHP_SPL_VERSION,
	unsafe.Sizeof(ZendSplGlobals),
	SplGlobals,
	(func(any))(ZmGlobalsCtorSpl),
	nil,
)
