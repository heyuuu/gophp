package mathkit

import (
	"github.com/heyuuu/gophp/shim/cmp"
	"math"
)

type number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func Sign[T number](n T) int {
	return cmp.Compare(n, 0)
}

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
