package types

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
	"sort"
)

/**
 * ArrayKey
 * 	- 可直接比较
 *  - 零值为数字 0
 * 为减少内存占用，省略类型标识字段，采用以下方式确认类型:
 * - str == ""        时: int 类型，值为 idx
 * - str[0] != '\x00' 时: string 类型，值为 str
 * - str[0] == '\x00' 时: string 类型，值为 str[1:]
 */
type ArrayKey struct {
	idx int
	str string
}

func IdxKey(index int) ArrayKey { return ArrayKey{index, ""} }
func StrKey(str string) ArrayKey {
	if str == "" || str[0] == '\x00' {
		str = "\x00" + str
	}
	return ArrayKey{0, str}
}

func (k ArrayKey) IsStrKey() bool { return k.str != "" }
func (k ArrayKey) IdxKey() int    { return k.idx }
func (k ArrayKey) StrKey() string {
	if k.str != "" && k.str[0] == '\x00' {
		return k.str[1:]
	} else {
		return k.str
	}
}
func (k ArrayKey) Keys() (index int, key string, isStrKey bool) {
	return k.IdxKey(), k.StrKey(), k.IsStrKey()
}
func (k ArrayKey) ToZval() *Zval {
	if k.IsStrKey() {
		return NewZvalString(k.StrKey())
	} else {
		return NewZvalLong(k.IdxKey())
	}
}

/**
 * ArrayPair
 */
type ArrayPair struct {
	key ArrayKey
	val *Zval
}

func MakeArrayPair(key ArrayKey, val *Zval) ArrayPair {
	return ArrayPair{key: key, val: val}
}
func NewArrayPair(key ArrayKey, val *Zval) *ArrayPair {
	return &ArrayPair{key: key, val: val}
}
func (p ArrayPair) GetKey() ArrayKey { return p.key }
func (p ArrayPair) GetVal() *Zval    { return p.val }

/**
 * Bucket
 */
type Bucket struct {
	val Zval
	key ArrayKey
}

func NewBucket(key ArrayKey, zval *Zval) *Bucket {
	var bucket = &Bucket{key: key}
	ZVAL_COPY_VALUE(&bucket.val, zval)
	return bucket
}

func (this *Bucket) IsStrKey() bool            { return this.key.IsStrKey() }
func (this *Bucket) StrKey() string            { return this.key.StrKey() }
func (this *Bucket) IndexKey() int             { return this.key.IdxKey() }
func (this *Bucket) Keys() (int, string, bool) { return this.key.Keys() }

func (this *Bucket) SetStrKey(key string)  { this.key = StrKey(key) }
func (this *Bucket) SetIndexKey(index int) { this.key = IdxKey(index) }
func (this *Bucket) GetArrayKey() ArrayKey { return this.key }

func (this *Bucket) GetVal() *Zval     { return &this.val }
func (this *Bucket) SetVal(zval *Zval) { ZVAL_COPY_VALUE(&this.val, zval) }

func (this *Bucket) IsValid() bool { return !this.val.IsUndef() }

func (this *Bucket) CopyFrom(from *Bucket) {
	this.SetVal(from.GetVal())
	this.key = from.key
}

/**
 * Array
 * HashTable Data Layout
 * =====================
 *
 *                 +=============================+
 *                 | HT_HASH(ht, ht->nTableMask) |
 *                 | ...                         |
 *                 | HT_HASH(ht, -1)             |
 *                 +-----------------------------+
 * ht->arData ---> | Bucket[0]                   |
 *                 | ...                         |
 *                 | Bucket[ht->tableSize-1]    |
 *                 +=============================+
 */
type ArrayPosition = uint32
type Array struct {
	ZendRefcounted
	flags           uint8
	elementsCount   uint32
	internalPointer uint32
	nextFreeElement int

	data    []Bucket            // 实际存储数据的地方
	indexes map[ArrayKey]uint32 // 索引到具体位置的映射
	arData  *Bucket             // C 源码中存储数据的地方，实际不使用

	data0 ArrayData
}

