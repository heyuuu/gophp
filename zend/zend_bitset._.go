package zend

import (
	b "sik/builtin"
)

type ZendBitset *ZendUlong

const ZEND_BITSET_ELM_SIZE = b.SizeOf("zend_ulong")
