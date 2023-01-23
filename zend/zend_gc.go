// <<generate>>

package zend

import (
	r "sik/runtime"
	g "sik/runtime/grammar"
)

// Source: <Zend/zend_gc.h>

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
   | Authors: David Wang <planetbeing@gmail.com>                          |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

// #define ZEND_GC_H

// @type ZendGcStatus struct

var GcCollectCycles func() int

/* enable/disable automatic start of GC collection */

/* enable/disable possible root additions */

/* The default implementation of the gc_collect_cycles callback. */

// #define GC_REMOVE_FROM_BUFFER(p) do { zend_refcounted * _p = ( zend_refcounted * ) ( p ) ; if ( GC_TYPE_INFO ( _p ) & GC_INFO_MASK ) { gc_remove_from_buffer ( _p ) ; } } while ( 0 )

// #define GC_MAY_LEAK(ref) ( ( GC_TYPE_INFO ( ref ) & ( GC_INFO_MASK | ( GC_COLLECTABLE << GC_FLAGS_SHIFT ) ) ) == ( GC_COLLECTABLE << GC_FLAGS_SHIFT ) )

func GcCheckPossibleRoot(ref *ZendRefcounted) {
	if ref.GetGc().GetTypeInfo() == 10 {
		var zv *Zval = &((*ZendReference)(ref)).val
		if (zv.GetTypeFlags() & 1 << 1) == 0 {
			return
		}
		ref = zv.GetValue().GetCounted()
	}
	if (ref.GetGc().GetTypeInfo() & (0xfffffc00 | 1<<4<<0)) == 1<<4<<0 {
		GcPossibleRoot(ref)
	}
}

// Source: <Zend/zend_gc.c>

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
   | Authors: David Wang <planetbeing@gmail.com>                          |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

// # include "zend.h"

// # include "zend_API.h"

// #define GC_BENCH       0

// #define ZEND_GC_DEBUG       0

/* GC_INFO layout */

// #define GC_ADDRESS       0x0fffffu

// #define GC_COLOR       0x300000u

// #define GC_BLACK       0x000000u

// #define GC_WHITE       0x100000u

// #define GC_GREY       0x200000u

// #define GC_PURPLE       0x300000u

/* Debug tracing */

// #define GC_TRACE_REF(ref,format)

// #define GC_TRACE_SET_COLOR(ref,new_color)

// #define GC_TRACE(str)

/* GC_INFO access */

// #define GC_REF_ADDRESS(ref) ( ( ( GC_TYPE_INFO ( ref ) ) & ( GC_ADDRESS << GC_INFO_SHIFT ) ) >> GC_INFO_SHIFT )

// #define GC_REF_COLOR(ref) ( ( ( GC_TYPE_INFO ( ref ) ) & ( GC_COLOR << GC_INFO_SHIFT ) ) >> GC_INFO_SHIFT )

// #define GC_REF_CHECK_COLOR(ref,color) ( ( GC_TYPE_INFO ( ref ) & ( GC_COLOR << GC_INFO_SHIFT ) ) == ( ( color ) << GC_INFO_SHIFT ) )

// #define GC_REF_SET_INFO(ref,info) do { GC_TYPE_INFO ( ref ) = ( GC_TYPE_INFO ( ref ) & ( GC_TYPE_MASK | GC_FLAGS_MASK ) ) | ( ( info ) << GC_INFO_SHIFT ) ; } while ( 0 )

// #define GC_REF_SET_COLOR(ref,c) do { GC_TRACE_SET_COLOR ( ref , c ) ; GC_TYPE_INFO ( ref ) = ( GC_TYPE_INFO ( ref ) & ~ ( GC_COLOR << GC_INFO_SHIFT ) ) | ( ( c ) << GC_INFO_SHIFT ) ; } while ( 0 )

// #define GC_REF_SET_BLACK(ref) do { GC_TRACE_SET_COLOR ( ref , GC_BLACK ) ; GC_TYPE_INFO ( ref ) &= ~ ( GC_COLOR << GC_INFO_SHIFT ) ; } while ( 0 )

// #define GC_REF_SET_PURPLE(ref) do { GC_TRACE_SET_COLOR ( ref , GC_PURPLE ) ; GC_TYPE_INFO ( ref ) |= ( GC_COLOR << GC_INFO_SHIFT ) ; } while ( 0 )

/* bit stealing tags for gc_root_buffer.ref */

// #define GC_BITS       0x3

// #define GC_ROOT       0x0

// #define GC_UNUSED       0x1

// #define GC_GARBAGE       0x2

// #define GC_DTOR_GARBAGE       0x3

// #define GC_GET_PTR(ptr) ( ( void * ) ( ( ( uintptr_t ) ( ptr ) ) & ~ GC_BITS ) )

// #define GC_IS_ROOT(ptr) ( ( ( ( uintptr_t ) ( ptr ) ) & GC_BITS ) == GC_ROOT )

// #define GC_IS_UNUSED(ptr) ( ( ( ( uintptr_t ) ( ptr ) ) & GC_BITS ) == GC_UNUSED )

// #define GC_IS_GARBAGE(ptr) ( ( ( ( uintptr_t ) ( ptr ) ) & GC_BITS ) == GC_GARBAGE )

// #define GC_IS_DTOR_GARBAGE(ptr) ( ( ( ( uintptr_t ) ( ptr ) ) & GC_BITS ) == GC_DTOR_GARBAGE )

// #define GC_MAKE_GARBAGE(ptr) ( ( void * ) ( ( ( uintptr_t ) ( ptr ) ) | GC_GARBAGE ) )

// #define GC_MAKE_DTOR_GARBAGE(ptr) ( ( void * ) ( ( ( uintptr_t ) ( ptr ) ) | GC_DTOR_GARBAGE ) )

/* GC address conversion */

// #define GC_IDX2PTR(idx) ( GC_G ( buf ) + ( idx ) )

// #define GC_PTR2IDX(ptr) ( ( ptr ) - GC_G ( buf ) )

// #define GC_IDX2LIST(idx) ( ( void * ) ( uintptr_t ) ( ( ( idx ) * sizeof ( void * ) ) | GC_UNUSED ) )

// #define GC_LIST2IDX(list) ( ( ( uint32_t ) ( uintptr_t ) ( list ) ) / sizeof ( void * ) )

