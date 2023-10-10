package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

func ZendStackIntTop(stack *ZendStack) int {
	// todo 这里有多个 int / uintptr / null 的隐式转换，需要替换
	var e *int = stack.Top().(*int)
	if e != nil {
		return *e
	} else {
		return types.FAILURE
	}
}
func ZendStackDestroy(stack *ZendStack) { stack.Destroy() }
func ZendStackApplyWithArgument(stack *ZendStack, type_ int, apply_function func(element any, arg any) int, arg any) {
	stack.ApplyWithArgument(type_, apply_function, arg)
}
func ZendStackClean(stack *ZendStack, func_ func(any), free_elements bool) {
	stack.Clean(func_, free_elements)
}
