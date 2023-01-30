// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/core"
	r "sik/runtime"
)

func ZEND_MM_ALIGNED_SIZE(size int) int {
	return size + core.ZEND_MM_ALIGNMENT - 1&ZEND_MM_ALIGNMENT_MASK
}
func ZEND_MM_ALIGNED_SIZE_EX(size int, alignment int) int {
	return size + (alignment-1) & ^(alignment-1)
}
func EfreeSize(ptr any, size __auto__)                { Efree(ptr) }
func EfreeSizeRel(ptr any, size __auto__)             { EfreeRel(ptr) }
func Emalloc(size int) any                            { return _emalloc(size) }
func EmallocLarge(size int) __auto__                  { return _emallocLarge(size) }
func EmallocHuge(size __auto__) __auto__              { return _emallocHuge(size) }
func SafeEmalloc(nmemb int, size int, offset int) any { return _safeEmalloc(nmemb, size, offset) }
func Efree(ptr any)                                   { _efree(ptr) }
func EfreeLarge(ptr __auto__) __auto__                { return _efreeLarge(ptr) }
func EfreeHuge(ptr __auto__) __auto__                 { return _efreeHuge(ptr) }
func Ecalloc(nmemb int, size int) any                 { return _ecalloc(nmemb, size) }
func Erealloc(ptr any, size int) any                  { return _erealloc(ptr, size) }
func Erealloc2(ptr any, size int, copy_size int) any  { return _erealloc2(ptr, size, copy_size) }
func SafeErealloc(ptr any, nmemb int, size int, offset int) any {
	return _safeErealloc(ptr, nmemb, size, offset)
}
func EreallocRecoverable(ptr any, size int) any { return _erealloc(ptr, size) }
func Erealloc2Recoverable(ptr any, size int, copy_size int) any {
	return _erealloc2(ptr, size, copy_size)
}
func Estrdup(s string) *byte                             { return _estrdup(s) }
func Estrndup(s *byte, length int) *byte                 { return _estrndup(s, length) }
func ZendMemBlockSize(ptr any) int                       { return _zendMemBlockSize(ptr) }
func EmallocRel(size int) any                            { return _emalloc(size) }
func SafeEmallocRel(nmemb int, size int, offset int) any { return _safeEmalloc(nmemb, size, offset) }
func EfreeRel(ptr any)                                   { _efree(ptr) }
func EcallocRel(nmemb int, size int) any                 { return _ecalloc(nmemb, size) }
func EreallocRel(ptr any, size int) any                  { return _erealloc(ptr, size) }
func Erealloc2Rel(ptr any, size int, copy_size int) any  { return _erealloc2(ptr, size, copy_size) }
func EreallocRecoverableRel(ptr any, size int) any       { return _erealloc(ptr, size) }
func Erealloc2RecoverableRel(ptr any, size int, copy_size int) any {
	return _erealloc2(ptr, size, copy_size)
}
func SafeEreallocRel(ptr any, nmemb int, size int, offset int) any {
	return _safeErealloc(ptr, nmemb, size, offset)
}
func EstrdupRel(s *byte) *byte              { return _estrdup(s) }
func EstrndupRel(s *byte, length int) *byte { return _estrndup(s, length) }
func ZendMemBlockSizeRel(ptr any) int       { return _zendMemBlockSize(ptr) }
func Pemalloc(size int, persistent int) any {
	if persistent != 0 {
		return __zendMalloc(size)
	} else {
		return Emalloc(size)
	}
}
func SafePemalloc(nmemb int, size int, offset int, persistent int) any {
	if persistent != 0 {
		return _safeMalloc(nmemb, size, offset)
	} else {
		return SafeEmalloc(nmemb, size, offset)
	}
}
func Pefree(ptr any, persistent int) {
	if persistent != 0 {
		return Free(ptr)
	} else {
		return Efree(ptr)
	}
}
func PefreeSize(ptr any, size __auto__, persistent __auto__) {
	if persistent {
		Free(ptr)
	} else {
		EfreeSize(ptr, size)
	}
}
func Pecalloc(nmemb int, size int, persistent uint8) any {
	if persistent != 0 {
		return __zendCalloc(nmemb, size)
	} else {
		return Ecalloc(nmemb, size)
	}
}
func Perealloc(ptr any, size int, persistent uint8) any {
	if persistent != 0 {
		return __zendRealloc(ptr, size)
	} else {
		return Erealloc(ptr, size)
	}
}
func Perealloc2(ptr any, size int, copy_size int, persistent int) any {
	if persistent != 0 {
		return __zendRealloc(ptr, size)
	} else {
		return Erealloc2(ptr, size, copy_size)
	}
}
func SafePerealloc(ptr any, nmemb int, size int, offset int, persistent int) any {
	if persistent != 0 {
		return _safeRealloc(ptr, nmemb, size, offset)
	} else {
		return SafeErealloc(ptr, nmemb, size, offset)
	}
}
func PereallocRecoverable(ptr any, size int, persistent __auto__) any {
	if persistent {
		return realloc(ptr, size)
	} else {
		return EreallocRecoverable(ptr, size)
	}
}
func Perealloc2Recoverable(ptr any, size int, persistent __auto__) any {
	if persistent {
		return realloc(ptr, size)
	} else {
		return Erealloc2Recoverable(ptr, size, copy_size)
	}
}
func Pestrdup(s *byte, persistent int) *byte {
	if persistent != 0 {
		return strdup(s)
	} else {
		return Estrdup(s)
	}
}
func Pestrndup(s *byte, length int, persistent int) *byte {
	if persistent != 0 {
		return ZendStrndup(s, length)
	} else {
		return Estrndup(s, length)
	}
}
func PemallocRel(size int, persistent __auto__) any {
	if persistent {
		return __zendMalloc(size)
	} else {
		return EmallocRel(size)
	}
}
func PefreeRel(ptr any, persistent __auto__) {
	if persistent {
		return Free(ptr)
	} else {
		return EfreeRel(ptr)
	}
}
func PecallocRel(nmemb int, size int, persistent __auto__) any {
	if persistent {
		return __zendCalloc(nmemb, size)
	} else {
		return EcallocRel(nmemb, size)
	}
}
func PereallocRel(ptr any, size int, persistent __auto__) any {
	if persistent {
		return __zendRealloc(ptr, size)
	} else {
		return EreallocRel(ptr, size)
	}
}
func Perealloc2Rel(ptr any, size int, copy_size int, persistent __auto__) any {
	if persistent {
		return __zendRealloc(ptr, size)
	} else {
		return Erealloc2Rel(ptr, size, copy_size)
	}
}
func PereallocRecoverableRel(ptr any, size int, persistent __auto__) any {
	if persistent {
		return realloc(ptr, size)
	} else {
		return EreallocRecoverableRel(ptr, size)
	}
}
func Perealloc2RecoverableRel(ptr any, size int, copy_size int, persistent __auto__) any {
	if persistent {
		return realloc(ptr, size)
	} else {
		return Erealloc2RecoverableRel(ptr, size, copy_size)
	}
}
func PestrdupRel(s *byte, persistent __auto__) *byte {
	if persistent {
		return strdup(s)
	} else {
		return EstrdupRel(s)
	}
}
func ALLOC_HASHTABLE(ht *HashTable) *HashTable {
	ht = (*HashTable)(Emalloc(b.SizeOf("HashTable")))
	return ht
}
func FREE_HASHTABLE(ht *HashTable) { EfreeSize(ht, b.SizeOf("HashTable")) }
func ALLOC_HASHTABLE_REL(ht *HashTable) *HashTable {
	ht = (*HashTable)(EmallocRel(b.SizeOf("HashTable")))
	return ht
}
func FREE_HASHTABLE_REL(ht any)                           { EfreeSizeRel(ht, b.SizeOf("HashTable")) }
func ZendMmAlloc(heap *ZendMmHeap, size int) any          { return _zendMmAlloc(heap, size) }
func ZendMmFree(heap *ZendMmHeap, p any)                  { _zendMmFree(heap, p) }
func ZendMmRealloc(heap *ZendMmHeap, p any, size int) any { return _zendMmRealloc(heap, p, size) }
func ZendMmRealloc2(heap *ZendMmHeap, p any, size int, copy_size int) any {
	return _zendMmRealloc2(heap, p, size, copy_size)
}
func ZendMmBlockSize(heap *ZendMmHeap, p any) int            { return _zendMmBlockSize(heap, p) }
func ZendMmAllocRel(heap *ZendMmHeap, size int) any          { return _zendMmAlloc(heap, size) }
func ZendMmFreeRel(heap *ZendMmHeap, p any)                  { _zendMmFree(heap, p) }
func ZendMmReallocRel(heap *ZendMmHeap, p any, size int) any { return _zendMmRealloc(heap, p, size) }
func ZendMmRealloc2Rel(heap *ZendMmHeap, p any, size int, copy_size int) any {
	return _zendMmRealloc2(heap, p, size, copy_size)
}
func ZendMmBlockSizeRel(heap *ZendMmHeap, p any) int { return _zendMmBlockSize(heap, p) }
func ZEND_MM_CHECK(condition __auto__, message string) {
	if !condition {
		ZendMmPanic(message)
	}
}
func ZEND_MM_ALIGNED_OFFSET(size __auto__, alignment int) int { return size_t(size)&alignment - 1 }
func ZEND_MM_ALIGNED_BASE(size __auto__, alignment int) int   { return size_t(size) & ^(alignment - 1) }
func ZEND_MM_SIZE_TO_NUM(size int, alignment int) int {
	return (size_t(size) + (alignment - 1)) / alignment
}
func ZEND_MM_LRUN_PAGES(info ZendMmPageInfo) int {
	return (info & ZEND_MM_LRUN_PAGES_MASK) >> ZEND_MM_LRUN_PAGES_OFFSET
}
func ZEND_MM_SRUN_BIN_NUM(info ZendMmPageInfo) int {
	return (info & ZEND_MM_SRUN_BIN_NUM_MASK) >> ZEND_MM_SRUN_BIN_NUM_OFFSET
}
func ZEND_MM_SRUN_FREE_COUNTER(info ZendMmPageInfo) int {
	return (info & ZEND_MM_SRUN_FREE_COUNTER_MASK) >> ZEND_MM_SRUN_FREE_COUNTER_OFFSET
}
func ZEND_MM_NRUN_OFFSET(info ZendMmPageInfo) int {
	return (info & ZEND_MM_NRUN_OFFSET_MASK) >> ZEND_MM_NRUN_OFFSET_OFFSET
}
func ZEND_MM_FRUN() __auto__ { return ZEND_MM_IS_FRUN }
func ZEND_MM_LRUN(count __auto__) int {
	return ZEND_MM_IS_LRUN | count<<ZEND_MM_LRUN_PAGES_OFFSET
}
func ZEND_MM_SRUN(bin_num __auto__) int {
	return ZEND_MM_IS_SRUN | bin_num<<ZEND_MM_SRUN_BIN_NUM_OFFSET
}
func ZEND_MM_SRUN_EX(bin_num uint32, count uint32) int {
	return ZEND_MM_IS_SRUN | bin_num<<ZEND_MM_SRUN_BIN_NUM_OFFSET | count<<ZEND_MM_SRUN_FREE_COUNTER_OFFSET
}
func ZEND_MM_NRUN(bin_num uint32, offset uint32) int {
	return ZEND_MM_IS_SRUN | ZEND_MM_IS_LRUN | bin_num<<ZEND_MM_SRUN_BIN_NUM_OFFSET | offset<<ZEND_MM_NRUN_OFFSET_OFFSET
}
func ZEND_MM_PAGE_ADDR(chunk *ZendMmChunk, page_num *ZendMmPage) any {
	return any((*ZendMmPage)(chunk) + page_num)
}
func ZendMmPanic(message string) {
	r.Fprintf(stderr, "%s\n", message)

	/* See http://support.microsoft.com/kb/190351 */

	exit(1)

	/* See http://support.microsoft.com/kb/190351 */
}
func ZendMmSafeError(heap *ZendMmHeap, format string, limit int, size int) {
	heap.SetOverflow(1)
	var __orig_bailout *JMP_BUF = __EG().GetBailout()
	var __bailout JMP_BUF
	__EG().SetBailout(&__bailout)
	if SETJMP(__bailout) == 0 {
		ZendErrorNoreturn(E_ERROR, format, limit, size)
	} else {
		__EG().SetBailout(__orig_bailout)
	}
	__EG().SetBailout(__orig_bailout)
	heap.SetOverflow(0)
	ZendBailout()
	exit(1)
}
func ZendMmMmapFixed(addr any, size int) any {
	var flags int = MAP_PRIVATE | MAP_ANON

	/* MAP_FIXED leads to discarding of the old mapping, so it can't be used. */

	var ptr any = mmap(addr, size, PROT_READ|PROT_WRITE, flags, ZEND_MM_FD, 0)
	if ptr == MAP_FAILED {
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
	ptr = mmap(nil, size, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANON, ZEND_MM_FD, 0)
	if ptr == MAP_FAILED {
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
func ZendMmBitsetNts(bitset ZendMmBitset) int {
	var n int
	if bitset == zend_mm_bitset-1 {
		return ZEND_MM_BITSET_LEN
	}
	n = 0
	if b.SizeOf("zend_mm_bitset") == 8 {
		if (bitset & 0xffffffff) == 0xffffffff {
			n += 32
			bitset = bitset >> uint64(32)
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
func ZendMmBitsetIsSet(bitset *ZendMmBitset, bit int) int { return ZEND_BIT_TEST(bitset, bit) }
func ZendMmBitsetSetBit(bitset *ZendMmBitset, bit int) {
	bitset[bit/ZEND_MM_BITSET_LEN] |= uint64(1) << (bit&ZEND_MM_BITSET_LEN - 1)
}
func ZendMmBitsetResetBit(bitset *ZendMmBitset, bit int) {
	bitset[bit/ZEND_MM_BITSET_LEN] &= ^(uint64(1) << (bit&ZEND_MM_BITSET_LEN - 1))
}
func ZendMmBitsetSetRange(bitset *ZendMmBitset, start int, len_ int) {
	if len_ == 1 {
		ZendMmBitsetSetBit(bitset, start)
	} else {
		var pos int = start / ZEND_MM_BITSET_LEN
		var end int = (start + len_ - 1) / ZEND_MM_BITSET_LEN
		var bit int = start&ZEND_MM_BITSET_LEN - 1
		var tmp ZendMmBitset
		if pos != end {

			/* set bits from "bit" to ZEND_MM_BITSET_LEN-1 */

			tmp = zend_mm_bitset - 1<<bit
			bitset[b.PostInc(&pos)] |= tmp
			for pos != end {

				/* set all bits */

				bitset[b.PostInc(&pos)] = zend_mm_bitset - 1

				/* set all bits */

			}
			end = start + len_ - 1&ZEND_MM_BITSET_LEN - 1

			/* set bits from "0" to "end" */

			tmp = zend_mm_bitset - 1>>ZEND_MM_BITSET_LEN - 1 - end
			bitset[pos] |= tmp
		} else {
			end = start + len_ - 1&ZEND_MM_BITSET_LEN - 1

			/* set bits from "bit" to "end" */

			tmp = zend_mm_bitset - 1<<bit
			tmp &= zend_mm_bitset - 1>>ZEND_MM_BITSET_LEN - 1 - end
			bitset[pos] |= tmp
		}
	}
}
func ZendMmBitsetResetRange(bitset *ZendMmBitset, start int, len_ int) {
	if len_ == 1 {
		ZendMmBitsetResetBit(bitset, start)
	} else {
		var pos int = start / ZEND_MM_BITSET_LEN
		var end int = (start + len_ - 1) / ZEND_MM_BITSET_LEN
		var bit int = start&ZEND_MM_BITSET_LEN - 1
		var tmp ZendMmBitset
		if pos != end {

			/* reset bits from "bit" to ZEND_MM_BITSET_LEN-1 */

			tmp = ^((uint64(1) << bit) - 1)
			bitset[b.PostInc(&pos)] &= ^tmp
			for pos != end {

				/* set all bits */

				bitset[b.PostInc(&pos)] = 0

				/* set all bits */

			}
			end = start + len_ - 1&ZEND_MM_BITSET_LEN - 1

			/* reset bits from "0" to "end" */

			tmp = zend_mm_bitset - 1>>ZEND_MM_BITSET_LEN - 1 - end
			bitset[pos] &= ^tmp
		} else {
			end = start + len_ - 1&ZEND_MM_BITSET_LEN - 1

			/* reset bits from "bit" to "end" */

			tmp = zend_mm_bitset - 1<<bit
			tmp &= zend_mm_bitset - 1>>ZEND_MM_BITSET_LEN - 1 - end
			bitset[pos] &= ^tmp
		}
	}
}
func ZendMmBitsetIsFreeRange(bitset *ZendMmBitset, start int, len_ int) int {
	if len_ == 1 {
		return !(ZendMmBitsetIsSet(bitset, start))
	} else {
		var pos int = start / ZEND_MM_BITSET_LEN
		var end int = (start + len_ - 1) / ZEND_MM_BITSET_LEN
		var bit int = start&ZEND_MM_BITSET_LEN - 1
		var tmp ZendMmBitset
		if pos != end {

			/* set bits from "bit" to ZEND_MM_BITSET_LEN-1 */

			tmp = zend_mm_bitset - 1<<bit
			if (bitset[b.PostInc(&pos)] & tmp) != 0 {
				return 0
			}
			for pos != end {

				/* set all bits */

				if bitset[b.PostInc(&pos)] != 0 {
					return 0
				}

				/* set all bits */

			}
			end = start + len_ - 1&ZEND_MM_BITSET_LEN - 1

			/* set bits from "0" to "end" */

			tmp = zend_mm_bitset - 1>>ZEND_MM_BITSET_LEN - 1 - end
			return (bitset[pos] & tmp) == 0
		} else {
			end = start + len_ - 1&ZEND_MM_BITSET_LEN - 1

			/* set bits from "bit" to "end" */

			tmp = zend_mm_bitset - 1<<bit
			tmp &= zend_mm_bitset - 1>>ZEND_MM_BITSET_LEN - 1 - end
			return (bitset[pos] & tmp) == 0
		}
	}
}
func ZendMmChunkAllocInt(size int, alignment int) any {
	var ptr any = ZendMmMmap(size)
	if ptr == nil {
		return nil
	} else if ZEND_MM_ALIGNED_OFFSET(ptr, alignment) == 0 {
		return ptr
	} else {
		var offset int

		/* chunk has to be aligned */

		ZendMmMunmap(ptr, size)
		ptr = ZendMmMmap(size + alignment - REAL_PAGE_SIZE)
		offset = ZEND_MM_ALIGNED_OFFSET(ptr, alignment)
		if offset != 0 {
			offset = alignment - offset
			ZendMmMunmap(ptr, offset)
			ptr = (*byte)(ptr + offset)
			alignment -= offset
		}
		if alignment > REAL_PAGE_SIZE {
			ZendMmMunmap((*byte)(ptr+size), alignment-REAL_PAGE_SIZE)
		}
		return ptr
	}
}
func ZendMmChunkAlloc(heap *ZendMmHeap, size int, alignment int) any {
	if heap.GetStorage() != nil {
		var ptr any = heap.GetStorage().GetHandlers().GetChunkAlloc()(heap.GetStorage(), size, alignment)
		ZEND_ASSERT((zend_uintptr_t((*byte)(ptr+(alignment-1)))&alignment - 1) == ZendUintptrT(ptr))
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

	chunk.SetFreePages(ZEND_MM_PAGES - ZEND_MM_FIRST_PAGE)
	chunk.SetFreeTail(ZEND_MM_FIRST_PAGE)

	/* the younger chunks have bigger number */

	chunk.SetNum(chunk.GetPrev().GetNum() + 1)

	/* mark first pages as allocated */

	chunk.GetFreeMap()[0] = (1 << ZEND_MM_FIRST_PAGE) - 1
	chunk.GetMap()[0] = ZEND_MM_LRUN(ZEND_MM_FIRST_PAGE)
}
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
			var best_len uint32 = ZEND_MM_PAGES
			var free_tail uint32 = chunk.GetFreeTail()
			var bitset *ZendMmBitset = chunk.GetFreeMap()
			var tmp ZendMmBitset = *(b.PostInc(&bitset))
			var i uint32 = 0
			for true {

				/* skip allocated blocks */

				for tmp == zend_mm_bitset-1 {
					i += ZEND_MM_BITSET_LEN
					if i == ZEND_MM_PAGES {
						if best > 0 {
							page_num = best
							goto found
						} else {
							goto not_found
						}
					}
					tmp = *(b.PostInc(&bitset))
				}

				/* find first 0 bit */

				page_num = i + ZendMmBitsetNts(tmp)

				/* reset bits from 0 to "bit" */

				tmp &= tmp + 1

				/* skip free blocks */

				for tmp == 0 {
					i += ZEND_MM_BITSET_LEN
					if i >= free_tail || i == ZEND_MM_PAGES {
						len_ = ZEND_MM_PAGES - page_num
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
					tmp = *(b.PostInc(&bitset))
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
				if ZEND_MM_CHUNK_SIZE > heap.GetLimit()-heap.GetRealSize() {
					if ZendMmGc(heap) != 0 {
						goto get_chunk
					} else if heap.GetOverflow() == 0 {
						ZendMmSafeError(heap, "Allowed memory size of %zu bytes exhausted (tried to allocate %zu bytes)", heap.GetLimit(), ZEND_MM_PAGE_SIZE*pages_count)
						return nil
					}
				}
				chunk = (*ZendMmChunk)(ZendMmChunkAlloc(heap, ZEND_MM_CHUNK_SIZE, ZEND_MM_CHUNK_SIZE))
				if chunk == nil {

					/* insufficient memory */

					if ZendMmGc(heap) != 0 && b.Assign(&chunk, (*ZendMmChunk)(ZendMmChunkAlloc(heap, ZEND_MM_CHUNK_SIZE, ZEND_MM_CHUNK_SIZE))) != nil {

					} else {
						ZendMmSafeError(heap, "Out of memory (allocated %zu) (tried to allocate %zu bytes)", heap.GetRealSize(), ZEND_MM_PAGE_SIZE*pages_count)
						return nil
					}

					/* insufficient memory */

				}
				var size int = heap.GetRealSize() + ZEND_MM_CHUNK_SIZE
				var peak int = MAX(heap.GetRealPeak(), size)
				heap.SetRealSize(size)
				heap.SetRealPeak(peak)
			}
			heap.GetChunksCount()++
			if heap.GetChunksCount() > heap.GetPeakChunksCount() {
				heap.SetPeakChunksCount(heap.GetChunksCount())
			}
			ZendMmChunkInit(heap, chunk)
			page_num = ZEND_MM_FIRST_PAGE
			len_ = ZEND_MM_PAGES - ZEND_MM_FIRST_PAGE
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
	chunk.GetMap()[page_num] = ZEND_MM_LRUN(pages_count)
	if page_num == chunk.GetFreeTail() {
		chunk.SetFreeTail(page_num + pages_count)
	}
	return ZEND_MM_PAGE_ADDR(chunk, page_num)
}
func ZendMmAllocLargeEx(heap *ZendMmHeap, size int) any {
	var pages_count int = int(ZEND_MM_SIZE_TO_NUM(size, ZEND_MM_PAGE_SIZE))
	var ptr any = ZendMmAllocPages(heap, pages_count)
	var size__1 int = heap.GetSize() + pages_count*ZEND_MM_PAGE_SIZE
	var peak int = MAX(heap.GetPeak(), size__1)
	heap.SetSize(size__1)
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
		heap.SetRealSize(heap.GetRealSize() - ZEND_MM_CHUNK_SIZE)
		if heap.GetCachedChunks() == nil {
			if heap.GetChunksCount() != heap.GetLastChunksDeleteBoundary() {
				heap.SetLastChunksDeleteBoundary(heap.GetChunksCount())
				heap.SetLastChunksDeleteCount(0)
			} else {
				heap.GetLastChunksDeleteCount()++
			}
		}
		if heap.GetCachedChunks() == nil || chunk.GetNum() > heap.GetCachedChunks().GetNum() {
			ZendMmChunkFree(heap, chunk, ZEND_MM_CHUNK_SIZE)
		} else {

			//TODO: select the best chunk to delete???

			chunk.SetNext(heap.GetCachedChunks().GetNext())
			ZendMmChunkFree(heap, heap.GetCachedChunks(), ZEND_MM_CHUNK_SIZE)
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
	if free_chunk != 0 && chunk != heap.GetMainChunk() && chunk.GetFreePages() == ZEND_MM_PAGES-ZEND_MM_FIRST_PAGE {
		ZendMmDeleteChunk(heap, chunk)
	}
}
func ZendMmFreePages(heap *ZendMmHeap, chunk *ZendMmChunk, page_num int, pages_count int) {
	ZendMmFreePagesEx(heap, chunk, page_num, pages_count, 1)
}
func ZendMmFreeLarge(heap *ZendMmHeap, chunk *ZendMmChunk, page_num int, pages_count int) {
	heap.SetSize(heap.GetSize() - pages_count*ZEND_MM_PAGE_SIZE)
	ZendMmFreePages(heap, chunk, page_num, pages_count)
}
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
func ZEND_MM_SMALL_SIZE_TO_BIN(size int) int { return ZendMmSmallSizeToBin(size) }
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
	chunk = (*ZendMmChunk)(ZEND_MM_ALIGNED_BASE(bin, ZEND_MM_CHUNK_SIZE))
	page_num = ZEND_MM_ALIGNED_OFFSET(bin, ZEND_MM_CHUNK_SIZE) / ZEND_MM_PAGE_SIZE
	chunk.GetMap()[page_num] = ZEND_MM_SRUN(bin_num)
	if BinPages[bin_num] > 1 {
		var i uint32 = 1
		for {
			chunk.GetMap()[page_num+i] = ZEND_MM_NRUN(bin_num, i)
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
	var peak int = MAX(heap.GetPeak(), size)
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
func ZendMmAllocHeap(heap *ZendMmHeap, size int) any {
	var ptr any
	if size <= ZEND_MM_MAX_SMALL_SIZE {
		ptr = ZendMmAllocSmall(heap, ZEND_MM_SMALL_SIZE_TO_BIN(size))
		return ptr
	} else if size <= ZEND_MM_MAX_LARGE_SIZE {
		ptr = ZendMmAllocLarge(heap, size)
		return ptr
	} else {
		return ZendMmAllocHuge(heap, size)
	}
}
func ZendMmFreeHeap(heap *ZendMmHeap, ptr any) {
	var page_offset int = ZEND_MM_ALIGNED_OFFSET(ptr, ZEND_MM_CHUNK_SIZE)
	if page_offset == 0 {
		if ptr != nil {
			ZendMmFreeHuge(heap, ptr)
		}
	} else {
		var chunk *ZendMmChunk = (*ZendMmChunk)(ZEND_MM_ALIGNED_BASE(ptr, ZEND_MM_CHUNK_SIZE))
		var page_num int = int(page_offset / ZEND_MM_PAGE_SIZE)
		var info ZendMmPageInfo = chunk.GetMap()[page_num]
		ZEND_MM_CHECK(chunk.GetHeap() == heap, "zend_mm_heap corrupted")
		if (info & ZEND_MM_IS_SRUN) != 0 {
			ZendMmFreeSmall(heap, ptr, ZEND_MM_SRUN_BIN_NUM(info))
		} else {
			var pages_count int = ZEND_MM_LRUN_PAGES(info)
			ZEND_MM_CHECK(ZEND_MM_ALIGNED_OFFSET(page_offset, ZEND_MM_PAGE_SIZE) == 0, "zend_mm_heap corrupted")
			ZendMmFreeLarge(heap, chunk, page_num, pages_count)
		}
	}
}
func ZendMmSize(heap *ZendMmHeap, ptr any) int {
	var page_offset int = ZEND_MM_ALIGNED_OFFSET(ptr, ZEND_MM_CHUNK_SIZE)
	if page_offset == 0 {
		return ZendMmGetHugeBlockSize(heap, ptr)
	} else {
		var chunk *ZendMmChunk
		var page_num int
		var info ZendMmPageInfo
		chunk = (*ZendMmChunk)(ZEND_MM_ALIGNED_BASE(ptr, ZEND_MM_CHUNK_SIZE))
		page_num = int(page_offset / ZEND_MM_PAGE_SIZE)
		info = chunk.GetMap()[page_num]
		ZEND_MM_CHECK(chunk.GetHeap() == heap, "zend_mm_heap corrupted")
		if (info & ZEND_MM_IS_SRUN) != 0 {
			return BinDataSize[ZEND_MM_SRUN_BIN_NUM(info)]
		} else {
			return ZEND_MM_LRUN_PAGES(info) * ZEND_MM_PAGE_SIZE
		}
	}
}
func ZendMmReallocSlow(heap *ZendMmHeap, ptr any, size int, copy_size int) any {
	var ret any
	var orig_peak int = heap.GetPeak()
	ret = ZendMmAllocHeap(heap, size)
	memcpy(ret, ptr, copy_size)
	ZendMmFreeHeap(heap, ptr)
	heap.SetPeak(MAX(orig_peak, heap.GetSize()))
	return ret
}
func ZendMmReallocHuge(heap *ZendMmHeap, ptr any, size int, copy_size int) any {
	var old_size int
	var new_size int
	old_size = ZendMmGetHugeBlockSize(heap, ptr)
	if size > ZEND_MM_MAX_LARGE_SIZE {
		new_size = ZEND_MM_ALIGNED_SIZE_EX(size, REAL_PAGE_SIZE)
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
				heap.SetRealPeak(MAX(heap.GetRealPeak(), heap.GetRealSize()))
				heap.SetSize(heap.GetSize() + new_size - old_size)
				heap.SetPeak(MAX(heap.GetPeak(), heap.GetSize()))
				ZendMmChangeHugeBlockSize(heap, ptr, new_size)
				return ptr
			}

			/* try to map tail right after this block */

		}
	}
	return ZendMmReallocSlow(heap, ptr, size, MIN(old_size, copy_size))
}
func ZendMmReallocHeap(heap *ZendMmHeap, ptr any, size int, use_copy_size ZendBool, copy_size int) any {
	var page_offset int
	var old_size int
	var new_size int
	var ret any
	page_offset = ZEND_MM_ALIGNED_OFFSET(ptr, ZEND_MM_CHUNK_SIZE)
	if page_offset == 0 {
		if ptr == nil {
			return _zendMmAlloc(heap, size)
		} else {
			return ZendMmReallocHuge(heap, ptr, size, copy_size)
		}
	} else {
		var chunk *ZendMmChunk = (*ZendMmChunk)(ZEND_MM_ALIGNED_BASE(ptr, ZEND_MM_CHUNK_SIZE))
		var page_num int = int(page_offset / ZEND_MM_PAGE_SIZE)
		var info ZendMmPageInfo = chunk.GetMap()[page_num]
		ZEND_MM_CHECK(chunk.GetHeap() == heap, "zend_mm_heap corrupted")
		if (info & ZEND_MM_IS_SRUN) != 0 {
			var old_bin_num int = ZEND_MM_SRUN_BIN_NUM(info)
			for {
				old_size = BinDataSize[old_bin_num]

				/* Check if requested size fits into current bin */

				if size <= old_size {

					/* Check if truncation is necessary */

					if old_bin_num > 0 && size < BinDataSize[old_bin_num-1] {

						/* truncation */

						ret = ZendMmAllocSmall(heap, ZEND_MM_SMALL_SIZE_TO_BIN(size))
						if use_copy_size != 0 {
							copy_size = MIN(size, copy_size)
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

				} else if size <= ZEND_MM_MAX_SMALL_SIZE {

					/* small extension */

					var orig_peak int = heap.GetPeak()
					ret = ZendMmAllocSmall(heap, ZEND_MM_SMALL_SIZE_TO_BIN(size))
					if use_copy_size != 0 {
						copy_size = MIN(old_size, copy_size)
					} else {
						copy_size = old_size
					}
					memcpy(ret, ptr, copy_size)
					ZendMmFreeSmall(heap, ptr, old_bin_num)
					heap.SetPeak(MAX(orig_peak, heap.GetSize()))

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
			ZEND_MM_CHECK(ZEND_MM_ALIGNED_OFFSET(page_offset, ZEND_MM_PAGE_SIZE) == 0, "zend_mm_heap corrupted")
			old_size = ZEND_MM_LRUN_PAGES(info) * ZEND_MM_PAGE_SIZE
			if size > ZEND_MM_MAX_SMALL_SIZE && size <= ZEND_MM_MAX_LARGE_SIZE {
				new_size = ZEND_MM_ALIGNED_SIZE_EX(size, ZEND_MM_PAGE_SIZE)
				if new_size == old_size {
					return ptr
				} else if new_size < old_size {

					/* free tail pages */

					var new_pages_count int = int(new_size / ZEND_MM_PAGE_SIZE)
					var rest_pages_count int = int((old_size - new_size) / ZEND_MM_PAGE_SIZE)
					heap.SetSize(heap.GetSize() - rest_pages_count*ZEND_MM_PAGE_SIZE)
					chunk.GetMap()[page_num] = ZEND_MM_LRUN(new_pages_count)
					chunk.SetFreePages(chunk.GetFreePages() + rest_pages_count)
					ZendMmBitsetResetRange(chunk.GetFreeMap(), page_num+new_pages_count, rest_pages_count)
					return ptr
				} else {
					var new_pages_count int = int(new_size / ZEND_MM_PAGE_SIZE)
					var old_pages_count int = int(old_size / ZEND_MM_PAGE_SIZE)

					/* try to allocate tail pages after this block */

					if page_num+new_pages_count <= ZEND_MM_PAGES && ZendMmBitsetIsFreeRange(chunk.GetFreeMap(), page_num+old_pages_count, new_pages_count-old_pages_count) != 0 {
						var size int = heap.GetSize() + (new_size - old_size)
						var peak int = MAX(heap.GetPeak(), size)
						heap.SetSize(size)
						heap.SetPeak(peak)
						chunk.SetFreePages(chunk.GetFreePages() - new_pages_count - old_pages_count)
						ZendMmBitsetSetRange(chunk.GetFreeMap(), page_num+old_pages_count, new_pages_count-old_pages_count)
						chunk.GetMap()[page_num] = ZEND_MM_LRUN(new_pages_count)
						return ptr
					}

					/* try to allocate tail pages after this block */

				}
			}
		}
	}
	copy_size = MIN(old_size, copy_size)
	return ZendMmReallocSlow(heap, ptr, size, copy_size)
}
func ZendMmAddHugeBlock(heap *ZendMmHeap, ptr any, size int) {
	var list *ZendMmHugeList = (*ZendMmHugeList)(ZendMmAllocHeap(heap, b.SizeOf("zend_mm_huge_list")))
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
	ZEND_MM_CHECK(0, "zend_mm_heap corrupted")
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
	ZEND_MM_CHECK(0, "zend_mm_heap corrupted")
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
	var alignment int = REAL_PAGE_SIZE
	var new_size int = ZEND_MM_ALIGNED_SIZE_EX(size, alignment)
	var ptr any
	if new_size < size {
		ZendErrorNoreturn(E_ERROR, "Possible integer overflow in memory allocation (%zu + %zu)", size, alignment)
	}
	if new_size > heap.GetLimit()-heap.GetRealSize() {
		if ZendMmGc(heap) != 0 && new_size <= heap.GetLimit()-heap.GetRealSize() {

		} else if heap.GetOverflow() == 0 {
			ZendMmSafeError(heap, "Allowed memory size of %zu bytes exhausted (tried to allocate %zu bytes)", heap.GetLimit(), size)
			return nil
		}
	}
	ptr = ZendMmChunkAlloc(heap, new_size, ZEND_MM_CHUNK_SIZE)
	if ptr == nil {

		/* insufficient memory */

		if ZendMmGc(heap) != 0 && b.Assign(&ptr, ZendMmChunkAlloc(heap, new_size, ZEND_MM_CHUNK_SIZE)) != nil {

		} else {
			ZendMmSafeError(heap, "Out of memory (allocated %zu) (tried to allocate %zu bytes)", heap.GetRealSize(), size)
			return nil
		}

		/* insufficient memory */

	}
	ZendMmAddHugeBlock(heap, ptr, new_size)
	var size__1 int = heap.GetRealSize() + new_size
	var peak int = MAX(heap.GetRealPeak(), size__1)
	heap.SetRealSize(size__1)
	heap.SetRealPeak(peak)
	var size__2 int = heap.GetSize() + new_size
	var peak__1 int = MAX(heap.GetPeak(), size__2)
	heap.SetSize(size__2)
	heap.SetPeak(peak__1)
	return ptr
}
func ZendMmFreeHuge(heap *ZendMmHeap, ptr any) {
	var size int
	ZEND_MM_CHECK(ZEND_MM_ALIGNED_OFFSET(ptr, ZEND_MM_CHUNK_SIZE) == 0, "zend_mm_heap corrupted")
	size = ZendMmDelHugeBlock(heap, ptr)
	ZendMmChunkFree(heap, ptr, size)
	heap.SetRealSize(heap.GetRealSize() - size)
	heap.SetSize(heap.GetSize() - size)
}
func ZendMmInit() *ZendMmHeap {
	var chunk *ZendMmChunk = (*ZendMmChunk)(ZendMmChunkAllocInt(ZEND_MM_CHUNK_SIZE, ZEND_MM_CHUNK_SIZE))
	var heap *ZendMmHeap
	if chunk == nil {
		r.Fprintf(stderr, "\nCan't initialize heap: [%d] %s\n", errno, strerror(errno))
		return nil
	}
	heap = chunk.GetHeapSlot()
	chunk.SetHeap(heap)
	chunk.SetNext(chunk)
	chunk.SetPrev(chunk)
	chunk.SetFreePages(ZEND_MM_PAGES - ZEND_MM_FIRST_PAGE)
	chunk.SetFreeTail(ZEND_MM_FIRST_PAGE)
	chunk.SetNum(0)
	chunk.GetFreeMap()[0] = (int64(1) << ZEND_MM_FIRST_PAGE) - 1
	chunk.GetMap()[0] = ZEND_MM_LRUN(ZEND_MM_FIRST_PAGE)
	heap.SetMainChunk(chunk)
	heap.SetCachedChunks(nil)
	heap.SetChunksCount(1)
	heap.SetPeakChunksCount(1)
	heap.SetCachedChunksCount(0)
	heap.SetAvgChunksCount(1.0)
	heap.SetLastChunksDeleteBoundary(0)
	heap.SetLastChunksDeleteCount(0)
	heap.SetRealSize(ZEND_MM_CHUNK_SIZE)
	heap.SetRealPeak(ZEND_MM_CHUNK_SIZE)
	heap.SetSize(0)
	heap.SetPeak(0)
	heap.SetLimit(int(int64(-1) >> int(int64(1))))
	heap.SetOverflow(0)
	heap.SetUseCustomHeap(ZEND_MM_CUSTOM_HEAP_NONE)
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
	for i = 0; i < ZEND_MM_BINS; i++ {
		has_free_pages = 0
		p = heap.GetFreeSlot()[i]
		for p != nil {
			chunk = (*ZendMmChunk)(ZEND_MM_ALIGNED_BASE(p, ZEND_MM_CHUNK_SIZE))
			ZEND_MM_CHECK(chunk.GetHeap() == heap, "zend_mm_heap corrupted")
			page_offset = ZEND_MM_ALIGNED_OFFSET(p, ZEND_MM_CHUNK_SIZE)
			ZEND_ASSERT(page_offset != 0)
			page_num = int(page_offset / ZEND_MM_PAGE_SIZE)
			info = chunk.GetMap()[page_num]
			ZEND_ASSERT((info & ZEND_MM_IS_SRUN) != 0)
			if (info & ZEND_MM_IS_LRUN) != 0 {
				page_num -= ZEND_MM_NRUN_OFFSET(info)
				info = chunk.GetMap()[page_num]
				ZEND_ASSERT((info & ZEND_MM_IS_SRUN) != 0)
				ZEND_ASSERT((info & ZEND_MM_IS_LRUN) == 0)
			}
			ZEND_ASSERT(ZEND_MM_SRUN_BIN_NUM(info) == i)
			free_counter = ZEND_MM_SRUN_FREE_COUNTER(info) + 1
			if free_counter == BinElements[i] {
				has_free_pages = 1
			}
			chunk.GetMap()[page_num] = ZEND_MM_SRUN_EX(i, free_counter)
			p = p.GetNextFreeSlot()
		}
		if has_free_pages == 0 {
			continue
		}
		q = heap.GetFreeSlot()[i]
		p = *q
		for p != nil {
			chunk = (*ZendMmChunk)(ZEND_MM_ALIGNED_BASE(p, ZEND_MM_CHUNK_SIZE))
			ZEND_MM_CHECK(chunk.GetHeap() == heap, "zend_mm_heap corrupted")
			page_offset = ZEND_MM_ALIGNED_OFFSET(p, ZEND_MM_CHUNK_SIZE)
			ZEND_ASSERT(page_offset != 0)
			page_num = int(page_offset / ZEND_MM_PAGE_SIZE)
			info = chunk.GetMap()[page_num]
			ZEND_ASSERT((info & ZEND_MM_IS_SRUN) != 0)
			if (info & ZEND_MM_IS_LRUN) != 0 {
				page_num -= ZEND_MM_NRUN_OFFSET(info)
				info = chunk.GetMap()[page_num]
				ZEND_ASSERT((info & ZEND_MM_IS_SRUN) != 0)
				ZEND_ASSERT((info & ZEND_MM_IS_LRUN) == 0)
			}
			ZEND_ASSERT(ZEND_MM_SRUN_BIN_NUM(info) == i)
			if ZEND_MM_SRUN_FREE_COUNTER(info) == BinElements[i] {

				/* remove from cache */

				p = p.GetNextFreeSlot()
				*q = p
			} else {
				q = p.GetNextFreeSlot()
				p = *q
			}
		}
	}
	chunk = heap.GetMainChunk()
	for {
		i = ZEND_MM_FIRST_PAGE
		for i < chunk.GetFreeTail() {
			if ZendMmBitsetIsSet(chunk.GetFreeMap(), i) != 0 {
				info = chunk.GetMap()[i]
				if (info & ZEND_MM_IS_SRUN) != 0 {
					var bin_num int = ZEND_MM_SRUN_BIN_NUM(info)
					var pages_count int = BinPages[bin_num]
					if ZEND_MM_SRUN_FREE_COUNTER(info) == BinElements[bin_num] {

						/* all elements are free */

						ZendMmFreePagesEx(heap, chunk, i, pages_count, 0)
						collected += pages_count
					} else {

						/* reset counter */

						chunk.GetMap()[i] = ZEND_MM_SRUN(bin_num)

						/* reset counter */

					}
					i += BinPages[bin_num]
				} else {
					i += ZEND_MM_LRUN_PAGES(info)
				}
			} else {
				i++
			}
		}
		if chunk.GetFreePages() == ZEND_MM_PAGES-ZEND_MM_FIRST_PAGE {
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
	return collected * ZEND_MM_PAGE_SIZE
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
			ZendMmChunkFree(heap, p, ZEND_MM_CHUNK_SIZE)
		}

		/* free the first chunk */

		ZendMmChunkFree(heap, heap.GetMainChunk(), ZEND_MM_CHUNK_SIZE)

		/* free the first chunk */

	} else {

		/* free some cached chunks to keep average count */

		heap.SetAvgChunksCount((heap.GetAvgChunksCount() + float64(heap.GetPeakChunksCount())) / 2.0)
		for float64(heap.GetCachedChunksCount()+0.9 > heap.GetAvgChunksCount() && heap.GetCachedChunks() != nil) {
			p = heap.GetCachedChunks()
			heap.SetCachedChunks(p.GetNext())
			ZendMmChunkFree(heap, p, ZEND_MM_CHUNK_SIZE)
			heap.GetCachedChunksCount()--
		}

		/* clear cached chunks */

		p = heap.GetCachedChunks()
		for p != nil {
			var q *ZendMmChunk = p.GetNext()
			memset(p, 0, b.SizeOf("zend_mm_chunk"))
			p.SetNext(q)
			p = q
		}

		/* reinitialize the first chunk and heap */

		p = heap.GetMainChunk()
		p.SetHeap(p.GetHeapSlot())
		p.SetNext(p)
		p.SetPrev(p)
		p.SetFreePages(ZEND_MM_PAGES - ZEND_MM_FIRST_PAGE)
		p.SetFreeTail(ZEND_MM_FIRST_PAGE)
		p.SetNum(0)
		heap.SetPeak(0)
		heap.SetSize(heap.GetPeak())
		memset(heap.GetFreeSlot(), 0, b.SizeOf("heap -> free_slot"))
		heap.SetRealSize(ZEND_MM_CHUNK_SIZE)
		heap.SetRealPeak(ZEND_MM_CHUNK_SIZE)
		heap.SetChunksCount(1)
		heap.SetPeakChunksCount(1)
		heap.SetLastChunksDeleteBoundary(0)
		heap.SetLastChunksDeleteCount(0)
		memset(p.GetFreeMap(), 0, b.SizeOf("p -> free_map")+b.SizeOf("p -> map"))
		p.GetFreeMap()[0] = (1 << ZEND_MM_FIRST_PAGE) - 1
		p.GetMap()[0] = ZEND_MM_LRUN(ZEND_MM_FIRST_PAGE)
	}
}
func _zendMmAlloc(heap *ZendMmHeap, size int) any { return ZendMmAllocHeap(heap, size) }
func _zendMmFree(heap *ZendMmHeap, ptr any)       { ZendMmFreeHeap(heap, ptr) }
func _zendMmRealloc(heap *ZendMmHeap, ptr any, size int) any {
	return ZendMmReallocHeap(heap, ptr, size, 0, size)
}
func _zendMmRealloc2(heap *ZendMmHeap, ptr any, size int, copy_size int) any {
	return ZendMmReallocHeap(heap, ptr, size, 1, copy_size)
}
func _zendMmBlockSize(heap *ZendMmHeap, ptr any) int { return ZendMmSize(heap, ptr) }
func AG(v __auto__) __auto__                         { return AllocGlobals.v }
func IsZendMm() int                                  { return !(AG(mm_heap).use_custom_heap) }
func IsZendPtr(ptr any) int {
	if AG(mm_heap).use_custom_heap {
		return 0
	}
	if AG(mm_heap).main_chunk {
		var chunk *ZendMmChunk = AG(mm_heap).main_chunk
		for {
			if ptr >= any(chunk != nil && ptr < any((*byte)(chunk+ZEND_MM_CHUNK_SIZE))) {
				return 1
			}
			chunk = chunk.GetNext()
			if chunk == AG(mm_heap).main_chunk {
				break
			}
		}
	}
	if AG(mm_heap).huge_list {
		var block *ZendMmHugeList = AG(mm_heap).huge_list
		for {
			if ptr >= any(block != nil && ptr < any((*byte)(block+block.GetSize()))) {
				return 1
			}
			block = block.GetNext()
			if block == AG(mm_heap).huge_list {
				break
			}
		}
	}
	return 0
}
func _mallocCustom(size int) any {
	return AG(mm_heap).custom_heap.std._malloc(size)
}
func _efreeCustom(ptr any) { AG(mm_heap).custom_heap.std._free(ptr) }
func _reallocCustom(ptr any, size int) any {
	return AG(mm_heap).custom_heap.std._realloc(ptr, size)
}
func _emalloc(size int) any {
	if AG(mm_heap).use_custom_heap {
		return _mallocCustom(size)
	}
	return ZendMmAllocHeap(AG(mm_heap), size)
}
func _efree(ptr any) {
	if AG(mm_heap).use_custom_heap {
		_efreeCustom(ptr)
		return
	}
	ZendMmFreeHeap(AG(mm_heap), ptr)
}
func _erealloc(ptr any, size int) any {
	if AG(mm_heap).use_custom_heap {
		return _reallocCustom(ptr, size)
	}
	return ZendMmReallocHeap(AG(mm_heap), ptr, size, 0, size)
}
func _erealloc2(ptr any, size int, copy_size int) any {
	if AG(mm_heap).use_custom_heap {
		return _reallocCustom(ptr, size)
	}
	return ZendMmReallocHeap(AG(mm_heap), ptr, size, 1, copy_size)
}
func _zendMemBlockSize(ptr any) int {
	if AG(mm_heap).use_custom_heap {
		return 0
	}
	return ZendMmSize(AG(mm_heap), ptr)
}
func _safeEmalloc(nmemb int, size int, offset int) any {
	return _emalloc(ZendSafeAddressGuarded(nmemb, size, offset))
}
func _safeMalloc(nmemb int, size int, offset int) any {
	return Pemalloc(ZendSafeAddressGuarded(nmemb, size, offset), 1)
}
func _safeErealloc(ptr any, nmemb int, size int, offset int) any {
	return _erealloc(ptr, ZendSafeAddressGuarded(nmemb, size, offset))
}
func _safeRealloc(ptr any, nmemb int, size int, offset int) any {
	return Perealloc(ptr, ZendSafeAddressGuarded(nmemb, size, offset), 1)
}
func _ecalloc(nmemb int, size int) any {
	var p any
	size = ZendSafeAddressGuarded(nmemb, size, 0)
	p = _emalloc(size)
	memset(p, 0, size)
	return p
}
func _estrdup(s *byte) *byte {
	var length int
	var p *byte
	length = strlen(s)
	if length+1 == 0 {
		ZendErrorNoreturn(E_ERROR, "Possible integer overflow in memory allocation (1 * %zu + 1)", length)
	}
	p = (*byte)(_emalloc(length + 1))
	memcpy(p, s, length+1)
	return p
}
func _estrndup(s *byte, length int) *byte {
	var p *byte
	if length+1 == 0 {
		ZendErrorNoreturn(E_ERROR, "Possible integer overflow in memory allocation (1 * %zu + 1)", length)
	}
	p = (*byte)(_emalloc(length + 1))
	memcpy(p, s, length)
	p[length] = 0
	return p
}
func ZendStrndup(s *byte, length int) *byte {
	var p *byte
	if length+1 == 0 {
		ZendErrorNoreturn(E_ERROR, "Possible integer overflow in memory allocation (1 * %zu + 1)", length)
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
	if memory_limit < ZEND_MM_CHUNK_SIZE {
		memory_limit = ZEND_MM_CHUNK_SIZE
	}
	if memory_limit < AG(mm_heap).real_size {
		return FAILURE
	}
	AG(mm_heap).limit = memory_limit
	return SUCCESS
}
func ZendMemoryUsage(real_usage int) int {
	if real_usage != 0 {
		return AG(mm_heap).real_size
	} else {
		var usage int = AG(mm_heap).size
		return usage
	}
	return 0
}
func ZendMemoryPeakUsage(real_usage int) int {
	if real_usage != 0 {
		return AG(mm_heap).real_peak
	} else {
		return AG(mm_heap).peak
	}
	return 0
}
func ShutdownMemoryManager(silent int, full_shutdown int) {
	ZendMmShutdown(AG(mm_heap), full_shutdown, silent)
}
func TrackedMalloc(size int) any {
	var ptr any = __zendMalloc(size)
	var h ZendUlong = uintPtr(ptr) >> core.ZEND_MM_ALIGNMENT_LOG2
	ZEND_ASSERT(any(uintPtr(h<<core.ZEND_MM_ALIGNMENT_LOG2) == ptr))
	ZendHashIndexAddEmptyElement(AG(mm_heap).tracked_allocs, h)
	return ptr
}
func TrackedFree(ptr any) {
	var h ZendUlong = uintPtr(ptr) >> core.ZEND_MM_ALIGNMENT_LOG2
	ZendHashIndexDel(AG(mm_heap).tracked_allocs, h)
	Free(ptr)
}
func TrackedRealloc(ptr any, new_size int) any {
	var h ZendUlong = uintPtr(ptr) >> core.ZEND_MM_ALIGNMENT_LOG2
	ZendHashIndexDel(AG(mm_heap).tracked_allocs, h)
	ptr = __zendRealloc(ptr, new_size)
	h = uintPtr(ptr) >> core.ZEND_MM_ALIGNMENT_LOG2
	ZEND_ASSERT(any(uintPtr(h<<core.ZEND_MM_ALIGNMENT_LOG2) == ptr))
	ZendHashIndexAddEmptyElement(AG(mm_heap).tracked_allocs, h)
	return ptr
}
func TrackedFreeAll() {
	var tracked_allocs *HashTable = AG(mm_heap).tracked_allocs
	var h ZendUlong
	var __ht *HashTable = tracked_allocs
	for _, _p := range __ht.foreachData() {
		var _z *Zval = _p.GetVal()

		h = _p.GetH()
		var ptr any = any(uintPtr(h << core.ZEND_MM_ALIGNMENT_LOG2))
		Free(ptr)
	}
}
func AllocGlobalsCtor(alloc_globals *ZendAllocGlobals) {
	var tmp *byte
	tmp = getenv("USE_ZEND_ALLOC")
	if tmp != nil && ZendAtoi(tmp, 0) == 0 {
		var tracked ZendBool = b.Assign(&tmp, getenv("USE_TRACKED_ALLOC")) && ZendAtoi(tmp, 0) != 0
		var mm_heap *ZendMmHeap = b.Assign(&(alloc_globals.GetMmHeap()), Malloc(b.SizeOf("zend_mm_heap")))
		memset(mm_heap, 0, b.SizeOf("zend_mm_heap"))
		mm_heap.SetUseCustomHeap(ZEND_MM_CUSTOM_HEAP_STD)
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
			mm_heap.SetTrackedAllocs(Malloc(b.SizeOf("HashTable")))
			ZendHashInit(mm_heap.GetTrackedAllocs(), 1024, nil, nil, 1)
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
	old_heap = AG(mm_heap)
	AG(mm_heap) = (*ZendMmHeap)(new_heap)
	return (*ZendMmHeap)(old_heap)
}
func ZendMmGetHeap() *ZendMmHeap                  { return AG(mm_heap) }
func ZendMmIsCustomHeap(new_heap *ZendMmHeap) int { return AG(mm_heap).use_custom_heap }
func ZendMmSetCustomHandlers(heap *ZendMmHeap, _malloc func(int) any, _free func(any), _realloc func(any, int) any) {
	var _heap *ZendMmHeap = (*ZendMmHeap)(heap)
	if _malloc == nil && _free == nil && _realloc == nil {
		_heap.SetUseCustomHeap(ZEND_MM_CUSTOM_HEAP_NONE)
	} else {
		_heap.SetUseCustomHeap(ZEND_MM_CUSTOM_HEAP_STD)
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
	memcpy((*ZendMmHandlers)(tmp_storage.GetHandlers()), handlers, b.SizeOf("zend_mm_handlers"))
	tmp_storage.SetData(data)
	chunk = (*ZendMmChunk)(handlers.GetChunkAlloc()(&tmp_storage, ZEND_MM_CHUNK_SIZE, ZEND_MM_CHUNK_SIZE))
	if chunk == nil {
		r.Fprintf(stderr, "\nCan't initialize heap: [%d] %s\n", errno, strerror(errno))
		return nil
	}
	heap = chunk.GetHeapSlot()
	chunk.SetHeap(heap)
	chunk.SetNext(chunk)
	chunk.SetPrev(chunk)
	chunk.SetFreePages(ZEND_MM_PAGES - ZEND_MM_FIRST_PAGE)
	chunk.SetFreeTail(ZEND_MM_FIRST_PAGE)
	chunk.SetNum(0)
	chunk.GetFreeMap()[0] = (int64(1) << ZEND_MM_FIRST_PAGE) - 1
	chunk.GetMap()[0] = ZEND_MM_LRUN(ZEND_MM_FIRST_PAGE)
	heap.SetMainChunk(chunk)
	heap.SetCachedChunks(nil)
	heap.SetChunksCount(1)
	heap.SetPeakChunksCount(1)
	heap.SetCachedChunksCount(0)
	heap.SetAvgChunksCount(1.0)
	heap.SetLastChunksDeleteBoundary(0)
	heap.SetLastChunksDeleteCount(0)
	heap.SetRealSize(ZEND_MM_CHUNK_SIZE)
	heap.SetRealPeak(ZEND_MM_CHUNK_SIZE)
	heap.SetSize(0)
	heap.SetPeak(0)
	heap.SetLimit(int64(-1) >> int64(1))
	heap.SetOverflow(0)
	heap.SetUseCustomHeap(0)
	heap.SetStorage(&tmp_storage)
	heap.SetHugeList(nil)
	memset(heap.GetFreeSlot(), 0, b.SizeOf("heap -> free_slot"))
	storage = _zendMmAlloc(heap, b.SizeOf("zend_mm_storage")+data_size)
	if storage == nil {
		handlers.GetChunkFree()(&tmp_storage, chunk, ZEND_MM_CHUNK_SIZE)
		r.Fprintf(stderr, "\nCan't initialize heap: [%d] %s\n", errno, strerror(errno))
		return nil
	}
	memcpy(storage, &tmp_storage, b.SizeOf("zend_mm_storage"))
	if data {
		storage.SetData(any((*byte)(storage + b.SizeOf("zend_mm_storage"))))
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
