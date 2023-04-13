package streams

import (
	"github.com/heyuuu/gophp/zend"
)

func PemallocRelOrig(size int) any { return zend.Pemalloc(size) }
func EmallocRelOrig(size int) any  { return zend.Emalloc(size) }