/* GC buffers */

// #define GC_INVALID       0

// #define GC_FIRST_ROOT       1

// #define GC_DEFAULT_BUF_SIZE       ( 16 * 1024 )

// #define GC_BUF_GROW_STEP       ( 128 * 1024 )

// #define GC_MAX_UNCOMPRESSED       ( 512 * 1024 )

// #define GC_MAX_BUF_SIZE       0x40000000

// #define GC_THRESHOLD_DEFAULT       10000

// #define GC_THRESHOLD_STEP       10000

// #define GC_THRESHOLD_MAX       1000000000

// #define GC_THRESHOLD_TRIGGER       100

/* GC flags */

// #define GC_HAS_DESTRUCTORS       ( 1 << 0 )

/* unused buffers */

// #define GC_HAS_UNUSED() ( GC_G ( unused ) != GC_INVALID )

// #define GC_FETCH_UNUSED() gc_fetch_unused ( )

// #define GC_LINK_UNUSED(root) gc_link_unused ( root )

// #define GC_HAS_NEXT_UNUSED_UNDER_THRESHOLD() ( GC_G ( first_unused ) < GC_G ( gc_threshold ) )

// #define GC_HAS_NEXT_UNUSED() ( GC_G ( first_unused ) != GC_G ( buf_size ) )

// #define GC_FETCH_NEXT_UNUSED() gc_fetch_next_unused ( )

// @type GcRootBuffer struct

// @type ZendGcGlobals struct

// #define GC_G(v) ( gc_globals . v )

var GcGlobals ZendGcGlobals

// #define GC_BENCH_INC(counter)

// #define GC_BENCH_DEC(counter)

// #define GC_BENCH_PEAK(peak,counter)

// #define GC_STACK_SEGMENT_SIZE       ( ( ( 4096 - ZEND_MM_OVERHEAD ) / sizeof ( void * ) ) - 2 )

// @type GcStack struct

// #define GC_STACK_DCL(init) gc_stack * _stack = init ; size_t _top = 0 ;

// #define GC_STACK_PUSH(ref) gc_stack_push ( & _stack , & _top , ref ) ;

// #define GC_STACK_POP() gc_stack_pop ( & _stack , & _top )

