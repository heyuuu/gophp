// <<generate>>

package spl

import (
	"sik/zend"
	"sik/zend/types"
)

/**
 * SplFixedarray
 */
type SplFixedarray struct {
	size     zend.ZendLong
	elements *types.Zval
}

// func MakeSplFixedarray(size zend.ZendLong, elements *zend.Zval) SplFixedarray {
//     return SplFixedarray{
//         size:size,
//         elements:elements,
//     }
// }
func (this *SplFixedarray) GetSize() zend.ZendLong        { return this.size }
func (this *SplFixedarray) SetSize(value zend.ZendLong)   { this.size = value }
func (this *SplFixedarray) GetElements() *types.Zval      { return this.elements }
func (this *SplFixedarray) SetElements(value *types.Zval) { this.elements = value }

/**
 * SplFixedarrayObject
 */
type SplFixedarrayObject struct {
	array           SplFixedarray
	fptr_offset_get *zend.ZendFunction
	fptr_offset_set *zend.ZendFunction
	fptr_offset_has *zend.ZendFunction
	fptr_offset_del *zend.ZendFunction
	fptr_count      *zend.ZendFunction
	current         int
	flags           int
	ce_get_iterator *zend.ZendClassEntry
	std             types.ZendObject
}

//             func MakeSplFixedarrayObject(
// array SplFixedarray,
// fptr_offset_get *zend.ZendFunction,
// fptr_offset_set *zend.ZendFunction,
// fptr_offset_has *zend.ZendFunction,
// fptr_offset_del *zend.ZendFunction,
// fptr_count *zend.ZendFunction,
// current int,
// flags int,
// ce_get_iterator *zend.ZendClassEntry,
// std zend.ZendObject,
// ) SplFixedarrayObject {
//                 return SplFixedarrayObject{
//                     array:array,
//                     fptr_offset_get:fptr_offset_get,
//                     fptr_offset_set:fptr_offset_set,
//                     fptr_offset_has:fptr_offset_has,
//                     fptr_offset_del:fptr_offset_del,
//                     fptr_count:fptr_count,
//                     current:current,
//                     flags:flags,
//                     ce_get_iterator:ce_get_iterator,
//                     std:std,
//                 }
//             }
func (this *SplFixedarrayObject) GetArray() SplFixedarray              { return this.array }
func (this *SplFixedarrayObject) SetArray(value SplFixedarray)         { this.array = value }
func (this *SplFixedarrayObject) GetFptrOffsetGet() *zend.ZendFunction { return this.fptr_offset_get }
func (this *SplFixedarrayObject) SetFptrOffsetGet(value *zend.ZendFunction) {
	this.fptr_offset_get = value
}
func (this *SplFixedarrayObject) GetFptrOffsetSet() *zend.ZendFunction { return this.fptr_offset_set }
func (this *SplFixedarrayObject) SetFptrOffsetSet(value *zend.ZendFunction) {
	this.fptr_offset_set = value
}
func (this *SplFixedarrayObject) GetFptrOffsetHas() *zend.ZendFunction { return this.fptr_offset_has }
func (this *SplFixedarrayObject) SetFptrOffsetHas(value *zend.ZendFunction) {
	this.fptr_offset_has = value
}
func (this *SplFixedarrayObject) GetFptrOffsetDel() *zend.ZendFunction { return this.fptr_offset_del }
func (this *SplFixedarrayObject) SetFptrOffsetDel(value *zend.ZendFunction) {
	this.fptr_offset_del = value
}
func (this *SplFixedarrayObject) GetFptrCount() *zend.ZendFunction      { return this.fptr_count }
func (this *SplFixedarrayObject) SetFptrCount(value *zend.ZendFunction) { this.fptr_count = value }
func (this *SplFixedarrayObject) GetCurrent() int                       { return this.current }
func (this *SplFixedarrayObject) SetCurrent(value int)                  { this.current = value }

// func (this *SplFixedarrayObject)  GetFlags() int      { return this.flags }
func (this *SplFixedarrayObject) SetFlags(value int)                     { this.flags = value }
func (this *SplFixedarrayObject) GetCeGetIterator() *zend.ZendClassEntry { return this.ce_get_iterator }
func (this *SplFixedarrayObject) SetCeGetIterator(value *zend.ZendClassEntry) {
	this.ce_get_iterator = value
}
func (this *SplFixedarrayObject) GetStd() types.ZendObject { return this.std }

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
