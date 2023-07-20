package slices

func Map[T any, R any](slice []T, mapper func(T) R) []R {
	var result []R
	for _, item := range slice {
		result = append(result, mapper(item))
	}
	return result
}

func Each[T any, R any](slice []T, handler func(T)) []R {
	var result []R
	for _, item := range slice {
		handler(item)
	}
	return result
}
