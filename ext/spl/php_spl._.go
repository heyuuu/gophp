package spl

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
	"sik/zend/types"
)

const PHP_SPL_VERSION = core.PHP_VERSION
const PhpextSplPtr = &SplModuleEntry

var SplGlobals ZendSplGlobals

const SPL_DEFAULT_FILE_EXTENSIONS = ".inc,.php"

var SplAutoloadFn *zend.ZendFunction = nil
var SplAutoloadCallFn *zend.ZendFunction = nil

var SplFunctions []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("spl_classes", 0, ZifSplClasses, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	DefZifSplAutoload,
	DefZifSplAutoloadExtensions,
	DefZifSplAutoloadRegister,
	DefZifSplAutoloadUnregister,
	types.MakeZendFunctionEntryEx("spl_autoload_functions", 0, ZifSplAutoloadFunctions, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
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

var SplModuleEntry zend.ZendModuleEntry = zend.MakeZendModuleEntry(b.SizeOf("zend_module_entry"), zend.ZEND_MODULE_API_NO, 0, zend.USING_ZTS, nil, nil, "SPL", SplFunctions, ZmStartupSpl, nil, ZmActivateSpl, ZmDeactivateSpl, ZmInfoSpl, PHP_SPL_VERSION, core.PHP_MODULE_GLOBALS(spl), (func(any))(ZmGlobalsCtorSpl), nil, nil, 0, 0, nil, 0, "API"+"ZEND_MODULE_API_NO"+zend.ZEND_BUILD_TS)
