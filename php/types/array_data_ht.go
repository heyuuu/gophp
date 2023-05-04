package types

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"sort"
)

var _ ArrayData = (*ArrayDataHt)(nil)

type ArrayDataHt struct {
	elementsCount   uint32
	internalPointer uint32
	nextFreeElement int
	data            []Bucket            // 实际存储数据的地方
	indexes         map[ArrayKey]uint32 // 索引到具体位置的映射

	hasIndirect bool
	writable    bool
}

func (ht *ArrayDataHt) HasIndirect() bool { return ht.hasIndirect }

func (ht *ArrayDataHt) Len() int { return int(ht.elementsCount) }
func (ht *ArrayDataHt) Cap() int { return len(ht.data) }
func (ht *ArrayDataHt) Exists(key ArrayKey) bool {
	_, ok := ht.indexes[key]
	return ok
}
func (ht *ArrayDataHt) Find(key ArrayKey) *Zval {
	if pos, ok := ht.indexes[key]; ok {
		return ht.data[pos].GetVal()
	}
	return nil
}
func (ht *ArrayDataHt) Add(key ArrayKey, data *Zval) bool {
	ht.assertWritable()
	if ht.Exists(key) {
		return false
	}

	ht.appendBucket(key, data)
	return true
}
func (ht *ArrayDataHt) Update(key ArrayKey, data *Zval) {
	ht.assertWritable()
	if pos, ok := ht.indexes[key]; ok {
		p := ht.data[pos]
		oldData := p.GetVal()
		b.Assert(oldData != data)
		p.SetVal(data)
		return
	}
	ht.appendBucket(key, data)
}
func (ht *ArrayDataHt) Delete(key ArrayKey) bool {
	ht.assertWritable()
	if pos, ok := ht.indexes[key]; ok {
		ht.deleteBucket(pos)
		return true
	}
	return false
}
func (ht *ArrayDataHt) Push(data *Zval) int {
	ht.assertWritable()
	idx := ht.nextFreeElement
	key := IdxKey(idx)
	if ht.Exists(key) {
		panic("")
	}
	ht.appendBucket(key, data)
	return idx
}
func (ht *ArrayDataHt) Clean() {
	ht.assertWritable()

	// clear data
	ht.elementsCount = 0
	ht.internalPointer = 0
	ht.nextFreeElement = 0
	ht.data = nil
	ht.indexes = make(map[ArrayKey]uint32)
}

func (ht *ArrayDataHt) appendBucket(key ArrayKey, zv *Zval) *Bucket {
	bucket := NewBucket(key, zv)

	// 尝试 resize
	ht.resizeIfFull()

	// 添加到 data
	var idx = uint32(len(ht.data))
	ht.elementsCount++
	ht.data = append(ht.data, *bucket)

	// 更新 map
	ht.indexes[bucket.key] = idx

	if !key.IsStrKey() {
		var idxKey = key.IdxKey()
		// 更新 nextFreeElement
		if idxKey > ht.nextFreeElement {
			if idxKey < MaxLong {
				ht.nextFreeElement = idxKey + 1
			} else {
				ht.nextFreeElement = MaxLong
			}
		}
	}

	return &ht.data[idx]
}

func (ht *ArrayDataHt) resizeIfFull() {
	dataSize := len(ht.data)
	if dataSize == cap(ht.data) {
		// 若空隙率过高，重新压缩；否则，跳过扩容 (后面会由 append(ht.data) 触发自动扩容)
		if dataSize > int(ht.elementsCount+(ht.elementsCount>>5)) {
			ht.rehash()
		} else if dataSize >= MaxArraySize {
			triggerError(fmt.Sprintf("Possible integer overflow in memory allocation (%d)", dataSize*2))
		}
	}
}

func (ht *ArrayDataHt) deleteBucket(pos uint32) {
	ht.assertWritable()
	assert(int(pos) < len(ht.data))

	var p = &ht.data[pos]
	assert(p.IsValid())

	// 移除映射
	delete(ht.indexes, p.key)

	// 减少有效元素
	ht.elementsCount--

	// 更新内部指针和遍历器指针
	if ht.internalPointer == pos {
		var newIdx = ht.validPosVal(pos + 1)
		if ht.internalPointer == pos {
			ht.internalPointer = newIdx
		}
	}

	// 设置数据不可用
	p.GetVal().SetUndef()

	// 若删除队尾元素，尝试清除 data 队尾无用数据
	dataSize := uint32(len(ht.data))
	if pos == dataSize-1 {
		newDataSize := dataSize
		for newDataSize > 0 && !ht.data[newDataSize-1].IsValid() {
			newDataSize--
		}

		ht.data = ht.data[:newDataSize]
		if ht.internalPointer > newDataSize {
			ht.internalPointer = newDataSize
		}
	}
}

