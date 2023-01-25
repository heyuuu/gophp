// <<generate>>

package zend

func ZEND_STACK_ELEMENT(stack *ZendStack, n int) any {
	return any((*byte)(stack.GetElements() + stack.GetSize()*n))
}
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
		stack.SetMax(stack.GetMax() + STACK_BLOCK_SIZE)
		stack.SetElements(SafeErealloc(stack.GetElements(), stack.GetSize(), stack.GetMax(), 0))
	}
	memcpy(ZEND_STACK_ELEMENT(stack, stack.GetTop()), element, stack.GetSize())
	stack.GetTop()++
	return stack.GetTop() - 1
}
func ZendStackTop(stack *ZendStack) any {
	if stack.GetTop() > 0 {
		return ZEND_STACK_ELEMENT(stack, stack.GetTop()-1)
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
		Efree(stack.GetElements())
		stack.SetElements(nil)
	}
	return SUCCESS
}
func ZendStackBase(stack *ZendStack) any  { return stack.GetElements() }
func ZendStackCount(stack *ZendStack) int { return stack.GetTop() }
func ZendStackApply(stack *ZendStack, type_ int, apply_function func(element any) int) {
	var i int
	switch type_ {
	case ZEND_STACK_APPLY_TOPDOWN:
		for i = stack.GetTop() - 1; i >= 0; i-- {
			if apply_function(ZEND_STACK_ELEMENT(stack, i)) != 0 {
				break
			}
		}
		break
	case ZEND_STACK_APPLY_BOTTOMUP:
		for i = 0; i < stack.GetTop(); i++ {
			if apply_function(ZEND_STACK_ELEMENT(stack, i)) != 0 {
				break
			}
		}
		break
	}
}
func ZendStackApplyWithArgument(stack *ZendStack, type_ int, apply_function func(element any, arg any) int, arg any) {
	var i int
	switch type_ {
	case ZEND_STACK_APPLY_TOPDOWN:
		for i = stack.GetTop() - 1; i >= 0; i-- {
			if apply_function(ZEND_STACK_ELEMENT(stack, i), arg) != 0 {
				break
			}
		}
		break
	case ZEND_STACK_APPLY_BOTTOMUP:
		for i = 0; i < stack.GetTop(); i++ {
			if apply_function(ZEND_STACK_ELEMENT(stack, i), arg) != 0 {
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
			func_(ZEND_STACK_ELEMENT(stack, i))
		}
	}
	if free_elements != 0 {
		if stack.GetElements() {
			Efree(stack.GetElements())
			stack.SetElements(nil)
		}
		stack.SetMax(0)
		stack.SetTop(stack.GetMax())
	}
}
