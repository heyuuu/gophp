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
