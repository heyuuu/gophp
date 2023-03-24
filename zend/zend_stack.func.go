package zend

import "sik/zend/types"

func ZendStackInit(stack *ZendStack, size int)        { stack.Init() }
func ZendStackPush(stack *ZendStack, element any) int { return stack.Push(element) }
func ZendStackTop(stack *ZendStack) any               { return stack.Top() }
func ZendStackDelTop(stack *ZendStack)                { stack.DelTop() }
func ZendStackIntTop(stack *ZendStack) int {
	// todo 这里有多个 int / uintptr / null 的隐式转换，需要替换
	var e *int = stack.Top().(*int)
	if e != nil {
		return *e
	} else {
		return types.FAILURE
	}
}
func ZendStackIsEmpty(stack *ZendStack) types.ZendBool { return types.IntBool(stack.IsEmpty()) }
func ZendStackDestroy(stack *ZendStack)                { stack.Destroy() }
func ZendStackApplyWithArgument(stack *ZendStack, type_ int, apply_function func(element any, arg any) int, arg any) {
	stack.ApplyWithArgument(type_, apply_function, arg)
}
func ZendStackClean(stack *ZendStack, func_ func(any), free_elements types.ZendBool) {
	stack.Clean(func_, free_elements)
}
