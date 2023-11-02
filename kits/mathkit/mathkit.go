package mathkit

import (
	"math"
)

func IsInf(f float64) bool {
	return f > math.MaxFloat64 || f < -math.MaxFloat64
}

func IsNaN(f float64) bool {
	// same to
	return math.IsNaN(f)
}

func IsFinite(f float64) bool {
	if IsNaN(f) || IsInf(f) {
		return false
	}
	return true
}
