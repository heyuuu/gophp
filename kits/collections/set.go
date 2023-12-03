package collections

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (set Set[T]) Add(value T) bool {
	if _, exists := set[value]; exists {
		return false
	}
	set[value] = struct{}{}
	return true
}
func (set Set[T]) Set(value T) {
	set[value] = struct{}{}
}
func (set Set[T]) Del(value T) {
	delete(set, value)
}
func (set Set[T]) Exists(value T) bool {
	_, exists := set[value]
	return exists
}
func (set StringSet) Len() int {
	return len(set)
}

type StringSet = Set[string]
