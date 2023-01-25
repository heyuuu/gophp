// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/core"
)

// Source: <Zend/zend_alloc.h>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

// failed # include "../TSRM/TSRM.h"

const ZEND_MM_ALIGNMENT_MASK = ^(core.ZEND_MM_ALIGNMENT - 1)

const ZEND_MM_OVERHEAD = 0

/* _emalloc() & _efree() specialization */

const _emallocLarge = _emalloc
const _emallocHuge = _emalloc
const _efreeLarge = _efree
const _efreeHuge = _efree

/* Standard wrapper macros */

/* Relay wrapper macros */

/* Selective persistent/non persistent allocation macros */

/* fast cache for HashTables */

/* Heap functions */

const ZEND_MM_CUSTOM_HEAP_NONE = 0
const ZEND_MM_CUSTOM_HEAP_STD = 1
const ZEND_MM_CUSTOM_HEAP_DEBUG = 2

type ZendMmChunkAllocT func(storage *ZendMmStorage, size int, alignment int) any
type ZendMmChunkFreeT func(storage *ZendMmStorage, chunk any, size int)
type ZendMmChunkTruncateT func(storage *ZendMmStorage, chunk any, old_size int, new_size int) int
type ZendMmChunkExtendT func(storage *ZendMmStorage, chunk any, old_size int, new_size int) int

/*

// The following example shows how to use zend_mm_heap API with custom storage

static zend_mm_heap *apc_heap = NULL;
static HashTable    *apc_ht = NULL;

typedef struct _apc_data {
    void     *mem;
    uint32_t  free_pages;
} apc_data;

static void *apc_chunk_alloc(zend_mm_storage *storage, size_t size, size_t alignment)
{
    apc_data *data = (apc_data*)(storage->data);
    size_t real_size = ((size + (ZEND_MM_CHUNK_SIZE-1)) & ~(ZEND_MM_CHUNK_SIZE-1));
    uint32_t count = real_size / ZEND_MM_CHUNK_SIZE;
    uint32_t first, last, i;

    ZEND_ASSERT(alignment == ZEND_MM_CHUNK_SIZE);

    for (first = 0; first < 32; first++) {
        if (!(data->free_pages & (1 << first))) {
            last = first;
            do {
                if (last - first == count - 1) {
                    for (i = first; i <= last; i++) {
                        data->free_pages |= (1 << i);
                    }
                    return (void *)(((char*)(data->mem)) + ZEND_MM_CHUNK_SIZE * (1 << first));
                }
                last++;
            } while (last < 32 && !(data->free_pages & (1 << last)));
            first = last;
        }
    }
    return NULL;
}

static void apc_chunk_free(zend_mm_storage *storage, void *chunk, size_t size)
{
    apc_data *data = (apc_data*)(storage->data);
    uint32_t i;

    ZEND_ASSERT(((uintptr_t)chunk & (ZEND_MM_CHUNK_SIZE - 1)) == 0);

    i = ((uintptr_t)chunk - (uintptr_t)(data->mem)) / ZEND_MM_CHUNK_SIZE;
    while (1) {
        data->free_pages &= ~(1 << i);
        if (size <= ZEND_MM_CHUNK_SIZE) {
            break;
        }
        size -= ZEND_MM_CHUNK_SIZE;
    }
}

static void apc_init_heap(void)
{
    zend_mm_handlers apc_handlers = {
        apc_chunk_alloc,
        apc_chunk_free,
        NULL,
        NULL,
    };
    apc_data tmp_data;
    zend_mm_heap *old_heap;

    // Preallocate properly aligned SHM chunks (64MB)
    tmp_data.mem = shm_memalign(ZEND_MM_CHUNK_SIZE, ZEND_MM_CHUNK_SIZE * 32);

    // Initialize temporary storage data
    tmp_data.free_pages = 0;

    // Create heap
    apc_heap = zend_mm_startup_ex(&apc_handlers, &tmp_data, sizeof(tmp_data));

    // Allocate some data in the heap
    old_heap = zend_mm_set_heap(apc_heap);
    ALLOC_HASHTABLE(apc_ht);
    zend_hash_init(apc_ht, 64, NULL, ZVAL_PTR_DTOR, 0);
    zend_mm_set_heap(old_heap);
}

*/

// Source: <Zend/zend_alloc.c>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

const MAP_FAILED = any(-1)
const MAP_POPULATE = 0
const REAL_PAGE_SIZE = ZEND_MM_PAGE_SIZE

/* NetBSD has an mremap() function with a signature that is incompatible with Linux (WTF?),
 * so pretend it doesn't exist. */

const ZEND_MM_FD = -1
const ZEND_MM_STAT = 1
const ZEND_MM_LIMIT = 1
const ZEND_MM_CUSTOM = 1

/* USE_ZEND_ALLOC=0 may switch to system malloc() */

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

/*
 * Memory is retrieved from OS by chunks of fixed size 2MB.
 * Inside chunk it's managed by pages of fixed size 4096B.
 * So each chunk consists from 512 pages.
 * The first page of each chunk is reserved for chunk header.
 * It contains service information about all pages.
 *
 * free_pages - current number of free pages in this chunk
 *
 * free_tail  - number of continuous free pages at the end of chunk
 *
 * free_map   - bitset (a bit for each page). The bit is set if the corresponding
 *              page is allocated. Allocator for "lage sizes" may easily find a
 *              free page (or a continuous number of pages) searching for zero
 *              bits.
 *
 * map        - contains service information for each page. (32-bits for each
 *              page).
 *    usage:
 *                (2 bits)
 *                 FRUN - free page,
 *              LRUN - first page of "large" allocation
 *              SRUN - first page of a bin used for "small" allocation
 *
 *    lrun_pages:
 *              (10 bits) number of allocated pages
 *
 *    srun_bin_num:
 *              (5 bits) bin number (e.g. 0 for sizes 0-2, 1 for 3-4,
 *               2 for 5-8, 3 for 9-16 etc) see zend_alloc_sizes.h
 */

/*
 * bin - is one or few continuous pages (up to 8) used for allocation of
 * a particular "small size".
 */

var BinDataSize []uint32 = []uint32{8, 16, 24, 32, 40, 48, 56, 64, 80, 96, 112, 128, 160, 192, 224, 256, 320, 384, 448, 512, 640, 768, 896, 1024, 1280, 1536, 1792, 2048, 2560, 3072}
var BinElements []uint32 = []uint32{512, 256, 170, 128, 102, 85, 73, 64, 51, 42, 36, 32, 25, 21, 18, 16, 64, 32, 9, 8, 32, 16, 9, 8, 16, 8, 16, 8, 8, 4}
var BinPages []uint32 = []uint32{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 3, 1, 1, 5, 3, 2, 2, 5, 3, 7, 4, 5, 3}

/*****************/

/***********/

/**********/

/***********************/

/**************/

/**************/

/********/

/*********************/

/******************/

/**************/

/**********************/

var AllocGlobals ZendAllocGlobals
