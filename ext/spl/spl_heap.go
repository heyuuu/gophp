// <<generate>>

package spl

import (
	"sik/core"
	r "sik/runtime"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/spl/spl_heap.h>

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
   | Authors: Etienne Kneuss <colder@php.net>                             |
   +----------------------------------------------------------------------+
*/

// #define SPL_HEAP_H

// # include "php.h"

// # include "php_spl.h"

var spl_ce_SplHeap *zend.ZendClassEntry
var spl_ce_SplMinHeap *zend.ZendClassEntry
var spl_ce_SplMaxHeap *zend.ZendClassEntry
var spl_ce_SplPriorityQueue *zend.ZendClassEntry

// Source: <ext/spl/spl_heap.c>

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
   | Authors: Etienne Kneuss <colder@php.net>                             |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "zend_exceptions.h"

// # include "php_spl.h"

// # include "spl_functions.h"

// # include "spl_engine.h"

// # include "spl_iterators.h"

// # include "spl_heap.h"

// # include "spl_exceptions.h"

// #define PTR_HEAP_BLOCK_SIZE       64

// #define SPL_HEAP_CORRUPTED       0x00000001

// #define SPL_PQUEUE_EXTR_MASK       0x00000003

// #define SPL_PQUEUE_EXTR_BOTH       0x00000003

// #define SPL_PQUEUE_EXTR_DATA       0x00000001

// #define SPL_PQUEUE_EXTR_PRIORITY       0x00000002

var spl_handler_SplHeap zend.ZendObjectHandlers
var spl_handler_SplPriorityQueue zend.ZendObjectHandlers

type SplPtrHeapDtorFunc func(any)
type SplPtrHeapCtorFunc func(any)
type SplPtrHeapCmpFunc func(any, any, *zend.Zval) int

/* define an __special__  overloaded iterator structure */

func SplHeapFromObj(obj *zend.ZendObject) *SplHeapObject {
	return (*SplHeapObject)((*byte)(obj - zend_long((*byte)(&((*SplHeapObject)(nil).GetStd()))-(*byte)(nil))))
}

/* }}} */

// #define Z_SPLHEAP_P(zv) spl_heap_from_obj ( Z_OBJ_P ( ( zv ) ) )

func SplHeapElem(heap *SplPtrHeap, i int) any {
	return any((*byte)(heap.GetElements() + heap.GetElemSize()*i))
}
func SplHeapElemCopy(heap *SplPtrHeap, to any, from any) {
	r.Assert(to != from)
	memcpy(to, from, heap.GetElemSize())
}
func SplPtrHeapZvalDtor(elem any) { zend.ZvalPtrDtor((*zend.Zval)(elem)) }

/* }}} */

func SplPtrHeapZvalCtor(elem any) {
	if (*zend.Zval)(elem).u1.v.type_flags != 0 {
		zend.ZvalAddrefP((*zend.Zval)(elem))
	}
}

/* }}} */

func SplPtrHeapPqueueElemDtor(elem any) {
	var pq_elem *SplPqueueElem = elem
	zend.ZvalPtrDtor(&pq_elem.data)
	zend.ZvalPtrDtor(&pq_elem.priority)
}

/* }}} */

func SplPtrHeapPqueueElemCtor(elem any) {
	var pq_elem *SplPqueueElem = elem
	if &pq_elem.data.u1.v.type_flags != 0 {
		zend.ZvalAddrefP(&pq_elem.data)
	}
	if &pq_elem.priority.u1.v.type_flags != 0 {
		zend.ZvalAddrefP(&pq_elem.priority)
	}
}

/* }}} */

func SplPtrHeapCmpCbHelper(object *zend.Zval, heap_object *SplHeapObject, a *zend.Zval, b *zend.Zval, result *zend.ZendLong) int {
	var zresult zend.Zval
	zend.ZendCallMethod(object, heap_object.std.ce, &heap_object.fptr_cmp, "compare", g.SizeOf("\"compare\"")-1, &zresult, 2, a, b)
	if zend.EG.exception != nil {
		return zend.FAILURE
	}
	*result = zend.ZvalGetLong(&zresult)
	zend.ZvalPtrDtor(&zresult)
	return zend.SUCCESS
}

/* }}} */

func SplPqueueExtractHelper(result *zend.Zval, elem *SplPqueueElem, flags int) {
	if (flags & 0x3) == 0x3 {
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = result
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		if &(elem.GetData()).u1.v.type_flags != 0 {
			zend.ZvalAddrefP(&(elem.GetData()))
		}
		zend.AddAssocZvalEx(result, "data", g.SizeOf("\"data\"")-1, &elem.data)
		if &(elem.GetPriority()).u1.v.type_flags != 0 {
			zend.ZvalAddrefP(&(elem.GetPriority()))
		}
		zend.AddAssocZvalEx(result, "priority", g.SizeOf("\"priority\"")-1, &elem.priority)
		return
	}
	if (flags & 0x1) != 0 {
		var _z1 *zend.Zval = result
		var _z2 *zend.Zval = &elem.data
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		if (_t & 0xff00) != 0 {
			zend.ZendGcAddref(&_gc.gc)
		}
		return
	}
	if (flags & 0x2) != 0 {
		var _z1 *zend.Zval = result
		var _z2 *zend.Zval = &elem.priority
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		if (_t & 0xff00) != 0 {
			zend.ZendGcAddref(&_gc.gc)
		}
		return
	}
	r.Assert(false)
}

/* }}} */

func SplPtrHeapZvalMaxCmp(x any, y any, object *zend.Zval) int {
	var a *zend.Zval = x
	var b *zend.Zval = y
	var result zend.Zval
	if zend.EG.exception != nil {
		return 0
	}
	if object != nil {
		var heap_object *SplHeapObject = SplHeapFromObj(object.value.obj)
		if heap_object.GetFptrCmp() != nil {
			var lval zend.ZendLong = 0
			if SplPtrHeapCmpCbHelper(object, heap_object, a, b, &lval) == zend.FAILURE {

				/* exception or call failure */

				return 0

				/* exception or call failure */

			}
			if lval != 0 {
				if lval < 0 {
					return -1
				} else {
					return 1
				}
			} else {
				return 0
			}
		}
	}
	zend.CompareFunction(&result, a, b)
	return int(result.value.lval)
}

/* }}} */

func SplPtrHeapZvalMinCmp(x any, y any, object *zend.Zval) int {
	var a *zend.Zval = x
	var b *zend.Zval = y
	var result zend.Zval
	if zend.EG.exception != nil {
		return 0
	}
	if object != nil {
		var heap_object *SplHeapObject = SplHeapFromObj(object.value.obj)
		if heap_object.GetFptrCmp() != nil {
			var lval zend.ZendLong = 0
			if SplPtrHeapCmpCbHelper(object, heap_object, a, b, &lval) == zend.FAILURE {

				/* exception or call failure */

				return 0

				/* exception or call failure */

			}
			if lval != 0 {
				if lval < 0 {
					return -1
				} else {
					return 1
				}
			} else {
				return 0
			}
		}
	}
	zend.CompareFunction(&result, b, a)
	return int(result.value.lval)
}

/* }}} */

func SplPtrPqueueElemCmp(x any, y any, object *zend.Zval) int {
	var a *SplPqueueElem = x
	var b *SplPqueueElem = y
	var a_priority_p *zend.Zval = &a.priority
	var b_priority_p *zend.Zval = &b.priority
	var result zend.Zval
	if zend.EG.exception != nil {
		return 0
	}
	if object != nil {
		var heap_object *SplHeapObject = SplHeapFromObj(object.value.obj)
		if heap_object.GetFptrCmp() != nil {
			var lval zend.ZendLong = 0
			if SplPtrHeapCmpCbHelper(object, heap_object, a_priority_p, b_priority_p, &lval) == zend.FAILURE {

				/* exception or call failure */

				return 0

				/* exception or call failure */

			}
			if lval != 0 {
				if lval < 0 {
					return -1
				} else {
					return 1
				}
			} else {
				return 0
			}
		}
	}
	zend.CompareFunction(&result, a_priority_p, b_priority_p)
	return int(result.value.lval)
}

/* }}} */

func SplPtrHeapInit(cmp SplPtrHeapCmpFunc, ctor SplPtrHeapCtorFunc, dtor SplPtrHeapDtorFunc, elem_size int) *SplPtrHeap {
	var heap *SplPtrHeap = zend._emalloc(g.SizeOf("spl_ptr_heap"))
	heap.SetDtor(dtor)
	heap.SetCtor(ctor)
	heap.SetCmp(cmp)
	heap.SetElements(zend._ecalloc(64, elem_size))
	heap.SetMaxSize(64)
	heap.SetCount(0)
	heap.SetFlags(0)
	heap.SetElemSize(elem_size)
	return heap
}

/* }}} */

func SplPtrHeapInsert(heap *SplPtrHeap, elem any, cmp_userdata any) {
	var i int
	if heap.GetCount()+1 > heap.GetMaxSize() {
		var alloc_size int = heap.GetMaxSize() * heap.GetElemSize()

		/* we need to allocate more memory */

		heap.SetElements(zend._erealloc(heap.GetElements(), 2*alloc_size))
		memset((*byte)(heap.GetElements()+alloc_size), 0, alloc_size)
		heap.SetMaxSize(heap.GetMaxSize() * 2)
	}

	/* sifting up */

	for i = heap.GetCount(); i > 0 && heap.GetCmp()(SplHeapElem(heap, (i-1)/2), elem, cmp_userdata) < 0; i = (i - 1) / 2 {
		SplHeapElemCopy(heap, SplHeapElem(heap, i), SplHeapElem(heap, (i-1)/2))
	}
	heap.GetCount()++
	if zend.EG.exception != nil {

		/* exception thrown during comparison */

		heap.SetFlags(heap.GetFlags() | 0x1)

		/* exception thrown during comparison */

	}
	SplHeapElemCopy(heap, SplHeapElem(heap, i), elem)
}

/* }}} */

func SplPtrHeapTop(heap *SplPtrHeap) any {
	if heap.GetCount() == 0 {
		return nil
	}
	return heap.GetElements()
}

/* }}} */

func SplPtrHeapDeleteTop(heap *SplPtrHeap, elem any, cmp_userdata any) int {
	var i int
	var j int
	var limit int = (heap.GetCount() - 1) / 2
	var bottom any
	if heap.GetCount() == 0 {
		return zend.FAILURE
	}
	if elem {
		SplHeapElemCopy(heap, elem, SplHeapElem(heap, 0))
	} else {
		heap.GetDtor()(SplHeapElem(heap, 0))
	}
	bottom = SplHeapElem(heap, g.PreDec(&(heap.GetCount())))
	for i = 0; i < limit; i = j {

		/* Find smaller child */

		j = i*2 + 1
		if j != heap.GetCount() && heap.GetCmp()(SplHeapElem(heap, j+1), SplHeapElem(heap, j), cmp_userdata) > 0 {
			j++
		}

		/* swap elements between two levels */

		if heap.GetCmp()(bottom, SplHeapElem(heap, j), cmp_userdata) < 0 {
			SplHeapElemCopy(heap, SplHeapElem(heap, i), SplHeapElem(heap, j))
		} else {
			break
		}

		/* swap elements between two levels */

	}
	if zend.EG.exception != nil {

		/* exception thrown during comparison */

		heap.SetFlags(heap.GetFlags() | 0x1)

		/* exception thrown during comparison */

	}
	var to any = SplHeapElem(heap, i)
	if to != bottom {
		SplHeapElemCopy(heap, to, bottom)
	}
	return zend.SUCCESS
}

/* }}} */

func SplPtrHeapClone(from *SplPtrHeap) *SplPtrHeap {
	var i int
	var heap *SplPtrHeap = zend._emalloc(g.SizeOf("spl_ptr_heap"))
	heap.SetDtor(from.GetDtor())
	heap.SetCtor(from.GetCtor())
	heap.SetCmp(from.GetCmp())
	heap.SetMaxSize(from.GetMaxSize())
	heap.SetCount(from.GetCount())
	heap.SetFlags(from.GetFlags())
	heap.SetElemSize(from.GetElemSize())
	heap.SetElements(zend._safeEmalloc(from.GetElemSize(), from.GetMaxSize(), 0))
	memcpy(heap.GetElements(), from.GetElements(), from.GetElemSize()*from.GetMaxSize())
	for i = 0; i < heap.GetCount(); i++ {
		heap.GetCtor()(SplHeapElem(heap, i))
	}
	return heap
}

/* }}} */

func SplPtrHeapDestroy(heap *SplPtrHeap) {
	var i int
	for i = 0; i < heap.GetCount(); i++ {
		heap.GetDtor()(SplHeapElem(heap, i))
	}
	zend._efree(heap.GetElements())
	zend._efree(heap)
}

/* }}} */

func SplPtrHeapCount(heap *SplPtrHeap) int { return heap.GetCount() }

/* }}} */

func SplHeapObjectFreeStorage(object *zend.ZendObject) {
	var intern *SplHeapObject = SplHeapFromObj(object)
	zend.ZendObjectStdDtor(&intern.std)
	SplPtrHeapDestroy(intern.GetHeap())
}

/* }}} */

func SplHeapObjectNewEx(class_type *zend.ZendClassEntry, orig *zend.Zval, clone_orig int) *zend.ZendObject {
	var intern *SplHeapObject
	var parent *zend.ZendClassEntry = class_type
	var inherited int = 0
	intern = zend.ZendObjectAlloc(g.SizeOf("spl_heap_object"), parent)
	zend.ZendObjectStdInit(&intern.std, class_type)
	zend.ObjectPropertiesInit(&intern.std, class_type)
	if orig != nil {
		var other *SplHeapObject = SplHeapFromObj(orig.value.obj)
		intern.std.handlers = other.std.handlers
		intern.SetCeGetIterator(other.GetCeGetIterator())
		if clone_orig != 0 {
			intern.SetHeap(SplPtrHeapClone(other.GetHeap()))
		} else {
			intern.SetHeap(other.GetHeap())
		}
		intern.SetFlags(other.GetFlags())
		intern.SetFptrCmp(other.GetFptrCmp())
		intern.SetFptrCount(other.GetFptrCount())
		return &intern.std
	}
	for parent != nil {
		if parent == spl_ce_SplPriorityQueue {
			intern.SetHeap(SplPtrHeapInit(SplPtrPqueueElemCmp, SplPtrHeapPqueueElemCtor, SplPtrHeapPqueueElemDtor, g.SizeOf("spl_pqueue_elem")))
			intern.std.handlers = &spl_handler_SplPriorityQueue
			intern.SetFlags(0x1)
			break
		}
		if parent == spl_ce_SplMinHeap || parent == spl_ce_SplMaxHeap || parent == spl_ce_SplHeap {
			intern.SetHeap(SplPtrHeapInit(g.Cond(parent == spl_ce_SplMinHeap, SplPtrHeapZvalMinCmp, SplPtrHeapZvalMaxCmp), SplPtrHeapZvalCtor, SplPtrHeapZvalDtor, g.SizeOf("zval")))
			intern.std.handlers = &spl_handler_SplHeap
			break
		}
		parent = parent.parent
		inherited = 1
	}
	if parent == nil {
		core.PhpErrorDocref(nil, 1<<6, "Internal compiler error, Class is not child of SplHeap")
	}
	if inherited != 0 {
		intern.SetFptrCmp(zend.ZendHashStrFindPtr(&class_type.function_table, "compare", g.SizeOf("\"compare\"")-1))
		if intern.GetFptrCmp().common.scope == parent {
			intern.SetFptrCmp(nil)
		}
		intern.SetFptrCount(zend.ZendHashStrFindPtr(&class_type.function_table, "count", g.SizeOf("\"count\"")-1))
		if intern.GetFptrCount().common.scope == parent {
			intern.SetFptrCount(nil)
		}
	}
	return &intern.std
}

/* }}} */

func SplHeapObjectNew(class_type *zend.ZendClassEntry) *zend.ZendObject {
	return SplHeapObjectNewEx(class_type, nil, 0)
}

/* }}} */

func SplHeapObjectClone(zobject *zend.Zval) *zend.ZendObject {
	var old_object *zend.ZendObject
	var new_object *zend.ZendObject
	old_object = zobject.value.obj
	new_object = SplHeapObjectNewEx(old_object.ce, zobject, 1)
	zend.ZendObjectsCloneMembers(new_object, old_object)
	return new_object
}

/* }}} */

func SplHeapObjectCountElements(object *zend.Zval, count *zend.ZendLong) int {
	var intern *SplHeapObject = SplHeapFromObj(object.value.obj)
	if intern.GetFptrCount() != nil {
		var rv zend.Zval
		zend.ZendCallMethod(object, intern.std.ce, &intern.fptr_count, "count", g.SizeOf("\"count\"")-1, &rv, 0, nil, nil)
		if rv.u1.v.type_ != 0 {
			*count = zend.ZvalGetLong(&rv)
			zend.ZvalPtrDtor(&rv)
			return zend.SUCCESS
		}
		*count = 0
		return zend.FAILURE
	}
	*count = SplPtrHeapCount(intern.GetHeap())
	return zend.SUCCESS
}

/* }}} */

func SplHeapObjectGetDebugInfo(ce *zend.ZendClassEntry, obj *zend.Zval) *zend.HashTable {
	var intern *SplHeapObject = SplHeapFromObj(obj.value.obj)
	var tmp zend.Zval
	var heap_array zend.Zval
	var pnstr *zend.ZendString
	var debug_info *zend.HashTable
	var i int
	if intern.std.properties == nil {
		zend.RebuildObjectProperties(&intern.std)
	}
	debug_info = zend._zendNewArray(intern.std.properties.nNumOfElements + 1)
	zend.ZendHashCopy(debug_info, intern.std.properties, zend.CopyCtorFuncT(zend.ZvalAddRef))
	pnstr = SplGenPrivatePropName(ce, "flags", g.SizeOf("\"flags\"")-1)
	var __z *zend.Zval = &tmp
	__z.value.lval = intern.GetFlags()
	__z.u1.type_info = 4
	zend.ZendHashUpdate(debug_info, pnstr, &tmp)
	zend.ZendStringReleaseEx(pnstr, 0)
	pnstr = SplGenPrivatePropName(ce, "isCorrupted", g.SizeOf("\"isCorrupted\"")-1)
	if (intern.GetHeap().GetFlags() & 0x1) != 0 {
		&tmp.u1.type_info = 3
	} else {
		&tmp.u1.type_info = 2
	}
	zend.ZendHashUpdate(debug_info, pnstr, &tmp)
	zend.ZendStringReleaseEx(pnstr, 0)
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = &heap_array
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	for i = 0; i < intern.GetHeap().GetCount(); i++ {
		if ce == spl_ce_SplPriorityQueue {
			var pq_elem *SplPqueueElem = SplHeapElem(intern.GetHeap(), i)
			var elem zend.Zval
			SplPqueueExtractHelper(&elem, pq_elem, 0x3)
			zend.AddIndexZval(&heap_array, i, &elem)
		} else {
			var elem *zend.Zval = SplHeapElem(intern.GetHeap(), i)
			zend.AddIndexZval(&heap_array, i, elem)
			if elem.u1.v.type_flags != 0 {
				zend.ZvalAddrefP(elem)
			}
		}
	}
	pnstr = SplGenPrivatePropName(ce, "heap", g.SizeOf("\"heap\"")-1)
	zend.ZendHashUpdate(debug_info, pnstr, &heap_array)
	zend.ZendStringReleaseEx(pnstr, 0)
	return debug_info
}

/* }}} */

func SplHeapObjectGetGc(obj *zend.Zval, gc_data **zend.Zval, gc_data_count *int) *zend.HashTable {
	var intern *SplHeapObject = SplHeapFromObj(obj.value.obj)
	*gc_data = (*zend.Zval)(intern.GetHeap().GetElements())
	*gc_data_count = intern.GetHeap().GetCount()
	return zend.ZendStdGetProperties(obj)
}

/* }}} */

func SplPqueueObjectGetGc(obj *zend.Zval, gc_data **zend.Zval, gc_data_count *int) *zend.HashTable {
	var intern *SplHeapObject = SplHeapFromObj(obj.value.obj)
	*gc_data = (*zend.Zval)(intern.GetHeap().GetElements())

	/* Two zvals (value and priority) per pqueue entry */

	*gc_data_count = 2 * intern.GetHeap().GetCount()
	return zend.ZendStdGetProperties(obj)
}

/* }}} */

func zim_spl_SplHeap_count(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var count zend.ZendLong
	var intern *SplHeapObject = SplHeapFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	count = SplPtrHeapCount(intern.GetHeap())
	var __z *zend.Zval = return_value
	__z.value.lval = count
	__z.u1.type_info = 4
	return
}

/* }}} */

func zim_spl_SplHeap_isEmpty(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplHeapObject = SplHeapFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if SplPtrHeapCount(intern.GetHeap()) == 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func zim_spl_SplHeap_insert(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var value *zend.Zval
	var intern *SplHeapObject
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "z", &value) == zend.FAILURE {
		return
	}
	intern = SplHeapFromObj(&(execute_data.This).value.obj)
	if (intern.GetHeap().GetFlags() & 0x1) != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "Heap is corrupted, heap properties are no longer ensured.", 0)
		return
	}
	if value.u1.v.type_flags != 0 {
		zend.ZvalAddrefP(value)
	}
	SplPtrHeapInsert(intern.GetHeap(), value, &(execute_data.This))
	return_value.u1.type_info = 3
	return
}

