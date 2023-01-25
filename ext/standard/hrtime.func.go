// <<generate>>

package standard

import (
	"sik/core"
	"sik/zend"
)

func _timerInit() int {
	/* Timer unavailable. */

	return -1
	return 0
}
func ZmStartupHrtime(type_ int, module_number int) int {
	if 0 > _timerInit() {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Failed to initialize high-resolution timer")
		return zend.FAILURE
	}
	return zend.SUCCESS
}
func _timerCurrent() PhpHrtimeT { return 0 }
func PHP_RETURN_HRTIME(t __auto__) {
	zend.RETVAL_LONG(zend.ZendLong(t))
	return
}
func ZifHrtime(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	zend.RETVAL_FALSE
	return
}
func PhpHrtimeCurrent() PhpHrtimeT { return _timerCurrent() }
