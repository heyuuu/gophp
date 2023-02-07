// <<generate>>

package spl

import (
	b "sik/builtin"
	"sik/core"
	"sik/ext/standard"
	"sik/zend"
)

func SPL_LLIST_DELREF(elem any) {
	if !(b.PreDec(&(elem.rc))) {
		zend.Efree(elem)
	}
}
func SPL_LLIST_CHECK_DELREF(elem any) {
	if elem && !(b.PreDec(&(elem.rc))) {
		zend.Efree(elem)
	}
}
func SPL_LLIST_ADDREF(elem __auto__) int {
	elem.rc++
	return elem.rc - 1
}
func SPL_LLIST_CHECK_ADDREF(elem *SplPtrLlistElement) {
	if elem != nil {
		elem.GetRc()++
	}
}
func SplDllistFromObj(obj *zend.ZendObject) *SplDllistObject {
	return (*SplDllistObject)((*byte)(obj - zend_long((*byte)(&((*SplDllistObject)(nil).GetStd()))-(*byte)(nil))))
}
func Z_SPLDLLIST_P(zv *zend.Zval) *SplDllistObject { return SplDllistFromObj(zv.GetObj()) }
func SplPtrLlistZvalDtor(elem *SplPtrLlistElement) {
	if !(elem.GetData().IsUndef()) {
		zend.ZvalPtrDtor(elem.GetData())
		elem.GetData().SetUndef()
	}
}
func SplPtrLlistZvalCtor(elem *SplPtrLlistElement) {
	if elem.GetData().IsRefcounted() {
		elem.GetData().AddRefcount()
	}
}
func SplPtrLlistInit(ctor SplPtrLlistCtorFunc, dtor SplPtrLlistDtorFunc) *SplPtrLlist {
	var llist *SplPtrLlist = zend.Emalloc(b.SizeOf("spl_ptr_llist"))
	llist.SetHead(nil)
	llist.SetTail(nil)
	llist.SetCount(0)
	llist.SetDtor(dtor)
	llist.SetCtor(ctor)
	return llist
}
func SplPtrLlistCount(llist *SplPtrLlist) zend.ZendLong { return zend.ZendLong(llist.GetCount()) }
func SplPtrLlistDestroy(llist *SplPtrLlist) {
	var current *SplPtrLlistElement = llist.GetHead()
	var next *SplPtrLlistElement
	var dtor SplPtrLlistDtorFunc = llist.GetDtor()
	for current != nil {
		next = current.GetNext()
		if dtor != nil {
			dtor(current)
		}
		SPL_LLIST_DELREF(current)
		current = next
	}
	zend.Efree(llist)
}
func SplPtrLlistOffset(llist *SplPtrLlist, offset zend.ZendLong, backward int) *SplPtrLlistElement {
	var current *SplPtrLlistElement
	var pos int = 0
	if backward != 0 {
		current = llist.GetTail()
	} else {
		current = llist.GetHead()
	}
	for current != nil && pos < offset {
		pos++
		if backward != 0 {
			current = current.GetPrev()
		} else {
			current = current.GetNext()
		}
	}
	return current
}
func SplPtrLlistUnshift(llist *SplPtrLlist, data *zend.Zval) {
	var elem *SplPtrLlistElement = zend.Emalloc(b.SizeOf("spl_ptr_llist_element"))
	elem.SetRc(1)
	elem.SetPrev(nil)
	elem.SetNext(llist.GetHead())
	zend.ZVAL_COPY_VALUE(elem.GetData(), data)
	if llist.GetHead() != nil {
		llist.GetHead().SetPrev(elem)
	} else {
		llist.SetTail(elem)
	}
	llist.SetHead(elem)
	llist.GetCount()++
	if llist.GetCtor() != nil {
		llist.GetCtor()(elem)
	}
}
func SplPtrLlistPush(llist *SplPtrLlist, data *zend.Zval) {
	var elem *SplPtrLlistElement = zend.Emalloc(b.SizeOf("spl_ptr_llist_element"))
	elem.SetRc(1)
	elem.SetPrev(llist.GetTail())
	elem.SetNext(nil)
	zend.ZVAL_COPY_VALUE(elem.GetData(), data)
	if llist.GetTail() != nil {
		llist.GetTail().SetNext(elem)
	} else {
		llist.SetHead(elem)
	}
	llist.SetTail(elem)
	llist.GetCount()++
	if llist.GetCtor() != nil {
		llist.GetCtor()(elem)
	}
}
func SplPtrLlistPop(llist *SplPtrLlist, ret *zend.Zval) {
	var tail *SplPtrLlistElement = llist.GetTail()
	if tail == nil {
		ret.SetUndef()
		return
	}
	if tail.GetPrev() != nil {
		tail.GetPrev().SetNext(nil)
	} else {
		llist.SetHead(nil)
	}
	llist.SetTail(tail.GetPrev())
	llist.GetCount()--
	zend.ZVAL_COPY(ret, tail.GetData())
	tail.SetPrev(nil)
	if llist.GetDtor() != nil {
		llist.GetDtor()(tail)
	}
	tail.GetData().SetUndef()
	SPL_LLIST_DELREF(tail)
}
func SplPtrLlistLast(llist *SplPtrLlist) *zend.Zval {
	var tail *SplPtrLlistElement = llist.GetTail()
	if tail == nil {
		return nil
	} else {
		return tail.GetData()
	}
}
func SplPtrLlistFirst(llist *SplPtrLlist) *zend.Zval {
	var head *SplPtrLlistElement = llist.GetHead()
	if head == nil {
		return nil
	} else {
		return head.GetData()
	}
}
func SplPtrLlistShift(llist *SplPtrLlist, ret *zend.Zval) {
	var head *SplPtrLlistElement = llist.GetHead()
	if head == nil {
		ret.SetUndef()
		return
	}
	if head.GetNext() != nil {
		head.GetNext().SetPrev(nil)
	} else {
		llist.SetTail(nil)
	}
	llist.SetHead(head.GetNext())
	llist.GetCount()--
	zend.ZVAL_COPY(ret, head.GetData())
	head.SetNext(nil)
	if llist.GetDtor() != nil {
		llist.GetDtor()(head)
	}
	head.GetData().SetUndef()
	SPL_LLIST_DELREF(head)
}
func SplPtrLlistCopy(from *SplPtrLlist, to *SplPtrLlist) {
	var current *SplPtrLlistElement = from.GetHead()
	var next *SplPtrLlistElement

	//???    spl_ptr_llist_ctor_func ctor = from->ctor;

	for current != nil {
		next = current.GetNext()

		/*??? FIXME
		  if (ctor) {
		      ctor(current);
		  }
		*/

		SplPtrLlistPush(to, current.GetData())
		current = next
	}

	//???    spl_ptr_llist_ctor_func ctor = from->ctor;
}
func SplDllistObjectFreeStorage(object *zend.ZendObject) {
	var intern *SplDllistObject = SplDllistFromObj(object)
	var tmp zend.Zval
	zend.ZendObjectStdDtor(intern.GetStd())
	for intern.GetLlist().GetCount() > 0 {
		SplPtrLlistPop(intern.GetLlist(), &tmp)
		zend.ZvalPtrDtor(&tmp)
	}
	if intern.GetGcData() != nil {
		zend.Efree(intern.GetGcData())
	}
	SplPtrLlistDestroy(intern.GetLlist())
	SPL_LLIST_CHECK_DELREF(intern.GetTraversePointer())
}
func SplDllistObjectNewEx(class_type *zend.ZendClassEntry, orig *zend.Zval, clone_orig int) *zend.ZendObject {
	var intern *SplDllistObject
	var parent *zend.ZendClassEntry = class_type
	var inherited int = 0
	intern = zend.ZendObjectAlloc(b.SizeOf("spl_dllist_object"), parent)
	zend.ZendObjectStdInit(intern.GetStd(), class_type)
	zend.ObjectPropertiesInit(intern.GetStd(), class_type)
	intern.SetFlags(0)
	intern.SetTraversePosition(0)
	if orig != nil {
		var other *SplDllistObject = Z_SPLDLLIST_P(orig)
		intern.SetCeGetIterator(other.GetCeGetIterator())
		if clone_orig != 0 {
			intern.SetLlist((*SplPtrLlist)(SplPtrLlistInit(other.GetLlist().GetCtor(), other.GetLlist().GetDtor())))
			SplPtrLlistCopy(other.GetLlist(), intern.GetLlist())
			intern.SetTraversePointer(intern.GetLlist().GetHead())
			SPL_LLIST_CHECK_ADDREF(intern.GetTraversePointer())
		} else {
			intern.SetLlist(other.GetLlist())
			intern.SetTraversePointer(intern.GetLlist().GetHead())
			SPL_LLIST_CHECK_ADDREF(intern.GetTraversePointer())
		}
		intern.SetFlags(other.GetFlags())
	} else {
		intern.SetLlist((*SplPtrLlist)(SplPtrLlistInit(SplPtrLlistZvalCtor, SplPtrLlistZvalDtor)))
		intern.SetTraversePointer(intern.GetLlist().GetHead())
		SPL_LLIST_CHECK_ADDREF(intern.GetTraversePointer())
	}
	for parent != nil {
		if parent == spl_ce_SplStack {
			intern.AddFlags(SPL_DLLIST_IT_FIX | SPL_DLLIST_IT_LIFO)
			intern.GetStd().SetHandlers(&spl_handler_SplDoublyLinkedList)
		} else if parent == spl_ce_SplQueue {
			intern.SetIsItFix(true)
			intern.GetStd().SetHandlers(&spl_handler_SplDoublyLinkedList)
		}
		if parent == spl_ce_SplDoublyLinkedList {
			intern.GetStd().SetHandlers(&spl_handler_SplDoublyLinkedList)
			break
		}
		parent = parent.GetParent()
		inherited = 1
	}
	if parent == nil {
		core.PhpErrorDocref(nil, zend.E_COMPILE_ERROR, "Internal compiler error, Class is not child of SplDoublyLinkedList")
	}
	if inherited != 0 {
		intern.SetFptrOffsetGet(zend.ZendHashStrFindPtr(class_type.GetFunctionTable(), "offsetget", b.SizeOf("\"offsetget\"")-1))
		if intern.GetFptrOffsetGet().GetScope() == parent {
			intern.SetFptrOffsetGet(nil)
		}
		intern.SetFptrOffsetSet(zend.ZendHashStrFindPtr(class_type.GetFunctionTable(), "offsetset", b.SizeOf("\"offsetset\"")-1))
		if intern.GetFptrOffsetSet().GetScope() == parent {
			intern.SetFptrOffsetSet(nil)
		}
		intern.SetFptrOffsetHas(zend.ZendHashStrFindPtr(class_type.GetFunctionTable(), "offsetexists", b.SizeOf("\"offsetexists\"")-1))
		if intern.GetFptrOffsetHas().GetScope() == parent {
			intern.SetFptrOffsetHas(nil)
		}
		intern.SetFptrOffsetDel(zend.ZendHashStrFindPtr(class_type.GetFunctionTable(), "offsetunset", b.SizeOf("\"offsetunset\"")-1))
		if intern.GetFptrOffsetDel().GetScope() == parent {
			intern.SetFptrOffsetDel(nil)
		}
		intern.SetFptrCount(zend.ZendHashStrFindPtr(class_type.GetFunctionTable(), "count", b.SizeOf("\"count\"")-1))
		if intern.GetFptrCount().GetScope() == parent {
			intern.SetFptrCount(nil)
		}
	}
	return intern.GetStd()
}
func SplDllistObjectNew(class_type *zend.ZendClassEntry) *zend.ZendObject {
	return SplDllistObjectNewEx(class_type, nil, 0)
}
func SplDllistObjectClone(zobject *zend.Zval) *zend.ZendObject {
	var old_object *zend.ZendObject
	var new_object *zend.ZendObject
	old_object = zobject.GetObj()
	new_object = SplDllistObjectNewEx(old_object.GetCe(), zobject, 1)
	zend.ZendObjectsCloneMembers(new_object, old_object)
	return new_object
}
func SplDllistObjectCountElements(object *zend.Zval, count *zend.ZendLong) int {
	var intern *SplDllistObject = Z_SPLDLLIST_P(object)
	if intern.GetFptrCount() != nil {
		var rv zend.Zval
		zend.ZendCallMethodWith0Params(object, intern.GetStd().GetCe(), intern.GetFptrCount(), "count", &rv)
		if !(rv.IsUndef()) {
			*count = zend.ZvalGetLong(&rv)
			zend.ZvalPtrDtor(&rv)
			return zend.SUCCESS
		}
		*count = 0
		return zend.FAILURE
	}
	*count = SplPtrLlistCount(intern.GetLlist())
	return zend.SUCCESS
}
func SplDllistObjectGetDebugInfo(obj *zend.Zval) *zend.HashTable {
	var intern *SplDllistObject = Z_SPLDLLIST_P(obj)
	var current *SplPtrLlistElement = intern.GetLlist().GetHead()
	var next *SplPtrLlistElement
	var tmp zend.Zval
	var dllist_array zend.Zval
	var pnstr *zend.ZendString
	var i int = 0
	var debug_info *zend.HashTable
	if intern.GetStd().GetProperties() == nil {
		zend.RebuildObjectProperties(intern.GetStd())
	}
	debug_info = zend.ZendNewArray(1)
	zend.ZendHashCopy(debug_info, intern.GetStd().GetProperties(), zend.CopyCtorFuncT(zend.ZvalAddRef))
	pnstr = SplGenPrivatePropName(spl_ce_SplDoublyLinkedList, "flags", b.SizeOf("\"flags\"")-1)
	tmp.SetLong(intern.GetFlags())
	debug_info.KeyAdd(pnstr.GetStr(), &tmp)
	zend.ZendStringReleaseEx(pnstr, 0)
	zend.ArrayInit(&dllist_array)
	for current != nil {
		next = current.GetNext()
		zend.AddIndexZval(&dllist_array, i, current.GetData())
		if current.GetData().IsRefcounted() {
			current.GetData().AddRefcount()
		}
		i++
		current = next
	}
	pnstr = SplGenPrivatePropName(spl_ce_SplDoublyLinkedList, "dllist", b.SizeOf("\"dllist\"")-1)
	debug_info.KeyAdd(pnstr.GetStr(), &dllist_array)
	zend.ZendStringReleaseEx(pnstr, 0)
	return debug_info
}
func SplDllistObjectGetGc(obj *zend.Zval, gc_data **zend.Zval, gc_data_count *int) *zend.HashTable {
	var intern *SplDllistObject = Z_SPLDLLIST_P(obj)
	var current *SplPtrLlistElement = intern.GetLlist().GetHead()
	var i int = 0
	if intern.GetGcDataCount() < intern.GetLlist().GetCount() {
		intern.SetGcDataCount(intern.GetLlist().GetCount())
		intern.SetGcData(zend.SafeErealloc(intern.GetGcData(), intern.GetGcDataCount(), b.SizeOf("zval"), 0))
	}
	for current != nil {
		zend.ZVAL_COPY_VALUE(intern.GetGcData()[b.PostInc(&i)], current.GetData())
		current = current.GetNext()
	}
	*gc_data = intern.GetGcData()
	*gc_data_count = i
	return zend.ZendStdGetProperties(obj)
}
func zim_spl_SplDoublyLinkedList_push(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var value *zend.Zval
	var intern *SplDllistObject
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "z", &value) == zend.FAILURE {
		return
	}
	intern = Z_SPLDLLIST_P(zend.ZEND_THIS)
	SplPtrLlistPush(intern.GetLlist(), value)
	return_value.SetTrue()
	return
}
func zim_spl_SplDoublyLinkedList_unshift(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var value *zend.Zval
	var intern *SplDllistObject
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "z", &value) == zend.FAILURE {
		return
	}
	intern = Z_SPLDLLIST_P(zend.ZEND_THIS)
	SplPtrLlistUnshift(intern.GetLlist(), value)
	return_value.SetTrue()
	return
}
func zim_spl_SplDoublyLinkedList_pop(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDllistObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	intern = Z_SPLDLLIST_P(zend.ZEND_THIS)
	SplPtrLlistPop(intern.GetLlist(), return_value)
	if return_value.IsUndef() {
		zend.ZendThrowException(spl_ce_RuntimeException, "Can't pop from an empty datastructure", 0)
		return_value.SetNull()
		return
	}
}
func zim_spl_SplDoublyLinkedList_shift(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDllistObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	intern = Z_SPLDLLIST_P(zend.ZEND_THIS)
	SplPtrLlistShift(intern.GetLlist(), return_value)
	if return_value.IsUndef() {
		zend.ZendThrowException(spl_ce_RuntimeException, "Can't shift from an empty datastructure", 0)
		return_value.SetNull()
		return
	}
}
func zim_spl_SplDoublyLinkedList_top(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var value *zend.Zval
	var intern *SplDllistObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	intern = Z_SPLDLLIST_P(zend.ZEND_THIS)
	value = SplPtrLlistLast(intern.GetLlist())
	if value == nil || value.IsUndef() {
		zend.ZendThrowException(spl_ce_RuntimeException, "Can't peek at an empty datastructure", 0)
		return
	}
	zend.ZVAL_COPY_DEREF(return_value, value)
}
func zim_spl_SplDoublyLinkedList_bottom(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var value *zend.Zval
	var intern *SplDllistObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	intern = Z_SPLDLLIST_P(zend.ZEND_THIS)
	value = SplPtrLlistFirst(intern.GetLlist())
	if value == nil || value.IsUndef() {
		zend.ZendThrowException(spl_ce_RuntimeException, "Can't peek at an empty datastructure", 0)
		return
	}
	zend.ZVAL_COPY_DEREF(return_value, value)
}
func zim_spl_SplDoublyLinkedList_count(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var count zend.ZendLong
	var intern *SplDllistObject = Z_SPLDLLIST_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	count = SplPtrLlistCount(intern.GetLlist())
	return_value.SetLong(count)
	return
}
func zim_spl_SplDoublyLinkedList_isEmpty(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var count zend.ZendLong
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	SplDllistObjectCountElements(zend.ZEND_THIS, &count)
	zend.ZVAL_BOOL(return_value, count == 0)
	return
}
func zim_spl_SplDoublyLinkedList_setIteratorMode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var value zend.ZendLong
	var intern *SplDllistObject
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "l", &value) == zend.FAILURE {
		return
	}
	intern = Z_SPLDLLIST_P(zend.ZEND_THIS)
	if intern.IsItFix() && (intern.GetFlags()&SPL_DLLIST_IT_LIFO) != (value&SPL_DLLIST_IT_LIFO) {
		zend.ZendThrowException(spl_ce_RuntimeException, "Iterators' LIFO/FIFO modes for SplStack/SplQueue objects are frozen", 0)
		return
	}
	intern.SetFlags(value&SPL_DLLIST_IT_MASK | intern.GetFlags()&SPL_DLLIST_IT_FIX)
	return_value.SetLong(intern.GetFlags())
	return
}
func zim_spl_SplDoublyLinkedList_getIteratorMode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDllistObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	intern = Z_SPLDLLIST_P(zend.ZEND_THIS)
	return_value.SetLong(intern.GetFlags())
	return
}
func zim_spl_SplDoublyLinkedList_offsetExists(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zindex *zend.Zval
	var intern *SplDllistObject
	var index zend.ZendLong
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "z", &zindex) == zend.FAILURE {
		return
	}
	intern = Z_SPLDLLIST_P(zend.ZEND_THIS)
	index = SplOffsetConvertToLong(zindex)
	zend.ZVAL_BOOL(return_value, index >= 0 && index < intern.GetLlist().GetCount())
	return
}
func zim_spl_SplDoublyLinkedList_offsetGet(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zindex *zend.Zval
	var index zend.ZendLong
	var intern *SplDllistObject
	var element *SplPtrLlistElement
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "z", &zindex) == zend.FAILURE {
		return
	}
	intern = Z_SPLDLLIST_P(zend.ZEND_THIS)
	index = SplOffsetConvertToLong(zindex)
	if index < 0 || index >= intern.GetLlist().GetCount() {
		zend.ZendThrowException(spl_ce_OutOfRangeException, "Offset invalid or out of range", 0)
		return
	}
	element = SplPtrLlistOffset(intern.GetLlist(), index, intern.GetFlags()&SPL_DLLIST_IT_LIFO)
	if element != nil {
		var value *zend.Zval = element.GetData()
		zend.ZVAL_COPY_DEREF(return_value, value)
	} else {
		zend.ZendThrowException(spl_ce_OutOfRangeException, "Offset invalid", 0)
	}
}
func zim_spl_SplDoublyLinkedList_offsetSet(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zindex *zend.Zval
	var value *zend.Zval
	var intern *SplDllistObject
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "zz", &zindex, &value) == zend.FAILURE {
		return
	}
	intern = Z_SPLDLLIST_P(zend.ZEND_THIS)
	if zindex.IsType(zend.IS_NULL) {

		/* $obj[] = ... */

		SplPtrLlistPush(intern.GetLlist(), value)

		/* $obj[] = ... */

	} else {

		/* $obj[$foo] = ... */

		var index zend.ZendLong
		var element *SplPtrLlistElement
		index = SplOffsetConvertToLong(zindex)
		if index < 0 || index >= intern.GetLlist().GetCount() {
			zend.ZendThrowException(spl_ce_OutOfRangeException, "Offset invalid or out of range", 0)
			return
		}
		element = SplPtrLlistOffset(intern.GetLlist(), index, intern.GetFlags()&SPL_DLLIST_IT_LIFO)
		if element != nil {

			/* call dtor on the old element as in spl_ptr_llist_pop */

			if intern.GetLlist().GetDtor() != nil {
				intern.GetLlist().GetDtor()(element)
			}

			/* the element is replaced, delref the old one as in
			 * SplDoublyLinkedList::pop() */

			zend.ZvalPtrDtor(element.GetData())
			zend.ZVAL_COPY_VALUE(element.GetData(), value)

			/* new element, call ctor as in spl_ptr_llist_push */

			if intern.GetLlist().GetCtor() != nil {
				intern.GetLlist().GetCtor()(element)
			}

			/* new element, call ctor as in spl_ptr_llist_push */

		} else {
			zend.ZvalPtrDtor(value)
			zend.ZendThrowException(spl_ce_OutOfRangeException, "Offset invalid", 0)
			return
		}
	}
}
func zim_spl_SplDoublyLinkedList_offsetUnset(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zindex *zend.Zval
	var index zend.ZendLong
	var intern *SplDllistObject
	var element *SplPtrLlistElement
	var llist *SplPtrLlist
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "z", &zindex) == zend.FAILURE {
		return
	}
	intern = Z_SPLDLLIST_P(zend.ZEND_THIS)
	index = SplOffsetConvertToLong(zindex)
	llist = intern.GetLlist()
	if index < 0 || index >= intern.GetLlist().GetCount() {
		zend.ZendThrowException(spl_ce_OutOfRangeException, "Offset out of range", 0)
		return
	}
	element = SplPtrLlistOffset(intern.GetLlist(), index, intern.GetFlags()&SPL_DLLIST_IT_LIFO)
	if element != nil {

		/* connect the neightbors */

		if element.GetPrev() != nil {
			element.GetPrev().SetNext(element.GetNext())
		}
		if element.GetNext() != nil {
			element.GetNext().SetPrev(element.GetPrev())
		}

		/* take care of head/tail */

		if element == llist.GetHead() {
			llist.SetHead(element.GetNext())
		}
		if element == llist.GetTail() {
			llist.SetTail(element.GetPrev())
		}

		/* finally, delete the element */

		llist.GetCount()--
		if llist.GetDtor() != nil {
			llist.GetDtor()(element)
		}
		if intern.GetTraversePointer() == element {
			SPL_LLIST_DELREF(element)
			intern.SetTraversePointer(nil)
		}
		zend.ZvalPtrDtor(element.GetData())
		element.GetData().SetUndef()
		SPL_LLIST_DELREF(element)
	} else {
		zend.ZendThrowException(spl_ce_OutOfRangeException, "Offset invalid", 0)
		return
	}
}
func SplDllistItDtor(iter *zend.ZendObjectIterator) {
	var iterator *SplDllistIt = (*SplDllistIt)(iter)
	SPL_LLIST_CHECK_DELREF(iterator.GetTraversePointer())
	zend.ZendUserItInvalidateCurrent(iter)
	zend.ZvalPtrDtor(iterator.GetIntern().GetIt().GetData())
}
func SplDllistItHelperRewind(traverse_pointer_ptr **SplPtrLlistElement, traverse_position_ptr *int, llist *SplPtrLlist, flags int) {
	SPL_LLIST_CHECK_DELREF(*traverse_pointer_ptr)
	if (flags & SPL_DLLIST_IT_LIFO) != 0 {
		*traverse_position_ptr = llist.GetCount() - 1
		*traverse_pointer_ptr = llist.GetTail()
	} else {
		*traverse_position_ptr = 0
		*traverse_pointer_ptr = llist.GetHead()
	}
	SPL_LLIST_CHECK_ADDREF(*traverse_pointer_ptr)
}
func SplDllistItHelperMoveForward(traverse_pointer_ptr **SplPtrLlistElement, traverse_position_ptr *int, llist *SplPtrLlist, flags int) {
	if (*traverse_pointer_ptr) != nil {
		var old *SplPtrLlistElement = *traverse_pointer_ptr
		if (flags & SPL_DLLIST_IT_LIFO) != 0 {
			*traverse_pointer_ptr = old.GetPrev()
			*traverse_position_ptr--
			if (flags & SPL_DLLIST_IT_DELETE) != 0 {
				var prev zend.Zval
				SplPtrLlistPop(llist, &prev)
				zend.ZvalPtrDtor(&prev)
			}
		} else {
			*traverse_pointer_ptr = old.GetNext()
			if (flags & SPL_DLLIST_IT_DELETE) != 0 {
				var prev zend.Zval
				SplPtrLlistShift(llist, &prev)
				zend.ZvalPtrDtor(&prev)
			} else {
				*traverse_position_ptr++
			}
		}
		SPL_LLIST_DELREF(old)
		SPL_LLIST_CHECK_ADDREF(*traverse_pointer_ptr)
	}
}
func SplDllistItRewind(iter *zend.ZendObjectIterator) {
	var iterator *SplDllistIt = (*SplDllistIt)(iter)
	var object *SplDllistObject = Z_SPLDLLIST_P(iter.GetData())
	var llist *SplPtrLlist = object.GetLlist()
	SplDllistItHelperRewind(iterator.GetTraversePointer(), iterator.GetTraversePosition(), llist, object.GetFlags())
}
func SplDllistItValid(iter *zend.ZendObjectIterator) int {
	var iterator *SplDllistIt = (*SplDllistIt)(iter)
	var element *SplPtrLlistElement = iterator.GetTraversePointer()
	if element != nil {
		return zend.SUCCESS
	} else {
		return zend.FAILURE
	}
}
func SplDllistItGetCurrentData(iter *zend.ZendObjectIterator) *zend.Zval {
	var iterator *SplDllistIt = (*SplDllistIt)(iter)
	var element *SplPtrLlistElement = iterator.GetTraversePointer()
	if element == nil || element.GetData().IsUndef() {
		return nil
	}
	return element.GetData()
}
func SplDllistItGetCurrentKey(iter *zend.ZendObjectIterator, key *zend.Zval) {
	var iterator *SplDllistIt = (*SplDllistIt)(iter)
	key.SetLong(iterator.GetTraversePosition())
}
func SplDllistItMoveForward(iter *zend.ZendObjectIterator) {
	var iterator *SplDllistIt = (*SplDllistIt)(iter)
	var object *SplDllistObject = Z_SPLDLLIST_P(iter.GetData())
	zend.ZendUserItInvalidateCurrent(iter)
	SplDllistItHelperMoveForward(iterator.GetTraversePointer(), iterator.GetTraversePosition(), object.GetLlist(), object.GetFlags())
}
func zim_spl_SplDoublyLinkedList_key(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDllistObject = Z_SPLDLLIST_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	return_value.SetLong(intern.GetTraversePosition())
	return
}
func zim_spl_SplDoublyLinkedList_prev(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDllistObject = Z_SPLDLLIST_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	SplDllistItHelperMoveForward(intern.GetTraversePointer(), intern.GetTraversePosition(), intern.GetLlist(), intern.GetFlags()^SPL_DLLIST_IT_LIFO)
}
func zim_spl_SplDoublyLinkedList_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDllistObject = Z_SPLDLLIST_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	SplDllistItHelperMoveForward(intern.GetTraversePointer(), intern.GetTraversePosition(), intern.GetLlist(), intern.GetFlags())
}
func zim_spl_SplDoublyLinkedList_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDllistObject = Z_SPLDLLIST_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZVAL_BOOL(return_value, intern.GetTraversePointer() != nil)
	return
}
func zim_spl_SplDoublyLinkedList_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDllistObject = Z_SPLDLLIST_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	SplDllistItHelperRewind(intern.GetTraversePointer(), intern.GetTraversePosition(), intern.GetLlist(), intern.GetFlags())
}
func zim_spl_SplDoublyLinkedList_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDllistObject = Z_SPLDLLIST_P(zend.ZEND_THIS)
	var element *SplPtrLlistElement = intern.GetTraversePointer()
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if element == nil || element.GetData().IsUndef() {
		return_value.SetNull()
		return
	} else {
		var value *zend.Zval = element.GetData()
		zend.ZVAL_COPY_DEREF(return_value, value)
	}
}
func zim_spl_SplDoublyLinkedList_serialize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDllistObject = Z_SPLDLLIST_P(zend.ZEND_THIS)
	var buf zend.SmartStr = zend.SmartStr{0}
	var current *SplPtrLlistElement = intern.GetLlist().GetHead()
	var next *SplPtrLlistElement
	var flags zend.Zval
	var var_hash standard.PhpSerializeDataT
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	standard.PHP_VAR_SERIALIZE_INIT(var_hash)

	/* flags */

	flags.SetLong(intern.GetFlags())
	standard.PhpVarSerialize(&buf, &flags, &var_hash)

	/* elements */

	for current != nil {
		buf.AppendByte(':')
		next = current.GetNext()
		standard.PhpVarSerialize(&buf, current.GetData(), &var_hash)
		current = next
	}
	buf.ZeroTail()

	/* done */

	standard.PHP_VAR_SERIALIZE_DESTROY(var_hash)
	if buf.GetS() != nil {
		return_value.SetString(buf.GetS())
		return
	} else {
		return_value.SetNull()
		return
	}
}
func zim_spl_SplDoublyLinkedList_unserialize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDllistObject = Z_SPLDLLIST_P(zend.ZEND_THIS)
	var flags *zend.Zval
	var elem *zend.Zval
	var buf *byte
	var buf_len int
	var p *uint8
	var s *uint8
	var var_hash standard.PhpUnserializeDataT
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "s", &buf, &buf_len) == zend.FAILURE {
		return
	}
	if buf_len == 0 {
		return
	}
	for intern.GetLlist().GetCount() > 0 {
		var tmp zend.Zval
		SplPtrLlistPop(intern.GetLlist(), &tmp)
		zend.ZvalPtrDtor(&tmp)
	}
	p = (*uint8)(buf)
	s = p
	standard.PHP_VAR_UNSERIALIZE_INIT(var_hash)

	/* flags */

	flags = standard.VarTmpVar(&var_hash)
	if standard.PhpVarUnserialize(flags, &p, s+buf_len, &var_hash) == 0 || flags.GetType() != zend.IS_LONG {
		goto error
	}
	intern.SetFlags(int(flags.GetLval()))

	/* elements */

	for (*p) == ':' {
		p++
		elem = standard.VarTmpVar(&var_hash)
		if standard.PhpVarUnserialize(elem, &p, s+buf_len, &var_hash) == 0 {
			goto error
		}
		standard.VarPushDtor(&var_hash, elem)
		SplPtrLlistPush(intern.GetLlist(), elem)
	}
	if (*p) != '0' {
		goto error
	}
	standard.PHP_VAR_UNSERIALIZE_DESTROY(var_hash)
	return
