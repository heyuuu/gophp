// <<generate>>

package spl

/**
 * SplArrayObject
 */
type SplArrayObject struct {
	array           zend.Zval
	ht_iter         uint32
	ar_flags        int
	nApplyCount     uint8
	fptr_offset_get *zend.ZendFunction
	fptr_offset_set *zend.ZendFunction
	fptr_offset_has *zend.ZendFunction
	fptr_offset_del *zend.ZendFunction
	fptr_count      *zend.ZendFunction
	ce_get_iterator *zend.ZendClassEntry
	std             zend.ZendObject
}

func (this SplArrayObject) GetArray() zend.Zval                        { return this.array }
func (this *SplArrayObject) SetArray(value zend.Zval)                  { this.array = value }
func (this SplArrayObject) GetHtIter() uint32                          { return this.ht_iter }
func (this *SplArrayObject) SetHtIter(value uint32)                    { this.ht_iter = value }
func (this SplArrayObject) GetArFlags() int                            { return this.ar_flags }
func (this *SplArrayObject) SetArFlags(value int)                      { this.ar_flags = value }
func (this SplArrayObject) GetNApplyCount() uint8                      { return this.nApplyCount }
func (this *SplArrayObject) SetNApplyCount(value uint8)                { this.nApplyCount = value }
func (this SplArrayObject) GetFptrOffsetGet() *zend.ZendFunction       { return this.fptr_offset_get }
func (this *SplArrayObject) SetFptrOffsetGet(value *zend.ZendFunction) { this.fptr_offset_get = value }
func (this SplArrayObject) GetFptrOffsetSet() *zend.ZendFunction       { return this.fptr_offset_set }
func (this *SplArrayObject) SetFptrOffsetSet(value *zend.ZendFunction) { this.fptr_offset_set = value }
func (this SplArrayObject) GetFptrOffsetHas() *zend.ZendFunction       { return this.fptr_offset_has }
func (this *SplArrayObject) SetFptrOffsetHas(value *zend.ZendFunction) { this.fptr_offset_has = value }
func (this SplArrayObject) GetFptrOffsetDel() *zend.ZendFunction       { return this.fptr_offset_del }
func (this *SplArrayObject) SetFptrOffsetDel(value *zend.ZendFunction) { this.fptr_offset_del = value }
func (this SplArrayObject) GetFptrCount() *zend.ZendFunction           { return this.fptr_count }
func (this *SplArrayObject) SetFptrCount(value *zend.ZendFunction)     { this.fptr_count = value }
func (this SplArrayObject) GetCeGetIterator() *zend.ZendClassEntry     { return this.ce_get_iterator }
func (this *SplArrayObject) SetCeGetIterator(value *zend.ZendClassEntry) {
	this.ce_get_iterator = value
}
func (this SplArrayObject) GetStd() zend.ZendObject       { return this.std }
func (this *SplArrayObject) SetStd(value zend.ZendObject) { this.std = value }
