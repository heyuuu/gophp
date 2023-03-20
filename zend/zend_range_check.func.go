package zend

import (
	"sik/core"
)

func ZEND_SIZE_T_INT_OVFL(size int) bool { return size > int(core.INT_MAX) }
