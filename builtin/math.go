package builtin

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}
type number interface {
	integer | ~float32 | ~float64
}

func Min[T integer](a T, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T integer](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

func FixRange[T integer](num T, min T, max T) T {
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

func Compare[T number](a T, b T) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}
