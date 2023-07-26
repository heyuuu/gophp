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
	TypeLongLong     = uint(types.IsLong)<<4 | uint(types.IsLong)
	TypeLongDouble   = uint(types.IsLong)<<4 | uint(types.IsDouble)
	TypeDoubleLong   = uint(types.IsDouble)<<4 | uint(types.IsLong)
	TypeDoubleDouble = uint(types.IsDouble)<<4 | uint(types.IsDouble)
	TypeArrayArray   = uint(types.IsArray)<<4 | uint(types.IsArray)
	TypeNullNull     = uint(types.IsNull)<<4 | uint(types.IsNull)
	TypeNullFalse    = uint(types.IsNull)<<4 | uint(types.IsFalse)
	TypeFalseNull    = uint(types.IsFalse)<<4 | uint(types.IsNull)
	TypeFalseFalse   = uint(types.IsFalse)<<4 | uint(types.IsFalse)
	TypeTrueTrue     = uint(types.IsTrue)<<4 | uint(types.IsTrue)
)
