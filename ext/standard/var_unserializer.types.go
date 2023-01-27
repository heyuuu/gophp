// <<generate>>

package standard

import (
	"sik/zend"
)

/**
 * VarEntries
 */
type VarEntries struct {
	used_slots zend.ZendLong
	next       any
	data       []*zend.Zval
}

func (this *VarEntries) GetUsedSlots() zend.ZendLong      { return this.used_slots }
func (this *VarEntries) SetUsedSlots(value zend.ZendLong) { this.used_slots = value }
func (this *VarEntries) GetNext() any                     { return this.next }
func (this *VarEntries) SetNext(value any)                { this.next = value }
func (this *VarEntries) GetData() []*zend.Zval            { return this.data }
func (this *VarEntries) SetData(value []*zend.Zval)       { this.data = value }

/**
 * VarDtorEntries
 */
type VarDtorEntries struct {
	used_slots zend.ZendLong
	next       any
	data       []zend.Zval
}

func (this *VarDtorEntries) GetUsedSlots() zend.ZendLong      { return this.used_slots }
func (this *VarDtorEntries) SetUsedSlots(value zend.ZendLong) { this.used_slots = value }
func (this *VarDtorEntries) GetNext() any                     { return this.next }
func (this *VarDtorEntries) SetNext(value any)                { this.next = value }
func (this *VarDtorEntries) GetData() []zend.Zval             { return this.data }
func (this *VarDtorEntries) SetData(value []zend.Zval)        { this.data = value }

/**
 * PhpUnserializeData
 */
type PhpUnserializeData struct {
	last            *VarEntries
	first_dtor      *VarDtorEntries
	last_dtor       *VarDtorEntries
	allowed_classes *zend.HashTable
	ref_props       *zend.HashTable
	cur_depth       zend.ZendLong
	max_depth       zend.ZendLong
	entries         VarEntries
}

func (this *PhpUnserializeData) GetLast() *VarEntries               { return this.last }
func (this *PhpUnserializeData) SetLast(value *VarEntries)          { this.last = value }
func (this *PhpUnserializeData) GetFirstDtor() *VarDtorEntries      { return this.first_dtor }
func (this *PhpUnserializeData) SetFirstDtor(value *VarDtorEntries) { this.first_dtor = value }
func (this *PhpUnserializeData) GetLastDtor() *VarDtorEntries       { return this.last_dtor }
func (this *PhpUnserializeData) SetLastDtor(value *VarDtorEntries)  { this.last_dtor = value }
func (this *PhpUnserializeData) GetAllowedClasses() *zend.HashTable { return this.allowed_classes }
func (this *PhpUnserializeData) SetAllowedClasses(value *zend.HashTable) {
	this.allowed_classes = value
}
func (this *PhpUnserializeData) GetRefProps() *zend.HashTable      { return this.ref_props }
func (this *PhpUnserializeData) SetRefProps(value *zend.HashTable) { this.ref_props = value }
func (this *PhpUnserializeData) GetCurDepth() zend.ZendLong        { return this.cur_depth }
func (this *PhpUnserializeData) SetCurDepth(value zend.ZendLong)   { this.cur_depth = value }
func (this *PhpUnserializeData) GetMaxDepth() zend.ZendLong        { return this.max_depth }
func (this *PhpUnserializeData) SetMaxDepth(value zend.ZendLong)   { this.max_depth = value }
func (this *PhpUnserializeData) GetEntries() VarEntries            { return this.entries }
func (this *PhpUnserializeData) SetEntries(value VarEntries)       { this.entries = value }
