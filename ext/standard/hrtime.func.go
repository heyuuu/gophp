// <<generate>>

package standard

import (
	"sik/core"
	"sik/zend"
	"sik/zend/faults"
	"sik/zend/types"
)

func _timerInit() int {
	/* Timer unavailable. */

	return -1
	return 0
}
func ZmStartupHrtime(type_ int, module_number int) int {
	if 0 > _timerInit() {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Failed to initialize high-resolution timer")
		return types.FAILURE
	}
	return types.SUCCESS
}
func _timerCurrent() PhpHrtimeT { return 0 }
func PHP_RETURN_HRTIME(t __auto__) {
	return_value.SetLong(zend.ZendLong(t))
	return
}
func ZifHrtime(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	return_value.SetFalse()
	return
}
func PhpHrtimeCurrent() PhpHrtimeT { return _timerCurrent() }
