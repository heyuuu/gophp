package types

import (
	"github.com/heyuuu/gophp/kits/ascii"
)

/**
 * 内部频繁使用的 map, 有如下特征
 * - 有序
 * - key 为字符串，可无视大小写(可选)
 * - value 为引用且不为空。get(key) 方法返回 nil 表示 key 不存在
 */
type Table[T any] struct {
	keys       []string
	m          map[string]T
	caseIgnore bool
}

func NewLcTable[T any]() *Table[T] {
	return &Table[T]{
		keys:       nil,
		m:          make(map[string]T),
		caseIgnore: true,
	}
}
func NewTable[T any]() *Table[T] {
	return &Table[T]{
		keys:       nil,
		m:          make(map[string]T),
		caseIgnore: false,
	}
}
func (t *Table[T]) realKey(key string) string {
	if t.caseIgnore {
		return ascii.StrToLower(key)
	}
	return key
}

func (t *Table[T]) Clean()                 { t.keys, t.m = nil, make(map[string]T) }
func (t *Table[T]) Len() int               { return len(t.m) }
func (t *Table[T]) Get(key string) T       { return t.m[t.realKey(key)] }
func (t *Table[T]) Exists(key string) bool { _, ok := t.m[t.realKey(key)]; return ok }
func (t *Table[T]) Add(key string, val T) bool {
	key = t.realKey(key)
	if _, ok := t.m[key]; ok {
		return false
	} else {
		t.keys = append(t.keys, key)
		t.m[key] = val
		return true
	}
}
func (t *Table[T]) Set(key string, val T) {
	key = t.realKey(key)
	if _, ok := t.m[key]; ok {
		t.m[key] = val
	} else {
		t.keys = append(t.keys, key)
		t.m[key] = val
	}
}
func (t *Table[T]) Del(key string) {
	key = t.realKey(key)
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

func (t *Table[T]) Each(handler func(string, T)) {
	for _, k := range t.keys {
		val := t.m[k]
		handler(k, val)
	}
}

func (t *Table[T]) EachEx(handler func(string, T) bool) bool {
	for _, k := range t.keys {
		val := t.m[k]
		if !handler(k, val) {
			return false
		}
	}
	return true
}

func (t *Table[T]) Filter(handler func(string, T) bool) {
	newKeys := make([]string, 0, len(t.keys))
	for _, key := range t.keys {
		val := t.m[key]
		if handler(key, val) {
			newKeys = append(newKeys, key)
		} else {
			delete(t.m, key)
		}
	}
	t.keys = newKeys
}
