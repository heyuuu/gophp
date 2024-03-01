package cmpkit

import "cmp"

type number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func Normalize[T number](n T) int {
	return cmp.Compare(n, 0)
}
