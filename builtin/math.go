package builtin

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
