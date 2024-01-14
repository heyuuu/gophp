package collections

type orderedMapNode[K comparable, V any] struct {
	key  K
	val  V
	prev *orderedMapNode[K, V]
	next *orderedMapNode[K, V]
}

type OrderedMap[K comparable, V any] struct {
	m    map[K]*orderedMapNode[K, V]
	head *orderedMapNode[K, V]
	tail *orderedMapNode[K, V]
}

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		m:    make(map[K]*orderedMapNode[K, V]),
		head: nil,
		tail: nil,
	}
}

func (m *OrderedMap[K, V]) Len() int {
	return len(m.m)
}

func (m *OrderedMap[K, V]) Get(key K) (val V, ok bool) {
	if node, ok := m.m[key]; ok {
		return node.val, true
	}
	return
}

func (m *OrderedMap[K, V]) GetVal(key K) (val V) {
	if node, ok := m.m[key]; ok {
		return node.val
	}
	return
}

func (m *OrderedMap[K, V]) Set(key K, val V) {
	if node, ok := m.m[key]; ok {
		node.val = val
		return
	}

	node := &orderedMapNode[K, V]{key: key, val: val, prev: m.tail, next: nil}
	m.m[key] = node
	if m.tail != nil {
		m.tail.next = node
	}
	m.tail = node
	if m.head == nil {
		m.head = node
	}
}

func (m *OrderedMap[K, V]) Del(key K, val V) {
	node, ok := m.m[key]
	if !ok {
		return
	}

	delete(m.m, node.key)
	if node.prev == nil {
		m.head = node.next
	} else {
		node.prev.next = node.next
	}
	if node.next == nil {
		m.tail = node.prev
	} else {
		node.next.prev = node.prev
	}
}

func (m *OrderedMap[K, V]) Each(f func(key K, val V) error) error {
	for curr := m.head; curr != nil; curr = curr.next {
		err := f(curr.key, curr.val)
		if err != nil {
			return err
		}
	}
	return nil
}
