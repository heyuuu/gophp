package types

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
	"math"
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
func (k ArrayKey) Keys() (idx int, key string, isStrKey bool) {
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
func (p ArrayPair) GetKey() ArrayKey        { return p.key }
func (p ArrayPair) GetVal() *Zval           { return p.val }
func (p ArrayPair) Pair() (ArrayKey, *Zval) { return p.key, p.val }

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
	flags           uint8
	elementsCount   uint32
	internalPointer uint32
	nextFreeElement int

	data []Bucket // 实际存储数据的地方

	data0 ArrayData

	// flags todo 待合并
	protected bool
	writable  bool
}

/**
 * Constructor && Init
 */
func NewArray() *Array {
	return NewArrayCap(0)
}
func NewArrayCap(cap int) *Array {
	var data []Bucket
	if cap > 0 {
		data = make([]Bucket, 0, cap)
	}

	var ht = &Array{
		// 数据存储
		data: data,
	}

	return ht
}

// 实际元素个数，从使用者角度的数组大小
func (ht *Array) Len() int                  { return ht.data0.Len() }
func (ht *Array) Cap() int                  { return ht.data0.Cap() }
func (ht *Array) Bucket(pos uint32) *Bucket { return &ht.data[pos] }

/* misc */
func (ht *Array) assertWritable() { assert(ht.writable) }

/** Array.flags */
func (ht *Array) CopyFlags(arr *Array) { ht.flags = arr.flags }
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
	arr := NewArrayCap(ht.Len())
	ht.Foreach(func(key ArrayKey, value *Zval) {
		newKey, newValue := mapper(key, value)
		arr.Add(newKey, newValue)
	})
	ht.ResetInternalPointer()
	return arr
}

/**
 * Open methods
 */
const maxArrayPosition uint32 = math.MaxUint32

func (ht *Array) KeyFindValAndPos(strKey string) (*Zval, uint32) {
	key := StrKey(strKey)
	return ht.data0.FindEx(key)
}
func (ht *Array) KeyAddValAndPos(strKey string, pData *Zval) (*Zval, uint32) {
	key := StrKey(strKey)
	ok := ht.data0.Add(key, pData)
	if !ok {
		return nil, maxArrayPosition
	}
	return ht.data0.FindEx(key)
}
func (ht *Array) KeyUpdateValAndPos(strKey string, pData *Zval) (*Zval, uint32) {
	key := StrKey(strKey)
	ht.data0.Update(key, pData)
	return ht.data0.FindEx(key)
}
func (ht *Array) PosValue(pos uint32) *Zval {
	p := ht.data0.Pos(pos)
	if p != nil {
		return p.GetVal()
	}
	return nil
}

func (ht *Array) GetNNumUsed() uint32         { return uint32(ht.Len()) }
func (ht *Array) GetNInternalPointer() uint32 { return ht.internalPointer }
func (ht *Array) GetNNextFreeElement() int    { return ht.nextFreeElement }

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

// Pos 相关
func (ht *Array) NextPos(pos ArrayPosition) ArrayPosition {
	if pos == maxArrayPosition {
		return maxArrayPosition
	}
	return pos + 1
}
func (ht *Array) PrevPos(pos ArrayPosition) ArrayPosition {
	if pos == 0 {
		return 0
	}
	return pos - 1
}

func (ht *Array) findPos(pos ArrayPosition) (*ArrayPair, ArrayPosition) {
	return ht.data0.FindPos(pos)
}
func (ht *Array) findPosIndirect(pos ArrayPosition) (*ArrayPair, ArrayPosition) {
	return ht.data0.FindPosIndirect(pos)
}
func (ht *Array) findPosReserve(pos ArrayPosition) (*ArrayPair, ArrayPosition) {
	return ht.data0.FindPosReserve(pos)
}
func (ht *Array) findPosReserveIndirect(pos ArrayPosition) (*ArrayPair, ArrayPosition) {
	return ht.data0.FindPosReserveIndirect(pos)
}

func (ht *Array) CurrentEx(pos ArrayPosition) (*ArrayPair, ArrayPosition) {
	return ht.findPos(pos)
}
func (ht *Array) NextEx(pos ArrayPosition) (*ArrayPair, ArrayPosition) {
	p, pos := ht.findPos(pos)
	return p, ht.NextPos(pos)
}
func (ht *Array) PrevEx(pos ArrayPosition) (*ArrayPair, ArrayPosition) {
	p, pos := ht.findPosReserve(pos)
	return p, ht.PrevPos(pos)
}

