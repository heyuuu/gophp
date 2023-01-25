// <<generate>>

package zend

const HASH_KEY_IS_STRING = 1
const HASH_KEY_IS_LONG = 2
const HASH_KEY_NON_EXISTENT = 3
const HASH_UPDATE uint32 = 1 << 0
const HASH_ADD uint32 = 1 << 1
const HASH_UPDATE_INDIRECT = 1 << 2
const HASH_ADD_NEW uint32 = 1 << 3
const HASH_ADD_NEXT = 1 << 4
const HASH_FLAG_CONSISTENCY = 1<<0 | 1<<1
const HASH_FLAG_PACKED = 1 << 2
const HASH_FLAG_UNINITIALIZED uint32 = 1 << 3
const HASH_FLAG_STATIC_KEYS ZendUchar = 1 << 4
const HASH_FLAG_HAS_EMPTY_IND = 1 << 5
const HASH_FLAG_ALLOW_COW_VIOLATION = 1 << 6
const HASH_FLAG_MASK = 0xff

var ZendEmptyArray HashTable

type MergeCheckerFuncT func(target_ht *HashTable, source_data *Zval, hash_key *ZendHashKey, pParam any) ZendBool

const ZEND_HASH_APPLY_KEEP = 0
const ZEND_HASH_APPLY_REMOVE = 1 << 0
const ZEND_HASH_APPLY_STOP = 1 << 1

type ApplyFuncT func(pDest *Zval) int
type ApplyFuncArgT func(pDest *Zval, argument any) int
type ApplyFuncArgsT func(pDest *Zval, num_args int, args ...any, hash_key *ZendHashKey) int

const HT_POISONED_PTR *HashTable = (*HashTable)(intptr_t - 1)

var UninitializedBucket []uint32 = []uint32{HT_INVALID_IDX, HT_INVALID_IDX}
