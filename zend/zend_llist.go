// <<generate>>

package zend

import (
	g "sik/runtime/grammar"
)

// Source: <Zend/zend_llist.h>

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
   +----------------------------------------------------------------------+
*/

// #define ZEND_LLIST_H

// @type ZendLlistElement struct

type LlistDtorFuncT func(any)
type LlistCompareFuncT func(**ZendLlistElement, **ZendLlistElement) int
type LlistApplyWithArgsFuncT func(data any, num_args int, args va_list)
type LlistApplyWithArgFuncT func(data any, arg any)
type LlistApplyFuncT func(any)

// @type ZendLlist struct

type ZendLlistPosition *ZendLlistElement

/* traversal */

// #define zend_llist_get_first(l) zend_llist_get_first_ex ( l , NULL )

// #define zend_llist_get_last(l) zend_llist_get_last_ex ( l , NULL )

// #define zend_llist_get_next(l) zend_llist_get_next_ex ( l , NULL )

// #define zend_llist_get_prev(l) zend_llist_get_prev_ex ( l , NULL )

// Source: <Zend/zend_llist.c>

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
   +----------------------------------------------------------------------+
*/

// # include "zend.h"

// # include "zend_llist.h"

// # include "zend_sort.h"

func ZendLlistInit(l *ZendLlist, size int, dtor LlistDtorFuncT, persistent uint8) {
	l.SetHead(nil)
	l.SetTail(nil)
	l.SetCount(0)
	l.SetSize(size)
	l.SetDtor(dtor)
	l.SetPersistent(persistent)
}
func ZendLlistAddElement(l *ZendLlist, element any) {
	var tmp *ZendLlistElement = g.CondF(l.GetPersistent() != 0, func() any { return __zendMalloc(g.SizeOf("zend_llist_element") + l.GetSize() - 1) }, func() any { return _emalloc(g.SizeOf("zend_llist_element") + l.GetSize() - 1) })
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
	var tmp *ZendLlistElement = g.CondF(l.GetPersistent() != 0, func() any { return __zendMalloc(g.SizeOf("zend_llist_element") + l.GetSize() - 1) }, func() any { return _emalloc(g.SizeOf("zend_llist_element") + l.GetSize() - 1) })
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

// #define DEL_LLIST_ELEMENT(current,l) if ( ( current ) -> prev ) { ( current ) -> prev -> next = ( current ) -> next ; } else { ( l ) -> head = ( current ) -> next ; } if ( ( current ) -> next ) { ( current ) -> next -> prev = ( current ) -> prev ; } else { ( l ) -> tail = ( current ) -> prev ; } if ( ( l ) -> dtor ) { ( l ) -> dtor ( ( current ) -> data ) ; } pefree ( ( current ) , ( l ) -> persistent ) ; -- l -> count ;

func ZendLlistDelElement(l *ZendLlist, element any, compare func(element1 any, element2 any) int) {
	var current *ZendLlistElement = l.GetHead()
	for current != nil {
		if compare(current.GetData(), element) != 0 {
			if current.GetPrev() != nil {
				current.GetPrev().SetNext(current.GetNext())
			} else {
				l.SetHead(current.GetNext())
			}
			if current.GetNext() != nil {
				current.GetNext().SetPrev(current.GetPrev())
			} else {
				l.SetTail(current.GetPrev())
			}
			if l.GetDtor() != nil {
				l.GetDtor()(current.GetData())
			}
			g.CondF(l.GetPersistent() != 0, func() { return Free(current) }, func() { return _efree(current) })
			l.GetCount()--
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
		g.CondF(l.GetPersistent() != 0, func() { return Free(current) }, func() { return _efree(current) })
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
	g.CondF(l.GetPersistent() != 0, func() { return Free(old_tail) }, func() { return _efree(old_tail) })
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
			if element.GetPrev() != nil {
				element.GetPrev().SetNext(element.GetNext())
			} else {
				l.SetHead(element.GetNext())
			}
			if element.GetNext() != nil {
				element.GetNext().SetPrev(element.GetPrev())
			} else {
				l.SetTail(element.GetPrev())
			}
			if l.GetDtor() != nil {
				l.GetDtor()(element.GetData())
			}
			g.CondF(l.GetPersistent() != 0, func() { return Free(element) }, func() { return _efree(element) })
			l.GetCount()--
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
	elements = (**ZendLlistElement)(_emalloc(l.GetCount() * g.SizeOf("zend_llist_element *")))
	ptr = &elements[0]
	for element = l.GetHead(); element != nil; element = element.GetNext() {
		g.PostInc(&(*ptr)) = element
	}
	ZendSort(elements, l.GetCount(), g.SizeOf("zend_llist_element *"), CompareFuncT(comp_func), SwapFuncT(ZendLlistSwap))
	l.SetHead(elements[0])
	elements[0].SetPrev(nil)
	for i = 1; i < l.GetCount(); i++ {
		elements[i].SetPrev(elements[i-1])
		elements[i-1].SetNext(elements[i])
	}
	elements[i-1].SetNext(nil)
	l.SetTail(elements[i-1])
	_efree(elements)
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
	var current *ZendLlistPosition = g.CondF2(pos != nil, pos, func() *ZendLlistElement { return &l.traverse_ptr })
	*current = l.GetHead()
	if (*current) != nil {
		return (*current).GetData()
	} else {
		return nil
	}
}
func ZendLlistGetLastEx(l *ZendLlist, pos *ZendLlistPosition) any {
	var current *ZendLlistPosition = g.CondF2(pos != nil, pos, func() *ZendLlistElement { return &l.traverse_ptr })
	*current = l.GetTail()
	if (*current) != nil {
		return (*current).GetData()
	} else {
		return nil
	}
}
func ZendLlistGetNextEx(l *ZendLlist, pos *ZendLlistPosition) any {
	var current *ZendLlistPosition = g.CondF2(pos != nil, pos, func() *ZendLlistElement { return &l.traverse_ptr })
	if (*current) != nil {
		*current = (*current).GetNext()
		if (*current) != nil {
			return (*current).GetData()
		}
	}
	return nil
}
func ZendLlistGetPrevEx(l *ZendLlist, pos *ZendLlistPosition) any {
	var current *ZendLlistPosition = g.CondF2(pos != nil, pos, func() *ZendLlistElement { return &l.traverse_ptr })
	if (*current) != nil {
		*current = (*current).GetPrev()
		if (*current) != nil {
			return (*current).GetData()
		}
	}
	return nil
}
