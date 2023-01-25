// <<generate>>

package zend

import (
	b "sik/builtin"
)

func ZendLlistGetFirst(l *ZendLlist) any { return ZendLlistGetFirstEx(l, nil) }
func ZendLlistGetLast(l *ZendLlist) any  { return ZendLlistGetLastEx(l, nil) }
func ZendLlistGetNext(l *ZendLlist) any  { return ZendLlistGetNextEx(l, nil) }
func ZendLlistGetPrev(l *ZendLlist) any  { return ZendLlistGetPrevEx(l, nil) }
func ZendLlistInit(l *ZendLlist, size int, dtor LlistDtorFuncT, persistent uint8) {
	l.SetHead(nil)
	l.SetTail(nil)
	l.SetCount(0)
	l.SetSize(size)
	l.SetDtor(dtor)
	l.SetPersistent(persistent)
}
func ZendLlistAddElement(l *ZendLlist, element any) {
	var tmp *ZendLlistElement = Pemalloc(b.SizeOf("zend_llist_element")+l.GetSize()-1, l.GetPersistent())
	tmp.SetPrev(l.GetTail())
	tmp.SetNext(nil)
	if l.GetTail() != nil {
		l.GetTail().SetNext(tmp)
	} else {
		l.SetHead(tmp)
	}
	l.SetTail(tmp)
	memcpy(tmp.GetData(), element, l.GetSize())
	l.GetCount()++
}
func ZendLlistPrependElement(l *ZendLlist, element any) {
	var tmp *ZendLlistElement = Pemalloc(b.SizeOf("zend_llist_element")+l.GetSize()-1, l.GetPersistent())
	tmp.SetNext(l.GetHead())
	tmp.SetPrev(nil)
	if l.GetHead() != nil {
		l.GetHead().SetPrev(tmp)
	} else {
		l.SetTail(tmp)
	}
	l.SetHead(tmp)
	memcpy(tmp.GetData(), element, l.GetSize())
	l.GetCount()++
}
func DEL_LLIST_ELEMENT(current any, l *ZendLlist) {
	if current.prev {
		current.prev.next = current.next
	} else {
		l.SetHead(current.next)
	}
	if current.next {
		current.next.prev = current.prev
	} else {
		l.SetTail(current.prev)
	}
	if l.GetDtor() != nil {
		l.GetDtor()(current.data)
	}
	Pefree(current, l.GetPersistent())
	l.GetCount()--
}
func ZendLlistDelElement(l *ZendLlist, element any, compare func(element1 any, element2 any) int) {
	var current *ZendLlistElement = l.GetHead()
	for current != nil {
		if compare(current.GetData(), element) != 0 {
			DEL_LLIST_ELEMENT(current, l)
			break
		}
		current = current.GetNext()
	}
}
func ZendLlistDestroy(l *ZendLlist) {
	var current *ZendLlistElement = l.GetHead()
	var next *ZendLlistElement
	for current != nil {
		next = current.GetNext()
		if l.GetDtor() != nil {
			l.GetDtor()(current.GetData())
		}
		Pefree(current, l.GetPersistent())
		current = next
	}
	l.SetCount(0)
}
func ZendLlistClean(l *ZendLlist) {
	ZendLlistDestroy(l)
	l.SetTail(nil)
	l.SetHead(l.GetTail())
}
func ZendLlistRemoveTail(l *ZendLlist) {
	var old_tail *ZendLlistElement = l.GetTail()
	if old_tail == nil {
		return
	}
	if old_tail.GetPrev() != nil {
		old_tail.GetPrev().SetNext(nil)
	} else {
		l.SetHead(nil)
	}
	l.SetTail(old_tail.GetPrev())
	l.GetCount()--
	if l.GetDtor() != nil {
		l.GetDtor()(old_tail.GetData())
	}
	Pefree(old_tail, l.GetPersistent())
}
func ZendLlistCopy(dst *ZendLlist, src *ZendLlist) {
	var ptr *ZendLlistElement
	ZendLlistInit(dst, src.GetSize(), src.GetDtor(), src.GetPersistent())
	ptr = src.GetHead()
	for ptr != nil {
		ZendLlistAddElement(dst, ptr.GetData())
		ptr = ptr.GetNext()
	}
}
func ZendLlistApplyWithDel(l *ZendLlist, func_ func(data any) int) {
	var element *ZendLlistElement
	var next *ZendLlistElement
	element = l.GetHead()
	for element != nil {
		next = element.GetNext()
		if func_(element.GetData()) != 0 {
			DEL_LLIST_ELEMENT(element, l)
		}
		element = next
	}
}
func ZendLlistApply(l *ZendLlist, func_ LlistApplyFuncT) {
	var element *ZendLlistElement
	for element = l.GetHead(); element != nil; element = element.GetNext() {
		func_(element.GetData())
	}
}
func ZendLlistSwap(p **ZendLlistElement, q **ZendLlistElement) {
	var t *ZendLlistElement
	t = *p
	*p = *q
	*q = t
}
func ZendLlistSort(l *ZendLlist, comp_func LlistCompareFuncT) {
	var i int
	var elements **ZendLlistElement
	var element *ZendLlistElement
	var ptr **ZendLlistElement
	if l.GetCount() == 0 {
		return
	}
	elements = (**ZendLlistElement)(Emalloc(l.GetCount() * b.SizeOf("zend_llist_element *")))
	ptr = &elements[0]
	for element = l.GetHead(); element != nil; element = element.GetNext() {
		b.PostInc(&(*ptr)) = element
	}
	ZendSort(elements, l.GetCount(), b.SizeOf("zend_llist_element *"), CompareFuncT(comp_func), SwapFuncT(ZendLlistSwap))
	l.SetHead(elements[0])
	elements[0].SetPrev(nil)
	for i = 1; i < l.GetCount(); i++ {
		elements[i].SetPrev(elements[i-1])
		elements[i-1].SetNext(elements[i])
	}
	elements[i-1].SetNext(nil)
	l.SetTail(elements[i-1])
	Efree(elements)
}
func ZendLlistApplyWithArgument(l *ZendLlist, func_ LlistApplyWithArgFuncT, arg any) {
	var element *ZendLlistElement
	for element = l.GetHead(); element != nil; element = element.GetNext() {
		func_(element.GetData(), arg)
	}
}
func ZendLlistApplyWithArguments(l *ZendLlist, func_ LlistApplyWithArgsFuncT, num_args int, _ ...any) {
	var element *ZendLlistElement
	var args va_list
	va_start(args, num_args)
	for element = l.GetHead(); element != nil; element = element.GetNext() {
		func_(element.GetData(), num_args, args)
	}
	va_end(args)
}
func ZendLlistCount(l *ZendLlist) int { return l.GetCount() }
func ZendLlistGetFirstEx(l *ZendLlist, pos *ZendLlistPosition) any {
	var current *ZendLlistPosition = b.CondF2(pos != nil, pos, func() *ZendLlistElement { return &l.traverse_ptr })
	*current = l.GetHead()
	if (*current) != nil {
		return (*current).GetData()
	} else {
		return nil
	}
}
func ZendLlistGetLastEx(l *ZendLlist, pos *ZendLlistPosition) any {
	var current *ZendLlistPosition = b.CondF2(pos != nil, pos, func() *ZendLlistElement { return &l.traverse_ptr })
	*current = l.GetTail()
	if (*current) != nil {
		return (*current).GetData()
	} else {
		return nil
	}
}
func ZendLlistGetNextEx(l *ZendLlist, pos *ZendLlistPosition) any {
	var current *ZendLlistPosition = b.CondF2(pos != nil, pos, func() *ZendLlistElement { return &l.traverse_ptr })
	if (*current) != nil {
		*current = (*current).GetNext()
		if (*current) != nil {
			return (*current).GetData()
		}
	}
	return nil
}
func ZendLlistGetPrevEx(l *ZendLlist, pos *ZendLlistPosition) any {
	var current *ZendLlistPosition = b.CondF2(pos != nil, pos, func() *ZendLlistElement { return &l.traverse_ptr })
	if (*current) != nil {
		*current = (*current).GetPrev()
		if (*current) != nil {
			return (*current).GetData()
		}
	}
	return nil
}
