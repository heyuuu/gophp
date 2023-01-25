package builtin

type SlicePtr[T any] struct {
	items []T
	index uint
}

func NewSlicePtr[T any](items []T) *SlicePtr[T] {
	return &SlicePtr[T]{items: items, index: 0}
}

func (p *SlicePtr[T]) Curr() *T { return &p.items[p.index] }

func (p *SlicePtr[T]) PreInc() *T {
	p.index++
	return p.Curr()
}

func (p *SlicePtr[T]) PreDec() *T {
	p.index--
	return p.Curr()
}

func (p *SlicePtr[T]) PostInc() *T {
	var ptr = p.Curr()
	p.index++
	return ptr
}

func (p *SlicePtr[T]) PostDec() *T {
	var ptr = p.Curr()
	p.index--
	return ptr
}
