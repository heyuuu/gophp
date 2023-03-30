package zend

import (
	r "github.com/heyuuu/gophp/builtin/file"
)

type ZendStreamFsizerT func(handle any) int
type ZendStreamReaderT func(handle any, buf *byte, len_ int) ssize_t
type ZendStreamCloserT func(handle any)

const ZEND_MMAP_AHEAD = 32

type ZendStreamType = int

const (
	ZEND_HANDLE_FILENAME = iota + 1
	ZEND_HANDLE_FP
	ZEND_HANDLE_STREAM
)

type ZendStatT = __struct__stat

const ZendFseek = r.Fseek
const ZendFtell = r.Ftell
const ZendLseek = lseek
const ZendFstat = fstat
const ZendStat = stat

var Isatty func(fd int) int