/* }}} */

func zim_spl_SplHeap_extract(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplHeapObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	intern = SplHeapFromObj(&(execute_data.This).value.obj)
	if (intern.GetHeap().GetFlags() & 0x1) != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "Heap is corrupted, heap properties are no longer ensured.", 0)
		return
	}
	if SplPtrHeapDeleteTop(intern.GetHeap(), return_value, &(execute_data.This)) == zend.FAILURE {
		zend.ZendThrowException(spl_ce_RuntimeException, "Can't extract from an empty heap", 0)
		return
	}
}

/* }}} */

func zim_spl_SplPriorityQueue_insert(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var data *zend.Zval
	var priority *zend.Zval
	var intern *SplHeapObject
	var elem SplPqueueElem
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "zz", &data, &priority) == zend.FAILURE {
		return
	}
	intern = SplHeapFromObj(&(execute_data.This).value.obj)
	if (intern.GetHeap().GetFlags() & 0x1) != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "Heap is corrupted, heap properties are no longer ensured.", 0)
		return
	}
	var _z1 *zend.Zval = &elem.data
	var _z2 *zend.Zval = data
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
	if (_t & 0xff00) != 0 {
		zend.ZendGcAddref(&_gc.gc)
	}
	var _z1 *zend.Zval = &elem.priority
	var _z2 *zend.Zval = priority
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
	if (_t & 0xff00) != 0 {
		zend.ZendGcAddref(&_gc.gc)
	}
	SplPtrHeapInsert(intern.GetHeap(), &elem, &(execute_data.This))
	return_value.u1.type_info = 3
	return
}

