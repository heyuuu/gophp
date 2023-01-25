// <<generate>>

package zend

import (
	b "sik/builtin"
)

// Source: <Zend/zend_ptr_stack.h>

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

// #define ZEND_PTR_STACK_H

const PTR_STACK_BLOCK_SIZE = 64

// #define ZEND_PTR_STACK_RESIZE_IF_NEEDED(stack,count) if ( stack -> top + count > stack -> max ) { do { stack -> max += PTR_STACK_BLOCK_SIZE ; } while ( stack -> top + count > stack -> max ) ; stack -> elements = ( void * * ) perealloc ( stack -> elements , ( sizeof ( void * ) * ( stack -> max ) ) , stack -> persistent ) ; stack -> top_element = stack -> elements + stack -> top ; }

/*    Not doing this with a macro because of the loop unrolling in the element assignment.
Just using a macro for 3 in the body for readability sake. */

func ZendPtrStack3Push(stack *ZendPtrStack, a any, b any, c any) {
	// #define ZEND_PTR_STACK_NUM_ARGS       3

	if stack.GetTop()+3 > stack.GetMax() {
		for {
			stack.SetMax(stack.GetMax() + PTR_STACK_BLOCK_SIZE)
			if stack.GetTop()+3 <= stack.GetMax() {
				break
			}
		}
		stack.SetElements((*any)(Perealloc(stack.GetElements(), b.SizeOf("void *")*stack.GetMax(), stack.GetPersistent())))
		stack.SetTopElement(stack.GetElements() + stack.GetTop())
	}
	stack.SetTop(stack.GetTop() + 3)
	*(b.PostInc(&(stack.GetTopElement()))) = a
	*(b.PostInc(&(stack.GetTopElement()))) = b
	*(b.PostInc(&(stack.GetTopElement()))) = c
}
func ZendPtrStack2Push(stack *ZendPtrStack, a any, b any) {
	// #define ZEND_PTR_STACK_NUM_ARGS       2

	if stack.GetTop()+2 > stack.GetMax() {
		for {
			stack.SetMax(stack.GetMax() + PTR_STACK_BLOCK_SIZE)
			if stack.GetTop()+2 <= stack.GetMax() {
				break
			}
		}
		stack.SetElements((*any)(Perealloc(stack.GetElements(), b.SizeOf("void *")*stack.GetMax(), stack.GetPersistent())))
		stack.SetTopElement(stack.GetElements() + stack.GetTop())
	}
	stack.SetTop(stack.GetTop() + 2)
	*(b.PostInc(&(stack.GetTopElement()))) = a
	*(b.PostInc(&(stack.GetTopElement()))) = b
}
func ZendPtrStack3Pop(stack *ZendPtrStack, a *any, b *any, c *any) {
	*a = *(b.PreDec(&(stack.GetTopElement())))
	*b = *(b.PreDec(&(stack.GetTopElement())))
	*c = *(b.PreDec(&(stack.GetTopElement())))
	stack.SetTop(stack.GetTop() - 3)
}
func ZendPtrStack2Pop(stack *ZendPtrStack, a *any, b *any) {
	*a = *(b.PreDec(&(stack.GetTopElement())))
	*b = *(b.PreDec(&(stack.GetTopElement())))
	stack.SetTop(stack.GetTop() - 2)
}
func ZendPtrStackPush(stack *ZendPtrStack, ptr any) {
	if stack.GetTop()+1 > stack.GetMax() {
		for {
			stack.SetMax(stack.GetMax() + PTR_STACK_BLOCK_SIZE)
			if stack.GetTop()+1 <= stack.GetMax() {
				break
			}
		}
		stack.SetElements((*any)(Perealloc(stack.GetElements(), b.SizeOf("void *")*stack.GetMax(), stack.GetPersistent())))
		stack.SetTopElement(stack.GetElements() + stack.GetTop())
	}
	stack.GetTop()++
	*(b.PostInc(&(stack.GetTopElement()))) = ptr
}
func ZendPtrStackPop(stack *ZendPtrStack) any {
	stack.GetTop()--
	return *(b.PreDec(&(stack.GetTopElement())))
}
func ZendPtrStackTop(stack *ZendPtrStack) any {
	return stack.GetElements()[stack.GetTop()-1]
}

// Source: <Zend/zend_ptr_stack.c>

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

// # include "zend_ptr_stack.h"

// # include < stdarg . h >

func ZendPtrStackInitEx(stack *ZendPtrStack, persistent ZendBool) {
	stack.SetElements(nil)
	stack.SetTopElement(stack.GetElements())
	stack.SetMax(0)
	stack.SetTop(stack.GetMax())
	stack.SetPersistent(persistent)
}
func ZendPtrStackInit(stack *ZendPtrStack) { ZendPtrStackInitEx(stack, 0) }
func ZendPtrStackNPush(stack *ZendPtrStack, count int, _ ...any) {
	var ptr va_list
	var elem any
	if stack.GetTop()+count > stack.GetMax() {
		for {
			stack.SetMax(stack.GetMax() + PTR_STACK_BLOCK_SIZE)
			if stack.GetTop()+count <= stack.GetMax() {
				break
			}
		}
		stack.SetElements((*any)(Perealloc(stack.GetElements(), b.SizeOf("void *")*stack.GetMax(), stack.GetPersistent())))
		stack.SetTopElement(stack.GetElements() + stack.GetTop())
	}
	va_start(ptr, count)
	for count > 0 {
		elem = __va_arg(ptr, any(_))
		stack.GetTop()++
		*(b.PostInc(&(stack.GetTopElement()))) = elem
		count--
	}
	va_end(ptr)
}
func ZendPtrStackNPop(stack *ZendPtrStack, count int, _ ...any) {
	var ptr va_list
	var elem *any
	va_start(ptr, count)
	for count > 0 {
		elem = __va_arg(ptr, (*any)(_))
		*elem = *(b.PreDec(&(stack.GetTopElement())))
		stack.GetTop()--
		count--
	}
	va_end(ptr)
}
func ZendPtrStackDestroy(stack *ZendPtrStack) {
	if stack.GetElements() != nil {
		Pefree(stack.GetElements(), stack.GetPersistent())
	}
}
func ZendPtrStackApply(stack *ZendPtrStack, func_ func(any)) {
	var i int = stack.GetTop()
	for b.PreDec(&i) >= 0 {
		func_(stack.GetElements()[i])
	}
}
func ZendPtrStackReverseApply(stack *ZendPtrStack, func_ func(any)) {
	var i int = 0
	for i < stack.GetTop() {
		func_(stack.GetElements()[b.PostInc(&i)])
	}
}
func ZendPtrStackClean(stack *ZendPtrStack, func_ func(any), free_elements ZendBool) {
	ZendPtrStackApply(stack, func_)
	if free_elements != 0 {
		var i int = stack.GetTop()
		for b.PreDec(&i) >= 0 {
			Pefree(stack.GetElements()[i], stack.GetPersistent())
		}
	}
	stack.SetTop(0)
	stack.SetTopElement(stack.GetElements())
}
func ZendPtrStackNumElements(stack *ZendPtrStack) int { return stack.GetTop() }