func (ht *Array) Current() *ArrayPair {
	var p *ArrayPair
	p, _ = ht.findPosIndirect(ht.internalPointer)
	return p
}
func (ht *Array) MoveNext() {
	_, ht.internalPointer = ht.NextEx(ht.internalPointer)
}
func (ht *Array) MovePrev() {
	_, ht.internalPointer = ht.PrevEx(ht.internalPointer - 1)
}
func (ht *Array) MoveFirst() {
	ht.internalPointer = ht.FirstPos()
}
func (ht *Array) MoveEnd() {
	ht.internalPointer = ht.LastPos()
}

func (ht *Array) FirstPos() ArrayPosition {
	_, pos := ht.findPos(0)
	return pos
}
func (ht *Array) First() *ArrayPair {
	pair, _ := ht.findPos(0)
	return pair
}
func (ht *Array) FirstIndirect() *ArrayPair {
	pair, _ := ht.findPosIndirect(0)
	return pair
}

func (ht *Array) LastPos() ArrayPosition {
	_, pos := ht.findPosReserve(ht.data0.MaxPos())
	return pos
}
func (ht *Array) Last() *ArrayPair {
	pair, _ := ht.findPosReserve(ht.data0.MaxPos())
	return pair
}

func (ht *Array) LastIndirect() *ArrayPair {
	pair, _ := ht.findPosReserveIndirect(ht.data0.MaxPos())
	return pair
}

/**
 * Sort
 */
type ArrayLessComparer func(p1, p2 ArrayPair) bool
type ArrayComparer func(p1, p2 ArrayPair) int

func (ht *Array) Sort(comparer ArrayComparer, renumber bool) {
	ht.assertWritable()

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
	ht.assertWritable()

	ht.elementsCount = 0
	ht.internalPointer = 0
	ht.nextFreeElement = 0

	// todo 确认 clean 逻辑
}

func (ht *Array) Destroy() { ht.Clean() }

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

func (ht *Array) KeyFindIndirect(key string) *Zval {
	var zv = ht.KeyFind(key)
	if zv != nil && zv.IsIndirect() {
		if zv.Indirect().IsNotUndef() {
			return zv.Indirect()
		} else {
			return nil
		}
	} else {
		return zv
	}
}

func (ht *Array) KeyAddIndirect(key string, pData *Zval) *Zval {
	ht.assertWritable()

	if data := ht.KeyFind(key); data != nil && data.IsIndirect() {
		b.Assert(data != pData)
		if !data.Indirect().IsUndef() {
			return nil
		}
		data.SetIndirect(pData)
		return data
	}

	return ht.Add(StrKey(key), pData)
}
func (ht *Array) KeyUpdateIndirect(key string, pData *Zval) *Zval {
	ht.assertWritable()

	if data := ht.KeyFind(key); data != nil && data.IsIndirect() {
		b.Assert(data != pData)
		data.SetIndirect(pData)
		return data
	}

	return ht.Update(StrKey(key), pData)
}
func (ht *Array) KeyDeleteIndirect(key string) bool {
	ht.assertWritable()

	if data := ht.KeyFind(key); data != nil && data.IsIndirect() {
		if data.Indirect().IsUndef() {
			return false
		}
		data.Indirect().SetUndef()
		ht.MarkHasEmptyIndex()
		return true
	}

	return ht.KeyDelete(key)
}

/**
 * Add / Update by ArrayKey
 */

func (ht *Array) Exists(key ArrayKey) bool { return ht.data0.Exists(key) }
func (ht *Array) Find(key ArrayKey) *Zval  { return ht.data0.Find(key) }
func (ht *Array) Add(key ArrayKey, pData *Zval) *Zval {
	ht.assertWritable()
	ok := ht.data0.Add(key, pData)
	if !ok {
		return nil
	}
	return ht.Find(key)
}

func (ht *Array) AddNew(key ArrayKey, pData *Zval) *Zval {
	// todo 此操作要求提前确认 key 不冲突
	ht.assertWritable()
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
	ht.assertWritable()
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
	ht.assertWritable()
	return ht.data0.Delete(key)
}

