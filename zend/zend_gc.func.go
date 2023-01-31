// <<generate>>

package zend

import (
	b "sik/builtin"
)

func GC_REMOVE_FROM_BUFFER(p __auto__) {
	var _p *ZendRefcounted = (*ZendRefcounted)(p)
	if (_p.GetGcTypeInfo() & GC_INFO_MASK) != 0 {
		GcRemoveFromBuffer(_p)
	}
}
func GC_MAY_LEAK(ref *ZendRefcounted) bool {
	return (ref.GetGcTypeInfo() & (GC_INFO_MASK | GC_COLLECTABLE<<GC_FLAGS_SHIFT)) == GC_COLLECTABLE<<GC_FLAGS_SHIFT
}
func GcCheckPossibleRoot(ref *ZendRefcounted) {
	if ref.GetGcTypeInfo() == IS_REFERENCE {
		var zv *Zval = (*ZendReference)(ref).GetVal()
		if !(zv.IsCollectable()) {
			return
		}
		ref = zv.GetCounted()
	}
	if GC_MAY_LEAK(ref) {
		GcPossibleRoot(ref)
	}
}
func GC_REF_ADDRESS(ref __auto__) int {
	return (ref.GetGcTypeInfo() & GC_ADDRESS << GC_INFO_SHIFT) >> GC_INFO_SHIFT
}
func GC_REF_COLOR(ref *ZendRefcounted) int {
	return (ref.GetGcTypeInfo() & GC_COLOR << GC_INFO_SHIFT) >> GC_INFO_SHIFT
}
func GC_REF_CHECK_COLOR(ref __auto__, color __auto__) bool {
	return (ref.GetGcTypeInfo() & GC_COLOR << GC_INFO_SHIFT) == color<<GC_INFO_SHIFT
}
func GC_REF_SET_INFO(ref *ZendRefcounted, info int) {
	ref.GetGcTypeInfo() = ref.GetGcTypeInfo()&(GC_TYPE_MASK|GC_FLAGS_MASK) | info<<GC_INFO_SHIFT
}
func GC_REF_SET_COLOR(ref __auto__, c __auto__) {
	ref.GetGcTypeInfo() = ref.GetGcTypeInfo() & ^(GC_COLOR<<GC_INFO_SHIFT) | c<<GC_INFO_SHIFT
}
func GC_REF_SET_BLACK(ref __auto__) {
	ref.GetGcTypeInfo() &= ^(GC_COLOR << GC_INFO_SHIFT)
}
func GC_REF_SET_PURPLE(ref __auto__) {
	ref.GetGcTypeInfo() |= GC_COLOR << GC_INFO_SHIFT
}
func GC_GET_PTR(ptr *ZendRefcounted) any { return any(uintPtr(ptr) & ^GC_BITS) }
func GC_IS_ROOT(ptr *ZendRefcounted) bool {
	return (uintPtr(ptr) & GC_BITS) == GC_ROOT
}
func GC_IS_UNUSED(ptr *ZendRefcounted) bool {
	return (uintPtr(ptr) & GC_BITS) == GC_UNUSED
}
func GC_IS_GARBAGE(ptr *ZendRefcounted) bool {
	return (uintPtr(ptr) & GC_BITS) == GC_GARBAGE
}
func GC_IS_DTOR_GARBAGE(ptr *ZendRefcounted) bool {
	return (uintPtr(ptr) & GC_BITS) == GC_DTOR_GARBAGE
}
func GC_MAKE_GARBAGE(ptr __auto__) any { return any(uintPtr(ptr) | GC_GARBAGE) }
func GC_MAKE_DTOR_GARBAGE(ptr *ZendObject) any {
	return any(uintPtr(ptr) | GC_DTOR_GARBAGE)
}
func GC_IDX2PTR(idx *GcRootBuffer) __auto__ { return GC_G__().GetBuf() + idx }
func GC_PTR2IDX(ptr *GcRootBuffer) int      { return ptr - GC_G__().GetBuf() }
func GC_IDX2LIST(idx uint32) any {
	return any(uintPtr(idx*b.SizeOf("void *") | GC_UNUSED))
}
func GC_LIST2IDX(list *ZendRefcounted) int {
	return uint32(uintPtr)(list) / b.SizeOf("void *")
}
func GC_HAS_UNUSED() bool {
	return GC_G__().GetUnused() != GC_INVALID
}
func GC_FETCH_UNUSED() uint32           { return GcFetchUnused() }
func GC_LINK_UNUSED(root *GcRootBuffer) { GcLinkUnused(root) }
func GC_HAS_NEXT_UNUSED_UNDER_THRESHOLD() bool {
	return GC_G__().GetFirstUnused() < GC_G__().GetGcThreshold()
}
func GC_HAS_NEXT_UNUSED() bool {
	return GC_G__().GetFirstUnused() != GC_G__().GetBufSize()
}
func GC_FETCH_NEXT_UNUSED() uint32 { return GcFetchNextUnused() }
func GC_STACK_DCL(init *GcStack) {
	var _stack *GcStack = init
	var _top int = 0
}
func GC_STACK_PUSH(ref *ZendRefcounted) { GcStackPush(&_stack, &_top, ref) }
func GC_STACK_POP() *ZendRefcounted     { return GcStackPop(&_stack, &_top) }
func GcStackNext(stack *GcStack) *GcStack {
	if stack.GetNext() == nil {
		var segment *GcStack = Emalloc(b.SizeOf("gc_stack"))
		segment.SetPrev(stack)
		segment.SetNext(nil)
		stack.SetNext(segment)
	}
	return stack.GetNext()
}
func GcStackPush(stack **GcStack, top *int, ref *ZendRefcounted) {
	if (*top) == GC_STACK_SEGMENT_SIZE {
		*stack = GcStackNext(*stack)
		*top = 0
	}
	stack.GetData()[b.PostInc(&(*top))] = ref
}
func GcStackPop(stack **GcStack, top *int) *ZendRefcounted {
	if (*top) == 0 {
		if stack.GetPrev() == nil {
			return nil
		} else {
			*stack = stack.GetPrev()
			*top = GC_STACK_SEGMENT_SIZE - 1
			return stack.GetData()[GC_STACK_SEGMENT_SIZE-1]
		}
	} else {
		return stack.GetData()[b.PreDec(&(*top))]
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
	if idx < GC_MAX_UNCOMPRESSED {
		return idx
	}
	return idx%GC_MAX_UNCOMPRESSED | GC_MAX_UNCOMPRESSED
}
func GcDecompress(ref *ZendRefcounted, idx uint32) *GcRootBuffer {
	var root *GcRootBuffer = GC_IDX2PTR(idx)
	if GC_GET_PTR(root.GetRef()) == ref {
		return root
	}
	for true {
		idx += GC_MAX_UNCOMPRESSED
		ZEND_ASSERT(idx < GC_G__().GetFirstUnused())
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
	idx = GC_G__().GetUnused()
	root = GC_IDX2PTR(idx)
	ZEND_ASSERT(GC_IS_UNUSED(root.GetRef()))
	GC_G__().SetUnused(GC_LIST2IDX(root.GetRef()))
	return idx
}
func GcLinkUnused(root *GcRootBuffer) {
	root.SetRef(GC_IDX2LIST(GC_G__().GetUnused()))
	GC_G__().SetUnused(GC_PTR2IDX(root))
}
func GcFetchNextUnused() uint32 {
	var idx uint32
	ZEND_ASSERT(GC_HAS_NEXT_UNUSED())
	idx = GC_G__().GetFirstUnused()
	GC_G__().SetFirstUnused(GC_G__().GetFirstUnused() + 1)
	return idx
}
func GcRemoveFromRoots(root *GcRootBuffer) {
	GC_LINK_UNUSED(root)
	GC_G__().GetNumRoots()--
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
	if GC_G__().GetBuf() != nil {
		GC_G__().SetGcActive(0)
		GC_G__().SetGcProtected(0)
		GC_G__().SetGcFull(0)
		GC_G__().SetUnused(GC_INVALID)
		GC_G__().SetFirstUnused(GC_FIRST_ROOT)
		GC_G__().SetNumRoots(0)
		GC_G__().SetGcRuns(0)
		GC_G__().SetCollected(0)
	}
}
func GcEnable(enable ZendBool) ZendBool {
	var old_enabled ZendBool = GC_G__().GetGcEnabled()
	GC_G__().SetGcEnabled(enable)
	if enable != 0 && old_enabled == 0 && GC_G__().GetBuf() == nil {
		GC_G__().SetBuf((*GcRootBuffer)(Pemalloc(b.SizeOf("gc_root_buffer")*GC_DEFAULT_BUF_SIZE, 1)))
		GC_G__().GetBuf()[0].SetRef(nil)
		GC_G__().SetBufSize(GC_DEFAULT_BUF_SIZE)
		GC_G__().SetGcThreshold(GC_THRESHOLD_DEFAULT + GC_FIRST_ROOT)
		GcReset()
	}
	return old_enabled
}
func GcEnabled() ZendBool { return GC_G__().GetGcEnabled() }
func GcProtect(protect ZendBool) ZendBool {
	var old_protected ZendBool = GC_G__().GetGcProtected()
	GC_G__().SetGcProtected(protect)
	return old_protected
}
func GcProtected() ZendBool { return GC_G__().GetGcProtected() }
func GcGrowRootBuffer() {
	var new_size int
	if GC_G__().GetBufSize() >= GC_MAX_BUF_SIZE {
		if GC_G__().GetGcFull() == 0 {
			ZendError(E_WARNING, "GC buffer overflow (GC disabled)\n")
			GC_G__().SetGcActive(1)
			GC_G__().SetGcProtected(1)
			GC_G__().SetGcFull(1)
			return
		}
	}
	if GC_G__().GetBufSize() < GC_BUF_GROW_STEP {
		new_size = GC_G__().GetBufSize() * 2
	} else {
		new_size = GC_G__().GetBufSize() + GC_BUF_GROW_STEP
	}
	if new_size > GC_MAX_BUF_SIZE {
		new_size = GC_MAX_BUF_SIZE
	}
	GC_G__().SetBuf(Perealloc(GC_G__().GetBuf(), b.SizeOf("gc_root_buffer")*new_size, 1))
	GC_G__().SetBufSize(new_size)
}
func GcAdjustThreshold(count int) {
	var new_threshold uint32

	/* TODO Very simple heuristic for dynamic GC buffer resizing:
	 * If there are "too few" collections, increase the collection threshold
	 * by a fixed step */

	if count < GC_THRESHOLD_TRIGGER {

		/* increase */

		if GC_G__().GetGcThreshold() < GC_THRESHOLD_MAX {
			new_threshold = GC_G__().GetGcThreshold() + GC_THRESHOLD_STEP
			if new_threshold > GC_THRESHOLD_MAX {
				new_threshold = GC_THRESHOLD_MAX
			}
			if new_threshold > GC_G__().GetBufSize() {
				GcGrowRootBuffer()
			}
			if new_threshold <= GC_G__().GetBufSize() {
				GC_G__().SetGcThreshold(new_threshold)
			}
		}

		/* increase */

	} else if GC_G__().GetGcThreshold() > GC_THRESHOLD_DEFAULT {
		new_threshold = GC_G__().GetGcThreshold() - GC_THRESHOLD_STEP
		if new_threshold < GC_THRESHOLD_DEFAULT {
			new_threshold = GC_THRESHOLD_DEFAULT
		}
		GC_G__().SetGcThreshold(new_threshold)
	}

	/* TODO Very simple heuristic for dynamic GC buffer resizing:
	 * If there are "too few" collections, increase the collection threshold
	 * by a fixed step */
}
func GcPossibleRootWhenFull(ref *ZendRefcounted) {
	var idx uint32
	var newRoot *GcRootBuffer
	ZEND_ASSERT(ref.GetGcType() == IS_ARRAY || ref.GetGcType() == IS_OBJECT)
	ZEND_ASSERT(ref.GetGcInfo() == 0)
	if GC_G__().GetGcEnabled() != 0 && GC_G__().GetGcActive() == 0 {
		ref.AddRefcount()
		GcAdjustThreshold(GcCollectCycles())
		if ref.DelRefcount() == 0 {
			RcDtorFunc(ref)
			return
		} else if ref.GetGcInfo() != 0 {
			return
		}
	}
	if GC_HAS_UNUSED() {
		idx = GC_FETCH_UNUSED()
	} else if GC_HAS_NEXT_UNUSED() {
		idx = GC_FETCH_NEXT_UNUSED()
	} else {
		GcGrowRootBuffer()
		if !(GC_HAS_NEXT_UNUSED()) {
			return
		}
		idx = GC_FETCH_NEXT_UNUSED()
	}
	newRoot = GC_IDX2PTR(idx)
	newRoot.SetRef(ref)
	idx = GcCompress(idx)
	GC_REF_SET_INFO(ref, idx|GC_PURPLE)
	GC_G__().GetNumRoots()++
}
func GcPossibleRoot(ref *ZendRefcounted) {
	var idx uint32
	var newRoot *GcRootBuffer
	if GC_G__().GetGcProtected() != 0 {
		return
	}
	if GC_HAS_UNUSED() {
		idx = GC_FETCH_UNUSED()
	} else if GC_HAS_NEXT_UNUSED_UNDER_THRESHOLD() {
		idx = GC_FETCH_NEXT_UNUSED()
	} else {
		GcPossibleRootWhenFull(ref)
		return
	}
	ZEND_ASSERT(ref.GetGcType() == IS_ARRAY || ref.GetGcType() == IS_OBJECT)
	ZEND_ASSERT(ref.GetGcInfo() == 0)
	newRoot = GC_IDX2PTR(idx)
	newRoot.SetRef(ref)
	idx = GcCompress(idx)
	GC_REF_SET_INFO(ref, idx|GC_PURPLE)
	GC_G__().GetNumRoots()++
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

	if GC_G__().GetFirstUnused() >= GC_MAX_UNCOMPRESSED {
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
	if ref.GetGcType() == IS_OBJECT {
		var obj *ZendObject = (*ZendObject)(ref)
		if (ref.GetGcFlags() & IS_OBJ_FREE_CALLED) == 0 {
			var n int
			var zv *Zval
			var end *Zval
			var tmp Zval
			ZVAL_OBJ(&tmp, obj)
			ht = obj.GetHandlers().GetGetGc()(&tmp, &zv, &n)
			end = zv + n
			if ht == nil || GC_REF_CHECK_COLOR(ht, GC_BLACK) {
				ht = nil
				if n == 0 {
					goto next
				}
				for !(b.PreDec(&end).IsRefcounted()) {
					if zv == end {
						goto next
					}
				}
			} else {
				GC_REF_SET_BLACK(ht)
			}
			for zv != end {
				if zv.IsRefcounted() {
					ref = zv.GetCounted()
					ref.AddRefcount()
					if !(GC_REF_CHECK_COLOR(ref, GC_BLACK)) {
						GC_REF_SET_BLACK(ref)
						GC_STACK_PUSH(ref)
					}
				}
				zv++
			}
			if ht == nil {
				ref = zv.GetCounted()
				ref.AddRefcount()
				if !(GC_REF_CHECK_COLOR(ref, GC_BLACK)) {
					GC_REF_SET_BLACK(ref)
					goto tail_call
				}
				goto next
			}
		} else {
			goto next
		}
	} else if ref.GetGcType() == IS_ARRAY {
		if (*ZendArray)(ref != EG__().GetSymbolTable()) != nil {
			ht = (*ZendArray)(ref)
		} else {
			goto next
		}
	} else if ref.GetGcType() == IS_REFERENCE {
		if (*ZendReference)(ref).GetVal().IsRefcounted() {
			ref = (*ZendReference)(ref).GetVal().GetCounted()
			ref.AddRefcount()
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
		zv = end.GetVal()
		if zv.IsIndirect() {
			zv = zv.GetZv()
		}
		if zv.IsRefcounted() {
			break
		}
		if p == end {
			goto next
		}
	}
	for p != end {
		zv = p.GetVal()
		if zv.IsIndirect() {
			zv = zv.GetZv()
		}
		if zv.IsRefcounted() {
			ref = zv.GetCounted()
			ref.AddRefcount()
			if !(GC_REF_CHECK_COLOR(ref, GC_BLACK)) {
				GC_REF_SET_BLACK(ref)
				GC_STACK_PUSH(ref)
			}
		}
		p++
	}
	zv = p.GetVal()
	if zv.IsIndirect() {
		zv = zv.GetZv()
	}
	ref = zv.GetCounted()
	ref.AddRefcount()
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
		if ref.GetGcType() == IS_OBJECT {
			var obj *ZendObject = (*ZendObject)(ref)
			if (ref.GetGcFlags() & IS_OBJ_FREE_CALLED) == 0 {
				var n int
				var zv *Zval
				var end *Zval
				var tmp Zval
				ZVAL_OBJ(&tmp, obj)
				ht = obj.GetHandlers().GetGetGc()(&tmp, &zv, &n)
				end = zv + n
				if ht == nil || GC_REF_CHECK_COLOR(ht, GC_GREY) {
					ht = nil
					if n == 0 {
						goto next
					}
					for !(b.PreDec(&end).IsRefcounted()) {
						if zv == end {
							goto next
						}
					}
				} else {
					GC_REF_SET_COLOR(ht, GC_GREY)
				}
				for zv != end {
					if zv.IsRefcounted() {
						ref = zv.GetCounted()
						ref.DelRefcount()
						if !(GC_REF_CHECK_COLOR(ref, GC_GREY)) {
							GC_REF_SET_COLOR(ref, GC_GREY)
							GC_STACK_PUSH(ref)
						}
					}
					zv++
				}
				if ht == nil {
					ref = zv.GetCounted()
					ref.DelRefcount()
					if !(GC_REF_CHECK_COLOR(ref, GC_GREY)) {
						GC_REF_SET_COLOR(ref, GC_GREY)
						continue
					}
					goto next
				}
			} else {
				goto next
			}
		} else if ref.GetGcType() == IS_ARRAY {
			if (*ZendArray)(ref) == EG__().GetSymbolTable() {
				GC_REF_SET_BLACK(ref)
				goto next
			} else {
				ht = (*ZendArray)(ref)
			}
		} else if ref.GetGcType() == IS_REFERENCE {
			if (*ZendReference)(ref).GetVal().IsRefcounted() {
				ref = (*ZendReference)(ref).GetVal().GetCounted()
				ref.DelRefcount()
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
			zv = end.GetVal()
			if zv.IsIndirect() {
				zv = zv.GetZv()
			}
			if zv.IsRefcounted() {
				break
			}
			if p == end {
				goto next
			}
		}
		for p != end {
			zv = p.GetVal()
			if zv.IsIndirect() {
				zv = zv.GetZv()
			}
			if zv.IsRefcounted() {
				ref = zv.GetCounted()
				ref.DelRefcount()
				if !(GC_REF_CHECK_COLOR(ref, GC_GREY)) {
					GC_REF_SET_COLOR(ref, GC_GREY)
					GC_STACK_PUSH(ref)
				}
			}
			p++
		}
		zv = p.GetVal()
		if zv.IsIndirect() {
			zv = zv.GetZv()
		}
		ref = zv.GetCounted()
		ref.DelRefcount()
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
	if GC_G__().GetNumRoots()+GC_FIRST_ROOT != GC_G__().GetFirstUnused() {
		if GC_G__().GetNumRoots() != 0 {
			var free *GcRootBuffer = GC_IDX2PTR(GC_FIRST_ROOT)
			var scan *GcRootBuffer = GC_IDX2PTR(GC_G__().GetFirstUnused() - 1)
			var end *GcRootBuffer = GC_IDX2PTR(GC_G__().GetNumRoots())
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
		GC_G__().SetUnused(GC_INVALID)
		GC_G__().SetFirstUnused(GC_G__().GetNumRoots() + GC_FIRST_ROOT)
	}
}
func GcMarkRoots(stack *GcStack) {
	var current *GcRootBuffer
	var last *GcRootBuffer
	GcCompact()
	current = GC_IDX2PTR(GC_FIRST_ROOT)
	last = GC_IDX2PTR(GC_G__().GetFirstUnused())
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
		if ref.GetRefcount() > 0 {
			if !(GC_REF_CHECK_COLOR(ref, GC_BLACK)) {
				GC_REF_SET_BLACK(ref)
				if !(_stack.next) {
					GcStackNext(_stack)
				}

				/* Split stack and reuse the tail */

				_stack.next.prev = nil
				GcScanBlack(ref, _stack.next)
				_stack.next.prev = _stack
			}
		} else {
			if ref.GetGcType() == IS_OBJECT {
				var obj *ZendObject = (*ZendObject)(ref)
				if (ref.GetGcFlags() & IS_OBJ_FREE_CALLED) == 0 {
					var n int
					var zv *Zval
					var end *Zval
					var tmp Zval
					ZVAL_OBJ(&tmp, obj)
					ht = obj.GetHandlers().GetGetGc()(&tmp, &zv, &n)
					end = zv + n
					if ht == nil || !(GC_REF_CHECK_COLOR(ht, GC_GREY)) {
						ht = nil
						if n == 0 {
							goto next
						}
						for !(b.PreDec(&end).IsRefcounted()) {
							if zv == end {
								goto next
							}
						}
					} else {
						GC_REF_SET_COLOR(ht, GC_WHITE)
					}
					for zv != end {
						if zv.IsRefcounted() {
							ref = zv.GetCounted()
							if GC_REF_CHECK_COLOR(ref, GC_GREY) {
								GC_REF_SET_COLOR(ref, GC_WHITE)
								GC_STACK_PUSH(ref)
							}
						}
						zv++
					}
					if ht == nil {
						ref = zv.GetCounted()
						if GC_REF_CHECK_COLOR(ref, GC_GREY) {
							GC_REF_SET_COLOR(ref, GC_WHITE)
							goto tail_call
						}
						goto next
					}
				} else {
					goto next
				}
			} else if ref.GetGcType() == IS_ARRAY {
				if (*ZendArray)(ref == EG__().GetSymbolTable()) != nil {
					GC_REF_SET_BLACK(ref)
					goto next
				} else {
					ht = (*ZendArray)(ref)
				}
			} else if ref.GetGcType() == IS_REFERENCE {
				if (*ZendReference)(ref).GetVal().IsRefcounted() {
					ref = (*ZendReference)(ref).GetVal().GetCounted()
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
				zv = end.GetVal()
				if zv.IsIndirect() {
					zv = zv.GetZv()
				}
				if zv.IsRefcounted() {
					break
				}
				if p == end {
					goto next
				}
			}
			for p != end {
				zv = p.GetVal()
				if zv.IsIndirect() {
					zv = zv.GetZv()
				}
				if zv.IsRefcounted() {
					ref = zv.GetCounted()
					if GC_REF_CHECK_COLOR(ref, GC_GREY) {
						GC_REF_SET_COLOR(ref, GC_WHITE)
						GC_STACK_PUSH(ref)
					}
				}
				p++
			}
			zv = p.GetVal()
			if zv.IsIndirect() {
				zv = zv.GetZv()
			}
			ref = zv.GetCounted()
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
	var last *GcRootBuffer = GC_IDX2PTR(GC_G__().GetFirstUnused())
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
		if !(GC_HAS_NEXT_UNUSED()) {
			return
		}
		idx = GC_FETCH_NEXT_UNUSED()
	}
	buf = GC_IDX2PTR(idx)
	buf.SetRef(GC_MAKE_GARBAGE(ref))
	idx = GcCompress(idx)
	GC_REF_SET_INFO(ref, idx|GC_BLACK)
	GC_G__().GetNumRoots()++
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

		if ref.GetGcType() != IS_REFERENCE {
			count++
		}
		if ref.GetGcType() == IS_OBJECT {
			var obj *ZendObject = (*ZendObject)(ref)
			if (ref.GetGcFlags() & IS_OBJ_FREE_CALLED) == 0 {
				var n int
				var zv *Zval
				var end *Zval
				var tmp Zval

				/* optimization: color is GC_BLACK (0) */

				if ref.GetGcInfo() == 0 {
					GcAddGarbage(ref)
				}
				if (obj.GetGcFlags()&IS_OBJ_DESTRUCTOR_CALLED) == 0 && (obj.GetHandlers().GetDtorObj() != ZendObjectsDestroyObject || obj.GetCe().GetDestructor() != nil) {
					*flags |= GC_HAS_DESTRUCTORS
				}
				ZVAL_OBJ(&tmp, obj)
				ht = obj.GetHandlers().GetGetGc()(&tmp, &zv, &n)
				end = zv + n
				if ht == nil || GC_REF_CHECK_COLOR(ht, GC_BLACK) {
					ht = nil
					if n == 0 {
						goto next
					}
					for !(b.PreDec(&end).IsRefcounted()) {
						if zv == end {
							goto next
						}
					}
				} else {
					GC_REF_SET_BLACK(ht)
				}
				for zv != end {
					if zv.IsRefcounted() {
						ref = zv.GetCounted()
						ref.AddRefcount()
						if GC_REF_CHECK_COLOR(ref, GC_WHITE) {
							GC_REF_SET_BLACK(ref)
							GC_STACK_PUSH(ref)
						}
					}
					zv++
				}
				if ht == nil {
					ref = zv.GetCounted()
					ref.AddRefcount()
					if GC_REF_CHECK_COLOR(ref, GC_WHITE) {
						GC_REF_SET_BLACK(ref)
						continue
					}
					goto next
				}
			} else {
				goto next
			}
		} else if ref.GetGcType() == IS_ARRAY {

			/* optimization: color is GC_BLACK (0) */

			if ref.GetGcInfo() == 0 {
				GcAddGarbage(ref)
			}
			ht = (*ZendArray)(ref)
		} else if ref.GetGcType() == IS_REFERENCE {
			if (*ZendReference)(ref).GetVal().IsRefcounted() {
				ref = (*ZendReference)(ref).GetVal().GetCounted()
				ref.AddRefcount()
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
			zv = end.GetVal()
			if zv.IsIndirect() {
				zv = zv.GetZv()
			}
			if zv.IsRefcounted() {
				break
			}
			if p == end {
				goto next
			}
		}
		for p != end {
			zv = p.GetVal()
			if zv.IsIndirect() {
				zv = zv.GetZv()
			}
			if zv.IsRefcounted() {
				ref = zv.GetCounted()
				ref.AddRefcount()
				if GC_REF_CHECK_COLOR(ref, GC_WHITE) {
					GC_REF_SET_BLACK(ref)
					GC_STACK_PUSH(ref)
				}
			}
			p++
		}
		zv = p.GetVal()
		if zv.IsIndirect() {
			zv = zv.GetZv()
		}
		ref = zv.GetCounted()
		ref.AddRefcount()
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
	var last *GcRootBuffer = GC_IDX2PTR(GC_G__().GetFirstUnused())

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
	end = GC_G__().GetFirstUnused()
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
		} else if ref.GetGcType() == IS_REFERENCE {
			if (*ZendReference)(ref).GetVal().IsRefcounted() {
				ref = (*ZendReference)(ref).GetVal().GetCounted()
				goto tail_call
			}
			return count
		} else {
			return count
		}
		if ref.GetGcType() == IS_OBJECT {
			var obj *ZendObject = (*ZendObject)(ref)
			if (ref.GetGcFlags() & IS_OBJ_FREE_CALLED) == 0 {
				var n int
				var zv *Zval
				var end *Zval
				var tmp Zval
				ZVAL_OBJ(&tmp, obj)
				ht = obj.GetHandlers().GetGetGc()(&tmp, &zv, &n)
				end = zv + n
				if ht == nil {
					if n == 0 {
						return count
					}
					for !(b.PreDec(&end).IsRefcounted()) {
						if zv == end {
							return count
						}
					}
				}
				for zv != end {
					if zv.IsRefcounted() {
						ref = zv.GetCounted()
						count += GcRemoveNestedDataFromBuffer(ref, nil)
					}
					zv++
				}
				if ht == nil {
					ref = zv.GetCounted()
					goto tail_call
				}
				if GC_REF_ADDRESS(ht) != 0 && GC_REF_CHECK_COLOR(ht, GC_BLACK) {
					GC_REMOVE_FROM_BUFFER(ht)
				}
			} else {
				return count
			}
		} else if ref.GetGcType() == IS_ARRAY {
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
			zv = end.GetVal()
			if zv.IsIndirect() {
				zv = zv.GetZv()
			}
			if zv.IsRefcounted() {
				break
			}
			if p == end {
				return count
			}
		}
		for p != end {
			zv = p.GetVal()
			if zv.IsIndirect() {
				zv = zv.GetZv()
			}
			if zv.IsRefcounted() {
				ref = zv.GetCounted()
				count += GcRemoveNestedDataFromBuffer(ref, nil)
			}
			p++
		}
		zv = p.GetVal()
		if zv.IsIndirect() {
			zv = zv.GetZv()
		}
		ref = zv.GetCounted()
		goto tail_call
		break
	}
}
func ZendGcCollectCycles() int {
	var count int = 0
	if GC_G__().GetNumRoots() != 0 {
		var current *GcRootBuffer
		var last *GcRootBuffer
		var p *ZendRefcounted
		var gc_flags uint32 = 0
		var idx uint32
		var end uint32
		var stack GcStack
		stack.SetPrev(nil)
		stack.SetNext(nil)
		if GC_G__().GetGcActive() != 0 {
			return 0
		}
		GC_G__().GetGcRuns()++
		GC_G__().SetGcActive(1)
		GcMarkRoots(&stack)
		GcScanRoots(&stack)
		count = GcCollectRoots(&gc_flags, &stack)
		GcStackFree(&stack)
		if GC_G__().GetNumRoots() == 0 {

			/* nothing to free */

			GC_G__().SetGcActive(0)
			return 0
		}
		end = GC_G__().GetFirstUnused()
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
					if p.GetGcType() == IS_OBJECT && (p.GetGcFlags()&IS_OBJ_DESTRUCTOR_CALLED) == 0 {
						var obj *ZendObject = (*ZendObject)(p)
						if obj.GetHandlers().GetDtorObj() != ZendObjectsDestroyObject || obj.GetCe().GetDestructor() != nil {
							current.SetRef(GC_MAKE_DTOR_GARBAGE(obj))
							GC_REF_SET_COLOR(obj, GC_PURPLE)
						} else {
							obj.AddGcFlags(IS_OBJ_DESTRUCTOR_CALLED)
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

					if (p.GetGcFlags() & IS_OBJ_DESTRUCTOR_CALLED) == 0 {
						var obj *ZendObject = (*ZendObject)(p)
						obj.AddGcFlags(IS_OBJ_DESTRUCTOR_CALLED)
						obj.AddRefcount()
						obj.GetHandlers().GetDtorObj()(obj)
						obj.DelRefcount()
					}

					/* Double check that the destructor hasn't been called yet. It could have
					 * already been invoked indirectly by some other destructor. */

				}
				idx++
			}
			if GC_G__().GetGcProtected() != 0 {

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
				if p.GetGcType() == IS_OBJECT {
					var obj *ZendObject = (*ZendObject)(p)
					EG__().GetObjectsStore().GetObjectBuckets()[obj.GetHandle()] = SET_OBJ_INVALID(obj)
					obj.GetGcTypeInfo() = IS_NULL | obj.GetGcTypeInfo() & ^GC_TYPE_MASK

					/* Modify current before calling free_obj (bug #78811: free_obj() can cause the root buffer (with current) to be reallocated.) */

					current.SetRef(GC_MAKE_GARBAGE((*byte)(obj) - obj.GetHandlers().GetOffset()))
					if (obj.GetGcFlags() & IS_OBJ_FREE_CALLED) == 0 {
						obj.AddGcFlags(IS_OBJ_FREE_CALLED)
						obj.AddRefcount()
						obj.GetHandlers().GetFreeObj()(obj)
						obj.DelRefcount()
					}
					ZEND_OBJECTS_STORE_ADD_TO_FREE_LIST(obj.GetHandle())
				} else if p.GetGcType() == IS_ARRAY {
					var arr *ZendArray = (*ZendArray)(p)
					arr.GetGcTypeInfo() = IS_NULL | arr.GetGcTypeInfo() & ^GC_TYPE_MASK

					/* GC may destroy arrays with rc>1. This is valid and safe. */

					arr.Destroy()
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
				GC_G__().GetNumRoots()--
				Efree(p)
			}
			current++
		}
		GC_G__().SetCollected(GC_G__().GetCollected() + count)
		GC_G__().SetGcActive(0)
	}
	GcCompact()
	return count
}
func ZendGcGetStatus(status *ZendGcStatus) {
	status.SetRuns(GC_G__().GetGcRuns())
	status.SetCollected(GC_G__().GetCollected())
	status.SetThreshold(GC_G__().GetGcThreshold())
	status.SetNumRoots(GC_G__().GetNumRoots())
}
