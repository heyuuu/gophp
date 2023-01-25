// <<generate>>

package zend

import (
	b "sik/builtin"
)

func ZendArenaCreate(size int) *ZendArena {
	var arena *ZendArena = (*ZendArena)(Emalloc(size))
	arena.SetPtr((*byte)(arena + ZEND_MM_ALIGNED_SIZE(b.SizeOf("zend_arena"))))
	arena.SetEnd((*byte)(arena + size))
	arena.SetPrev(nil)
	return arena
}
func ZendArenaDestroy(arena *ZendArena) {
	for {
		var prev *ZendArena = arena.GetPrev()
		Efree(arena)
		arena = prev
		if arena == nil {
			break
		}
	}
}
func ZendArenaAlloc(arena_ptr **ZendArena, size int) any {
	var arena *ZendArena = *arena_ptr
	var ptr *byte = arena.GetPtr()
	size = ZEND_MM_ALIGNED_SIZE(size)
	if size <= size_t(arena.GetEnd()-ptr) {
		arena.SetPtr(ptr + size)
	} else {
		var arena_size int = b.CondF(size+ZEND_MM_ALIGNED_SIZE(b.SizeOf("zend_arena")) > size_t(arena.GetEnd()-(*byte)(arena)), func() int { return size + ZEND_MM_ALIGNED_SIZE(b.SizeOf("zend_arena")) }, func() __auto__ { return size_t(arena.GetEnd() - (*byte)(arena)) })
		var new_arena *ZendArena = (*ZendArena)(Emalloc(arena_size))
		ptr = (*byte)(new_arena + ZEND_MM_ALIGNED_SIZE(b.SizeOf("zend_arena")))
		new_arena.SetPtr((*byte)(new_arena + ZEND_MM_ALIGNED_SIZE(b.SizeOf("zend_arena")) + size))
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
		ZendError(E_ERROR, "Possible integer overflow in zend_arena_calloc() (%zu * %zu)", unit_size, count)
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
		Efree(arena)
		arena = prev
		*arena_ptr = arena
	}
	ZEND_ASSERT((*byte)(checkpoint > (*byte)(arena != nil && (*byte)(checkpoint <= arena.GetEnd()) != nil)) != nil)
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
