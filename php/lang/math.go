package lang

import "cmp"

func FixRange[T cmp.Ordered](num T, min T, max T) T {
	if min > max {
		min, max = max, min
	}
	if num < min {
		num = min
	} else if num > max {
		num = max
	}
	return num
}
