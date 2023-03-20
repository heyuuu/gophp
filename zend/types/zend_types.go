package types

import (
	b "sik/builtin"
	"sik/zend"
)

/**
 * ZendType
 */
type ZendType uintptr

func (this ZendType) IsSet() bool   { return this > 0x3 }
func (this ZendType) IsCode() bool  { return this > 0x3 && this <= 0x3ff }
func (this ZendType) IsClass() bool { return this > 0x3ff }
func (this ZendType) IsCe() bool    { return b.FlagMatch(this, 0x2) }
func (this ZendType) IsName() bool  { return this.IsClass() && !(this.IsCe()) }
func (this ZendType) Name() *String {
	var ptr = this &^ 0x3
	return b.CastPtr[String](ptr)
}
func (this ZendType) Ce() *ClassEntry {
	var ptr = this &^ 0x3
	return b.CastPtr[ClassEntry](ptr)
}
func (this ZendType) Code() int       { return this >> int64(2) }
func (this ZendType) AllowNull() bool { return b.FlagMatch(this, 0x1) }

// todo 不明确作用
func (this ZendType) TypeName() string {
	//return (*byte)(this)
	return b.CastStrAuto(b.CastPtr[byte](this))
}

/**
 * HashTableIterator
 */
type HashTableIterator struct {
	ht  *Array
	pos HashPosition
}

func (this *HashTableIterator) GetHt() *Array             { return this.ht }
func (this *HashTableIterator) SetHt(value *Array)        { this.ht = value }
func (this *HashTableIterator) GetPos() HashPosition      { return this.pos }
func (this *HashTableIterator) SetPos(value HashPosition) { this.pos = value }

/**
 * ZendObject
 */
type ZendObject struct {
	ZendRefcounted
	handle           uint32
	ce               *ClassEntry
	handlers         *zend.ZendObjectHandlers
	properties       *Array
	properties_table []Zval
}

var _ IRefcounted = &ZendObject{}

func (this *ZendObject) GetHandle() uint32                          { return this.handle }
func (this *ZendObject) SetHandle(value uint32)                     { this.handle = value }
func (this *ZendObject) GetCe() *ClassEntry                         { return this.ce }
func (this *ZendObject) SetCe(value *ClassEntry)                    { this.ce = value }
func (this *ZendObject) GetHandlers() *zend.ZendObjectHandlers      { return this.handlers }
func (this *ZendObject) SetHandlers(value *zend.ZendObjectHandlers) { this.handlers = value }
func (this *ZendObject) GetProperties() *Array                      { return this.properties }
func (this *ZendObject) SetProperties(value *Array)                 { this.properties = value }
func (this *ZendObject) GetPropertiesTable() []Zval                 { return this.properties_table }
func (this *ZendObject) SetPropertiesTable(value []Zval)            { this.properties_table = value }

/**
 * ZendResource
 */
type ZendResource struct {
	ZendRefcounted
	handle int
	type_  int
	ptr    any
}

var _ IRefcounted = &ZendResource{}

func NewZendResource(handle int, ptr any, type_ int) *ZendResource {
	return NewZendResourcePersistent(handle, ptr, type_, false)
}
func NewZendResourcePersistent(handle int, ptr any, type_ int, persistent bool) *ZendResource {
	var res = &ZendResource{
		handle: handle,
		type_:  type_,
		ptr:    ptr,
	}

	res.SetRefcount(1)
	res.SetGcTypeInfo(IS_RESOURCE)
	if persistent {
		res.SetPersistent()
	}

	return res
}

func (this *ZendResource) GetHandle() int      { return this.handle }
func (this *ZendResource) SetHandle(value int) { this.handle = value }
func (this *ZendResource) GetType() int        { return this.type_ }
func (this *ZendResource) SetType(value int)   { this.type_ = value }
func (this *ZendResource) GetPtr() any         { return this.ptr }
func (this *ZendResource) SetPtr(value any)    { this.ptr = value }

/**
 * ZendPropertyInfoList
 */
type ZendPropertyInfoList struct {
	num           int
	num_allocated int
	ptr           []*zend.ZendPropertyInfo
}

func (this *ZendPropertyInfoList) GetNum() int                           { return this.num }
func (this *ZendPropertyInfoList) SetNum(value int)                      { this.num = value }
func (this *ZendPropertyInfoList) GetNumAllocated() int                  { return this.num_allocated }
func (this *ZendPropertyInfoList) SetNumAllocated(value int)             { this.num_allocated = value }
func (this *ZendPropertyInfoList) GetPtr() []*zend.ZendPropertyInfo      { return this.ptr }
func (this *ZendPropertyInfoList) SetPtr(value []*zend.ZendPropertyInfo) { this.ptr = value }

/**
 * ZendPropertyInfoSourceList
 */
type ZendPropertyInfoSourceList struct /* union */ {
	ptr  *zend.ZendPropertyInfo
	list uintptr
}

func (this *ZendPropertyInfoSourceList) GetPtr() *zend.ZendPropertyInfo      { return this.ptr }
func (this *ZendPropertyInfoSourceList) SetPtr(value *zend.ZendPropertyInfo) { this.ptr = value }
func (this *ZendPropertyInfoSourceList) GetList() uintptr                    { return this.list }
func (this *ZendPropertyInfoSourceList) SetList(value uintptr)               { this.list = value }

/**
 * ZendReference
 */
type ZendReference struct {
	ZendRefcounted
	val     Zval
	sources ZendPropertyInfoSourceList
}

var _ IRefcounted = &ZendReference{}

func NewZendReference(val *Zval) *ZendReference {
	var ref = &ZendReference{}

	ZVAL_COPY_VALUE(ref.GetVal(), val)
	ref.sources.SetPtr(nil)

	ref.SetRefcount(1)
	ref.SetGcTypeInfo(IS_REFERENCE)

	return ref
}

func (this *ZendReference) GetVal() *Zval                               { return &this.val }
func (this *ZendReference) SetVal(value Zval)                           { this.val = value }
func (this *ZendReference) GetSources() *ZendPropertyInfoSourceList     { return &this.sources }
func (this *ZendReference) SetSources(value ZendPropertyInfoSourceList) { this.sources = value }

/**
 * ZendAstRef
 */
type ZendAstRef struct {
	ZendRefcounted
	//
	ast *zend.ZendAst
}

var _ IRefcounted = &ZendAstRef{}

func (this ZendAstRef) GcAst() *zend.ZendAst {
	//func GC_AST(p *ZendAstRef) *ZendAst {
	//	return (*ZendAst)((*byte)(p) + b.SizeOf("zend_ast_ref"))
	//}
	return this.ast
}