/* }}} */

func zim_spl_SplPriorityQueue_extract(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var elem SplPqueueElem
	var intern *SplHeapObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	intern = SplHeapFromObj(&(execute_data.This).value.obj)
	if (intern.GetHeap().GetFlags() & 0x1) != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "Heap is corrupted, heap properties are no longer ensured.", 0)
		return
	}
	if SplPtrHeapDeleteTop(intern.GetHeap(), &elem, &(execute_data.This)) == zend.FAILURE {
		zend.ZendThrowException(spl_ce_RuntimeException, "Can't extract from an empty heap", 0)
		return
	}
	SplPqueueExtractHelper(return_value, &elem, intern.GetFlags())
	SplPtrHeapPqueueElemDtor(&elem)
}

/* }}} */

func zim_spl_SplPriorityQueue_top(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplHeapObject
	var elem *SplPqueueElem
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	intern = SplHeapFromObj(&(execute_data.This).value.obj)
	if (intern.GetHeap().GetFlags() & 0x1) != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "Heap is corrupted, heap properties are no longer ensured.", 0)
		return
	}
	elem = SplPtrHeapTop(intern.GetHeap())
	if elem == nil {
		zend.ZendThrowException(spl_ce_RuntimeException, "Can't peek at an empty heap", 0)
		return
	}
	SplPqueueExtractHelper(return_value, elem, intern.GetFlags())
}