func GcStackNext(stack *GcStack) *GcStack {
	if stack.GetNext() == nil {
		var segment *GcStack = _emalloc(g.SizeOf("gc_stack"))
		segment.SetPrev(stack)
		segment.SetNext(nil)
		stack.SetNext(segment)
	}
	return stack.GetNext()
}
func GcStackPush(stack **GcStack, top *int, ref *ZendRefcounted) {
	if (*top) == (4096-0)/g.SizeOf("void *")-2 {
		*stack = GcStackNext(*stack)
		*top = 0
	}
	(*stack).GetData()[g.PostInc(&(*top))] = ref
}
func GcStackPop(stack **GcStack, top *int) *ZendRefcounted {
	if (*top) == 0 {
		if (*stack).GetPrev() == nil {
			return nil
		} else {
			*stack = (*stack).GetPrev()
			*top = (4096-0)/g.SizeOf("void *") - 2 - 1
			return (*stack).GetData()[(4096-0)/g.SizeOf("void *")-2-1]
		}
	} else {
		return (*stack).GetData()[g.PreDec(&(*top))]
	}
}
func GcStackFree(stack *GcStack) {
	var p *GcStack = stack.GetNext()
	for p != nil {
		stack = p.GetNext()
		_efree(p)
		p = stack
	}
}
func GcCompress(idx uint32) uint32 {
	if idx < 512*1024 {
		return idx
	}
	return idx%(512*1024) | 512*1024
}
func GcDecompress(ref *ZendRefcounted, idx uint32) *GcRootBuffer {
	var root *GcRootBuffer = GC_G.GetBuf() + idx
	if any(uintptr_t(root.GetRef()) & ^0x3) == ref {
		return root
	}
	for true {
		idx += 512 * 1024
		r.Assert(idx < GC_G.GetFirstUnused())
		root = GC_G.GetBuf() + idx
		if any(uintptr_t(root.GetRef()) & ^0x3) == ref {
			return root
		}
	}
}
func GcFetchUnused() uint32 {
	var idx uint32
	var root *GcRootBuffer
	r.Assert(GC_G.GetUnused() != 0)
	idx = GC_G.GetUnused()
	root = GC_G.GetBuf() + idx
	r.Assert((uintptr_t(root.GetRef()) & 0x3) == 0x1)
	GC_G.SetUnused(uint32(uintptr_t)(root.GetRef()) / g.SizeOf("void *"))
	return idx
}
func GcLinkUnused(root *GcRootBuffer) {
	root.SetRef(any(uintptr_t(GC_G.GetUnused()*g.SizeOf("void *") | 0x1)))
	GC_G.SetUnused(root - GC_G.GetBuf())
}
func GcFetchNextUnused() uint32 {
	var idx uint32
	r.Assert(GC_G.GetFirstUnused() != GC_G.GetBufSize())
	idx = GC_G.GetFirstUnused()
	GC_G.SetFirstUnused(GC_G.GetFirstUnused() + 1)
	return idx
}
func GcRemoveFromRoots(root *GcRootBuffer) {
	GcLinkUnused(root)
	GC_G.GetNumRoots()--
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
	gc_globals.SetUnused(0)
	gc_globals.SetFirstUnused(0)
	gc_globals.SetGcThreshold(0)
	gc_globals.SetBufSize(0)
	gc_globals.SetNumRoots(0)
	gc_globals.SetGcRuns(0)
	gc_globals.SetCollected(0)
}
func GcGlobalsCtor() { GcGlobalsCtorEx(&GcGlobals) }
func GcGlobalsDtor() { RootBufferDtor(&GcGlobals) }
func GcReset() {
	if GC_G.GetBuf() != nil {
		GC_G.SetGcActive(0)
		GC_G.SetGcProtected(0)
		GC_G.SetGcFull(0)
		GC_G.SetUnused(0)
		GC_G.SetFirstUnused(1)
		GC_G.SetNumRoots(0)
		GC_G.SetGcRuns(0)
		GC_G.SetCollected(0)
	}
}
func GcEnable(enable ZendBool) ZendBool {
	var old_enabled ZendBool = GC_G.GetGcEnabled()
	GC_G.SetGcEnabled(enable)
	if enable != 0 && old_enabled == 0 && GC_G.GetBuf() == nil {
		GC_G.SetBuf((*GcRootBuffer)(g.CondF(true, func() any { return __zendMalloc(g.SizeOf("gc_root_buffer") * (16 * 1024)) }, func() any { return _emalloc(g.SizeOf("gc_root_buffer") * (16 * 1024)) })))
		GC_G.GetBuf()[0].SetRef(nil)
		GC_G.SetBufSize(16 * 1024)
		GC_G.SetGcThreshold(10000 + 1)
		GcReset()
	}
	return old_enabled
}
func GcEnabled() ZendBool { return GC_G.GetGcEnabled() }
func GcProtect(protect ZendBool) ZendBool {
	var old_protected ZendBool = GC_G.GetGcProtected()
	GC_G.SetGcProtected(protect)
	return old_protected
}
func GcProtected() ZendBool { return GC_G.GetGcProtected() }
func GcGrowRootBuffer() {
	var new_size int
	if GC_G.GetBufSize() >= 0x40000000 {
		if GC_G.GetGcFull() == 0 {
			ZendError(1<<1, "GC buffer overflow (GC disabled)\n")
			GC_G.SetGcActive(1)
			GC_G.SetGcProtected(1)
			GC_G.SetGcFull(1)
			return
		}
	}
	if GC_G.GetBufSize() < 128*1024 {
		new_size = GC_G.GetBufSize() * 2
	} else {
		new_size = GC_G.GetBufSize() + 128*1024
	}
	if new_size > 0x40000000 {
		new_size = 0x40000000
	}
	GC_G.SetBuf(__zendRealloc(GC_G.GetBuf(), g.SizeOf("gc_root_buffer")*new_size))
	GC_G.SetBufSize(new_size)
}
func GcAdjustThreshold(count int) {
	var new_threshold uint32

	/* TODO Very simple heuristic for dynamic GC buffer resizing:
	 * If there are "too few" collections, increase the collection threshold
	 * by a fixed step */

	if count < 100 {

		/* increase */

		if GC_G.GetGcThreshold() < 1000000000 {
			new_threshold = GC_G.GetGcThreshold() + 10000
			if new_threshold > 1000000000 {
				new_threshold = 1000000000
			}
			if new_threshold > GC_G.GetBufSize() {
				GcGrowRootBuffer()
			}
			if new_threshold <= GC_G.GetBufSize() {
				GC_G.SetGcThreshold(new_threshold)
			}
		}

		/* increase */

	} else if GC_G.GetGcThreshold() > 10000 {
		new_threshold = GC_G.GetGcThreshold() - 10000
		if new_threshold < 10000 {
			new_threshold = 10000
		}
		GC_G.SetGcThreshold(new_threshold)
	}

	/* TODO Very simple heuristic for dynamic GC buffer resizing:
	 * If there are "too few" collections, increase the collection threshold
	 * by a fixed step */
}
func GcPossibleRootWhenFull(ref *ZendRefcounted) {
	var idx uint32
	var newRoot *GcRootBuffer
	r.Assert(ZvalGcType(ref.GetGc().GetTypeInfo()) == 7 || ZvalGcType(ref.GetGc().GetTypeInfo()) == 8)
	r.Assert(ZvalGcInfo(ref.GetGc().GetTypeInfo()) == 0)
	if GC_G.GetGcEnabled() != 0 && GC_G.GetGcActive() == 0 {
		ZendGcAddref(&ref.gc)
		GcAdjustThreshold(GcCollectCycles())
		if ZendGcDelref(&ref.gc) == 0 {
			RcDtorFunc(ref)
			return
		} else if ZvalGcInfo(ref.GetGc().GetTypeInfo()) != 0 {
			return
		}
	}
	if GC_G.GetUnused() != 0 {
		idx = GcFetchUnused()
	} else if GC_G.GetFirstUnused() != GC_G.GetBufSize() {
		idx = GcFetchNextUnused()
	} else {
		GcGrowRootBuffer()
		if GC_G.GetFirstUnused() == GC_G.GetBufSize() {
			return
		}
		idx = GcFetchNextUnused()
	}
	newRoot = GC_G.GetBuf() + idx
	newRoot.SetRef(ref)
	idx = GcCompress(idx)
	ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo()&(0xf|0x3f0) | (idx|0x300000)<<10)
	GC_G.GetNumRoots()++
}
func GcPossibleRoot(ref *ZendRefcounted) {
	var idx uint32
	var newRoot *GcRootBuffer
	if GC_G.GetGcProtected() != 0 {
		return
	}
	if GC_G.GetUnused() != 0 {
		idx = GcFetchUnused()
	} else if GC_G.GetFirstUnused() < GC_G.GetGcThreshold() {
		idx = GcFetchNextUnused()
	} else {
		GcPossibleRootWhenFull(ref)
		return
	}
	r.Assert(ZvalGcType(ref.GetGc().GetTypeInfo()) == 7 || ZvalGcType(ref.GetGc().GetTypeInfo()) == 8)
	r.Assert(ZvalGcInfo(ref.GetGc().GetTypeInfo()) == 0)
	newRoot = GC_G.GetBuf() + idx
	newRoot.SetRef(ref)
	idx = GcCompress(idx)
	ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo()&(0xf|0x3f0) | (idx|0x300000)<<10)
	GC_G.GetNumRoots()++
}
func GcRemoveCompressed(ref *ZendRefcounted, idx uint32) {
	var root *GcRootBuffer = GcDecompress(ref, idx)
	GcRemoveFromRoots(root)
}
func GcRemoveFromBuffer(ref *ZendRefcounted) {
	var root *GcRootBuffer
	var idx uint32 = (ref.GetGc().GetTypeInfo() & 0xfffff << 10) >> 10
	if (ref.GetGc().GetTypeInfo() & 0x300000 << 10) != 0x0<<10 {

	}
	ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo()&(0xf|0x3f0) | 0<<10)

	/* Perform decompression only in case of large buffers */

	if GC_G.GetFirstUnused() >= 512*1024 {
		GcRemoveCompressed(ref, idx)
		return
	}
	r.Assert(idx != 0)
	root = GC_G.GetBuf() + idx
	GcRemoveFromRoots(root)
}
func GcScanBlack(ref *ZendRefcounted, stack *GcStack) {
	var ht *HashTable = nil
	var p *Bucket
	var end *Bucket
	var zv *Zval
	var _stack *GcStack = stack
	var _top int = 0
tail_call:
	if ZvalGcType(ref.GetGc().GetTypeInfo()) == 8 {
		var obj *ZendObject = (*ZendObject)(ref)
		if (ZvalGcFlags(ref.GetGc().GetTypeInfo()) & 1 << 9) == 0 {
			var n int
			var zv *Zval
			var end *Zval
			var tmp Zval
			var __z *Zval = &tmp
			__z.GetValue().SetObj(obj)
			__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
			ht = obj.GetHandlers().GetGetGc()(&tmp, &zv, &n)
			end = zv + n
			if ht == nil || (ht.GetGc().GetTypeInfo()&0x300000<<10) == 0x0<<10 {
				ht = nil
				if n == 0 {
					goto next
				}
				for g.PreDec(&end).u1.v.type_flags == 0 {
					if zv == end {
						goto next
					}
				}
			} else {
				ht.GetGc().SetTypeInfo(ht.GetGc().GetTypeInfo() &^ (0x300000 << 10))
			}
			for zv != end {
				if zv.GetTypeFlags() != 0 {
					ref = zv.GetValue().GetCounted()
					ZendGcAddref(&ref.gc)
					if (ref.GetGc().GetTypeInfo() & 0x300000 << 10) != 0x0<<10 {
						ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo() &^ (0x300000 << 10))
						GcStackPush(&_stack, &_top, ref)
					}
				}
				zv++
			}
			if ht == nil {
				ref = zv.GetValue().GetCounted()
				ZendGcAddref(&ref.gc)
				if (ref.GetGc().GetTypeInfo() & 0x300000 << 10) != 0x0<<10 {
					ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo() &^ (0x300000 << 10))
					goto tail_call
				}
				goto next
			}
		} else {
			goto next
		}
	} else if ZvalGcType(ref.GetGc().GetTypeInfo()) == 7 {
		if (*ZendArray)(ref != &EG.symbol_table) != nil {
			ht = (*ZendArray)(ref)
		} else {
			goto next
		}
	} else if ZvalGcType(ref.GetGc().GetTypeInfo()) == 10 {
		if (*ZendReference)(ref).GetVal().GetTypeFlags() != 0 {
			ref = (*ZendReference)(ref).GetVal().GetValue().GetCounted()
			ZendGcAddref(&ref.gc)
			if (ref.GetGc().GetTypeInfo() & 0x300000 << 10) != 0x0<<10 {
				ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo() &^ (0x300000 << 10))
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
		if zv.GetType() == 13 {
			zv = zv.GetValue().GetZv()
		}
		if zv.GetTypeFlags() != 0 {
			break
		}
		if p == end {
			goto next
		}
	}
	for p != end {
		zv = &p.val
		if zv.GetType() == 13 {
			zv = zv.GetValue().GetZv()
		}
		if zv.GetTypeFlags() != 0 {
			ref = zv.GetValue().GetCounted()
			ZendGcAddref(&ref.gc)
			if (ref.GetGc().GetTypeInfo() & 0x300000 << 10) != 0x0<<10 {
				ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo() &^ (0x300000 << 10))
				GcStackPush(&_stack, &_top, ref)
			}
		}
		p++
	}
	zv = &p.val
	if zv.GetType() == 13 {
		zv = zv.GetValue().GetZv()
	}
	ref = zv.GetValue().GetCounted()
	ZendGcAddref(&ref.gc)
	if (ref.GetGc().GetTypeInfo() & 0x300000 << 10) != 0x0<<10 {
		ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo() &^ (0x300000 << 10))
		goto tail_call
	}
