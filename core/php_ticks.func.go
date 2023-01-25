// <<generate>>

package core

import (
	b "sik/builtin"
	"sik/zend"
)

func PhpStartupTicks() int {
	zend.ZendLlistInit(&PG(tick_functions), b.SizeOf("struct st_tick_function"), nil, 1)
	return zend.SUCCESS
}
func PhpDeactivateTicks() { zend.ZendLlistClean(&PG(tick_functions)) }
func PhpShutdownTicks() {
	zend.ZendLlistDestroy(&PG(tick_functions))
}
func PhpCompareTickFunctions(elem1 any, elem2 any) int {
	var e1 *StTickFunction = (*StTickFunction)(elem1)
	var e2 *StTickFunction = (*StTickFunction)(elem2)
	return e1.GetFunc() == e2.GetFunc() && e1.GetArg() == e2.GetArg()
}
func PhpAddTickFunction(func_ func(int, any), arg any) {
	var tmp StTickFunction = StTickFunction{func_, arg}
	zend.ZendLlistAddElement(&PG(tick_functions), any(&tmp))
}
func PhpRemoveTickFunction(func_ func(int, any), arg any) {
	var tmp StTickFunction = StTickFunction{func_, arg}
	zend.ZendLlistDelElement(&PG(tick_functions), any(&tmp), (func(any, any) int)(PhpCompareTickFunctions))
}
func PhpTickIterator(d any, arg any) {
	var data *StTickFunction = (*StTickFunction)(d)
	data.GetFunc()(*((*int)(arg)), data.GetArg())
}
func PhpRunTicks(count int) {
	zend.ZendLlistApplyWithArgument(&PG(tick_functions), zend.LlistApplyWithArgFuncT(PhpTickIterator), &count)
}
