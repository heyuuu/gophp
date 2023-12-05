package operators

import (
	"github.com/heyuuu/gophp/php/types"
)

type ZvalTypePair uint

// inline
func typePair(v1, v2 Val) ZvalTypePair {
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

// fast type and functions
type Val = *types.Zval

var NewVal = types.NewZvalUndef
var Null = types.NewZvalNull
var False = types.NewZvalFalse
var True = types.NewZvalTrue
var Bool = types.NewZvalBool
var Long = types.NewZvalLong
var Double = types.NewZvalDouble
var String = types.NewZvalString
var Array = types.NewZvalArray

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

func opScalarGetNumber(op1, op2 Val) (Val, Val) {
	return opScalarGetNumberEx(op1, op2, false)
}
func opScalarGetNumberEx(op1, op2 Val, silent bool) (Val, Val) {
	if op1 != op2 {
		op1 = ScalarGetNumber(op1, silent)
		op2 = ScalarGetNumber(op2, silent)
	} else {
		op1 = ScalarGetNumber(op1, silent)
		op2 = op1
	}
	return op1, op2
}

func objectCompare(obj *types.Object, op1, op2 Val) (result int, ok bool) {
	// todo object 比较
	return 1, true
}

func getPrecision() int {
	//  php.EG__().GetPrecision()
	return 14
}

func hasException() bool {
	// todo EG__().HasException()
	return false
}

func throwIfExecuting(ce *types.Class, message string) {
	// todo 触发错误或异常信息
}

func normalizeSign[T int | float64](n T) int {
	if n == 0 {
		return 0
	} else if n < 0 {
		return -1
	} else {
		return 1
	}
}
