package types

/* Array.flags */
const HASH_FLAG_PACKED = 1 << 2
const HASH_FLAG_HAS_EMPTY_IND = 1 << 5

const HASH_KEY_IS_STRING = 1
const HASH_KEY_IS_LONG = 2
const HASH_KEY_NON_EXISTENT = 3

type MergeCheckerFuncT func(target_ht *Array, source_data *Zval, hash_key *ArrayKey, pParam any) ZendBool

const (
	ArrayApplyKeep = 1 << iota
	ArrayApplyRemove
	ArrayApplyStop
)

type ApplyFuncT func(pDest *Zval) int
type ApplyFuncArgT func(pDest *Zval, argument any) int
type ApplyFuncArgsT func(pDest *Zval, num_args int, args []any, hash_key *ArrayKey) int

const HT_POISONED_PTR *Array = (*Array)(intptr_t - 1)
