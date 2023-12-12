package operators

import (
	"github.com/heyuuu/gophp/kits/mathkit"
	"github.com/heyuuu/gophp/php/lang"
	"strings"
)

// SmartStrCompare: zendi_smart_strcmp
func SmartStrCompare(s1 string, s2 string) int {
	v1, overflow1 := ParseNumberEx(s1)
	v2, overflow2 := ParseNumberEx(s2)
	if v1 == nil || v2 == nil {
		goto stringCmp
	}

	if overflow1 != 0 && overflow1 == overflow2 && v1.Double()-v2.Double() == 0.0 {
		/* both values are integers overflown to the same side, and the
		 * double comparison may have resulted in crucial accuracy lost */
		goto stringCmp
	}
	if v1.IsDouble() || v2.IsDouble() {
		dval1, dval2 := v1.Double(), v2.Double()
		if v1.IsLong() {
			if overflow2 != 0 {
				/* 2nd operand is integer > LONG_MAX (oflow2==1) or < LONG_MIN (-1) */
				return -1 * overflow2
			}
			dval1 = float64(v1.Long())
		} else if v2.IsLong() {
			if overflow1 != 0 {
				return overflow1
			}
			dval2 = float64(v2.Long())
		} else if v1.Double() == v2.Double() && !(mathkit.IsFinite(v1.Double())) {
			/* Both values overflowed and have the same sign,
			 * so a numeric comparison would be inaccurate */
			goto stringCmp
		}
		return lang.Compare(dval1, dval2)
	} else {
		return lang.Compare(v1.Long(), v2.Long())
	}

stringCmp:
	return strings.Compare(s1, s2)
}

// SmartStrEquals: zend_fast_equal_strings
func SmartStrEquals(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	} else if len(s1) > 0 && s1[0] <= '9' && len(s2) > 0 && s2[0] <= '9' {
		return SmartStrCompare(s1, s2) == 0
	} else {
		return false
	}
}
