// <<generate>>

package zend

import (
	b "sik/builtin"
)

// Source: <Zend/zend_weakrefs.h>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: krakjoe@php.net                                             |
   +----------------------------------------------------------------------+
*/

var ZendCeWeakref *ZendClassEntry

// Source: <Zend/zend_weakrefs.c>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: krakjoe@php.net                                             |
   +----------------------------------------------------------------------+
*/

var ZendWeakrefHandlers ZendObjectHandlers
var ZendWeakrefCreateArginfo []ZendInternalArgInfo = []ZendInternalArgInfo{
	{
		(*byte)(zend_uintptr_t(1)),
		ZEND_TYPE_ENCODE_CLASS_CONST("WeakReference", 0),
		0,
		0,
	},
	{"referent", ZEND_TYPE_ENCODE(IS_OBJECT, 0), 0, 0},
}
var ZendWeakrefGetArginfo []ZendInternalArgInfo = []ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(0)), ZEND_TYPE_ENCODE(IS_OBJECT, 1), 0, 0},
}
var ZendWeakrefMethods []ZendFunctionEntry = []ZendFunctionEntry{
	{
		"__construct",
		zim_WeakReference___construct,
		nil,
		uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC,
	},
	{
		"create",
		zim_WeakReference_create,
		ZendWeakrefCreateArginfo,
		uint32(b.SizeOf("zend_weakref_create_arginfo")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC | ZEND_ACC_STATIC,
	},
	{
		"get",
		zim_WeakReference_get,
		ZendWeakrefGetArginfo,
		uint32(b.SizeOf("zend_weakref_get_arginfo")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}
