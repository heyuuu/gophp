// <<generate>>

package spl

import (
	b "sik/builtin"
	"sik/core"
	r "sik/runtime"
	"sik/zend"
)

func SplHeapFromObj(obj *zend.ZendObject) *SplHeapObject {
	return (*SplHeapObject)((*byte)(obj - zend_long((*byte)(&((*SplHeapObject)(nil).GetStd()))-(*byte)(nil))))
}
func Z_SPLHEAP_P(zv *zend.Zval) *SplHeapObject { return SplHeapFromObj(zend.Z_OBJ_P(zv)) }
func SplHeapElem(heap *SplPtrHeap, i int) any {
	return any((*byte)(heap.GetElements() + heap.GetElemSize()*i))
}
func SplHeapElemCopy(heap *SplPtrHeap, to any, from any) {
	r.Assert(to != from)
	memcpy(to, from, heap.GetElemSize())
}
func SplPtrHeapZvalDtor(elem any) { zend.ZvalPtrDtor((*zend.Zval)(elem)) }
func SplPtrHeapZvalCtor(elem any) { zend.Z_TRY_ADDREF_P((*zend.Zval)(elem)) }
func SplPtrHeapPqueueElemDtor(elem any) {
	var pq_elem *SplPqueueElem = elem
	zend.ZvalPtrDtor(&pq_elem.data)
	zend.ZvalPtrDtor(&pq_elem.priority)
}
func SplPtrHeapPqueueElemCtor(elem any) {
	var pq_elem *SplPqueueElem = elem
	zend.Z_TRY_ADDREF_P(&pq_elem.data)
	zend.Z_TRY_ADDREF_P(&pq_elem.priority)
}
func SplPtrHeapCmpCbHelper(object *zend.Zval, heap_object *SplHeapObject, a *zend.Zval, b *zend.Zval, result *zend.ZendLong) int {
	var zresult zend.Zval
	zend.ZendCallMethodWith2Params(object, heap_object.std.ce, &heap_object.fptr_cmp, "compare", &zresult, a, b)
	if zend.ExecutorGlobals.exception != nil {
		return zend.FAILURE
	}
	*result = zend.ZvalGetLong(&zresult)
	zend.ZvalPtrDtor(&zresult)
	return zend.SUCCESS
}
func SplPqueueExtractHelper(result *zend.Zval, elem *SplPqueueElem, flags int) {
	if (flags & SPL_PQUEUE_EXTR_BOTH) == SPL_PQUEUE_EXTR_BOTH {
		zend.ArrayInit(result)
		zend.Z_TRY_ADDREF(elem.GetData())
		zend.AddAssocZvalEx(result, "data", b.SizeOf("\"data\"")-1, &elem.data)
		zend.Z_TRY_ADDREF(elem.GetPriority())
		zend.AddAssocZvalEx(result, "priority", b.SizeOf("\"priority\"")-1, &elem.priority)
		return
	}
	if (flags & SPL_PQUEUE_EXTR_DATA) != 0 {
		zend.ZVAL_COPY(result, &elem.data)
		return
	}
	if (flags & SPL_PQUEUE_EXTR_PRIORITY) != 0 {
		zend.ZVAL_COPY(result, &elem.priority)
		return
	}
	zend.ZEND_ASSERT(false)
}
func SplPtrHeapZvalMaxCmp(x any, y any, object *zend.Zval) int {
	var a *zend.Zval = x
	var b *zend.Zval = y
	var result zend.Zval
	if zend.ExecutorGlobals.exception != nil {
		return 0
	}
	if object != nil {
		var heap_object *SplHeapObject = Z_SPLHEAP_P(object)
		if heap_object.GetFptrCmp() != nil {
			var lval zend.ZendLong = 0
			if SplPtrHeapCmpCbHelper(object, heap_object, a, b, &lval) == zend.FAILURE {

				/* exception or call failure */

				return 0

				/* exception or call failure */

			}
			return zend.ZEND_NORMALIZE_BOOL(lval)
		}
	}
	zend.CompareFunction(&result, a, b)
	return int(zend.Z_LVAL(result))
}
func SplPtrHeapZvalMinCmp(x any, y any, object *zend.Zval) int {
	var a *zend.Zval = x
	var b *zend.Zval = y
	var result zend.Zval
	if zend.ExecutorGlobals.exception != nil {
		return 0
	}
	if object != nil {
		var heap_object *SplHeapObject = Z_SPLHEAP_P(object)
		if heap_object.GetFptrCmp() != nil {
			var lval zend.ZendLong = 0
			if SplPtrHeapCmpCbHelper(object, heap_object, a, b, &lval) == zend.FAILURE {

				/* exception or call failure */

				return 0

				/* exception or call failure */

			}
			return zend.ZEND_NORMALIZE_BOOL(lval)
		}
	}
	zend.CompareFunction(&result, b, a)
	return int(zend.Z_LVAL(result))
}
func SplPtrPqueueElemCmp(x any, y any, object *zend.Zval) int {
	var a *SplPqueueElem = x
	var b *SplPqueueElem = y
	var a_priority_p *zend.Zval = &a.priority
	var b_priority_p *zend.Zval = &b.priority
	var result zend.Zval
	if zend.ExecutorGlobals.exception != nil {
		return 0
	}
	if object != nil {
		var heap_object *SplHeapObject = Z_SPLHEAP_P(object)
		if heap_object.GetFptrCmp() != nil {
			var lval zend.ZendLong = 0
			if SplPtrHeapCmpCbHelper(object, heap_object, a_priority_p, b_priority_p, &lval) == zend.FAILURE {

				/* exception or call failure */

				return 0

				/* exception or call failure */

			}
			return zend.ZEND_NORMALIZE_BOOL(lval)
		}
	}
	zend.CompareFunction(&result, a_priority_p, b_priority_p)
	return int(zend.Z_LVAL(result))
}
func SplPtrHeapInit(cmp SplPtrHeapCmpFunc, ctor SplPtrHeapCtorFunc, dtor SplPtrHeapDtorFunc, elem_size int) *SplPtrHeap {
	var heap *SplPtrHeap = zend.Emalloc(b.SizeOf("spl_ptr_heap"))
	heap.SetDtor(dtor)
	heap.SetCtor(ctor)
	heap.SetCmp(cmp)
	heap.SetElements(zend.Ecalloc(PTR_HEAP_BLOCK_SIZE, elem_size))
	heap.SetMaxSize(PTR_HEAP_BLOCK_SIZE)
	heap.SetCount(0)
	heap.SetFlags(0)
	heap.SetElemSize(elem_size)
	return heap
}
func SplPtrHeapInsert(heap *SplPtrHeap, elem any, cmp_userdata any) {
	var i int
	if heap.GetCount()+1 > heap.GetMaxSize() {
		var alloc_size int = heap.GetMaxSize() * heap.GetElemSize()

		/* we need to allocate more memory */

		heap.SetElements(zend.Erealloc(heap.GetElements(), 2*alloc_size))
		memset((*byte)(heap.GetElements()+alloc_size), 0, alloc_size)
		heap.SetMaxSize(heap.GetMaxSize() * 2)
	}

	/* sifting up */

	for i = heap.GetCount(); i > 0 && heap.GetCmp()(SplHeapElem(heap, (i-1)/2), elem, cmp_userdata) < 0; i = (i - 1) / 2 {
		SplHeapElemCopy(heap, SplHeapElem(heap, i), SplHeapElem(heap, (i-1)/2))
	}
	heap.GetCount()++
	if zend.ExecutorGlobals.exception != nil {

		/* exception thrown during comparison */

		heap.SetFlags(heap.GetFlags() | SPL_HEAP_CORRUPTED)

		/* exception thrown during comparison */

	}
	SplHeapElemCopy(heap, SplHeapElem(heap, i), elem)
}
func SplPtrHeapTop(heap *SplPtrHeap) any {
	if heap.GetCount() == 0 {
		return nil
	}
	return heap.GetElements()
}
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
	bottom = SplHeapElem(heap, b.PreDec(&(heap.GetCount())))
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
	if zend.ExecutorGlobals.exception != nil {

		/* exception thrown during comparison */

		heap.SetFlags(heap.GetFlags() | SPL_HEAP_CORRUPTED)

		/* exception thrown during comparison */

	}
	var to any = SplHeapElem(heap, i)
	if to != bottom {
		SplHeapElemCopy(heap, to, bottom)
	}
	return zend.SUCCESS
}
func SplPtrHeapClone(from *SplPtrHeap) *SplPtrHeap {
	var i int
	var heap *SplPtrHeap = zend.Emalloc(b.SizeOf("spl_ptr_heap"))
	heap.SetDtor(from.GetDtor())
	heap.SetCtor(from.GetCtor())
	heap.SetCmp(from.GetCmp())
	heap.SetMaxSize(from.GetMaxSize())
	heap.SetCount(from.GetCount())
	heap.SetFlags(from.GetFlags())
	heap.SetElemSize(from.GetElemSize())
	heap.SetElements(zend.SafeEmalloc(from.GetElemSize(), from.GetMaxSize(), 0))
	memcpy(heap.GetElements(), from.GetElements(), from.GetElemSize()*from.GetMaxSize())
	for i = 0; i < heap.GetCount(); i++ {
		heap.GetCtor()(SplHeapElem(heap, i))
	}
	return heap
}
func SplPtrHeapDestroy(heap *SplPtrHeap) {
	var i int
	for i = 0; i < heap.GetCount(); i++ {
		heap.GetDtor()(SplHeapElem(heap, i))
	}
	zend.Efree(heap.GetElements())
	zend.Efree(heap)
}
func SplPtrHeapCount(heap *SplPtrHeap) int { return heap.GetCount() }
func SplHeapObjectFreeStorage(object *zend.ZendObject) {
	var intern *SplHeapObject = SplHeapFromObj(object)
	zend.ZendObjectStdDtor(&intern.std)
	SplPtrHeapDestroy(intern.GetHeap())
}
func SplHeapObjectNewEx(class_type *zend.ZendClassEntry, orig *zend.Zval, clone_orig int) *zend.ZendObject {
	var intern *SplHeapObject
	var parent *zend.ZendClassEntry = class_type
	var inherited int = 0
	intern = zend.ZendObjectAlloc(b.SizeOf("spl_heap_object"), parent)
	zend.ZendObjectStdInit(&intern.std, class_type)
	zend.ObjectPropertiesInit(&intern.std, class_type)
	if orig != nil {
		var other *SplHeapObject = Z_SPLHEAP_P(orig)
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
			intern.SetHeap(SplPtrHeapInit(SplPtrPqueueElemCmp, SplPtrHeapPqueueElemCtor, SplPtrHeapPqueueElemDtor, b.SizeOf("spl_pqueue_elem")))
			intern.std.handlers = &spl_handler_SplPriorityQueue
			intern.SetFlags(SPL_PQUEUE_EXTR_DATA)
			break
		}
		if parent == spl_ce_SplMinHeap || parent == spl_ce_SplMaxHeap || parent == spl_ce_SplHeap {
			intern.SetHeap(SplPtrHeapInit(b.Cond(parent == spl_ce_SplMinHeap, SplPtrHeapZvalMinCmp, SplPtrHeapZvalMaxCmp), SplPtrHeapZvalCtor, SplPtrHeapZvalDtor, b.SizeOf("zval")))
			intern.std.handlers = &spl_handler_SplHeap
			break
		}
		parent = parent.parent
		inherited = 1
	}
	if parent == nil {
		core.PhpErrorDocref(nil, zend.E_COMPILE_ERROR, "Internal compiler error, Class is not child of SplHeap")
	}
	if inherited != 0 {
		intern.SetFptrCmp(zend.ZendHashStrFindPtr(&class_type.function_table, "compare", b.SizeOf("\"compare\"")-1))
		if intern.GetFptrCmp().common.scope == parent {
			intern.SetFptrCmp(nil)
		}
		intern.SetFptrCount(zend.ZendHashStrFindPtr(&class_type.function_table, "count", b.SizeOf("\"count\"")-1))
		if intern.GetFptrCount().common.scope == parent {
			intern.SetFptrCount(nil)
		}
	}
	return &intern.std
}
func SplHeapObjectNew(class_type *zend.ZendClassEntry) *zend.ZendObject {
	return SplHeapObjectNewEx(class_type, nil, 0)
}
func SplHeapObjectClone(zobject *zend.Zval) *zend.ZendObject {
	var old_object *zend.ZendObject
	var new_object *zend.ZendObject
	old_object = zend.Z_OBJ_P(zobject)
	new_object = SplHeapObjectNewEx(old_object.ce, zobject, 1)
	zend.ZendObjectsCloneMembers(new_object, old_object)
	return new_object
}
func SplHeapObjectCountElements(object *zend.Zval, count *zend.ZendLong) int {
	var intern *SplHeapObject = Z_SPLHEAP_P(object)
	if intern.GetFptrCount() != nil {
		var rv zend.Zval
		zend.ZendCallMethodWith0Params(object, intern.std.ce, &intern.fptr_count, "count", &rv)
		if !(zend.Z_ISUNDEF(rv)) {
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
func SplHeapObjectGetDebugInfo(ce *zend.ZendClassEntry, obj *zend.Zval) *zend.HashTable {
	var intern *SplHeapObject = Z_SPLHEAP_P(obj)
	var tmp zend.Zval
	var heap_array zend.Zval
	var pnstr *zend.ZendString
	var debug_info *zend.HashTable
	var i int
	if intern.std.properties == nil {
		zend.RebuildObjectProperties(&intern.std)
	}
	debug_info = zend.ZendNewArray(zend.ZendHashNumElements(intern.std.properties) + 1)
	zend.ZendHashCopy(debug_info, intern.std.properties, zend.CopyCtorFuncT(zend.ZvalAddRef))
	pnstr = SplGenPrivatePropName(ce, "flags", b.SizeOf("\"flags\"")-1)
	zend.ZVAL_LONG(&tmp, intern.GetFlags())
	zend.ZendHashUpdate(debug_info, pnstr, &tmp)
	zend.ZendStringReleaseEx(pnstr, 0)
	pnstr = SplGenPrivatePropName(ce, "isCorrupted", b.SizeOf("\"isCorrupted\"")-1)
	zend.ZVAL_BOOL(&tmp, intern.GetHeap().GetFlags()&SPL_HEAP_CORRUPTED)
	zend.ZendHashUpdate(debug_info, pnstr, &tmp)
	zend.ZendStringReleaseEx(pnstr, 0)
	zend.ArrayInit(&heap_array)
	for i = 0; i < intern.GetHeap().GetCount(); i++ {
		if ce == spl_ce_SplPriorityQueue {
			var pq_elem *SplPqueueElem = SplHeapElem(intern.GetHeap(), i)
			var elem zend.Zval
			SplPqueueExtractHelper(&elem, pq_elem, SPL_PQUEUE_EXTR_BOTH)
			zend.AddIndexZval(&heap_array, i, &elem)
		} else {
			var elem *zend.Zval = SplHeapElem(intern.GetHeap(), i)
			zend.AddIndexZval(&heap_array, i, elem)
			zend.Z_TRY_ADDREF_P(elem)
		}
	}
	pnstr = SplGenPrivatePropName(ce, "heap", b.SizeOf("\"heap\"")-1)
	zend.ZendHashUpdate(debug_info, pnstr, &heap_array)
	zend.ZendStringReleaseEx(pnstr, 0)
	return debug_info
}
func SplHeapObjectGetGc(obj *zend.Zval, gc_data **zend.Zval, gc_data_count *int) *zend.HashTable {
	var intern *SplHeapObject = Z_SPLHEAP_P(obj)
	*gc_data = (*zend.Zval)(intern.GetHeap().GetElements())
	*gc_data_count = intern.GetHeap().GetCount()
	return zend.ZendStdGetProperties(obj)
}
func SplPqueueObjectGetGc(obj *zend.Zval, gc_data **zend.Zval, gc_data_count *int) *zend.HashTable {
	var intern *SplHeapObject = Z_SPLHEAP_P(obj)
	*gc_data = (*zend.Zval)(intern.GetHeap().GetElements())

	/* Two zvals (value and priority) per pqueue entry */

	*gc_data_count = 2 * intern.GetHeap().GetCount()
	return zend.ZendStdGetProperties(obj)
}
func zim_spl_SplHeap_count(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var count zend.ZendLong
	var intern *SplHeapObject = Z_SPLHEAP_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	count = SplPtrHeapCount(intern.GetHeap())
	zend.RETVAL_LONG(count)
	return
}
func zim_spl_SplHeap_isEmpty(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplHeapObject = Z_SPLHEAP_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_BOOL(SplPtrHeapCount(intern.GetHeap()) == 0)
	return
}
func zim_spl_SplHeap_insert(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var value *zend.Zval
	var intern *SplHeapObject
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "z", &value) == zend.FAILURE {
		return
	}
	intern = Z_SPLHEAP_P(zend.ZEND_THIS)
	if (intern.GetHeap().GetFlags() & SPL_HEAP_CORRUPTED) != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "Heap is corrupted, heap properties are no longer ensured.", 0)
		return
	}
	zend.Z_TRY_ADDREF_P(value)
	SplPtrHeapInsert(intern.GetHeap(), value, zend.ZEND_THIS)
	zend.RETVAL_TRUE
	return
}
func zim_spl_SplHeap_extract(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplHeapObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	intern = Z_SPLHEAP_P(zend.ZEND_THIS)
	if (intern.GetHeap().GetFlags() & SPL_HEAP_CORRUPTED) != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "Heap is corrupted, heap properties are no longer ensured.", 0)
		return
	}
	if SplPtrHeapDeleteTop(intern.GetHeap(), return_value, zend.ZEND_THIS) == zend.FAILURE {
		zend.ZendThrowException(spl_ce_RuntimeException, "Can't extract from an empty heap", 0)
		return
	}
}
func zim_spl_SplPriorityQueue_insert(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var data *zend.Zval
	var priority *zend.Zval
	var intern *SplHeapObject
	var elem SplPqueueElem
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "zz", &data, &priority) == zend.FAILURE {
		return
	}
	intern = Z_SPLHEAP_P(zend.ZEND_THIS)
	if (intern.GetHeap().GetFlags() & SPL_HEAP_CORRUPTED) != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "Heap is corrupted, heap properties are no longer ensured.", 0)
		return
	}
	zend.ZVAL_COPY(&elem.data, data)
	zend.ZVAL_COPY(&elem.priority, priority)
	SplPtrHeapInsert(intern.GetHeap(), &elem, zend.ZEND_THIS)
	zend.RETVAL_TRUE
	return
}
func zim_spl_SplPriorityQueue_extract(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var elem SplPqueueElem
	var intern *SplHeapObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	intern = Z_SPLHEAP_P(zend.ZEND_THIS)
	if (intern.GetHeap().GetFlags() & SPL_HEAP_CORRUPTED) != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "Heap is corrupted, heap properties are no longer ensured.", 0)
		return
	}
	if SplPtrHeapDeleteTop(intern.GetHeap(), &elem, zend.ZEND_THIS) == zend.FAILURE {
		zend.ZendThrowException(spl_ce_RuntimeException, "Can't extract from an empty heap", 0)
		return
	}
	SplPqueueExtractHelper(return_value, &elem, intern.GetFlags())
	SplPtrHeapPqueueElemDtor(&elem)
}
func zim_spl_SplPriorityQueue_top(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplHeapObject
	var elem *SplPqueueElem
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	intern = Z_SPLHEAP_P(zend.ZEND_THIS)
	if (intern.GetHeap().GetFlags() & SPL_HEAP_CORRUPTED) != 0 {
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
func zim_spl_SplPriorityQueue_setExtractFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var value zend.ZendLong
	var intern *SplHeapObject
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "l", &value) == zend.FAILURE {
		return
	}
	value &= SPL_PQUEUE_EXTR_MASK
	if value == 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "Must specify at least one extract flag", 0)
		return
	}
	intern = Z_SPLHEAP_P(zend.ZEND_THIS)
	intern.SetFlags(value)
	zend.RETVAL_LONG(intern.GetFlags())
	return
}
func zim_spl_SplPriorityQueue_getExtractFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplHeapObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	intern = Z_SPLHEAP_P(zend.ZEND_THIS)
	zend.RETVAL_LONG(intern.GetFlags())
	return
}
func zim_spl_SplHeap_recoverFromCorruption(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplHeapObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	intern = Z_SPLHEAP_P(zend.ZEND_THIS)
	intern.GetHeap().SetFlags(intern.GetHeap().GetFlags() & ^SPL_HEAP_CORRUPTED)
	zend.RETVAL_TRUE
	return
}
func zim_spl_SplHeap_isCorrupted(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplHeapObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	intern = Z_SPLHEAP_P(zend.ZEND_THIS)
	zend.RETVAL_BOOL((intern.GetHeap().GetFlags() & SPL_HEAP_CORRUPTED) != 0)
	return
}
func zim_spl_SplPriorityQueue_compare(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var a *zend.Zval
	var b *zend.Zval
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "zz", &a, &b) == zend.FAILURE {
		return
	}
	zend.RETVAL_LONG(SplPtrHeapZvalMaxCmp(a, b, nil))
	return
}
func zim_spl_SplHeap_top(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var value *zend.Zval
	var intern *SplHeapObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	intern = Z_SPLHEAP_P(zend.ZEND_THIS)
	if (intern.GetHeap().GetFlags() & SPL_HEAP_CORRUPTED) != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "Heap is corrupted, heap properties are no longer ensured.", 0)
		return
	}
	value = SplPtrHeapTop(intern.GetHeap())
	if value == nil {
		zend.ZendThrowException(spl_ce_RuntimeException, "Can't peek at an empty heap", 0)
		return
	}
	zend.ZVAL_COPY_DEREF(return_value, value)
}
func zim_spl_SplMinHeap_compare(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var a *zend.Zval
	var b *zend.Zval
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "zz", &a, &b) == zend.FAILURE {
		return
	}
	zend.RETVAL_LONG(SplPtrHeapZvalMinCmp(a, b, nil))
	return
}
func zim_spl_SplMaxHeap_compare(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var a *zend.Zval
	var b *zend.Zval
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "zz", &a, &b) == zend.FAILURE {
		return
	}
	zend.RETVAL_LONG(SplPtrHeapZvalMaxCmp(a, b, nil))
	return
}
func SplHeapItDtor(iter *zend.ZendObjectIterator) {
	var iterator *SplHeapIt = (*SplHeapIt)(iter)
	zend.ZendUserItInvalidateCurrent(iter)
	zend.ZvalPtrDtor(&iterator.intern.it.data)
}
func SplHeapItRewind(iter *zend.ZendObjectIterator) {}
func SplHeapItValid(iter *zend.ZendObjectIterator) int {
	if Z_SPLHEAP_P(&iter.data).GetHeap().GetCount() != 0 {
		return zend.SUCCESS
	} else {
		return zend.FAILURE
	}
}
func SplHeapItGetCurrentData(iter *zend.ZendObjectIterator) *zend.Zval {
	var object *SplHeapObject = Z_SPLHEAP_P(&iter.data)
	if (object.GetHeap().GetFlags() & SPL_HEAP_CORRUPTED) != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "Heap is corrupted, heap properties are no longer ensured.", 0)
		return nil
	}
	if object.GetHeap().GetCount() == 0 {
		return nil
	} else {
		return SplHeapElem(object.GetHeap(), 0)
	}
}
func SplPqueueItGetCurrentData(iter *zend.ZendObjectIterator) *zend.Zval {
	var user_it *zend.ZendUserIterator = (*zend.ZendUserIterator)(iter)
	var object *SplHeapObject = Z_SPLHEAP_P(&iter.data)
	if (object.GetHeap().GetFlags() & SPL_HEAP_CORRUPTED) != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "Heap is corrupted, heap properties are no longer ensured.", 0)
		return nil
	}
	if object.GetHeap().GetCount() == 0 {
		return nil
	}
	if zend.Z_ISUNDEF(user_it.value) {
		var elem *SplPqueueElem = SplHeapElem(object.GetHeap(), 0)
		SplPqueueExtractHelper(&user_it.value, elem, object.GetFlags())
	}
	return &user_it.value
}
func SplHeapItGetCurrentKey(iter *zend.ZendObjectIterator, key *zend.Zval) {
	var object *SplHeapObject = Z_SPLHEAP_P(&iter.data)
	zend.ZVAL_LONG(key, object.GetHeap().GetCount()-1)
}
func SplHeapItMoveForward(iter *zend.ZendObjectIterator) {
	var object *SplHeapObject = Z_SPLHEAP_P(&iter.data)
	if (object.GetHeap().GetFlags() & SPL_HEAP_CORRUPTED) != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "Heap is corrupted, heap properties are no longer ensured.", 0)
		return
	}
	SplPtrHeapDeleteTop(object.GetHeap(), nil, &iter.data)
	zend.ZendUserItInvalidateCurrent(iter)
}
func zim_spl_SplHeap_key(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplHeapObject = Z_SPLHEAP_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_LONG(intern.GetHeap().GetCount() - 1)
	return
}
func zim_spl_SplHeap_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplHeapObject = Z_SPLHEAP_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	SplPtrHeapDeleteTop(intern.GetHeap(), nil, zend.ZEND_THIS)
}
func zim_spl_SplHeap_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplHeapObject = Z_SPLHEAP_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_BOOL(intern.GetHeap().GetCount() != 0)
	return
}
func zim_spl_SplHeap_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
}
func zim_spl_SplHeap_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplHeapObject = Z_SPLHEAP_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if intern.GetHeap().GetCount() == 0 {
		zend.RETVAL_NULL()
		return
	} else {
		var element *zend.Zval = SplHeapElem(intern.GetHeap(), 0)
		zend.ZVAL_COPY_DEREF(return_value, element)
	}
}
func zim_spl_SplPriorityQueue_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplHeapObject = Z_SPLHEAP_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if intern.GetHeap().GetCount() == 0 {
		zend.RETVAL_NULL()
		return
	} else {
		var elem *SplPqueueElem = SplHeapElem(intern.GetHeap(), 0)
		SplPqueueExtractHelper(return_value, elem, intern.GetFlags())
	}
}
func zim_spl_SplHeap___debugInfo(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_ARR(SplHeapObjectGetDebugInfo(spl_ce_SplHeap, zend.getThis()))
	return
}
func zim_spl_SplPriorityQueue___debugInfo(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_ARR(SplHeapObjectGetDebugInfo(spl_ce_SplPriorityQueue, zend.getThis()))
	return
}
func SplHeapGetIterator(ce *zend.ZendClassEntry, object *zend.Zval, by_ref int) *zend.ZendObjectIterator {
	var iterator *SplHeapIt
	var heap_object *SplHeapObject = Z_SPLHEAP_P(object)
	if by_ref != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "An iterator cannot be used with foreach by reference", 0)
		return nil
	}
	iterator = zend.Emalloc(b.SizeOf("spl_heap_it"))
	zend.ZendIteratorInit(&iterator.intern.it)
	zend.Z_ADDREF_P(object)
	zend.ZVAL_OBJ(&iterator.intern.it.data, zend.Z_OBJ_P(object))
	iterator.intern.it.funcs = &SplHeapItFuncs
	iterator.intern.ce = ce
	iterator.SetFlags(heap_object.GetFlags())
	zend.ZVAL_UNDEF(&iterator.intern.value)
	return &iterator.intern.it
}
func SplPqueueGetIterator(ce *zend.ZendClassEntry, object *zend.Zval, by_ref int) *zend.ZendObjectIterator {
	var iterator *SplHeapIt
	var heap_object *SplHeapObject = Z_SPLHEAP_P(object)
	if by_ref != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "An iterator cannot be used with foreach by reference", 0)
		return nil
	}
	iterator = zend.Emalloc(b.SizeOf("spl_heap_it"))
	zend.ZendIteratorInit((*zend.ZendObjectIterator)(iterator))
	zend.Z_ADDREF_P(object)
	zend.ZVAL_OBJ(&iterator.intern.it.data, zend.Z_OBJ_P(object))
	iterator.intern.it.funcs = &SplPqueueItFuncs
	iterator.intern.ce = ce
	iterator.SetFlags(heap_object.GetFlags())
	zend.ZVAL_UNDEF(&iterator.intern.value)
	return &iterator.intern.it
}
func ZmStartupSplHeap(type_ int, module_number int) int {
	SplRegisterStdClass(&spl_ce_SplHeap, "SplHeap", SplHeapObjectNew, spl_funcs_SplHeap)
	memcpy(&spl_handler_SplHeap, &zend.StdObjectHandlers, b.SizeOf("zend_object_handlers"))
	spl_handler_SplHeap.offset = zend_long((*byte)(&((*SplHeapObject)(nil).GetStd())) - (*byte)(nil))
	spl_handler_SplHeap.clone_obj = SplHeapObjectClone
	spl_handler_SplHeap.count_elements = SplHeapObjectCountElements
	spl_handler_SplHeap.get_gc = SplHeapObjectGetGc
	spl_handler_SplHeap.dtor_obj = zend.ZendObjectsDestroyObject
	spl_handler_SplHeap.free_obj = SplHeapObjectFreeStorage
	zend.ZendClassImplements(spl_ce_SplHeap, 1, spl_ce_Iterator)
	zend.ZendClassImplements(spl_ce_SplHeap, 1, spl_ce_Countable)
	spl_ce_SplHeap.get_iterator = SplHeapGetIterator
	SplRegisterSubClass(&spl_ce_SplMinHeap, spl_ce_SplHeap, "SplMinHeap", SplHeapObjectNew, spl_funcs_SplMinHeap)
	SplRegisterSubClass(&spl_ce_SplMaxHeap, spl_ce_SplHeap, "SplMaxHeap", SplHeapObjectNew, spl_funcs_SplMaxHeap)
	spl_ce_SplMaxHeap.get_iterator = SplHeapGetIterator
	spl_ce_SplMinHeap.get_iterator = SplHeapGetIterator
	SplRegisterStdClass(&spl_ce_SplPriorityQueue, "SplPriorityQueue", SplHeapObjectNew, spl_funcs_SplPriorityQueue)
	memcpy(&spl_handler_SplPriorityQueue, &zend.StdObjectHandlers, b.SizeOf("zend_object_handlers"))
	spl_handler_SplPriorityQueue.offset = zend_long((*byte)(&((*SplHeapObject)(nil).GetStd())) - (*byte)(nil))
	spl_handler_SplPriorityQueue.clone_obj = SplHeapObjectClone
	spl_handler_SplPriorityQueue.count_elements = SplHeapObjectCountElements
	spl_handler_SplPriorityQueue.get_gc = SplPqueueObjectGetGc
	spl_handler_SplPriorityQueue.dtor_obj = zend.ZendObjectsDestroyObject
	spl_handler_SplPriorityQueue.free_obj = SplHeapObjectFreeStorage
	zend.ZendClassImplements(spl_ce_SplPriorityQueue, 1, spl_ce_Iterator)
	zend.ZendClassImplements(spl_ce_SplPriorityQueue, 1, spl_ce_Countable)
	spl_ce_SplPriorityQueue.get_iterator = SplPqueueGetIterator
	zend.ZendDeclareClassConstantLong(spl_ce_SplPriorityQueue, "EXTR_BOTH", b.SizeOf("\"EXTR_BOTH\"")-1, zend.ZendLong(SPL_PQUEUE_EXTR_BOTH))
	zend.ZendDeclareClassConstantLong(spl_ce_SplPriorityQueue, "EXTR_PRIORITY", b.SizeOf("\"EXTR_PRIORITY\"")-1, zend.ZendLong(SPL_PQUEUE_EXTR_PRIORITY))
	zend.ZendDeclareClassConstantLong(spl_ce_SplPriorityQueue, "EXTR_DATA", b.SizeOf("\"EXTR_DATA\"")-1, zend.ZendLong(SPL_PQUEUE_EXTR_DATA))
	return zend.SUCCESS
}
