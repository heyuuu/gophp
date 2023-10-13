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

func Get[K cmp.Ordered, V any](m map[K]V, key K) V {
	var element V
	if m != nil {
		element = m[key]
	}
	return element
}

func Add[K string, V any](m *map[K]V, key K, value V) bool {
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
