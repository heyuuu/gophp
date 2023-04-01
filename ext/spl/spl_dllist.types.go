package spl

import (
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/types"
)

/**
 * SplPtrLlistElement
 */
type SplPtrLlistElement struct {
	prev *SplPtrLlistElement
	next *SplPtrLlistElement
	rc   int
	data types.Zval
}

// func MakeSplPtrLlistElement(prev *SplPtrLlistElement, next *SplPtrLlistElement, rc int, data zend.Zval) SplPtrLlistElement {
//     return SplPtrLlistElement{
//         prev:prev,
//         next:next,
//         rc:rc,
//         data:data,
//     }
// }
func (this *SplPtrLlistElement) GetPrev() *SplPtrLlistElement      { return this.prev }
func (this *SplPtrLlistElement) SetPrev(value *SplPtrLlistElement) { this.prev = value }
func (this *SplPtrLlistElement) GetNext() *SplPtrLlistElement      { return this.next }
func (this *SplPtrLlistElement) SetNext(value *SplPtrLlistElement) { this.next = value }
func (this *SplPtrLlistElement) GetRc() int                        { return this.rc }
func (this *SplPtrLlistElement) SetRc(value int)                   { this.rc = value }
func (this *SplPtrLlistElement) GetData() types.Zval               { return this.data }

// func (this *SplPtrLlistElement) SetData(value zend.Zval) { this.data = value }

/**
 * SplPtrLlist
 */
type SplPtrLlist struct {
	head  *SplPtrLlistElement
	tail  *SplPtrLlistElement
	dtor  SplPtrLlistDtorFunc
	ctor  SplPtrLlistCtorFunc
	count int
}

// func MakeSplPtrLlist(head *SplPtrLlistElement, tail *SplPtrLlistElement, dtor SplPtrLlistDtorFunc, ctor SplPtrLlistCtorFunc, count int) SplPtrLlist {
//     return SplPtrLlist{
//         head:head,
//         tail:tail,
//         dtor:dtor,
//         ctor:ctor,
//         count:count,
//     }
// }
func (this *SplPtrLlist) GetHead() *SplPtrLlistElement      { return this.head }
func (this *SplPtrLlist) SetHead(value *SplPtrLlistElement) { this.head = value }
func (this *SplPtrLlist) GetTail() *SplPtrLlistElement      { return this.tail }
func (this *SplPtrLlist) SetTail(value *SplPtrLlistElement) { this.tail = value }
func (this *SplPtrLlist) GetDtor() SplPtrLlistDtorFunc      { return this.dtor }
func (this *SplPtrLlist) SetDtor(value SplPtrLlistDtorFunc) { this.dtor = value }
func (this *SplPtrLlist) GetCtor() SplPtrLlistCtorFunc      { return this.ctor }
func (this *SplPtrLlist) SetCtor(value SplPtrLlistCtorFunc) { this.ctor = value }
func (this *SplPtrLlist) GetCount() int                     { return this.count }
func (this *SplPtrLlist) SetCount(value int)                { this.count = value }

/**
 * SplDllistObject
 */
type SplDllistObject struct {
	llist             *SplPtrLlist
	traverse_position int
	traverse_pointer  *SplPtrLlistElement
	flags             int
	fptr_offset_get   types.IFunction
	fptr_offset_set   types.IFunction
	fptr_offset_has   types.IFunction
	fptr_offset_del   types.IFunction
	fptr_count        types.IFunction
	ce_get_iterator   *types.ClassEntry
	gc_data           *types.Zval
	gc_data_count     int
	std               types.ZendObject
}

