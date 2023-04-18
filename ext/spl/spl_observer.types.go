package spl

import (
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

/**
 * spl_SplObjectStorage
 */
type spl_SplObjectStorage struct {
	storage       types2.Array
	index         zend.ZendLong
	pos           types2.ArrayPosition
	flags         zend.ZendLong
	fptr_get_hash types2.IFunction
	gcdata        *types2.Zval
	gcdata_num    int
	std           types2.ZendObject
}

//             func Makespl_SplObjectStorage(
// storage zend.HashTable,
// index zend.ZendLong,
// pos zend.ArrayPosition,
// flags zend.ZendLong,
// fptr_get_hash *zend.ZendFunction,
// gcdata *zend.Zval,
// gcdata_num int,
// std zend.ZendObject,
// ) spl_SplObjectStorage {
//                 return spl_SplObjectStorage{
//                     storage:storage,
//                     index:index,
//                     pos:pos,
//                     flags:flags,
//                     fptr_get_hash:fptr_get_hash,
//                     gcdata:gcdata,
//                     gcdata_num:gcdata_num,
//                     std:std,
//                 }
//             }
func (this *spl_SplObjectStorage) GetStorage() *types2.Array { return &this.storage }

// func (this *spl_SplObjectStorage) SetStorage(value zend.HashTable) { this.storage = value }
func (this *spl_SplObjectStorage) GetIndex() zend.ZendLong           { return this.index }
func (this *spl_SplObjectStorage) SetIndex(value zend.ZendLong)      { this.index = value }
func (this *spl_SplObjectStorage) GetPos() types2.ArrayPosition      { return this.pos }
func (this *spl_SplObjectStorage) SetPos(value types2.ArrayPosition) { this.pos = value }
func (this *spl_SplObjectStorage) GetFlags() zend.ZendLong           { return this.flags }
func (this *spl_SplObjectStorage) SetFlags(value zend.ZendLong)      { this.flags = value }
func (this *spl_SplObjectStorage) GetFptrGetHash() types2.IFunction  { return this.fptr_get_hash }
func (this *spl_SplObjectStorage) SetFptrGetHash(value types2.IFunction) {
	this.fptr_get_hash = value
}
func (this *spl_SplObjectStorage) GetGcdata() *types2.Zval      { return this.gcdata }
func (this *spl_SplObjectStorage) SetGcdata(value *types2.Zval) { this.gcdata = value }
func (this *spl_SplObjectStorage) GetGcdataNum() int            { return this.gcdata_num }
func (this *spl_SplObjectStorage) SetGcdataNum(value int)       { this.gcdata_num = value }
func (this *spl_SplObjectStorage) GetStd() *types2.ZendObject   { return &this.std }

// func (this *spl_SplObjectStorage) SetStd(value zend.ZendObject) { this.std = value }

/* spl_SplObjectStorage.flags */
func (this *spl_SplObjectStorage) AddFlags(value zend.ZendLong)      { this.flags |= value }
func (this *spl_SplObjectStorage) SubFlags(value zend.ZendLong)      { this.flags &^= value }
func (this *spl_SplObjectStorage) HasFlags(value zend.ZendLong) bool { return this.flags&value != 0 }
func (this *spl_SplObjectStorage) SwitchFlags(value zend.ZendLong, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}
func (this spl_SplObjectStorage) IsNeedAll() bool           { return this.HasFlags(MIT_NEED_ALL) }
func (this spl_SplObjectStorage) IsKeysAssoc() bool         { return this.HasFlags(MIT_KEYS_ASSOC) }
func (this *spl_SplObjectStorage) SetIsNeedAll(cond bool)   { this.SwitchFlags(MIT_NEED_ALL, cond) }
func (this *spl_SplObjectStorage) SetIsKeysAssoc(cond bool) { this.SwitchFlags(MIT_KEYS_ASSOC, cond) }

/**
 * spl_SplObjectStorageElement
 */
type spl_SplObjectStorageElement struct {
	obj types2.Zval
	inf types2.Zval
}

func (this *spl_SplObjectStorageElement) GetObj() types2.Zval { return this.obj }
func (this *spl_SplObjectStorageElement) GetInf() types2.Zval { return this.inf }
