package php

import (
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
)

// fast type and functions
type Val = types.Zval

var False = types.ZvalFalse
var True = types.ZvalTrue
var Bool = types.ZvalBool
var Long = types.ZvalLong
var Double = types.ZvalDouble
var String = types.ZvalString
var Array = types.ZvalArray
var Resource = types.ZvalResource

type ZvalTypePair uint

// inline
func TypePair(v1, v2 Val) ZvalTypePair {
	return ZvalTypePair(v1.Type())<<8 | ZvalTypePair(v2.Type())
}

const (
	IsNullNull     ZvalTypePair = ZvalTypePair(types.IsNull)<<8 | ZvalTypePair(types.IsNull)
	IsNullFalse                 = ZvalTypePair(types.IsNull)<<8 | ZvalTypePair(types.IsFalse)
	IsNullTrue                  = ZvalTypePair(types.IsNull)<<8 | ZvalTypePair(types.IsTrue)
	IsNullLong                  = ZvalTypePair(types.IsNull)<<8 | ZvalTypePair(types.IsLong)
	IsNullDouble                = ZvalTypePair(types.IsNull)<<8 | ZvalTypePair(types.IsDouble)
	IsNullString                = ZvalTypePair(types.IsNull)<<8 | ZvalTypePair(types.IsString)
	IsNullArray                 = ZvalTypePair(types.IsNull)<<8 | ZvalTypePair(types.IsArray)
	IsFalseNull                 = ZvalTypePair(types.IsFalse)<<8 | ZvalTypePair(types.IsNull)
	IsFalseFalse                = ZvalTypePair(types.IsFalse)<<8 | ZvalTypePair(types.IsFalse)
	IsFalseTrue                 = ZvalTypePair(types.IsFalse)<<8 | ZvalTypePair(types.IsTrue)
	IsFalseLong                 = ZvalTypePair(types.IsFalse)<<8 | ZvalTypePair(types.IsLong)
	IsFalseDouble               = ZvalTypePair(types.IsFalse)<<8 | ZvalTypePair(types.IsDouble)
	IsFalseString               = ZvalTypePair(types.IsFalse)<<8 | ZvalTypePair(types.IsString)
	IsFalseArray                = ZvalTypePair(types.IsFalse)<<8 | ZvalTypePair(types.IsArray)
	IsTrueNull                  = ZvalTypePair(types.IsTrue)<<8 | ZvalTypePair(types.IsNull)
	IsTrueFalse                 = ZvalTypePair(types.IsTrue)<<8 | ZvalTypePair(types.IsFalse)
	IsTrueTrue                  = ZvalTypePair(types.IsTrue)<<8 | ZvalTypePair(types.IsTrue)
	IsTrueLong                  = ZvalTypePair(types.IsTrue)<<8 | ZvalTypePair(types.IsLong)
	IsTrueDouble                = ZvalTypePair(types.IsTrue)<<8 | ZvalTypePair(types.IsDouble)
	IsTrueString                = ZvalTypePair(types.IsTrue)<<8 | ZvalTypePair(types.IsString)
	IsTrueArray                 = ZvalTypePair(types.IsTrue)<<8 | ZvalTypePair(types.IsArray)
	IsLongNull                  = ZvalTypePair(types.IsLong)<<8 | ZvalTypePair(types.IsNull)
	IsLongFalse                 = ZvalTypePair(types.IsLong)<<8 | ZvalTypePair(types.IsFalse)
	IsLongTrue                  = ZvalTypePair(types.IsLong)<<8 | ZvalTypePair(types.IsTrue)
	IsLongLong                  = ZvalTypePair(types.IsLong)<<8 | ZvalTypePair(types.IsLong)
	IsLongDouble                = ZvalTypePair(types.IsLong)<<8 | ZvalTypePair(types.IsDouble)
	IsLongString                = ZvalTypePair(types.IsLong)<<8 | ZvalTypePair(types.IsString)
	IsLongArray                 = ZvalTypePair(types.IsLong)<<8 | ZvalTypePair(types.IsArray)
	IsDoubleNull                = ZvalTypePair(types.IsDouble)<<8 | ZvalTypePair(types.IsNull)
	IsDoubleFalse               = ZvalTypePair(types.IsDouble)<<8 | ZvalTypePair(types.IsFalse)
	IsDoubleTrue                = ZvalTypePair(types.IsDouble)<<8 | ZvalTypePair(types.IsTrue)
	IsDoubleLong                = ZvalTypePair(types.IsDouble)<<8 | ZvalTypePair(types.IsLong)
	IsDoubleDouble              = ZvalTypePair(types.IsDouble)<<8 | ZvalTypePair(types.IsDouble)
	IsDoubleString              = ZvalTypePair(types.IsDouble)<<8 | ZvalTypePair(types.IsString)
	IsDoubleArray               = ZvalTypePair(types.IsDouble)<<8 | ZvalTypePair(types.IsArray)
	IsStringNull                = ZvalTypePair(types.IsString)<<8 | ZvalTypePair(types.IsNull)
	IsStringFalse               = ZvalTypePair(types.IsString)<<8 | ZvalTypePair(types.IsFalse)
	IsStringTrue                = ZvalTypePair(types.IsString)<<8 | ZvalTypePair(types.IsTrue)
	IsStringLong                = ZvalTypePair(types.IsString)<<8 | ZvalTypePair(types.IsLong)
	IsStringDouble              = ZvalTypePair(types.IsString)<<8 | ZvalTypePair(types.IsDouble)
	IsStringString              = ZvalTypePair(types.IsString)<<8 | ZvalTypePair(types.IsString)
	IsStringArray               = ZvalTypePair(types.IsString)<<8 | ZvalTypePair(types.IsArray)
	IsArrayNull                 = ZvalTypePair(types.IsArray)<<8 | ZvalTypePair(types.IsNull)
	IsArrayFalse                = ZvalTypePair(types.IsArray)<<8 | ZvalTypePair(types.IsFalse)
	IsArrayTrue                 = ZvalTypePair(types.IsArray)<<8 | ZvalTypePair(types.IsTrue)
	IsArrayLong                 = ZvalTypePair(types.IsArray)<<8 | ZvalTypePair(types.IsLong)
	IsArrayDouble               = ZvalTypePair(types.IsArray)<<8 | ZvalTypePair(types.IsDouble)
	IsArrayString               = ZvalTypePair(types.IsArray)<<8 | ZvalTypePair(types.IsString)
	IsArrayArray                = ZvalTypePair(types.IsArray)<<8 | ZvalTypePair(types.IsArray)
	IsObjectNull                = ZvalTypePair(types.IsObject)<<8 | ZvalTypePair(types.IsNull)
	IsNullObject                = ZvalTypePair(types.IsNull)<<8 | ZvalTypePair(types.IsObject)
)

// helpers
func Assert(cond bool)                   { perr.Assert(cond) }
func AssertEx(cond bool, message string) { perr.AssertEx(cond, message) }

// internal functions
func sign(i int) int {
	if i > 0 {
		return 1
	}
	return 0
}

func fastGetDouble(v Val) float64 {
	if v.IsLong() {
		return float64(v.Long())
	} else if v.IsDouble() {
		return v.Double()
	} else {
		return 0
	}
}
