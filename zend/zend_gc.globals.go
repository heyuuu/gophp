// <<generate>>

package zend

import (
	b "sik/builtin"
)

var GcCollectCycles func() int

const GC_BENCH = 0
const ZEND_GC_DEBUG = 0
const GC_ADDRESS = 0xfffff
const GC_COLOR = 0x300000
const GC_BLACK = 0x0
const GC_WHITE = 0x100000
const GC_GREY = 0x200000
const GC_PURPLE = 0x300000
const GC_BITS = 0x3
const GC_ROOT = 0x0
const GC_UNUSED = 0x1
const GC_GARBAGE = 0x2
const GC_DTOR_GARBAGE = 0x3
const GC_INVALID = 0
const GC_FIRST_ROOT = 1
const GC_DEFAULT_BUF_SIZE uint32 = 16 * 1024
const GC_BUF_GROW_STEP uint32 = 128 * 1024
const GC_MAX_UNCOMPRESSED = 512 * 1024
const GC_MAX_BUF_SIZE = 0x40000000
const GC_THRESHOLD_DEFAULT = 10000
const GC_THRESHOLD_STEP = 10000
const GC_THRESHOLD_MAX = 1000000000
const GC_THRESHOLD_TRIGGER = 100
const GC_HAS_DESTRUCTORS = 1 << 0

var GcGlobals ZendGcGlobals

const GC_STACK_SEGMENT_SIZE = (4096-ZEND_MM_OVERHEAD)/b.SizeOf("void *") - 2
