package zend

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

/**
 * ZendLlist
 */
type ZendLlist[T any] struct {
	head  *ZendLlistElement[T]
	tail  *ZendLlistElement[T]
	count int
}

func (l *ZendLlist[T]) Init()  { *l = ZendLlist[T]{} }
func (l *ZendLlist[T]) Clean() { l.Init() }

func (l *ZendLlist[T]) AddLast(element T) {
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

func (l *ZendLlist[T]) AddFirst(element T) {
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

func (l *ZendLlist[T]) deleteElement(curr *ZendLlistElement[T]) {
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
	l.count--
}

func (l *ZendLlist[T]) EachElement(handler func(element *ZendLlistElement[T])) {
	for curr := l.head; curr != nil; {
		handler(curr)
	}
}

func (l *ZendLlist[T]) Each(handler func(T)) {
	for curr := l.head; curr != nil; {
		handler(curr.data)
	}
}

func (l *ZendLlist[T]) FindFunc(check func(T) bool) (data T, ok bool) {
	for curr := l.head; curr != nil; {
		if check(curr.data) {
			return curr.data, true
		}
	}
	return
}

func (l *ZendLlist[T]) Filter(handler func(T) bool) {
	for curr := l.head; curr != nil; {
		next := curr.next
		if !handler(curr.data) {
			l.deleteElement(curr)
		}
		curr = next
	}
}
