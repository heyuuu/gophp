package types

import "github.com/heyuuu/gophp/shim/slices"

// ArrayData
type ArrayData interface {
	Clone() ArrayData

	/**
	 * 读操作
	 */
	// 元素个数
	Len() int
	// 有效元素个数. 与 Len() 不同，它需要过滤 IS_INDIRECT 元素为 IS_UNDEF 的情况
	Count() int
	// 元素容量
	Cap() int
	// 已使用的容量，同时也相当于 pos 的值范围
	Used() int
	Exists(key ArrayKey) bool
	Find(key ArrayKey) (value *Zval, pos ArrayPosition)
	Each(func(key ArrayKey, value *Zval) error) error
	EachReserve(func(key ArrayKey, value *Zval) error) error

	/**
	 * pos 表示数据在内部的具体位置，具体值由具体 ArrayData 实现确定，满足如下条件
	 * - 每个 pos 对应 0~1 具体合法值，在无写操作的情况下对应值不变
	 * - 在 [0, Used()] 之间可能有对应值，其他位置肯定无对应值
	 * - 未进行写操作时，Used() 值及各 pos 对应的值不会发生变化
	 */
	Pos(pos ArrayPosition) (key ArrayKey, value *Zval)

	/**
	 * 写操作
	 * - 返回值中的 error 目前只有一种非 nil 值 arrayDataUnsupported，表示当前类型 ArrayData 不支持此写行为
	 */
	Add(key ArrayKey, value *Zval) (bool, error)
	Update(key ArrayKey, value *Zval) error
	Delete(key ArrayKey) (bool, error)
	Append(value *Zval) (int, error)
}

// errors
type arrayDataOpError string

func (e arrayDataOpError) Error() string { return string(e) }

var arrayDataUnsupported = arrayDataOpError("this kind of array data cannot support this operate")

// emptyArrayData
var emptyArrayData ArrayData = emptyArrayDataType{}

type emptyArrayDataType struct{}

func (d emptyArrayDataType) Clone() ArrayData                                          { return d }
func (d emptyArrayDataType) Len() int                                                  { return 0 }
func (d emptyArrayDataType) Count() int                                                { return 0 }
func (d emptyArrayDataType) Cap() int                                                  { return 0 }
func (d emptyArrayDataType) Used() int                                                 { return 0 }
func (d emptyArrayDataType) Exists(key ArrayKey) bool                                  { return false }
func (d emptyArrayDataType) Find(key ArrayKey) (*Zval, ArrayPosition)                  { return nil, InvalidArrayPos }
func (d emptyArrayDataType) Each(f func(key ArrayKey, value *Zval) error) error        { return nil }
func (d emptyArrayDataType) EachReserve(f func(key ArrayKey, value *Zval) error) error { return nil }
func (d emptyArrayDataType) Pos(pos ArrayPosition) (key ArrayKey, value *Zval)         { return }

func (d emptyArrayDataType) Add(key ArrayKey, value *Zval) (bool, error) {
	return false, arrayDataUnsupported
}
func (d emptyArrayDataType) Update(key ArrayKey, value *Zval) error {
	return arrayDataUnsupported
}
func (d emptyArrayDataType) Delete(key ArrayKey) (bool, error) {
	return false, arrayDataUnsupported
}
func (d emptyArrayDataType) Append(value *Zval) (int, error) {
	return 0, arrayDataUnsupported
}

// ArrayDataList
var _ ArrayData = (*ArrayDataList[any])(nil)

type ArrayDataList[T any] struct {
	data   []T
	wrap   func(T) *Zval
	unwrap func(*Zval) (T, bool)
}

func NewArrayDataList[T any](data []T, wrap func(T) *Zval, unwrap func(*Zval) (T, bool)) *ArrayDataList[T] {
	return &ArrayDataList[T]{data: data, wrap: wrap, unwrap: unwrap}
}

func (l *ArrayDataList[T]) Clone() ArrayData {
	return NewArrayDataList(slices.Clone(l.data), l.wrap, l.unwrap)
}
func (l *ArrayDataList[T]) Len() int   { return len(l.data) }
func (l *ArrayDataList[T]) Count() int { return len(l.data) }
func (l *ArrayDataList[T]) Cap() int   { return len(l.data) }
func (l *ArrayDataList[T]) Used() int  { return len(l.data) }
func (l *ArrayDataList[T]) Exists(key ArrayKey) bool {
	return key.IsIdxKey() && 0 <= key.IdxKey() && key.IdxKey() < len(l.data)
}

func (l *ArrayDataList[T]) Find(key ArrayKey) (value *Zval, pos ArrayPosition) {
	if l.Exists(key) {
		idx := key.IdxKey()
		return l.wrap(l.data[idx]), ArrayPosition(idx)
	}
	return nil, InvalidArrayPos
}

func (l *ArrayDataList[T]) Each(f func(key ArrayKey, value *Zval) error) error {
	for idx, value := range l.data {
		if err := f(IdxKey(idx), l.wrap(value)); err != nil {
			return err
		}
	}
	return nil
}

func (l *ArrayDataList[T]) EachReserve(f func(key ArrayKey, value *Zval) error) error {
	for idx := len(l.data) - 1; idx >= 0; idx-- {
		value := l.data[idx]
		if err := f(IdxKey(idx), l.wrap(value)); err != nil {
			return err
		}
	}
	return nil
}

func (l *ArrayDataList[T]) Pos(pos ArrayPosition) (key ArrayKey, value *Zval) {
	if 0 <= pos && pos < len(l.data) {
		return IdxKey(pos), l.wrap(l.data[pos])
	}
	return
}

func (l *ArrayDataList[T]) Add(key ArrayKey, value *Zval) (bool, error) {
	if key.IsStrKey() || key.IdxKey() < 0 || key.IdxKey() > len(l.data) {
		return false, arrayDataUnsupported
	}

	val, ok := l.unwrap(value)
	if !ok {
		return false, arrayDataUnsupported
	}

	if key.IdxKey() == len(l.data) {
		l.data = append(l.data, val)
		return true, nil
	} else {
		return false, nil
	}
}

func (l *ArrayDataList[T]) Update(key ArrayKey, value *Zval) error {
	if key.IsStrKey() || key.IdxKey() < 0 || key.IdxKey() > len(l.data) {
		return arrayDataUnsupported
	}

	val, ok := l.unwrap(value)
	if !ok {
		return arrayDataUnsupported
	}

	if key.IdxKey() == len(l.data) {
		l.data = append(l.data, val)
	} else {
		l.data[key.IdxKey()] = val
	}
	return nil
}

func (l *ArrayDataList[T]) Delete(key ArrayKey) (bool, error) {
	if !l.Exists(key) {
		return false, nil
	}

	if key.IdxKey() == len(l.data)-1 { // 支持删除最后一个元素
		l.data = l.data[:len(l.data)-1]
		return true, nil
	} else {
		return false, arrayDataUnsupported
	}
}

func (l *ArrayDataList[T]) Append(value *Zval) (int, error) {
	val, ok := l.unwrap(value)
	if !ok {
		return 0, arrayDataUnsupported
	}

	l.data = append(l.data, val)
	return len(l.data) - 1, nil
}
