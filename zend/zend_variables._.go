// <<generate>>

package zend

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

const ZvalDtorFunc = RcDtorFunc
const ZvalPtrDtorWrapper = ZvalPtrDtor
const ZvalInternalPtrDtorWrapper = ZvalInternalPtrDtor
const ZVAL_PTR_DTOR DtorFuncT = ZvalPtrDtor
const ZVAL_INTERNAL_PTR_DTOR = ZvalInternalPtrDtor

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

const ZendStringDestroy = _efree

type ZendRcDtorFuncT func(p *ZendRefcounted)

var ZendRcDtorFunc []ZendRcDtorFuncT = []ZendRcDtorFuncT{ZendRcDtorFuncT(ZendEmptyDestroy), ZendRcDtorFuncT(ZendEmptyDestroy), ZendRcDtorFuncT(ZendEmptyDestroy), ZendRcDtorFuncT(ZendEmptyDestroy), ZendRcDtorFuncT(ZendEmptyDestroy), ZendRcDtorFuncT(ZendEmptyDestroy), ZendRcDtorFuncT(ZendStringDestroy), ZendRcDtorFuncT(ZendArrayDestroy), ZendRcDtorFuncT(ZendObjectsStoreDel), ZendRcDtorFuncT(ZendListFree), ZendRcDtorFuncT(ZendReferenceDestroy), ZendRcDtorFuncT(ZendAstRefDestroy)}
