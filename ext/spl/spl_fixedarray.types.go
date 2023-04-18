package spl

import (
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

/**
 * SplFixedarray
 */
type SplFixedarray struct {
	size     zend.ZendLong
	elements *types2.Zval
}

// func MakeSplFixedarray(size zend.ZendLong, elements *zend.Zval) SplFixedarray {
//     return SplFixedarray{
//         size:size,
//         elements:elements,
//     }
// }
func (this *SplFixedarray) GetSize() zend.ZendLong         { return this.size }
func (this *SplFixedarray) SetSize(value zend.ZendLong)    { this.size = value }
func (this *SplFixedarray) GetElements() *types2.Zval      { return this.elements }
func (this *SplFixedarray) SetElements(value *types2.Zval) { this.elements = value }

/**
 * SplFixedarrayObject
 */
type SplFixedarrayObject struct {
	array           SplFixedarray
	fptr_offset_get types2.IFunction
	fptr_offset_set types2.IFunction
	fptr_offset_has types2.IFunction
	fptr_offset_del types2.IFunction
	fptr_count      types2.IFunction
	current         int
	flags           int
	ce_get_iterator *types2.ClassEntry
	std             types2.ZendObject
}

func (this *SplFixedarrayObject) GetArray() SplFixedarray            { return this.array }
func (this *SplFixedarrayObject) SetArray(value SplFixedarray)       { this.array = value }
func (this *SplFixedarrayObject) GetFptrOffsetGet() types2.IFunction { return this.fptr_offset_get }
func (this *SplFixedarrayObject) SetFptrOffsetGet(value types2.IFunction) {
	this.fptr_offset_get = value
}
func (this *SplFixedarrayObject) GetFptrOffsetSet() types2.IFunction { return this.fptr_offset_set }
func (this *SplFixedarrayObject) SetFptrOffsetSet(value types2.IFunction) {
	this.fptr_offset_set = value
}
func (this *SplFixedarrayObject) GetFptrOffsetHas() types2.IFunction { return this.fptr_offset_has }
func (this *SplFixedarrayObject) SetFptrOffsetHas(value types2.IFunction) {
	this.fptr_offset_has = value
}
func (this *SplFixedarrayObject) GetFptrOffsetDel() types2.IFunction { return this.fptr_offset_del }
func (this *SplFixedarrayObject) SetFptrOffsetDel(value types2.IFunction) {
	this.fptr_offset_del = value
}
func (this *SplFixedarrayObject) GetFptrCount() types2.IFunction      { return this.fptr_count }
func (this *SplFixedarrayObject) SetFptrCount(value types2.IFunction) { this.fptr_count = value }
func (this *SplFixedarrayObject) GetCurrent() int                     { return this.current }
func (this *SplFixedarrayObject) SetCurrent(value int)                { this.current = value }

// func (this *SplFixedarrayObject)  GetFlags() int      { return this.flags }
func (this *SplFixedarrayObject) SetFlags(value int)                   { this.flags = value }
func (this *SplFixedarrayObject) GetCeGetIterator() *types2.ClassEntry { return this.ce_get_iterator }
func (this *SplFixedarrayObject) SetCeGetIterator(value *types2.ClassEntry) {
	this.ce_get_iterator = value
}
func (this *SplFixedarrayObject) GetStd() types2.ZendObject { return this.std }

// func (this *SplFixedarrayObject) SetStd(value zend.ZendObject) { this.std = value }

/* SplFixedarrayObject.flags */
func (this *SplFixedarrayObject) AddFlags(value int)      { this.flags |= value }
func (this *SplFixedarrayObject) SubFlags(value int)      { this.flags &^= value }
func (this *SplFixedarrayObject) HasFlags(value int) bool { return this.flags&value != 0 }
func (this *SplFixedarrayObject) SwitchFlags(value int, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}
func (this SplFixedarrayObject) IsRewind() bool {
	return this.HasFlags(SPL_FIXEDARRAY_OVERLOADED_REWIND)
}
func (this SplFixedarrayObject) IsValid() bool { return this.HasFlags(SPL_FIXEDARRAY_OVERLOADED_VALID) }
func (this SplFixedarrayObject) IsCurrent() bool {
	return this.HasFlags(SPL_FIXEDARRAY_OVERLOADED_CURRENT)
}
func (this SplFixedarrayObject) IsKey() bool  { return this.HasFlags(SPL_FIXEDARRAY_OVERLOADED_KEY) }
func (this SplFixedarrayObject) IsNext() bool { return this.HasFlags(SPL_FIXEDARRAY_OVERLOADED_NEXT) }
func (this *SplFixedarrayObject) SetIsRewind(cond bool) {
	this.SwitchFlags(SPL_FIXEDARRAY_OVERLOADED_REWIND, cond)
}
func (this *SplFixedarrayObject) SetIsValid(cond bool) {
	this.SwitchFlags(SPL_FIXEDARRAY_OVERLOADED_VALID, cond)
}
func (this *SplFixedarrayObject) SetIsCurrent(cond bool) {
	this.SwitchFlags(SPL_FIXEDARRAY_OVERLOADED_CURRENT, cond)
}
func (this *SplFixedarrayObject) SetIsKey(cond bool) {
	this.SwitchFlags(SPL_FIXEDARRAY_OVERLOADED_KEY, cond)
}
func (this *SplFixedarrayObject) SetIsNext(cond bool) {
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
