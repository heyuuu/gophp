package builtin

type Iterator[K any, V any] interface {
	Key() K
	Current() V
	Valid() bool
	Next()
}
