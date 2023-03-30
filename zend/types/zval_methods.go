package types

import (
	b "github.com/heyuuu/gophp/builtin"
)

/**
 * GC - Refcount
 */
func (zv *Zval) IsRefcounted() bool  { return zv.GetTypeFlags() != 0 }
func (zv *Zval) IsCollectable() bool { return b.FlagMatch(zv.GetTypeFlags(), IS_TYPE_COLLECTABLE) }

func (zv *Zval) GetRefcount() uint32 {
	b.Assert(zv.IsRefcounted())
	return zv.GetCounted().GetRefcount()
}
func (zv *Zval) SetRefcount(rc uint32) uint32 {
	b.Assert(zv.IsRefcounted())
	return zv.GetCounted().SetRefcount(rc)
}
func (zv *Zval) AddRefcount() uint32 {
	b.Assert(zv.IsRefcounted())
	return zv.GetCounted().AddRefcount()
}
func (zv *Zval) DelRefcount() uint32 {
	b.Assert(zv.IsRefcounted())
	return zv.GetCounted().DelRefcount()
}
func (zv *Zval) TryAddRefcount() {
	if zv.IsRefcounted() {
		zv.GetCounted().AddRefcount()
	}
}
func (zv *Zval) TryDelRefcount() {
	if zv.IsRefcounted() {
		zv.GetCounted().DelRefcount()
	}
}

/**
 * GC - GC_PROTECTED
 */
func (zv *Zval) IsRecursive() bool   { return zv.GetCounted().IsRecursive() }
func (zv *Zval) ProtectRecursive()   { zv.GetCounted().ProtectRecursive() }
func (zv *Zval) UnprotectRecursive() { zv.GetCounted().UnprotectRecursive() }
