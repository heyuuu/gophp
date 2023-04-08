package core

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/types"
)

func PhpStartupTicks() int {
	PG__().tick_functions.Init(b.SizeOf("struct st_tick_function"), nil, 1)
	return types.SUCCESS
}
func PhpDeactivateTicks() { PG__().tick_functions.Clean() }
func PhpShutdownTicks()   { PG__().tick_functions.Destroy() }
func PhpAddTickFunction(func_ func(int, any), arg any) {
	var tmp StTickFunction = MakeStTickFunction(func_, arg)
	PG__().tick_functions.AddElement(any(&tmp))
}
func PhpTickIterator(d any, arg any) {
	var data *StTickFunction = (*StTickFunction)(d)
	data.GetFunc()(*((*int)(arg)), data.GetArg())
}
func PhpRunTicks(count int) {
	PG__().tick_functions.ApplyWithArgument(zend.LlistApplyWithArgFuncT(PhpTickIterator), &count)
}