/* hash -> Array.indexMap & Array.keyMap */
func (ht *ArrayDataHt) rehash() {
	// reset hash
	ht.assertWritable()
	ht.indexes = make(map[ArrayKey]uint32)

	// 空数组快速清空
	if ht.elementsCount == 0 {
		ht.data = nil
		return
	}

	// 移除 data 中的空位
	ht.removeHoles()

	// 重建 hash
	for pos, p := range ht.data {
		// removeHoles 后，此处不用判断 p.IsValid()
		ht.indexes[p.key] = uint32(pos)
	}
}

/* misc */
func (ht *ArrayDataHt) assertWritable()  { assert(ht.writable) }
func (ht *ArrayDataHt) DataSize() uint32 { return uint32(len(ht.data)) }
func (ht *ArrayDataHt) Count() int {
	if !ht.hasIndirect {
		return ht.Len()
	}

	// 计算有效元素个数(与 elementsCount 不同，它需要过滤 IS_INDIRECT 元素为 IS_UNDEF 的情况)
	var num = 0
	for _, p := range ht.data {
		data := p.GetVal()
		if data.IsIndirect() {
			data = data.Indirect()
		}
		if data.IsUndef() {
			continue
		}
		num++
	}
	return num
}

/**
 * Sort
 */
func (ht *ArrayDataHt) Sort(comparer ArrayComparer, renumber bool) {
	ht.assertWritable()

	if ht.elementsCount == 0 || (ht.elementsCount == 1 && !renumber) {
		return
	}

	ht.removeHolesAndCleanInternalPointer()

	sort.SliceStable(ht.data, func(i, j int) bool {
		b1 := &ht.data[i]
		b2 := &ht.data[j]
		p1 := MakeArrayPair(b1.GetArrayKey(), b1.GetVal())
		p2 := MakeArrayPair(b2.GetArrayKey(), b2.GetVal())
		ret := comparer(p1, p2)
		return ret < 0
	})

	if renumber {
		ht.eachBucket(func(pos uint32, p *Bucket) {
			p.SetIndexKey(int(pos))
		})
		ht.nextFreeElement = len(ht.data)
	}

	ht.rehash()
}

func (ht *ArrayDataHt) GracefulReverseDestroy() {
	ht.assertWritable()
	for idx := ht.DataSize(); idx > 0; idx-- {
		pos := idx - 1
		p := &ht.data[pos]
		if p.IsValid() {
			ht.deleteBucket(idx)
		}
	}
}

/**
 * each
 */
func (ht *ArrayDataHt) Foreach(handler func(key ArrayKey, value *Zval)) {
	for i, _ := range ht.data {
		p := &ht.data[i]
		if p.IsValid() {
			continue
		}
		handler(p.GetArrayKey(), p.GetVal())
	}
}
func (ht *ArrayDataHt) ForeachEx(handler func(key ArrayKey, value *Zval) bool) bool {
	for i, _ := range ht.data {
		p := &ht.data[i]
		if p.IsValid() {
			continue
		}
		if !handler(p.GetArrayKey(), p.GetVal()) {
			return false
		}
	}
	return true
}
func (ht *ArrayDataHt) ForeachReserve(handler func(key ArrayKey, value *Zval)) {
	for i := len(ht.data) - 1; i >= 0; i-- {
		p := &ht.data[i]
		if p.IsValid() {
			continue
		}
		handler(p.GetArrayKey(), p.GetVal())
	}
}

