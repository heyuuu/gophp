package zend

import (
	b "github.com/heyuuu/gophp/builtin"
)

/**
 * todo ZendArena 本用于申请整块内存，然后在此块内存内操作以提高性能、减少内存碎片；在 go 中已不需要此功能，待移除
 */

type ZendArena struct{}

func ZendArenaAlloc(arena_ptr **ZendArena, size int) any {
	// 直接申请一块内存
	return b.Malloc(size)
}
