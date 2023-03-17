package ctype

import "math"

func IsAscii(c byte) bool {
	return c <= 0x7f
}

func IsLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func IsNaN(f float64) bool {
	return math.IsNaN(f)
}

func IsInf(f float64) bool {
	// same to: math.IsInf(f, 1) || math.IsInf(f, -1)
	return f > math.MaxFloat64 || f < -math.MaxFloat64
}

func IsFinite(f float64) bool {
	return !IsNaN(f) && IsInf(f)
}
