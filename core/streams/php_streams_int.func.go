package streams

import (
	"github.com/heyuuu/gophp/zend"
)

func PemallocRelOrig(size int, persistent __auto__) any { return zend.Pemalloc(size, persistent) }
func EmallocRelOrig(size int) any                       { return zend.Emalloc(size) }
