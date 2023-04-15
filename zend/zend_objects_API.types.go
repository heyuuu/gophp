package zend

import "github.com/heyuuu/gophp/zend/types"

/**
 * ObjectsStore
 * 简化为仅作为对象 ZendObject.handle 的发号器
 */
type ObjectsStore struct {
	lastHandle uint32 // 最后一个object的handle值，初始为0
}

func (s *ObjectsStore) PutObject(obj *types.ZendObject) {
	// notice: 此处的 handle 是从 1 开始的，以保持索引 == ZendObject.handle 且 ZendObject.handle 始终大于 0
	s.lastHandle++
	handle := s.lastHandle
	obj.SetHandle(handle)
}
