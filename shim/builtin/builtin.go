package builtin

import "github.com/heyuuu/gophp/shim/cmp"

// @version go1.21
func Max[T cmp.Ordered](x T, y ...T) T {
	for _, other := range y {
		if other > x {
			x = other
		}
	}
	return x
}

// @version go1.21
func Min[T cmp.Ordered](x T, y ...T) T {
	for _, other := range y {
		if other < x {
			x = other
		}
	}
	return x
}
