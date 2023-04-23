package types

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
	"math"
	"runtime"
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

func NewStrKeyBucket(strKey string, zval *Zval) *Bucket {
	var key = StrKey(strKey)
	return NewBucket(key, zval)
}

func NewIndexBucket(indexKey int, zval *Zval) *Bucket {
	var key = IdxKey(indexKey)
	return NewBucket(key, zval)
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
func (this *Bucket) MarkInvalid()  { this.val.SetUndef() }

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
	flags           ZendUchar
	elementsCount   uint32
	internalPointer uint32
	nextFreeElement int
	destructor      DtorFuncT

	data     []Bucket          // 实际存储数据的地方
	indexMap map[int]uint32    // 数字索引到具体位置的映射
	keyMap   map[string]uint32 // 字符串索引到具体位置的映射

	arData *Bucket // C 源码中存储数据的地方，实际不使用

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
	ht.InitEx(size, nil)
}
func (ht *Array) InitEx(size int, pDestructor DtorFuncT) {
	var data []Bucket
	if size > 0 {
		data = make([]Bucket, 0, size)
	}

	*ht = Array{
		destructor: pDestructor,

		// 数据存储
		data:     data,
		indexMap: make(map[int]uint32),    // todo 改为 nil，延迟初始化
		keyMap:   make(map[string]uint32), // todo 改为 nil，延迟初始化
	}

	// GC 信息
	ht.SetRefcount(1)
	ht.SetGcTypeInfo(uint32(IS_ARRAY))

	// 析构函数
	if pDestructor != nil {
		runtime.SetFinalizer(ht, ht.DestroyEx)
	}
}

/* init */
func (ht *Array) SetBy(arr *Array) {
	ht.flags = arr.flags
	ht.elementsCount = arr.elementsCount
	ht.nextFreeElement = arr.nextFreeElement
	ht.destructor = arr.destructor

	ht.data = arr.data
	ht.indexMap = arr.indexMap
	ht.keyMap = arr.keyMap

	ZendHashInternalPointerReset(ht)
}

// 实际元素个数，从使用者角度的数组大小
func (ht *Array) Len() int                  { return int(ht.elementsCount) }
func (ht *Array) Cap() int                  { return cap(ht.data) }
func (ht *Array) Bucket(pos uint32) *Bucket { return &ht.data[pos] }

/* data -> Array.data */
func (ht *Array) clearData() {
	ht.assertRc1()

	ht.elementsCount = 0
	ht.internalPointer = 0
	ht.nextFreeElement = 0
	ht.data = nil
	ht.indexMap = make(map[int]uint32)
	ht.keyMap = make(map[string]uint32)
}
func (ht *Array) appendBucketStr(strKey string, zv *Zval) *Bucket {
	var bucket = NewStrKeyBucket(strKey, zv)
	return ht.appendBucket(bucket)
}

