package types

import b "sik/builtin"

type Number struct {
	isInt bool
	lval  int
	dval  float64
}

func IntNumber(l int) Number       { return Number{isInt: true, lval: l, dval: float64(l)} }
func FloatNumber(d float64) Number { return Number{isInt: true, lval: 0, dval: d} }
func (n Number) Int() int          { return n.lval }
func (n Number) Float() float64    { return n.dval }

func (num1 Number) CompareTo(num2 Number) int {
	if num1.isInt && num2.isInt {
		return b.Compare(num1.lval, num2.lval)
	} else {
		return b.Compare(num1.dval, num2.dval)
	}
}
