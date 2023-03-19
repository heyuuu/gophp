// <<generate>>

package spl

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
	"sik/zend/types"
)

func SplHeapFromObj(obj *types.ZendObject) *SplHeapObject {
	return (*SplHeapObject)((*byte)(obj - zend_long((*byte)(&((*SplHeapObject)(nil).GetStd()))-(*byte)(nil))))
}
func Z_SPLHEAP_P(zv *types.Zval) *SplHeapObject { return SplHeapFromObj(zv.GetObj()) }
func SplHeapElem(heap *SplPtrHeap, i int) any {
	return any((*byte)(heap.GetElements() + heap.GetElemSize()*i))
}
func SplHeapElemCopy(heap *SplPtrHeap, to any, from any) {
	b.Assert(to != from)
	memcpy(to, from, heap.GetElemSize())
}
func SplPtrHeapZvalDtor(elem any) { zend.ZvalPtrDtor((*types.Zval)(elem)) }
func SplPtrHeapZvalCtor(elem any) { (*types.Zval)(elem).TryAddRefcount() }
func SplPtrHeapPqueueElemDtor(elem any) {
	var pq_elem *SplPqueueElem = elem
	zend.ZvalPtrDtor(pq_elem.GetData())
	zend.ZvalPtrDtor(pq_elem.GetPriority())
}
func SplPtrHeapPqueueElemCtor(elem any) {
	var pq_elem *SplPqueueElem = elem
	pq_elem.GetData().TryAddRefcount()
	pq_elem.GetPriority().TryAddRefcount()
}
func SplPtrHeapCmpCbHelper(object *types.Zval, heap_object *SplHeapObject, a *types.Zval, b *types.Zval, result *zend.ZendLong) int {
	var zresult types.Zval
	zend.ZendCallMethodWith2Params(object, heap_object.GetStd().GetCe(), heap_object.GetFptrCmp(), "compare", &zresult, a, b)
	if zend.EG__().GetException() != nil {
		return types.FAILURE
	}
	*result = zend.ZvalGetLong(&zresult)
	zend.ZvalPtrDtor(&zresult)
	return types.SUCCESS
}
func SplPqueueExtractHelper(result *types.Zval, elem *SplPqueueElem, flags int) {
	if (flags & SPL_PQUEUE_EXTR_BOTH) == SPL_PQUEUE_EXTR_BOTH {
		zend.ArrayInit(result)
		elem.GetData().TryAddRefcount()
		zend.AddAssocZvalEx(result, "data", elem.GetData())
		elem.GetPriority().TryAddRefcount()
		zend.AddAssocZvalEx(result, "priority", elem.GetPriority())
		return
	}
	if (flags & SPL_PQUEUE_EXTR_DATA) != 0 {
		types.ZVAL_COPY(result, elem.GetData())
		return
	}
	if (flags & SPL_PQUEUE_EXTR_PRIORITY) != 0 {
		types.ZVAL_COPY(result, elem.GetPriority())
		return
	}
	b.Assert(false)
}
func SplPtrHeapZvalMaxCmp(x any, y any, object *types.Zval) int {
	var a *types.Zval = x
	var b *types.Zval = y
	var result types.Zval
	if zend.EG__().GetException() != nil {
		return 0
	}
	if object != nil {
		var heap_object *SplHeapObject = Z_SPLHEAP_P(object)
		if heap_object.GetFptrCmp() != nil {
			var lval zend.ZendLong = 0
			if SplPtrHeapCmpCbHelper(object, heap_object, a, b, &lval) == types.FAILURE {

				/* exception or call failure */

				return 0

				/* exception or call failure */

			}
			return zend.ZEND_NORMALIZE_BOOL(lval)
		}
	}
	zend.CompareFunction(&result, a, b)
	return int(result.GetLval())
}
func SplPtrHeapZvalMinCmp(x any, y any, object *types.Zval) int {
	var a *types.Zval = x
	var b *types.Zval = y
	var result types.Zval
	if zend.EG__().GetException() != nil {
		return 0
	}
	if object != nil {
		var heap_object *SplHeapObject = Z_SPLHEAP_P(object)
		if heap_object.GetFptrCmp() != nil {
			var lval zend.ZendLong = 0
			if SplPtrHeapCmpCbHelper(object, heap_object, a, b, &lval) == types.FAILURE {

				/* exception or call failure */

				return 0

				/* exception or call failure */

			}
			return zend.ZEND_NORMALIZE_BOOL(lval)
		}
	}
	zend.CompareFunction(&result, b, a)
	return int(result.GetLval())
}
func SplPtrPqueueElemCmp(x any, y any, object *types.Zval) int {
	var a *SplPqueueElem = x
	var b *SplPqueueElem = y
	var a_priority_p *types.Zval = a.GetPriority()
	var b_priority_p *types.Zval = b.GetPriority()
	var result types.Zval
	if zend.EG__().GetException() != nil {
		return 0
	}
	if object != nil {
		var heap_object *SplHeapObject = Z_SPLHEAP_P(object)
		if heap_object.GetFptrCmp() != nil {
			var lval zend.ZendLong = 0
			if SplPtrHeapCmpCbHelper(object, heap_object, a_priority_p, b_priority_p, &lval) == types.FAILURE {

				/* exception or call failure */

				return 0

				/* exception or call failure */

			}
			return zend.ZEND_NORMALIZE_BOOL(lval)
		}
	}
	zend.CompareFunction(&result, a_priority_p, b_priority_p)
	return int(result.GetLval())
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
	if zend.EG__().GetException() != nil {

		/* exception thrown during comparison */

		heap.SetIsHeapCorrupted(true)

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
		return types.FAILURE
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
	if zend.EG__().GetException() != nil {

		/* exception thrown during comparison */

		heap.SetIsHeapCorrupted(true)

		/* exception thrown during comparison */

	}
	var to any = SplHeapElem(heap, i)
	if to != bottom {
		SplHeapElemCopy(heap, to, bottom)
	}
	return types.SUCCESS
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
func SplHeapObjectFreeStorage(object *types.ZendObject) {
	var intern *SplHeapObject = SplHeapFromObj(object)
	zend.ZendObjectStdDtor(intern.GetStd())
	SplPtrHeapDestroy(intern.GetHeap())
}
func SplHeapObjectNewEx(class_type *types.ClassEntry, orig *types.Zval, clone_orig int) *types.ZendObject {
	var intern *SplHeapObject
	var parent *types.ClassEntry = class_type
	var inherited int = 0
	intern = zend.ZendObjectAlloc(b.SizeOf("spl_heap_object"), parent)
	zend.ZendObjectStdInit(intern.GetStd(), class_type)
	zend.ObjectPropertiesInit(intern.GetStd(), class_type)
	if orig != nil {
		var other *SplHeapObject = Z_SPLHEAP_P(orig)
		intern.GetStd().SetHandlers(other.GetStd().GetHandlers())
		intern.SetCeGetIterator(other.GetCeGetIterator())
		if clone_orig != 0 {
			intern.SetHeap(SplPtrHeapClone(other.GetHeap()))
		} else {
			intern.SetHeap(other.GetHeap())
		}
		intern.SetFlags(other.GetFlags())
		intern.SetFptrCmp(other.GetFptrCmp())
		intern.SetFptrCount(other.GetFptrCount())
		return intern.GetStd()
	}
	for parent != nil {
		if parent == spl_ce_SplPriorityQueue {
			intern.SetHeap(SplPtrHeapInit(SplPtrPqueueElemCmp, SplPtrHeapPqueueElemCtor, SplPtrHeapPqueueElemDtor, b.SizeOf("spl_pqueue_elem")))
			intern.GetStd().SetHandlers(&spl_handler_SplPriorityQueue)
			intern.SetFlags(SPL_PQUEUE_EXTR_DATA)
			break
		}
		if parent == spl_ce_SplMinHeap || parent == spl_ce_SplMaxHeap || parent == spl_ce_SplHeap {
			intern.SetHeap(SplPtrHeapInit(b.Cond(parent == spl_ce_SplMinHeap, SplPtrHeapZvalMinCmp, SplPtrHeapZvalMaxCmp), SplPtrHeapZvalCtor, SplPtrHeapZvalDtor, b.SizeOf("zval")))
			intern.GetStd().SetHandlers(&spl_handler_SplHeap)
			break
		}
		parent = parent.GetParent()
		inherited = 1
	}
	if parent == nil {
		core.PhpErrorDocref(nil, zend.E_COMPILE_ERROR, "Internal compiler error, Class is not child of SplHeap")
	}
	if inherited != 0 {
		intern.SetFptrCmp(zend.ZendHashStrFindPtr(class_type.GetFunctionTable(), "compare", b.SizeOf("\"compare\"")-1))
		if intern.GetFptrCmp().GetScope() == parent {
			intern.SetFptrCmp(nil)
		}
		intern.SetFptrCount(zend.ZendHashStrFindPtr(class_type.GetFunctionTable(), "count", b.SizeOf("\"count\"")-1))
		if intern.GetFptrCount().GetScope() == parent {
			intern.SetFptrCount(nil)
		}
	}
	return intern.GetStd()
}
func SplHeapObjectNew(class_type *types.ClassEntry) *types.ZendObject {
	return SplHeapObjectNewEx(class_type, nil, 0)
}
func SplHeapObjectClone(zobject *types.Zval) *types.ZendObject {
	var old_object *types.ZendObject
	var new_object *types.ZendObject
	old_object = zobject.GetObj()
	new_object = SplHeapObjectNewEx(old_object.GetCe(), zobject, 1)
	zend.ZendObjectsCloneMembers(new_object, old_object)
	return new_object
}
func SplHeapObjectCountElements(object *types.Zval, count *zend.ZendLong) int {
	var intern *SplHeapObject = Z_SPLHEAP_P(object)
	if intern.GetFptrCount() != nil {
		var rv types.Zval
		zend.ZendCallMethodWith0Params(object, intern.GetStd().GetCe(), intern.GetFptrCount(), "count", &rv)
		if !(rv.IsUndef()) {
			*count = zend.ZvalGetLong(&rv)
			zend.ZvalPtrDtor(&rv)
			return types.SUCCESS
		}
		*count = 0
		return types.FAILURE
	}
	*count = intern.GetHeap().GetCount()
	return types.SUCCESS
}
func SplHeapObjectGetDebugInfo(ce *types.ClassEntry, obj *types.Zval) *types.HashTable {
	var intern *SplHeapObject = Z_SPLHEAP_P(obj)
	var tmp types.Zval
	var heap_array types.Zval
	var pnstr *types.ZendString
	var debug_info *types.HashTable
	var i int
	if intern.GetStd().GetProperties() == nil {
		zend.RebuildObjectProperties(intern.GetStd())
	}
	debug_info = zend.ZendNewArray(intern.GetStd().GetProperties().GetNNumOfElements() + 1)
	zend.ZendHashCopy(debug_info, intern.GetStd().GetProperties(), types.CopyCtorFuncT(zend.ZvalAddRef))
	pnstr = SplGenPrivatePropName(ce, "flags")
	tmp.SetLong(intern.GetFlags())
	debug_info.KeyUpdate(pnstr.GetStr(), &tmp)
	types.ZendStringReleaseEx(pnstr, 0)
	pnstr = SplGenPrivatePropName(ce, "isCorrupted")
	types.ZVAL_BOOL(&tmp, intern.GetHeap().IsHeapCorrupted())
	debug_info.KeyUpdate(pnstr.GetStr(), &tmp)
	types.ZendStringReleaseEx(pnstr, 0)
	zend.ArrayInit(&heap_array)
	for i = 0; i < intern.GetHeap().GetCount(); i++ {
		if ce == spl_ce_SplPriorityQueue {
			var pq_elem *SplPqueueElem = SplHeapElem(intern.GetHeap(), i)
			var elem types.Zval
			SplPqueueExtractHelper(&elem, pq_elem, SPL_PQUEUE_EXTR_BOTH)
			zend.AddIndexZval(&heap_array, i, &elem)
		} else {
			var elem *types.Zval = SplHeapElem(intern.GetHeap(), i)
			zend.AddIndexZval(&heap_array, i, elem)
			elem.TryAddRefcount()
		}
	}
	pnstr = SplGenPrivatePropName(ce, "heap")
	debug_info.KeyUpdate(pnstr.GetStr(), &heap_array)
	types.ZendStringReleaseEx(pnstr, 0)
	return debug_info
}
func SplHeapObjectGetGc(obj *types.Zval, gc_data **types.Zval, gc_data_count *int) *types.HashTable {
	var intern *SplHeapObject = Z_SPLHEAP_P(obj)
	*gc_data = (*types.Zval)(intern.GetHeap().GetElements())
	*gc_data_count = intern.GetHeap().GetCount()
	return zend.ZendStdGetProperties(obj)
}
func SplPqueueObjectGetGc(obj *types.Zval, gc_data **types.Zval, gc_data_count *int) *types.HashTable {
	var intern *SplHeapObject = Z_SPLHEAP_P(obj)
	*gc_data = (*types.Zval)(intern.GetHeap().GetElements())

	/* Two zvals (value and priority) per pqueue entry */

	*gc_data_count = 2 * intern.GetHeap().GetCount()
	return zend.ZendStdGetProperties(obj)
}
func zim_spl_SplHeap_count(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var count zend.ZendLong
	var intern *SplHeapObject = Z_SPLHEAP_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	count = intern.GetHeap().GetCount()
	return_value.SetLong(count)
	return
}
func zim_spl_SplHeap_isEmpty(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplHeapObject = Z_SPLHEAP_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	types.ZVAL_BOOL(return_value, intern.GetHeap().GetCount() == 0)
	return
}
func zim_spl_SplHeap_insert(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var value *types.Zval
	var intern *SplHeapObject
	if zend.ZendParseParameters(executeData.NumArgs(), "z", &value) == types.FAILURE {
		return
	}
	intern = Z_SPLHEAP_P(zend.ZEND_THIS(executeData))
	if intern.GetHeap().IsHeapCorrupted() {
		zend.ZendThrowException(spl_ce_RuntimeException, "Heap is corrupted, heap properties are no longer ensured.", 0)
		return
	}
	value.TryAddRefcount()
	SplPtrHeapInsert(intern.GetHeap(), value, zend.ZEND_THIS(executeData))
	return_value.SetTrue()
	return
}
func zim_spl_SplHeap_extract(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplHeapObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	intern = Z_SPLHEAP_P(zend.ZEND_THIS(executeData))
	if intern.GetHeap().IsHeapCorrupted() {
		zend.ZendThrowException(spl_ce_RuntimeException, "Heap is corrupted, heap properties are no longer ensured.", 0)
		return
	}
	if SplPtrHeapDeleteTop(intern.GetHeap(), return_value, zend.ZEND_THIS(executeData)) == types.FAILURE {
		zend.ZendThrowException(spl_ce_RuntimeException, "Can't extract from an empty heap", 0)
		return
	}
}
func zim_spl_SplPriorityQueue_insert(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var data *types.Zval
	var priority *types.Zval
	var intern *SplHeapObject
	var elem SplPqueueElem
	if zend.ZendParseParameters(executeData.NumArgs(), "zz", &data, &priority) == types.FAILURE {
		return
	}
	intern = Z_SPLHEAP_P(zend.ZEND_THIS(executeData))
	if intern.GetHeap().IsHeapCorrupted() {
		zend.ZendThrowException(spl_ce_RuntimeException, "Heap is corrupted, heap properties are no longer ensured.", 0)
		return
	}
	types.ZVAL_COPY(elem.GetData(), data)
	types.ZVAL_COPY(elem.GetPriority(), priority)
	SplPtrHeapInsert(intern.GetHeap(), &elem, zend.ZEND_THIS(executeData))
	return_value.SetTrue()
	return
}
func zim_spl_SplPriorityQueue_extract(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var elem SplPqueueElem
	var intern *SplHeapObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	intern = Z_SPLHEAP_P(zend.ZEND_THIS(executeData))
	if intern.GetHeap().IsHeapCorrupted() {
		zend.ZendThrowException(spl_ce_RuntimeException, "Heap is corrupted, heap properties are no longer ensured.", 0)
		return
	}
	if SplPtrHeapDeleteTop(intern.GetHeap(), &elem, zend.ZEND_THIS(executeData)) == types.FAILURE {
		zend.ZendThrowException(spl_ce_RuntimeException, "Can't extract from an empty heap", 0)
		return
	}
	SplPqueueExtractHelper(return_value, &elem, intern.GetFlags())
	SplPtrHeapPqueueElemDtor(&elem)
}
func zim_spl_SplPriorityQueue_top(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplHeapObject
	var elem *SplPqueueElem
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	intern = Z_SPLHEAP_P(zend.ZEND_THIS(executeData))
	if intern.GetHeap().IsHeapCorrupted() {
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
func zim_spl_SplPriorityQueue_setExtractFlags(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var value zend.ZendLong
	var intern *SplHeapObject
	if zend.ZendParseParameters(executeData.NumArgs(), "l", &value) == types.FAILURE {
		return
	}
	value &= SPL_PQUEUE_EXTR_MASK
	if value == 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "Must specify at least one extract flag", 0)
		return
	}
	intern = Z_SPLHEAP_P(zend.ZEND_THIS(executeData))
	intern.SetFlags(value)
	return_value.SetLong(intern.GetFlags())
	return
}
func zim_spl_SplPriorityQueue_getExtractFlags(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplHeapObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	intern = Z_SPLHEAP_P(zend.ZEND_THIS(executeData))
	return_value.SetLong(intern.GetFlags())
	return
}
func zim_spl_SplHeap_recoverFromCorruption(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplHeapObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	intern = Z_SPLHEAP_P(zend.ZEND_THIS(executeData))
	intern.GetHeap().SetFlags(intern.GetHeap().GetFlags() & ^SPL_HEAP_CORRUPTED)
	return_value.SetTrue()
	return
}
func zim_spl_SplHeap_isCorrupted(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplHeapObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	intern = Z_SPLHEAP_P(zend.ZEND_THIS(executeData))
	types.ZVAL_BOOL(return_value, intern.GetHeap().IsHeapCorrupted())
	return
}
func zim_spl_SplPriorityQueue_compare(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var a *types.Zval
	var b *types.Zval
	if zend.ZendParseParameters(executeData.NumArgs(), "zz", &a, &b) == types.FAILURE {
		return
	}
	return_value.SetLong(SplPtrHeapZvalMaxCmp(a, b, nil))
	return
}
func zim_spl_SplHeap_top(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var value *types.Zval
	var intern *SplHeapObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	intern = Z_SPLHEAP_P(zend.ZEND_THIS(executeData))
	if intern.GetHeap().IsHeapCorrupted() {
		zend.ZendThrowException(spl_ce_RuntimeException, "Heap is corrupted, heap properties are no longer ensured.", 0)
		return
	}
	value = SplPtrHeapTop(intern.GetHeap())
	if value == nil {
		zend.ZendThrowException(spl_ce_RuntimeException, "Can't peek at an empty heap", 0)
		return
	}
	types.ZVAL_COPY_DEREF(return_value, value)
}
func zim_spl_SplMinHeap_compare(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var a *types.Zval
	var b *types.Zval
	if zend.ZendParseParameters(executeData.NumArgs(), "zz", &a, &b) == types.FAILURE {
		return
	}
	return_value.SetLong(SplPtrHeapZvalMinCmp(a, b, nil))
	return
}
func zim_spl_SplMaxHeap_compare(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var a *types.Zval
	var b *types.Zval
	if zend.ZendParseParameters(executeData.NumArgs(), "zz", &a, &b) == types.FAILURE {
		return
	}
	return_value.SetLong(SplPtrHeapZvalMaxCmp(a, b, nil))
	return
}
func SplHeapItDtor(iter *zend.ZendObjectIterator) {
	var iterator *SplHeapIt = (*SplHeapIt)(iter)
	zend.ZendUserItInvalidateCurrent(iter)
	zend.ZvalPtrDtor(iterator.GetIntern().GetIt().GetData())
}
func SplHeapItRewind(iter *zend.ZendObjectIterator) {}
func SplHeapItValid(iter *zend.ZendObjectIterator) int {
	if Z_SPLHEAP_P(iter.GetData()).GetHeap().GetCount() != 0 {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func SplHeapItGetCurrentData(iter *zend.ZendObjectIterator) *types.Zval {
	var object *SplHeapObject = Z_SPLHEAP_P(iter.GetData())
	if object.GetHeap().IsHeapCorrupted() {
		zend.ZendThrowException(spl_ce_RuntimeException, "Heap is corrupted, heap properties are no longer ensured.", 0)
		return nil
	}
	if object.GetHeap().GetCount() == 0 {
		return nil
	} else {
		return SplHeapElem(object.GetHeap(), 0)
	}
}
func SplPqueueItGetCurrentData(iter *zend.ZendObjectIterator) *types.Zval {
	var user_it *zend.ZendUserIterator = (*zend.ZendUserIterator)(iter)
	var object *SplHeapObject = Z_SPLHEAP_P(iter.GetData())
	if object.GetHeap().IsHeapCorrupted() {
		zend.ZendThrowException(spl_ce_RuntimeException, "Heap is corrupted, heap properties are no longer ensured.", 0)
		return nil
	}
	if object.GetHeap().GetCount() == 0 {
		return nil
	}
	if user_it.GetValue().IsUndef() {
		var elem *SplPqueueElem = SplHeapElem(object.GetHeap(), 0)
		SplPqueueExtractHelper(user_it.GetValue(), elem, object.GetFlags())
	}
	return user_it.GetValue()
}
func SplHeapItGetCurrentKey(iter *zend.ZendObjectIterator, key *types.Zval) {
	var object *SplHeapObject = Z_SPLHEAP_P(iter.GetData())
	key.SetLong(object.GetHeap().GetCount() - 1)
}
func SplHeapItMoveForward(iter *zend.ZendObjectIterator) {
	var object *SplHeapObject = Z_SPLHEAP_P(iter.GetData())
	if object.GetHeap().IsHeapCorrupted() {
		zend.ZendThrowException(spl_ce_RuntimeException, "Heap is corrupted, heap properties are no longer ensured.", 0)
		return
	}
	SplPtrHeapDeleteTop(object.GetHeap(), nil, iter.GetData())
	zend.ZendUserItInvalidateCurrent(iter)
}
func zim_spl_SplHeap_key(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplHeapObject = Z_SPLHEAP_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	return_value.SetLong(intern.GetHeap().GetCount() - 1)
	return
}
func zim_spl_SplHeap_next(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplHeapObject = Z_SPLHEAP_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	SplPtrHeapDeleteTop(intern.GetHeap(), nil, zend.ZEND_THIS(executeData))
}
func zim_spl_SplHeap_valid(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplHeapObject = Z_SPLHEAP_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	types.ZVAL_BOOL(return_value, intern.GetHeap().GetCount() != 0)
	return
}
func zim_spl_SplHeap_rewind(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
}
func zim_spl_SplHeap_current(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplHeapObject = Z_SPLHEAP_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	if intern.GetHeap().GetCount() == 0 {
		return_value.SetNull()
		return
	} else {
		var element *types.Zval = SplHeapElem(intern.GetHeap(), 0)
		types.ZVAL_COPY_DEREF(return_value, element)
	}
}
func zim_spl_SplPriorityQueue_current(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplHeapObject = Z_SPLHEAP_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	if intern.GetHeap().GetCount() == 0 {
		return_value.SetNull()
		return
	} else {
		var elem *SplPqueueElem = SplHeapElem(intern.GetHeap(), 0)
		SplPqueueExtractHelper(return_value, elem, intern.GetFlags())
	}
}
func zim_spl_SplHeap___debugInfo(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	return_value.SetArray(SplHeapObjectGetDebugInfo(spl_ce_SplHeap, zend.getThis()))
	return
}
func zim_spl_SplPriorityQueue___debugInfo(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	return_value.SetArray(SplHeapObjectGetDebugInfo(spl_ce_SplPriorityQueue, zend.getThis()))
	return
}
func SplHeapGetIterator(ce *types.ClassEntry, object *types.Zval, by_ref int) *zend.ZendObjectIterator {
	var iterator *SplHeapIt
	var heap_object *SplHeapObject = Z_SPLHEAP_P(object)
	if by_ref != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "An iterator cannot be used with foreach by reference", 0)
		return nil
	}
	iterator = zend.Emalloc(b.SizeOf("spl_heap_it"))
	zend.ZendIteratorInit(iterator.GetIntern().GetIt())
	object.AddRefcount()
	iterator.GetIntern().GetIt().GetData().SetObject(object.GetObj())
	iterator.GetIntern().GetIt().SetFuncs(&SplHeapItFuncs)
	iterator.GetIntern().SetCe(ce)
	iterator.SetFlags(heap_object.GetFlags())
	iterator.GetIntern().GetValue().SetUndef()
	return iterator.GetIntern().GetIt()
}
func SplPqueueGetIterator(ce *types.ClassEntry, object *types.Zval, by_ref int) *zend.ZendObjectIterator {
	var iterator *SplHeapIt
	var heap_object *SplHeapObject = Z_SPLHEAP_P(object)
	if by_ref != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "An iterator cannot be used with foreach by reference", 0)
		return nil
	}
	iterator = zend.Emalloc(b.SizeOf("spl_heap_it"))
	zend.ZendIteratorInit((*zend.ZendObjectIterator)(iterator))
	object.AddRefcount()
	iterator.GetIntern().GetIt().GetData().SetObject(object.GetObj())
	iterator.GetIntern().GetIt().SetFuncs(&SplPqueueItFuncs)
	iterator.GetIntern().SetCe(ce)
	iterator.SetFlags(heap_object.GetFlags())
	iterator.GetIntern().GetValue().SetUndef()
	return iterator.GetIntern().GetIt()
}
func ZmStartupSplHeap(type_ int, module_number int) int {
	SplRegisterStdClass(&spl_ce_SplHeap, "SplHeap", SplHeapObjectNew, spl_funcs_SplHeap)
	memcpy(&spl_handler_SplHeap, &zend.StdObjectHandlers, b.SizeOf("zend_object_handlers"))
	spl_handler_SplHeap.SetOffset(zend_long((*byte)(&((*SplHeapObject)(nil).GetStd())) - (*byte)(nil)))
	spl_handler_SplHeap.SetCloneObj(SplHeapObjectClone)
	spl_handler_SplHeap.SetCountElements(SplHeapObjectCountElements)
	spl_handler_SplHeap.SetGetGc(SplHeapObjectGetGc)
	spl_handler_SplHeap.SetDtorObj(zend.ZendObjectsDestroyObject)
	spl_handler_SplHeap.SetFreeObj(SplHeapObjectFreeStorage)
	zend.ZendClassImplements(spl_ce_SplHeap, 1, spl_ce_Iterator)
	zend.ZendClassImplements(spl_ce_SplHeap, 1, spl_ce_Countable)
	spl_ce_SplHeap.SetGetIterator(SplHeapGetIterator)
	SplRegisterSubClass(&spl_ce_SplMinHeap, spl_ce_SplHeap, "SplMinHeap", SplHeapObjectNew, spl_funcs_SplMinHeap)
	SplRegisterSubClass(&spl_ce_SplMaxHeap, spl_ce_SplHeap, "SplMaxHeap", SplHeapObjectNew, spl_funcs_SplMaxHeap)
	spl_ce_SplMaxHeap.SetGetIterator(SplHeapGetIterator)
	spl_ce_SplMinHeap.SetGetIterator(SplHeapGetIterator)
	SplRegisterStdClass(&spl_ce_SplPriorityQueue, "SplPriorityQueue", SplHeapObjectNew, spl_funcs_SplPriorityQueue)
	memcpy(&spl_handler_SplPriorityQueue, &zend.StdObjectHandlers, b.SizeOf("zend_object_handlers"))
	spl_handler_SplPriorityQueue.SetOffset(zend_long((*byte)(&((*SplHeapObject)(nil).GetStd())) - (*byte)(nil)))
	spl_handler_SplPriorityQueue.SetCloneObj(SplHeapObjectClone)
	spl_handler_SplPriorityQueue.SetCountElements(SplHeapObjectCountElements)
	spl_handler_SplPriorityQueue.SetGetGc(SplPqueueObjectGetGc)
	spl_handler_SplPriorityQueue.SetDtorObj(zend.ZendObjectsDestroyObject)
	spl_handler_SplPriorityQueue.SetFreeObj(SplHeapObjectFreeStorage)
	zend.ZendClassImplements(spl_ce_SplPriorityQueue, 1, spl_ce_Iterator)
	zend.ZendClassImplements(spl_ce_SplPriorityQueue, 1, spl_ce_Countable)
	spl_ce_SplPriorityQueue.SetGetIterator(SplPqueueGetIterator)
	zend.ZendDeclareClassConstantLong(spl_ce_SplPriorityQueue, "EXTR_BOTH", b.SizeOf("\"EXTR_BOTH\"")-1, zend.ZendLong(SPL_PQUEUE_EXTR_BOTH))
	zend.ZendDeclareClassConstantLong(spl_ce_SplPriorityQueue, "EXTR_PRIORITY", b.SizeOf("\"EXTR_PRIORITY\"")-1, zend.ZendLong(SPL_PQUEUE_EXTR_PRIORITY))
	zend.ZendDeclareClassConstantLong(spl_ce_SplPriorityQueue, "EXTR_DATA", b.SizeOf("\"EXTR_DATA\"")-1, zend.ZendLong(SPL_PQUEUE_EXTR_DATA))
	return types.SUCCESS
}
