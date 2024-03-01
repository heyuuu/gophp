package types

import (
	"github.com/heyuuu/gophp/kits/mapkit"
	"github.com/heyuuu/gophp/php/assert"
	"maps"
	"slices"
	"sort"
)

/**
 * ArrayDataHt 默认数据数据
 * - ArrayData 的基础数据类型，能实现除符号表(SymtableData)外的所有数据操作
 * - 此类型内所有 Zval 值都不为 IsIndirect 类型
 */
type ArrayDataHt struct {
	elementsCount   int
	internalPointer int
	nextFreeElement int
	data            []ArrayPair      // 实际存储数据的地方
	indexes         map[ArrayKey]int // 索引到具体位置的映射
	writable        bool
}

func newArrayDataHt(cap int) *ArrayDataHt {
	return &ArrayDataHt{
		indexes:  make(map[ArrayKey]int, cap),
		data:     make([]ArrayPair, 0, cap),
		writable: true,
	}
}

func newArrayDataHtByData(pairs []ArrayPair) *ArrayDataHt {
	ht := &ArrayDataHt{
		elementsCount: len(pairs),
		indexes:       make(map[ArrayKey]int, len(pairs)),
		data:          slices.Clone(pairs),
		writable:      true,
	}
	ht.rehash()
	return ht
}

func (ht *ArrayDataHt) Clone() ArrayData {
	return &ArrayDataHt{
		elementsCount:   ht.elementsCount,
		internalPointer: ht.internalPointer,
		nextFreeElement: ht.nextFreeElement,
		data:            slices.Clone(ht.data),
		indexes:         maps.Clone(ht.indexes),
		writable:        true,
	}
}

func (ht *ArrayDataHt) Len() int   { return ht.elementsCount }
func (ht *ArrayDataHt) Used() int  { return len(ht.data) }
func (ht *ArrayDataHt) Cap() int   { return cap(ht.data) }
func (ht *ArrayDataHt) Count() int { return ht.elementsCount }
func (ht *ArrayDataHt) Exists(key ArrayKey) bool {
	_, ok := ht.indexes[key]
	return ok
}
func (ht *ArrayDataHt) Find(key ArrayKey) (Zval, ArrayPosition) {
	if pos, ok := ht.indexes[key]; ok {
		return ht.data[pos].Val, pos
	}
	return Undef, InvalidArrayPos
}
func (ht *ArrayDataHt) Each(handler func(key ArrayKey, value Zval) error) error {
	for i, _ := range ht.data {
		if !ht.isValid(i) {
			continue
		}

		p := ht.data[i]
		err := handler(p.Key, p.Val)
		if err != nil {
			return err
		}
	}
	return nil
}
func (ht *ArrayDataHt) EachReserve(handler func(key ArrayKey, value Zval) error) error {
	for i := len(ht.data) - 1; i >= 0; i-- {
		if !ht.isValid(i) {
			continue
		}

		p := ht.data[i]
		err := handler(p.Key, p.Val)
		if err != nil {
			return err
		}
	}
	return nil
}
func (ht *ArrayDataHt) Pos(pos ArrayPosition) ArrayPair {
	if pos < 0 || pos >= len(ht.data) || !ht.isValid(pos) {
		return invalidArrayPair
	}

	return ht.data[pos]
}

func (ht *ArrayDataHt) Add(key ArrayKey, value Zval) (bool, error) {
	ht.assertWritable()
	if _, exists := ht.indexes[key]; exists {
		return false, nil
	} else {
		ht.appendBucket(key, value)
		return true, nil
	}
}
func (ht *ArrayDataHt) Update(key ArrayKey, value Zval) error {
	ht.assertWritable()
	if pos, exists := ht.indexes[key]; exists {
		// todo 此处需要做懒复制(主要针对数组)
		ht.data[pos].Val = value
	} else {
		ht.appendBucket(key, value)
	}
	return nil
}
func (ht *ArrayDataHt) Delete(key ArrayKey) (bool, error) {
	ht.assertWritable()
	if pos, exists := ht.indexes[key]; exists {
		ht.deleteBucket(pos)
		return true, nil
	} else {
		return false, nil
	}
}
func (ht *ArrayDataHt) Append(value Zval) (int, error) {
	ht.assertWritable()
	idx := ht.nextFreeElement
	key := IdxKey(idx)
	assert.Assert(!ht.Exists(key))
	ht.appendBucket(key, value)
	return idx, nil
}
func (ht *ArrayDataHt) Clean() {
	ht.assertWritable()

	ht.elementsCount = 0
	ht.internalPointer = 0
	ht.nextFreeElement = 0
	ht.data = nil
	ht.indexes = make(map[ArrayKey]int)
}

