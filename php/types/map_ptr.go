package types

/**
 * 替代 ZEND_MAP_PTR_* 系列宏，用于表示常驻对象中的非公共数据(比如内部类在不同请求中对应的不同类静态属性)
 * todo 目前没有做不同请求上下文的区分，只保留切片用于后续迭代
 */
type MapPtr[T any] struct {
	ptr    *T
	offset bool
}

func NewMapPtr[T any]() *MapPtr[T]        { return &MapPtr[T]{ptr: nil, offset: true} }
func InitMapPtr[T any](ptr *T) *MapPtr[T] { return &MapPtr[T]{ptr: ptr, offset: false} }

func (ptr *MapPtr[T]) IsOffset() bool { return ptr.offset }
func (ptr *MapPtr[T]) Get() *T        { return ptr.ptr }
func (ptr *MapPtr[T]) Set(p *T) *T {
	ptr.ptr = p
	ptr.offset = false
	return ptr.ptr
}
func (ptr *MapPtr[T]) SetMapPtr(p *MapPtr[T]) *T {
	ptr.ptr = p.ptr
	ptr.offset = p.offset
	return ptr.ptr
}
