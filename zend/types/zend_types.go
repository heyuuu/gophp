package types

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
	"runtime"
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
func (this ZendType) Code() ZendUchar { return this >> int64(2) }
func (this ZendType) AllowNull() bool { return b.FlagMatch(this, 0x1) }

// todo 不明确作用
func (this ZendType) TypeName() string {
	//return (*byte)(this)
	return b.CastStrAuto(b.CastPtr[byte](this))
}

/**
 * ZendObject
 */
type ZendObject struct {
	ZendRefcounted
	handle          uint32
	ce              *ClassEntry
	handlers        *zend.ZendObjectHandlers
	properties      *Array
	propertiesTable []Zval
}

var _ IRefcounted = &ZendObject{}

func NewObject(ce *ClassEntry, handle uint32, handlers *zend.ZendObjectHandlers) *ZendObject {
	propertyCount := ce.GetDefaultPropertiesCount()
	if ce.IsUseGuards() {
		propertyCount++
	}

	o := &ZendObject{}
	o.handlers = handlers
	o.propertiesTable = make([]Zval, propertyCount)

	o.Init(ce, handle)
	return o
}

func (o *ZendObject) Init(ce *ClassEntry, handle uint32) {
	o.SetRefcount(1)
	o.SetGcTypeInfo(uint32(IS_OBJECT) | GC_COLLECTABLE<<GC_FLAGS_SHIFT)

	o.handle = handle
	o.ce = ce
	o.properties = nil

	if ce.IsUseGuards() {
		o.propertiesTable[ce.GetDefaultPropertiesCount()].SetUndef()
	}

	runtime.SetFinalizer(o, zend.AutoGlobalDtor)
}

func (o *ZendObject) GetHandle() uint32                          { return o.handle }
func (o *ZendObject) SetHandle(value uint32)                     { o.handle = value }
func (o *ZendObject) GetCe() *ClassEntry                         { return o.ce }
func (o *ZendObject) SetCe(value *ClassEntry)                    { o.ce = value }
func (o *ZendObject) GetHandlers() *zend.ZendObjectHandlers      { return o.handlers }
func (o *ZendObject) SetHandlers(value *zend.ZendObjectHandlers) { o.handlers = value }
func (o *ZendObject) GetProperties() *Array                      { return o.properties }
func (o *ZendObject) SetProperties(value *Array)                 { o.properties = value }
func (o *ZendObject) GetPropertiesTable() []Zval                 { return o.propertiesTable }
func (o *ZendObject) SetPropertiesTable(value []Zval)            { o.propertiesTable = value }

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

	runtime.SetFinalizer(res, zend.ZendListFree)

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
	ref.SetGcTypeInfo(uint32(IS_REFERENCE))

	runtime.SetFinalizer(ref, zend.ZendReferenceDestroy)

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

func NewAstRef(ast *zend.ZendAst) *ZendAstRef {
	b.Assert(ast != nil)

	// init
	var ref *ZendAstRef = &ZendAstRef{}
	//tree_size = zend.ZendAstTreeSize(ast) + b.SizeOf("zend_ast_ref")
	//ref = zend.Emalloc(tree_size)
	zend.ZendAstTreeCopy(ast, ref.ast)
	ref.SetRefcount(1)
	ref.SetGcTypeInfo(uint32(IS_CONSTANT_AST))

	// dtor
	runtime.SetFinalizer(ref, zend.ZendAstRefDestroy)

	return ref
}

func (this ZendAstRef) GcAst() *zend.ZendAst {
	//func GC_AST(p *ZendAstRef) *ZendAst {
	//	return (*ZendAst)((*byte)(p) + b.SizeOf("zend_ast_ref"))
	//}
	return this.ast
}
