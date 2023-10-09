package operator

import "github.com/heyuuu/gophp/php/types"

type ZvalTypePair = uint

// inline
func typePair(v1, v2 *types.Zval) ZvalTypePair {
	return ZvalTypePair(uint(v1.Type())<<8 | uint(v2.Type()))
}

const (
	IsNullNull     ZvalTypePair = uint(types.IsNull)<<8 | uint(types.IsNull)
	IsNullFalse                 = uint(types.IsNull)<<8 | uint(types.IsFalse)
	IsNullTrue                  = uint(types.IsNull)<<8 | uint(types.IsTrue)
	IsNullLong                  = uint(types.IsNull)<<8 | uint(types.IsLong)
	IsNullDouble                = uint(types.IsNull)<<8 | uint(types.IsDouble)
	IsNullString                = uint(types.IsNull)<<8 | uint(types.IsString)
	IsNullArray                 = uint(types.IsNull)<<8 | uint(types.IsArray)
	IsFalseNull                 = uint(types.IsFalse)<<8 | uint(types.IsNull)
	IsFalseFalse                = uint(types.IsFalse)<<8 | uint(types.IsFalse)
	IsFalseTrue                 = uint(types.IsFalse)<<8 | uint(types.IsTrue)
	IsFalseLong                 = uint(types.IsFalse)<<8 | uint(types.IsLong)
	IsFalseDouble               = uint(types.IsFalse)<<8 | uint(types.IsDouble)
	IsFalseString               = uint(types.IsFalse)<<8 | uint(types.IsString)
	IsFalseArray                = uint(types.IsFalse)<<8 | uint(types.IsArray)
	IsTrueNull                  = uint(types.IsTrue)<<8 | uint(types.IsNull)
	IsTrueFalse                 = uint(types.IsTrue)<<8 | uint(types.IsFalse)
	IsTrueTrue                  = uint(types.IsTrue)<<8 | uint(types.IsTrue)
	IsTrueLong                  = uint(types.IsTrue)<<8 | uint(types.IsLong)
	IsTrueDouble                = uint(types.IsTrue)<<8 | uint(types.IsDouble)
	IsTrueString                = uint(types.IsTrue)<<8 | uint(types.IsString)
	IsTrueArray                 = uint(types.IsTrue)<<8 | uint(types.IsArray)
	IsLongNull                  = uint(types.IsLong)<<8 | uint(types.IsNull)
	IsLongFalse                 = uint(types.IsLong)<<8 | uint(types.IsFalse)
	IsLongTrue                  = uint(types.IsLong)<<8 | uint(types.IsTrue)
	IsLongLong                  = uint(types.IsLong)<<8 | uint(types.IsLong)
	IsLongDouble                = uint(types.IsLong)<<8 | uint(types.IsDouble)
	IsLongString                = uint(types.IsLong)<<8 | uint(types.IsString)
	IsLongArray                 = uint(types.IsLong)<<8 | uint(types.IsArray)
	IsDoubleNull                = uint(types.IsDouble)<<8 | uint(types.IsNull)
	IsDoubleFalse               = uint(types.IsDouble)<<8 | uint(types.IsFalse)
	IsDoubleTrue                = uint(types.IsDouble)<<8 | uint(types.IsTrue)
	IsDoubleLong                = uint(types.IsDouble)<<8 | uint(types.IsLong)
	IsDoubleDouble              = uint(types.IsDouble)<<8 | uint(types.IsDouble)
	IsDoubleString              = uint(types.IsDouble)<<8 | uint(types.IsString)
	IsDoubleArray               = uint(types.IsDouble)<<8 | uint(types.IsArray)
	IsStringNull                = uint(types.IsString)<<8 | uint(types.IsNull)
	IsStringFalse               = uint(types.IsString)<<8 | uint(types.IsFalse)
	IsStringTrue                = uint(types.IsString)<<8 | uint(types.IsTrue)
	IsStringLong                = uint(types.IsString)<<8 | uint(types.IsLong)
	IsStringDouble              = uint(types.IsString)<<8 | uint(types.IsDouble)
	IsStringString              = uint(types.IsString)<<8 | uint(types.IsString)
	IsStringArray               = uint(types.IsString)<<8 | uint(types.IsArray)
	IsArrayNull                 = uint(types.IsArray)<<8 | uint(types.IsNull)
	IsArrayFalse                = uint(types.IsArray)<<8 | uint(types.IsFalse)
	IsArrayTrue                 = uint(types.IsArray)<<8 | uint(types.IsTrue)
	IsArrayLong                 = uint(types.IsArray)<<8 | uint(types.IsLong)
	IsArrayDouble               = uint(types.IsArray)<<8 | uint(types.IsDouble)
	IsArrayString               = uint(types.IsArray)<<8 | uint(types.IsString)
	IsArrayArray                = uint(types.IsArray)<<8 | uint(types.IsArray)
)

// fast functions
var Long = types.NewZvalLong
var Double = types.NewZvalDouble

// internal functions
func sign(i int) int {
	if i > 0 {
		return 1
	}
	return 0
}