/* }}} */

func zim_spl_SplPriorityQueue_setExtractFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var value zend.ZendLong
	var intern *SplHeapObject
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "l", &value) == zend.FAILURE {
		return
	}
	value &= 0x3
	if value == 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "Must specify at least one extract flag", 0)
		return
	}
	intern = SplHeapFromObj(&(execute_data.This).value.obj)
	intern.SetFlags(value)
	var __z *zend.Zval = return_value
	__z.value.lval = intern.GetFlags()
	__z.u1.type_info = 4
	return
}

/* }}} */

func zim_spl_SplPriorityQueue_getExtractFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplHeapObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	intern = SplHeapFromObj(&(execute_data.This).value.obj)
	var __z *zend.Zval = return_value
	__z.value.lval = intern.GetFlags()
	__z.u1.type_info = 4
	return
}

/* }}} */

func zim_spl_SplHeap_recoverFromCorruption(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplHeapObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	intern = SplHeapFromObj(&(execute_data.This).value.obj)
	intern.GetHeap().SetFlags(intern.GetHeap().GetFlags() & ^0x1)
	return_value.u1.type_info = 3
	return
}

/* }}} */

func zim_spl_SplHeap_isCorrupted(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplHeapObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	intern = SplHeapFromObj(&(execute_data.This).value.obj)
	if (intern.GetHeap().GetFlags() & 0x1) != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func zim_spl_SplPriorityQueue_compare(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var a *zend.Zval
	var b *zend.Zval
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "zz", &a, &b) == zend.FAILURE {
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = SplPtrHeapZvalMaxCmp(a, b, nil)
	__z.u1.type_info = 4
	return
}

/* }}} */

func zim_spl_SplHeap_top(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var value *zend.Zval
	var intern *SplHeapObject
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	intern = SplHeapFromObj(&(execute_data.This).value.obj)
	if (intern.GetHeap().GetFlags() & 0x1) != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "Heap is corrupted, heap properties are no longer ensured.", 0)
		return
	}
	value = SplPtrHeapTop(intern.GetHeap())
	if value == nil {
		zend.ZendThrowException(spl_ce_RuntimeException, "Can't peek at an empty heap", 0)
		return
	}
	var _z3 *zend.Zval = value
	if (_z3.u1.type_info & 0xff00) != 0 {
		if (_z3.u1.type_info & 0xff) == 10 {
			_z3 = &(*_z3).value.ref.val
			if (_z3.u1.type_info & 0xff00) != 0 {
				zend.ZvalAddrefP(_z3)
			}
		} else {
			zend.ZvalAddrefP(_z3)
		}
	}
	var _z1 *zend.Zval = return_value
	var _z2 *zend.Zval = _z3
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
}

