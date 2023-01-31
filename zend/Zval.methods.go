package zend

import b "sik/builtin"

/**
 * Refcount
 */
func (this *Zval) IsRefcounted() bool  { return this.GetTypeFlags() != 0 }
func (this *Zval) IsCollectable() bool { return b.FlagMatch(this.GetTypeFlags(), IS_TYPE_COLLECTABLE) }

func (this *Zval) GetRefcount() uint32 {
	ZEND_ASSERT(this.IsRefcounted())
	return this.GetCounted().GetRefcount()
}
func (this *Zval) SetRefcount(rc uint32) uint32 {
	ZEND_ASSERT(this.IsRefcounted())
	return this.GetCounted().SetRefcount(rc)
}
func (this *Zval) AddRefcount() uint32 {
	ZEND_ASSERT(this.IsRefcounted())
	return this.GetCounted().AddRefcount()
}
func (this *Zval) DelRefcount() uint32 {
	ZEND_ASSERT(this.IsRefcounted())
	return this.GetCounted().DelRefcount()
}
