// <<generate>>

package zend

import (
	r "sik/runtime"
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
	if zval_ptr.GetTypeFlags() != 0 && ZvalDelrefP(zval_ptr) == 0 {
		RcDtorFunc(zval_ptr.GetValue().GetCounted())
	}
}
func IZvalPtrDtor(zval_ptr *Zval) {
	if zval_ptr.GetTypeFlags() != 0 {
		var ref *ZendRefcounted = zval_ptr.GetValue().GetCounted()
		if ZendGcDelref(&ref.gc) == 0 {
			RcDtorFunc(ref)
		} else {
			GcCheckPossibleRoot(ref)
		}
	}
}
func ZvalCopyCtor(zvalue *Zval) {
	if zvalue.GetType() == 7 {
		var __arr *ZendArray = ZendArrayDup(zvalue.GetValue().GetArr())
		var __z *Zval = zvalue
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	} else if zvalue.GetTypeFlags() != 0 {
		ZvalAddrefP(zvalue)
	}
}
func ZvalOptCopyCtor(zvalue *Zval) {
	if (zvalue.GetTypeInfo() & 0xff) == 7 {
		var __arr *ZendArray = ZendArrayDup(zvalue.GetValue().GetArr())
		var __z *Zval = zvalue
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	} else if (zvalue.GetTypeInfo() & 0xff00) != 0 {
		ZvalAddrefP(zvalue)
	}
}
func ZvalPtrDtorStr(zval_ptr *Zval) {
	if zval_ptr.GetTypeFlags() != 0 && ZvalDelrefP(zval_ptr) == 0 {
		r.Assert(zval_ptr.GetType() == 6)
		r.Assert((ZvalGcFlags(zval_ptr.GetValue().GetStr().GetGc().GetTypeInfo()) & 1 << 6) == 0)
		r.Assert((ZvalGcFlags(zval_ptr.GetValue().GetStr().GetGc().GetTypeInfo()) & 1 << 7) == 0)
		_efree(zval_ptr.GetValue().GetStr())
	}
}

/* Kept for compatibility */

// #define zval_dtor(zvalue) zval_ptr_dtor_nogc ( zvalue )

// #define zval_internal_dtor(zvalue) zval_internal_ptr_dtor ( zvalue )

// #define zval_dtor_func       rc_dtor_func

// #define zval_ptr_dtor_wrapper       zval_ptr_dtor

// #define zval_internal_ptr_dtor_wrapper       zval_internal_ptr_dtor

// #define ZVAL_PTR_DTOR       zval_ptr_dtor

// #define ZVAL_INTERNAL_PTR_DTOR       zval_internal_ptr_dtor

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

// #define zend_string_destroy       _efree

type ZendRcDtorFuncT func(p *ZendRefcounted)

var ZendRcDtorFunc []ZendRcDtorFuncT = []ZendRcDtorFuncT{ZendRcDtorFuncT(ZendEmptyDestroy), ZendRcDtorFuncT(ZendEmptyDestroy), ZendRcDtorFuncT(ZendEmptyDestroy), ZendRcDtorFuncT(ZendEmptyDestroy), ZendRcDtorFuncT(ZendEmptyDestroy), ZendRcDtorFuncT(ZendEmptyDestroy), ZendRcDtorFuncT(_efree), ZendRcDtorFuncT(ZendArrayDestroy), ZendRcDtorFuncT(ZendObjectsStoreDel), ZendRcDtorFuncT(ZendListFree), ZendRcDtorFuncT(ZendReferenceDestroy), ZendRcDtorFuncT(ZendAstRefDestroy)}

func RcDtorFunc(p *ZendRefcounted) {
	r.Assert(ZvalGcType(p.GetGc().GetTypeInfo()) <= 11)
	ZendRcDtorFunc[ZvalGcType(p.GetGc().GetTypeInfo())](p)
}
func ZendReferenceDestroy(ref *ZendReference) {
	r.Assert(ref.GetSources().GetPtr() == nil)
	IZvalPtrDtor(&ref.val)
	_efree(ref)
}
func ZendEmptyDestroy(ref *ZendReference) {}
func ZvalPtrDtor(zval_ptr *Zval)          { IZvalPtrDtor(zval_ptr) }

/* }}} */

func ZvalInternalPtrDtor(zval_ptr *Zval) {
	if zval_ptr.GetTypeFlags() != 0 {
		var ref *ZendRefcounted = zval_ptr.GetValue().GetCounted()
		if ZendGcDelref(&ref.gc) == 0 {
			if zval_ptr.GetType() == 6 {
				var str *ZendString = (*ZendString)(ref)
				r.Assert((ZvalGcFlags(str.GetGc().GetTypeInfo()) & 1 << 6) == 0)
				r.Assert((ZvalGcFlags(str.GetGc().GetTypeInfo()) & 1 << 7) != 0)
				Free(str)
			} else {
				ZendErrorNoreturn(1<<4, "Internal zval's can't be arrays, objects, resources or reference")
			}
		}
	}
}

/* }}} */

func ZvalAddRef(p *Zval) {
	if p.GetTypeFlags() != 0 {
		if p.GetType() == 10 && ZvalRefcountP(p) == 1 {
			var _z1 *Zval = p
			var _z2 *Zval = &(*p).value.GetRef().GetVal()
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		} else {
			ZvalAddrefP(p)
		}
	}
}
func ZvalCopyCtorFunc(zvalue *Zval) {
	if zvalue.GetType() == 7 {
		var __arr *ZendArray = ZendArrayDup(zvalue.GetValue().GetArr())
		var __z *Zval = zvalue
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	} else if zvalue.GetType() == 6 {
		r.Assert((ZvalGcFlags(zvalue.GetValue().GetStr().GetGc().GetTypeInfo()) & 1 << 6) == 0)
		var __z *Zval = zvalue
		var __s *ZendString = ZendStringDup(zvalue.GetValue().GetStr(), 0)
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
}
