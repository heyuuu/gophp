// <<generate>>

package spl

import (
	"sik/zend"
)

/**
 * spl_SplObjectStorage
 */
type spl_SplObjectStorage struct {
	storage       zend.HashTable
	index         zend.ZendLong
	pos           zend.HashPosition
	flags         zend.ZendLong
	fptr_get_hash *zend.ZendFunction
	gcdata        *zend.Zval
	gcdata_num    int
	std           zend.ZendObject
}

func (this *spl_SplObjectStorage) GetStorage() zend.HashTable         { return this.storage }
func (this *spl_SplObjectStorage) SetStorage(value zend.HashTable)    { this.storage = value }
func (this *spl_SplObjectStorage) GetIndex() zend.ZendLong            { return this.index }
func (this *spl_SplObjectStorage) SetIndex(value zend.ZendLong)       { this.index = value }
func (this *spl_SplObjectStorage) GetPos() zend.HashPosition          { return this.pos }
func (this *spl_SplObjectStorage) SetPos(value zend.HashPosition)     { this.pos = value }
func (this *spl_SplObjectStorage) GetFlags() zend.ZendLong            { return this.flags }
func (this *spl_SplObjectStorage) SetFlags(value zend.ZendLong)       { this.flags = value }
func (this *spl_SplObjectStorage) GetFptrGetHash() *zend.ZendFunction { return this.fptr_get_hash }
func (this *spl_SplObjectStorage) SetFptrGetHash(value *zend.ZendFunction) {
	this.fptr_get_hash = value
}
func (this *spl_SplObjectStorage) GetGcdata() *zend.Zval        { return this.gcdata }
func (this *spl_SplObjectStorage) SetGcdata(value *zend.Zval)   { this.gcdata = value }
func (this *spl_SplObjectStorage) GetGcdataNum() int            { return this.gcdata_num }
func (this *spl_SplObjectStorage) SetGcdataNum(value int)       { this.gcdata_num = value }
func (this *spl_SplObjectStorage) GetStd() zend.ZendObject      { return this.std }
func (this *spl_SplObjectStorage) SetStd(value zend.ZendObject) { this.std = value }

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
	obj zend.Zval
	inf zend.Zval
}

func (this *spl_SplObjectStorageElement) GetObj() zend.Zval      { return this.obj }
func (this *spl_SplObjectStorageElement) SetObj(value zend.Zval) { this.obj = value }
func (this *spl_SplObjectStorageElement) GetInf() zend.Zval      { return this.inf }
func (this *spl_SplObjectStorageElement) SetInf(value zend.Zval) { this.inf = value }
