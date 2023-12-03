package iterator

type Iterator[T any] interface {
	Current() T
	Next()
	Valid() bool
}

type Rewindable interface {
	Rewind()
}

type RewindableIterator[T any] interface {
	Iterator[T]
	Rewindable
}

// slices iterator
type slicesIterator[T any] struct {
	slice []T
	index int
}

func (it *slicesIterator[T]) Current() T {
	if it.index < len(it.slice) {
		return it.slice[it.index]
	}

	var tmp T
	return tmp
}
func (it *slicesIterator[T]) Next()       { it.index++ }
func (it *slicesIterator[T]) Valid() bool { return it.index < len(it.slice) }
func (it *slicesIterator[T]) Rewind()     { it.index = 0 }

func Slices[T any](slice []T) RewindableIterator[T] {
	return &slicesIterator[T]{slice: slice}
}
