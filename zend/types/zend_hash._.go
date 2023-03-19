package types

const HASH_KEY_IS_STRING = 1
const HASH_KEY_IS_LONG = 2
const HASH_KEY_NON_EXISTENT = 3

var ZendEmptyArray HashTable

type MergeCheckerFuncT func(target_ht *HashTable, source_data *Zval, hash_key *ZendHashKey, pParam any) ZendBool

const ZEND_HASH_APPLY_KEEP = 0
const ZEND_HASH_APPLY_REMOVE = 1 << 0
const ZEND_HASH_APPLY_STOP = 1 << 1

type ApplyFuncT func(pDest *Zval) int
type ApplyFuncArgT func(pDest *Zval, argument any) int
type ApplyFuncArgsT func(pDest *Zval, num_args int, args []any, hash_key *ZendHashKey) int

const HT_POISONED_PTR *HashTable = (*HashTable)(intptr_t - 1)
