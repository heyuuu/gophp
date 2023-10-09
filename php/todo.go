package php

import "math"

func getPrec() int {
	// todo zend.EG__().GetPrecision()
	return 20
}

func DvalToLval(f float64) int {
	if !isFinite(f) {
		return 0
	} else if !doubleFitsLong(f) {
		return dvalToLvalSlow(f)
	}
	return int(f)
}

func DvalToLvalCap(f float64) int {
	if !isFinite(f) {
		return 0
	} else if !doubleFitsLong(f) {
		if f > 0 {
			return math.MaxInt
		} else {
			return math.MinInt
		}
	}
	return int(f)
}

func doubleFitsLong(f float64) bool {
	return math.MinInt <= f && f < math.MaxInt
}

func isFinite(f float64) bool {
	return f == f && f <= math.MaxFloat64 && f >= -math.MaxFloat64
}

func dvalToLvalSlow(f float64) int {
	dmod := math.Mod(f, 1<<64)
	if f > math.MaxInt {
		dmod -= 1 << 64
	} else if dmod < math.MinInt {
		dmod += 1 << 64
	}
	return int(dmod)
}
