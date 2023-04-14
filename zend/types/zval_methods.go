package types

import (
	b "github.com/heyuuu/gophp/builtin"
)

/**
 * GC - Refcount
 */
func (zv *Zval) IsRefcounted() bool {
	switch zv.typ {
	case IS_ARRAY, // 不包含 _IS_IMMUTABLE_ARRAY
		IS_OBJECT,
		IS_RESOURCE,
		IS_REFERENCE:
		return true
	default:
		return false
	}
}
func (zv *Zval) IsCollectable() bool {
	return zv.typ == IS_ARRAY || zv.typ == IS_OBJECT
}

func (zv *Zval) GetRefcount() uint32 {
	b.Assert(zv.IsRefcounted())
	return zv.RefCounted().GetRefcount()
}
func (zv *Zval) SetRefcount(rc uint32) uint32 {
	b.Assert(zv.IsRefcounted())
	return zv.RefCounted().SetRefcount(rc)
}
func (zv *Zval) AddRefcount() uint32 {
	b.Assert(zv.IsRefcounted())
	return zv.RefCounted().AddRefcount()
}
func (zv *Zval) DelRefcount() uint32 {
	b.Assert(zv.IsRefcounted())
	return zv.RefCounted().DelRefcount()
}
func (zv *Zval) TryAddRefcount() {
	if zv.IsRefcounted() {
		zv.RefCounted().AddRefcount()
	}
}
func (zv *Zval) TryDelRefcount() {
	if zv.IsRefcounted() {
		zv.RefCounted().DelRefcount()
	}
}

/**
 * GC - GC_PROTECTED
 */
func (zv *Zval) IsRecursive() bool   { return zv.RefCounted().IsRecursive() }
func (zv *Zval) ProtectRecursive()   { zv.RefCounted().ProtectRecursive() }
func (zv *Zval) UnprotectRecursive() { zv.RefCounted().UnprotectRecursive() }
