// <<generate>>

package streams

import (
	"sik/zend"
)

func PemallocRelOrig(size int, persistent __auto__) any { return zend.Pemalloc(size, persistent) }
func PereallocRelOrig(ptr any, size int, persistent uint8) any {
	return zend.Perealloc(ptr, size, persistent)
}
func EmallocRelOrig(size int) any { return zend.Emalloc(size) }
