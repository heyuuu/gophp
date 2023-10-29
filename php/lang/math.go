package lang

import "github.com/heyuuu/gophp/shim/cmp"

func Min[T cmp.Ordered](a T, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T cmp.Ordered](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

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

func Compare[T cmp.Ordered](a T, b T) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}
