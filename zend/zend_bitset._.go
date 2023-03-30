package zend

import (
	b "github.com/heyuuu/gophp/builtin"
)

type ZendBitset *ZendUlong

const ZEND_BITSET_ELM_SIZE = b.SizeOf("zend_ulong")
