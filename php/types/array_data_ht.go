package types

import (
	"github.com/heyuuu/gophp/kits/mapkit"
	"sort"
)

// bucket
type bucket struct {
	key ArrayKey
	val Zval
}

func makeBucket(key ArrayKey, val *Zval) bucket {
	return bucket{key: key, val: *val}
}

func (bkt *bucket) Key() ArrayKey       { return bkt.key }
func (bkt *bucket) SetKey(key ArrayKey) { bkt.key = key }
func (bkt *bucket) Val() *Zval          { return &bkt.val }
func (bkt *bucket) SetVal(val *Zval)    { bkt.val = *val }
func (bkt *bucket) IsValid() bool       { return !bkt.val.IsUndef() }
func (bkt *bucket) MarkInvalid()        { bkt.val.SetUndef() }

// ArrayDataHt
type ArrayDataHt struct {
	elementsCount   int
	nextFreeElement int
	data            []bucket         // 实际存储数据的地方
	indexes         map[ArrayKey]int // 索引到具体位置的映射
	writable        bool
}

func newArrayDataHt(cap int) *ArrayDataHt {
	return &ArrayDataHt{
		indexes:  make(map[ArrayKey]int, cap),
		data:     make([]bucket, 0, cap),
		writable: true,
	}
}

func (ht *ArrayDataHt) Clone() ArrayData {
	//TODO implement me
	panic("implement me")
}

func (ht *ArrayDataHt) Len() int  { return ht.elementsCount }
func (ht *ArrayDataHt) Used() int { return len(ht.data) }
func (ht *ArrayDataHt) Cap() int  { return cap(ht.data) }
func (ht *ArrayDataHt) Count() int {
	var num = 0
	for i, _ := range ht.data {
		if !ht.data[i].IsValid() {
			continue
		}
		num++
	}
	return num
}

func (ht *ArrayDataHt) Exists(key ArrayKey) bool {
	_, ok := ht.indexes[key]
	return ok
}
func (ht *ArrayDataHt) Find(key ArrayKey) (*Zval, ArrayPosition) {
	if pos, ok := ht.indexes[key]; ok {
		return ht.data[pos].Val(), pos
	}
	return nil, InvalidArrayPos
}
func (ht *ArrayDataHt) Each(handler func(key ArrayKey, value *Zval) error) error {
	for i, _ := range ht.data {
		bkt := &ht.data[i]
		if !bkt.IsValid() {
			continue
		}
		err := handler(bkt.Key(), bkt.Val())
		if err != nil {
			return err
		}
	}
	return nil
}
func (ht *ArrayDataHt) EachReserve(handler func(key ArrayKey, value *Zval) error) error {
	for i := len(ht.data) - 1; i >= 0; i-- {
		bkt := &ht.data[i]
		if !bkt.IsValid() {
			continue
		}
		err := handler(bkt.Key(), bkt.Val())
		if err != nil {
			return err
		}
	}
	return nil
}
func (ht *ArrayDataHt) Pos(pos ArrayPosition) (key ArrayKey, value *Zval) {
	if pos < 0 || pos >= len(ht.data) {
		return
	}

	bkt := &ht.data[pos]
	if !bkt.IsValid() {
		return
	}

	return bkt.Key(), bkt.Val()
}

func (ht *ArrayDataHt) Add(key ArrayKey, value *Zval) (bool, error) {
	ht.assertWritable()
	if _, exists := ht.indexes[key]; exists {
		return false, nil
	} else {
		ht.appendBucket(key, value)
		return true, nil
	}
}
func (ht *ArrayDataHt) Update(key ArrayKey, value *Zval) error {
	ht.assertWritable()
	if pos, exists := ht.indexes[key]; exists {
		ht.data[pos].SetVal(value)
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
func (ht *ArrayDataHt) Append(value *Zval) (int, error) {
	ht.assertWritable()
	idx := ht.nextFreeElement
	key := IdxKey(idx)
	assert(!ht.Exists(key))
	ht.appendBucket(key, value)
	return idx, nil
}

func (ht *ArrayDataHt) Sort(comparer ArrayComparer, renumber bool) {
	ht.assertWritable()

	if ht.elementsCount == 0 || (ht.elementsCount == 1 && !renumber) {
		return
	}

	ht.removeHoles()

	sort.SliceStable(ht.data, func(i, j int) bool {
		bkt1 := &ht.data[i]
		bkt2 := &ht.data[j]
		ret := comparer.Compare(bkt1.Key(), bkt1.Val(), bkt2.Key(), bkt2.Val())
		return ret < 0
	})

	if renumber {
		for pos, _ := range ht.data {
			ht.data[pos].SetKey(IdxKey(pos))
		}
		ht.nextFreeElement = len(ht.data)
	}

	ht.rehash()
}

func (ht *ArrayDataHt) assertWritable() { assert(ht.writable) }

func (ht *ArrayDataHt) appendBucket(key ArrayKey, value *Zval) {
	// 尝试 resize
	ht.resizeIfFull()

	// 添加数据，更新元素计数
	ht.elementsCount++
	ht.indexes[key] = len(ht.data)
	ht.data = append(ht.data, makeBucket(key, value))

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
	assert(0 <= pos && pos < len(ht.data))

	bkt := &ht.data[pos]
	assert(bkt.IsValid())

	// 移除数据，更新元素计数
	ht.elementsCount--
	delete(ht.indexes, bkt.Key())
	bkt.MarkInvalid()

	// 若删除队尾元素，尝试清除 data 队尾无用数据
	if pos == len(ht.data)-1 {
		newDataSize := len(ht.data) - 1
		for newDataSize > 0 && !ht.data[newDataSize-1].IsValid() {
			newDataSize--
		}

		ht.data = ht.data[:newDataSize]
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
		key := ht.data[pos].Key()
		ht.indexes[key] = pos
	}
}

// 移除 this.data 数据中的 holes
func (ht *ArrayDataHt) removeHoles() {
	ht.assertWritable()

	if len(ht.data) == ht.elementsCount {
		return
	}

	newPos := 0
	for pos, _ := range ht.data {
		if !ht.data[pos].IsValid() {
			continue
		}
		if newPos != pos {
			ht.data[newPos] = ht.data[pos]
		}
		newPos++
	}

	// 截取数据，记录有效元素数
	ht.data = ht.data[:newPos]
	ht.elementsCount = newPos
}