/* }}} */

func zim_spl_SplMinHeap_compare(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var a *zend.Zval
	var b *zend.Zval
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "zz", &a, &b) == zend.FAILURE {
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = SplPtrHeapZvalMinCmp(a, b, nil)
	__z.u1.type_info = 4
	return
}

/* }}} */

func zim_spl_SplMaxHeap_compare(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var a *zend.Zval
	var b *zend.Zval
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "zz", &a, &b) == zend.FAILURE {
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = SplPtrHeapZvalMaxCmp(a, b, nil)
	__z.u1.type_info = 4
	return
}

/* }}} */

func SplHeapItDtor(iter *zend.ZendObjectIterator) {
	var iterator *SplHeapIt = (*SplHeapIt)(iter)
	zend.ZendUserItInvalidateCurrent(iter)
	zend.ZvalPtrDtor(&iterator.intern.it.data)
}

/* }}} */

func SplHeapItRewind(iter *zend.ZendObjectIterator) {}

/* }}} */

func SplHeapItValid(iter *zend.ZendObjectIterator) int {
	if SplHeapFromObj(&iter.data.value.obj).GetHeap().GetCount() != 0 {
		return zend.SUCCESS
	} else {
		return zend.FAILURE
	}
}

/* }}} */

func SplHeapItGetCurrentData(iter *zend.ZendObjectIterator) *zend.Zval {
	var object *SplHeapObject = SplHeapFromObj(&iter.data.value.obj)
	if (object.GetHeap().GetFlags() & 0x1) != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "Heap is corrupted, heap properties are no longer ensured.", 0)
		return nil
	}
	if object.GetHeap().GetCount() == 0 {
		return nil
	} else {
		return SplHeapElem(object.GetHeap(), 0)
	}
}

