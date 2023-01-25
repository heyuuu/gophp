// <<generate>>

package zend

import (
	b "sik/builtin"
)

func GC_REMOVE_FROM_BUFFER(p __auto__) {
	var _p *ZendRefcounted = (*ZendRefcounted)(p)
	if (GC_TYPE_INFO(_p) & GC_INFO_MASK) != 0 {
		GcRemoveFromBuffer(_p)
	}
}
func GC_MAY_LEAK(ref *ZendRefcounted) bool {
	return (GC_TYPE_INFO(ref) & (GC_INFO_MASK | GC_COLLECTABLE<<GC_FLAGS_SHIFT)) == GC_COLLECTABLE<<GC_FLAGS_SHIFT
}
func GcCheckPossibleRoot(ref *ZendRefcounted) {
	if EXPECTED(GC_TYPE_INFO(ref) == IS_REFERENCE) {
		var zv *Zval = &((*ZendReference)(ref)).val
		if !(Z_COLLECTABLE_P(zv)) {
			return
		}
		ref = Z_COUNTED_P(zv)
	}
	if UNEXPECTED(GC_MAY_LEAK(ref)) {
		GcPossibleRoot(ref)
	}
}
func GC_REF_ADDRESS(ref __auto__) int {
	return (GC_TYPE_INFO(ref) & GC_ADDRESS << GC_INFO_SHIFT) >> GC_INFO_SHIFT
}
func GC_REF_COLOR(ref *ZendRefcounted) int {
	return (GC_TYPE_INFO(ref) & GC_COLOR << GC_INFO_SHIFT) >> GC_INFO_SHIFT
}
func GC_REF_CHECK_COLOR(ref __auto__, color __auto__) bool {
	return (GC_TYPE_INFO(ref) & GC_COLOR << GC_INFO_SHIFT) == color<<GC_INFO_SHIFT
}
func GC_REF_SET_INFO(ref *ZendRefcounted, info int) {
	GC_TYPE_INFO(ref) = GC_TYPE_INFO(ref)&(GC_TYPE_MASK|GC_FLAGS_MASK) | info<<GC_INFO_SHIFT
}
func GC_REF_SET_COLOR(ref __auto__, c __auto__) {
	GC_TYPE_INFO(ref) = GC_TYPE_INFO(ref) & ^(GC_COLOR<<GC_INFO_SHIFT) | c<<GC_INFO_SHIFT
}
func GC_REF_SET_BLACK(ref __auto__) {
	GC_TYPE_INFO(ref) &= ^(GC_COLOR << GC_INFO_SHIFT)
}
func GC_REF_SET_PURPLE(ref __auto__) {
	GC_TYPE_INFO(ref) |= GC_COLOR << GC_INFO_SHIFT
}
func GC_GET_PTR(ptr *ZendRefcounted) any { return any(uintptr_t(ptr) & ^GC_BITS) }
func GC_IS_ROOT(ptr *ZendRefcounted) bool {
	return (uintptr_t(ptr) & GC_BITS) == GC_ROOT
}
func GC_IS_UNUSED(ptr *ZendRefcounted) bool {
	return (uintptr_t(ptr) & GC_BITS) == GC_UNUSED
}
func GC_IS_GARBAGE(ptr *ZendRefcounted) bool {
	return (uintptr_t(ptr) & GC_BITS) == GC_GARBAGE
}
func GC_IS_DTOR_GARBAGE(ptr *ZendRefcounted) bool {
	return (uintptr_t(ptr) & GC_BITS) == GC_DTOR_GARBAGE
}
func GC_MAKE_GARBAGE(ptr __auto__) any { return any(uintptr_t(ptr) | GC_GARBAGE) }
func GC_MAKE_DTOR_GARBAGE(ptr *ZendObject) any {
	return any(uintptr_t(ptr) | GC_DTOR_GARBAGE)
}
func GC_IDX2PTR(idx __auto__) __auto__ { return GC_G(buf) + idx }
func GC_PTR2IDX(ptr *GcRootBuffer) int { return ptr - GC_G(buf) }
func GC_IDX2LIST(idx uint32) any {
	return any(uintptr_t(idx*b.SizeOf("void *") | GC_UNUSED))
}
func GC_LIST2IDX(list *ZendRefcounted) int {
	return uint32_t(uintptr_t)(list) / b.SizeOf("void *")
}
func GC_HAS_UNUSED() bool               { return GC_G(unused) != GC_INVALID }
func GC_FETCH_UNUSED() uint32           { return GcFetchUnused() }
func GC_LINK_UNUSED(root *GcRootBuffer) { GcLinkUnused(root) }
func GC_HAS_NEXT_UNUSED_UNDER_THRESHOLD() bool {
	return GC_G(first_unused) < GC_G(gc_threshold)
}
func GC_HAS_NEXT_UNUSED() bool {
	return GC_G(first_unused) != GC_G(buf_size)
}
func GC_FETCH_NEXT_UNUSED() uint32 { return GcFetchNextUnused() }
func GC_STACK_DCL(init *GcStack) {
	var _stack *GcStack = init
	var _top int = 0
}
func GC_STACK_PUSH(ref *ZendRefcounted) { GcStackPush(&_stack, &_top, ref) }
func GC_STACK_POP() *ZendRefcounted     { return GcStackPop(&_stack, &_top) }
func GcStackNext(stack *GcStack) *GcStack {
	if UNEXPECTED(stack.GetNext() == nil) {
		var segment *GcStack = Emalloc(b.SizeOf("gc_stack"))
		segment.SetPrev(stack)
		segment.SetNext(nil)
		stack.SetNext(segment)
	}
	return stack.GetNext()
}
func GcStackPush(stack **GcStack, top *int, ref *ZendRefcounted) {
	if UNEXPECTED((*top) == GC_STACK_SEGMENT_SIZE) {
		*stack = GcStackNext(*stack)
		*top = 0
	}
	(*stack).GetData()[b.PostInc(&(*top))] = ref
}
func GcStackPop(stack **GcStack, top *int) *ZendRefcounted {
	if UNEXPECTED((*top) == 0) {
		if (*stack).GetPrev() == nil {
			return nil
		} else {
			*stack = (*stack).GetPrev()
			*top = GC_STACK_SEGMENT_SIZE - 1
			return (*stack).GetData()[GC_STACK_SEGMENT_SIZE-1]
		}
	} else {
		return (*stack).GetData()[b.PreDec(&(*top))]
	}
}
func GcStackFree(stack *GcStack) {
	var p *GcStack = stack.GetNext()
	for p != nil {
		stack = p.GetNext()
		Efree(p)
		p = stack
	}
}
func GcCompress(idx uint32) uint32 {
	if EXPECTED(idx < GC_MAX_UNCOMPRESSED) {
		return idx
	}
	return idx%GC_MAX_UNCOMPRESSED | GC_MAX_UNCOMPRESSED
}
func GcDecompress(ref *ZendRefcounted, idx uint32) *GcRootBuffer {
	var root *GcRootBuffer = GC_IDX2PTR(idx)
	if EXPECTED(GC_GET_PTR(root.GetRef()) == ref) {
		return root
	}
	for true {
		idx += GC_MAX_UNCOMPRESSED
		ZEND_ASSERT(idx < GcGlobals.GetFirstUnused())
		root = GC_IDX2PTR(idx)
		if GC_GET_PTR(root.GetRef()) == ref {
			return root
		}
	}
}
func GcFetchUnused() uint32 {
	var idx uint32
	var root *GcRootBuffer
	ZEND_ASSERT(GC_HAS_UNUSED())
	idx = GcGlobals.GetUnused()
	root = GC_IDX2PTR(idx)
	ZEND_ASSERT(GC_IS_UNUSED(root.GetRef()))
	GcGlobals.SetUnused(GC_LIST2IDX(root.GetRef()))
	return idx
}
func GcLinkUnused(root *GcRootBuffer) {
	root.SetRef(GC_IDX2LIST(GcGlobals.GetUnused()))
	GcGlobals.SetUnused(GC_PTR2IDX(root))
}
func GcFetchNextUnused() uint32 {
	var idx uint32
	ZEND_ASSERT(GC_HAS_NEXT_UNUSED())
	idx = GcGlobals.GetFirstUnused()
	GcGlobals.SetFirstUnused(GcGlobals.GetFirstUnused() + 1)
	return idx
}
func GcRemoveFromRoots(root *GcRootBuffer) {
	GC_LINK_UNUSED(root)
	GcGlobals.GetNumRoots()--
}
func RootBufferDtor(gc_globals *ZendGcGlobals) {
	if gc_globals.GetBuf() != nil {
		Free(gc_globals.GetBuf())
		gc_globals.SetBuf(nil)
	}
}
func GcGlobalsCtorEx(gc_globals *ZendGcGlobals) {
	gc_globals.SetGcEnabled(0)
	gc_globals.SetGcActive(0)
	gc_globals.SetGcProtected(1)
	gc_globals.SetGcFull(0)
	gc_globals.SetBuf(nil)
	gc_globals.SetUnused(GC_INVALID)
	gc_globals.SetFirstUnused(GC_INVALID)
	gc_globals.SetGcThreshold(GC_INVALID)
	gc_globals.SetBufSize(GC_INVALID)
	gc_globals.SetNumRoots(0)
	gc_globals.SetGcRuns(0)
	gc_globals.SetCollected(0)
}
func GcGlobalsCtor() { GcGlobalsCtorEx(&GcGlobals) }
func GcGlobalsDtor() { RootBufferDtor(&GcGlobals) }
func GcReset() {
	if GcGlobals.GetBuf() != nil {
		GcGlobals.SetGcActive(0)
		GcGlobals.SetGcProtected(0)
		GcGlobals.SetGcFull(0)
		GcGlobals.SetUnused(GC_INVALID)
		GcGlobals.SetFirstUnused(GC_FIRST_ROOT)
		GcGlobals.SetNumRoots(0)
		GcGlobals.SetGcRuns(0)
		GcGlobals.SetCollected(0)
	}
}
func GcEnable(enable ZendBool) ZendBool {
	var old_enabled ZendBool = GcGlobals.GetGcEnabled()
	GcGlobals.SetGcEnabled(enable)
	if enable != 0 && old_enabled == 0 && GcGlobals.GetBuf() == nil {
		GcGlobals.SetBuf((*GcRootBuffer)(Pemalloc(b.SizeOf("gc_root_buffer")*GC_DEFAULT_BUF_SIZE, 1)))
		GcGlobals.GetBuf()[0].SetRef(nil)
		GcGlobals.SetBufSize(GC_DEFAULT_BUF_SIZE)
		GcGlobals.SetGcThreshold(GC_THRESHOLD_DEFAULT + GC_FIRST_ROOT)
		GcReset()
	}
	return old_enabled
}
func GcEnabled() ZendBool { return GcGlobals.GetGcEnabled() }
func GcProtect(protect ZendBool) ZendBool {
	var old_protected ZendBool = GcGlobals.GetGcProtected()
	GcGlobals.SetGcProtected(protect)
	return old_protected
}
func GcProtected() ZendBool { return GcGlobals.GetGcProtected() }
func GcGrowRootBuffer() {
	var new_size int
	if GcGlobals.GetBufSize() >= GC_MAX_BUF_SIZE {
		if GcGlobals.GetGcFull() == 0 {
			ZendError(E_WARNING, "GC buffer overflow (GC disabled)\n")
			GcGlobals.SetGcActive(1)
			GcGlobals.SetGcProtected(1)
			GcGlobals.SetGcFull(1)
			return
		}
	}
	if GcGlobals.GetBufSize() < GC_BUF_GROW_STEP {
		new_size = GcGlobals.GetBufSize() * 2
	} else {
		new_size = GcGlobals.GetBufSize() + GC_BUF_GROW_STEP
	}
	if new_size > GC_MAX_BUF_SIZE {
		new_size = GC_MAX_BUF_SIZE
	}
	GcGlobals.SetBuf(Perealloc(GcGlobals.GetBuf(), b.SizeOf("gc_root_buffer")*new_size, 1))
	GcGlobals.SetBufSize(new_size)
}
func GcAdjustThreshold(count int) {
	var new_threshold uint32

	/* TODO Very simple heuristic for dynamic GC buffer resizing:
	 * If there are "too few" collections, increase the collection threshold
	 * by a fixed step */

	if count < GC_THRESHOLD_TRIGGER {

		/* increase */

		if GcGlobals.GetGcThreshold() < GC_THRESHOLD_MAX {
			new_threshold = GcGlobals.GetGcThreshold() + GC_THRESHOLD_STEP
			if new_threshold > GC_THRESHOLD_MAX {
				new_threshold = GC_THRESHOLD_MAX
			}
			if new_threshold > GcGlobals.GetBufSize() {
				GcGrowRootBuffer()
			}
			if new_threshold <= GcGlobals.GetBufSize() {
				GcGlobals.SetGcThreshold(new_threshold)
			}
		}

		/* increase */

	} else if GcGlobals.GetGcThreshold() > GC_THRESHOLD_DEFAULT {
		new_threshold = GcGlobals.GetGcThreshold() - GC_THRESHOLD_STEP
		if new_threshold < GC_THRESHOLD_DEFAULT {
			new_threshold = GC_THRESHOLD_DEFAULT
		}
		GcGlobals.SetGcThreshold(new_threshold)
	}

	/* TODO Very simple heuristic for dynamic GC buffer resizing:
	 * If there are "too few" collections, increase the collection threshold
	 * by a fixed step */
}
func GcPossibleRootWhenFull(ref *ZendRefcounted) {
	var idx uint32
	var newRoot *GcRootBuffer
	ZEND_ASSERT(GC_TYPE(ref) == IS_ARRAY || GC_TYPE(ref) == IS_OBJECT)
	ZEND_ASSERT(GC_INFO(ref) == 0)
	if GcGlobals.GetGcEnabled() != 0 && GcGlobals.GetGcActive() == 0 {
		GC_ADDREF(ref)
		GcAdjustThreshold(GcCollectCycles())
		if UNEXPECTED(GC_DELREF(ref) != 0) == 0 {
			RcDtorFunc(ref)
			return
		} else if UNEXPECTED(GC_INFO(ref) != 0) {
			return
		}
	}
	if GC_HAS_UNUSED() {
		idx = GC_FETCH_UNUSED()
	} else if EXPECTED(GC_HAS_NEXT_UNUSED()) {
		idx = GC_FETCH_NEXT_UNUSED()
	} else {
		GcGrowRootBuffer()
		if UNEXPECTED(!(GC_HAS_NEXT_UNUSED())) {
			return
		}
		idx = GC_FETCH_NEXT_UNUSED()
	}
	newRoot = GC_IDX2PTR(idx)
	newRoot.SetRef(ref)
	idx = GcCompress(idx)
	GC_REF_SET_INFO(ref, idx|GC_PURPLE)
	GcGlobals.GetNumRoots()++
}
func GcPossibleRoot(ref *ZendRefcounted) {
	var idx uint32
	var newRoot *GcRootBuffer
	if UNEXPECTED(GcGlobals.GetGcProtected() != 0) {
		return
	}
	if EXPECTED(GC_HAS_UNUSED()) {
		idx = GC_FETCH_UNUSED()
	} else if EXPECTED(GC_HAS_NEXT_UNUSED_UNDER_THRESHOLD()) {
		idx = GC_FETCH_NEXT_UNUSED()
	} else {
		GcPossibleRootWhenFull(ref)
		return
	}
	ZEND_ASSERT(GC_TYPE(ref) == IS_ARRAY || GC_TYPE(ref) == IS_OBJECT)
	ZEND_ASSERT(GC_INFO(ref) == 0)
	newRoot = GC_IDX2PTR(idx)
	newRoot.SetRef(ref)
	idx = GcCompress(idx)
	GC_REF_SET_INFO(ref, idx|GC_PURPLE)
	GcGlobals.GetNumRoots()++
}
func GcRemoveCompressed(ref *ZendRefcounted, idx uint32) {
	var root *GcRootBuffer = GcDecompress(ref, idx)
	GcRemoveFromRoots(root)
}
func GcRemoveFromBuffer(ref *ZendRefcounted) {
	var root *GcRootBuffer
	var idx uint32 = GC_REF_ADDRESS(ref)
	if !(GC_REF_CHECK_COLOR(ref, GC_BLACK)) {

	}
	GC_REF_SET_INFO(ref, 0)

	/* Perform decompression only in case of large buffers */

	if UNEXPECTED(GcGlobals.GetFirstUnused() >= GC_MAX_UNCOMPRESSED) {
		GcRemoveCompressed(ref, idx)
		return
	}
	ZEND_ASSERT(idx != 0)
	root = GC_IDX2PTR(idx)
	GcRemoveFromRoots(root)
}
func GcScanBlack(ref *ZendRefcounted, stack *GcStack) {
	var ht *HashTable = nil
	var p *Bucket
	var end *Bucket
	var zv *Zval
	GC_STACK_DCL(stack)
tail_call:
	if GC_TYPE(ref) == IS_OBJECT {
		var obj *ZendObject = (*ZendObject)(ref)
		if EXPECTED((OBJ_FLAGS(ref) & IS_OBJ_FREE_CALLED) == 0) {
			var n int
			var zv *Zval
			var end *Zval
			var tmp Zval
			ZVAL_OBJ(&tmp, obj)
			ht = obj.GetHandlers().GetGetGc()(&tmp, &zv, &n)
			end = zv + n
			if EXPECTED(ht == nil) || UNEXPECTED(GC_REF_CHECK_COLOR(ht, GC_BLACK)) {
				ht = nil
				if n == 0 {
					goto next
				}
				for !(Z_REFCOUNTED_P(b.PreDec(&end))) {
					if zv == end {
						goto next
					}
				}
			} else {
				GC_REF_SET_BLACK(ht)
			}
			for zv != end {
				if Z_REFCOUNTED_P(zv) {
					ref = Z_COUNTED_P(zv)
					GC_ADDREF(ref)
					if !(GC_REF_CHECK_COLOR(ref, GC_BLACK)) {
						GC_REF_SET_BLACK(ref)
						GC_STACK_PUSH(ref)
					}
				}
				zv++
			}
			if EXPECTED(ht == nil) {
				ref = Z_COUNTED_P(zv)
				GC_ADDREF(ref)
				if !(GC_REF_CHECK_COLOR(ref, GC_BLACK)) {
					GC_REF_SET_BLACK(ref)
					goto tail_call
				}
				goto next
			}
		} else {
			goto next
		}
	} else if GC_TYPE(ref) == IS_ARRAY {
		if (*ZendArray)(ref != &(ExecutorGlobals.GetSymbolTable())) != nil {
			ht = (*ZendArray)(ref)
		} else {
			goto next
		}
	} else if GC_TYPE(ref) == IS_REFERENCE {
		if Z_REFCOUNTED((*ZendReference)(ref).GetVal()) {
			ref = Z_COUNTED((*ZendReference)(ref).GetVal())
			GC_ADDREF(ref)
			if !(GC_REF_CHECK_COLOR(ref, GC_BLACK)) {
				GC_REF_SET_BLACK(ref)
				goto tail_call
			}
		}
		goto next
	} else {
		goto next
	}
	if ht.GetNNumUsed() == 0 {
		goto next
	}
	p = ht.GetArData()
	end = p + ht.GetNNumUsed()
	for true {
		end--
		zv = &end.val
		if Z_TYPE_P(zv) == IS_INDIRECT {
			zv = Z_INDIRECT_P(zv)
		}
		if Z_REFCOUNTED_P(zv) {
			break
		}
		if p == end {
			goto next
		}
	}
	for p != end {
		zv = &p.val
		if Z_TYPE_P(zv) == IS_INDIRECT {
			zv = Z_INDIRECT_P(zv)
		}
		if Z_REFCOUNTED_P(zv) {
			ref = Z_COUNTED_P(zv)
			GC_ADDREF(ref)
			if !(GC_REF_CHECK_COLOR(ref, GC_BLACK)) {
				GC_REF_SET_BLACK(ref)
				GC_STACK_PUSH(ref)
			}
		}
		p++
	}
	zv = &p.val
	if Z_TYPE_P(zv) == IS_INDIRECT {
		zv = Z_INDIRECT_P(zv)
	}
	ref = Z_COUNTED_P(zv)
	GC_ADDREF(ref)
	if !(GC_REF_CHECK_COLOR(ref, GC_BLACK)) {
		GC_REF_SET_BLACK(ref)
		goto tail_call
	}
next:
	ref = GC_STACK_POP()
	if ref != nil {
		goto tail_call
	}
}
func GcMarkGrey(ref *ZendRefcounted, stack *GcStack) {
	var ht *HashTable = nil
	var p *Bucket
	var end *Bucket
	var zv *Zval
	GC_STACK_DCL(stack)
	for {
		if GC_TYPE(ref) == IS_OBJECT {
			var obj *ZendObject = (*ZendObject)(ref)
			if EXPECTED((OBJ_FLAGS(ref) & IS_OBJ_FREE_CALLED) == 0) {
				var n int
				var zv *Zval
				var end *Zval
				var tmp Zval
				ZVAL_OBJ(&tmp, obj)
				ht = obj.GetHandlers().GetGetGc()(&tmp, &zv, &n)
				end = zv + n
				if EXPECTED(ht == nil) || UNEXPECTED(GC_REF_CHECK_COLOR(ht, GC_GREY)) {
					ht = nil
					if n == 0 {
						goto next
					}
					for !(Z_REFCOUNTED_P(b.PreDec(&end))) {
						if zv == end {
							goto next
						}
					}
				} else {
					GC_REF_SET_COLOR(ht, GC_GREY)
				}
				for zv != end {
					if Z_REFCOUNTED_P(zv) {
						ref = Z_COUNTED_P(zv)
						GC_DELREF(ref)
						if !(GC_REF_CHECK_COLOR(ref, GC_GREY)) {
							GC_REF_SET_COLOR(ref, GC_GREY)
							GC_STACK_PUSH(ref)
						}
					}
					zv++
				}
				if EXPECTED(ht == nil) {
					ref = Z_COUNTED_P(zv)
					GC_DELREF(ref)
					if !(GC_REF_CHECK_COLOR(ref, GC_GREY)) {
						GC_REF_SET_COLOR(ref, GC_GREY)
						continue
					}
					goto next
				}
			} else {
				goto next
			}
		} else if GC_TYPE(ref) == IS_ARRAY {
			if (*ZendArray)(ref) == &(ExecutorGlobals.GetSymbolTable()) {
				GC_REF_SET_BLACK(ref)
				goto next
			} else {
				ht = (*ZendArray)(ref)
			}
		} else if GC_TYPE(ref) == IS_REFERENCE {
			if Z_REFCOUNTED((*ZendReference)(ref).GetVal()) {
				ref = Z_COUNTED((*ZendReference)(ref).GetVal())
				GC_DELREF(ref)
				if !(GC_REF_CHECK_COLOR(ref, GC_GREY)) {
					GC_REF_SET_COLOR(ref, GC_GREY)
					continue
				}
			}
			goto next
		} else {
			goto next
		}
		if ht.GetNNumUsed() == 0 {
			goto next
		}
		p = ht.GetArData()
		end = p + ht.GetNNumUsed()
		for true {
			end--
			zv = &end.val
			if Z_TYPE_P(zv) == IS_INDIRECT {
				zv = Z_INDIRECT_P(zv)
			}
			if Z_REFCOUNTED_P(zv) {
				break
			}
			if p == end {
				goto next
			}
		}
		for p != end {
			zv = &p.val
			if Z_TYPE_P(zv) == IS_INDIRECT {
				zv = Z_INDIRECT_P(zv)
			}
			if Z_REFCOUNTED_P(zv) {
				ref = Z_COUNTED_P(zv)
				GC_DELREF(ref)
				if !(GC_REF_CHECK_COLOR(ref, GC_GREY)) {
					GC_REF_SET_COLOR(ref, GC_GREY)
					GC_STACK_PUSH(ref)
				}
			}
			p++
		}
		zv = &p.val
		if Z_TYPE_P(zv) == IS_INDIRECT {
			zv = Z_INDIRECT_P(zv)
		}
		ref = Z_COUNTED_P(zv)
		GC_DELREF(ref)
		if !(GC_REF_CHECK_COLOR(ref, GC_GREY)) {
			GC_REF_SET_COLOR(ref, GC_GREY)
			continue
		}
	next:
		ref = GC_STACK_POP()
		if ref == nil {
			break
		}
	}
}
func GcCompact() {
	if GcGlobals.GetNumRoots()+GC_FIRST_ROOT != GcGlobals.GetFirstUnused() {
		if GcGlobals.GetNumRoots() != 0 {
			var free *GcRootBuffer = GC_IDX2PTR(GC_FIRST_ROOT)
			var scan *GcRootBuffer = GC_IDX2PTR(GcGlobals.GetFirstUnused() - 1)
			var end *GcRootBuffer = GC_IDX2PTR(GcGlobals.GetNumRoots())
			var idx uint32
			var p *ZendRefcounted
			for free < scan {
				for !(GC_IS_UNUSED(free.GetRef())) {
					free++
				}
				for GC_IS_UNUSED(scan.GetRef()) {
					scan--
				}
				if scan > free {
					p = scan.GetRef()
					free.SetRef(p)
					p = GC_GET_PTR(p)
					idx = GcCompress(GC_PTR2IDX(free))
					GC_REF_SET_INFO(p, idx|GC_REF_COLOR(p))
					free++
					scan--
					if scan <= end {
						break
					}
				}
			}
		}
		GcGlobals.SetUnused(GC_INVALID)
		GcGlobals.SetFirstUnused(GcGlobals.GetNumRoots() + GC_FIRST_ROOT)
	}
}
func GcMarkRoots(stack *GcStack) {
	var current *GcRootBuffer
	var last *GcRootBuffer
	GcCompact()
	current = GC_IDX2PTR(GC_FIRST_ROOT)
	last = GC_IDX2PTR(GcGlobals.GetFirstUnused())
	for current != last {
		if GC_IS_ROOT(current.GetRef()) {
			if GC_REF_CHECK_COLOR(current.GetRef(), GC_PURPLE) {
				GC_REF_SET_COLOR(current.GetRef(), GC_GREY)
				GcMarkGrey(current.GetRef(), stack)
			}
		}
		current++
	}
}
func GcScan(ref *ZendRefcounted, stack *GcStack) {
	var ht *HashTable = nil
	var p *Bucket
	var end *Bucket
	var zv *Zval
	GC_STACK_DCL(stack)
tail_call:
	if GC_REF_CHECK_COLOR(ref, GC_WHITE) {
		if GC_REFCOUNT(ref) > 0 {
			if !(GC_REF_CHECK_COLOR(ref, GC_BLACK)) {
				GC_REF_SET_BLACK(ref)
				if UNEXPECTED(!(_stack.next)) {
					GcStackNext(_stack)
				}

				/* Split stack and reuse the tail */

				_stack.next.prev = nil
				GcScanBlack(ref, _stack.next)
				_stack.next.prev = _stack
			}
		} else {
			if GC_TYPE(ref) == IS_OBJECT {
				var obj *ZendObject = (*ZendObject)(ref)
				if EXPECTED((OBJ_FLAGS(ref) & IS_OBJ_FREE_CALLED) == 0) {
					var n int
					var zv *Zval
					var end *Zval
					var tmp Zval
					ZVAL_OBJ(&tmp, obj)
					ht = obj.GetHandlers().GetGetGc()(&tmp, &zv, &n)
					end = zv + n
					if EXPECTED(ht == nil) || UNEXPECTED(!(GC_REF_CHECK_COLOR(ht, GC_GREY))) {
						ht = nil
						if n == 0 {
							goto next
						}
						for !(Z_REFCOUNTED_P(b.PreDec(&end))) {
							if zv == end {
								goto next
							}
						}
					} else {
						GC_REF_SET_COLOR(ht, GC_WHITE)
					}
					for zv != end {
						if Z_REFCOUNTED_P(zv) {
							ref = Z_COUNTED_P(zv)
							if GC_REF_CHECK_COLOR(ref, GC_GREY) {
								GC_REF_SET_COLOR(ref, GC_WHITE)
								GC_STACK_PUSH(ref)
							}
						}
						zv++
					}
					if EXPECTED(ht == nil) {
						ref = Z_COUNTED_P(zv)
						if GC_REF_CHECK_COLOR(ref, GC_GREY) {
							GC_REF_SET_COLOR(ref, GC_WHITE)
							goto tail_call
						}
						goto next
					}
				} else {
					goto next
				}
			} else if GC_TYPE(ref) == IS_ARRAY {
				if (*ZendArray)(ref == &(ExecutorGlobals.GetSymbolTable())) != nil {
					GC_REF_SET_BLACK(ref)
					goto next
				} else {
					ht = (*ZendArray)(ref)
				}
			} else if GC_TYPE(ref) == IS_REFERENCE {
				if Z_REFCOUNTED((*ZendReference)(ref).GetVal()) {
					ref = Z_COUNTED((*ZendReference)(ref).GetVal())
					if GC_REF_CHECK_COLOR(ref, GC_GREY) {
						GC_REF_SET_COLOR(ref, GC_WHITE)
						goto tail_call
					}
				}
				goto next
			} else {
				goto next
			}
			if ht.GetNNumUsed() == 0 {
				goto next
			}
			p = ht.GetArData()
			end = p + ht.GetNNumUsed()
			for true {
				end--
				zv = &end.val
				if Z_TYPE_P(zv) == IS_INDIRECT {
					zv = Z_INDIRECT_P(zv)
				}
				if Z_REFCOUNTED_P(zv) {
					break
				}
				if p == end {
					goto next
				}
			}
			for p != end {
				zv = &p.val
				if Z_TYPE_P(zv) == IS_INDIRECT {
					zv = Z_INDIRECT_P(zv)
				}
				if Z_REFCOUNTED_P(zv) {
					ref = Z_COUNTED_P(zv)
					if GC_REF_CHECK_COLOR(ref, GC_GREY) {
						GC_REF_SET_COLOR(ref, GC_WHITE)
						GC_STACK_PUSH(ref)
					}
				}
				p++
			}
			zv = &p.val
			if Z_TYPE_P(zv) == IS_INDIRECT {
				zv = Z_INDIRECT_P(zv)
			}
			ref = Z_COUNTED_P(zv)
			if GC_REF_CHECK_COLOR(ref, GC_GREY) {
				GC_REF_SET_COLOR(ref, GC_WHITE)
				goto tail_call
			}
		}
	}
next:
	ref = GC_STACK_POP()
	if ref != nil {
		goto tail_call
	}
}
func GcScanRoots(stack *GcStack) {
	var current *GcRootBuffer = GC_IDX2PTR(GC_FIRST_ROOT)
	var last *GcRootBuffer = GC_IDX2PTR(GcGlobals.GetFirstUnused())
	for current != last {
		if GC_IS_ROOT(current.GetRef()) {
			if GC_REF_CHECK_COLOR(current.GetRef(), GC_GREY) {
				GC_REF_SET_COLOR(current.GetRef(), GC_WHITE)
				GcScan(current.GetRef(), stack)
			}
		}
		current++
	}
}
func GcAddGarbage(ref *ZendRefcounted) {
	var idx uint32
	var buf *GcRootBuffer
	if GC_HAS_UNUSED() {
		idx = GC_FETCH_UNUSED()
	} else if GC_HAS_NEXT_UNUSED() {
		idx = GC_FETCH_NEXT_UNUSED()
	} else {
		GcGrowRootBuffer()
		if UNEXPECTED(!(GC_HAS_NEXT_UNUSED())) {
			return
		}
		idx = GC_FETCH_NEXT_UNUSED()
	}
	buf = GC_IDX2PTR(idx)
	buf.SetRef(GC_MAKE_GARBAGE(ref))
	idx = GcCompress(idx)
	GC_REF_SET_INFO(ref, idx|GC_BLACK)
	GcGlobals.GetNumRoots()++
}
func GcCollectWhite(ref *ZendRefcounted, flags *uint32, stack *GcStack) int {
	var count int = 0
	var ht *HashTable = nil
	var p *Bucket
	var end *Bucket
	var zv *Zval
	GC_STACK_DCL(stack)
	for {

		/* don't count references for compatibility ??? */

		if GC_TYPE(ref) != IS_REFERENCE {
			count++
		}
		if GC_TYPE(ref) == IS_OBJECT {
			var obj *ZendObject = (*ZendObject)(ref)
			if EXPECTED((OBJ_FLAGS(ref) & IS_OBJ_FREE_CALLED) == 0) {
				var n int
				var zv *Zval
				var end *Zval
				var tmp Zval

				/* optimization: color is GC_BLACK (0) */

				if GC_INFO(ref) == 0 {
					GcAddGarbage(ref)
				}
				if (OBJ_FLAGS(obj)&IS_OBJ_DESTRUCTOR_CALLED) == 0 && (obj.GetHandlers().GetDtorObj() != ZendObjectsDestroyObject || obj.GetCe().GetDestructor() != nil) {
					*flags |= GC_HAS_DESTRUCTORS
				}
				ZVAL_OBJ(&tmp, obj)
				ht = obj.GetHandlers().GetGetGc()(&tmp, &zv, &n)
				end = zv + n
				if EXPECTED(ht == nil) || UNEXPECTED(GC_REF_CHECK_COLOR(ht, GC_BLACK)) {
					ht = nil
					if n == 0 {
						goto next
					}
					for !(Z_REFCOUNTED_P(b.PreDec(&end))) {
						if zv == end {
							goto next
						}
					}
				} else {
					GC_REF_SET_BLACK(ht)
				}
				for zv != end {
					if Z_REFCOUNTED_P(zv) {
						ref = Z_COUNTED_P(zv)
						GC_ADDREF(ref)
						if GC_REF_CHECK_COLOR(ref, GC_WHITE) {
							GC_REF_SET_BLACK(ref)
							GC_STACK_PUSH(ref)
						}
					}
					zv++
				}
				if EXPECTED(ht == nil) {
					ref = Z_COUNTED_P(zv)
					GC_ADDREF(ref)
					if GC_REF_CHECK_COLOR(ref, GC_WHITE) {
						GC_REF_SET_BLACK(ref)
						continue
					}
					goto next
				}
			} else {
				goto next
			}
		} else if GC_TYPE(ref) == IS_ARRAY {

			/* optimization: color is GC_BLACK (0) */

			if GC_INFO(ref) == 0 {
				GcAddGarbage(ref)
			}
			ht = (*ZendArray)(ref)
		} else if GC_TYPE(ref) == IS_REFERENCE {
			if Z_REFCOUNTED((*ZendReference)(ref).GetVal()) {
				ref = Z_COUNTED((*ZendReference)(ref).GetVal())
				GC_ADDREF(ref)
				if GC_REF_CHECK_COLOR(ref, GC_WHITE) {
					GC_REF_SET_BLACK(ref)
					continue
				}
			}
			goto next
		} else {
			goto next
		}
		if ht.GetNNumUsed() == 0 {
			goto next
		}
		p = ht.GetArData()
		end = p + ht.GetNNumUsed()
		for true {
			end--
			zv = &end.val
			if Z_TYPE_P(zv) == IS_INDIRECT {
				zv = Z_INDIRECT_P(zv)
			}
			if Z_REFCOUNTED_P(zv) {
				break
			}
			if p == end {
				goto next
			}
		}
		for p != end {
			zv = &p.val
			if Z_TYPE_P(zv) == IS_INDIRECT {
				zv = Z_INDIRECT_P(zv)
			}
			if Z_REFCOUNTED_P(zv) {
				ref = Z_COUNTED_P(zv)
				GC_ADDREF(ref)
				if GC_REF_CHECK_COLOR(ref, GC_WHITE) {
					GC_REF_SET_BLACK(ref)
					GC_STACK_PUSH(ref)
				}
			}
			p++
		}
		zv = &p.val
		if Z_TYPE_P(zv) == IS_INDIRECT {
			zv = Z_INDIRECT_P(zv)
		}
		ref = Z_COUNTED_P(zv)
		GC_ADDREF(ref)
		if GC_REF_CHECK_COLOR(ref, GC_WHITE) {
			GC_REF_SET_BLACK(ref)
			continue
		}
	next:
		ref = GC_STACK_POP()
		if ref == nil {
			break
		}
	}
	return count
}
func GcCollectRoots(flags *uint32, stack *GcStack) int {
	var idx uint32
	var end uint32
	var ref *ZendRefcounted
	var count int = 0
	var current *GcRootBuffer = GC_IDX2PTR(GC_FIRST_ROOT)
	var last *GcRootBuffer = GC_IDX2PTR(GcGlobals.GetFirstUnused())

	/* remove non-garbage from the list */

	for current != last {
		if GC_IS_ROOT(current.GetRef()) {
			if GC_REF_CHECK_COLOR(current.GetRef(), GC_BLACK) {
				GC_REF_SET_INFO(current.GetRef(), 0)
				GcRemoveFromRoots(current)
			}
		}
		current++
	}
	GcCompact()

	/* Root buffer might be reallocated during gc_collect_white,
	 * make sure to reload pointers. */

	idx = GC_FIRST_ROOT
	end = GcGlobals.GetFirstUnused()
	for idx != end {
		current = GC_IDX2PTR(idx)
		ref = current.GetRef()
		ZEND_ASSERT(GC_IS_ROOT(ref))
		current.SetRef(GC_MAKE_GARBAGE(ref))
		if GC_REF_CHECK_COLOR(ref, GC_WHITE) {
			GC_REF_SET_BLACK(ref)
			count += GcCollectWhite(ref, flags, stack)
		}
		idx++
	}
	return count
}
func GcRemoveNestedDataFromBuffer(ref *ZendRefcounted, root *GcRootBuffer) int {
	var ht *HashTable = nil
	var p *Bucket
	var end *Bucket
	var zv *Zval
	var count int = 0
tail_call:
	for {
		if root != nil {
			root = nil
			count++
		} else if GC_REF_ADDRESS(ref) != 0 && GC_REF_CHECK_COLOR(ref, GC_BLACK) {
			GC_REMOVE_FROM_BUFFER(ref)
			count++
		} else if GC_TYPE(ref) == IS_REFERENCE {
			if Z_REFCOUNTED((*ZendReference)(ref).GetVal()) {
				ref = Z_COUNTED((*ZendReference)(ref).GetVal())
				goto tail_call
			}
			return count
		} else {
			return count
		}
		if GC_TYPE(ref) == IS_OBJECT {
			var obj *ZendObject = (*ZendObject)(ref)
			if EXPECTED((OBJ_FLAGS(ref) & IS_OBJ_FREE_CALLED) == 0) {
				var n int
				var zv *Zval
				var end *Zval
				var tmp Zval
				ZVAL_OBJ(&tmp, obj)
				ht = obj.GetHandlers().GetGetGc()(&tmp, &zv, &n)
				end = zv + n
				if EXPECTED(ht == nil) {
					if n == 0 {
						return count
					}
					for !(Z_REFCOUNTED_P(b.PreDec(&end))) {
						if zv == end {
							return count
						}
					}
				}
				for zv != end {
					if Z_REFCOUNTED_P(zv) {
						ref = Z_COUNTED_P(zv)
						count += GcRemoveNestedDataFromBuffer(ref, nil)
					}
					zv++
				}
				if EXPECTED(ht == nil) {
					ref = Z_COUNTED_P(zv)
					goto tail_call
				}
				if GC_REF_ADDRESS(ht) != 0 && GC_REF_CHECK_COLOR(ht, GC_BLACK) {
					GC_REMOVE_FROM_BUFFER(ht)
				}
			} else {
				return count
			}
		} else if GC_TYPE(ref) == IS_ARRAY {
			ht = (*ZendArray)(ref)
		} else {
			return count
		}
		if ht.GetNNumUsed() == 0 {
			return count
		}
		p = ht.GetArData()
		end = p + ht.GetNNumUsed()
		for true {
			end--
			zv = &end.val
			if Z_TYPE_P(zv) == IS_INDIRECT {
				zv = Z_INDIRECT_P(zv)
			}
			if Z_REFCOUNTED_P(zv) {
				break
			}
			if p == end {
				return count
			}
		}
		for p != end {
			zv = &p.val
			if Z_TYPE_P(zv) == IS_INDIRECT {
				zv = Z_INDIRECT_P(zv)
			}
			if Z_REFCOUNTED_P(zv) {
				ref = Z_COUNTED_P(zv)
				count += GcRemoveNestedDataFromBuffer(ref, nil)
			}
			p++
		}
		zv = &p.val
		if Z_TYPE_P(zv) == IS_INDIRECT {
			zv = Z_INDIRECT_P(zv)
		}
		ref = Z_COUNTED_P(zv)
		goto tail_call
		break
	}
}
func ZendGcCollectCycles() int {
	var count int = 0
	if GcGlobals.GetNumRoots() != 0 {
		var current *GcRootBuffer
		var last *GcRootBuffer
		var p *ZendRefcounted
		var gc_flags uint32 = 0
		var idx uint32
		var end uint32
		var stack GcStack
		stack.SetPrev(nil)
		stack.SetNext(nil)
		if GcGlobals.GetGcActive() != 0 {
			return 0
		}
		GcGlobals.GetGcRuns()++
		GcGlobals.SetGcActive(1)
		GcMarkRoots(&stack)
		GcScanRoots(&stack)
		count = GcCollectRoots(&gc_flags, &stack)
		GcStackFree(&stack)
		if GcGlobals.GetNumRoots() == 0 {

			/* nothing to free */

			GcGlobals.SetGcActive(0)
			return 0
		}
		end = GcGlobals.GetFirstUnused()
		if (gc_flags & GC_HAS_DESTRUCTORS) != 0 {

			/* During a destructor call, new externally visible references to nested data may
			 * be introduced. These references can be introduced in a way that does not
			 * modify any refcounts, so we have no real way to detect this situation
			 * short of rerunning full GC tracing. What we do instead is to only run
			 * destructors at this point, and leave the actual freeing of the objects
			 * until the next GC run. */

			idx = GC_FIRST_ROOT
			current = GC_IDX2PTR(GC_FIRST_ROOT)
			for idx != end {
				if GC_IS_GARBAGE(current.GetRef()) {
					p = GC_GET_PTR(current.GetRef())
					if GC_TYPE(p) == IS_OBJECT && (OBJ_FLAGS(p)&IS_OBJ_DESTRUCTOR_CALLED) == 0 {
						var obj *ZendObject = (*ZendObject)(p)
						if obj.GetHandlers().GetDtorObj() != ZendObjectsDestroyObject || obj.GetCe().GetDestructor() != nil {
							current.SetRef(GC_MAKE_DTOR_GARBAGE(obj))
							GC_REF_SET_COLOR(obj, GC_PURPLE)
						} else {
							GC_ADD_FLAGS(obj, IS_OBJ_DESTRUCTOR_CALLED)
						}
					}
				}
				current++
				idx++
			}

			/* Remove nested data for objects on which a destructor will be called.
			 * This will not remove the objects themselves, as they have been colored
			 * purple. */

			idx = GC_FIRST_ROOT
			current = GC_IDX2PTR(GC_FIRST_ROOT)
			for idx != end {
				if GC_IS_DTOR_GARBAGE(current.GetRef()) {
					p = GC_GET_PTR(current.GetRef())
					count -= GcRemoveNestedDataFromBuffer(p, current)
				}
				current++
				idx++
			}

			/* Actually call destructors.
			 *
			 * The root buffer might be reallocated during destructors calls,
			 * make sure to reload pointers as necessary. */

			idx = GC_FIRST_ROOT
			for idx != end {
				current = GC_IDX2PTR(idx)
				if GC_IS_DTOR_GARBAGE(current.GetRef()) {
					p = GC_GET_PTR(current.GetRef())

					/* Mark this is as a normal root for the next GC run,
					 * it's no longer garbage for this run. */

					current.SetRef(p)

					/* Double check that the destructor hasn't been called yet. It could have
					 * already been invoked indirectly by some other destructor. */

					if (OBJ_FLAGS(p) & IS_OBJ_DESTRUCTOR_CALLED) == 0 {
						var obj *ZendObject = (*ZendObject)(p)
						GC_ADD_FLAGS(obj, IS_OBJ_DESTRUCTOR_CALLED)
						GC_ADDREF(obj)
						obj.GetHandlers().GetDtorObj()(obj)
						GC_DELREF(obj)
					}

					/* Double check that the destructor hasn't been called yet. It could have
					 * already been invoked indirectly by some other destructor. */

				}
				idx++
			}
			if GcGlobals.GetGcProtected() != 0 {

				/* something went wrong */

				return 0

				/* something went wrong */

			}
		}

		/* Destroy zvals. The root buffer may be reallocated. */

		idx = GC_FIRST_ROOT
		for idx != end {
			current = GC_IDX2PTR(idx)
			if GC_IS_GARBAGE(current.GetRef()) {
				p = GC_GET_PTR(current.GetRef())
				if GC_TYPE(p) == IS_OBJECT {
					var obj *ZendObject = (*ZendObject)(p)
					ExecutorGlobals.GetObjectsStore().GetObjectBuckets()[obj.GetHandle()] = SET_OBJ_INVALID(obj)
					GC_TYPE_INFO(obj) = IS_NULL | GC_TYPE_INFO(obj) & ^GC_TYPE_MASK

					/* Modify current before calling free_obj (bug #78811: free_obj() can cause the root buffer (with current) to be reallocated.) */

					current.SetRef(GC_MAKE_GARBAGE((*byte)(obj) - obj.GetHandlers().GetOffset()))
					if (OBJ_FLAGS(obj) & IS_OBJ_FREE_CALLED) == 0 {
						GC_ADD_FLAGS(obj, IS_OBJ_FREE_CALLED)
						GC_ADDREF(obj)
						obj.GetHandlers().GetFreeObj()(obj)
						GC_DELREF(obj)
					}
					ZEND_OBJECTS_STORE_ADD_TO_FREE_LIST(obj.GetHandle())
				} else if GC_TYPE(p) == IS_ARRAY {
					var arr *ZendArray = (*ZendArray)(p)
					GC_TYPE_INFO(arr) = IS_NULL | GC_TYPE_INFO(arr) & ^GC_TYPE_MASK

					/* GC may destroy arrays with rc>1. This is valid and safe. */

					ZendHashDestroy(arr)
				}
			}
			idx++
		}

		/* Free objects */

		current = GC_IDX2PTR(GC_FIRST_ROOT)
		last = GC_IDX2PTR(end)
		for current != last {
			if GC_IS_GARBAGE(current.GetRef()) {
				p = GC_GET_PTR(current.GetRef())
				GC_LINK_UNUSED(current)
				GcGlobals.GetNumRoots()--
				Efree(p)
			}
			current++
		}
		GcGlobals.SetCollected(GcGlobals.GetCollected() + count)
		GcGlobals.SetGcActive(0)
	}
	GcCompact()
	return count
}
func ZendGcGetStatus(status *ZendGcStatus) {
	status.SetRuns(GcGlobals.GetGcRuns())
	status.SetCollected(GcGlobals.GetCollected())
	status.SetThreshold(GcGlobals.GetGcThreshold())
	status.SetNumRoots(GcGlobals.GetNumRoots())
}
