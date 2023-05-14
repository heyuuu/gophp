package spl

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

/**
 * spl_SplObjectStorage
 */
type SplObjectStorage struct {
	std           *types.ZendObject
	storage       *types.Array
	index         zend.ZendLong
	pos           types.ArrayPosition
	flags         zend.ZendLong
	fptr_get_hash types.IFunction
	gcdata        *types.Zval
	gcdata_num    int
}

func NewSplObjectStorage(ce *types.ClassEntry) *SplObjectStorage {
	storage := &SplObjectStorage{
		std:     types.NewObjectEx(ce, &spl_handler_SplObjectStorage),
		storage: types.NewArray(0),
		pos:     0,
	}
	return storage
}

func (this *SplObjectStorage) GetStd() *types.ZendObject { return this.std }
func (this *SplObjectStorage) GetStorage() *types.Array  { return this.storage }

func (this *SplObjectStorage) GetIndex() zend.ZendLong          { return this.index }
func (this *SplObjectStorage) SetIndex(value zend.ZendLong)     { this.index = value }
func (this *SplObjectStorage) GetPos() types.ArrayPosition      { return this.pos }
func (this *SplObjectStorage) SetPos(value types.ArrayPosition) { this.pos = value }
func (this *SplObjectStorage) GetFlags() zend.ZendLong          { return this.flags }
func (this *SplObjectStorage) SetFlags(value zend.ZendLong)     { this.flags = value }
func (this *SplObjectStorage) GetFptrGetHash() types.IFunction  { return this.fptr_get_hash }
func (this *SplObjectStorage) SetFptrGetHash(value types.IFunction) {
	this.fptr_get_hash = value
}
func (this *SplObjectStorage) GetGcdata() *types.Zval      { return this.gcdata }
func (this *SplObjectStorage) SetGcdata(value *types.Zval) { this.gcdata = value }
func (this *SplObjectStorage) GetGcdataNum() int           { return this.gcdata_num }
func (this *SplObjectStorage) SetGcdataNum(value int)      { this.gcdata_num = value }

/* spl_SplObjectStorage.flags */
func (this *SplObjectStorage) AddFlags(value zend.ZendLong)      { this.flags |= value }
func (this *SplObjectStorage) SubFlags(value zend.ZendLong)      { this.flags &^= value }
func (this *SplObjectStorage) HasFlags(value zend.ZendLong) bool { return this.flags&value != 0 }
func (this *SplObjectStorage) SwitchFlags(value zend.ZendLong, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}
func (this SplObjectStorage) IsNeedAll() bool           { return this.HasFlags(MIT_NEED_ALL) }
func (this SplObjectStorage) IsKeysAssoc() bool         { return this.HasFlags(MIT_KEYS_ASSOC) }
func (this *SplObjectStorage) SetIsNeedAll(cond bool)   { this.SwitchFlags(MIT_NEED_ALL, cond) }
func (this *SplObjectStorage) SetIsKeysAssoc(cond bool) { this.SwitchFlags(MIT_KEYS_ASSOC, cond) }

/**
 * spl_SplObjectStorageElement
 */
type SplObjectStorageElement struct {
	obj types.Zval
	inf types.Zval
}

func (this *SplObjectStorageElement) GetObj() *types.Zval { return &this.obj }
func (this *SplObjectStorageElement) GetInf() *types.Zval { return &this.inf }