func (ht *Array) appendBucketIndex(indexKey int, zv *Zval) *Bucket {
	var bucket = NewIndexBucket(indexKey, zv)
	return ht.appendBucket(bucket)
}
func (ht *Array) appendBucket(bucket *Bucket) *Bucket {
	// 尝试 resize
	ht.resizeIfFull()

	// 添加到 data
	var idx = uint32(len(ht.data))
	ht.elementsCount++
	ht.data = append(ht.data, *bucket)

	// 更新 map
	if bucket.IsStrKey() {
		ht.keyMap[bucket.StrKey()] = idx
	} else {
		ht.indexMap[bucket.IndexKey()] = idx
	}

	if !bucket.IsStrKey() {
		var indexKey = bucket.IndexKey()
		// 更新 nextFreeElement
		if indexKey > ht.nextFreeElement {
			if indexKey < MaxLong {
				ht.nextFreeElement = indexKey + 1
			} else {
				ht.nextFreeElement = MaxLong
			}
		}
	}

	return &ht.data[idx]
}
func (ht *Array) resizeIfFull() {
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
func (ht *Array) deleteBucket(pos uint32) {
	ht.assertRc1()
	assert(int(pos) < len(ht.data))

	var p = &ht.data[pos]
	assert(p.IsValid())

	// 移除映射
	if p.key.IsStrKey() {
		delete(ht.keyMap, p.key.StrKey())
	} else {
		delete(ht.indexMap, p.key.IdxKey())
	}

	// 减少有效元素
	ht.elementsCount--

	// 更新内部指针和遍历器指针
	if ht.internalPointer == pos {
		var newIdx = ht.validPosVal(pos + 1)
		if ht.internalPointer == pos {
			ht.internalPointer = newIdx
		}
	}

	// 析构函数
	if ht.destructor != nil {
		var tmp Zval
		ZVAL_COPY_VALUE(&tmp, p.GetVal())
		ht.destructor(&tmp)
	}

	// 设置数据不可用
	p.MarkInvalid()

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
	if p.IsStrKey() {
		if pos, ok := ht.keyMap[p.StrKey()]; ok {
			return pos, true
		}
		return 0, false
	} else {
		if pos, ok := ht.indexMap[p.IndexKey()]; ok {
			return pos, true
		}
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

/* hash -> Array.indexMap & Array.keyMap */
func (ht *Array) Rehash() {
	// reset hash
	ht.assertRc1()
	ht.indexMap = make(map[int]uint32)
	ht.keyMap = make(map[string]uint32)

	// 空数组快速清空
	if ht.elementsCount == 0 {
		ht.data = nil
		return
	}

	// 移除 data 中的空位
	ht.removeHoles()

	// 重建 hash
	ht.eachBucket(func(pos uint32, p *Bucket) {
		if p.IsStrKey() {
			ht.keyMap[p.StrKey()] = pos
		} else {
			ht.indexMap[p.IndexKey()] = pos
		}
	})
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
func (ht *Array) GetPDestructor() DtorFuncT      { return ht.destructor }
func (ht *Array) SetPDestructor(value DtorFuncT) { ht.destructor = value }

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
	ht.EachValidBucketIndirect(func(pos uint32, p *Bucket, data *Zval) {
		num++
	})
	return num
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
		ht.nextFreeElement = int(ht.DataSize())
	}

	ht.Rehash()
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
	if ht.elementsCount != 0 && ht.destructor != nil {
		ht.eachValidBucket(func(pos uint32, p *Bucket) {
			ht.destructor(p.GetVal())
		})
	}
	ht.clearData()
}

func (ht *Array) Destroy() {
	if ht.elementsCount != 0 && ht.destructor != nil {
		ht.eachValidBucket(func(pos uint32, p *Bucket) {
			ht.destructor(p.GetVal())
		})
	}
}

func (ht *Array) DestroyEx() {
	/* break possible cycles */
	//GC_REMOVE_FROM_BUFFER(ht)
	//ht.SetGcTypeInfo(IS_NULL)
	ht.Destroy()
}

func (ht *Array) GracefulReverseDestroy() {
	ht.assertRc1()
	for idx := ht.DataSize(); idx > 0; idx-- {
		pos := idx - 1
		p := &ht.data[pos]
		if p.IsValid() {
			ht.deleteBucket(idx)
		}
	}
}

/**
 * Methods use index key
 */
func (ht *Array) IndexFind(index int) *Zval {
	if pos, ok := ht.indexMap[index]; ok {
		return &ht.data[pos].val
	}
	return nil
}
func (ht *Array) IndexExists(index int) bool {
	_, ok := ht.indexMap[index]
	return ok
}
func (ht *Array) IndexAdd(index int, pData *Zval) *Zval {
	ht.assertRc1()
	if ht.IndexExists(index) {
		return nil
	}
	return ht.appendBucketIndex(index, pData).GetVal()
}
func (ht *Array) IndexAddNew(index int, pData *Zval) *Zval {
	ht.assertRc1()
	return ht.appendBucketIndex(index, pData).GetVal()
}
func (ht *Array) IndexUpdate(index int, pData *Zval) *Zval {
	ht.assertRc1()

	// 若找到则更新
	if zv := ht.IndexFind(index); zv != nil {
		if ht.destructor != nil {
			ht.destructor(zv)
		}
		zv.CopyValueFrom(pData)
		return zv
	}

	// 插入后返回
	return ht.appendBucketIndex(index, pData).GetVal()
}
func (ht *Array) NextIndexInsert(pData *Zval) *Zval {
	ht.assertRc1()

	var index = ht.nextFreeElement

	if ht.IndexExists(index) {
		return nil
	}

	var p = ht.appendBucketIndex(index, pData)
	return p.GetVal()
}
func (ht *Array) NextIndexInsertNew(pData *Zval) *Zval {
	ht.assertRc1()

	var index = ht.nextFreeElement
	var p = ht.appendBucketIndex(index, pData)
	return p.GetVal()
}
func (ht *Array) IndexDelete(index int) bool {
	if idx, ok := ht.indexMap[index]; ok {
		ht.deleteBucket(idx)
		return true
	}
	return false
}

/**
 * Methods use string key
 */

func (ht *Array) KeyFind(key string) *Zval {
	if pos, ok := ht.keyMap[key]; ok {
		return &ht.data[pos].val
	}
	return nil
}
func (ht *Array) KeyFindPtr(key string) any {
	var zv = ht.KeyFind(key)
	if zv != nil {
		return zv.Ptr()
	}
	return nil
}
func (ht *Array) KeyExists(key string) bool {
	_, ok := ht.keyMap[key]
	return ok
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
func (ht *Array) KeyAdd(key string, pData *Zval) *Zval {
	ht.assertRc1()
	if ht.KeyExists(key) {
		return nil
	}

	var p = ht.appendBucketStr(key, pData)
	return p.GetVal()
}
func (ht *Array) KeyAddNew(key string, pData *Zval) *Zval {
	ht.assertRc1()

	var p = ht.appendBucketStr(key, pData)
	return p.GetVal()
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
		if ht.GetPDestructor() != nil {
			ht.GetPDestructor()(data)
		}
		data.CopyValueFrom(pData)
		return data
	}

	return ht.appendBucketStr(key, pData).GetVal()
}
func (ht *Array) KeyUpdate(key string, pData *Zval) *Zval {
	ht.assertRc1()

	if data := ht.KeyFind(key); data != nil {
		b.Assert(data != pData)
		if ht.GetPDestructor() != nil {
			ht.GetPDestructor()(data)
		}
		data.CopyValueFrom(pData)
		return data
	}

	return ht.appendBucketStr(key, pData).GetVal()
}
func (ht *Array) KeyUpdateIndirect(key string, pData *Zval) *Zval {
	ht.assertRc1()

	if data := ht.KeyFind(key); data != nil {
		b.Assert(data != pData)
		if data.IsType(IS_INDIRECT) {
			data = data.Indirect()
		}
		if ht.GetPDestructor() != nil {
			ht.GetPDestructor()(data)
		}
		data.CopyValueFrom(pData)
		return data
	}

	return ht.appendBucketStr(key, pData).GetVal()
}
func (ht *Array) KeyDelete(key string) bool {
	if pos, ok := ht.keyMap[key]; ok {
		ht.deleteBucket(pos)
		return true
	}
	return false
}

func (ht *Array) KeyDeleteIndirect(key string) bool {
	ht.assertRc1()
	if pos, ok := ht.keyMap[key]; ok {
		var p = &ht.data[pos]
		if p.GetVal().IsType(IS_INDIRECT) {
			var data *Zval = p.GetVal().Indirect()
			if data.IsType(IS_UNDEF) {
				return false
			} else {
				if ht.GetPDestructor() != nil {
					var tmp Zval
					ZVAL_COPY_VALUE(&tmp, data)
					data.SetUndef()
					ht.GetPDestructor()(&tmp)
				} else {
					data.SetUndef()
				}
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

func (ht *Array) Find(key ArrayKey) *Zval {
	if key.IsStrKey() {
		return ht.KeyFind(key.StrKey())
	} else {
		return ht.IndexFind(key.IdxKey())
	}
}

func (ht *Array) Add(key ArrayKey, pData *Zval) *Zval {
	if key.IsStrKey() {
		return ht.KeyAdd(key.StrKey(), pData)
	} else {
		return ht.IndexAdd(key.IdxKey(), pData)
	}
}

func (ht *Array) AddIndirect(key ArrayKey, pData *Zval) *Zval {
	if key.IsStrKey() {
		return ht.KeyAddIndirect(key.StrKey(), pData)
	} else {
		return ht.IndexAdd(key.IdxKey(), pData)
	}
}

func (ht *Array) Update(key ArrayKey, pData *Zval) *Zval {
	if key.IsStrKey() {
		return ht.KeyUpdate(key.StrKey(), pData)
	} else {
		return ht.IndexUpdate(key.IdxKey(), pData)
	}
}

func (ht *Array) UpdateIndirect(key ArrayKey, pData *Zval) *Zval {
	if key.IsStrKey() {
		return ht.KeyUpdateIndirect(key.StrKey(), pData)
	} else {
		return ht.IndexUpdate(key.IdxKey(), pData)
	}
}

func (ht *Array) Delete(key ArrayKey) bool {
	if key.IsStrKey() {
		return ht.KeyDelete(key.StrKey())
	} else {
		return ht.IndexDelete(key.IdxKey())
	}
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
			return
		}
		handler(p.GetArrayKey(), data)
	}
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

func (ht *Array) Iterator() *ArrayIterator {
	htReader := ArrayLazyDup(ht)
	return &ArrayIterator{ht: htReader, pos: 0}
}

func (ht *Array) IteratorEx(pos uint32) *ArrayIterator {
	htReader := ArrayLazyDup(ht)
	return &ArrayIterator{ht: htReader, pos: pos}
}

// todo 逐渐替换为 Foreach 或其他更高效代码
func (ht *Array) ForeachData() []*Bucket {
	var data = make([]*Bucket, 0)
	ht.eachValidBucket(func(_ uint32, p *Bucket) {
		data = append(data, p)
	})
	return data
}

// todo 逐渐替换为 ForeachReserve 或其他更高效代码
func (ht *Array) ForeachDataReserve() []*Bucket {
	var data = make([]*Bucket, 0)

	for i := len(ht.data) - 1; i >= 0; i-- {
		var p = &ht.data[i]
		if p.IsValid() {
			continue
		}
		data = append(data, p)
	}

	return data
}

func (ht *Array) eachBucket(handler func(pos uint32, p *Bucket)) {
	for i, _ := range ht.data {
		p := &ht.data[i]
		handler(uint32(i), p)
	}
}
func (ht *Array) eachValidBucket(handler func(pos uint32, p *Bucket)) {
	for i, _ := range ht.data {
		p := &ht.data[i]
		if p.IsValid() {
			continue
		}
		handler(uint32(i), p)
	}
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
func (ht *Array) copyDataAndHash(source *Array) {
	ht.data = make([]Bucket, len(source.data))
	copy(ht.data, source.data)

	ht.indexMap = make(map[int]uint32)
	for i, pos := range source.indexMap {
		ht.indexMap[i] = pos
	}

	ht.keyMap = make(map[string]uint32)
	for i, pos := range source.keyMap {
		ht.keyMap[i] = pos
	}
}

// 移除 this.data 数据中的 holes, 返回是否移动 bucket
func (ht *Array) removeHoles() {
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
func (ht *Array) removeHolesAndCleanInternalPointer() bool {
	ht.removeHoles()
	ht.internalPointer = 0
	return true
}
