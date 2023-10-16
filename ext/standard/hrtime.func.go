package standard

import (
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func _timerInit() int {
	/* Timer unavailable. */
	return -1
}
func ZmStartupHrtime() int {
	if 0 > _timerInit() {
		core.PhpErrorDocref("", faults.E_WARNING, "Failed to initialize high-resolution timer")
		return types.FAILURE
	}
	return types.SUCCESS
}
func ZifHrtime() bool {
	return false
}
