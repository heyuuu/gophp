package zend

import (
	r "github.com/heyuuu/gophp/builtin/file"
)

const ZEND_MMAP_AHEAD = 32

type ZendStatT = __struct__stat

const ZendFseek = r.Fseek
const ZendFtell = r.Ftell
const ZendLseek = lseek
const ZendFstat = fstat
const ZendStat = stat

var Isatty func(fd int) int
