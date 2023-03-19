// <<generate>>

package spl

import (
	"sik/zend"
	"sik/zend/types"
)

/**
 * SplPtrHeap
 */
type SplPtrHeap struct {
	elements  any
	ctor      SplPtrHeapCtorFunc
	dtor      SplPtrHeapDtorFunc
	cmp       SplPtrHeapCmpFunc
	count     int
	flags     int
	max_size  int
	elem_size int
}

//             func MakeSplPtrHeap(
// elements any,
// ctor SplPtrHeapCtorFunc,
// dtor SplPtrHeapDtorFunc,
// cmp SplPtrHeapCmpFunc,
// count int,
// flags int,
// max_size int,
// elem_size int,
// ) SplPtrHeap {
//                 return SplPtrHeap{
//                     elements:elements,
//                     ctor:ctor,
//                     dtor:dtor,
//                     cmp:cmp,
//                     count:count,
//                     flags:flags,
//                     max_size:max_size,
//                     elem_size:elem_size,
//                 }
//             }
func (this *SplPtrHeap) GetElements() any                 { return this.elements }
func (this *SplPtrHeap) SetElements(value any)            { this.elements = value }
func (this *SplPtrHeap) GetCtor() SplPtrHeapCtorFunc      { return this.ctor }
func (this *SplPtrHeap) SetCtor(value SplPtrHeapCtorFunc) { this.ctor = value }
func (this *SplPtrHeap) GetDtor() SplPtrHeapDtorFunc      { return this.dtor }
func (this *SplPtrHeap) SetDtor(value SplPtrHeapDtorFunc) { this.dtor = value }
func (this *SplPtrHeap) GetCmp() SplPtrHeapCmpFunc        { return this.cmp }
func (this *SplPtrHeap) SetCmp(value SplPtrHeapCmpFunc)   { this.cmp = value }
func (this *SplPtrHeap) GetCount() int                    { return this.count }
func (this *SplPtrHeap) SetCount(value int)               { this.count = value }
func (this *SplPtrHeap) GetFlags() int                    { return this.flags }
func (this *SplPtrHeap) SetFlags(value int)               { this.flags = value }
func (this *SplPtrHeap) GetMaxSize() int                  { return this.max_size }
func (this *SplPtrHeap) SetMaxSize(value int)             { this.max_size = value }
func (this *SplPtrHeap) GetElemSize() int                 { return this.elem_size }
func (this *SplPtrHeap) SetElemSize(value int)            { this.elem_size = value }

/* SplPtrHeap.flags */
func (this *SplPtrHeap) AddFlags(value int)      { this.flags |= value }
func (this *SplPtrHeap) SubFlags(value int)      { this.flags &^= value }
func (this *SplPtrHeap) HasFlags(value int) bool { return this.flags&value != 0 }
func (this *SplPtrHeap) SwitchFlags(value int, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}
func (this SplPtrHeap) IsHeapCorrupted() bool         { return this.HasFlags(SPL_HEAP_CORRUPTED) }
func (this *SplPtrHeap) SetIsHeapCorrupted(cond bool) { this.SwitchFlags(SPL_HEAP_CORRUPTED, cond) }

/**
 * SplHeapObject
 */
type SplHeapObject struct {
	heap            *SplPtrHeap
	flags           int
	ce_get_iterator *types.ClassEntry
	fptr_cmp        *zend.ZendFunction
	fptr_count      *zend.ZendFunction
	std             types.ZendObject
}

//             func MakeSplHeapObject(
// heap *SplPtrHeap,
// flags int,
// ce_get_iterator *zend.ClassEntry,
// fptr_cmp *zend.ZendFunction,
// fptr_count *zend.ZendFunction,
// std zend.ZendObject,
// ) SplHeapObject {
//                 return SplHeapObject{
//                     heap:heap,
//                     flags:flags,
//                     ce_get_iterator:ce_get_iterator,
//                     fptr_cmp:fptr_cmp,
//                     fptr_count:fptr_count,
//                     std:std,
//                 }
//             }
func (this *SplHeapObject) GetHeap() *SplPtrHeap                     { return this.heap }
func (this *SplHeapObject) SetHeap(value *SplPtrHeap)                { this.heap = value }
func (this *SplHeapObject) GetFlags() int                            { return this.flags }
func (this *SplHeapObject) SetFlags(value int)                       { this.flags = value }
func (this *SplHeapObject) GetCeGetIterator() *types.ClassEntry      { return this.ce_get_iterator }
func (this *SplHeapObject) SetCeGetIterator(value *types.ClassEntry) { this.ce_get_iterator = value }
func (this *SplHeapObject) GetFptrCmp() *zend.ZendFunction           { return this.fptr_cmp }
func (this *SplHeapObject) SetFptrCmp(value *zend.ZendFunction)      { this.fptr_cmp = value }
func (this *SplHeapObject) GetFptrCount() *zend.ZendFunction         { return this.fptr_count }
func (this *SplHeapObject) SetFptrCount(value *zend.ZendFunction)    { this.fptr_count = value }
func (this *SplHeapObject) GetStd() types.ZendObject                 { return this.std }

// func (this *SplHeapObject) SetStd(value zend.ZendObject) { this.std = value }

/* SplHeapObject.flags */
func (this *SplHeapObject) AddFlags(value int)      { this.flags |= value }
func (this *SplHeapObject) SubFlags(value int)      { this.flags &^= value }
func (this *SplHeapObject) HasFlags(value int) bool { return this.flags&value != 0 }
func (this *SplHeapObject) SwitchFlags(value int, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}

/**
 * SplHeapIt
 */
type SplHeapIt struct {
	intern zend.ZendUserIterator
	flags  int
}

// func MakeSplHeapIt(intern zend.ZendUserIterator, flags int) SplHeapIt {
//     return SplHeapIt{
//         intern:intern,
//         flags:flags,
//     }
// }
func (this *SplHeapIt) GetIntern() zend.ZendUserIterator { return this.intern }

// func (this *SplHeapIt) SetIntern(value zend.ZendUserIterator) { this.intern = value }
// func (this *SplHeapIt)  GetFlags() int      { return this.flags }
func (this *SplHeapIt) SetFlags(value int) { this.flags = value }

/* SplHeapIt.flags */
func (this *SplHeapIt) AddFlags(value int)      { this.flags |= value }
func (this *SplHeapIt) SubFlags(value int)      { this.flags &^= value }
func (this *SplHeapIt) HasFlags(value int) bool { return this.flags&value != 0 }
func (this *SplHeapIt) SwitchFlags(value int, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}

/**
 * SplPqueueElem
 */
type SplPqueueElem struct {
	data     types.Zval
	priority types.Zval
}

// func MakeSplPqueueElem(data zend.Zval, priority zend.Zval) SplPqueueElem {
//     return SplPqueueElem{
//         data:data,
//         priority:priority,
//     }
// }
func (this *SplPqueueElem) GetData() types.Zval { return this.data }

// func (this *SplPqueueElem) SetData(value zend.Zval) { this.data = value }
func (this *SplPqueueElem) GetPriority() types.Zval { return this.priority }

// func (this *SplPqueueElem) SetPriority(value zend.Zval) { this.priority = value }
