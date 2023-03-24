package standard

import (
	"sik/zend"
	"sik/zend/types"
)

/**
 * VarEntries
 */
type VarEntries struct {
	used_slots zend.ZendLong
	next       any
	data       []*types.Zval
}

// func MakeVarEntries(used_slots zend.ZendLong, next any, data []*zend.Zval) VarEntries {
//     return VarEntries{
//         used_slots:used_slots,
//         next:next,
//         data:data,
//     }
// }
func (this *VarEntries) GetUsedSlots() zend.ZendLong      { return this.used_slots }
func (this *VarEntries) SetUsedSlots(value zend.ZendLong) { this.used_slots = value }
func (this *VarEntries) GetNext() any                     { return this.next }
func (this *VarEntries) SetNext(value any)                { this.next = value }
func (this *VarEntries) GetData() []*types.Zval           { return this.data }

// func (this *VarEntries) SetData(value []*zend.Zval) { this.data = value }

/**
 * VarDtorEntries
 */
type VarDtorEntries struct {
	used_slots zend.ZendLong
	next       any
	data       []types.Zval
}

// func MakeVarDtorEntries(used_slots zend.ZendLong, next any, data []zend.Zval) VarDtorEntries {
//     return VarDtorEntries{
//         used_slots:used_slots,
//         next:next,
//         data:data,
//     }
// }
func (this *VarDtorEntries) GetUsedSlots() zend.ZendLong      { return this.used_slots }
func (this *VarDtorEntries) SetUsedSlots(value zend.ZendLong) { this.used_slots = value }
func (this *VarDtorEntries) GetNext() any                     { return this.next }
func (this *VarDtorEntries) SetNext(value any)                { this.next = value }
func (this *VarDtorEntries) GetData() []types.Zval            { return this.data }

// func (this *VarDtorEntries) SetData(value []zend.Zval) { this.data = value }

/**
 * PhpUnserializeData
 */
type PhpUnserializeData struct {
	last            *VarEntries
	first_dtor      *VarDtorEntries
	last_dtor       *VarDtorEntries
	allowed_classes *types.Array
	ref_props       *types.Array
	cur_depth       zend.ZendLong
	max_depth       zend.ZendLong
	entries         VarEntries
}

//             func MakePhpUnserializeData(
// last *VarEntries,
// first_dtor *VarDtorEntries,
// last_dtor *VarDtorEntries,
// allowed_classes *zend.HashTable,
// ref_props *zend.HashTable,
// cur_depth zend.ZendLong,
// max_depth zend.ZendLong,
// entries VarEntries,
// ) PhpUnserializeData {
//                 return PhpUnserializeData{
//                     last:last,
//                     first_dtor:first_dtor,
//                     last_dtor:last_dtor,
//                     allowed_classes:allowed_classes,
//                     ref_props:ref_props,
//                     cur_depth:cur_depth,
//                     max_depth:max_depth,
//                     entries:entries,
//                 }
//             }
func (this *PhpUnserializeData) GetLast() *VarEntries               { return this.last }
func (this *PhpUnserializeData) SetLast(value *VarEntries)          { this.last = value }
func (this *PhpUnserializeData) GetFirstDtor() *VarDtorEntries      { return this.first_dtor }
func (this *PhpUnserializeData) SetFirstDtor(value *VarDtorEntries) { this.first_dtor = value }
func (this *PhpUnserializeData) GetLastDtor() *VarDtorEntries       { return this.last_dtor }
func (this *PhpUnserializeData) SetLastDtor(value *VarDtorEntries)  { this.last_dtor = value }
func (this *PhpUnserializeData) GetAllowedClasses() *types.Array    { return this.allowed_classes }
func (this *PhpUnserializeData) SetAllowedClasses(value *types.Array) {
	this.allowed_classes = value
}
func (this *PhpUnserializeData) GetRefProps() *types.Array       { return this.ref_props }
func (this *PhpUnserializeData) SetRefProps(value *types.Array)  { this.ref_props = value }
func (this *PhpUnserializeData) GetCurDepth() zend.ZendLong      { return this.cur_depth }
func (this *PhpUnserializeData) SetCurDepth(value zend.ZendLong) { this.cur_depth = value }
func (this *PhpUnserializeData) GetMaxDepth() zend.ZendLong      { return this.max_depth }
func (this *PhpUnserializeData) SetMaxDepth(value zend.ZendLong) { this.max_depth = value }
func (this *PhpUnserializeData) GetEntries() VarEntries          { return this.entries }

// func (this *PhpUnserializeData) SetEntries(value VarEntries) { this.entries = value }
