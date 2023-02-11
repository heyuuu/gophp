package builtin

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Size() int     { return len(s.elements) }
func (s *Stack[T]) IsEmpty() bool { return len(s.elements) == 0 }

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}
func (s *Stack[T]) Pop() T {
	if len(s.elements) > 0 {
		element := s.elements[len(s.elements)-1]
		s.elements = s.elements[:len(s.elements)-1]
		return element
	}

	var element T
	return element
}
func (s *Stack[T]) Top() T {
	if len(s.elements) > 0 {
		return s.elements[len(s.elements)-1]
	}

	var element T
	return element
}

func (s *Stack[T]) Clean() {
	s.elements = nil
}

func (s *Stack[T]) Apply(func_ func(T)) {
	s.RangeTopDown(func(elem T) bool {
		func_(elem)
		return false
	})
}

func (s *Stack[T]) ApplyReverse(func_ func(T)) {
	s.RangeBottomUp(func(elem T) bool {
		func_(elem)
		return false
	})
}

func (s *Stack[T]) RangeTopDown(f func(element T) bool) {
	for i := len(s.elements) - 1; i >= 0; i-- {
		element := s.elements[i]
		if f(element) {
			break
		}
	}
}

func (s *Stack[T]) RangeBottomUp(f func(element T) bool) {
	for _, element := range s.elements {
		if f(element) {
			break
		}
	}
}

func (s *Stack[T]) Copy() Stack[T] {
	newElements := make([]T, len(s.elements))
	copy(newElements, s.elements)
	return Stack[T]{elements: newElements}
}
