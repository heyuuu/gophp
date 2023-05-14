package spl

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

/**
 * SplFixedarray
 */
type SplFixedarray struct {
	size     zend.ZendLong
	elements *types.Zval
}

func (this *SplFixedarray) GetSize() zend.ZendLong        { return this.size }
func (this *SplFixedarray) SetSize(value zend.ZendLong)   { this.size = value }
func (this *SplFixedarray) GetElements() *types.Zval      { return this.elements }
func (this *SplFixedarray) SetElements(value *types.Zval) { this.elements = value }

/**
 * SplFixedArrayObject
 */
type SplFixedArrayObject struct {
	std *types.ZendObject

	array           SplFixedarray
	fptr_offset_get types.IFunction
	fptr_offset_set types.IFunction
	fptr_offset_has types.IFunction
	fptr_offset_del types.IFunction
	fptr_count      types.IFunction
	current         int
	flags           int
	ce_get_iterator *types.ClassEntry
}

func NewSplFixedArrayObject(ce *types.ClassEntry) *SplFixedArrayObject {
	return &SplFixedArrayObject{
		std:     types.NewObjectEx(ce, &spl_handler_SplFixedArray),
		current: 0,
		flags:   0,
	}
}

func (this *SplFixedArrayObject) GetArray() SplFixedarray           { return this.array }
func (this *SplFixedArrayObject) SetArray(value SplFixedarray)      { this.array = value }
func (this *SplFixedArrayObject) GetFptrOffsetGet() types.IFunction { return this.fptr_offset_get }
func (this *SplFixedArrayObject) SetFptrOffsetGet(value types.IFunction) {
	this.fptr_offset_get = value
}
func (this *SplFixedArrayObject) GetFptrOffsetSet() types.IFunction { return this.fptr_offset_set }
func (this *SplFixedArrayObject) SetFptrOffsetSet(value types.IFunction) {
	this.fptr_offset_set = value
}
func (this *SplFixedArrayObject) GetFptrOffsetHas() types.IFunction { return this.fptr_offset_has }
func (this *SplFixedArrayObject) SetFptrOffsetHas(value types.IFunction) {
	this.fptr_offset_has = value
}
func (this *SplFixedArrayObject) GetFptrOffsetDel() types.IFunction { return this.fptr_offset_del }
func (this *SplFixedArrayObject) SetFptrOffsetDel(value types.IFunction) {
	this.fptr_offset_del = value
}
func (this *SplFixedArrayObject) GetFptrCount() types.IFunction       { return this.fptr_count }
func (this *SplFixedArrayObject) SetFptrCount(value types.IFunction)  { this.fptr_count = value }
func (this *SplFixedArrayObject) GetCurrent() int                     { return this.current }
func (this *SplFixedArrayObject) SetCurrent(value int)                { this.current = value }
func (this *SplFixedArrayObject) SetFlags(value int)                  { this.flags = value }
func (this *SplFixedArrayObject) GetCeGetIterator() *types.ClassEntry { return this.ce_get_iterator }
func (this *SplFixedArrayObject) SetCeGetIterator(value *types.ClassEntry) {
	this.ce_get_iterator = value
}
func (this *SplFixedArrayObject) GetStd() *types.ZendObject { return this.std }

/* SplFixedArrayObject.flags */
func (this *SplFixedArrayObject) AddFlags(value int)      { this.flags |= value }
func (this *SplFixedArrayObject) SubFlags(value int)      { this.flags &^= value }
func (this *SplFixedArrayObject) HasFlags(value int) bool { return this.flags&value != 0 }
func (this *SplFixedArrayObject) SwitchFlags(value int, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}
func (this SplFixedArrayObject) IsRewind() bool {
	return this.HasFlags(SPL_FIXEDARRAY_OVERLOADED_REWIND)
}
func (this SplFixedArrayObject) IsValid() bool { return this.HasFlags(SPL_FIXEDARRAY_OVERLOADED_VALID) }
func (this SplFixedArrayObject) IsCurrent() bool {
	return this.HasFlags(SPL_FIXEDARRAY_OVERLOADED_CURRENT)
}
func (this SplFixedArrayObject) IsKey() bool  { return this.HasFlags(SPL_FIXEDARRAY_OVERLOADED_KEY) }
func (this SplFixedArrayObject) IsNext() bool { return this.HasFlags(SPL_FIXEDARRAY_OVERLOADED_NEXT) }
func (this *SplFixedArrayObject) SetIsRewind(cond bool) {
	this.SwitchFlags(SPL_FIXEDARRAY_OVERLOADED_REWIND, cond)
}
func (this *SplFixedArrayObject) SetIsValid(cond bool) {
	this.SwitchFlags(SPL_FIXEDARRAY_OVERLOADED_VALID, cond)
}
func (this *SplFixedArrayObject) SetIsCurrent(cond bool) {
	this.SwitchFlags(SPL_FIXEDARRAY_OVERLOADED_CURRENT, cond)
}
func (this *SplFixedArrayObject) SetIsKey(cond bool) {
	this.SwitchFlags(SPL_FIXEDARRAY_OVERLOADED_KEY, cond)
}
func (this *SplFixedArrayObject) SetIsNext(cond bool) {
	this.SwitchFlags(SPL_FIXEDARRAY_OVERLOADED_NEXT, cond)
}

/**
 * SplFixedarrayIt
 */
type SplFixedarrayIt struct {
	intern zend.ZendUserIterator
}

// func MakeSplFixedarrayIt(intern zend.ZendUserIterator) SplFixedarrayIt {
//     return SplFixedarrayIt{
//         intern:intern,
//     }
// }
func (this *SplFixedarrayIt) GetIntern() zend.ZendUserIterator { return this.intern }

// func (this *SplFixedarrayIt) SetIntern(value zend.ZendUserIterator) { this.intern = value }