/* }}} */

func SplPqueueItGetCurrentData(iter *zend.ZendObjectIterator) *zend.Zval {
	var user_it *zend.ZendUserIterator = (*zend.ZendUserIterator)(iter)
	var object *SplHeapObject = SplHeapFromObj(&iter.data.value.obj)
	if (object.GetHeap().GetFlags() & 0x1) != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "Heap is corrupted, heap properties are no longer ensured.", 0)
		return nil
	}
	if object.GetHeap().GetCount() == 0 {
		return nil
	}
	if user_it.value.u1.v.type_ == 0 {
		var elem *SplPqueueElem = SplHeapElem(object.GetHeap(), 0)
		SplPqueueExtractHelper(&user_it.value, elem, object.GetFlags())
	}
	return &user_it.value
}

/* }}} */

func SplHeapItGetCurrentKey(iter *zend.ZendObjectIterator, key *zend.Zval) {
	var object *SplHeapObject = SplHeapFromObj(&iter.data.value.obj)
	var __z *zend.Zval = key
	__z.value.lval = object.GetHeap().GetCount() - 1
	__z.u1.type_info = 4
}

/* }}} */

func SplHeapItMoveForward(iter *zend.ZendObjectIterator) {
	var object *SplHeapObject = SplHeapFromObj(&iter.data.value.obj)
	if (object.GetHeap().GetFlags() & 0x1) != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "Heap is corrupted, heap properties are no longer ensured.", 0)
		return
	}
	SplPtrHeapDeleteTop(object.GetHeap(), nil, &iter.data)
	zend.ZendUserItInvalidateCurrent(iter)
}

/* }}} */

func zim_spl_SplHeap_key(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplHeapObject = SplHeapFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = intern.GetHeap().GetCount() - 1
	__z.u1.type_info = 4
	return
}

/* }}} */

func zim_spl_SplHeap_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplHeapObject = SplHeapFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	SplPtrHeapDeleteTop(intern.GetHeap(), nil, &(execute_data.This))
}

/* }}} */

func zim_spl_SplHeap_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplHeapObject = SplHeapFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if intern.GetHeap().GetCount() != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func zim_spl_SplHeap_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
}

/* }}} */

func zim_spl_SplHeap_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplHeapObject = SplHeapFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if intern.GetHeap().GetCount() == 0 {
		return_value.u1.type_info = 1
		return
	} else {
		var element *zend.Zval = SplHeapElem(intern.GetHeap(), 0)
		var _z3 *zend.Zval = element
		if (_z3.u1.type_info & 0xff00) != 0 {
			if (_z3.u1.type_info & 0xff) == 10 {
				_z3 = &(*_z3).value.ref.val
				if (_z3.u1.type_info & 0xff00) != 0 {
					zend.ZvalAddrefP(_z3)
				}
			} else {
				zend.ZvalAddrefP(_z3)
			}
		}
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = _z3
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
	}
}

/* }}} */

func zim_spl_SplPriorityQueue_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplHeapObject = SplHeapFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if intern.GetHeap().GetCount() == 0 {
		return_value.u1.type_info = 1
		return
	} else {
		var elem *SplPqueueElem = SplHeapElem(intern.GetHeap(), 0)
		SplPqueueExtractHelper(return_value, elem, intern.GetFlags())
	}
}

/* }}} */

func zim_spl_SplHeap___debugInfo(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __arr *zend.ZendArray = SplHeapObjectGetDebugInfo(spl_ce_SplHeap, g.CondF1(&(execute_data.This).u1.v.type_ == 8, func() *zend.Zval { return &(execute_data.This) }, nil))
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	return
}

/* {{{ proto void SplPriorityQueue::__debugInfo() */

func zim_spl_SplPriorityQueue___debugInfo(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __arr *zend.ZendArray = SplHeapObjectGetDebugInfo(spl_ce_SplPriorityQueue, g.CondF1(&(execute_data.This).u1.v.type_ == 8, func() *zend.Zval { return &(execute_data.This) }, nil))
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	return
}

/* iterator handler table */

var SplHeapItFuncs zend.ZendObjectIteratorFuncs = zend.ZendObjectIteratorFuncs{SplHeapItDtor, SplHeapItValid, SplHeapItGetCurrentData, SplHeapItGetCurrentKey, SplHeapItMoveForward, SplHeapItRewind, nil}
var SplPqueueItFuncs zend.ZendObjectIteratorFuncs = zend.ZendObjectIteratorFuncs{SplHeapItDtor, SplHeapItValid, SplPqueueItGetCurrentData, SplHeapItGetCurrentKey, SplHeapItMoveForward, SplHeapItRewind, nil}

