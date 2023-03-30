package collections

type orderedBucket[V any] struct {
	prev  *orderedBucket[V]
	next  *orderedBucket[V]
	value V
}

type OrderedMap[K comparable, V any] struct {
	m    map[K]*orderedBucket[V]
	head *orderedBucket[V]
	tail *orderedBucket[V]
}

func (o *OrderedMap[K, V]) Len() int {
	return len(o.m)
}

func (o *OrderedMap[K, V]) Get(key K) (value V, exists bool) {
	if o.m == nil {
		return value, false
	}

	if b, ok := o.m[key]; ok {
		return b.value, true
	} else {
		return value, false
	}
}

func (o *OrderedMap[K, V]) Set(key K, value V) {
	if o.m == nil {
		o.m = make(map[K]*orderedBucket[V])
	}
	if b, ok := o.m[key]; ok {
		b.value = value
	} else {
		newBucket := &orderedBucket[V]{
			prev:  o.tail,
			next:  nil,
			value: value,
		}
		o.m[key] = newBucket
		if o.head == nil {
			o.head = newBucket
		}
		if o.tail != nil {
			o.tail.next = newBucket
		}
		o.tail = newBucket
	}
}

func (o *OrderedMap[K, V]) Del(key K) {
	if o.m == nil {
		return
	}
	if b, ok := o.m[key]; ok {
		if o.head == b {
			o.head = b.next
		}
		if o.tail == b {
			o.tail = b.prev
		}
		if b.prev != nil {
			b.prev.next = b.next
		}
		if b.next != nil {
			b.next.prev = b.prev
		}
		delete(o.m, key)
	}
}