next:
	ref = GcStackPop(&_stack, &_top)
	if ref != nil {
		goto tail_call
	}
}
func GcMarkGrey(ref *ZendRefcounted, stack *GcStack) {
	var ht *HashTable = nil
	var p *Bucket
	var end *Bucket
	var zv *Zval
	var _stack *GcStack = stack
	var _top int = 0
	for {
		if ZvalGcType(ref.GetGc().GetTypeInfo()) == 8 {
			var obj *ZendObject = (*ZendObject)(ref)
			if (ZvalGcFlags(ref.GetGc().GetTypeInfo()) & 1 << 9) == 0 {
				var n int
				var zv *Zval
				var end *Zval
				var tmp Zval
				var __z *Zval = &tmp
				__z.GetValue().SetObj(obj)
				__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
				ht = obj.GetHandlers().GetGetGc()(&tmp, &zv, &n)
				end = zv + n
				if ht == nil || (ht.GetGc().GetTypeInfo()&0x300000<<10) == 0x200000<<10 {
					ht = nil
					if n == 0 {
						goto next
					}
					for g.PreDec(&end).u1.v.type_flags == 0 {
						if zv == end {
							goto next
						}
					}
				} else {
					ht.GetGc().SetTypeInfo(ht.GetGc().GetTypeInfo() & ^(0x300000<<10) | 0x200000<<10)
				}
				for zv != end {
					if zv.GetTypeFlags() != 0 {
						ref = zv.GetValue().GetCounted()
						ZendGcDelref(&ref.gc)
						if (ref.GetGc().GetTypeInfo() & 0x300000 << 10) != 0x200000<<10 {
							ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo() & ^(0x300000<<10) | 0x200000<<10)
							GcStackPush(&_stack, &_top, ref)
						}
					}
					zv++
				}
				if ht == nil {
					ref = zv.GetValue().GetCounted()
					ZendGcDelref(&ref.gc)
					if (ref.GetGc().GetTypeInfo() & 0x300000 << 10) != 0x200000<<10 {
						ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo() & ^(0x300000<<10) | 0x200000<<10)
						continue
					}
					goto next
				}
			} else {
				goto next
			}
		} else if ZvalGcType(ref.GetGc().GetTypeInfo()) == 7 {
			if (*ZendArray)(ref) == &EG.symbol_table {
				ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo() &^ (0x300000 << 10))
				goto next
			} else {
				ht = (*ZendArray)(ref)
			}
		} else if ZvalGcType(ref.GetGc().GetTypeInfo()) == 10 {
			if (*ZendReference)(ref).GetVal().GetTypeFlags() != 0 {
				ref = (*ZendReference)(ref).GetVal().GetValue().GetCounted()
				ZendGcDelref(&ref.gc)
				if (ref.GetGc().GetTypeInfo() & 0x300000 << 10) != 0x200000<<10 {
					ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo() & ^(0x300000<<10) | 0x200000<<10)
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
			if zv.GetType() == 13 {
				zv = zv.GetValue().GetZv()
			}
			if zv.GetTypeFlags() != 0 {
				break
			}
			if p == end {
				goto next
			}
		}
		for p != end {
			zv = &p.val
			if zv.GetType() == 13 {
				zv = zv.GetValue().GetZv()
			}
			if zv.GetTypeFlags() != 0 {
				ref = zv.GetValue().GetCounted()
				ZendGcDelref(&ref.gc)
				if (ref.GetGc().GetTypeInfo() & 0x300000 << 10) != 0x200000<<10 {
					ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo() & ^(0x300000<<10) | 0x200000<<10)
					GcStackPush(&_stack, &_top, ref)
				}
			}
			p++
		}
		zv = &p.val
		if zv.GetType() == 13 {
			zv = zv.GetValue().GetZv()
		}
		ref = zv.GetValue().GetCounted()
		ZendGcDelref(&ref.gc)
		if (ref.GetGc().GetTypeInfo() & 0x300000 << 10) != 0x200000<<10 {
			ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo() & ^(0x300000<<10) | 0x200000<<10)
			continue
		}
	next:
		ref = GcStackPop(&_stack, &_top)
		if ref == nil {
			break
		}
	}
}

