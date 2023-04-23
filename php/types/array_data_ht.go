package types

import (
	"fmt"
	"math"
)

var _ ArrayData = (*ArrayDataHt)(nil)

type ArrayDataHt struct {
	elementsCount   uint32
	internalPointer uint32
	nextFreeElement int
	data            []Bucket          // 实际存储数据的地方
	indexMap        map[int]uint32    // 数字索引到具体位置的映射
	keyMap          map[string]uint32 // 字符串索引到具体位置的映射
}

func (ht *ArrayDataHt) Len() int {
	return int(ht.elementsCount)
}

func (ht *ArrayDataHt) Exists(key ArrayKey) bool {
	if pos := ht.index(key); pos >= 0 {
		return true
	}
	return false
}

func (ht *ArrayDataHt) Find(key ArrayKey) *Zval {
	if pos := ht.index(key); pos >= 0 {
		return ht.data[pos].GetVal()
	}
	return nil
}

func (ht *ArrayDataHt) Add(key ArrayKey, data *Zval) bool {
	if pos := ht.index(key); pos >= 0 {
		return false
	}

	bucket := NewBucket(key, data)

	//TODO implement me
	panic("implement me")
}

func (ht *ArrayDataHt) Update(key ArrayKey, data *Zval) {
	//TODO implement me
	panic("implement me")
}

func (ht *ArrayDataHt) Delete(key ArrayKey) bool {
	//TODO implement me
	panic("implement me")
}

func (ht *ArrayDataHt) index(key ArrayKey) int {
	var pos uint32
	var ok bool
	if key.IsStrKey() {
		pos, ok = ht.keyMap[key.StrKey()]
	} else {
		pos, ok = ht.indexMap[key.IdxKey()]
	}
	if ok {
		return int(pos)
	} else {
		return -1
	}
}

func (ht *ArrayDataHt) appendBucket(key ArrayKey, value *Zval) {
	ht.resizeIfFull()

	bucket := NewBucket(key, value)

	pos := uint32(len(ht.data))
	ht.elementsCount++
	ht.data = append(ht.data, *bucket)

	ht.addHash(bucket.GetArrayKey(), pos)

	if !key.IsStrKey() {
		idxKey := key.IdxKey()
		if idxKey > ht.nextFreeElement {
			ht.nextFreeElement = idxKey + 1
		}
	}

}

func (ht *ArrayDataHt) addHash(key ArrayKey, pos uint32) {
	if key.IsStrKey() {
		ht.keyMap[key.StrKey()] = pos
	} else {
		ht.indexMap[key.IdxKey()] = pos
	}
}

func (ht *ArrayDataHt) resizeIfFull() {
	dataSize := len(ht.data)
	if dataSize == cap(ht.data) {
		// 若空隙率过高，重新压缩；否则，跳过扩容 (后面会由 append(ht.data) 触发自动扩容)
		if dataSize > int(ht.elementsCount+(ht.elementsCount>>5)) {
			ht.Rehash()
		} else if dataSize >= math.MaxInt32 {
			triggerError(fmt.Sprintf("Possible integer overflow in memory allocation (%d)", dataSize*2))
		}
	}
}
