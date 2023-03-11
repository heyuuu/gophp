// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/core"
)

func ZEND_MM_ALIGNED_SIZE(size int) int {
	return size + core.ZEND_MM_ALIGNMENT - 1&ZEND_MM_ALIGNMENT_MASK
}
func ZEND_MM_ALIGNED_SIZE_EX(size int, alignment int) int {
	return size + (alignment-1) & ^(alignment-1)
}
func EfreeSize(ptr any, size int)                     { b.Free(ptr) }
func Emalloc(size int) any                            { return b.Malloc(size) }
func SafeEmalloc(nmemb int, size int, offset int) any { return b.Malloc(nmemb, size, offset) }
func Efree(ptr any)                                   { b.Free(ptr) }
func Ecalloc(nmemb int, size int) any                 { return b.Calloc(nmemb, size) }
func Erealloc(ptr any, size int) any                  { return b.Realloc(ptr, size) }
func SafeErealloc(ptr any, nmemb int, size int, offset int) any {
	return b.Realloc(ptr, nmemb, size, offset)
}
func Estrdup(s string) *byte                             { return b.Strdup(s) }
func Estrndup(s *byte, length int) *byte                 { return b.Strndup(s, length) }
func Pemalloc(size int, persistent int) any              { return b.Malloc(size) }
func Pefree(ptr any, persistent int)                     { b.Free(ptr) }
func Pecalloc(nmemb int, size int, persistent uint8) any { return b.Calloc(nmemb, size) }
func Perealloc(ptr any, size int, persistent uint8) any {
	return b.Realloc(ptr, size)
}
func SafePerealloc(ptr any, nmemb int, size int, offset int, persistent int) any {
	return b.Realloc(ptr, size)
}
func Pestrdup(s *byte, persistent int) *byte {
	return b.Strdup(s)
}
func Pestrndup(s *byte, length int, persistent int) *byte {
	return b.Strndup(s, length)
}
func ALLOC_HASHTABLE(ht *HashTable) *HashTable { return NewZendArray(0) }
func FREE_HASHTABLE(ht *HashTable)             { b.Free(ht) }
func ZendMmGc(__ZendMmHeap any) int            { return 0 }
func IsZendMm() int                            { return 0 }
func ZendStrndup(s *byte, length int) *byte {
	var p *byte
	if length+1 == 0 {
		ZendErrorNoreturn(E_ERROR, "Possible integer overflow in memory allocation (1 * %zu + 1)", length)
	}
	var str = b.CastStr(s, length)
	return b.CastStrPtr(str)
}
func ZendSetMemoryLimit(memory_limit int) int {
	// notice memory 限制失效
	return SUCCESS
}
func ZendMemoryUsage(real_usage int) int {
	// notice 获取 memory 使用情况失效
	return 0
}
func ZendMemoryPeakUsage(real_usage int) int {
	// notice 获取 memory 使用情况失效
	return 0
}
func ShutdownMemoryManager(silent int, full_shutdown int) {}
func StartMemoryManager()                                 {}
func ZendMmGetHeap() any/* *ZendMmHeap   */ { return nil }
