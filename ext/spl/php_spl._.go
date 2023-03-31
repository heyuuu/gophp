package spl

import (
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/types"
	"unsafe"
)

const PHP_SPL_VERSION = core.PHP_VERSION
const PhpextSplPtr = &SplModuleEntry

var SplGlobals ZendSplGlobals

const SPL_DEFAULT_FILE_EXTENSIONS = ".inc,.php"

var SplAutoloadFn *types.ZendFunction = nil
var SplAutoloadCallFn *types.ZendFunction = nil

var SplFunctions []types.ZendFunctionEntry = []types.ZendFunctionEntry{
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
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Traversable", 0))),
		zend.MakeArgName("use_keys"),
	}),
	types.MakeZendFunctionEntryEx("iterator_count", 0, ZifIteratorCount, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Traversable", 0))),
	}),
	types.MakeZendFunctionEntryEx("iterator_apply", 0, ZifIteratorApply, []zend.ArgInfo{zend.MakeReturnArgInfo(2),
		zend.MakeArgInfo("iterator", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Traversable", 0))),
		zend.MakeArgName("function"),
		zend.MakeArgInfo("args", zend.ArgInfoType(types.ZEND_TYPE_ENCODE(types.IS_ARRAY, 1))),
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