/* Two-Finger compaction algorithm */

func GcCompact() {
	if GC_G.GetNumRoots()+1 != GC_G.GetFirstUnused() {
		if GC_G.GetNumRoots() != 0 {
			var free *GcRootBuffer = GC_G.GetBuf() + 1
			var scan *GcRootBuffer = GC_G.GetBuf() + (GC_G.GetFirstUnused() - 1)
			var end *GcRootBuffer = GC_G.GetBuf() + GC_G.GetNumRoots()
			var idx uint32
			var p *ZendRefcounted
			for free < scan {
				for (uintptr_t(free.GetRef()) & 0x3) != 0x1 {
					free++
				}
				for (uintptr_t(scan.GetRef()) & 0x3) == 0x1 {
					scan--
				}
				if scan > free {
					p = scan.GetRef()
					free.SetRef(p)
					p = any(uintptr_t(p) & ^0x3)
					idx = GcCompress(free - GC_G.GetBuf())
					p.GetGc().SetTypeInfo(p.GetGc().GetTypeInfo()&(0xf|0x3f0) | (idx|(p.GetGc().GetTypeInfo()&0x300000<<10)>>10)<<10)
					free++
					scan--
					if scan <= end {
						break
					}
				}
			}
		}
		GC_G.SetUnused(0)
		GC_G.SetFirstUnused(GC_G.GetNumRoots() + 1)
	}
}
func GcMarkRoots(stack *GcStack) {
	var current *GcRootBuffer
	var last *GcRootBuffer
	GcCompact()
	current = GC_G.GetBuf() + 1
	last = GC_G.GetBuf() + GC_G.GetFirstUnused()
	for current != last {
		if (uintptr_t(current.GetRef()) & 0x3) == 0x0 {
			if (current.GetRef().GetGc().GetTypeInfo() & 0x300000 << 10) == 0x300000<<10 {
				current.GetRef().GetGc().SetTypeInfo(current.GetRef().GetGc().GetTypeInfo() & ^(0x300000<<10) | 0x200000<<10)
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
	var _stack *GcStack = stack
	var _top int = 0
tail_call:
	if (ref.GetGc().GetTypeInfo() & 0x300000 << 10) == 0x100000<<10 {
		if ZendGcRefcount(&ref.gc) > 0 {
			if (ref.GetGc().GetTypeInfo() & 0x300000 << 10) != 0x0<<10 {
				ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo() &^ (0x300000 << 10))
				if _stack.GetNext() == nil {
					GcStackNext(_stack)
				}

				/* Split stack and reuse the tail */

				_stack.GetNext().SetPrev(nil)
				GcScanBlack(ref, _stack.GetNext())
				_stack.GetNext().SetPrev(_stack)
			}
		} else {
			if ZvalGcType(ref.GetGc().GetTypeInfo()) == 8 {
				var obj *ZendObject = (*ZendObject)(ref)
				if (ZvalGcFlags(ref.GetGc().GetTypeInfo()) & 1 << 9) == 0 {
					var n int
					var zv *Zval
					var end *Zval
					var tmp Zval
					var __z *Zval = &tmp
					__z.GetValue().SetObj(obj)
					__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
					ht = obj.GetHandlers().GetGetGc()(&tmp, &zv, &n)
					end = zv + n
					if ht == nil || (ht.GetGc().GetTypeInfo()&0x300000<<10) != 0x200000<<10 {
						ht = nil
						if n == 0 {
							goto next
						}
						for g.PreDec(&end).u1.v.type_flags == 0 {
							if zv == end {
								goto next
							}
						}
					} else {
						ht.GetGc().SetTypeInfo(ht.GetGc().GetTypeInfo() & ^(0x300000<<10) | 0x100000<<10)
					}
					for zv != end {
						if zv.GetTypeFlags() != 0 {
							ref = zv.GetValue().GetCounted()
							if (ref.GetGc().GetTypeInfo() & 0x300000 << 10) == 0x200000<<10 {
								ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo() & ^(0x300000<<10) | 0x100000<<10)
								GcStackPush(&_stack, &_top, ref)
							}
						}
						zv++
					}
					if ht == nil {
						ref = zv.GetValue().GetCounted()
						if (ref.GetGc().GetTypeInfo() & 0x300000 << 10) == 0x200000<<10 {
							ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo() & ^(0x300000<<10) | 0x100000<<10)
							goto tail_call
						}
						goto next
					}
				} else {
					goto next
				}
			} else if ZvalGcType(ref.GetGc().GetTypeInfo()) == 7 {
				if (*ZendArray)(ref == &EG.symbol_table) != nil {
					ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo() &^ (0x300000 << 10))
					goto next
				} else {
					ht = (*ZendArray)(ref)
				}
			} else if ZvalGcType(ref.GetGc().GetTypeInfo()) == 10 {
				if (*ZendReference)(ref).GetVal().GetTypeFlags() != 0 {
					ref = (*ZendReference)(ref).GetVal().GetValue().GetCounted()
					if (ref.GetGc().GetTypeInfo() & 0x300000 << 10) == 0x200000<<10 {
						ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo() & ^(0x300000<<10) | 0x100000<<10)
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
				if zv.GetType() == 13 {
					zv = zv.GetValue().GetZv()
				}
				if zv.GetTypeFlags() != 0 {
					break
				}
				if p == end {
					goto next
				}
			}
			for p != end {
				zv = &p.val
				if zv.GetType() == 13 {
					zv = zv.GetValue().GetZv()
				}
				if zv.GetTypeFlags() != 0 {
					ref = zv.GetValue().GetCounted()
					if (ref.GetGc().GetTypeInfo() & 0x300000 << 10) == 0x200000<<10 {
						ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo() & ^(0x300000<<10) | 0x100000<<10)
						GcStackPush(&_stack, &_top, ref)
					}
				}
				p++
			}
			zv = &p.val
			if zv.GetType() == 13 {
				zv = zv.GetValue().GetZv()
			}
			ref = zv.GetValue().GetCounted()
			if (ref.GetGc().GetTypeInfo() & 0x300000 << 10) == 0x200000<<10 {
				ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo() & ^(0x300000<<10) | 0x100000<<10)
				goto tail_call
			}
		}
	}
