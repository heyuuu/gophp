// <<generate>>

package zend

import (
	b "sik/builtin"
)

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

// #define ZEND_VARIABLES_H

// # include "zend_types.h"

// # include "zend_gc.h"

func ZvalPtrDtorNogc(zval_ptr *Zval) {
	if Z_REFCOUNTED_P(zval_ptr) && Z_DELREF_P(zval_ptr) == 0 {
		RcDtorFunc(Z_COUNTED_P(zval_ptr))
	}
}
func IZvalPtrDtor(zval_ptr *Zval) {
	if Z_REFCOUNTED_P(zval_ptr) {
		var ref *ZendRefcounted = Z_COUNTED_P(zval_ptr)
		if GC_DELREF(ref) == 0 {
			RcDtorFunc(ref)
		} else {
			GcCheckPossibleRoot(ref)
		}
	}
}
func ZvalCopyCtor(zvalue *Zval) {
	if Z_TYPE_P(zvalue) == IS_ARRAY {
		ZVAL_ARR(zvalue, ZendArrayDup(Z_ARR_P(zvalue)))
	} else if Z_REFCOUNTED_P(zvalue) {
		Z_ADDREF_P(zvalue)
	}
}
func ZvalOptCopyCtor(zvalue *Zval) {
	if Z_OPT_TYPE_P(zvalue) == IS_ARRAY {
		ZVAL_ARR(zvalue, ZendArrayDup(Z_ARR_P(zvalue)))
	} else if Z_OPT_REFCOUNTED_P(zvalue) {
		Z_ADDREF_P(zvalue)
	}
}
func ZvalPtrDtorStr(zval_ptr *Zval) {
	if Z_REFCOUNTED_P(zval_ptr) && Z_DELREF_P(zval_ptr) == 0 {
		ZEND_ASSERT(Z_TYPE_P(zval_ptr) == IS_STRING)
		ZEND_ASSERT(ZSTR_IS_INTERNED(Z_STR_P(zval_ptr)) == 0)
		ZEND_ASSERT((GC_FLAGS(Z_STR_P(zval_ptr)) & IS_STR_PERSISTENT) == 0)
		Efree(Z_STR_P(zval_ptr))
	}
}

/* Kept for compatibility */

func ZvalDtor(zvalue *Zval)         { ZvalPtrDtorNogc(zvalue) }
func ZvalInternalDtor(zvalue *Zval) { ZvalInternalPtrDtor(zvalue) }

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

// # include < stdio . h >

// # include "zend.h"

// # include "zend_API.h"

// # include "zend_ast.h"

// # include "zend_globals.h"

// # include "zend_constants.h"

// # include "zend_list.h"

const ZendStringDestroy = _efree

type ZendRcDtorFuncT func(p *ZendRefcounted)

var ZendRcDtorFunc []ZendRcDtorFuncT = []ZendRcDtorFuncT{ZendRcDtorFuncT(ZendEmptyDestroy), ZendRcDtorFuncT(ZendEmptyDestroy), ZendRcDtorFuncT(ZendEmptyDestroy), ZendRcDtorFuncT(ZendEmptyDestroy), ZendRcDtorFuncT(ZendEmptyDestroy), ZendRcDtorFuncT(ZendEmptyDestroy), ZendRcDtorFuncT(ZendStringDestroy), ZendRcDtorFuncT(ZendArrayDestroy), ZendRcDtorFuncT(ZendObjectsStoreDel), ZendRcDtorFuncT(ZendListFree), ZendRcDtorFuncT(ZendReferenceDestroy), ZendRcDtorFuncT(ZendAstRefDestroy)}

func RcDtorFunc(p *ZendRefcounted) {
	ZEND_ASSERT(GC_TYPE(p) <= IS_CONSTANT_AST)
	ZendRcDtorFunc[GC_TYPE(p)](p)
}
func ZendReferenceDestroy(ref *ZendReference) {
	ZEND_ASSERT(!(ZEND_REF_HAS_TYPE_SOURCES(ref)))
	IZvalPtrDtor(&ref.val)
	EfreeSize(ref, b.SizeOf("zend_reference"))
}
func ZendEmptyDestroy(ref *ZendReference) {}
func ZvalPtrDtor(zval_ptr *Zval)          { IZvalPtrDtor(zval_ptr) }

/* }}} */

func ZvalInternalPtrDtor(zval_ptr *Zval) {
	if Z_REFCOUNTED_P(zval_ptr) {
		var ref *ZendRefcounted = Z_COUNTED_P(zval_ptr)
		if GC_DELREF(ref) == 0 {
			if Z_TYPE_P(zval_ptr) == IS_STRING {
				var str *ZendString = (*ZendString)(ref)
				ZEND_ASSERT(ZSTR_IS_INTERNED(str) == 0)
				ZEND_ASSERT((GC_FLAGS(str) & IS_STR_PERSISTENT) != 0)
				Free(str)
			} else {
				ZendErrorNoreturn(E_CORE_ERROR, "Internal zval's can't be arrays, objects, resources or reference")
			}
		}
	}
}

/* }}} */

func ZvalAddRef(p *Zval) {
	if Z_REFCOUNTED_P(p) {
		if Z_ISREF_P(p) && Z_REFCOUNT_P(p) == 1 {
			ZVAL_COPY(p, Z_REFVAL_P(p))
		} else {
			Z_ADDREF_P(p)
		}
	}
}
func ZvalCopyCtorFunc(zvalue *Zval) {
	if EXPECTED(Z_TYPE_P(zvalue) == IS_ARRAY) {
		ZVAL_ARR(zvalue, ZendArrayDup(Z_ARRVAL_P(zvalue)))
	} else if EXPECTED(Z_TYPE_P(zvalue) == IS_STRING) {
		ZEND_ASSERT(ZSTR_IS_INTERNED(Z_STR_P(zvalue)) == 0)
		ZVAL_NEW_STR(zvalue, ZendStringDup(Z_STR_P(zvalue), 0))
	}
}
