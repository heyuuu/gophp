// <<generate>>

package core

import (
	b "sik/builtin"
	"sik/zend"
)

func PhpStartupTicks() int {
	PG(tick_functions).Init(b.SizeOf("struct st_tick_function"), nil, 1)
	return zend.SUCCESS
}
func PhpDeactivateTicks() { PG(tick_functions).Clean() }
func PhpShutdownTicks()   { PG(tick_functions).Destroy() }
func PhpCompareTickFunctions(elem1 any, elem2 any) int {
	var e1 *StTickFunction = (*StTickFunction)(elem1)
	var e2 *StTickFunction = (*StTickFunction)(elem2)
	return e1.GetFunc() == e2.GetFunc() && e1.GetArg() == e2.GetArg()
}
func PhpAddTickFunction(func_ func(int, any), arg any) {
	var tmp StTickFunction = MakeStTickFunction(func_, arg)
	PG(tick_functions).AddElement(any(&tmp))
}
func PhpRemoveTickFunction(func_ func(int, any), arg any) {
	var tmp StTickFunction = MakeStTickFunction(func_, arg)
	zend.ZendLlistDelElement(&(PG(tick_functions)), any(&tmp), (func(any, any) int)(PhpCompareTickFunctions))
}
func PhpTickIterator(d any, arg any) {
	var data *StTickFunction = (*StTickFunction)(d)
	data.GetFunc()(*((*int)(arg)), data.GetArg())
}
func PhpRunTicks(count int) {
	PG(tick_functions).ApplyWithArgument(zend.LlistApplyWithArgFuncT(PhpTickIterator), &count)
}
