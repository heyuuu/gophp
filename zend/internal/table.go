package internal

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/builtin/ascii"
	"sort"
)

/**
 * 内部频繁使用的 map, 有如下特征
 * - 有序
 * - key 为字符串且无视大小写
 * - value 为引用且不为空。get(key) 方法返回 nil 表示 key 不存在
 * - 支持元素析构函数
 */
type LcTable[T any] struct {
	keys       []string
	m          map[string]*T
	destructor func(*T)
}

func NewLcTable[T any](destructor func(*T)) *LcTable[T] {
	return &LcTable[T]{
		keys:       nil,
		m:          make(map[string]*T),
		destructor: destructor,
	}
}

func (t *LcTable[T]) realKey(key string) string { return ascii.StrToLower(key) }

func (t *LcTable[T]) Clean()                 { t.keys, t.m = nil, make(map[string]*T) }
func (t *LcTable[T]) Len() int               { return len(t.m) }
func (t *LcTable[T]) Get(key string) *T      { return t.m[t.realKey(key)] }
func (t *LcTable[T]) Exists(key string) bool { return t.m[t.realKey(key)] != nil }
func (t *LcTable[T]) Add(key string, val *T) bool {
	if val == nil {
		panic("LcTable.Add(key, val) 方法参数 val 不可为 nil")
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
func (t *LcTable[T]) Update(key string, val *T) {
	if val == nil {
		panic("LcTable.Add(key, val) 方法参数 val 不可为 nil")
	}

	key = t.realKey(key)
	if oldVal, ok := t.m[key]; ok {
		b.Assert(val != oldVal)
		if t.destructor != nil {
			t.destructor(oldVal)
		}
		t.m[key] = val
	} else {
		t.keys = append(t.keys, key)
		t.m[key] = val
	}
}
func (t *LcTable[T]) Del(key string) {
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
func (t *LcTable[T]) Values() []*T {
	var values []*T
	for _, key := range t.keys {
		values = append(values, t.m[key])
	}
	return values
}
func (t *LcTable[T]) Sort(less func(i, j *T) bool) {
	sort.SliceStable(t.keys, func(i, j int) bool {
		return less(t.m[t.keys[i]], t.m[t.keys[j]])
	})
}
func (t *LcTable[T]) Foreach(handler func(*T)) {
	for _, key := range t.keys {
		v := t.m[key]
		handler(v)
	}
}
func (t *LcTable[T]) ForeachReserve(handler func(*T)) {
	for i := len(t.keys) - 1; i >= 0; i-- {
		v := t.m[t.keys[i]]
		handler(v)
	}
}

// todo 此方法不是并发安全的，待优化
func (t *LcTable[T]) Filter(handler func(*T) bool) {
	var newKeys = make([]string, 0, cap(t.keys))
	for _, key := range t.keys {
		v := t.m[key]
		if v != nil && handler(v) {
			newKeys = append(newKeys, key)
		} else {
			delete(t.m, key)
		}
	}
	t.keys = newKeys
}

func (t *LcTable[T]) Destroy() {
	if t.destructor != nil {
		t.Foreach(t.destructor)
	}
	t.Clean()
}
