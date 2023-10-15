package zend

type LlistApplyWithArgFuncT func(data any, arg any)
type ZendLlistPosition[T any] *ZendLlistElement[T]

/**
 * ZendLlistElement
 */
type ZendLlistElement[T any] struct {
	next *ZendLlistElement[T]
	prev *ZendLlistElement[T]
	data T
}

func NewZendLlistElement[T any](data T) *ZendLlistElement[T] {
	return &ZendLlistElement[T]{
		prev: nil,
		next: nil,
		data: data,
	}
}

func (elem *ZendLlistElement[T]) Next() *ZendLlistElement[T] { return elem.next }
func (elem *ZendLlistElement[T]) Prev() *ZendLlistElement[T] { return elem.prev }
func (elem *ZendLlistElement[T]) Data() T                    { return elem.data }

func (elem *ZendLlistElement[T]) SetNext(value *ZendLlistElement[T]) { elem.next = value }
func (elem *ZendLlistElement[T]) SetPrev(value *ZendLlistElement[T]) { elem.prev = value }

/**
 * ZendLlist
 */
type ZendLlist[T any] struct {
	head        *ZendLlistElement[T]
	tail        *ZendLlistElement[T]
	count       int
	dtor        func(T)
	traversePtr ZendLlistPosition[T]
}

func (l *ZendLlist[T]) Init()               { *l = ZendLlist[T]{} }
func (l *ZendLlist[T]) InitEx(dtor func(T)) { *l = ZendLlist[T]{dtor: dtor} }

func (l *ZendLlist[T]) AddElement(element T) {
	node := NewZendLlistElement(element)

	node.prev = l.tail
	node.next = nil
	if l.tail != nil {
		l.tail.next = node
		l.tail = node
	} else {
		l.head = node
		l.tail = node
	}

	l.count++
}

func (l *ZendLlist[T]) PrependElement(element T) {
	node := NewZendLlistElement(element)

	node.prev = nil
	node.next = l.head
	if l.head != nil {
		l.head.prev = node
		l.head = node
	} else {
		l.head = node
		l.tail = node
	}

	l.count++
}

func (l *ZendLlist[T]) DelElement(curr *ZendLlistElement[T]) {
	if curr == nil {
		return
	}
	if curr.prev != nil {
		curr.prev.next = curr.next
	} else {
		l.head = curr.next
	}
	if curr.next != nil {
		curr.next.prev = curr.prev
	} else {
		l.tail = curr.prev
	}
	if l.dtor != nil {
		l.dtor(curr.data)
	}
	l.count--
}

func (l *ZendLlist[T]) Filter(handler func(T) bool) {
	for curr := l.head; curr != nil; {
		next := curr.next
		if !handler(curr.data) {
			l.DelElement(curr)
		}
		curr = next
	}
}

func (l *ZendLlist[T]) DelElementByData(elementData T, compare func(T, T) int) {
	current := l.head
	for current != nil {
		if compare(current.data, elementData) != 0 {
			l.DelElement(current)
			break
		}
		current = current.next
	}
}

func (l *ZendLlist[T]) Clean() {
	current := l.head
	for current != nil {
		next := current.next
		if l.dtor != nil {
			l.dtor(current.data)
		}
		current = next
	}
	l.count = 0
	l.tail = nil
	l.head = nil
}

func (l *ZendLlist[T]) RemoveTail() {
	oldTail := l.tail
	if oldTail == nil {
		return
	}

	if oldTail.prev != nil {
		l.tail = oldTail.prev
		l.tail.next = nil
	} else {
		l.head = nil
		l.tail = nil
	}
	l.count--
	if l.dtor != nil {
		l.dtor(oldTail.data)
	}
}

func (l *ZendLlist[T]) CopyFrom(src *ZendLlist[T]) {
	l.InitEx(src.dtor)

	ptr := src.head
	for ptr != nil {
		l.AddElement(ptr.data)
		ptr = ptr.next
	}
}

func (l *ZendLlist[T]) ApplyWithDel(f func(data any) int) {
	element := l.head
	for element != nil {
		next := element.next
		if f(element.data) != 0 {
			l.DelElement(element)
		}
		element = next
	}
}

func (l *ZendLlist[T]) Apply(f func(data any)) {
	for element := l.head; element != nil; element = element.next {
		f(element.data)
	}
}

func (l *ZendLlist[T]) ApplyWithArgument(f func(data any, arg any), arg any) {
	for element := l.head; element != nil; element = element.next {
		f(element.data, arg)
	}
}

func (l *ZendLlist[T]) ElementsData() []any {
	if l.count == 0 {
		return nil
	}

	elements := make([]any, 0, l.count)
	for element := l.head; element != nil; element = element.next {
		elements = append(elements, element.data)
	}
	return elements
}

func (l *ZendLlist[T]) GetFirstEx(pos *ZendLlistPosition[T]) T {
	var current = pos
	if pos == nil {
		current = &l.traversePtr
	}

	*current = l.head
	if (*current) != nil {
		return (*current).data
	} else {
		return nil
	}
}
func (l *ZendLlist[T]) GetLastEx(pos *ZendLlistPosition[T]) any {
	var current = pos
	if pos == nil {
		current = &l.traversePtr
	}

	*current = l.tail
	if (*current) != nil {
		return (*current).data
	} else {
		return nil
	}
}
func (l *ZendLlist[T]) GetNextEx(pos *ZendLlistPosition[T]) any {
	var current = pos
	if pos == nil {
		current = &l.traversePtr
	}
	if (*current) != nil {
		*current = (*current).next
		if (*current) != nil {
			return (*current).data
		}
	}
	return nil
}

func (l *ZendLlist[T]) GetHead() *ZendLlistElement[T] { return l.head }
func (l *ZendLlist[T]) GetCount() int                 { return l.count }

func (l *ZendLlist[T]) SetHead(value *ZendLlistElement[T]) { l.head = value }
func (l *ZendLlist[T]) SetTail(value *ZendLlistElement[T]) { l.tail = value }