error:
	standard.PHP_VAR_UNSERIALIZE_DESTROY(var_hash)
	zend.ZendThrowExceptionEx(spl_ce_UnexpectedValueException, 0, "Error at offset %zd of %zd bytes", (*byte)(p-buf), buf_len)
	return
}
func zim_spl_SplDoublyLinkedList___serialize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDllistObject = Z_SPLDLLIST_P(zend.ZEND_THIS)
	var current *SplPtrLlistElement = intern.GetLlist().GetHead()
	var tmp zend.Zval
	if zend.ZendParseParametersNoneThrow() == zend.FAILURE {
		return
	}
	zend.ArrayInit(return_value)

	/* flags */

	tmp.SetLong(intern.GetFlags())
	return_value.GetArr().NextIndexInsert(&tmp)

	/* elements */

	zend.ArrayInitSize(&tmp, intern.GetLlist().GetCount())
	for current != nil {
		tmp.GetArr().NextIndexInsert(current.GetData())
		current.GetData().TryAddRefcount()
		current = current.GetNext()
	}
	return_value.GetArr().NextIndexInsert(&tmp)

	/* members */

	tmp.SetArray(zend.ZendStdGetProperties(zend.ZEND_THIS))
	tmp.TryAddRefcount()
	return_value.GetArr().NextIndexInsert(&tmp)
}
func zim_spl_SplDoublyLinkedList___unserialize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDllistObject = Z_SPLDLLIST_P(zend.ZEND_THIS)
	var data *zend.HashTable
	var flags_zv *zend.Zval
	var storage_zv *zend.Zval
	var members_zv *zend.Zval
	var elem *zend.Zval
	if zend.ZendParseParametersThrow(zend.ZEND_NUM_ARGS(), "h", &data) == zend.FAILURE {
		return
	}
	flags_zv = data.IndexFindH(0)
	storage_zv = data.IndexFindH(1)
	members_zv = data.IndexFindH(2)
	if flags_zv == nil || storage_zv == nil || members_zv == nil || flags_zv.GetType() != zend.IS_LONG || storage_zv.GetType() != zend.IS_ARRAY || members_zv.GetType() != zend.IS_ARRAY {
		zend.ZendThrowException(spl_ce_UnexpectedValueException, "Incomplete or ill-typed serialization data", 0)
		return
	}
	intern.SetFlags(int(flags_zv.GetLval()))
	var __ht *zend.HashTable = storage_zv.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *zend.Zval = _p.GetVal()

		elem = _z
		SplPtrLlistPush(intern.GetLlist(), elem)
	}
	zend.ObjectPropertiesLoad(intern.GetStd(), members_zv.GetArr())
}
func zim_spl_SplDoublyLinkedList_add(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zindex *zend.Zval
	var value *zend.Zval
	var intern *SplDllistObject
	var element *SplPtrLlistElement
	var index zend.ZendLong
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "zz", &zindex, &value) == zend.FAILURE {
		return
	}
	intern = Z_SPLDLLIST_P(zend.ZEND_THIS)
	index = SplOffsetConvertToLong(zindex)
	if index < 0 || index > intern.GetLlist().GetCount() {
		zend.ZendThrowException(spl_ce_OutOfRangeException, "Offset invalid or out of range", 0)
		return
	}
	value.TryAddRefcount()
	if index == intern.GetLlist().GetCount() {

		/* If index is the last entry+1 then we do a push because we're not inserting before any entry */

		SplPtrLlistPush(intern.GetLlist(), value)

		/* If index is the last entry+1 then we do a push because we're not inserting before any entry */

	} else {

		/* Create the new element we want to insert */

		var elem *SplPtrLlistElement = zend.Emalloc(b.SizeOf("spl_ptr_llist_element"))

		/* Get the element we want to insert before */

		element = SplPtrLlistOffset(intern.GetLlist(), index, intern.GetFlags()&SPL_DLLIST_IT_LIFO)
		zend.ZVAL_COPY_VALUE(elem.GetData(), value)
		elem.SetRc(1)

		/* connect to the neighbours */

		elem.SetNext(element)
		elem.SetPrev(element.GetPrev())

		/* connect the neighbours to this new element */

		if elem.GetPrev() == nil {
			intern.GetLlist().SetHead(elem)
		} else {
			element.GetPrev().SetNext(elem)
		}
		element.SetPrev(elem)
		intern.GetLlist().GetCount()++
		if intern.GetLlist().GetCtor() != nil {
			intern.GetLlist().GetCtor()(elem)
		}
	}
}
func zim_spl_SplDoublyLinkedList___debugInfo(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	return_value.SetArray(SplDllistObjectGetDebugInfo(zend.getThis()))
	return
}
func SplDllistGetIterator(ce *zend.ZendClassEntry, object *zend.Zval, by_ref int) *zend.ZendObjectIterator {
	var iterator *SplDllistIt
	var dllist_object *SplDllistObject = Z_SPLDLLIST_P(object)
	if by_ref != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "An iterator cannot be used with foreach by reference", 0)
		return nil
	}
	iterator = zend.Emalloc(b.SizeOf("spl_dllist_it"))
	zend.ZendIteratorInit((*zend.ZendObjectIterator)(iterator))
	object.AddRefcount()
	iterator.GetIntern().GetIt().GetData().SetObject(object.GetObj())
	iterator.GetIntern().GetIt().SetFuncs(&SplDllistItFuncs)
	iterator.GetIntern().SetCe(ce)
	iterator.SetTraversePosition(dllist_object.GetTraversePosition())
	iterator.SetTraversePointer(dllist_object.GetTraversePointer())
	iterator.SetFlags(dllist_object.GetFlags() & SPL_DLLIST_IT_MASK)
	iterator.GetIntern().GetValue().SetUndef()
	SPL_LLIST_CHECK_ADDREF(iterator.GetTraversePointer())
	return iterator.GetIntern().GetIt()
}
func ZmStartupSplDllist(type_ int, module_number int) int {
	SplRegisterStdClass(&spl_ce_SplDoublyLinkedList, "SplDoublyLinkedList", SplDllistObjectNew, spl_funcs_SplDoublyLinkedList)
	memcpy(&spl_handler_SplDoublyLinkedList, &zend.StdObjectHandlers, b.SizeOf("zend_object_handlers"))
	spl_handler_SplDoublyLinkedList.SetOffset(zend_long((*byte)(&((*SplDllistObject)(nil).GetStd())) - (*byte)(nil)))
	spl_handler_SplDoublyLinkedList.SetCloneObj(SplDllistObjectClone)
	spl_handler_SplDoublyLinkedList.SetCountElements(SplDllistObjectCountElements)
	spl_handler_SplDoublyLinkedList.SetGetGc(SplDllistObjectGetGc)
	spl_handler_SplDoublyLinkedList.SetDtorObj(zend.ZendObjectsDestroyObject)
	spl_handler_SplDoublyLinkedList.SetFreeObj(SplDllistObjectFreeStorage)
	zend.ZendDeclareClassConstantLong(spl_ce_SplDoublyLinkedList, "IT_MODE_LIFO", b.SizeOf("\"IT_MODE_LIFO\"")-1, zend.ZendLong(SPL_DLLIST_IT_LIFO))
	zend.ZendDeclareClassConstantLong(spl_ce_SplDoublyLinkedList, "IT_MODE_FIFO", b.SizeOf("\"IT_MODE_FIFO\"")-1, zend.ZendLong(0))
	zend.ZendDeclareClassConstantLong(spl_ce_SplDoublyLinkedList, "IT_MODE_DELETE", b.SizeOf("\"IT_MODE_DELETE\"")-1, zend.ZendLong(SPL_DLLIST_IT_DELETE))
	zend.ZendDeclareClassConstantLong(spl_ce_SplDoublyLinkedList, "IT_MODE_KEEP", b.SizeOf("\"IT_MODE_KEEP\"")-1, zend.ZendLong(0))
	zend.ZendClassImplements(spl_ce_SplDoublyLinkedList, 1, spl_ce_Iterator)
	zend.ZendClassImplements(spl_ce_SplDoublyLinkedList, 1, spl_ce_Countable)
	zend.ZendClassImplements(spl_ce_SplDoublyLinkedList, 1, spl_ce_ArrayAccess)
	zend.ZendClassImplements(spl_ce_SplDoublyLinkedList, 1, spl_ce_Serializable)
	spl_ce_SplDoublyLinkedList.SetGetIterator(SplDllistGetIterator)
	SplRegisterSubClass(&spl_ce_SplQueue, spl_ce_SplDoublyLinkedList, "SplQueue", SplDllistObjectNew, spl_funcs_SplQueue)
	SplRegisterSubClass(&spl_ce_SplStack, spl_ce_SplDoublyLinkedList, "SplStack", SplDllistObjectNew, nil)
	spl_ce_SplQueue.SetGetIterator(SplDllistGetIterator)
	spl_ce_SplStack.SetGetIterator(SplDllistGetIterator)
	return zend.SUCCESS
}
