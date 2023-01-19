// <<generate>>

package zend

import g "sik/runtime/grammar"

// Source: <Zend/zend_arena.h>

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
   | Authors: Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

// #define _ZEND_ARENA_H_

// # include "zend.h"

// @type ZendArena struct
func ZendArenaCreate(size int) *ZendArena {
	var arena *ZendArena = (*ZendArena)(_emalloc(size))
	arena.SetPtr((*byte)(arena + (g.SizeOf("zend_arena") + 8 - 1 & ^(8-1))))
	arena.SetEnd((*byte)(arena + size))
	arena.SetPrev(nil)
	return arena
}
func ZendArenaDestroy(arena *ZendArena) {
	for {
		var prev *ZendArena = arena.GetPrev()
		_efree(arena)
		arena = prev
		if arena == nil {
			break
		}
	}
}
func ZendArenaAlloc(arena_ptr **ZendArena, size int) any {
	var arena *ZendArena = *arena_ptr
	var ptr *byte = arena.GetPtr()
	size = size + 8 - 1 & ^(8-1)
	if size <= size_t(arena.GetEnd()-ptr) {
		arena.SetPtr(ptr + size)
	} else {
		var arena_size int = g.CondF(size+(g.SizeOf("zend_arena")+8 - 1 & ^(8-1)) > size_t(arena.GetEnd()-(*byte)(arena)), func() int { return size + (g.SizeOf("zend_arena") + 8 - 1 & ^(8-1)) }, func() __auto__ { return size_t(arena.GetEnd() - (*byte)(arena)) })
		var new_arena *ZendArena = (*ZendArena)(_emalloc(arena_size))
		ptr = (*byte)(new_arena + (g.SizeOf("zend_arena") + 8 - 1 & ^(8-1)))
		new_arena.SetPtr((*byte)(new_arena + (g.SizeOf("zend_arena") + 8 - 1 & ^(8-1)) + size))
		new_arena.SetEnd((*byte)(new_arena + arena_size))
		new_arena.SetPrev(arena)
		*arena_ptr = new_arena
	}
	return any(ptr)
}
func ZendArenaCalloc(arena_ptr **ZendArena, count int, unit_size int) any {
	var overflow int
	var size int
	var ret any
	size = ZendSafeAddress(unit_size, count, 0, &overflow)
	if overflow != 0 {
		ZendError(1<<0, "Possible integer overflow in zend_arena_calloc() (%zu * %zu)", unit_size, count)
	}
	ret = ZendArenaAlloc(arena_ptr, size)
	memset(ret, 0, size)
	return ret
}
func ZendArenaCheckpoint(arena *ZendArena) any { return arena.GetPtr() }
func ZendArenaRelease(arena_ptr **ZendArena, checkpoint any) {
	var arena *ZendArena = *arena_ptr
	for (*byte)(checkpoint > arena.GetEnd() || (*byte)(checkpoint <= (*byte)(arena)) != nil) != nil {
		var prev *ZendArena = arena.GetPrev()
		_efree(arena)
		arena = prev
		*arena_ptr = arena
	}
	assert((*byte)(checkpoint > (*byte)(arena != nil && (*byte)(checkpoint <= arena.GetEnd()) != nil)) != nil)
	arena.SetPtr((*byte)(checkpoint))
}
func ZendArenaContains(arena *ZendArena, ptr any) ZendBool {
	for arena != nil {
		if (*byte)(ptr > (*byte)(arena != nil && (*byte)(ptr <= arena.GetPtr()) != nil)) != nil {
			return 1
		}
		arena = arena.GetPrev()
	}
	return 0
}
