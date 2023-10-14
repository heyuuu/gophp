package slicekit

func Map[T any, R any](slice []T, mapper func(T) R) []R {
	var result []R
	for _, item := range slice {
		result = append(result, mapper(item))
	}
	return result
}

func Each[T any](slice []T, handler func(T)) {
	for _, item := range slice {
		handler(item)
	}
}

func EachReserve[T any](slice []T, handler func(T)) {
	for i := len(slice) - 1; i >= 0; i-- {
		item := slice[i]
		handler(item)
	}
}

func EachEx[T any](slice []T, handler func(T) bool) {
	for _, item := range slice {
		ok := handler(item)
		if !ok {
			return
		}
	}
}

func EachReserveEx[T any](slice []T, handler func(T) bool) {
	for i := len(slice) - 1; i >= 0; i-- {
		item := slice[i]
		ok := handler(item)
		if !ok {
			return
		}
	}
}

func Push[T any](slicePtr *[]T, item T) {
	*slicePtr = append(*slicePtr, item)
}

func Pop[T any](slicePtr *[]T) (T, bool) {
	slice := *slicePtr
	if len(slice) == 0 {
		var tmp T
		return tmp, false
	}

	top := slice[len(slice)-1]
	*slicePtr = slice[:len(slice)-1]
	return top, true
}

func Last[T any](slicePtr *[]T) (T, bool) {
	slice := *slicePtr
	if len(slice) == 0 {
		var tmp T
		return tmp, false
	}

	top := slice[len(slice)-1]
	return top, true
}

func All[T any](slice []T, handler func(T) bool) bool {
	for _, item := range slice {
		if !handler(item) {
			return false
		}
	}
	return true
}

func Any[T any](slice []T, handler func(T) bool) bool {
	for _, item := range slice {
		if handler(item) {
			return true
		}
	}
	return false
}