//             func MakeSplDllistObject(
// llist *SplPtrLlist,
// traverse_position int,
// traverse_pointer *SplPtrLlistElement,
// flags int,
// fptr_offset_get *zend.ZendFunction,
// fptr_offset_set *zend.ZendFunction,
// fptr_offset_has *zend.ZendFunction,
// fptr_offset_del *zend.ZendFunction,
// fptr_count *zend.ZendFunction,
// ce_get_iterator *zend.ClassEntry,
// gc_data *zend.Zval,
// gc_data_count int,
// std zend.ZendObject,
// ) SplDllistObject {
//                 return SplDllistObject{
//                     llist:llist,
//                     traverse_position:traverse_position,
//                     traverse_pointer:traverse_pointer,
//                     flags:flags,
//                     fptr_offset_get:fptr_offset_get,
//                     fptr_offset_set:fptr_offset_set,
//                     fptr_offset_has:fptr_offset_has,
//                     fptr_offset_del:fptr_offset_del,
//                     fptr_count:fptr_count,
//                     ce_get_iterator:ce_get_iterator,
//                     gc_data:gc_data,
//                     gc_data_count:gc_data_count,
//                     std:std,
//                 }
//             }
func (this *SplDllistObject) GetLlist() *SplPtrLlist                  { return this.llist }
func (this *SplDllistObject) SetLlist(value *SplPtrLlist)             { this.llist = value }
func (this *SplDllistObject) GetTraversePosition() int                { return this.traverse_position }
func (this *SplDllistObject) SetTraversePosition(value int)           { this.traverse_position = value }
func (this *SplDllistObject) GetTraversePointer() *SplPtrLlistElement { return this.traverse_pointer }
func (this *SplDllistObject) SetTraversePointer(value *SplPtrLlistElement) {
	this.traverse_pointer = value
}
func (this *SplDllistObject) GetFlags() int                         { return this.flags }
func (this *SplDllistObject) SetFlags(value int)                    { this.flags = value }
func (this *SplDllistObject) GetFptrOffsetGet() *types.ZendFunction { return this.fptr_offset_get }
func (this *SplDllistObject) SetFptrOffsetGet(value *types.ZendFunction) {
	this.fptr_offset_get = value
}
func (this *SplDllistObject) GetFptrOffsetSet() *types.ZendFunction { return this.fptr_offset_set }
func (this *SplDllistObject) SetFptrOffsetSet(value *types.ZendFunction) {
	this.fptr_offset_set = value
}
func (this *SplDllistObject) GetFptrOffsetHas() *types.ZendFunction { return this.fptr_offset_has }
func (this *SplDllistObject) SetFptrOffsetHas(value *types.ZendFunction) {
	this.fptr_offset_has = value
}
func (this *SplDllistObject) GetFptrOffsetDel() *types.ZendFunction { return this.fptr_offset_del }
func (this *SplDllistObject) SetFptrOffsetDel(value *types.ZendFunction) {
	this.fptr_offset_del = value
}
func (this *SplDllistObject) GetFptrCount() *types.ZendFunction      { return this.fptr_count }
func (this *SplDllistObject) SetFptrCount(value *types.ZendFunction) { this.fptr_count = value }
func (this *SplDllistObject) GetCeGetIterator() *types.ClassEntry    { return this.ce_get_iterator }
func (this *SplDllistObject) SetCeGetIterator(value *types.ClassEntry) {
	this.ce_get_iterator = value
}
func (this *SplDllistObject) GetGcData() *types.Zval      { return this.gc_data }
func (this *SplDllistObject) SetGcData(value *types.Zval) { this.gc_data = value }
func (this *SplDllistObject) GetGcDataCount() int         { return this.gc_data_count }
func (this *SplDllistObject) SetGcDataCount(value int)    { this.gc_data_count = value }
func (this *SplDllistObject) GetStd() types.ZendObject    { return this.std }

// func (this *SplDllistObject) SetStd(value zend.ZendObject) { this.std = value }

/* SplDllistObject.flags */
func (this *SplDllistObject) AddFlags(value int)      { this.flags |= value }
func (this *SplDllistObject) SubFlags(value int)      { this.flags &^= value }
func (this *SplDllistObject) HasFlags(value int) bool { return this.flags&value != 0 }
func (this *SplDllistObject) SwitchFlags(value int, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}
func (this SplDllistObject) IsItFix() bool         { return this.HasFlags(SPL_DLLIST_IT_FIX) }
func (this *SplDllistObject) SetIsItFix(cond bool) { this.SwitchFlags(SPL_DLLIST_IT_FIX, cond) }

/**
 * SplDllistIt
 */
type SplDllistIt struct {
	intern            zend.ZendUserIterator
	traverse_pointer  *SplPtrLlistElement
	traverse_position int
	flags             int
}

// func MakeSplDllistIt(intern zend.ZendUserIterator, traverse_pointer *SplPtrLlistElement, traverse_position int, flags int) SplDllistIt {
//     return SplDllistIt{
//         intern:intern,
//         traverse_pointer:traverse_pointer,
//         traverse_position:traverse_position,
//         flags:flags,
//     }
// }
func (this *SplDllistIt) GetIntern() zend.ZendUserIterator { return this.intern }

// func (this *SplDllistIt) SetIntern(value zend.ZendUserIterator) { this.intern = value }
func (this *SplDllistIt) GetTraversePointer() *SplPtrLlistElement      { return this.traverse_pointer }
func (this *SplDllistIt) SetTraversePointer(value *SplPtrLlistElement) { this.traverse_pointer = value }
func (this *SplDllistIt) GetTraversePosition() int                     { return this.traverse_position }
func (this *SplDllistIt) SetTraversePosition(value int)                { this.traverse_position = value }

// func (this *SplDllistIt)  GetFlags() int      { return this.flags }
func (this *SplDllistIt) SetFlags(value int) { this.flags = value }

/* SplDllistIt.flags */
func (this *SplDllistIt) AddFlags(value int)      { this.flags |= value }
func (this *SplDllistIt) SubFlags(value int)      { this.flags &^= value }
func (this *SplDllistIt) HasFlags(value int) bool { return this.flags&value != 0 }
func (this *SplDllistIt) SwitchFlags(value int, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}
