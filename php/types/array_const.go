package types

/* Array.flags */
const HASH_FLAG_HAS_EMPTY_IND = 1 << 5

const (
	ArrayApplyKeep = 1 << iota
	ArrayApplyRemove
	ArrayApplyStop
)

type ApplyFuncT func(pDest *Zval) int
type ApplyFuncArgT func(pDest *Zval, argument any) int
type ApplyFuncArgsT func(pDest *Zval, num_args int, args []any, hash_key *ArrayKey) int
