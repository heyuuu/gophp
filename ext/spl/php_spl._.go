// <<generate>>

package spl

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

// Source: <ext/spl/php_spl.h>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Authors: Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

const PHP_SPL_VERSION = core.PHP_VERSION
const PhpextSplPtr = &SplModuleEntry

var ZmShutdownSpl func(type_ int, module_number int) int

var SplGlobals ZendSplGlobals

// Source: <ext/spl/php_spl.c>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Authors: Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

const SPL_DEFAULT_FILE_EXTENSIONS = ".inc,.php"

var SplAutoloadFn *zend.ZendFunction = nil
var SplAutoloadCallFn *zend.ZendFunction = nil

/* {{{ PHP_GINIT_FUNCTION
 */

/* {{{ proto array class_parents(object instance [, bool autoload = true])
Return an array containing the names of all parent classes */

/* {{{ proto array spl_classes()
Return an array containing the names of all clsses and interfaces defined in SPL */

/* {{{ proto void spl_autoload(string class_name [, string file_extensions])
Default implementation for __autoload() */

/* {{{ proto string spl_autoload_extensions([string file_extensions])
Register and return default file extensions for spl_autoload */

/* {{{ proto void spl_autoload_call(string class_name)
Try all registered autoload function to load the requested class */

/* {{{ proto bool spl_autoload_register([mixed autoload_function [, bool throw [, bool prepend]]])
Register given function as __autoload() implementation */

/* {{{ proto bool spl_autoload_unregister(mixed autoload_function)
Unregister given function as __autoload() implementation */

/* {{{ proto false|array spl_autoload_functions()
Return all registered __autoload() functionns */

/* {{{ proto string spl_object_hash(object obj)
Return hash id for given object */

/* {{{ PHP_MINFO(spl)
 */

var ArginfoIteratorToArray []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("iterator", ArgInfoType(zend.ZEND_TYPE_ENCODE_CLASS_CONST("Traversable", 0))),
	zend.MakeArgInfo("use_keys"),
}
var ArginfoIterator []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("iterator", ArgInfoType(zend.ZEND_TYPE_ENCODE_CLASS_CONST("Traversable", 0))),
}
var ArginfoIteratorApply []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(2),
	zend.MakeArgInfo("iterator", ArgInfoType(zend.ZEND_TYPE_ENCODE_CLASS_CONST("Traversable", 0))),
	zend.MakeArgInfo("function"),
	zend.MakeArgInfo("args", ArgInfoType(zend.ZEND_TYPE_ENCODE(zend.IS_ARRAY, 1))),
}
var ArginfoClassParents []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("instance"),
	zend.MakeArgInfo("autoload"),
}
var ArginfoClassImplements []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("what"),
	zend.MakeArgInfo("autoload"),
}
var ArginfoClassUses []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("what"),
	zend.MakeArgInfo("autoload"),
}
var ArginfoSplClasses []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoSplAutoloadFunctions []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var ArginfoSplAutoload []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("class_name"),
	zend.MakeArgInfo("file_extensions"),
}
var ArginfoSplAutoloadExtensions []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("file_extensions"),
}
var ArginfoSplAutoloadCall []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("class_name"),
}
var ArginfoSplAutoloadRegister []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("autoload_function"),
	zend.MakeArgInfo("throw"),
	zend.MakeArgInfo("prepend"),
}
var ArginfoSplAutoloadUnregister []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("autoload_function"),
}
var ArginfoSplObjectHash []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("obj"),
}
var ArginfoSplObjectId []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("obj"),
}
var SplFunctions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	zend.MakeZendFunctionEntryEx("spl_classes", 0, ZifSplClasses, ArginfoSplClasses),
	zend.MakeZendFunctionEntryEx("spl_autoload", 0, ZifSplAutoload, ArginfoSplAutoload),
	zend.MakeZendFunctionEntryEx("spl_autoload_extensions", 0, ZifSplAutoloadExtensions, ArginfoSplAutoloadExtensions),
	zend.MakeZendFunctionEntryEx("spl_autoload_register", 0, ZifSplAutoloadRegister, ArginfoSplAutoloadRegister),
	zend.MakeZendFunctionEntryEx("spl_autoload_unregister", 0, ZifSplAutoloadUnregister, ArginfoSplAutoloadUnregister),
	zend.MakeZendFunctionEntryEx("spl_autoload_functions", 0, ZifSplAutoloadFunctions, ArginfoSplAutoloadFunctions),
	zend.MakeZendFunctionEntryEx("spl_autoload_call", 0, ZifSplAutoloadCall, ArginfoSplAutoloadCall),
	zend.MakeZendFunctionEntryEx("class_parents", 0, ZifClassParents, ArginfoClassParents),
	zend.MakeZendFunctionEntryEx("class_implements", 0, ZifClassImplements, ArginfoClassImplements),
	zend.MakeZendFunctionEntryEx("class_uses", 0, ZifClassUses, ArginfoClassUses),
	zend.MakeZendFunctionEntryEx("spl_object_hash", 0, ZifSplObjectHash, ArginfoSplObjectHash),
	zend.MakeZendFunctionEntryEx("spl_object_id", 0, ZifSplObjectId, ArginfoSplObjectId),
	zend.MakeZendFunctionEntryEx("iterator_to_array", 0, ZifIteratorToArray, ArginfoIteratorToArray),
	zend.MakeZendFunctionEntryEx("iterator_count", 0, ZifIteratorCount, ArginfoIterator),
	zend.MakeZendFunctionEntryEx("iterator_apply", 0, ZifIteratorApply, ArginfoIteratorApply),
}

/* {{{ spl_module_entry
 */

var SplModuleEntry zend.ZendModuleEntry = zend.MakeZendModuleEntry(b.SizeOf("zend_module_entry"), zend.ZEND_MODULE_API_NO, 0, zend.USING_ZTS, nil, nil, "SPL", SplFunctions, ZmStartupSpl, nil, ZmActivateSpl, ZmDeactivateSpl, ZmInfoSpl, PHP_SPL_VERSION, core.PHP_MODULE_GLOBALS(spl), (func(any))(ZmGlobalsCtorSpl), nil, nil, 0, 0, nil, 0, "API"+"ZEND_MODULE_API_NO"+zend.ZEND_BUILD_TS)
