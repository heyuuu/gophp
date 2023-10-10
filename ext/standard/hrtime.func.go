package standard

import (
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/zpp"
)

func _timerInit() int {
	/* Timer unavailable. */

	return -1
	return 0
}
func ZmStartupHrtime(type_ int, module_number int) int {
	if 0 > _timerInit() {
		core.PhpErrorDocref("", faults.E_WARNING, "Failed to initialize high-resolution timer")
		return types.FAILURE
	}
	return types.SUCCESS
}
func _timerCurrent() PhpHrtimeT { return 0 }
func PHP_RETURN_HRTIME(t __auto__) {
	return_value.SetLong(zend.ZendLong(t))
	return
}
func ZifHrtime(executeData zpp.Ex, return_value zpp.Ret, getAsNumber *types.Zval) {
	return_value.SetFalse()
	return
}
func PhpHrtimeCurrent() PhpHrtimeT { return _timerCurrent() }