next:
	ref = GcStackPop(&_stack, &_top)
	if ref != nil {
		goto tail_call
	}
}
func GcScanRoots(stack *GcStack) {
	var current *GcRootBuffer = GC_G.GetBuf() + 1
	var last *GcRootBuffer = GC_G.GetBuf() + GC_G.GetFirstUnused()
	for current != last {
		if (uintptr_t(current.GetRef()) & 0x3) == 0x0 {
			if (current.GetRef().GetGc().GetTypeInfo() & 0x300000 << 10) == 0x200000<<10 {
				current.GetRef().GetGc().SetTypeInfo(current.GetRef().GetGc().GetTypeInfo() & ^(0x300000<<10) | 0x100000<<10)
				GcScan(current.GetRef(), stack)
			}
		}
		current++
	}
}
func GcAddGarbage(ref *ZendRefcounted) {
	var idx uint32
	var buf *GcRootBuffer
	if GC_G.GetUnused() != 0 {
		idx = GcFetchUnused()
	} else if GC_G.GetFirstUnused() != GC_G.GetBufSize() {
		idx = GcFetchNextUnused()
	} else {
		GcGrowRootBuffer()
		if GC_G.GetFirstUnused() == GC_G.GetBufSize() {
			return
		}
		idx = GcFetchNextUnused()
	}
	buf = GC_G.GetBuf() + idx
	buf.SetRef(any(uintptr_t(ref) | 0x2))
	idx = GcCompress(idx)
	ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo()&(0xf|0x3f0) | (idx|0x0)<<10)
	GC_G.GetNumRoots()++
}
func GcCollectWhite(ref *ZendRefcounted, flags *uint32, stack *GcStack) int {
	var count int = 0
	var ht *HashTable = nil
	var p *Bucket
	var end *Bucket
	var zv *Zval
	var _stack *GcStack = stack
	var _top int = 0
	for {

		/* don't count references for compatibility ??? */

		if ZvalGcType(ref.GetGc().GetTypeInfo()) != 10 {
			count++
		}
		if ZvalGcType(ref.GetGc().GetTypeInfo()) == 8 {
			var obj *ZendObject = (*ZendObject)(ref)
			if (ZvalGcFlags(ref.GetGc().GetTypeInfo()) & 1 << 9) == 0 {
				var n int
				var zv *Zval
				var end *Zval
				var tmp Zval

				/* optimization: color is GC_BLACK (0) */

				if ZvalGcInfo(ref.GetGc().GetTypeInfo()) == 0 {
					GcAddGarbage(ref)
				}
				if (ZvalGcFlags(obj.GetGc().GetTypeInfo())&1<<8) == 0 && (obj.GetHandlers().GetDtorObj() != ZendObjectsDestroyObject || obj.GetCe().GetDestructor() != nil) {
					*flags |= 1 << 0
				}
				var __z *Zval = &tmp
				__z.GetValue().SetObj(obj)
				__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
				ht = obj.GetHandlers().GetGetGc()(&tmp, &zv, &n)
				end = zv + n
				if ht == nil || (ht.GetGc().GetTypeInfo()&0x300000<<10) == 0x0<<10 {
					ht = nil
					if n == 0 {
						goto next
					}
					for g.PreDec(&end).u1.v.type_flags == 0 {
						if zv == end {
							goto next
						}
					}
				} else {
					ht.GetGc().SetTypeInfo(ht.GetGc().GetTypeInfo() &^ (0x300000 << 10))
				}
				for zv != end {
					if zv.GetTypeFlags() != 0 {
						ref = zv.GetValue().GetCounted()
						ZendGcAddref(&ref.gc)
						if (ref.GetGc().GetTypeInfo() & 0x300000 << 10) == 0x100000<<10 {
							ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo() &^ (0x300000 << 10))
							GcStackPush(&_stack, &_top, ref)
						}
					}
					zv++
				}
				if ht == nil {
					ref = zv.GetValue().GetCounted()
					ZendGcAddref(&ref.gc)
					if (ref.GetGc().GetTypeInfo() & 0x300000 << 10) == 0x100000<<10 {
						ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo() &^ (0x300000 << 10))
						continue
					}
					goto next
				}
			} else {
				goto next
			}
		} else if ZvalGcType(ref.GetGc().GetTypeInfo()) == 7 {

			/* optimization: color is GC_BLACK (0) */

			if ZvalGcInfo(ref.GetGc().GetTypeInfo()) == 0 {
				GcAddGarbage(ref)
			}
			ht = (*ZendArray)(ref)
		} else if ZvalGcType(ref.GetGc().GetTypeInfo()) == 10 {
			if (*ZendReference)(ref).GetVal().GetTypeFlags() != 0 {
				ref = (*ZendReference)(ref).GetVal().GetValue().GetCounted()
				ZendGcAddref(&ref.gc)
				if (ref.GetGc().GetTypeInfo() & 0x300000 << 10) == 0x100000<<10 {
					ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo() &^ (0x300000 << 10))
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
			if zv.GetType() == 13 {
				zv = zv.GetValue().GetZv()
			}
			if zv.GetTypeFlags() != 0 {
				break
			}
			if p == end {
				goto next
			}
		}
		for p != end {
			zv = &p.val
			if zv.GetType() == 13 {
				zv = zv.GetValue().GetZv()
			}
			if zv.GetTypeFlags() != 0 {
				ref = zv.GetValue().GetCounted()
				ZendGcAddref(&ref.gc)
				if (ref.GetGc().GetTypeInfo() & 0x300000 << 10) == 0x100000<<10 {
					ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo() &^ (0x300000 << 10))
					GcStackPush(&_stack, &_top, ref)
				}
			}
			p++
		}
		zv = &p.val
		if zv.GetType() == 13 {
			zv = zv.GetValue().GetZv()
		}
		ref = zv.GetValue().GetCounted()
		ZendGcAddref(&ref.gc)
		if (ref.GetGc().GetTypeInfo() & 0x300000 << 10) == 0x100000<<10 {
			ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo() &^ (0x300000 << 10))
			continue
		}
	next:
		ref = GcStackPop(&_stack, &_top)
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
	var current *GcRootBuffer = GC_G.GetBuf() + 1
	var last *GcRootBuffer = GC_G.GetBuf() + GC_G.GetFirstUnused()

	/* remove non-garbage from the list */

	for current != last {
		if (uintptr_t(current.GetRef()) & 0x3) == 0x0 {
			if (current.GetRef().GetGc().GetTypeInfo() & 0x300000 << 10) == 0x0<<10 {
				current.GetRef().GetGc().SetTypeInfo(current.GetRef().GetGc().GetTypeInfo()&(0xf|0x3f0) | 0<<10)
				GcRemoveFromRoots(current)
			}
		}
		current++
	}
	GcCompact()

	/* Root buffer might be reallocated during gc_collect_white,
	 * make sure to reload pointers. */

	idx = 1
	end = GC_G.GetFirstUnused()
	for idx != end {
		current = GC_G.GetBuf() + idx
		ref = current.GetRef()
		r.Assert((uintptr_t(ref) & 0x3) == 0x0)
		current.SetRef(any(uintptr_t(ref) | 0x2))
		if (ref.GetGc().GetTypeInfo() & 0x300000 << 10) == 0x100000<<10 {
			ref.GetGc().SetTypeInfo(ref.GetGc().GetTypeInfo() &^ (0x300000 << 10))
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
		} else if (ref.GetGc().GetTypeInfo()&0xfffff<<10)>>10 != 0 && (ref.GetGc().GetTypeInfo()&0x300000<<10) == 0x0<<10 {
			var _p *ZendRefcounted = (*ZendRefcounted)(ref)
			if (_p.GetGc().GetTypeInfo() & 0xfffffc00) != 0 {
				GcRemoveFromBuffer(_p)
			}
			count++
		} else if ZvalGcType(ref.GetGc().GetTypeInfo()) == 10 {
			if (*ZendReference)(ref).GetVal().GetTypeFlags() != 0 {
				ref = (*ZendReference)(ref).GetVal().GetValue().GetCounted()
				goto tail_call
			}
			return count
		} else {
			return count
		}
		if ZvalGcType(ref.GetGc().GetTypeInfo()) == 8 {
			var obj *ZendObject = (*ZendObject)(ref)
			if (ZvalGcFlags(ref.GetGc().GetTypeInfo()) & 1 << 9) == 0 {
				var n int
				var zv *Zval
				var end *Zval
				var tmp Zval
				var __z *Zval = &tmp
				__z.GetValue().SetObj(obj)
				__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
				ht = obj.GetHandlers().GetGetGc()(&tmp, &zv, &n)
				end = zv + n
				if ht == nil {
					if n == 0 {
						return count
					}
					for g.PreDec(&end).u1.v.type_flags == 0 {
						if zv == end {
							return count
						}
					}
				}
				for zv != end {
					if zv.GetTypeFlags() != 0 {
						ref = zv.GetValue().GetCounted()
						count += GcRemoveNestedDataFromBuffer(ref, nil)
					}
					zv++
				}
				if ht == nil {
					ref = zv.GetValue().GetCounted()
					goto tail_call
				}
				if (ht.GetGc().GetTypeInfo()&0xfffff<<10)>>10 != 0 && (ht.GetGc().GetTypeInfo()&0x300000<<10) == 0x0<<10 {
					var _p *ZendRefcounted = (*ZendRefcounted)(ht)
					if (_p.GetGc().GetTypeInfo() & 0xfffffc00) != 0 {
						GcRemoveFromBuffer(_p)
					}
				}
			} else {
				return count
			}
		} else if ZvalGcType(ref.GetGc().GetTypeInfo()) == 7 {
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
			if zv.GetType() == 13 {
				zv = zv.GetValue().GetZv()
			}
			if zv.GetTypeFlags() != 0 {
				break
			}
			if p == end {
				return count
			}
		}
		for p != end {
			zv = &p.val
			if zv.GetType() == 13 {
				zv = zv.GetValue().GetZv()
			}
			if zv.GetTypeFlags() != 0 {
				ref = zv.GetValue().GetCounted()
				count += GcRemoveNestedDataFromBuffer(ref, nil)
			}
			p++
		}
		zv = &p.val
		if zv.GetType() == 13 {
			zv = zv.GetValue().GetZv()
		}
		ref = zv.GetValue().GetCounted()
		goto tail_call
		break
	}
}
func ZendGcCollectCycles() int {
	var count int = 0
	if GC_G.GetNumRoots() != 0 {
		var current *GcRootBuffer
		var last *GcRootBuffer
		var p *ZendRefcounted
		var gc_flags uint32 = 0
		var idx uint32
		var end uint32
		var stack GcStack
		stack.SetPrev(nil)
		stack.SetNext(nil)
		if GC_G.GetGcActive() != 0 {
			return 0
		}
		GC_G.GetGcRuns()++
		GC_G.SetGcActive(1)
		GcMarkRoots(&stack)
		GcScanRoots(&stack)
		count = GcCollectRoots(&gc_flags, &stack)
		GcStackFree(&stack)
		if GC_G.GetNumRoots() == 0 {

			/* nothing to free */

			GC_G.SetGcActive(0)
			return 0
		}
		end = GC_G.GetFirstUnused()
		if (gc_flags & 1 << 0) != 0 {

			/* During a destructor call, new externally visible references to nested data may
			 * be introduced. These references can be introduced in a way that does not
			 * modify any refcounts, so we have no real way to detect this situation
			 * short of rerunning full GC tracing. What we do instead is to only run
			 * destructors at this point, and leave the actual freeing of the objects
			 * until the next GC run. */

			idx = 1
			current = GC_G.GetBuf() + 1
			for idx != end {
				if (uintptr_t(current.GetRef()) & 0x3) == 0x2 {
					p = any(uintptr_t(current.GetRef()) & ^0x3)
					if ZvalGcType(p.GetGc().GetTypeInfo()) == 8 && (ZvalGcFlags(p.GetGc().GetTypeInfo())&1<<8) == 0 {
						var obj *ZendObject = (*ZendObject)(p)
						if obj.GetHandlers().GetDtorObj() != ZendObjectsDestroyObject || obj.GetCe().GetDestructor() != nil {
							current.SetRef(any(uintptr_t(obj) | 0x3))
							obj.GetGc().SetTypeInfo(obj.GetGc().GetTypeInfo() & ^(0x300000<<10) | 0x300000<<10)
						} else {
							obj.GetGc().SetTypeInfo(obj.GetGc().GetTypeInfo() | 1<<8<<0)
						}
					}
				}
				current++
				idx++
			}

			/* Remove nested data for objects on which a destructor will be called.
			 * This will not remove the objects themselves, as they have been colored
			 * purple. */

			idx = 1
			current = GC_G.GetBuf() + 1
			for idx != end {
				if (uintptr_t(current.GetRef()) & 0x3) == 0x3 {
					p = any(uintptr_t(current.GetRef()) & ^0x3)
					count -= GcRemoveNestedDataFromBuffer(p, current)
				}
				current++
				idx++
			}

			/* Actually call destructors.
			 *
			 * The root buffer might be reallocated during destructors calls,
			 * make sure to reload pointers as necessary. */

			idx = 1
			for idx != end {
				current = GC_G.GetBuf() + idx
				if (uintptr_t(current.GetRef()) & 0x3) == 0x3 {
					p = any(uintptr_t(current.GetRef()) & ^0x3)

					/* Mark this is as a normal root for the next GC run,
					 * it's no longer garbage for this run. */

					current.SetRef(p)

					/* Double check that the destructor hasn't been called yet. It could have
					 * already been invoked indirectly by some other destructor. */

					if (ZvalGcFlags(p.GetGc().GetTypeInfo()) & 1 << 8) == 0 {
						var obj *ZendObject = (*ZendObject)(p)
						obj.GetGc().SetTypeInfo(obj.GetGc().GetTypeInfo() | 1<<8<<0)
						ZendGcAddref(&obj.gc)
						obj.GetHandlers().GetDtorObj()(obj)
						ZendGcDelref(&obj.gc)
					}

					/* Double check that the destructor hasn't been called yet. It could have
					 * already been invoked indirectly by some other destructor. */

				}
				idx++
			}
			if GC_G.GetGcProtected() != 0 {

				/* something went wrong */

				return 0

				/* something went wrong */

			}
		}

		/* Destroy zvals. The root buffer may be reallocated. */

		idx = 1
		for idx != end {
			current = GC_G.GetBuf() + idx
			if (uintptr_t(current.GetRef()) & 0x3) == 0x2 {
				p = any(uintptr_t(current.GetRef()) & ^0x3)
				if ZvalGcType(p.GetGc().GetTypeInfo()) == 8 {
					var obj *ZendObject = (*ZendObject)(p)
					EG.GetObjectsStore().GetObjectBuckets()[obj.GetHandle()] = (*ZendObject)(zend_uintptr_t(obj) | 1<<0)
					obj.GetGc().SetTypeInfo(1 | obj.GetGc().GetTypeInfo() & ^0xf)

					/* Modify current before calling free_obj (bug #78811: free_obj() can cause the root buffer (with current) to be reallocated.) */

					current.SetRef(any(uintptr_t((*byte)(obj)-obj.GetHandlers().GetOffset()) | 0x2))
					if (ZvalGcFlags(obj.GetGc().GetTypeInfo()) & 1 << 9) == 0 {
						obj.GetGc().SetTypeInfo(obj.GetGc().GetTypeInfo() | 1<<9<<0)
						ZendGcAddref(&obj.gc)
						obj.GetHandlers().GetFreeObj()(obj)
						ZendGcDelref(&obj.gc)
					}
					EG.GetObjectsStore().GetObjectBuckets()[obj.GetHandle()] = (*ZendObject)(zend_uintptr_t(EG.GetObjectsStore().GetFreeListHead())<<1 | 1<<0)
					EG.GetObjectsStore().SetFreeListHead(obj.GetHandle())
				} else if ZvalGcType(p.GetGc().GetTypeInfo()) == 7 {
					var arr *ZendArray = (*ZendArray)(p)
					arr.GetGc().SetTypeInfo(1 | arr.GetGc().GetTypeInfo() & ^0xf)

					/* GC may destroy arrays with rc>1. This is valid and safe. */

					ZendHashDestroy(arr)
				}
			}
			idx++
		}

		/* Free objects */

		current = GC_G.GetBuf() + 1
		last = GC_G.GetBuf() + end
		for current != last {
			if (uintptr_t(current.GetRef()) & 0x3) == 0x2 {
				p = any(uintptr_t(current.GetRef()) & ^0x3)
				GcLinkUnused(current)
				GC_G.GetNumRoots()--
				_efree(p)
			}
			current++
		}
		GC_G.SetCollected(GC_G.GetCollected() + count)
		GC_G.SetGcActive(0)
	}
	GcCompact()
	return count
}
func ZendGcGetStatus(status *ZendGcStatus) {
	status.SetRuns(GC_G.GetGcRuns())
	status.SetCollected(GC_G.GetCollected())
	status.SetThreshold(GC_G.GetGcThreshold())
	status.SetNumRoots(GC_G.GetNumRoots())
}
