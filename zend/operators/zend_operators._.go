package operators

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

const LONG_SIGN_MASK = zend.ZEND_LONG_MIN

const LOWER_CASE = 1
const UPPER_CASE = 2
const NUMERIC = 3

// inline
func TypePair(t1 types.ZvalType, t2 types.ZvalType) uint { return uint(t1)<<4 | uint(t2) }

const (
	TypeLongLong     = uint(types.IS_LONG)<<4 | uint(types.IS_LONG)
	TypeLongDouble   = uint(types.IS_LONG)<<4 | uint(types.IS_DOUBLE)
	TypeDoubleLong   = uint(types.IS_DOUBLE)<<4 | uint(types.IS_LONG)
	TypeDoubleDouble = uint(types.IS_DOUBLE)<<4 | uint(types.IS_DOUBLE)
	TypeArrayArray   = uint(types.IS_ARRAY)<<4 | uint(types.IS_ARRAY)
	TypeNullNull     = uint(types.IS_NULL)<<4 | uint(types.IS_NULL)
	TypeNullFalse    = uint(types.IS_NULL)<<4 | uint(types.IS_FALSE)
	TypeFalseNull    = uint(types.IS_FALSE)<<4 | uint(types.IS_NULL)
	TypeFalseFalse   = uint(types.IS_FALSE)<<4 | uint(types.IS_FALSE)
	TypeTrueTrue     = uint(types.IS_TRUE)<<4 | uint(types.IS_TRUE)
)
