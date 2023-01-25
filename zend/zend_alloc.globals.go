// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/core"
)

const ZEND_MM_ALIGNMENT_MASK = ^(core.ZEND_MM_ALIGNMENT - 1)

const ZEND_MM_OVERHEAD = 0
const _emallocLarge = _emalloc
const _emallocHuge = _emalloc
const _efreeLarge = _efree
const _efreeHuge = _efree

const ZEND_MM_CUSTOM_HEAP_NONE = 0
const ZEND_MM_CUSTOM_HEAP_STD = 1
const ZEND_MM_CUSTOM_HEAP_DEBUG = 2

type ZendMmChunkAllocT func(storage *ZendMmStorage, size int, alignment int) any
type ZendMmChunkFreeT func(storage *ZendMmStorage, chunk any, size int)
type ZendMmChunkTruncateT func(storage *ZendMmStorage, chunk any, old_size int, new_size int) int
type ZendMmChunkExtendT func(storage *ZendMmStorage, chunk any, old_size int, new_size int) int

const MAP_FAILED = any(-1)
const MAP_POPULATE = 0
const REAL_PAGE_SIZE = ZEND_MM_PAGE_SIZE
const ZEND_MM_FD = -1
const ZEND_MM_STAT = 1
const ZEND_MM_LIMIT = 1
const ZEND_MM_CUSTOM = 1
const ZEND_MM_STORAGE = 1
const ZEND_MM_ERROR = 1

type ZendMmPageInfo = uint32
type ZendMmBitset = ZendUlong

const ZEND_MM_BITSET_LEN = b.SizeOf("zend_mm_bitset") * 8
const ZEND_MM_PAGE_MAP_LEN = ZEND_MM_PAGES / ZEND_MM_BITSET_LEN

type ZendMmPageMap []ZendMmBitset

const ZEND_MM_IS_FRUN = 0x0
const ZEND_MM_IS_LRUN = 0x40000000
const ZEND_MM_IS_SRUN = 0x80000000
const ZEND_MM_LRUN_PAGES_MASK = 0x3ff
const ZEND_MM_LRUN_PAGES_OFFSET = 0
const ZEND_MM_SRUN_BIN_NUM_MASK = 0x1f
const ZEND_MM_SRUN_BIN_NUM_OFFSET = 0
const ZEND_MM_SRUN_FREE_COUNTER_MASK = 0x1ff0000
const ZEND_MM_SRUN_FREE_COUNTER_OFFSET = 16
const ZEND_MM_NRUN_OFFSET_MASK = 0x1ff0000
const ZEND_MM_NRUN_OFFSET_OFFSET = 16
const ZEND_MM_BINS = 30

var ZendMmUseHugePages int = 0
var BinDataSize []uint32 = []uint32{8, 16, 24, 32, 40, 48, 56, 64, 80, 96, 112, 128, 160, 192, 224, 256, 320, 384, 448, 512, 640, 768, 896, 1024, 1280, 1536, 1792, 2048, 2560, 3072}
var BinElements []uint32 = []uint32{512, 256, 170, 128, 102, 85, 73, 64, 51, 42, 36, 32, 25, 21, 18, 16, 64, 32, 9, 8, 32, 16, 9, 8, 16, 8, 16, 8, 8, 4}
var BinPages []uint32 = []uint32{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 3, 1, 1, 5, 3, 2, 2, 5, 3, 7, 4, 5, 3}

var AllocGlobals ZendAllocGlobals
