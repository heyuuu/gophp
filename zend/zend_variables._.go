// <<generate>>

package zend

import b "sik/builtin"

// Source: <Zend/zend_variables.h>

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
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

/* Kept for compatibility */

var ZVAL_PTR_DTOR DtorFuncT = ZvalPtrDtor

// Source: <Zend/zend_variables.c>

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
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

var ZendStringDestroy = func(ptr *any) { b.Free(ptr) }

type ZendRcDtorFuncT func(p *ZendRefcounted)

func __ZendRcDtorFuncTWrapper[T any](fun func(ptr *T)) ZendRcDtorFuncT {
	return func(p *ZendRefcounted) {
		fun(b.Cast[T](p))
	}
}

func __ZendRcDtorFuncTWrapper2[T any, R any](fun func(ptr *T) R) ZendRcDtorFuncT {
	return func(p *ZendRefcounted) {
		fun(b.Cast[T](p))
	}
}

var ZendRcDtorFunc []ZendRcDtorFuncT = []ZendRcDtorFuncT{
	__ZendRcDtorFuncTWrapper(ZendEmptyDestroy),     // IS_UNDEF = 0
	__ZendRcDtorFuncTWrapper(ZendEmptyDestroy),     // IS_NULL = 1
	__ZendRcDtorFuncTWrapper(ZendEmptyDestroy),     // IS_FALSE = 2
	__ZendRcDtorFuncTWrapper(ZendEmptyDestroy),     // IS_TRUE = 3
	__ZendRcDtorFuncTWrapper(ZendEmptyDestroy),     // IS_LONG = 4
	__ZendRcDtorFuncTWrapper(ZendEmptyDestroy),     // IS_DOUBLE = 5
	__ZendRcDtorFuncTWrapper(ZendStringDestroy),    // IS_STRING = 6
	__ZendRcDtorFuncTWrapper(ZendArrayDestroy),     // IS_ARRAY = 7
	__ZendRcDtorFuncTWrapper(ZendObjectsStoreDel),  // IS_OBJECT = 8
	__ZendRcDtorFuncTWrapper2(ZendListFree),        // IS_RESOURCE = 9
	__ZendRcDtorFuncTWrapper(ZendReferenceDestroy), // IS_REFERENCE = 10
	__ZendRcDtorFuncTWrapper(ZendAstRefDestroy),    // IS_CONSTANT_AST = 11
}
