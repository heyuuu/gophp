package spl

import (
	"github.com/heyuuu/gophp/core"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"unsafe"
)

const PHP_SPL_VERSION = core.PHP_VERSION
const PhpextSplPtr = &SplModuleEntry

var SplGlobals ZendSplGlobals

const SPL_DEFAULT_FILE_EXTENSIONS = ".inc,.php"

var SplAutoloadFn types2.IFunction = nil
var SplAutoloadCallFn types2.IFunction = nil

var SplFunctions []types2.FunctionEntry = []types2.FunctionEntry{
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
	types2.MakeZendFunctionEntryEx("iterator_to_array", 0, ZifIteratorToArray, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types2.ZEND_TYPE_ENCODE_CLASS_CONST("Traversable", 0))),
		zend.MakeArgName("use_keys"),
	}),
	types2.MakeZendFunctionEntryEx("iterator_count", 0, ZifIteratorCount, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types2.ZEND_TYPE_ENCODE_CLASS_CONST("Traversable", 0))),
	}),
	types2.MakeZendFunctionEntryEx("iterator_apply", 0, ZifIteratorApply, []zend.ArgInfo{zend.MakeReturnArgInfo(2),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types2.ZEND_TYPE_ENCODE_CLASS_CONST("Traversable", 0))),
		zend.MakeArgName("function"),
		zend.MakeArgInfo("args", zend.ArgInfoType(types2.ZEND_TYPE_ENCODE(types2.IS_ARRAY, 1))),
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
