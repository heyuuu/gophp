package types

import (
	b "github.com/heyuuu/gophp/builtin"
	"math"
)

type Number struct {
	isFloat bool
	iVal    int
	fVal    float64
}

func IntNumber(l int) Number       { return Number{isFloat: false, iVal: l, fVal: 0} }
func FloatNumber(d float64) Number { return Number{isFloat: true, iVal: 0, fVal: d} }
func (n Number) Int() int          { return n.iVal }
func (n Number) Float() float64    { return n.fVal }
func (n Number) IsInt() bool       { return !n.isFloat }
func (n Number) IsFloat() bool     { return n.isFloat }

func (n Number) AsFloat() float64 {
	if n.isFloat {
		return n.fVal
	} else {
		return float64(n.iVal)
	}
}
func (n Number) Floor() int {
	if n.isFloat {
		return int(math.Floor(n.fVal))
	} else {
		return n.iVal
	}
}

func (n Number) CompareTo(n2 Number) int {
	if n.isFloat && n2.isFloat {
		return b.Compare(n.iVal, n2.iVal)
	} else {
		return b.Compare(n.fVal, n2.fVal)
	}
}

func (n Number) IsInf() bool {
	return n.isFloat && (math.IsInf(n.fVal, 1) || math.IsInf(n.fVal, -1))
}