func (ht *ArrayDataHt) ForeachIndirect(handler func(key ArrayKey, value *Zval)) {
	for i, _ := range ht.data {
		p := &ht.data[i]
		data := p.GetVal()
		if data.IsIndirect() {
			data = data.Indirect()
		}
		if data.IsUndef() {
			continue
		}
		handler(p.GetArrayKey(), data)
	}
}
func (ht *ArrayDataHt) ForeachIndirectEx(handler func(key ArrayKey, value *Zval) bool) bool {
	for i, _ := range ht.data {
		p := &ht.data[i]
		data := p.GetVal()
		if data.IsIndirect() {
			data = data.Indirect()
		}
		if data.IsUndef() {
			continue
		}
		if !handler(p.GetArrayKey(), p.GetVal()) {
			return false
		}
	}
	return true
}
func (ht *ArrayDataHt) ForeachIndirectReserve(handler func(key ArrayKey, value *Zval)) {
	for i := len(ht.data) - 1; i >= 0; i-- {
		p := &ht.data[i]
		data := p.GetVal()
		if data.IsIndirect() {
			data = data.Indirect()
		}
		if data.IsUndef() {
			return
		}
		handler(p.GetArrayKey(), data)
	}
}

func (ht *ArrayDataHt) eachBucket(handler func(pos uint32, p *Bucket)) {
	for i, _ := range ht.data {
		p := &ht.data[i]
		handler(uint32(i), p)
	}
}
func (ht *ArrayDataHt) eachValidBucket(handler func(pos uint32, p *Bucket)) {
	for i, _ := range ht.data {
		p := &ht.data[i]
		if p.IsValid() {
			continue
		}
		handler(uint32(i), p)
	}
}
func (ht *ArrayDataHt) EachValidBucketIndirect(handler func(pos uint32, p *Bucket, data *Zval)) {
	for i, _ := range ht.data {
		p := &ht.data[i]
		data := p.GetVal()
		if data.IsIndirect() {
			data = data.Indirect()
		}
		if data.IsUndef() {
			return
		}
		handler(uint32(i), p, data)
	}
}

/**
 * Iterator & Pos
 */
// 查找从当前 pos 开始第一个有效 pos(含当前pos)
func (ht *ArrayDataHt) validPosEx(pos uint32, indirect bool) (uint32, bool) {
	dataSize := uint32(len(ht.data))
	for i := pos; i < dataSize; i++ {
		val := ht.data[i].GetVal()
		if indirect && val.IsIndirect() {
			val = val.Indirect()
		}
		if val.IsUndef() {
			continue
		}
		return i, true
	}
	// 没有有效pos，此时 pos == ht.DataSize()
	return pos, false
}

func (ht *ArrayDataHt) currentPos() (uint32, bool) {
	return ht.validPosEx(ht.internalPointer, false)
}

func (ht *ArrayDataHt) CurrentPosVal() uint32 {
	var pos, _ = ht.validPosEx(ht.internalPointer, false)
	return pos
}

func (ht *ArrayDataHt) currentPosVal() uint32 {
	var pos, _ = ht.validPosEx(ht.internalPointer, false)
	return pos
}

func (ht *ArrayDataHt) validPos(pos uint32) (uint32, bool) {
	return ht.validPosEx(pos, false)
}

func (ht *ArrayDataHt) validPosVal(pos uint32) uint32 {
	pos, _ = ht.validPosEx(pos, false)
	return pos
}

/**
 * Internal methods
 */
func (ht *ArrayDataHt) copyDataAndHash(source *Array) {
	ht.data = b.CopySlice(source.data)
	ht.indexes = b.CopyMap(source.indexes)
}

// 移除 this.data 数据中的 holes, 返回是否移动 bucket
func (ht *ArrayDataHt) removeHoles() {
	ht.assertWritable()

	var newPos uint32 = 0

	if len(ht.data) == int(ht.elementsCount) {
		return
	}

	ht.eachValidBucket(func(pos uint32, p *Bucket) {
		if newPos != pos {
			ht.data[newPos] = ht.data[pos]
			if ht.internalPointer == pos {
				ht.internalPointer = newPos
			}
		}
		newPos++
	})

	// 截取数据，记录有效元素数
	ht.data = ht.data[:newPos]
	ht.elementsCount = newPos
}

// 移除 data 的 holes, 不考虑 internalPointer 和 Iterators 内的 pos 指针
func (ht *ArrayDataHt) removeHolesAndCleanInternalPointer() bool {
	ht.removeHoles()
	ht.internalPointer = 0
	return true
}

func (ht *ArrayDataHt) MoveTailToHead() {
	var tmp Bucket = ht.data[len(ht.data)-1]
	copy(ht.data[1:], ht.data)
	ht.data[0] = tmp
	ht.rehash()
}
