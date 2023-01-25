// <<generate>>

package zend

import (
	b "sik/builtin"
)

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