func (ht *Array) Append(pData *Zval) *Zval {
	ht.assertWritable()
	idx := ht.data0.Push(pData)
	return ht.IndexFind(idx)
}
func (ht *Array) AppendNew(pData *Zval) *Zval {
	// todo 此操作要求提前确认 key 不冲突
	ht.assertWritable()
	return ht.Append(pData)
}

/**
 * each
 */
func (ht *Array) Foreach(handler func(key ArrayKey, value *Zval)) {
	var pair *ArrayPair
	var pos ArrayPosition = 0
	for {
		pair, pos = ht.findPos(pos)
		if pair == nil {
			break
		}
		handler(pair.GetKey(), pair.GetVal())
	}
}
func (ht *Array) ForeachEx(handler func(key ArrayKey, value *Zval) bool) bool {
	var pair *ArrayPair
	var pos ArrayPosition = 0
	for {
		pair, pos = ht.findPos(pos)
		if pair == nil {
			break
		}
		ret := handler(pair.GetKey(), pair.GetVal())
		if !ret {
			return false
		}
	}
	return true
}
func (ht *Array) ForeachReserve(handler func(key ArrayKey, value *Zval)) {
	var pair *ArrayPair
	var pos ArrayPosition = ArrayPosition(ht.data0.Cap())
	for {
		pair, pos = ht.findPosReserve(pos)
		if pair == nil {
			break
		}
		handler(pair.GetKey(), pair.GetVal())
	}
}

func (ht *Array) ForeachIndirect(handler func(key ArrayKey, value *Zval)) {
	var pair *ArrayPair
	var pos ArrayPosition = 0
	for {
		pair, pos = ht.findPos(pos)
		if pair == nil {
			break
		}
		data := pair.GetVal().Indirect()
		if data.IsUndef() {
			continue
		}
		handler(pair.GetKey(), data)
	}
}
func (ht *Array) ForeachIndirectEx(handler func(key ArrayKey, value *Zval) bool) bool {
	var pair *ArrayPair
	var pos ArrayPosition = 0
	for {
		pair, pos = ht.findPos(pos)
		if pair == nil {
			break
		}
		data := pair.GetVal().Indirect()
		if data.IsUndef() {
			continue
		}
		ret := handler(pair.GetKey(), data)
		if !ret {
			return false
		}
	}
	return true
}
func (ht *Array) ForeachIndirectReserve(handler func(key ArrayKey, value *Zval)) {
	var pair *ArrayPair
	var pos ArrayPosition = ArrayPosition(ht.data0.Cap())
	for {
		pair, pos = ht.findPosReserve(pos)
		if pair == nil {
			break
		}
		data := pair.GetVal().Indirect()
		if data.IsUndef() {
			continue
		}
		handler(pair.GetKey(), data)
	}
}

func (ht *Array) Filter(handler func(key ArrayKey, value *Zval) bool) bool {
	var removeKeys []ArrayKey
	ht.Foreach(func(key ArrayKey, value *Zval) {
		if !handler(key, value) {
			removeKeys = append(removeKeys, key)
		}
	})

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

/**
 * Iterator & Pos
 */
func (ht *Array) currentPosVal() uint32 {
	_, realPos := ht.findPos(ht.internalPointer)
	return realPos
}

func (ht *Array) validPos(pos uint32) (uint32, bool) {
	pair, realPos := ht.findPos(pos)
	return realPos, pair != nil
}

func (ht *Array) validPosVal(pos uint32) uint32 {
	_, realPos := ht.findPos(pos)
	return realPos
}

/**
 * Internal methods
 */
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

func (ht *Array) Copy() *Array {
	return ArrayLazyDup(ht)
}

func (ht *Array) CopyEx(cap int) *Array {
	return ArrayLazyDup(ht)
}

func (ht *Array) LazyDup() *Array {
	// todo 处理懒复制逻辑
	//if ht.GetRefcount() > 1 {
	//	return ZendArrayDup(ht)
	//}
	//return ht

	return ZendArrayDup(ht)
}

func (ht *Array) ResetInternalPointer() {
	ht.internalPointer = ht.validPosVal(0)
}

// recursive
func (ht *Array) IsRecursive() bool   { return ht.protected }
func (ht *Array) ProtectRecursive()   { ht.protected = true }
func (ht *Array) UnprotectRecursive() { ht.protected = false }
