package types

import (
	"github.com/heyuuu/gophp/php/assert"
	"github.com/heyuuu/gophp/shim/maps"
	"github.com/heyuuu/gophp/shim/slices"
)

// symtablePair
type symtablePair struct {
	key ArrayKey
	val *Zval
}

func makeSymtablePair(key ArrayKey, val Zval) symtablePair {
	return symtablePair{key: key, val: &val}
}
func (p *symtablePair) Key() ArrayKey   { return p.key }
func (p *symtablePair) Val() Zval       { return *p.val }
func (p *symtablePair) SetVal(val Zval) { *p.val = val }
func (p *symtablePair) IsValid() bool   { return !p.val.IsUndef() }
func (p *symtablePair) MarkInvalid()    { *p.val = Undef }

var _ ArrayData = (*SymtableData)(nil)

/**
 * SymtableData 符号表类型数据
 * - 与 ArrayDataHt 功能类似，主要差异在此处所有值都是 IsIndirect 类型的 Zval
 * - 用于符号表相关操作
 */
type SymtableData struct {
	elementsCount   int
	nextFreeElement int
	data            []symtablePair   // 实际存储数据的地方
	indexes         map[ArrayKey]int // 索引到具体位置的映射
	writable        bool
}

func newSymtableData(cap int) *SymtableData {
	return &SymtableData{
		indexes:  make(map[ArrayKey]int, cap),
		data:     make([]symtablePair, 0, cap),
		writable: true,
	}
}

func (ht *SymtableData) Clone() ArrayData {
	return &SymtableData{
		elementsCount:   ht.elementsCount,
		nextFreeElement: ht.nextFreeElement,
		data:            slices.Clone(ht.data),
		indexes:         maps.Clone(ht.indexes),
		writable:        true,
	}
}

func (ht *SymtableData) Len() int  { return ht.elementsCount }
func (ht *SymtableData) Used() int { return len(ht.data) }
func (ht *SymtableData) Cap() int  { return cap(ht.data) }
func (ht *SymtableData) Count() int {
	// 因 SymbolTable 会不通过 Array 直接修改 IsIndirect 内部数据，必须重新计算

	// 计算有效元素个数，过滤 IS_INDIRECT 元素为 IS_UNDEF 的情况
	var num = 0
	for i, _ := range ht.data {
		if !ht.isValid(i) {
			continue
		}
		num++
	}
	return num
}

func (ht *SymtableData) Exists(key ArrayKey) bool {
	_, ok := ht.indexes[key]
	return ok
}
func (ht *SymtableData) Find(key ArrayKey) (Zval, ArrayPosition) {
	if pos, ok := ht.indexes[key]; ok {
		return ht.data[pos].Val(), pos
	}
	return Undef, InvalidArrayPos
}
func (ht *SymtableData) Each(handler func(key ArrayKey, value Zval) error) error {
	for i, _ := range ht.data {
		if !ht.isValid(i) {
			continue
		}

		p := ht.data[i]
		err := handler(p.Key(), p.Val())
		if err != nil {
			return err
		}
	}
	return nil
}
func (ht *SymtableData) EachReserve(handler func(key ArrayKey, value Zval) error) error {
	for i := len(ht.data) - 1; i >= 0; i-- {
		if !ht.isValid(i) {
			continue
		}

		p := ht.data[i]
		err := handler(p.Key(), p.Val())
		if err != nil {
			return err
		}
	}
	return nil
}
func (ht *SymtableData) Pos(pos ArrayPosition) (key ArrayKey, value Zval) {
	if pos < 0 || pos >= len(ht.data) || !ht.isValid(pos) {
		return
	}

	p := ht.data[pos]
	return p.Key(), p.Val()
}

func (ht *SymtableData) Add(key ArrayKey, value Zval) (bool, error) {
	ht.assertWritable()
	if _, exists := ht.indexes[key]; exists {
		return false, nil
	} else {
		ht.appendBucket(key, value)
		return true, nil
	}
}
func (ht *SymtableData) Update(key ArrayKey, value Zval) error {
	ht.assertWritable()
	if pos, exists := ht.indexes[key]; exists {
		// todo 此处需要做懒复制(主要针对数组)
		ht.data[pos].SetVal(value)
	} else {
		ht.appendBucket(key, value)
	}
	return nil
}
func (ht *SymtableData) Delete(key ArrayKey) (bool, error) {
	ht.assertWritable()
	if pos, exists := ht.indexes[key]; exists {
		ht.deleteBucket(pos)
		return true, nil
	} else {
		return false, nil
	}
}
func (ht *SymtableData) Append(value Zval) (int, error) {
	ht.assertWritable()
	idx := ht.nextFreeElement
	key := IdxKey(idx)
	assert.Assert(!ht.Exists(key))
	ht.appendBucket(key, value)
	return idx, nil
}

func (ht *SymtableData) assertWritable() { assert.Assert(ht.writable) }

func (ht *SymtableData) appendBucket(key ArrayKey, value Zval) {
	// 添加数据，更新元素计数
	ht.elementsCount++
	ht.indexes[key] = len(ht.data)
	ht.data = append(ht.data, makeSymtablePair(key, value))

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

func (ht *SymtableData) deleteBucket(pos int) {
	ht.assertWritable()
	assert.Assert(0 <= pos && pos < len(ht.data))
	assert.Assert(ht.isValid(pos))

	// 移除数据，更新元素计数
	ht.elementsCount--
	delete(ht.indexes, ht.data[pos].Key())
	ht.markInvalid(pos)

	// 若删除队尾元素，尝试清除 data 队尾无用数据
	if pos == len(ht.data)-1 {
		newDataSize := len(ht.data) - 1
		for newDataSize > 0 && !ht.isValid(newDataSize-1) {
			newDataSize--
		}
		ht.data = ht.data[:newDataSize]
	}
}

func (ht *SymtableData) isValid(pos int) bool {
	return ht.data[pos].IsValid()
}
func (ht *SymtableData) markInvalid(pos int) {
	ht.data[pos].MarkInvalid()
}