func SplHeapGetIterator(ce *zend.ZendClassEntry, object *zend.Zval, by_ref int) *zend.ZendObjectIterator {
	var iterator *SplHeapIt
	var heap_object *SplHeapObject = SplHeapFromObj(object.value.obj)
	if by_ref != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "An iterator cannot be used with foreach by reference", 0)
		return nil
	}
	iterator = zend._emalloc(g.SizeOf("spl_heap_it"))
	zend.ZendIteratorInit(&iterator.intern.it)
	zend.ZvalAddrefP(object)
	var __z *zend.Zval = &iterator.intern.it.data
	__z.value.obj = object.value.obj
	__z.u1.type_info = 8 | 1<<0<<8 | 1<<1<<8
	iterator.intern.it.funcs = &SplHeapItFuncs
	iterator.intern.ce = ce
	iterator.SetFlags(heap_object.GetFlags())
	&iterator.intern.value.u1.type_info = 0
	return &iterator.intern.it
}

/* }}} */

func SplPqueueGetIterator(ce *zend.ZendClassEntry, object *zend.Zval, by_ref int) *zend.ZendObjectIterator {
	var iterator *SplHeapIt
	var heap_object *SplHeapObject = SplHeapFromObj(object.value.obj)
	if by_ref != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "An iterator cannot be used with foreach by reference", 0)
		return nil
	}
	iterator = zend._emalloc(g.SizeOf("spl_heap_it"))
	zend.ZendIteratorInit((*zend.ZendObjectIterator)(iterator))
	zend.ZvalAddrefP(object)
	var __z *zend.Zval = &iterator.intern.it.data
	__z.value.obj = object.value.obj
	__z.u1.type_info = 8 | 1<<0<<8 | 1<<1<<8
	iterator.intern.it.funcs = &SplPqueueItFuncs
	iterator.intern.ce = ce
	iterator.SetFlags(heap_object.GetFlags())
	&iterator.intern.value.u1.type_info = 0
	return &iterator.intern.it
}

/* }}} */

