// <<generate>>

package zend

import (
	b "sik/builtin"
)

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

var ZendCeTraversable *ZendClassEntry
var ZendCeAggregate *ZendClassEntry
var ZendCeIterator *ZendClassEntry
var ZendCeArrayaccess *ZendClassEntry
var ZendCeSerializable *ZendClassEntry
var ZendCeCountable *ZendClassEntry

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

var ZendFuncsAggregate []ZendFunctionEntry = []ZendFunctionEntry{
	MakeZendFunctionEntry("getIterator", nil, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry(nil, nil, nil, 0, 0),
}
var ZendFuncsIterator []ZendFunctionEntry = []ZendFunctionEntry{
	MakeZendFunctionEntry("current", nil, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry("next", nil, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry("key", nil, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry("valid", nil, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry("rewind", nil, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry(nil, nil, nil, 0, 0),
}
var ZendFuncsTraversable *ZendFunctionEntry = nil
var ArginfoArrayaccessOffset []ZendInternalArgInfo = []ZendInternalArgInfo{
	MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	MakeZendInternalArgInfo("offset", 0, 0, 0),
}
var ArginfoArrayaccessOffsetGet []ZendInternalArgInfo = []ZendInternalArgInfo{
	MakeZendInternalArgInfo((*byte)(zend_uintptr_t(1)), 0, 0, 0),
	MakeZendInternalArgInfo("offset", 0, 0, 0),
}
var ArginfoArrayaccessOffsetValue []ZendInternalArgInfo = []ZendInternalArgInfo{
	MakeZendInternalArgInfo((*byte)(zend_uintptr_t(2)), 0, 0, 0),
	MakeZendInternalArgInfo("offset", 0, 0, 0),
	MakeZendInternalArgInfo("value", 0, 0, 0),
}
var ZendFuncsArrayaccess []ZendFunctionEntry = []ZendFunctionEntry{
	MakeZendFunctionEntry("offsetExists", nil, ArginfoArrayaccessOffset, uint32(b.SizeOf("arginfo_arrayaccess_offset")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry("offsetGet", nil, ArginfoArrayaccessOffsetGet, uint32(b.SizeOf("arginfo_arrayaccess_offset_get")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry("offsetSet", nil, ArginfoArrayaccessOffsetValue, uint32(b.SizeOf("arginfo_arrayaccess_offset_value")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry("offsetUnset", nil, ArginfoArrayaccessOffset, uint32(b.SizeOf("arginfo_arrayaccess_offset")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry(nil, nil, nil, 0, 0),
}
var ArginfoSerializableSerialize []ZendInternalArgInfo = []ZendInternalArgInfo{
	MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, ZEND_RETURN_VALUE, 0),
	MakeZendInternalArgInfo("serialized", 0, 0, 0),
}
var ZendFuncsSerializable []ZendFunctionEntry = []ZendFunctionEntry{
	MakeZendFunctionEntry("serialize", nil, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry("unserialize", nil, ArginfoSerializableSerialize, uint32(b.SizeOf("arginfo_serializable_serialize")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry(nil, nil, nil, 0, 0),
}
var ArginfoCountableCount []ZendInternalArgInfo = []ZendInternalArgInfo{
	MakeZendInternalArgInfo((*byte)(zend_uintptr_t(-1)), 0, ZEND_RETURN_VALUE, 0),
}
var ZendFuncsCountable []ZendFunctionEntry = []ZendFunctionEntry{
	MakeZendFunctionEntry("count", nil, ArginfoCountableCount, uint32(b.SizeOf("arginfo_countable_count")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_ABSTRACT),
	MakeZendFunctionEntry(nil, nil, nil, 0, 0),
}
