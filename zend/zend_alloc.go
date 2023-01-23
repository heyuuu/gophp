// <<generate>>

package zend

import (
	r "sik/runtime"
	g "sik/runtime/grammar"
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

// #define ZEND_ALLOC_H

// # include < stdio . h >

// failed # include "../TSRM/TSRM.h"

// # include "zend.h"

// #define ZEND_MM_ALIGNMENT_MASK       ~ ( ZEND_MM_ALIGNMENT - 1 )

// #define ZEND_MM_ALIGNED_SIZE(size) ( ( ( size ) + ZEND_MM_ALIGNMENT - 1 ) & ZEND_MM_ALIGNMENT_MASK )

// #define ZEND_MM_ALIGNED_SIZE_EX(size,alignment) ( ( ( size ) + ( ( alignment ) - 1 ) ) & ~ ( ( alignment ) - 1 ) )

// @type ZendLeakInfo struct

// #define ZEND_MM_OVERHEAD       0

// # include "zend_alloc_sizes.h"

/* _emalloc() & _efree() specialization */

// #define efree_size(ptr,size) efree ( ptr )

// #define efree_size_rel(ptr,size) efree_rel ( ptr )

// #define _emalloc_large       _emalloc

// #define _emalloc_huge       _emalloc

// #define _efree_large       _efree

// #define _efree_huge       _efree

/* Standard wrapper macros */

// #define emalloc(size) _emalloc ( ( size ) ZEND_FILE_LINE_CC ZEND_FILE_LINE_EMPTY_CC )

// #define emalloc_large(size) _emalloc_large ( ( size ) ZEND_FILE_LINE_CC ZEND_FILE_LINE_EMPTY_CC )

// #define emalloc_huge(size) _emalloc_huge ( ( size ) ZEND_FILE_LINE_CC ZEND_FILE_LINE_EMPTY_CC )

// #define safe_emalloc(nmemb,size,offset) _safe_emalloc ( ( nmemb ) , ( size ) , ( offset ) ZEND_FILE_LINE_CC ZEND_FILE_LINE_EMPTY_CC )

// #define efree(ptr) _efree ( ( ptr ) ZEND_FILE_LINE_CC ZEND_FILE_LINE_EMPTY_CC )

// #define efree_large(ptr) _efree_large ( ( ptr ) ZEND_FILE_LINE_CC ZEND_FILE_LINE_EMPTY_CC )

// #define efree_huge(ptr) _efree_huge ( ( ptr ) ZEND_FILE_LINE_CC ZEND_FILE_LINE_EMPTY_CC )

// #define ecalloc(nmemb,size) _ecalloc ( ( nmemb ) , ( size ) ZEND_FILE_LINE_CC ZEND_FILE_LINE_EMPTY_CC )

// #define erealloc(ptr,size) _erealloc ( ( ptr ) , ( size ) ZEND_FILE_LINE_CC ZEND_FILE_LINE_EMPTY_CC )

// #define erealloc2(ptr,size,copy_size) _erealloc2 ( ( ptr ) , ( size ) , ( copy_size ) ZEND_FILE_LINE_CC ZEND_FILE_LINE_EMPTY_CC )

// #define safe_erealloc(ptr,nmemb,size,offset) _safe_erealloc ( ( ptr ) , ( nmemb ) , ( size ) , ( offset ) ZEND_FILE_LINE_CC ZEND_FILE_LINE_EMPTY_CC )

// #define erealloc_recoverable(ptr,size) _erealloc ( ( ptr ) , ( size ) ZEND_FILE_LINE_CC ZEND_FILE_LINE_EMPTY_CC )

// #define erealloc2_recoverable(ptr,size,copy_size) _erealloc2 ( ( ptr ) , ( size ) , ( copy_size ) ZEND_FILE_LINE_CC ZEND_FILE_LINE_EMPTY_CC )

// #define estrdup(s) _estrdup ( ( s ) ZEND_FILE_LINE_CC ZEND_FILE_LINE_EMPTY_CC )

// #define estrndup(s,length) _estrndup ( ( s ) , ( length ) ZEND_FILE_LINE_CC ZEND_FILE_LINE_EMPTY_CC )

// #define zend_mem_block_size(ptr) _zend_mem_block_size ( ( ptr ) ZEND_FILE_LINE_CC ZEND_FILE_LINE_EMPTY_CC )

/* Relay wrapper macros */

// #define emalloc_rel(size) _emalloc ( ( size ) ZEND_FILE_LINE_RELAY_CC ZEND_FILE_LINE_CC )

// #define safe_emalloc_rel(nmemb,size,offset) _safe_emalloc ( ( nmemb ) , ( size ) , ( offset ) ZEND_FILE_LINE_RELAY_CC ZEND_FILE_LINE_CC )

// #define efree_rel(ptr) _efree ( ( ptr ) ZEND_FILE_LINE_RELAY_CC ZEND_FILE_LINE_CC )

// #define ecalloc_rel(nmemb,size) _ecalloc ( ( nmemb ) , ( size ) ZEND_FILE_LINE_RELAY_CC ZEND_FILE_LINE_CC )

// #define erealloc_rel(ptr,size) _erealloc ( ( ptr ) , ( size ) ZEND_FILE_LINE_RELAY_CC ZEND_FILE_LINE_CC )

// #define erealloc2_rel(ptr,size,copy_size) _erealloc2 ( ( ptr ) , ( size ) , ( copy_size ) ZEND_FILE_LINE_RELAY_CC ZEND_FILE_LINE_CC )

// #define erealloc_recoverable_rel(ptr,size) _erealloc ( ( ptr ) , ( size ) ZEND_FILE_LINE_RELAY_CC ZEND_FILE_LINE_CC )

// #define erealloc2_recoverable_rel(ptr,size,copy_size) _erealloc2 ( ( ptr ) , ( size ) , ( copy_size ) ZEND_FILE_LINE_RELAY_CC ZEND_FILE_LINE_CC )

// #define safe_erealloc_rel(ptr,nmemb,size,offset) _safe_erealloc ( ( ptr ) , ( nmemb ) , ( size ) , ( offset ) ZEND_FILE_LINE_RELAY_CC ZEND_FILE_LINE_CC )

// #define estrdup_rel(s) _estrdup ( ( s ) ZEND_FILE_LINE_RELAY_CC ZEND_FILE_LINE_CC )

// #define estrndup_rel(s,length) _estrndup ( ( s ) , ( length ) ZEND_FILE_LINE_RELAY_CC ZEND_FILE_LINE_CC )

// #define zend_mem_block_size_rel(ptr) _zend_mem_block_size ( ( ptr ) ZEND_FILE_LINE_RELAY_CC ZEND_FILE_LINE_CC )

/* Selective persistent/non persistent allocation macros */

// #define pemalloc(size,persistent) ( ( persistent ) ? __zend_malloc ( size ) : emalloc ( size ) )

// #define safe_pemalloc(nmemb,size,offset,persistent) ( ( persistent ) ? _safe_malloc ( nmemb , size , offset ) : safe_emalloc ( nmemb , size , offset ) )

// #define pefree(ptr,persistent) ( ( persistent ) ? free ( ptr ) : efree ( ptr ) )

// #define pefree_size(ptr,size,persistent) do { if ( persistent ) { free ( ptr ) ; } else { efree_size ( ptr , size ) ; } } while ( 0 )

// #define pecalloc(nmemb,size,persistent) ( ( persistent ) ? __zend_calloc ( ( nmemb ) , ( size ) ) : ecalloc ( ( nmemb ) , ( size ) ) )

// #define perealloc(ptr,size,persistent) ( ( persistent ) ? __zend_realloc ( ( ptr ) , ( size ) ) : erealloc ( ( ptr ) , ( size ) ) )

// #define perealloc2(ptr,size,copy_size,persistent) ( ( persistent ) ? __zend_realloc ( ( ptr ) , ( size ) ) : erealloc2 ( ( ptr ) , ( size ) , ( copy_size ) ) )

// #define safe_perealloc(ptr,nmemb,size,offset,persistent) ( ( persistent ) ? _safe_realloc ( ( ptr ) , ( nmemb ) , ( size ) , ( offset ) ) : safe_erealloc ( ( ptr ) , ( nmemb ) , ( size ) , ( offset ) ) )

// #define perealloc_recoverable(ptr,size,persistent) ( ( persistent ) ? realloc ( ( ptr ) , ( size ) ) : erealloc_recoverable ( ( ptr ) , ( size ) ) )

// #define perealloc2_recoverable(ptr,size,persistent) ( ( persistent ) ? realloc ( ( ptr ) , ( size ) ) : erealloc2_recoverable ( ( ptr ) , ( size ) , ( copy_size ) ) )

// #define pestrdup(s,persistent) ( ( persistent ) ? strdup ( s ) : estrdup ( s ) )

// #define pestrndup(s,length,persistent) ( ( persistent ) ? zend_strndup ( ( s ) , ( length ) ) : estrndup ( ( s ) , ( length ) ) )

// #define pemalloc_rel(size,persistent) ( ( persistent ) ? __zend_malloc ( size ) : emalloc_rel ( size ) )

// #define pefree_rel(ptr,persistent) ( ( persistent ) ? free ( ptr ) : efree_rel ( ptr ) )

// #define pecalloc_rel(nmemb,size,persistent) ( ( persistent ) ? __zend_calloc ( ( nmemb ) , ( size ) ) : ecalloc_rel ( ( nmemb ) , ( size ) ) )

// #define perealloc_rel(ptr,size,persistent) ( ( persistent ) ? __zend_realloc ( ( ptr ) , ( size ) ) : erealloc_rel ( ( ptr ) , ( size ) ) )

// #define perealloc2_rel(ptr,size,copy_size,persistent) ( ( persistent ) ? __zend_realloc ( ( ptr ) , ( size ) ) : erealloc2_rel ( ( ptr ) , ( size ) , ( copy_size ) ) )

// #define perealloc_recoverable_rel(ptr,size,persistent) ( ( persistent ) ? realloc ( ( ptr ) , ( size ) ) : erealloc_recoverable_rel ( ( ptr ) , ( size ) ) )

// #define perealloc2_recoverable_rel(ptr,size,copy_size,persistent) ( ( persistent ) ? realloc ( ( ptr ) , ( size ) ) : erealloc2_recoverable_rel ( ( ptr ) , ( size ) , ( copy_size ) ) )

// #define pestrdup_rel(s,persistent) ( ( persistent ) ? strdup ( s ) : estrdup_rel ( s ) )

/* fast cache for HashTables */

// #define ALLOC_HASHTABLE(ht) ( ht ) = ( HashTable * ) emalloc ( sizeof ( HashTable ) )

// #define FREE_HASHTABLE(ht) efree_size ( ht , sizeof ( HashTable ) )

// #define ALLOC_HASHTABLE_REL(ht) ( ht ) = ( HashTable * ) emalloc_rel ( sizeof ( HashTable ) )

// #define FREE_HASHTABLE_REL(ht) efree_size_rel ( ht , sizeof ( HashTable ) )

/* Heap functions */

// #define zend_mm_alloc(heap,size) _zend_mm_alloc ( ( heap ) , ( size ) ZEND_FILE_LINE_CC ZEND_FILE_LINE_EMPTY_CC )

// #define zend_mm_free(heap,p) _zend_mm_free ( ( heap ) , ( p ) ZEND_FILE_LINE_CC ZEND_FILE_LINE_EMPTY_CC )

// #define zend_mm_realloc(heap,p,size) _zend_mm_realloc ( ( heap ) , ( p ) , ( size ) ZEND_FILE_LINE_CC ZEND_FILE_LINE_EMPTY_CC )

// #define zend_mm_realloc2(heap,p,size,copy_size) _zend_mm_realloc2 ( ( heap ) , ( p ) , ( size ) , ( copy_size ) ZEND_FILE_LINE_CC ZEND_FILE_LINE_EMPTY_CC )

// #define zend_mm_block_size(heap,p) _zend_mm_block_size ( ( heap ) , ( p ) ZEND_FILE_LINE_CC ZEND_FILE_LINE_EMPTY_CC )

// #define zend_mm_alloc_rel(heap,size) _zend_mm_alloc ( ( heap ) , ( size ) ZEND_FILE_LINE_RELAY_CC ZEND_FILE_LINE_CC )

// #define zend_mm_free_rel(heap,p) _zend_mm_free ( ( heap ) , ( p ) ZEND_FILE_LINE_RELAY_CC ZEND_FILE_LINE_CC )

// #define zend_mm_realloc_rel(heap,p,size) _zend_mm_realloc ( ( heap ) , ( p ) , ( size ) ZEND_FILE_LINE_RELAY_CC ZEND_FILE_LINE_CC )

// #define zend_mm_realloc2_rel(heap,p,size,copy_size) _zend_mm_realloc2 ( ( heap ) , ( p ) , ( size ) , ( copy_size ) ZEND_FILE_LINE_RELAY_CC ZEND_FILE_LINE_CC )

// #define zend_mm_block_size_rel(heap,p) _zend_mm_block_size ( ( heap ) , ( p ) ZEND_FILE_LINE_CC ZEND_FILE_LINE_EMPTY_CC )

// #define ZEND_MM_CUSTOM_HEAP_NONE       0

// #define ZEND_MM_CUSTOM_HEAP_STD       1

// #define ZEND_MM_CUSTOM_HEAP_DEBUG       2

type ZendMmChunkAllocT func(storage *ZendMmStorage, size int, alignment int) any
type ZendMmChunkFreeT func(storage *ZendMmStorage, chunk any, size int)
type ZendMmChunkTruncateT func(storage *ZendMmStorage, chunk any, old_size int, new_size int) int
type ZendMmChunkExtendT func(storage *ZendMmStorage, chunk any, old_size int, new_size int) int

// @type ZendMmHandlers struct

// @type ZendMmStorage struct

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

// # include "zend.h"

// # include "zend_alloc.h"

// # include "zend_globals.h"

// # include "zend_operators.h"

// # include "zend_multiply.h"

// # include "zend_bitset.h"

// # include < signal . h >

// # include < unistd . h >

// # include < stdio . h >

// # include < stdlib . h >

// # include < string . h >

// # include < sys / types . h >

// # include < sys / stat . h >

// # include < limits . h >

// # include < fcntl . h >

// # include < errno . h >

// # include < sys / mman . h >

// #define MAP_FAILED       ( ( void * ) - 1 )

// #define MAP_POPULATE       0

// #define REAL_PAGE_SIZE       ZEND_MM_PAGE_SIZE

/* NetBSD has an mremap() function with a signature that is incompatible with Linux (WTF?),
 * so pretend it doesn't exist. */

// #define ZEND_MM_FD       - 1

// #define ZEND_MM_STAT       1

// #define ZEND_MM_LIMIT       1

// #define ZEND_MM_CUSTOM       1

/* USE_ZEND_ALLOC=0 may switch to system malloc() */

// #define ZEND_MM_STORAGE       1

// #define ZEND_MM_ERROR       1

// #define ZEND_MM_CHECK(condition,message) do { if ( UNEXPECTED ( ! ( condition ) ) ) { zend_mm_panic ( message ) ; } } while ( 0 )

type ZendMmPageInfo = uint32
type ZendMmBitset = ZendUlong

// #define ZEND_MM_ALIGNED_OFFSET(size,alignment) ( ( ( size_t ) ( size ) ) & ( ( alignment ) - 1 ) )

// #define ZEND_MM_ALIGNED_BASE(size,alignment) ( ( ( size_t ) ( size ) ) & ~ ( ( alignment ) - 1 ) )

// #define ZEND_MM_SIZE_TO_NUM(size,alignment) ( ( ( size_t ) ( size ) + ( ( alignment ) - 1 ) ) / ( alignment ) )

// #define ZEND_MM_BITSET_LEN       ( sizeof ( zend_mm_bitset ) * 8 )

// #define ZEND_MM_PAGE_MAP_LEN       ( ZEND_MM_PAGES / ZEND_MM_BITSET_LEN )

type ZendMmPageMap []ZendMmBitset

// #define ZEND_MM_IS_FRUN       0x00000000

// #define ZEND_MM_IS_LRUN       0x40000000

// #define ZEND_MM_IS_SRUN       0x80000000

// #define ZEND_MM_LRUN_PAGES_MASK       0x000003ff

// #define ZEND_MM_LRUN_PAGES_OFFSET       0

// #define ZEND_MM_SRUN_BIN_NUM_MASK       0x0000001f

// #define ZEND_MM_SRUN_BIN_NUM_OFFSET       0

// #define ZEND_MM_SRUN_FREE_COUNTER_MASK       0x01ff0000

// #define ZEND_MM_SRUN_FREE_COUNTER_OFFSET       16

// #define ZEND_MM_NRUN_OFFSET_MASK       0x01ff0000

// #define ZEND_MM_NRUN_OFFSET_OFFSET       16

// #define ZEND_MM_LRUN_PAGES(info) ( ( ( info ) & ZEND_MM_LRUN_PAGES_MASK ) >> ZEND_MM_LRUN_PAGES_OFFSET )

// #define ZEND_MM_SRUN_BIN_NUM(info) ( ( ( info ) & ZEND_MM_SRUN_BIN_NUM_MASK ) >> ZEND_MM_SRUN_BIN_NUM_OFFSET )

// #define ZEND_MM_SRUN_FREE_COUNTER(info) ( ( ( info ) & ZEND_MM_SRUN_FREE_COUNTER_MASK ) >> ZEND_MM_SRUN_FREE_COUNTER_OFFSET )

// #define ZEND_MM_NRUN_OFFSET(info) ( ( ( info ) & ZEND_MM_NRUN_OFFSET_MASK ) >> ZEND_MM_NRUN_OFFSET_OFFSET )

// #define ZEND_MM_FRUN() ZEND_MM_IS_FRUN

// #define ZEND_MM_LRUN(count) ( ZEND_MM_IS_LRUN | ( ( count ) << ZEND_MM_LRUN_PAGES_OFFSET ) )

// #define ZEND_MM_SRUN(bin_num) ( ZEND_MM_IS_SRUN | ( ( bin_num ) << ZEND_MM_SRUN_BIN_NUM_OFFSET ) )

// #define ZEND_MM_SRUN_EX(bin_num,count) ( ZEND_MM_IS_SRUN | ( ( bin_num ) << ZEND_MM_SRUN_BIN_NUM_OFFSET ) | ( ( count ) << ZEND_MM_SRUN_FREE_COUNTER_OFFSET ) )

// #define ZEND_MM_NRUN(bin_num,offset) ( ZEND_MM_IS_SRUN | ZEND_MM_IS_LRUN | ( ( bin_num ) << ZEND_MM_SRUN_BIN_NUM_OFFSET ) | ( ( offset ) << ZEND_MM_NRUN_OFFSET_OFFSET ) )

// #define ZEND_MM_BINS       30

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

// @type ZendMmHeap struct
// @type ZendMmChunk struct
// @type ZendMmPage struct

/*
 * bin - is one or few continuous pages (up to 8) used for allocation of
 * a particular "small size".
 */

// @type ZendMmBin struct
// @type ZendMmFreeSlot struct
// @type ZendMmHugeList struct

// #define ZEND_MM_PAGE_ADDR(chunk,page_num) ( ( void * ) ( ( ( zend_mm_page * ) ( chunk ) ) + ( page_num ) ) )

// #define _BIN_DATA_SIZE(num,size,elements,pages,x,y) size ,

var BinDataSize []uint32 = []uint32{8, 16, 24, 32, 40, 48, 56, 64, 80, 96, 112, 128, 160, 192, 224, 256, 320, 384, 448, 512, 640, 768, 896, 1024, 1280, 1536, 1792, 2048, 2560, 3072}

// #define _BIN_DATA_ELEMENTS(num,size,elements,pages,x,y) elements ,

var BinElements []uint32 = []uint32{512, 256, 170, 128, 102, 85, 73, 64, 51, 42, 36, 32, 25, 21, 18, 16, 64, 32, 9, 8, 32, 16, 9, 8, 16, 8, 16, 8, 8, 4}

// #define _BIN_DATA_PAGES(num,size,elements,pages,x,y) pages ,

var BinPages []uint32 = []uint32{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 3, 1, 1, 5, 3, 2, 2, 5, 3, 7, 4, 5, 3}

func ZendMmPanic(message string) {
	r.Fprintf(stderr, "%s\n", message)

	/* See http://support.microsoft.com/kb/190351 */

	exit(1)

	/* See http://support.microsoft.com/kb/190351 */
}
func ZendMmSafeError(heap *ZendMmHeap, format string, limit int, size int) {
	heap.SetOverflow(1)
	var __orig_bailout *sigjmp_buf = EG.GetBailout()
	var __bailout sigjmp_buf
	EG.SetBailout(&__bailout)
	if sigsetjmp(__bailout, 0) == 0 {
		ZendErrorNoreturn(1<<0, format, limit, size)
	} else {
		EG.SetBailout(__orig_bailout)
	}
	EG.SetBailout(__orig_bailout)
	heap.SetOverflow(0)
	_zendBailout(__FILE__, __LINE__)
	exit(1)
}

/*****************/

func ZendMmMmapFixed(addr any, size int) any {
	var flags int = MAP_PRIVATE | MAP_ANON

	/* MAP_FIXED leads to discarding of the old mapping, so it can't be used. */

	var ptr any = mmap(addr, size, PROT_READ|PROT_WRITE, flags, -1, 0)
	if ptr == any(-1) {
		r.Fprintf(stderr, "\nmmap() failed: [%d] %s\n", errno, strerror(errno))
		return nil
	} else if ptr != addr {
		if munmap(ptr, size) != 0 {
			r.Fprintf(stderr, "\nmunmap() failed: [%d] %s\n", errno, strerror(errno))
		}
		return nil
	}
	return ptr
}
func ZendMmMmap(size int) any {
	var ptr any
	ptr = mmap(nil, size, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANON, -1, 0)
	if ptr == any(-1) {
		r.Fprintf(stderr, "\nmmap() failed: [%d] %s\n", errno, strerror(errno))
		return nil
	}
	return ptr
}
func ZendMmMunmap(addr any, size int) {
	if munmap(addr, size) != 0 {
		r.Fprintf(stderr, "\nmunmap() failed: [%d] %s\n", errno, strerror(errno))
	}
}

/***********/

func ZendMmBitsetNts(bitset ZendMmBitset) int {
	var n int
	if bitset == zend_mm_bitset-1 {
		return g.SizeOf("zend_mm_bitset") * 8
	}
	n = 0
	if g.SizeOf("zend_mm_bitset") == 8 {
		if (bitset & 0xffffffff) == 0xffffffff {
			n += 32
			bitset = bitset >> 32
		}
	}
	if (bitset & 0xffff) == 0xffff {
		n += 16
		bitset = bitset >> 16
	}
	if (bitset & 0xff) == 0xff {
		n += 8
		bitset = bitset >> 8
	}
	if (bitset & 0xf) == 0xf {
		n += 4
		bitset = bitset >> 4
	}
	if (bitset & 0x3) == 0x3 {
		n += 2
		bitset = bitset >> 2
	}
	return n + (bitset & 1)
}
func ZendMmBitsetIsSet(bitset *ZendMmBitset, bit int) int {
	return bitset[bit/(g.SizeOf("( bitset ) [ 0 ]")*8)] >> (bit&g.SizeOf("( bitset ) [ 0 ]")*8 - 1) & 1
}
func ZendMmBitsetSetBit(bitset *ZendMmBitset, bit int) {
	bitset[bit/(g.SizeOf("zend_mm_bitset")*8)] |= 1 << (bit&g.SizeOf("zend_mm_bitset")*8 - 1)
}
func ZendMmBitsetResetBit(bitset *ZendMmBitset, bit int) {
	bitset[bit/(g.SizeOf("zend_mm_bitset")*8)] &= ^(1 << (bit&g.SizeOf("zend_mm_bitset")*8 - 1))
}
func ZendMmBitsetSetRange(bitset *ZendMmBitset, start int, len_ int) {
	if len_ == 1 {
		ZendMmBitsetSetBit(bitset, start)
	} else {
		var pos int = start / (g.SizeOf("zend_mm_bitset") * 8)
		var end int = (start + len_ - 1) / (g.SizeOf("zend_mm_bitset") * 8)
		var bit int = start&g.SizeOf("zend_mm_bitset")*8 - 1
		var tmp ZendMmBitset
		if pos != end {

			/* set bits from "bit" to ZEND_MM_BITSET_LEN-1 */

			tmp = zend_mm_bitset - 1<<bit
			bitset[g.PostInc(&pos)] |= tmp
			for pos != end {

				/* set all bits */

				bitset[g.PostInc(&pos)] = zend_mm_bitset - 1

				/* set all bits */

			}
			end = start + len_ - 1&g.SizeOf("zend_mm_bitset")*8 - 1

			/* set bits from "0" to "end" */

			tmp = zend_mm_bitset - 1>>g.SizeOf("zend_mm_bitset")*8 - 1 - end
			bitset[pos] |= tmp
		} else {
			end = start + len_ - 1&g.SizeOf("zend_mm_bitset")*8 - 1

			/* set bits from "bit" to "end" */

			tmp = zend_mm_bitset - 1<<bit
			tmp &= zend_mm_bitset - 1>>g.SizeOf("zend_mm_bitset")*8 - 1 - end
			bitset[pos] |= tmp
		}
	}
}
func ZendMmBitsetResetRange(bitset *ZendMmBitset, start int, len_ int) {
	if len_ == 1 {
		ZendMmBitsetResetBit(bitset, start)
	} else {
		var pos int = start / (g.SizeOf("zend_mm_bitset") * 8)
		var end int = (start + len_ - 1) / (g.SizeOf("zend_mm_bitset") * 8)
		var bit int = start&g.SizeOf("zend_mm_bitset")*8 - 1
		var tmp ZendMmBitset
		if pos != end {

			/* reset bits from "bit" to ZEND_MM_BITSET_LEN-1 */

			tmp = ^((1 << bit) - 1)
			bitset[g.PostInc(&pos)] &= ^tmp
			for pos != end {

				/* set all bits */

				bitset[g.PostInc(&pos)] = 0

				/* set all bits */

			}
			end = start + len_ - 1&g.SizeOf("zend_mm_bitset")*8 - 1

			/* reset bits from "0" to "end" */

			tmp = zend_mm_bitset - 1>>g.SizeOf("zend_mm_bitset")*8 - 1 - end
			bitset[pos] &= ^tmp
		} else {
			end = start + len_ - 1&g.SizeOf("zend_mm_bitset")*8 - 1

			/* reset bits from "bit" to "end" */

			tmp = zend_mm_bitset - 1<<bit
			tmp &= zend_mm_bitset - 1>>g.SizeOf("zend_mm_bitset")*8 - 1 - end
			bitset[pos] &= ^tmp
		}
	}
}
func ZendMmBitsetIsFreeRange(bitset *ZendMmBitset, start int, len_ int) int {
	if len_ == 1 {
		return !(ZendMmBitsetIsSet(bitset, start))
	} else {
		var pos int = start / (g.SizeOf("zend_mm_bitset") * 8)
		var end int = (start + len_ - 1) / (g.SizeOf("zend_mm_bitset") * 8)
		var bit int = start&g.SizeOf("zend_mm_bitset")*8 - 1
		var tmp ZendMmBitset
		if pos != end {

			/* set bits from "bit" to ZEND_MM_BITSET_LEN-1 */

			tmp = zend_mm_bitset - 1<<bit
			if (bitset[g.PostInc(&pos)] & tmp) != 0 {
				return 0
			}
			for pos != end {

				/* set all bits */

				if bitset[g.PostInc(&pos)] != 0 {
					return 0
				}

				/* set all bits */

			}
			end = start + len_ - 1&g.SizeOf("zend_mm_bitset")*8 - 1

			/* set bits from "0" to "end" */

			tmp = zend_mm_bitset - 1>>g.SizeOf("zend_mm_bitset")*8 - 1 - end
			return (bitset[pos] & tmp) == 0
		} else {
			end = start + len_ - 1&g.SizeOf("zend_mm_bitset")*8 - 1

			/* set bits from "bit" to "end" */

			tmp = zend_mm_bitset - 1<<bit
			tmp &= zend_mm_bitset - 1>>g.SizeOf("zend_mm_bitset")*8 - 1 - end
			return (bitset[pos] & tmp) == 0
		}
	}
}

/**********/

func ZendMmChunkAllocInt(size int, alignment int) any {
	var ptr any = ZendMmMmap(size)
	if ptr == nil {
		return nil
	} else if (size_t(ptr)&alignment - 1) == 0 {
		return ptr
	} else {
		var offset int

		/* chunk has to be aligned */

		ZendMmMunmap(ptr, size)
		ptr = ZendMmMmap(size + alignment - 4*1024)
		offset = size_t(ptr)&alignment - 1
		if offset != 0 {
			offset = alignment - offset
			ZendMmMunmap(ptr, offset)
			ptr = (*byte)(ptr + offset)
			alignment -= offset
		}
		if alignment > 4*1024 {
			ZendMmMunmap((*byte)(ptr+size), alignment-4*1024)
		}
		return ptr
	}
}
func ZendMmChunkAlloc(heap *ZendMmHeap, size int, alignment int) any {
	if heap.GetStorage() != nil {
		var ptr any = heap.GetStorage().GetHandlers().GetChunkAlloc()(heap.GetStorage(), size, alignment)
		r.Assert((zend_uintptr_t((*byte)(ptr+(alignment-1)))&alignment - 1) == ZendUintptrT(ptr))
		return ptr
	}
	return ZendMmChunkAllocInt(size, alignment)
}
func ZendMmChunkFree(heap *ZendMmHeap, addr any, size int) {
	if heap.GetStorage() != nil {
		heap.GetStorage().GetHandlers().GetChunkFree()(heap.GetStorage(), addr, size)
		return
	}
	ZendMmMunmap(addr, size)
}
func ZendMmChunkTruncate(heap *ZendMmHeap, addr any, old_size int, new_size int) int {
	if heap.GetStorage() != nil {
		if heap.GetStorage().GetHandlers().GetChunkTruncate() != nil {
			return heap.GetStorage().GetHandlers().GetChunkTruncate()(heap.GetStorage(), addr, old_size, new_size)
		} else {
			return 0
		}
	}
	ZendMmMunmap((*byte)(addr+new_size), old_size-new_size)
	return 1
}
func ZendMmChunkExtend(heap *ZendMmHeap, addr any, old_size int, new_size int) int {
	if heap.GetStorage() != nil {
		if heap.GetStorage().GetHandlers().GetChunkExtend() != nil {
			return heap.GetStorage().GetHandlers().GetChunkExtend()(heap.GetStorage(), addr, old_size, new_size)
		} else {
			return 0
		}
	}
	return ZendMmMmapFixed((*byte)(addr+old_size), new_size-old_size) != nil
}
func ZendMmChunkInit(heap *ZendMmHeap, chunk *ZendMmChunk) {
	chunk.SetHeap(heap)
	chunk.SetNext(heap.GetMainChunk())
	chunk.SetPrev(heap.GetMainChunk().GetPrev())
	chunk.GetPrev().SetNext(chunk)
	chunk.GetNext().SetPrev(chunk)

	/* mark first pages as allocated */

	chunk.SetFreePages(2*1024*1024/(4*1024) - 1)
	chunk.SetFreeTail(1)

	/* the younger chunks have bigger number */

	chunk.SetNum(chunk.GetPrev().GetNum() + 1)

	/* mark first pages as allocated */

	chunk.GetFreeMap()[0] = (1 << 1) - 1
	chunk.GetMap()[0] = 0x40000000 | 1<<0
}

/***********************/

/**************/

func ZendMmAllocPages(heap *ZendMmHeap, pages_count uint32) any {
	var chunk *ZendMmChunk = heap.GetMainChunk()
	var page_num uint32
	var len_ uint32
	var steps int = 0
	for true {
		if chunk.GetFreePages() < pages_count {
			goto not_found
		} else {

			/* Best-Fit Search */

			var best int = -1
			var best_len uint32 = 2 * 1024 * 1024 / (4 * 1024)
			var free_tail uint32 = chunk.GetFreeTail()
			var bitset *ZendMmBitset = chunk.GetFreeMap()
			var tmp ZendMmBitset = *(g.PostInc(&bitset))
			var i uint32 = 0
			for true {

				/* skip allocated blocks */

				for tmp == zend_mm_bitset-1 {
					i += g.SizeOf("zend_mm_bitset") * 8
					if i == 2*1024*1024/(4*1024) {
						if best > 0 {
							page_num = best
							goto found
						} else {
							goto not_found
						}
					}
					tmp = *(g.PostInc(&bitset))
				}

				/* find first 0 bit */

				page_num = i + ZendMmBitsetNts(tmp)

				/* reset bits from 0 to "bit" */

				tmp &= tmp + 1

				/* skip free blocks */

				for tmp == 0 {
					i += g.SizeOf("zend_mm_bitset") * 8
					if i >= free_tail || i == 2*1024*1024/(4*1024) {
						len_ = 2*1024*1024/(4*1024) - page_num
						if len_ >= pages_count && len_ < best_len {
							chunk.SetFreeTail(page_num + pages_count)
							goto found
						} else {

							/* set accurate value */

							chunk.SetFreeTail(page_num)
							if best > 0 {
								page_num = best
								goto found
							} else {
								goto not_found
							}
						}
					}
					tmp = *(g.PostInc(&bitset))
				}

				/* find first 1 bit */

				len_ = i + ZendUlongNtz(tmp) - page_num
				if len_ >= pages_count {
					if len_ == pages_count {
						goto found
					} else if len_ < best_len {
						best_len = len_
						best = page_num
					}
				}

				/* set bits from 0 to "bit" */

				tmp |= tmp - 1

				/* set bits from 0 to "bit" */

			}
		}
	not_found:
		if chunk.GetNext() == heap.GetMainChunk() {
		get_chunk:
			if heap.GetCachedChunks() != nil {
				heap.GetCachedChunksCount()--
				chunk = heap.GetCachedChunks()
				heap.SetCachedChunks(chunk.GetNext())
			} else {
				if 2*1024*1024 > heap.GetLimit()-heap.GetRealSize() {
					if ZendMmGc(heap) != 0 {
						goto get_chunk
					} else if heap.GetOverflow() == 0 {
						ZendMmSafeError(heap, "Allowed memory size of %zu bytes exhausted (tried to allocate %zu bytes)", heap.GetLimit(), 4*1024*pages_count)
						return nil
					}
				}
				chunk = (*ZendMmChunk)(ZendMmChunkAlloc(heap, 2*1024*1024, 2*1024*1024))
				if chunk == nil {

					/* insufficient memory */

					if ZendMmGc(heap) != 0 && g.Assign(&chunk, (*ZendMmChunk)(ZendMmChunkAlloc(heap, 2*1024*1024, 2*1024*1024))) != nil {

					} else {
						ZendMmSafeError(heap, "Out of memory (allocated %zu) (tried to allocate %zu bytes)", heap.GetRealSize(), 4*1024*pages_count)
						return nil
					}

					/* insufficient memory */

				}
				var size int = heap.GetRealSize() + 2*1024*1024
				var peak int = g.CondF1(heap.GetRealPeak() > size, func() int { return heap.GetRealPeak() }, size)
				heap.SetRealSize(size)
				heap.SetRealPeak(peak)
			}
			heap.GetChunksCount()++
			if heap.GetChunksCount() > heap.GetPeakChunksCount() {
				heap.SetPeakChunksCount(heap.GetChunksCount())
			}
			ZendMmChunkInit(heap, chunk)
			page_num = 1
			len_ = 2*1024*1024/(4*1024) - 1
			goto found
		} else {
			chunk = chunk.GetNext()
			steps++
		}
	}
found:
	if steps > 2 && pages_count < 8 {

		/* move chunk into the head of the linked-list */

		chunk.GetPrev().SetNext(chunk.GetNext())
		chunk.GetNext().SetPrev(chunk.GetPrev())
		chunk.SetNext(heap.GetMainChunk().GetNext())
		chunk.SetPrev(heap.GetMainChunk())
		chunk.GetPrev().SetNext(chunk)
		chunk.GetNext().SetPrev(chunk)
	}

	/* mark run as allocated */

	chunk.SetFreePages(chunk.GetFreePages() - pages_count)
	ZendMmBitsetSetRange(chunk.GetFreeMap(), page_num, pages_count)
	chunk.GetMap()[page_num] = 0x40000000 | pages_count<<0
	if page_num == chunk.GetFreeTail() {
		chunk.SetFreeTail(page_num + pages_count)
	}
	return any((*ZendMmPage)(chunk) + page_num)
}
func ZendMmAllocLargeEx(heap *ZendMmHeap, size int) any {
	var pages_count int = int((size_t(size) + (4*1024 - 1)) / (4 * 1024))
	var ptr any = ZendMmAllocPages(heap, pages_count)
	var size int = heap.GetSize() + pages_count*(4*1024)
	var peak int = g.CondF1(heap.GetPeak() > size, func() int { return heap.GetPeak() }, size)
	heap.SetSize(size)
	heap.SetPeak(peak)
	return ptr
}
func ZendMmAllocLarge(heap *ZendMmHeap, size int) any { return ZendMmAllocLargeEx(heap, size) }
func ZendMmDeleteChunk(heap *ZendMmHeap, chunk *ZendMmChunk) {
	chunk.GetNext().SetPrev(chunk.GetPrev())
	chunk.GetPrev().SetNext(chunk.GetNext())
	heap.GetChunksCount()--
	if heap.GetChunksCount()+heap.GetCachedChunksCount() < heap.GetAvgChunksCount()+0.1 || heap.GetChunksCount() == heap.GetLastChunksDeleteBoundary() && heap.GetLastChunksDeleteCount() >= 4 {

		/* delay deletion */

		heap.GetCachedChunksCount()++
		chunk.SetNext(heap.GetCachedChunks())
		heap.SetCachedChunks(chunk)
	} else {
		heap.SetRealSize(heap.GetRealSize() - 2*1024*1024)
		if heap.GetCachedChunks() == nil {
			if heap.GetChunksCount() != heap.GetLastChunksDeleteBoundary() {
				heap.SetLastChunksDeleteBoundary(heap.GetChunksCount())
				heap.SetLastChunksDeleteCount(0)
			} else {
				heap.GetLastChunksDeleteCount()++
			}
		}
		if heap.GetCachedChunks() == nil || chunk.GetNum() > heap.GetCachedChunks().GetNum() {
			ZendMmChunkFree(heap, chunk, 2*1024*1024)
		} else {

			//TODO: select the best chunk to delete???

			chunk.SetNext(heap.GetCachedChunks().GetNext())
			ZendMmChunkFree(heap, heap.GetCachedChunks(), 2*1024*1024)
			heap.SetCachedChunks(chunk)
		}
	}
}
func ZendMmFreePagesEx(heap *ZendMmHeap, chunk *ZendMmChunk, page_num uint32, pages_count uint32, free_chunk int) {
	chunk.SetFreePages(chunk.GetFreePages() + pages_count)
	ZendMmBitsetResetRange(chunk.GetFreeMap(), page_num, pages_count)
	chunk.GetMap()[page_num] = 0
	if chunk.GetFreeTail() == page_num+pages_count {

		/* this setting may be not accurate */

		chunk.SetFreeTail(page_num)

		/* this setting may be not accurate */

	}
	if free_chunk != 0 && chunk != heap.GetMainChunk() && chunk.GetFreePages() == 2*1024*1024/(4*1024)-1 {
		ZendMmDeleteChunk(heap, chunk)
	}
}
func ZendMmFreePages(heap *ZendMmHeap, chunk *ZendMmChunk, page_num int, pages_count int) {
	ZendMmFreePagesEx(heap, chunk, page_num, pages_count, 1)
}
func ZendMmFreeLarge(heap *ZendMmHeap, chunk *ZendMmChunk, page_num int, pages_count int) {
	heap.SetSize(heap.GetSize() - pages_count*(4*1024))
	ZendMmFreePages(heap, chunk, page_num, pages_count)
}

/**************/

func ZendMmSmallSizeToBit(size int) int {
	var n int = 16
	if size <= 0xff {
		n -= 8
		size = size << 8
	}
	if size <= 0xfff {
		n -= 4
		size = size << 4
	}
	if size <= 0x3fff {
		n -= 2
		size = size << 2
	}
	if size <= 0x7fff {
		n -= 1
	}
	return n
}
func ZendMmSmallSizeToBin(size int) int {
	var t1 uint
	var t2 uint
	if size <= 64 {

		/* we need to support size == 0 ... */

		return size - !!size>>3

		/* we need to support size == 0 ... */

	} else {
		t1 = size - 1
		t2 = ZendMmSmallSizeToBit(t1) - 3
		t1 = t1 >> t2
		t2 = t2 - 3
		t2 = t2 << 2
		return int(t1 + t2)
	}
}

// #define ZEND_MM_SMALL_SIZE_TO_BIN(size) zend_mm_small_size_to_bin ( size )

func ZendMmAllocSmallSlow(heap *ZendMmHeap, bin_num uint32) any {
	var chunk *ZendMmChunk
	var page_num int
	var bin *ZendMmBin
	var p *ZendMmFreeSlot
	var end *ZendMmFreeSlot
	bin = (*ZendMmBin)(ZendMmAllocPages(heap, BinPages[bin_num]))
	if bin == nil {

		/* insufficient memory */

		return nil

		/* insufficient memory */

	}
	chunk = (*ZendMmChunk)(size_t(bin) & ^(2*1024*1024 - 1))
	page_num = (size_t(bin)&2*1024*1024 - 1) / (4 * 1024)
	chunk.GetMap()[page_num] = 0x80000000 | bin_num<<0
	if BinPages[bin_num] > 1 {
		var i uint32 = 1
		for {
			chunk.GetMap()[page_num+i] = 0x80000000 | 0x40000000 | bin_num<<0 | i<<16
			i++
			if i >= BinPages[bin_num] {
				break
			}
		}
	}

	/* create a linked list of elements from 1 to last */

	end = (*ZendMmFreeSlot)((*byte)(bin + BinDataSize[bin_num]*(BinElements[bin_num]-1)))
	p = (*ZendMmFreeSlot)((*byte)(bin + BinDataSize[bin_num]))
	heap.GetFreeSlot()[bin_num] = p
	for {
		p.SetNextFreeSlot((*ZendMmFreeSlot)((*byte)(p + BinDataSize[bin_num])))
		p = (*ZendMmFreeSlot)((*byte)(p + BinDataSize[bin_num]))
		if p == end {
			break
		}
	}

	/* terminate list using NULL */

	p.SetNextFreeSlot(nil)

	/* return first element */

	return (*byte)(bin)

	/* return first element */
}
func ZendMmAllocSmall(heap *ZendMmHeap, bin_num int) any {
	var size int = heap.GetSize() + BinDataSize[bin_num]
	var peak int = g.CondF1(heap.GetPeak() > size, func() int { return heap.GetPeak() }, size)
	heap.SetSize(size)
	heap.SetPeak(peak)
	if heap.GetFreeSlot()[bin_num] != nil {
		var p *ZendMmFreeSlot = heap.GetFreeSlot()[bin_num]
		heap.GetFreeSlot()[bin_num] = p.GetNextFreeSlot()
		return any(p)
	} else {
		return ZendMmAllocSmallSlow(heap, bin_num)
	}
}
func ZendMmFreeSmall(heap *ZendMmHeap, ptr any, bin_num int) {
	var p *ZendMmFreeSlot
	heap.SetSize(heap.GetSize() - BinDataSize[bin_num])
	p = (*ZendMmFreeSlot)(ptr)
	p.SetNextFreeSlot(heap.GetFreeSlot()[bin_num])
	heap.GetFreeSlot()[bin_num] = p
}

/********/

func ZendMmAllocHeap(heap *ZendMmHeap, size int) any {
	var ptr any
	if size <= 3072 {
		ptr = ZendMmAllocSmall(heap, ZendMmSmallSizeToBin(size))
		return ptr
	} else if size <= 2*1024*1024-4*1024*1 {
		ptr = ZendMmAllocLarge(heap, size)
		return ptr
	} else {
		return ZendMmAllocHuge(heap, size)
	}
}
func ZendMmFreeHeap(heap *ZendMmHeap, ptr any) {
	var page_offset int = size_t(ptr)&2*1024*1024 - 1
	if page_offset == 0 {
		if ptr != nil {
			ZendMmFreeHuge(heap, ptr)
		}
	} else {
		var chunk *ZendMmChunk = (*ZendMmChunk)(size_t(ptr) & ^(2*1024*1024 - 1))
		var page_num int = int(page_offset / (4 * 1024))
		var info ZendMmPageInfo = chunk.GetMap()[page_num]
		if chunk.GetHeap() != heap {
			ZendMmPanic("zend_mm_heap corrupted")
		}
		if (info & 0x80000000) != 0 {
			ZendMmFreeSmall(heap, ptr, (info&0x1f)>>0)
		} else {
			var pages_count int = (info & 0x3ff) >> 0
			if (size_t(page_offset)&4*1024 - 1) != 0 {
				ZendMmPanic("zend_mm_heap corrupted")
			}
			ZendMmFreeLarge(heap, chunk, page_num, pages_count)
		}
	}
}
func ZendMmSize(heap *ZendMmHeap, ptr any) int {
	var page_offset int = size_t(ptr)&2*1024*1024 - 1
	if page_offset == 0 {
		return ZendMmGetHugeBlockSize(heap, ptr)
	} else {
		var chunk *ZendMmChunk
		var page_num int
		var info ZendMmPageInfo
		chunk = (*ZendMmChunk)(size_t(ptr) & ^(2*1024*1024 - 1))
		page_num = int(page_offset / (4 * 1024))
		info = chunk.GetMap()[page_num]
		if chunk.GetHeap() != heap {
			ZendMmPanic("zend_mm_heap corrupted")
		}
		if (info & 0x80000000) != 0 {
			return BinDataSize[(info&0x1f)>>0]
		} else {
			return ((info & 0x3ff) >> 0) * (4 * 1024)
		}
	}
}
func ZendMmReallocSlow(heap *ZendMmHeap, ptr any, size int, copy_size int) any {
	var ret any
	var orig_peak int = heap.GetPeak()
	ret = ZendMmAllocHeap(heap, size)
	memcpy(ret, ptr, copy_size)
	ZendMmFreeHeap(heap, ptr)
	if orig_peak > heap.GetSize() {
		heap.SetPeak(orig_peak)
	} else {
		heap.SetPeak(heap.GetSize())
	}
	return ret
}
func ZendMmReallocHuge(heap *ZendMmHeap, ptr any, size int, copy_size int) any {
	var old_size int
	var new_size int
	old_size = ZendMmGetHugeBlockSize(heap, ptr)
	if size > 2*1024*1024-4*1024*1 {
		new_size = size + (4*1024-1) & ^(4*1024-1)
		if new_size == old_size {
			ZendMmChangeHugeBlockSize(heap, ptr, new_size)
			return ptr
		} else if new_size < old_size {

			/* unmup tail */

			if ZendMmChunkTruncate(heap, ptr, old_size, new_size) != 0 {
				heap.SetRealSize(heap.GetRealSize() - old_size - new_size)
				heap.SetSize(heap.GetSize() - old_size - new_size)
				ZendMmChangeHugeBlockSize(heap, ptr, new_size)
				return ptr
			}

			/* unmup tail */

		} else {
			if new_size-old_size > heap.GetLimit()-heap.GetRealSize() {
				if ZendMmGc(heap) != 0 && new_size-old_size <= heap.GetLimit()-heap.GetRealSize() {

				} else if heap.GetOverflow() == 0 {
					ZendMmSafeError(heap, "Allowed memory size of %zu bytes exhausted (tried to allocate %zu bytes)", heap.GetLimit(), size)
					return nil
				}
			}

			/* try to map tail right after this block */

			if ZendMmChunkExtend(heap, ptr, old_size, new_size) != 0 {
				heap.SetRealSize(heap.GetRealSize() + new_size - old_size)
				if heap.GetRealPeak() > heap.GetRealSize() {
					heap.SetRealPeak(heap.GetRealPeak())
				} else {
					heap.SetRealPeak(heap.GetRealSize())
				}
				heap.SetSize(heap.GetSize() + new_size - old_size)
				if heap.GetPeak() > heap.GetSize() {
					heap.SetPeak(heap.GetPeak())
				} else {
					heap.SetPeak(heap.GetSize())
				}
				ZendMmChangeHugeBlockSize(heap, ptr, new_size)
				return ptr
			}

			/* try to map tail right after this block */

		}
	}
	return ZendMmReallocSlow(heap, ptr, size, g.Cond(old_size < copy_size, old_size, copy_size))
}
func ZendMmReallocHeap(heap *ZendMmHeap, ptr any, size int, use_copy_size ZendBool, copy_size int) any {
	var page_offset int
	var old_size int
	var new_size int
	var ret any
	page_offset = size_t(ptr)&2*1024*1024 - 1
	if page_offset == 0 {
		if ptr == nil {
			return _zendMmAlloc(heap, size)
		} else {
			return ZendMmReallocHuge(heap, ptr, size, copy_size)
		}
	} else {
		var chunk *ZendMmChunk = (*ZendMmChunk)(size_t(ptr) & ^(2*1024*1024 - 1))
		var page_num int = int(page_offset / (4 * 1024))
		var info ZendMmPageInfo = chunk.GetMap()[page_num]
		if chunk.GetHeap() != heap {
			ZendMmPanic("zend_mm_heap corrupted")
		}
		if (info & 0x80000000) != 0 {
			var old_bin_num int = (info & 0x1f) >> 0
			for {
				old_size = BinDataSize[old_bin_num]

				/* Check if requested size fits into current bin */

				if size <= old_size {

					/* Check if truncation is necessary */

					if old_bin_num > 0 && size < BinDataSize[old_bin_num-1] {

						/* truncation */

						ret = ZendMmAllocSmall(heap, ZendMmSmallSizeToBin(size))
						if use_copy_size != 0 {
							if size < copy_size {
								copy_size = size
							} else {
								copy_size = copy_size
							}
						} else {
							copy_size = size
						}
						memcpy(ret, ptr, copy_size)
						ZendMmFreeSmall(heap, ptr, old_bin_num)
					} else {

						/* reallocation in-place */

						ret = ptr

						/* reallocation in-place */

					}

					/* Check if truncation is necessary */

				} else if size <= 3072 {

					/* small extension */

					var orig_peak int = heap.GetPeak()
					ret = ZendMmAllocSmall(heap, ZendMmSmallSizeToBin(size))
					if use_copy_size != 0 {
						if old_size < copy_size {
							copy_size = old_size
						} else {
							copy_size = copy_size
						}
					} else {
						copy_size = old_size
					}
					memcpy(ret, ptr, copy_size)
					ZendMmFreeSmall(heap, ptr, old_bin_num)
					if orig_peak > heap.GetSize() {
						heap.SetPeak(orig_peak)
					} else {
						heap.SetPeak(heap.GetSize())
					}

					/* small extension */

				} else {

					/* slow reallocation */

					break

					/* slow reallocation */

				}
				return ret
				break
			}
		} else {
			if (size_t(page_offset)&4*1024 - 1) != 0 {
				ZendMmPanic("zend_mm_heap corrupted")
			}
			old_size = ((info & 0x3ff) >> 0) * (4 * 1024)
			if size > 3072 && size <= 2*1024*1024-4*1024*1 {
				new_size = size + (4*1024-1) & ^(4*1024-1)
				if new_size == old_size {
					return ptr
				} else if new_size < old_size {

					/* free tail pages */

					var new_pages_count int = int(new_size / (4 * 1024))
					var rest_pages_count int = int((old_size - new_size) / (4 * 1024))
					heap.SetSize(heap.GetSize() - rest_pages_count*(4*1024))
					chunk.GetMap()[page_num] = 0x40000000 | new_pages_count<<0
					chunk.SetFreePages(chunk.GetFreePages() + rest_pages_count)
					ZendMmBitsetResetRange(chunk.GetFreeMap(), page_num+new_pages_count, rest_pages_count)
					return ptr
				} else {
					var new_pages_count int = int(new_size / (4 * 1024))
					var old_pages_count int = int(old_size / (4 * 1024))

					/* try to allocate tail pages after this block */

					if page_num+new_pages_count <= 2*1024*1024/(4*1024) && ZendMmBitsetIsFreeRange(chunk.GetFreeMap(), page_num+old_pages_count, new_pages_count-old_pages_count) != 0 {
						var size int = heap.GetSize() + (new_size - old_size)
						var peak int = g.CondF1(heap.GetPeak() > size, func() int { return heap.GetPeak() }, size)
						heap.SetSize(size)
						heap.SetPeak(peak)
						chunk.SetFreePages(chunk.GetFreePages() - new_pages_count - old_pages_count)
						ZendMmBitsetSetRange(chunk.GetFreeMap(), page_num+old_pages_count, new_pages_count-old_pages_count)
						chunk.GetMap()[page_num] = 0x40000000 | new_pages_count<<0
						return ptr
					}

					/* try to allocate tail pages after this block */

				}
			}
		}
	}
	if old_size < copy_size {
		copy_size = old_size
	} else {
		copy_size = copy_size
	}
	return ZendMmReallocSlow(heap, ptr, size, copy_size)
}

/*********************/

func ZendMmAddHugeBlock(heap *ZendMmHeap, ptr any, size int) {
	var list *ZendMmHugeList = (*ZendMmHugeList)(ZendMmAllocHeap(heap, g.SizeOf("zend_mm_huge_list")))
	list.SetPtr(ptr)
	list.SetSize(size)
	list.SetNext(heap.GetHugeList())
	heap.SetHugeList(list)
}
func ZendMmDelHugeBlock(heap *ZendMmHeap, ptr any) int {
	var prev *ZendMmHugeList = nil
	var list *ZendMmHugeList = heap.GetHugeList()
	for list != nil {
		if list.GetPtr() == ptr {
			var size int
			if prev != nil {
				prev.SetNext(list.GetNext())
			} else {
				heap.SetHugeList(list.GetNext())
			}
			size = list.GetSize()
			ZendMmFreeHeap(heap, list)
			return size
		}
		prev = list
		list = list.GetNext()
	}
	ZendMmPanic("zend_mm_heap corrupted")
	return 0
}
func ZendMmGetHugeBlockSize(heap *ZendMmHeap, ptr any) int {
	var list *ZendMmHugeList = heap.GetHugeList()
	for list != nil {
		if list.GetPtr() == ptr {
			return list.GetSize()
		}
		list = list.GetNext()
	}
	ZendMmPanic("zend_mm_heap corrupted")
	return 0
}
func ZendMmChangeHugeBlockSize(heap *ZendMmHeap, ptr any, size int) {
	var list *ZendMmHugeList = heap.GetHugeList()
	for list != nil {
		if list.GetPtr() == ptr {
			list.SetSize(size)
			return
		}
		list = list.GetNext()
	}
}
func ZendMmAllocHuge(heap *ZendMmHeap, size int) any {
	var alignment int = 4 * 1024
	var new_size int = size + (alignment-1) & ^(alignment-1)
	var ptr any
	if new_size < size {
		ZendErrorNoreturn(1<<0, "Possible integer overflow in memory allocation (%zu + %zu)", size, alignment)
	}
	if new_size > heap.GetLimit()-heap.GetRealSize() {
		if ZendMmGc(heap) != 0 && new_size <= heap.GetLimit()-heap.GetRealSize() {

		} else if heap.GetOverflow() == 0 {
			ZendMmSafeError(heap, "Allowed memory size of %zu bytes exhausted (tried to allocate %zu bytes)", heap.GetLimit(), size)
			return nil
		}
	}
	ptr = ZendMmChunkAlloc(heap, new_size, 2*1024*1024)
	if ptr == nil {

		/* insufficient memory */

		if ZendMmGc(heap) != 0 && g.Assign(&ptr, ZendMmChunkAlloc(heap, new_size, 2*1024*1024)) != nil {

		} else {
			ZendMmSafeError(heap, "Out of memory (allocated %zu) (tried to allocate %zu bytes)", heap.GetRealSize(), size)
			return nil
		}

		/* insufficient memory */

	}
	ZendMmAddHugeBlock(heap, ptr, new_size)
	var size int = heap.GetRealSize() + new_size
	var peak int = g.CondF1(heap.GetRealPeak() > size, func() int { return heap.GetRealPeak() }, size)
	heap.SetRealSize(size)
	heap.SetRealPeak(peak)
	var size int = heap.GetSize() + new_size
	var peak int = g.CondF1(heap.GetPeak() > size, func() int { return heap.GetPeak() }, size)
	heap.SetSize(size)
	heap.SetPeak(peak)
	return ptr
}
func ZendMmFreeHuge(heap *ZendMmHeap, ptr any) {
	var size int
	if (size_t(ptr)&2*1024*1024 - 1) != 0 {
		ZendMmPanic("zend_mm_heap corrupted")
	}
	size = ZendMmDelHugeBlock(heap, ptr)
	ZendMmChunkFree(heap, ptr, size)
	heap.SetRealSize(heap.GetRealSize() - size)
	heap.SetSize(heap.GetSize() - size)
}

/******************/

func ZendMmInit() *ZendMmHeap {
	var chunk *ZendMmChunk = (*ZendMmChunk)(ZendMmChunkAllocInt(2*1024*1024, 2*1024*1024))
	var heap *ZendMmHeap
	if chunk == nil {
		r.Fprintf(stderr, "\nCan't initialize heap: [%d] %s\n", errno, strerror(errno))
		return nil
	}
	heap = &chunk.heap_slot
	chunk.SetHeap(heap)
	chunk.SetNext(chunk)
	chunk.SetPrev(chunk)
	chunk.SetFreePages(2*1024*1024/(4*1024) - 1)
	chunk.SetFreeTail(1)
	chunk.SetNum(0)
	chunk.GetFreeMap()[0] = (1 << 1) - 1
	chunk.GetMap()[0] = 0x40000000 | 1<<0
	heap.SetMainChunk(chunk)
	heap.SetCachedChunks(nil)
	heap.SetChunksCount(1)
	heap.SetPeakChunksCount(1)
	heap.SetCachedChunksCount(0)
	heap.SetAvgChunksCount(1.0)
	heap.SetLastChunksDeleteBoundary(0)
	heap.SetLastChunksDeleteCount(0)
	heap.SetRealSize(2 * 1024 * 1024)
	heap.SetRealPeak(2 * 1024 * 1024)
	heap.SetSize(0)
	heap.SetPeak(0)
	heap.SetLimit(size_t - 1>>int(1))
	heap.SetOverflow(0)
	heap.SetUseCustomHeap(0)
	heap.SetStorage(nil)
	heap.SetHugeList(nil)
	return heap
}
func ZendMmGc(heap *ZendMmHeap) int {
	var p *ZendMmFreeSlot
	var q **ZendMmFreeSlot
	var chunk *ZendMmChunk
	var page_offset int
	var page_num int
	var info ZendMmPageInfo
	var i uint32
	var free_counter uint32
	var has_free_pages int
	var collected int = 0
	if heap.GetUseCustomHeap() != 0 {
		return 0
	}
	for i = 0; i < 30; i++ {
		has_free_pages = 0
		p = heap.GetFreeSlot()[i]
		for p != nil {
			chunk = (*ZendMmChunk)(size_t(p) & ^(2*1024*1024 - 1))
			if chunk.GetHeap() != heap {
				ZendMmPanic("zend_mm_heap corrupted")
			}
			page_offset = size_t(p)&2*1024*1024 - 1
			r.Assert(page_offset != 0)
			page_num = int(page_offset / (4 * 1024))
			info = chunk.GetMap()[page_num]
			r.Assert((info & 0x80000000) != 0)
			if (info & 0x40000000) != 0 {
				page_num -= (info & 0x1ff0000) >> 16
				info = chunk.GetMap()[page_num]
				r.Assert((info & 0x80000000) != 0)
				r.Assert((info & 0x40000000) == 0)
			}
			r.Assert((info&0x1f)>>0 == i)
			free_counter = ((info & 0x1ff0000) >> 16) + 1
			if free_counter == BinElements[i] {
				has_free_pages = 1
			}
			chunk.GetMap()[page_num] = 0x80000000 | i<<0 | free_counter<<16
			p = p.GetNextFreeSlot()
		}
		if has_free_pages == 0 {
			continue
		}
		q = &heap.free_slot[i]
		p = *q
		for p != nil {
			chunk = (*ZendMmChunk)(size_t(p) & ^(2*1024*1024 - 1))
			if chunk.GetHeap() != heap {
				ZendMmPanic("zend_mm_heap corrupted")
			}
			page_offset = size_t(p)&2*1024*1024 - 1
			r.Assert(page_offset != 0)
			page_num = int(page_offset / (4 * 1024))
			info = chunk.GetMap()[page_num]
			r.Assert((info & 0x80000000) != 0)
			if (info & 0x40000000) != 0 {
				page_num -= (info & 0x1ff0000) >> 16
				info = chunk.GetMap()[page_num]
				r.Assert((info & 0x80000000) != 0)
				r.Assert((info & 0x40000000) == 0)
			}
			r.Assert((info&0x1f)>>0 == i)
			if (info&0x1ff0000)>>16 == BinElements[i] {

				/* remove from cache */

				p = p.GetNextFreeSlot()
				*q = p
			} else {
				q = &p.next_free_slot
				p = *q
			}
		}
	}
	chunk = heap.GetMainChunk()
	for {
		i = 1
		for i < chunk.GetFreeTail() {
			if ZendMmBitsetIsSet(chunk.GetFreeMap(), i) != 0 {
				info = chunk.GetMap()[i]
				if (info & 0x80000000) != 0 {
					var bin_num int = (info & 0x1f) >> 0
					var pages_count int = BinPages[bin_num]
					if (info&0x1ff0000)>>16 == BinElements[bin_num] {

						/* all elements are free */

						ZendMmFreePagesEx(heap, chunk, i, pages_count, 0)
						collected += pages_count
					} else {

						/* reset counter */

						chunk.GetMap()[i] = 0x80000000 | bin_num<<0

						/* reset counter */

					}
					i += BinPages[bin_num]
				} else {
					i += (info & 0x3ff) >> 0
				}
			} else {
				i++
			}
		}
		if chunk.GetFreePages() == 2*1024*1024/(4*1024)-1 {
			var next_chunk *ZendMmChunk = chunk.GetNext()
			ZendMmDeleteChunk(heap, chunk)
			chunk = next_chunk
		} else {
			chunk = chunk.GetNext()
		}
		if chunk == heap.GetMainChunk() {
			break
		}
	}
	return collected * (4 * 1024)
}
func ZendMmShutdown(heap *ZendMmHeap, full int, silent int) {
	var p *ZendMmChunk
	var list *ZendMmHugeList
	if heap.GetUseCustomHeap() != 0 {
		if heap.GetCustomHeapStdMalloc() == TrackedMalloc {
			if silent != 0 {
				TrackedFreeAll()
			}
			ZendHashClean(heap.GetTrackedAllocs())
			if full != 0 {
				ZendHashDestroy(heap.GetTrackedAllocs())
				Free(heap.GetTrackedAllocs())

				/* Make sure the heap free below does not use tracked_free(). */

				heap.SetCustomHeapStdFree(Free)

				/* Make sure the heap free below does not use tracked_free(). */

			}
		}
		if full != 0 {
			heap.GetCustomHeapStdFree()(heap)
		}
		return
	}

	/* free huge blocks */

	list = heap.GetHugeList()
	heap.SetHugeList(nil)
	for list != nil {
		var q *ZendMmHugeList = list
		list = list.GetNext()
		ZendMmChunkFree(heap, q.GetPtr(), q.GetSize())
	}

	/* move all chunks except of the first one into the cache */

	p = heap.GetMainChunk().GetNext()
	for p != heap.GetMainChunk() {
		var q *ZendMmChunk = p.GetNext()
		p.SetNext(heap.GetCachedChunks())
		heap.SetCachedChunks(p)
		p = q
		heap.GetChunksCount()--
		heap.GetCachedChunksCount()++
	}
	if full != 0 {

		/* free all cached chunks */

		for heap.GetCachedChunks() != nil {
			p = heap.GetCachedChunks()
			heap.SetCachedChunks(p.GetNext())
			ZendMmChunkFree(heap, p, 2*1024*1024)
		}

		/* free the first chunk */

		ZendMmChunkFree(heap, heap.GetMainChunk(), 2*1024*1024)

		/* free the first chunk */

	} else {

		/* free some cached chunks to keep average count */

		heap.SetAvgChunksCount((heap.GetAvgChunksCount() + float64(heap.GetPeakChunksCount())) / 2.0)
		for float64(heap.GetCachedChunksCount()+0.9 > heap.GetAvgChunksCount() && heap.GetCachedChunks() != nil) {
			p = heap.GetCachedChunks()
			heap.SetCachedChunks(p.GetNext())
			ZendMmChunkFree(heap, p, 2*1024*1024)
			heap.GetCachedChunksCount()--
		}

		/* clear cached chunks */

		p = heap.GetCachedChunks()
		for p != nil {
			var q *ZendMmChunk = p.GetNext()
			memset(p, 0, g.SizeOf("zend_mm_chunk"))
			p.SetNext(q)
			p = q
		}

		/* reinitialize the first chunk and heap */

		p = heap.GetMainChunk()
		p.SetHeap(&p.heap_slot)
		p.SetNext(p)
		p.SetPrev(p)
		p.SetFreePages(2*1024*1024/(4*1024) - 1)
		p.SetFreeTail(1)
		p.SetNum(0)
		heap.SetPeak(0)
		heap.SetSize(heap.GetPeak())
		memset(heap.GetFreeSlot(), 0, g.SizeOf("heap -> free_slot"))
		heap.SetRealSize(2 * 1024 * 1024)
		heap.SetRealPeak(2 * 1024 * 1024)
		heap.SetChunksCount(1)
		heap.SetPeakChunksCount(1)
		heap.SetLastChunksDeleteBoundary(0)
		heap.SetLastChunksDeleteCount(0)
		memset(p.GetFreeMap(), 0, g.SizeOf("p -> free_map")+g.SizeOf("p -> map"))
		p.GetFreeMap()[0] = (1 << 1) - 1
		p.GetMap()[0] = 0x40000000 | 1<<0
	}
}

/**************/

func _zendMmAlloc(heap *ZendMmHeap, size int) any { return ZendMmAllocHeap(heap, size) }
func _zendMmFree(heap *ZendMmHeap, ptr any)       { ZendMmFreeHeap(heap, ptr) }
func _zendMmRealloc(heap *ZendMmHeap, ptr any, size int) any {
	return ZendMmReallocHeap(heap, ptr, size, 0, size)
}
func _zendMmRealloc2(heap *ZendMmHeap, ptr any, size int, copy_size int) any {
	return ZendMmReallocHeap(heap, ptr, size, 1, copy_size)
}
func _zendMmBlockSize(heap *ZendMmHeap, ptr any) int { return ZendMmSize(heap, ptr) }

/**********************/

// @type ZendAllocGlobals struct

// #define AG(v) ( alloc_globals . v )

var AllocGlobals ZendAllocGlobals

func IsZendMm() int {
	return !(AllocGlobals.GetMmHeap().GetUseCustomHeap())
}
func IsZendPtr(ptr any) int {
	if AllocGlobals.GetMmHeap().GetUseCustomHeap() != 0 {
		return 0
	}
	if AllocGlobals.GetMmHeap().GetMainChunk() != nil {
		var chunk *ZendMmChunk = AllocGlobals.GetMmHeap().GetMainChunk()
		for {
			if ptr >= any(chunk != nil && ptr < any((*byte)(chunk+2*1024*1024))) {
				return 1
			}
			chunk = chunk.GetNext()
			if chunk == AllocGlobals.GetMmHeap().GetMainChunk() {
				break
			}
		}
	}
	if AllocGlobals.GetMmHeap().GetHugeList() != nil {
		var block *ZendMmHugeList = AllocGlobals.GetMmHeap().GetHugeList()
		for {
			if ptr >= any(block != nil && ptr < any((*byte)(block+block.GetSize()))) {
				return 1
			}
			block = block.GetNext()
			if block == AllocGlobals.GetMmHeap().GetHugeList() {
				break
			}
		}
	}
	return 0
}
func _mallocCustom(size int) any {
	return AllocGlobals.GetMmHeap().GetCustomHeapStdMalloc()(size)
}
func _efreeCustom(ptr any) {
	AllocGlobals.GetMmHeap().GetCustomHeapStdFree()(ptr)
}
func _reallocCustom(ptr any, size int) any {
	return AllocGlobals.GetMmHeap().GetCustomHeapStdRealloc()(ptr, size)
}
func _emalloc(size int) any {
	if AllocGlobals.GetMmHeap().GetUseCustomHeap() != 0 {
		return _mallocCustom(size)
	}
	return ZendMmAllocHeap(AllocGlobals.GetMmHeap(), size)
}
func _efree(ptr any) {
	if AllocGlobals.GetMmHeap().GetUseCustomHeap() != 0 {
		_efreeCustom(ptr)
		return
	}
	ZendMmFreeHeap(AllocGlobals.GetMmHeap(), ptr)
}
func _erealloc(ptr any, size int) any {
	if AllocGlobals.GetMmHeap().GetUseCustomHeap() != 0 {
		return _reallocCustom(ptr, size)
	}
	return ZendMmReallocHeap(AllocGlobals.GetMmHeap(), ptr, size, 0, size)
}
func _erealloc2(ptr any, size int, copy_size int) any {
	if AllocGlobals.GetMmHeap().GetUseCustomHeap() != 0 {
		return _reallocCustom(ptr, size)
	}
	return ZendMmReallocHeap(AllocGlobals.GetMmHeap(), ptr, size, 1, copy_size)
}
func _zendMemBlockSize(ptr any) int {
	if AllocGlobals.GetMmHeap().GetUseCustomHeap() != 0 {
		return 0
	}
	return ZendMmSize(AllocGlobals.GetMmHeap(), ptr)
}
func _safeEmalloc(nmemb int, size int, offset int) any {
	return _emalloc(ZendSafeAddressGuarded(nmemb, size, offset))
}
func _safeMalloc(nmemb int, size int, offset int) any {
	return __zendMalloc(ZendSafeAddressGuarded(nmemb, size, offset))
}
func _safeErealloc(ptr any, nmemb int, size int, offset int) any {
	return _erealloc(ptr, ZendSafeAddressGuarded(nmemb, size, offset))
}
func _safeRealloc(ptr any, nmemb int, size int, offset int) any {
	return __zendRealloc(ptr, ZendSafeAddressGuarded(nmemb, size, offset))
}
func _ecalloc(nmemb int, size int) any {
	var p any
	size = ZendSafeAddressGuarded(nmemb, size, 0)
	p = _emalloc(size)
	memset(p, 0, size)
	return p
}
func _estrdup(s string) *byte {
	var length int
	var p *byte
	length = strlen(s)
	if length+1 == 0 {
		ZendErrorNoreturn(1<<0, "Possible integer overflow in memory allocation (1 * %zu + 1)", length)
	}
	p = (*byte)(_emalloc(length + 1))
	memcpy(p, s, length+1)
	return p
}
func _estrndup(s *byte, length int) *byte {
	var p *byte
	if length+1 == 0 {
		ZendErrorNoreturn(1<<0, "Possible integer overflow in memory allocation (1 * %zu + 1)", length)
	}
	p = (*byte)(_emalloc(length + 1))
	memcpy(p, s, length)
	p[length] = 0
	return p
}
func ZendStrndup(s *byte, length int) *byte {
	var p *byte
	if length+1 == 0 {
		ZendErrorNoreturn(1<<0, "Possible integer overflow in memory allocation (1 * %zu + 1)", length)
	}
	p = (*byte)(Malloc(length + 1))
	if p == nil {
		return p
	}
	if length != 0 {
		memcpy(p, s, length)
	}
	p[length] = 0
	return p
}
func ZendSetMemoryLimit(memory_limit int) int {
	if memory_limit < 2*1024*1024 {
		memory_limit = 2 * 1024 * 1024
	}
	if memory_limit < AllocGlobals.GetMmHeap().GetRealSize() {
		return FAILURE
	}
	AllocGlobals.GetMmHeap().SetLimit(memory_limit)
	return SUCCESS
}
func ZendMemoryUsage(real_usage int) int {
	if real_usage != 0 {
		return AllocGlobals.GetMmHeap().GetRealSize()
	} else {
		var usage int = AllocGlobals.GetMmHeap().GetSize()
		return usage
	}
	return 0
}
func ZendMemoryPeakUsage(real_usage int) int {
	if real_usage != 0 {
		return AllocGlobals.GetMmHeap().GetRealPeak()
	} else {
		return AllocGlobals.GetMmHeap().GetPeak()
	}
	return 0
}
func ShutdownMemoryManager(silent int, full_shutdown int) {
	ZendMmShutdown(AllocGlobals.GetMmHeap(), full_shutdown, silent)
}
func TrackedMalloc(size int) any {
	var ptr any = __zendMalloc(size)
	var h ZendUlong = uintPtr(ptr) >> 3
	r.Assert(any(uintptr_t(h<<3) == ptr))
	ZendHashIndexAddEmptyElement(AllocGlobals.GetMmHeap().GetTrackedAllocs(), h)
	return ptr
}
func TrackedFree(ptr any) {
	var h ZendUlong = uintPtr(ptr) >> 3
	ZendHashIndexDel(AllocGlobals.GetMmHeap().GetTrackedAllocs(), h)
	Free(ptr)
}
func TrackedRealloc(ptr any, new_size int) any {
	var h ZendUlong = uintPtr(ptr) >> 3
	ZendHashIndexDel(AllocGlobals.GetMmHeap().GetTrackedAllocs(), h)
	ptr = __zendRealloc(ptr, new_size)
	h = uintPtr(ptr) >> 3
	r.Assert(any(uintptr_t(h<<3) == ptr))
	ZendHashIndexAddEmptyElement(AllocGlobals.GetMmHeap().GetTrackedAllocs(), h)
	return ptr
}
func TrackedFreeAll() {
	var tracked_allocs *HashTable = AllocGlobals.GetMmHeap().GetTrackedAllocs()
	var h ZendUlong
	for {
		var __ht *HashTable = tracked_allocs
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			h = _p.GetH()
			var ptr any = any(uintptr_t(h << 3))
			Free(ptr)
		}
		break
	}
}
func AllocGlobalsCtor(alloc_globals *ZendAllocGlobals) {
	var tmp *byte
	tmp = getenv("USE_ZEND_ALLOC")
	if tmp != nil && ZendAtoi(tmp, 0) == 0 {
		var tracked ZendBool = g.Assign(&tmp, getenv("USE_TRACKED_ALLOC")) && ZendAtoi(tmp, 0) != 0
		var mm_heap *ZendMmHeap = g.Assign(&(alloc_globals.GetMmHeap()), Malloc(g.SizeOf("zend_mm_heap")))
		memset(mm_heap, 0, g.SizeOf("zend_mm_heap"))
		mm_heap.SetUseCustomHeap(1)
		if tracked == 0 {

			/* Use system allocator. */

			mm_heap.SetCustomHeapStdMalloc(__zendMalloc)
			mm_heap.SetCustomHeapStdFree(Free)
			mm_heap.SetCustomHeapStdRealloc(__zendRealloc)
		} else {

			/* Use system allocator and track allocations for auto-free. */

			mm_heap.SetCustomHeapStdMalloc(TrackedMalloc)
			mm_heap.SetCustomHeapStdFree(TrackedFree)
			mm_heap.SetCustomHeapStdRealloc(TrackedRealloc)
			mm_heap.SetTrackedAllocs(Malloc(g.SizeOf("HashTable")))
			_zendHashInit(mm_heap.GetTrackedAllocs(), 1024, nil, 1)
		}
		return
	}
	tmp = getenv("USE_ZEND_ALLOC_HUGE_PAGES")
	if tmp != nil && ZendAtoi(tmp, 0) != 0 {
		ZendMmUseHugePages = 1
	}
	alloc_globals.SetMmHeap(ZendMmInit())
}
func StartMemoryManager() { AllocGlobalsCtor(&AllocGlobals) }
func ZendMmSetHeap(new_heap *ZendMmHeap) *ZendMmHeap {
	var old_heap *ZendMmHeap
	old_heap = AllocGlobals.GetMmHeap()
	AllocGlobals.SetMmHeap((*ZendMmHeap)(new_heap))
	return (*ZendMmHeap)(old_heap)
}
func ZendMmGetHeap() *ZendMmHeap { return AllocGlobals.GetMmHeap() }
func ZendMmIsCustomHeap(new_heap *ZendMmHeap) int {
	return AllocGlobals.GetMmHeap().GetUseCustomHeap()
}
func ZendMmSetCustomHandlers(heap *ZendMmHeap, _malloc func(int) any, _free func(any), _realloc func(any, int) any) {
	var _heap *ZendMmHeap = (*ZendMmHeap)(heap)
	if _malloc == nil && _free == nil && _realloc == nil {
		_heap.SetUseCustomHeap(0)
	} else {
		_heap.SetUseCustomHeap(1)
		_heap.SetCustomHeapStdMalloc(_malloc)
		_heap.SetCustomHeapStdFree(_free)
		_heap.SetCustomHeapStdRealloc(_realloc)
	}
}
func ZendMmGetCustomHandlers(heap *ZendMmHeap, _malloc *func(int) any, _free *func(any), _realloc *func(any, int) any) {
	var _heap *ZendMmHeap = (*ZendMmHeap)(heap)
	if heap.GetUseCustomHeap() != 0 {
		*_malloc = _heap.GetCustomHeapStdMalloc()
		*_free = _heap.GetCustomHeapStdFree()
		*_realloc = _heap.GetCustomHeapStdRealloc()
	} else {
		*_malloc = nil
		*_free = nil
		*_realloc = nil
	}
}
func ZendMmGetStorage(heap *ZendMmHeap) *ZendMmStorage { return heap.GetStorage() }
func ZendMmStartup() *ZendMmHeap                       { return ZendMmInit() }
func ZendMmStartupEx(handlers *ZendMmHandlers, data any, data_size int) *ZendMmHeap {
	var tmp_storage ZendMmStorage
	var storage *ZendMmStorage
	var chunk *ZendMmChunk
	var heap *ZendMmHeap
	memcpy((*ZendMmHandlers)(&tmp_storage.handlers), handlers, g.SizeOf("zend_mm_handlers"))
	tmp_storage.SetData(data)
	chunk = (*ZendMmChunk)(handlers.GetChunkAlloc()(&tmp_storage, 2*1024*1024, 2*1024*1024))
	if chunk == nil {
		r.Fprintf(stderr, "\nCan't initialize heap: [%d] %s\n", errno, strerror(errno))
		return nil
	}
	heap = &chunk.heap_slot
	chunk.SetHeap(heap)
	chunk.SetNext(chunk)
	chunk.SetPrev(chunk)
	chunk.SetFreePages(2*1024*1024/(4*1024) - 1)
	chunk.SetFreeTail(1)
	chunk.SetNum(0)
	chunk.GetFreeMap()[0] = (1 << 1) - 1
	chunk.GetMap()[0] = 0x40000000 | 1<<0
	heap.SetMainChunk(chunk)
	heap.SetCachedChunks(nil)
	heap.SetChunksCount(1)
	heap.SetPeakChunksCount(1)
	heap.SetCachedChunksCount(0)
	heap.SetAvgChunksCount(1.0)
	heap.SetLastChunksDeleteBoundary(0)
	heap.SetLastChunksDeleteCount(0)
	heap.SetRealSize(2 * 1024 * 1024)
	heap.SetRealPeak(2 * 1024 * 1024)
	heap.SetSize(0)
	heap.SetPeak(0)
	heap.SetLimit(-1 >> 1)
	heap.SetOverflow(0)
	heap.SetUseCustomHeap(0)
	heap.SetStorage(&tmp_storage)
	heap.SetHugeList(nil)
	memset(heap.GetFreeSlot(), 0, g.SizeOf("heap -> free_slot"))
	storage = _zendMmAlloc(heap, g.SizeOf("zend_mm_storage")+data_size)
	if storage == nil {
		handlers.GetChunkFree()(&tmp_storage, chunk, 2*1024*1024)
		r.Fprintf(stderr, "\nCan't initialize heap: [%d] %s\n", errno, strerror(errno))
		return nil
	}
	memcpy(storage, &tmp_storage, g.SizeOf("zend_mm_storage"))
	if data {
		storage.SetData(any((*byte)(storage + g.SizeOf("zend_mm_storage"))))
		memcpy(storage.GetData(), data, data_size)
	}
	heap.SetStorage(storage)
	return heap
}
func ZendOutOfMemory() {
	r.Fprintf(stderr, "Out of memory\n")
	exit(1)
}
func __zendMalloc(len_ int) any {
	var tmp any = Malloc(len_)
	if tmp || len_ == 0 {
		return tmp
	}
	ZendOutOfMemory()
}
func __zendCalloc(nmemb int, len_ int) any {
	var tmp any
	len_ = ZendSafeAddressGuarded(nmemb, len_, 0)
	tmp = __zendMalloc(len_)
	memset(tmp, 0, len_)
	return tmp
}
func __zendRealloc(p any, len_ int) any {
	p = realloc(p, len_)
	if p || len_ == 0 {
		return p
	}
	ZendOutOfMemory()
}