var ArginfoHeapInsert []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"value", 0, 0, 0}}
var ArginfoHeapCompare []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"value1", 0, 0, 0}, {"value2", 0, 0, 0}}
var ArginfoPqueueInsert []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"value", 0, 0, 0}, {"priority", 0, 0, 0}}
var ArginfoPqueueSetflags []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"flags", 0, 0, 0}}
var ArginfoSplheapVoid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var spl_funcs_SplMinHeap []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"compare",
		zim_spl_SplMinHeap_compare,
		ArginfoHeapCompare,
		uint32(g.SizeOf("arginfo_heap_compare")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 1,
	},
	{nil, nil, nil, 0, 0},
}
var spl_funcs_SplMaxHeap []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"compare",
		zim_spl_SplMaxHeap_compare,
		ArginfoHeapCompare,
		uint32(g.SizeOf("arginfo_heap_compare")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 1,
	},
	{nil, nil, nil, 0, 0},
}
var spl_funcs_SplPriorityQueue []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"compare",
		zim_spl_SplPriorityQueue_compare,
		ArginfoHeapCompare,
		uint32(g.SizeOf("arginfo_heap_compare")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"insert",
		zim_spl_SplPriorityQueue_insert,
		ArginfoPqueueInsert,
		uint32(g.SizeOf("arginfo_pqueue_insert")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"setExtractFlags",
		zim_spl_SplPriorityQueue_setExtractFlags,
		ArginfoPqueueSetflags,
		uint32(g.SizeOf("arginfo_pqueue_setflags")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getExtractFlags",
		zim_spl_SplPriorityQueue_getExtractFlags,
		ArginfoSplheapVoid,
		uint32(g.SizeOf("arginfo_splheap_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"top",
		zim_spl_SplPriorityQueue_top,
		ArginfoSplheapVoid,
		uint32(g.SizeOf("arginfo_splheap_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"extract",
		zim_spl_SplPriorityQueue_extract,
		ArginfoSplheapVoid,
		uint32(g.SizeOf("arginfo_splheap_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"count",
		zim_spl_SplHeap_count,
		ArginfoSplheapVoid,
		uint32(g.SizeOf("arginfo_splheap_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"isEmpty",
		zim_spl_SplHeap_isEmpty,
		ArginfoSplheapVoid,
		uint32(g.SizeOf("arginfo_splheap_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"rewind",
		zim_spl_SplHeap_rewind,
		ArginfoSplheapVoid,
		uint32(g.SizeOf("arginfo_splheap_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"current",
		zim_spl_SplPriorityQueue_current,
		ArginfoSplheapVoid,
		uint32(g.SizeOf("arginfo_splheap_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"key",
		zim_spl_SplHeap_key,
		ArginfoSplheapVoid,
		uint32(g.SizeOf("arginfo_splheap_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"next",
		zim_spl_SplHeap_next,
		ArginfoSplheapVoid,
		uint32(g.SizeOf("arginfo_splheap_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"valid",
		zim_spl_SplHeap_valid,
		ArginfoSplheapVoid,
		uint32(g.SizeOf("arginfo_splheap_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"recoverFromCorruption",
		zim_spl_SplHeap_recoverFromCorruption,
		ArginfoSplheapVoid,
		uint32(g.SizeOf("arginfo_splheap_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"isCorrupted",
		zim_spl_SplHeap_isCorrupted,
		ArginfoSplheapVoid,
		uint32(g.SizeOf("arginfo_splheap_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"__debugInfo",
		zim_spl_SplPriorityQueue___debugInfo,
		ArginfoSplheapVoid,
		uint32(g.SizeOf("arginfo_splheap_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}
var spl_funcs_SplHeap []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"extract",
		zim_spl_SplHeap_extract,
		ArginfoSplheapVoid,
		uint32(g.SizeOf("arginfo_splheap_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"insert",
		zim_spl_SplHeap_insert,
		ArginfoHeapInsert,
		uint32(g.SizeOf("arginfo_heap_insert")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"top",
		zim_spl_SplHeap_top,
		ArginfoSplheapVoid,
		uint32(g.SizeOf("arginfo_splheap_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"count",
		zim_spl_SplHeap_count,
		ArginfoSplheapVoid,
		uint32(g.SizeOf("arginfo_splheap_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"isEmpty",
		zim_spl_SplHeap_isEmpty,
		ArginfoSplheapVoid,
		uint32(g.SizeOf("arginfo_splheap_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"rewind",
		zim_spl_SplHeap_rewind,
		ArginfoSplheapVoid,
		uint32(g.SizeOf("arginfo_splheap_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"current",
		zim_spl_SplHeap_current,
		ArginfoSplheapVoid,
		uint32(g.SizeOf("arginfo_splheap_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"key",
		zim_spl_SplHeap_key,
		ArginfoSplheapVoid,
		uint32(g.SizeOf("arginfo_splheap_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"next",
		zim_spl_SplHeap_next,
		ArginfoSplheapVoid,
		uint32(g.SizeOf("arginfo_splheap_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"valid",
		zim_spl_SplHeap_valid,
		ArginfoSplheapVoid,
		uint32(g.SizeOf("arginfo_splheap_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"recoverFromCorruption",
		zim_spl_SplHeap_recoverFromCorruption,
		ArginfoSplheapVoid,
		uint32(g.SizeOf("arginfo_splheap_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"isCorrupted",
		zim_spl_SplHeap_isCorrupted,
		ArginfoSplheapVoid,
		uint32(g.SizeOf("arginfo_splheap_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"__debugInfo",
		zim_spl_SplHeap___debugInfo,
		ArginfoSplheapVoid,
		uint32(g.SizeOf("arginfo_splheap_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"compare",
		nil,
		nil,
		uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1<<1 | 1<<6,
	},
	{nil, nil, nil, 0, 0},
}

/* }}} */

func ZmStartupSplHeap(type_ int, module_number int) int {
	SplRegisterStdClass(&spl_ce_SplHeap, "SplHeap", SplHeapObjectNew, spl_funcs_SplHeap)
	memcpy(&spl_handler_SplHeap, &zend.StdObjectHandlers, g.SizeOf("zend_object_handlers"))
	spl_handler_SplHeap.offset = zend_long((*byte)(&((*SplHeapObject)(nil).GetStd())) - (*byte)(nil))
	spl_handler_SplHeap.clone_obj = SplHeapObjectClone
	spl_handler_SplHeap.count_elements = SplHeapObjectCountElements
	spl_handler_SplHeap.get_gc = SplHeapObjectGetGc
	spl_handler_SplHeap.dtor_obj = zend.ZendObjectsDestroyObject
	spl_handler_SplHeap.free_obj = SplHeapObjectFreeStorage
	zend.ZendClassImplements(spl_ce_SplHeap, 1, zend.ZendCeIterator)
	zend.ZendClassImplements(spl_ce_SplHeap, 1, zend.ZendCeCountable)
	spl_ce_SplHeap.get_iterator = SplHeapGetIterator
	SplRegisterSubClass(&spl_ce_SplMinHeap, spl_ce_SplHeap, "SplMinHeap", SplHeapObjectNew, spl_funcs_SplMinHeap)
	SplRegisterSubClass(&spl_ce_SplMaxHeap, spl_ce_SplHeap, "SplMaxHeap", SplHeapObjectNew, spl_funcs_SplMaxHeap)
	spl_ce_SplMaxHeap.get_iterator = SplHeapGetIterator
	spl_ce_SplMinHeap.get_iterator = SplHeapGetIterator
	SplRegisterStdClass(&spl_ce_SplPriorityQueue, "SplPriorityQueue", SplHeapObjectNew, spl_funcs_SplPriorityQueue)
	memcpy(&spl_handler_SplPriorityQueue, &zend.StdObjectHandlers, g.SizeOf("zend_object_handlers"))
	spl_handler_SplPriorityQueue.offset = zend_long((*byte)(&((*SplHeapObject)(nil).GetStd())) - (*byte)(nil))
	spl_handler_SplPriorityQueue.clone_obj = SplHeapObjectClone
	spl_handler_SplPriorityQueue.count_elements = SplHeapObjectCountElements
	spl_handler_SplPriorityQueue.get_gc = SplPqueueObjectGetGc
	spl_handler_SplPriorityQueue.dtor_obj = zend.ZendObjectsDestroyObject
	spl_handler_SplPriorityQueue.free_obj = SplHeapObjectFreeStorage
	zend.ZendClassImplements(spl_ce_SplPriorityQueue, 1, zend.ZendCeIterator)
	zend.ZendClassImplements(spl_ce_SplPriorityQueue, 1, zend.ZendCeCountable)
	spl_ce_SplPriorityQueue.get_iterator = SplPqueueGetIterator
	zend.ZendDeclareClassConstantLong(spl_ce_SplPriorityQueue, "EXTR_BOTH", g.SizeOf("\"EXTR_BOTH\"")-1, zend.ZendLong(0x3))
	zend.ZendDeclareClassConstantLong(spl_ce_SplPriorityQueue, "EXTR_PRIORITY", g.SizeOf("\"EXTR_PRIORITY\"")-1, zend.ZendLong(0x2))
	zend.ZendDeclareClassConstantLong(spl_ce_SplPriorityQueue, "EXTR_DATA", g.SizeOf("\"EXTR_DATA\"")-1, zend.ZendLong(0x1))
	return zend.SUCCESS
}

/* }}} */
