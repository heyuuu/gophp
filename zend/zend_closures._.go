// <<generate>>

package zend

import (
	b "sik/builtin"
)

// Source: <Zend/zend_closures.h>

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
   | Authors: Christian Seiler <chris_se@gmx.net>                         |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

/* This macro depends on zend_closure structure layout */

var ZendCeClosure *ZendClassEntry

// Source: <Zend/zend_closures.c>

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
   | Authors: Christian Seiler <chris_se@gmx.net>                         |
   |          Dmitry Stogov <dmitry@php.net>                              |
   |          Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

const ZEND_CLOSURE_PRINT_NAME = "Closure object"

/* non-static since it needs to be referenced */

var ClosureHandlers ZendObjectHandlers
var ArginfoClosureBindto []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("newthis"),
	MakeArgInfo("newscope"),
}
var ArginfoClosureBind []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(2),
	MakeArgInfo("closure"),
	MakeArgInfo("newthis"),
	MakeArgInfo("newscope"),
}
var ArginfoClosureCall []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("newthis"),
	MakeArgInfo("parameters", ArgInfoVariadic()),
}
var ArginfoClosureFromcallable []ArgInfo = []ArgInfo{
	MakeReturnArgInfo(1),
	MakeArgInfo("callable"),
}
var ClosureFunctions []ZendFunctionEntry = []ZendFunctionEntry{
	MakeZendFunctionEntry("__construct", zim_Closure___construct, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PRIVATE),
	MakeZendFunctionEntry("bind", zim_Closure_bind, ArginfoClosureBind, uint32(b.SizeOf("arginfo_closure_bind")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_STATIC),
	MakeZendFunctionEntry("bindTo", zim_Closure_bind, ArginfoClosureBindto, uint32(b.SizeOf("arginfo_closure_bindto")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC),
	MakeZendFunctionEntry("call", zim_Closure_call, ArginfoClosureCall, uint32(b.SizeOf("arginfo_closure_call")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC),
	MakeZendFunctionEntry("fromCallable", zim_Closure_fromCallable, ArginfoClosureFromcallable, uint32(b.SizeOf("arginfo_closure_fromcallable")/b.SizeOf("struct _zend_internal_arg_info")-1), ZEND_ACC_PUBLIC|ZEND_ACC_STATIC),
	MakeZendFunctionEntry(nil, nil, nil, 0, 0),
}