var _ IRefcounted = &Array{}

/**
 * Constructor && Init
 */
func NewArray(size int) *Array {
	var ht = new(Array)
	ht.Init(size)
	return ht
}
func (ht *Array) Init(size int) {
	var data []Bucket
	if size > 0 {
		data = make([]Bucket, 0, size)
	}

	*ht = Array{
		// 数据存储
		data:    data,
		indexes: make(map[ArrayKey]uint32), // todo 改为 nil，延迟初始化
	}

	// GC 信息
	ht.SetRefcount(1)
	ht.SetGcTypeInfo(uint32(IS_ARRAY))
}

/* init */
func (ht *Array) SetBy(arr *Array) {
	ht.flags = arr.flags
	ht.elementsCount = arr.elementsCount
	ht.nextFreeElement = arr.nextFreeElement

	ht.data = arr.data
	ht.indexes = arr.indexes

	ZendHashInternalPointerReset(ht)
}

// 实际元素个数，从使用者角度的数组大小
func (ht *Array) Len() int                  { return ht.data0.Len() }
func (ht *Array) Cap() int                  { return ht.data0.Cap() }
func (ht *Array) Bucket(pos uint32) *Bucket { return &ht.data[pos] }

/* data -> Array.data */
func (ht *Array) clearData() {
	ht.assertRc1()

	ht.elementsCount = 0
	ht.internalPointer = 0
	ht.nextFreeElement = 0
	ht.data = nil
	ht.indexes = make(map[ArrayKey]uint32)
}
func (ht *Array) deleteBucket(pos uint32) {
	ht.assertRc1()
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

func (ht *Array) posBucket(p *Bucket) (uint32, bool) {
	if pos, ok := ht.indexes[p.key]; ok {
		return pos, true
	} else {
		return 0, false
	}
}

// 移动 bucket 到新位置
func (ht *Array) moveBucket(pos uint32, newPos uint32) {
	assert(newPos <= pos)
	if newPos == pos {
		return
	}
	(&ht.data[newPos]).CopyFrom(&ht.data[pos])
	if ht.internalPointer == pos {
		ht.internalPointer = newPos
	}
}

/* misc */
func (ht *Array) assertRc1()      { assert(ht.GetRefcount() == 1) }
func (ht *Array) assertWritable() { assert(ht.GetRefcount() == 1) }

/** Array.flags */
func (ht *Array) CopyFlags(arr *Array) { ht.flags = arr.flags }

func (ht *Array) IsPacked() bool       { return ht.flags&HASH_FLAG_PACKED != 0 }
func (ht *Array) HasEmptyIndex() bool  { return ht.flags&HASH_FLAG_HAS_EMPTY_IND != 0 }
func (ht *Array) MarkHasEmptyIndex()   { ht.flags |= HASH_FLAG_HAS_EMPTY_IND }
func (ht *Array) UnmarkHasEmptyIndex() { ht.flags &^= HASH_FLAG_HAS_EMPTY_IND }

func (ht *Array) Keys() []ArrayKey {
	var keys = make([]ArrayKey, 0, ht.Len())
	ht.Foreach(func(key ArrayKey, value *Zval) {
		keys = append(keys, key)
	})
	return keys
}
func (ht *Array) Values() []*Zval {
	var values = make([]*Zval, 0, ht.Len())
	ht.Foreach(func(key ArrayKey, value *Zval) {
		values = append(values, value)
	})
	return values
}
func (ht *Array) Pairs() []ArrayPair {
	var pairs = make([]ArrayPair, 0, ht.Len())
	ht.Foreach(func(key ArrayKey, value *Zval) {
		pairs = append(pairs, MakeArrayPair(key, value))
	})
	return pairs
}

func (ht *Array) MapWithKey(mapper func(key ArrayKey, value *Zval) (ArrayKey, *Zval)) *Array {
	// todo 考虑 rehash 等操作 或 对其他属性的处理
	arr := NewArray(ht.Len())
	ht.Foreach(func(key ArrayKey, value *Zval) {
		newKey, newValue := mapper(key, value)
		arr.Add(newKey, newValue)
	})
	return arr
}

/**
 * Open methods
 */
func (ht *Array) GetArData() *Bucket             { return ht.arData }
func (ht *Array) DataSize() uint32               { return uint32(len(ht.data)) }
func (ht *Array) GetNNumUsed() uint32            { return ht.DataSize() }
func (ht *Array) SetNNumOfElements(value uint32) { ht.elementsCount = value }
func (ht *Array) GetNInternalPointer() uint32    { return ht.internalPointer }
func (ht *Array) GetNNextFreeElement() int       { return ht.nextFreeElement }
func (ht *Array) SetNNextFreeElement(value int)  { ht.nextFreeElement = value }

func (ht *Array) Count() int {
	var num int
	if ht.HasEmptyIndex() {
		num = ht.recalcElements()
		if int(ht.elementsCount) == num {
			ht.UnmarkHasEmptyIndex()
		}
	} else if ht == zend.EG__().GetSymbolTable() { // todo
		num = ht.recalcElements()
	} else {
		num = ht.Len()
	}
	return num
}

// 重新计算有效元素个数(与 elementsCount 不同，它需要过滤 IS_INDIRECT 元素为 IS_UNDEF 的情况)
func (ht *Array) recalcElements() int {
	var num = 0
	ht.ForeachIndirect(func(key ArrayKey, value *Zval) {
		num++
	})
	return num
}

func (ht *Array) FirstPair() *ArrayPair {
	for _, p := range ht.data {
		if p.IsValid() {
			return NewArrayPair(p.GetArrayKey(), p.GetVal())
		}
	}
	return nil
}

func (ht *Array) First() (key ArrayKey, val *Zval) {
	for _, p := range ht.data {
		if p.IsValid() {
			return p.GetArrayKey(), p.GetVal()
		}
	}
	return
}

func (ht *Array) FirstIndirect() (key ArrayKey, val *Zval) {
	for _, p := range ht.data {
		v := p.GetVal()
		if v.IsIndirect() {
			v = v.Indirect()
		}
		if v.IsUndef() {
			continue
		}
		return p.GetArrayKey(), v
	}
	return
}

func (ht *Array) LastPair() *ArrayPair {
	for i := len(ht.data) - 1; i >= 0; i-- {
		p := ht.data[i]
		if p.IsValid() {
			return NewArrayPair(p.GetArrayKey(), p.GetVal())
		}
	}
	return nil
}

func (ht *Array) LastPairIndirect() *ArrayPair {
	for i := len(ht.data) - 1; i >= 0; i-- {
		p := ht.data[i]
		v := p.GetVal()
		if v.IsIndirect() {
			v = v.Indirect()
		}
		if v.IsUndef() {
			continue
		}
		return NewArrayPair(p.GetArrayKey(), v)
	}
	return nil
}

func (ht *Array) Last() (key ArrayKey, val *Zval) {
	for i := len(ht.data) - 1; i >= 0; i-- {
		p := ht.data[i]
		if p.IsValid() {
			return p.GetArrayKey(), p.GetVal()
		}
	}
	return
}

func (ht *Array) LastIndirect() (key ArrayKey, val *Zval) {
	for i := len(ht.data) - 1; i >= 0; i-- {
		p := ht.data[i]
		v := p.GetVal()
		if v.IsIndirect() {
			v = v.Indirect()
		}
		if v.IsUndef() {
			continue
		}
		return p.GetArrayKey(), v
	}
	return
}

/**
 * Sort
 */
type ArrayLessComparer func(p1, p2 ArrayPair) bool
type ArrayComparer func(p1, p2 ArrayPair) int

func (ht *Array) Sort(comparer ArrayComparer, renumber bool) {
	ht.assertRc1()

	if ht.elementsCount == 0 || (ht.elementsCount == 1 && !renumber) {
		return
	}

	ht.internalPointer = 0
	pairs := ht.Pairs()
	sort.SliceStable(pairs, func(i, j int) bool {
		ret := comparer(pairs[i], pairs[j])
		return ret < 0
	})

	if renumber {
		// todo 考虑生成 []*Zval 直接转 ArrayData
		for i, _ := range pairs {
			pairs[i].key = IdxKey(i)
		}
	}

	ht.data0 = ht.newDataOfPairs(pairs)
}
func (ht *Array) newDataOfPairs(pairs []ArrayPair) ArrayData {
	// todo
	panic("todo")
}

func (ht *Array) Min(comparer ArrayComparer) *ArrayPair {
	if ht.Len() == 0 {
		return nil
	}

	var res *ArrayPair
	ht.Foreach(func(key ArrayKey, value *Zval) {
		pair := MakeArrayPair(key, value)
		if res == nil || comparer(*res, pair) > 0 {
			res = &pair
		}
	})

	return res
}
func (ht *Array) Max(comparer ArrayComparer) *ArrayPair {
	if ht.Len() == 0 {
		return nil
	}

	var res *ArrayPair
	ht.Foreach(func(key ArrayKey, value *Zval) {
		pair := MakeArrayPair(key, value)
		if res == nil || comparer(*res, pair) < 0 {
			res = &pair
		}
	})
	return res
}

/**
 * Clean && Destroy
 */
func (ht *Array) Clean() {
	ht.assertRc1()
	ht.clearData()
}

func (ht *Array) Destroy()   { ht.Clean() }
func (ht *Array) DestroyEx() { ht.Clean() }

/**
 * Methods use index key
 */
func (ht *Array) IndexExists(index int) bool               { return ht.Exists(IdxKey(index)) }
func (ht *Array) IndexFind(index int) *Zval                { return ht.Find(IdxKey(index)) }
func (ht *Array) IndexAdd(index int, pData *Zval) *Zval    { return ht.Add(IdxKey(index), pData) }
func (ht *Array) IndexAddNew(index int, pData *Zval) *Zval { return ht.AddNew(IdxKey(index), pData) }
func (ht *Array) IndexUpdate(index int, pData *Zval) *Zval { return ht.Update(IdxKey(index), pData) }
func (ht *Array) IndexDelete(index int) bool               { return ht.Delete(IdxKey(index)) }

/**
 * Methods use string key
 */
func (ht *Array) KeyExists(key string) bool               { return ht.Exists(StrKey(key)) }
func (ht *Array) KeyFind(key string) *Zval                { return ht.Find(StrKey(key)) }
func (ht *Array) KeyAdd(key string, pData *Zval) *Zval    { return ht.Add(StrKey(key), pData) }
func (ht *Array) KeyAddNew(key string, pData *Zval) *Zval { return ht.AddNew(StrKey(key), pData) }
func (ht *Array) KeyUpdate(key string, pData *Zval) *Zval { return ht.Update(StrKey(key), pData) }
func (ht *Array) KeyDelete(key string) bool               { return ht.Delete(StrKey(key)) }

func (ht *Array) KeyFindPtr(key string) any {
	var zv = ht.KeyFind(key)
	if zv != nil {
		return zv.Ptr()
	}
	return nil
}
func (ht *Array) KeyExistsIndirect(key string) bool {
	var zv = ht.KeyFind(key)
	if zv == nil {
		return false
	}

	if zv.IsUndef() && zv.Indirect().IsUndef() {
		return false
	}

	return true
}
func (ht *Array) KeyAddIndirect(key string, pData *Zval) *Zval {
	ht.assertRc1()

	if data := ht.KeyFind(key); data != nil {
		b.Assert(data != pData)
		if data.IsIndirect() {
			data = data.Indirect()
			if !data.IsUndef() {
				return nil
			}
		} else {
			return nil
		}
		data.CopyValueFrom(pData)
		return data
	}

	return ht.Add(StrKey(key), pData)
}
func (ht *Array) KeyUpdateIndirect(key string, pData *Zval) *Zval {
	ht.assertRc1()

	if data := ht.KeyFind(key); data != nil {
		b.Assert(data != pData)
		if data.IsType(IS_INDIRECT) {
			data = data.Indirect()
		}
		data.CopyValueFrom(pData)
		return data
	}

	return ht.Update(StrKey(key), pData)
}
func (ht *Array) KeyDeleteIndirect(key string) bool {
	ht.assertRc1()
	if pos, ok := ht.indexes[StrKey(key)]; ok {
		var p = &ht.data[pos]
		if p.GetVal().IsType(IS_INDIRECT) {
			var data *Zval = p.GetVal().Indirect()
			if data.IsType(IS_UNDEF) {
				return false
			} else {
				data.SetUndef()
				ht.MarkHasEmptyIndex()
			}
		} else {
			ht.deleteBucket(pos)
		}
		return true
	}
	return false
}

/**
 * Add / Update by ArrayKey
 */

func (ht *Array) Exists(key ArrayKey) bool { return ht.data0.Exists(key) }
func (ht *Array) Find(key ArrayKey) *Zval  { return ht.data0.Find(key) }
func (ht *Array) Add(key ArrayKey, pData *Zval) *Zval {
	ht.assertRc1()
	ok := ht.data0.Add(key, pData)
	if !ok {
		return nil
	}
	return ht.Find(key)
}

func (ht *Array) AddNew(key ArrayKey, pData *Zval) *Zval {
	// todo 此操作要求提前确认 key 不冲突
	ht.assertRc1()
	return ht.Add(key, pData)
}

func (ht *Array) AddIndirect(key ArrayKey, pData *Zval) *Zval {
	if key.IsStrKey() {
		return ht.KeyAddIndirect(key.StrKey(), pData)
	} else {
		return ht.IndexAdd(key.IdxKey(), pData)
	}
}

func (ht *Array) Update(key ArrayKey, pData *Zval) *Zval {
	ht.assertRc1()
	ht.data0.Update(key, pData)
	return ht.data0.Find(key)
}

func (ht *Array) UpdateIndirect(key ArrayKey, pData *Zval) *Zval {
	if key.IsStrKey() {
		return ht.KeyUpdateIndirect(key.StrKey(), pData)
	} else {
		return ht.IndexUpdate(key.IdxKey(), pData)
	}
}

func (ht *Array) Delete(key ArrayKey) bool {
	ht.assertRc1()
	return ht.data0.Delete(key)
}

func (ht *Array) Append(pData *Zval) *Zval {
	ht.assertRc1()
	idx := ht.data0.Push(pData)
	return ht.IndexFind(idx)
}
func (ht *Array) AppendNew(pData *Zval) *Zval {
	// todo 此操作要求提前确认 key 不冲突
	ht.assertRc1()
	return ht.Append(pData)
}

/**
 * each
 */
func (ht *Array) Foreach(handler func(key ArrayKey, value *Zval)) {
	for i, _ := range ht.data {
		p := &ht.data[i]
		if p.IsValid() {
			continue
		}
		handler(p.GetArrayKey(), p.GetVal())
	}
}
func (ht *Array) ForeachEx(handler func(key ArrayKey, value *Zval) bool) bool {
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
func (ht *Array) ForeachReserve(handler func(key ArrayKey, value *Zval)) {
	for i := len(ht.data) - 1; i >= 0; i-- {
		p := &ht.data[i]
		if p.IsValid() {
			continue
		}
		handler(p.GetArrayKey(), p.GetVal())
	}
}

func (ht *Array) ForeachIndirect(handler func(key ArrayKey, value *Zval)) {
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
func (ht *Array) ForeachIndirectEx(handler func(key ArrayKey, value *Zval) bool) bool {
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
func (ht *Array) ForeachIndirectReserve(handler func(key ArrayKey, value *Zval)) {
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

func (ht *Array) Filter(handler func(key ArrayKey, value *Zval) bool) bool {
	var removeKeys []ArrayKey
	for i, _ := range ht.data {
		p := &ht.data[i]
		if p.IsValid() {
			continue
		}
		if !handler(p.GetArrayKey(), p.GetVal()) {
			removeKeys = append(removeKeys, p.GetArrayKey())
		}
	}

	if len(removeKeys) > 0 {
		for _, key := range removeKeys {
			ht.Delete(key)
		}
	}

	return true
}

func (ht *Array) Iterator() *ArrayIterator {
	htReader := ht.Copy()
	return &ArrayIterator{ht: htReader, pos: 0}
}

func (ht *Array) IteratorEx(pos uint32) *ArrayIterator {
	htReader := ht.Copy()
	return &ArrayIterator{ht: htReader, pos: pos}
}

// todo 逐渐替换为 Foreach 或其他更高效代码
func (ht *Array) ForeachData() []*Bucket {
	panic("todo replace")
}

func (ht *Array) EachValidBucketIndirect(handler func(pos uint32, p *Bucket, data *Zval)) {
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

func (ht *Array) applyValidBucket(apply_func func(p *Bucket) int) {
	for i, _ := range ht.data {
		p := &ht.data[i]
		if p.IsValid() {
			continue
		}
		result := apply_func(p)
		if b.FlagMatch(result, ArrayApplyRemove) {
			ht.deleteBucket(uint32(i))
		}
		if b.FlagMatch(result, ArrayApplyStop) {
			break
		}
	}
}
func (ht *Array) applyValidBucketReserve(apply_func func(p *Bucket) int) {
	for i := len(ht.data) - 1; i >= 0; i-- {
		p := &ht.data[i]
		if p.IsValid() {
			continue
		}
		result := apply_func(p)
		if b.FlagMatch(result, ArrayApplyRemove) {
			ht.deleteBucket(uint32(i))
		}
		if b.FlagMatch(result, ArrayApplyStop) {
			break
		}
	}
}

/**
 * Iterator & Pos
 */
// 查找从当前 pos 开始第一个有效 pos(含当前pos)
func (ht *Array) validPosEx(pos uint32, indirect bool) (uint32, bool) {
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

func (ht *Array) currentPos() (uint32, bool) {
	return ht.validPosEx(ht.internalPointer, false)
}

func (ht *Array) CurrentPosVal() uint32 {
	var pos, _ = ht.validPosEx(ht.internalPointer, false)
	return pos
}

func (ht *Array) currentPosVal() uint32 {
	var pos, _ = ht.validPosEx(ht.internalPointer, false)
	return pos
}

func (ht *Array) validPos(pos uint32) (uint32, bool) {
	return ht.validPosEx(pos, false)
}

func (ht *Array) validPosVal(pos uint32) uint32 {
	pos, _ = ht.validPosEx(pos, false)
	return pos
}

/**
 * Internal methods
 */
func (ht *Array) copyData0(source *Array) {
	// todo 待处理复制逻辑(参考逻辑: 建立一个reader，延迟复制，考虑内部指针)
	ht.data0 = source.data0
	ht.data = b.CopySlice(source.data)
	ht.indexes = b.CopyMap(source.indexes)
}

func (ht *Array) dupData0(source *Array) {
	// todo 待处理复制逻辑(参考逻辑: 非延迟复制，考虑内部指针)
	source.Foreach(func(key ArrayKey, value *Zval) {
		ht.Update(key, value)
	})

	ht.data0 = source.data0
	ht.data = b.CopySlice(source.data)
	ht.indexes = b.CopyMap(source.indexes)
}

func (ht *Array) MoveTailToHead() {
	if ht.Len() <= 1 {
		return
	}

	pairs := ht.Pairs()
	tmp := pairs[len(pairs)-1]
	copy(pairs[1:], pairs)
	pairs[0] = tmp

	ht.data0 = ht.newDataOfPairs(pairs)
}

func (ht *Array) Push(value *Zval) {
	ht.Append(value)
}

func (ht *Array) Copy() *Array {
	return ArrayLazyDup(ht)
}

func (ht *Array) CopyEx(cap int) *Array {
	return ArrayLazyDup(ht)
}

func (ht *Array) ResetInternalPointer() {
	ht.internalPointer = ht.validPosVal(0)
}
