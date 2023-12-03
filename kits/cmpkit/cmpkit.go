package cmpkit

import "github.com/heyuuu/gophp/shim/cmp"

func Normalize[T int | float64](n T) int {
	return cmp.Compare(n, 0)
}
