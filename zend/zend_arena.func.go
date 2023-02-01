// <<generate>>

package zend

import (
	b "sik/builtin"
)

/**
 * todo ZendArena 本用于申请整块内存，然后在此块内存内操作以提高性能、减少内存碎片；在 go 中已不需要此功能，待移除
 */

type ZendArena struct{}

func ZendArenaCreate(size int) *ZendArena {
	// 仅占位，无实际功能
	return &ZendArena{}
}
func ZendArenaDestroy(arena *ZendArena) {
	// 无需操作
}
func ZendArenaAlloc(arena_ptr **ZendArena, size int) any {
	// 直接申请一块内存
	return b.Malloc(size)
}
