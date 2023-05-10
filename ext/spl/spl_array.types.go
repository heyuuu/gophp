package spl

import (
	"github.com/heyuuu/gophp/php/types"
)

/**
 * SplArrayObject
 */
type SplArrayObject struct {
	array           types.Zval
	ht_iter         uint32
	ar_flags        int
	nApplyCount     uint8
	fptr_offset_get types.IFunction
	fptr_offset_set types.IFunction
	fptr_offset_has types.IFunction
	fptr_offset_del types.IFunction
	fptr_count      types.IFunction
	ce_get_iterator *types.ClassEntry
	std             types.ZendObject
}

func (this *SplArrayObject) GetArray() *types.Zval                  { return &this.array }
func (this *SplArrayObject) GetHtIter() uint32                      { return this.ht_iter }
func (this *SplArrayObject) SetHtIter(value uint32)                 { this.ht_iter = value }
func (this *SplArrayObject) GetArFlags() int                        { return this.ar_flags }
func (this *SplArrayObject) SetArFlags(value int)                   { this.ar_flags = value }
func (this *SplArrayObject) GetNApplyCount() uint8                  { return this.nApplyCount }
func (this *SplArrayObject) GetFptrOffsetGet() types.IFunction      { return this.fptr_offset_get }
func (this *SplArrayObject) SetFptrOffsetGet(value types.IFunction) { this.fptr_offset_get = value }
func (this *SplArrayObject) GetFptrOffsetSet() types.IFunction      { return this.fptr_offset_set }
func (this *SplArrayObject) SetFptrOffsetSet(value types.IFunction) { this.fptr_offset_set = value }
func (this *SplArrayObject) GetFptrOffsetHas() types.IFunction      { return this.fptr_offset_has }
func (this *SplArrayObject) SetFptrOffsetHas(value types.IFunction) { this.fptr_offset_has = value }
func (this *SplArrayObject) GetFptrOffsetDel() types.IFunction      { return this.fptr_offset_del }
func (this *SplArrayObject) SetFptrOffsetDel(value types.IFunction) { this.fptr_offset_del = value }
func (this *SplArrayObject) GetFptrCount() types.IFunction          { return this.fptr_count }
func (this *SplArrayObject) SetFptrCount(value types.IFunction)     { this.fptr_count = value }
func (this *SplArrayObject) GetCeGetIterator() *types.ClassEntry    { return this.ce_get_iterator }
func (this *SplArrayObject) SetCeGetIterator(value *types.ClassEntry) {
	this.ce_get_iterator = value
}
func (this *SplArrayObject) GetStd() types.ZendObject { return this.std }

/* SplArrayObject.ar_flags */
func (this *SplArrayObject) AddArFlags(value int)      { this.ar_flags |= value }
func (this *SplArrayObject) SubArFlags(value int)      { this.ar_flags &^= value }
func (this *SplArrayObject) HasArFlags(value int) bool { return this.ar_flags&value != 0 }
func (this *SplArrayObject) SwitchArFlags(value int, cond bool) {
	if cond {
		this.AddArFlags(value)
	} else {
		this.SubArFlags(value)
	}
}
func (this SplArrayObject) IsIsSelf() bool       { return this.HasArFlags(SPL_ARRAY_IS_SELF) }
func (this SplArrayObject) IsUseOther() bool     { return this.HasArFlags(SPL_ARRAY_USE_OTHER) }
func (this SplArrayObject) IsStdPropList() bool  { return this.HasArFlags(SPL_ARRAY_STD_PROP_LIST) }
func (this SplArrayObject) IsArrayAsProps() bool { return this.HasArFlags(SPL_ARRAY_ARRAY_AS_PROPS) }
func (this SplArrayObject) IsOverloadedValid() bool {
	return this.HasArFlags(SPL_ARRAY_OVERLOADED_VALID)
}
func (this SplArrayObject) IsOverloadedCurrent() bool {
	return this.HasArFlags(SPL_ARRAY_OVERLOADED_CURRENT)
}
func (this SplArrayObject) IsOverloadedKey() bool  { return this.HasArFlags(SPL_ARRAY_OVERLOADED_KEY) }
func (this SplArrayObject) IsOverloadedNext() bool { return this.HasArFlags(SPL_ARRAY_OVERLOADED_NEXT) }
func (this SplArrayObject) IsOverloadedRewind() bool {
	return this.HasArFlags(SPL_ARRAY_OVERLOADED_REWIND)
}
func (this SplArrayObject) IsChildArraysOnly() bool {
	return this.HasArFlags(SPL_ARRAY_CHILD_ARRAYS_ONLY)
}
func (this SplArrayObject) IsCloneMask() bool        { return this.HasArFlags(SPL_ARRAY_CLONE_MASK) }
func (this *SplArrayObject) SetIsIsSelf(cond bool)   { this.SwitchArFlags(SPL_ARRAY_IS_SELF, cond) }
func (this *SplArrayObject) SetIsUseOther(cond bool) { this.SwitchArFlags(SPL_ARRAY_USE_OTHER, cond) }
func (this *SplArrayObject) SetIsStdPropList(cond bool) {
	this.SwitchArFlags(SPL_ARRAY_STD_PROP_LIST, cond)
}
func (this *SplArrayObject) SetIsArrayAsProps(cond bool) {
	this.SwitchArFlags(SPL_ARRAY_ARRAY_AS_PROPS, cond)
}
func (this *SplArrayObject) SetIsOverloadedValid(cond bool) {
	this.SwitchArFlags(SPL_ARRAY_OVERLOADED_VALID, cond)
}
func (this *SplArrayObject) SetIsOverloadedCurrent(cond bool) {
	this.SwitchArFlags(SPL_ARRAY_OVERLOADED_CURRENT, cond)
}
func (this *SplArrayObject) SetIsOverloadedKey(cond bool) {
	this.SwitchArFlags(SPL_ARRAY_OVERLOADED_KEY, cond)
}
func (this *SplArrayObject) SetIsOverloadedNext(cond bool) {
	this.SwitchArFlags(SPL_ARRAY_OVERLOADED_NEXT, cond)
}
func (this *SplArrayObject) SetIsOverloadedRewind(cond bool) {
	this.SwitchArFlags(SPL_ARRAY_OVERLOADED_REWIND, cond)
}
func (this *SplArrayObject) SetIsChildArraysOnly(cond bool) {
	this.SwitchArFlags(SPL_ARRAY_CHILD_ARRAYS_ONLY, cond)
}
func (this *SplArrayObject) SetIsCloneMask(cond bool) { this.SwitchArFlags(SPL_ARRAY_CLONE_MASK, cond) }
