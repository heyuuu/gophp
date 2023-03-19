// <<generate>>

package zend

import "sik/zend/types"

// Source: <Zend/zend_interfaces.h>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

var ZendCeTraversable *types.ClassEntry
var ZendCeAggregate *types.ClassEntry
var ZendCeIterator *types.ClassEntry
var ZendCeArrayaccess *types.ClassEntry
var ZendCeSerializable *types.ClassEntry
var ZendCeCountable *types.ClassEntry

// Source: <Zend/zend_interfaces.c>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

/* {{{ zend_call_method
Only returns the returned zval if retval_ptr != NULL */

var ZendInterfaceIteratorFuncsIterator ZendObjectIteratorFuncs = MakeZendObjectIteratorFuncs(ZendUserItDtor, ZendUserItValid, ZendUserItGetCurrentData, ZendUserItGetCurrentKey, ZendUserItMoveForward, ZendUserItRewind, ZendUserItInvalidateCurrent)

/* {{{ zend_user_it_get_iterator */

/* }}}*/

/* }}}*/

/* }}}*/

var ZendFuncsAggregate []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("getIterator", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, nil),
}
var ZendFuncsIterator []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("current", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, nil),
	types.MakeZendFunctionEntryEx("next", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, nil),
	types.MakeZendFunctionEntryEx("key", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, nil),
	types.MakeZendFunctionEntryEx("valid", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, nil),
	types.MakeZendFunctionEntryEx("rewind", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, nil),
}
var ZendFuncsTraversable *types.ZendFunctionEntry = nil
var ArginfoArrayaccessOffset []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("offset"),
}
var ArginfoArrayaccessOffsetGet []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("offset"),
}
var ArginfoArrayaccessOffsetValue []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(2),
	MakeArgInfo("offset"),
	MakeArgInfo("value"),
}
var ZendFuncsArrayaccess []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("offsetExists", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, ArginfoArrayaccessOffset),
	types.MakeZendFunctionEntryEx("offsetGet", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, ArginfoArrayaccessOffsetGet),
	types.MakeZendFunctionEntryEx("offsetSet", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, ArginfoArrayaccessOffsetValue),
	types.MakeZendFunctionEntryEx("offsetUnset", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, ArginfoArrayaccessOffset),
}
var ArginfoSerializableSerialize []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(-1),
	MakeArgInfo("serialized"),
}
var ZendFuncsSerializable []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("serialize", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, nil),
	types.MakeZendFunctionEntryEx("unserialize", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, ArginfoSerializableSerialize),
}
var ArginfoCountableCount []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(-1),
}
var ZendFuncsCountable []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("count", ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT, nil, ArginfoCountableCount),
}
