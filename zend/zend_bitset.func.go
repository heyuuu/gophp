package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
)

func ZEND_BITSET_ELM_NUM(n uint32) int { return n >> 6 }
func ZEND_BITSET_BIT_NUM(n uint32) int { return zend_ulong(n) & uint64(0x3f) }
func ZEND_BITSET_ALLOCA(n uint32, use_heap __auto__) ZendBitset {
	return ZendBitset(DoAlloca(n*ZEND_BITSET_ELM_SIZE, use_heap))
}
func ZendBitsetLen(n uint32) uint32 {
	return (n + (b.SizeOf("zend_long")*8 - 1)) / (b.SizeOf("zend_long") * 8)
}
func ZendBitsetIn(set ZendBitset, n uint32) types.ZendBool { return ZEND_BIT_TEST(set, n) }
func ZendBitsetIncl(set ZendBitset, n uint32) {
	set[ZEND_BITSET_ELM_NUM(n)] |= uint64(1) << ZEND_BITSET_BIT_NUM(n)
}
func ZendBitsetClear(set ZendBitset, len_ uint32) {
	memset(set, 0, len_*ZEND_BITSET_ELM_SIZE)
}
