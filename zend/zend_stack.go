// <<generate>>

package zend

// Source: <Zend/zend_stack.h>

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

// #define ZEND_STACK_H

// @type ZendStack struct

// #define STACK_BLOCK_SIZE       16

// #define ZEND_STACK_APPLY_TOPDOWN       1

// #define ZEND_STACK_APPLY_BOTTOMUP       2

// Source: <Zend/zend_stack.c>

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

// # include "zend_stack.h"

// #define ZEND_STACK_ELEMENT(stack,n) ( ( void * ) ( ( char * ) ( stack ) -> elements + ( stack ) -> size * ( n ) ) )

func ZendStackInit(stack *ZendStack, size int) int {
	stack.SetSize(size)
	stack.SetTop(0)
	stack.SetMax(0)
	stack.SetElements(nil)
	return SUCCESS
}
func ZendStackPush(stack *ZendStack, element any) int {
	/* We need to allocate more memory */

	if stack.GetTop() >= stack.GetMax() {
		stack.SetMax(stack.GetMax() + 16)
		stack.SetElements(_safeErealloc(stack.GetElements(), stack.GetSize(), stack.GetMax(), 0))
	}
	memcpy(any((*byte)(stack.GetElements()+stack.GetSize()*stack.GetTop())), element, stack.GetSize())
	stack.GetTop()++
	return stack.GetTop() - 1
}
func ZendStackTop(stack *ZendStack) any {
	if stack.GetTop() > 0 {
		return any((*byte)(stack.GetElements() + stack.GetSize()*(stack.GetTop()-1)))
	} else {
		return nil
	}
}
func ZendStackDelTop(stack *ZendStack) int {
	stack.GetTop()--
	return SUCCESS
}
func ZendStackIntTop(stack *ZendStack) int {
	var e *int = ZendStackTop(stack)
	if e != nil {
		return *e
	} else {
		return FAILURE
	}
}
func ZendStackIsEmpty(stack *ZendStack) int { return stack.GetTop() == 0 }
func ZendStackDestroy(stack *ZendStack) int {
	if stack.GetElements() {
		_efree(stack.GetElements())
		stack.SetElements(nil)
	}
	return SUCCESS
}
func ZendStackBase(stack *ZendStack) any  { return stack.GetElements() }
func ZendStackCount(stack *ZendStack) int { return stack.GetTop() }
func ZendStackApply(stack *ZendStack, type_ int, apply_function func(element any) int) {
	var i int
	switch type_ {
	case 1:
		for i = stack.GetTop() - 1; i >= 0; i-- {
			if apply_function(any((*byte)(stack.GetElements()+stack.GetSize()*i))) != 0 {
				break
			}
		}
		break
	case 2:
		for i = 0; i < stack.GetTop(); i++ {
			if apply_function(any((*byte)(stack.GetElements()+stack.GetSize()*i))) != 0 {
				break
			}
		}
		break
	}
}
func ZendStackApplyWithArgument(stack *ZendStack, type_ int, apply_function func(element any, arg any) int, arg any) {
	var i int
	switch type_ {
	case 1:
		for i = stack.GetTop() - 1; i >= 0; i-- {
			if apply_function(any((*byte)(stack.GetElements()+stack.GetSize()*i)), arg) != 0 {
				break
			}
		}
		break
	case 2:
		for i = 0; i < stack.GetTop(); i++ {
			if apply_function(any((*byte)(stack.GetElements()+stack.GetSize()*i)), arg) != 0 {
				break
			}
		}
		break
	}
}
func ZendStackClean(stack *ZendStack, func_ func(any), free_elements ZendBool) {
	var i int
	if func_ != nil {
		for i = 0; i < stack.GetTop(); i++ {
			func_(any((*byte)(stack.GetElements() + stack.GetSize()*i)))
		}
	}
	if free_elements != 0 {
		if stack.GetElements() {
			_efree(stack.GetElements())
			stack.SetElements(nil)
		}
		stack.SetMax(0)
		stack.SetTop(stack.GetMax())
	}
}
