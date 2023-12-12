package types

import (
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/shim/slices"
)

/**
 * 内部频繁使用的 map, 有如下特征
 * - 有序
 * - key 为字符串，可无视大小写(可选)
 * - 主要操作是查询和添加，较少有删除操作
 */
type Table[T any] struct {
	keys []string
	m    map[string]T
}

func NewTable[T any]() *Table[T] {
	return &Table[T]{
		keys: nil,
		m:    make(map[string]T),
	}
}

func (t *Table[T]) Clean()                    { t.keys, t.m = nil, make(map[string]T) }
func (t *Table[T]) Len() int                  { return len(t.m) }
func (t *Table[T]) Get(key string) T          { return t.m[key] }
func (t *Table[T]) Find(key string) (T, bool) { v, ok := t.m[key]; return v, ok }
func (t *Table[T]) Exists(key string) bool    { _, ok := t.m[key]; return ok }
func (t *Table[T]) Add(key string, val T) bool {
	if _, ok := t.m[key]; ok {
		return false
	} else {
		t.keys = append(t.keys, key)
		t.m[key] = val
		return true
	}
}
func (t *Table[T]) AddToHead(key string, val T) bool {
	if _, ok := t.m[key]; ok {
		return false
	} else {
		t.keys = append([]string{key}, t.keys...)
		t.m[key] = val
		return true
	}
}
func (t *Table[T]) Set(key string, val T) {
	if _, ok := t.m[key]; ok {
		t.m[key] = val
	} else {
		t.keys = append(t.keys, key)
		t.m[key] = val
	}
}
func (t *Table[T]) Del(key string) {
	if _, exists := t.m[key]; !exists {
		return
	}

	for i, k := range t.keys {
		if k != key {
			continue
		}
		copy(t.keys[i:], t.keys[i+1:])
		t.keys = t.keys[:len(t.keys)-1]
		delete(t.m, key)
		break
	}
}

func (t *Table[T]) Keys() []string { return slices.Clone(t.keys) }
func (t *Table[T]) Values() []T {
	values := make([]T, len(t.keys))
	for i, key := range t.keys {
		values[i] = t.m[key]
	}
	return values
}

func (t *Table[T]) Each(handler func(string, T)) {
	for _, k := range t.keys {
		v := t.m[k]
		handler(k, v)
	}
}
func (t *Table[T]) EachReserve(handler func(string, T)) {
	for i := len(t.keys) - 1; i >= 0; i-- {
		k := t.keys[i]
		v := t.m[k]
		handler(k, v)
	}
}
func (t *Table[T]) EachEx(handler func(string, T) error) error {
	perr.Assert(t != nil)
	for _, k := range t.keys {
		v := t.m[k]
		err := handler(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *Table[T]) Filter(handler func(string, T) bool) {
	newKeys := make([]string, 0, len(t.keys))
	for _, k := range t.keys {
		v := t.m[k]
		if handler(k, v) {
			newKeys = append(newKeys, k)
		} else {
			delete(t.m, k)
		}
	}
	t.keys = newKeys
}
