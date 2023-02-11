// <<generate>>

package zend

import "sort"

/**
 * ZendLlistElement
 */
type ZendLlistElement struct {
	next *ZendLlistElement
	prev *ZendLlistElement
	data []byte
}

func NewZendLlistElement(data []byte) *ZendLlistElement {
	return &ZendLlistElement{
		prev: nil,
		next: nil,
		data: data,
	}
}

func (this *ZendLlistElement) GetNext() *ZendLlistElement      { return this.next }
func (this *ZendLlistElement) SetNext(value *ZendLlistElement) { this.next = value }
func (this *ZendLlistElement) GetPrev() *ZendLlistElement      { return this.prev }
func (this *ZendLlistElement) SetPrev(value *ZendLlistElement) { this.prev = value }
func (this *ZendLlistElement) GetData() []byte                 { return this.data }

// func (this *ZendLlistElement) SetData(value []byte) { this.data = value }

/**
 * ZendLlist
 */
type ZendLlist struct {
	head         *ZendLlistElement
	tail         *ZendLlistElement
	count        int
	size         int
	dtor         LlistDtorFuncT
	persistent   uint8
	traverse_ptr *ZendLlistElement
}

func NewZendLlist(size int, dtor LlistDtorFuncT, persistent uint8) *ZendLlist {
	l := &ZendLlist{}
	l.head = nil
	l.tail = nil
	l.count = 0
	l.size = size
	l.dtor = dtor
	l.persistent = persistent
	return l
}

func (l *ZendLlist) Init(size int, dtor LlistDtorFuncT, persistent uint8) {
	l.head = nil
	l.tail = nil
	l.count = 0
	l.size = size
	l.dtor = dtor
	l.persistent = persistent
}

func (l *ZendLlist) AddElement(element any) {
	// todo 这里 l.size 的作用需要明确下
	// todo 确认 element 类型
	data := make([]byte, l.size)
	copy(data, element.([]byte))

	node := NewZendLlistElement(data)

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

func (l *ZendLlist) PrependElement(element any) {
	// todo 这里 l.size 的作用需要明确下
	// todo 确认 element 类型
	data := make([]byte, l.size)
	copy(data, element.([]byte))

	node := NewZendLlistElement(data)

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

func (l *ZendLlist) DelElement(current *ZendLlistElement) {
	if current.prev != nil {
		current.prev.next = current.next
	} else {
		l.head = current.next
	}
	if current.next != nil {
		current.next.prev = current.prev
	} else {
		l.tail = current.prev
	}
	if l.dtor != nil {
		l.dtor(current.data)
	}
	if l.dtor != nil {
		l.dtor(current.data)
	}
	l.count--
}

func (l *ZendLlist) DelElementByData(elementData any, compare func(any, any) int) {
	current := l.head
	for current != nil {
		if compare(current.data, elementData) != 0 {
			l.DelElement(current)
			break
		}
		current = current.next
	}
}

func (l *ZendLlist) Destroy() {
	l.Clean()
}

func (l *ZendLlist) Clean() {
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

func (l *ZendLlist) RemoveTail() {
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

func (l *ZendLlist) CopyFrom(src *ZendLlist) {
	l.Init(src.size, src.dtor, src.persistent)

	ptr := src.head
	for ptr != nil {
		l.AddElement(ptr.data)
		ptr = ptr.next
	}
}

func (l *ZendLlist) ApplyWithDel(f func(data any) int) {
	element := l.head
	for element != nil {
		next := element.next
		if f(element.data) != 0 {
			l.DelElement(element)
		}
		element = next
	}
}

func (l *ZendLlist) Apply(f func(data any)) {
	for element := l.head; element != nil; element = element.next {
		f(element.data)
	}
}

func (l *ZendLlist) ApplyWithArgument(f func(data any, arg any), arg any) {
	for element := l.head; element != nil; element = element.next {
		f(element.data, arg)
	}
}

func (l *ZendLlist) ApplyWithArguments(f func(data any, args ...any), args ...any) {
	for element := l.head; element != nil; element = element.next {
		f(element.data, args)
	}
}

func (l *ZendLlist) GetFirst() {

}

func (l *ZendLlist) toSlice() []*ZendLlistElement {
	if l.count == 0 {
		return nil
	}

	elements := make([]*ZendLlistElement, 0, l.count)
	for element := l.head; element != nil; element = element.next {
		elements = append(elements, element)
	}
	return elements
}

func (l *ZendLlist) Sort(compFunc LlistCompareFuncT) {
	if l.count == 0 {
		return
	}

	// 排序
	elements := l.toSlice()
	sort.Slice(elements, func(i, j int) bool {
		p := &elements[i]
		q := &elements[j]
		// todo 确认正负号
		return compFunc(p, q) > 0
	})

	// 重新链接
	for i := 1; i < l.count; i++ {
		elements[i].prev = elements[i-1]
		elements[i-1].next = elements[i]
	}
	l.head = elements[0]
	l.head.prev = nil
	l.tail = elements[len(elements)-1]
	l.tail.next = nil
}

func (l *ZendLlist) GetFirstEx(pos *ZendLlistPosition) []byte {
	var current = pos
	if pos == nil {
		current = &l.traverse_ptr
	}

	*current = l.head
	if (*current) != nil {
		return current.data
	} else {
		return nil
	}
}
func (l *ZendLlist) GetLastEx(pos *ZendLlistPosition) any {
	var current = pos
	if pos == nil {
		current = &l.traverse_ptr
	}
	*current = l.tail
	if (*current) != nil {
		return current.data
	} else {
		return nil
	}
}
func (l *ZendLlist) GetNextEx(pos *ZendLlistPosition) any {
	var current = pos
	if pos == nil {
		current = &l.traverse_ptr
	}
	if (*current) != nil {
		*current = current.next
		if (*current) != nil {
			return current.data
		}
	}
	return nil
}

func (this *ZendLlist) GetHead() *ZendLlistElement      { return this.head }
func (this *ZendLlist) SetHead(value *ZendLlistElement) { this.head = value }
func (this *ZendLlist) SetTail(value *ZendLlistElement) { this.tail = value }
func (this *ZendLlist) GetCount() int                   { return this.count }
func (this *ZendLlist) SetDtor(value LlistDtorFuncT)    { this.dtor = value }
