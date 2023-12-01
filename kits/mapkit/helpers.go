package mapkit

import (
	"github.com/heyuuu/gophp/shim/cmp"
	"github.com/heyuuu/gophp/shim/maps"
	"github.com/heyuuu/gophp/shim/slices"
)

func SortedKeys[K cmp.Ordered, V any](m map[K]V) []K {
	keys := maps.Keys(m)
	slices.Sort(keys)
	return keys
}

func Each[K comparable, V any](m map[K]V, handler func(K, V)) {
	for k, v := range m {
		handler(k, v)
	}
}

func EachEx[K comparable, V any](m map[K]V, handler func(K, V) error) error {
	for k, v := range m {
		err := handler(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetValue[K comparable, V any](m map[K]V, key K) V {
	v, _ := Get(m, key)
	return v
}

func Get[K comparable, V any](m map[K]V, key K) (V, bool) {
	if m == nil {
		var temp V
		return temp, false
	} else {
		v, ok := m[key]
		return v, ok
	}
}

func Add[K comparable, V any](m *map[K]V, key K, value V) bool {
	if *m == nil {
		*m = map[K]V{key: value}
		return true
	} else if _, exists := (*m)[key]; exists {
		return false
	} else {
		(*m)[key] = value
		return true
	}
}
func Set[K comparable, V any](m *map[K]V, key K, value V) {
	if *m == nil {
		*m = map[K]V{key: value}
	} else {
		(*m)[key] = value
	}
}

func Clean[K comparable, V any](m map[K]V) {
	if m == nil {
		return
	}
	for k := range m {
		delete(m, k)
	}
}