func (ht *ArrayDataHt) Sort(comparer ArrayComparer, renumber bool) {
	ht.assertWritable()

	if ht.elementsCount == 0 || (ht.elementsCount == 1 && !renumber) {
		return
	}

	ht.removeHoles()
	ht.internalPointer = 0

	sort.SliceStable(ht.data, func(i, j int) bool {
		ret := comparer.Compare(ht.data[i], ht.data[j])
		return ret < 0
	})

	if renumber {
		for pos, _ := range ht.data {
			ht.data[pos].Key = IdxKey(pos)
		}
		ht.nextFreeElement = len(ht.data)
	}

	ht.rehash()
}

/**
 * Internal methods
 */

func (ht *ArrayDataHt) assertWritable() { assert.Assert(ht.writable) }

func (ht *ArrayDataHt) appendBucket(key ArrayKey, value Zval) {
	// 尝试 resize
	ht.resizeIfFull()

	// 添加数据，更新元素计数
	ht.elementsCount++
	ht.indexes[key] = len(ht.data)
	ht.data = append(ht.data, MakeArrayPair(key, value))

	// 更新 ht.nextFreeElement
	if key.IsIdxKey() {
		idx := key.IdxKey()
		if idx >= ht.nextFreeElement {
			if idx < MaxLong {
				ht.nextFreeElement = idx + 1
			} else {
				ht.nextFreeElement = MaxLong
			}
		}
	}
}

func (ht *ArrayDataHt) resizeIfFull() {
	dataSize := len(ht.data)
	if dataSize == cap(ht.data) {
		// 若空隙率过高，重新压缩；否则，跳过扩容 (后面会由 append(ht.data) 触发自动扩容)
		if dataSize > ht.elementsCount+ht.elementsCount>>5 {
			ht.rehash()
		}
	}
}

func (ht *ArrayDataHt) deleteBucket(pos int) {
	ht.assertWritable()
	assert.Assert(0 <= pos && pos < len(ht.data))
	assert.Assert(ht.isValid(pos))

	// 移除数据，更新元素计数
	ht.elementsCount--
	delete(ht.indexes, ht.data[pos].Key)
	ht.markInvalid(pos)

	// 更新内部指针和遍历器指针
	if ht.internalPointer == pos {
		ht.internalPointer = ht.validPos(pos + 1)
	}

	// 若删除队尾元素，尝试清除 data 队尾无用数据
	if pos == len(ht.data)-1 {
		newDataSize := len(ht.data) - 1
		for newDataSize > 0 && !ht.isValid(newDataSize-1) {
			newDataSize--
		}

		ht.data = ht.data[:newDataSize]
		if ht.internalPointer > newDataSize {
			ht.internalPointer = newDataSize
		}
	}
}

func (ht *ArrayDataHt) rehash() {
	// reset hash
	ht.assertWritable()
	mapkit.Clean(ht.indexes)

	// 空数组快速清空
	if ht.elementsCount == 0 {
		ht.data = nil
		return
	}

	// 移除 data 中的空位
	ht.removeHoles()

	// 重建 hash
	for pos, _ := range ht.data {
		// removeHoles 后，此处不用判断 p.IsValid()
		key := ht.data[pos].Key
		ht.indexes[key] = pos
	}
}

func (ht *ArrayDataHt) isValid(pos int) bool {
	return !ht.data[pos].Val.IsUndef()
}
func (ht *ArrayDataHt) markInvalid(pos int) {
	ht.data[pos].Val = Undef
}

// 移除 this.data 数据中的 holes
func (ht *ArrayDataHt) removeHoles() {
	ht.assertWritable()

	if len(ht.data) == ht.elementsCount {
		return
	}

	newPos := 0
	for pos, _ := range ht.data {
		if !ht.isValid(pos) {
			continue
		}
		if newPos != pos {
			ht.data[newPos] = ht.data[pos]
			if ht.internalPointer == pos {
				ht.internalPointer = newPos
			}
		}
		newPos++
	}

	// 截取数据，记录有效元素数
	ht.data = ht.data[:newPos]
	ht.elementsCount = newPos
}

func (ht *ArrayDataHt) validPos(pos int) int {
	for ; pos < len(ht.data); pos++ {
		if ht.isValid(pos) {
			return pos
		}
	}
	return len(ht.data)
}

func (ht *ArrayDataHt) validPosReserve(pos int) int {
	for ; pos >= 0; pos-- {
		if ht.isValid(pos) {
			return pos
		}
	}
	return InvalidArrayPos
}

func (ht *ArrayDataHt) Current() int {
	ht.internalPointer = ht.validPos(ht.internalPointer)
	return ht.internalPointer
}

func (ht *ArrayDataHt) ResetPointer() {
	ht.internalPointer = 0
}

func (ht *ArrayDataHt) MoveNext() int {
	ht.internalPointer = ht.validPos(ht.internalPointer + 1)
	return ht.internalPointer
}

func (ht *ArrayDataHt) MovePrev() int {
	ht.internalPointer = ht.validPosReserve(ht.internalPointer - 1)
	return ht.internalPointer
}

func (ht *ArrayDataHt) MoveEnd() int {
	ht.internalPointer = ht.validPosReserve(len(ht.data) - 1)
	return ht.internalPointer
}
