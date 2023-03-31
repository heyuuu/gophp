package collections

import "sort"

type OrderedMap[K comparable, V any] struct {
	keys []K
	m    map[K]V
}

func (o *OrderedMap[K, V]) Len() int {
	return len(o.m)
}

func (o *OrderedMap[K, V]) Exists(key K) bool {
	if o.m == nil {
		return false
	}

	_, ok := o.m[key]
	return ok
}

func (o *OrderedMap[K, V]) Get(key K) (V, bool) {
	if o.m == nil {
		return nil, false
	}

	val, ok := o.m[key]
	return val, ok
}

func (o *OrderedMap[K, V]) Set(key K, value V) {
	if o.m == nil {
		o.m = make(map[K]V)
	}
	if _, ok := o.m[key]; !ok {
		o.keys = append(o.keys, key)
	}
	o.m[key] = value
}

func (o *OrderedMap[K, V]) Del(key K) {
	if !o.Exists(key) {
		return
	}

	delete(o.m, key)
	for i, k := range o.keys {
		if k != key {
			continue
		}
		copy(o.keys[i:], o.keys[i+1:])
		o.keys = o.keys[:len(o.keys)-1]
		break
	}
}

func (o *OrderedMap[K, V]) Keys() []K { return o.keys }
func (o *OrderedMap[K, V]) Values() []V {
	var values []V
	for _, k := range o.keys {
		values = append(values, o.m[k])
	}
	return values
}

func (o *OrderedMap[K, V]) Sort(less func(i, j V) bool) {
	sort.SliceStable(o.keys, func(i, j int) bool {
		return less(o.m[o.keys[i]], o.m[o.keys[j]])
	})
}

func (o *OrderedMap[K, V]) Foreach(handler func(key K, value V)) {
	for _, k := range o.keys {
		handler(k, o.m[k])
	}
}

func (o *OrderedMap[K, V]) ForeachReserve(handler func(key K, value V)) {
	for i := len(o.keys) - 1; i >= 0; i-- {
		k := o.keys[i]
		v := o.m[k]
		handler(k, v)
	}
}

func (o *OrderedMap[K, V]) ForeachValues(handler func(value V)) {
	for _, k := range o.keys {
		handler(o.m[k])
	}
}

func (o *OrderedMap[K, V]) ForeachReserveValues(handler func(value V)) {
	for i := len(o.keys) - 1; i >= 0; i-- {
		k := o.keys[i]
		v := o.m[k]
		handler(v)
	}
}
