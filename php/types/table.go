package types

import (
	"github.com/heyuuu/gophp/builtin/ascii"
	"sort"
)

/**
 * 内部频繁使用的 map, 有如下特征
 * - 有序
 * - key 为字符串且可无视大小写(可选)
 * - value 为引用且不为空。get(key) 方法返回 nil 表示 key 不存在
 * - 支持元素析构函数
 */
type Table[T any] struct {
	keys       []string
	m          map[string]T
	destructor func(T)
	caseIgnore bool
}

func NewLcTable[T any](destructor func(T)) *Table[T] {
	return &Table[T]{
		keys:       nil,
		m:          make(map[string]T),
		destructor: destructor,
		caseIgnore: true,
	}
}
func NewTable[T any](destructor func(T)) *Table[T] {
	return &Table[T]{
		keys:       nil,
		m:          make(map[string]T),
		destructor: destructor,
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
func (t *Table[T]) Exists(key string) bool { return t.m[t.realKey(key)] != nil }
func (t *Table[T]) Add(key string, val T) bool {
	if val == nil {
		panic("Table.Add(key, val) 方法参数 val 不可为 nil")
	}

	key = t.realKey(key)
	if _, ok := t.m[key]; ok {
		return false
	} else {
		t.keys = append(t.keys, key)
		t.m[key] = val
		return true
	}
}
func (t *Table[T]) Update(key string, val T) {
	if val == nil {
		panic("Table.Update(key, val) 方法参数 val 不可为 nil")
	}

	key = t.realKey(key)
	if oldVal, ok := t.m[key]; ok {
		if t.destructor != nil {
			t.destructor(oldVal)
		}
		t.m[key] = val
	} else {
		t.keys = append(t.keys, key)
		t.m[key] = val
	}
}
func (t *Table[T]) UpdateDirect(key string, val T) {
	if val == nil {
		panic("Table.UpdateDirect(key, val) 方法参数 val 不可为 nil")
	}

	key = t.realKey(key)
	if _, ok := t.m[key]; ok {
		t.m[key] = val
	} else {
		t.keys = append(t.keys, key)
		t.m[key] = val
	}
}
func (t *Table[T]) Del(key string) {
	val := t.m[key]
	if val == nil {
		return
	}

	for i, k := range t.keys {
		if k != key {
			continue
		}
		copy(t.keys[i:], t.keys[i+1:])
		t.keys = t.keys[:len(t.keys)-1]
		break
	}
}
func (t *Table[T]) Values() []T {
	var values []T
	for _, key := range t.keys {
		values = append(values, t.m[key])
	}
	return values
}

func (t *Table[T]) SortByKey(less func(k1 string, k2 string) bool) {
	t.Sort(func(k1 string, v1 T, k2 string, v2 T) bool {
		return less(k1, k2)
	})
}

// todo 兼容用函数，后续会迁移
func (t *Table[T]) SortByArrayKey(less func(k1 ArrayKey, k2 ArrayKey) bool) {
	t.Sort(func(k1 string, v1 T, k2 string, v2 T) bool {
		return less(StrKey(k1), StrKey(k2))
	})
}

func (t *Table[T]) SortByValue(less func(i, j T) bool) {
	t.Sort(func(k1 string, v1 T, k2 string, v2 T) bool {
		return less(v1, v2)
	})
}

func (t *Table[T]) Sort(less func(k1 string, v1 T, k2 string, v2 T) bool) {
	sort.SliceStable(t.keys, func(i, j int) bool {
		k1 := t.keys[i]
		k2 := t.keys[j]
		v1 := t.m[k1]
		v2 := t.m[k2]
		return less(k1, v1, k2, v2)
	})
}

func (t *Table[T]) Foreach(handler func(string, T)) {
	for _, k := range t.keys {
		v := t.m[k]
		handler(k, v)
	}
}
func (t *Table[T]) ForeachReserve(handler func(string, T)) {
	for i := len(t.keys) - 1; i >= 0; i-- {
		k := t.keys[i]
		v := t.m[k]
		handler(k, v)
	}
}
func (t *Table[T]) ForeachEx(handler func(string, T) bool) bool {
	for _, k := range t.keys {
		v := t.m[k]
		if !handler(k, v) {
			return false
		}
	}
	return true
}

// todo 此方法不是并发安全的，待优化
func (t *Table[T]) Filter(handler func(string, T) bool) {
	var newKeys = make([]string, 0, cap(t.keys))
	for _, k := range t.keys {
		v := t.m[k]
		if v != nil && handler(k, v) {
			newKeys = append(newKeys, k)
		} else {
			delete(t.m, k)
		}
	}
	t.keys = newKeys
}

// todo 此方法不是并发安全的，待优化
func (t *Table[T]) FilterReserve(handler func(string, T) bool) {
	var newKeys = make([]string, 0, cap(t.keys))
	for i := len(t.keys) - 1; i >= 0; i-- {
		k := t.keys[i]
		v := t.m[k]
		if v != nil && handler(k, v) {
			newKeys = append(newKeys, k)
		} else {
			delete(t.m, k)
		}
	}
	t.keys = newKeys
}

func (t *Table[T]) Destroy() {
	if t.destructor != nil {
		t.Foreach(func(_ string, v T) { t.destructor(v) })
	}
	t.Clean()
}

func (t *Table[T]) DestroyReverse() {
	if t.destructor != nil {
		t.ForeachReserve(func(_ string, v T) { t.destructor(v) })
	}
	t.Clean()
}
