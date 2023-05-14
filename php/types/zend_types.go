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
func (this ZendType) Code() uint8     { return this >> int64(2) }
func (this ZendType) AllowNull() bool { return b.FlagMatch(this, 0x1) }

// todo 不明确作用
func (this ZendType) TypeName() string {
	//return (*byte)(this)
	return b.CastStrAuto(b.CastPtr[byte](this))
}

/**
 * ZendResource
 */
type ZendResource struct {
	handle int
	type_  int
	ptr    any
}

func NewZendResource(handle int, ptr any, type_ int) *ZendResource {
	var res = &ZendResource{
		handle: handle,
		type_:  type_,
		ptr:    ptr,
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
	ptr           []*PropertyInfo
}

func (this *ZendPropertyInfoList) GetNum() int                  { return this.num }
func (this *ZendPropertyInfoList) SetNum(value int)             { this.num = value }
func (this *ZendPropertyInfoList) GetNumAllocated() int         { return this.num_allocated }
func (this *ZendPropertyInfoList) SetNumAllocated(value int)    { this.num_allocated = value }
func (this *ZendPropertyInfoList) GetPtr() []*PropertyInfo      { return this.ptr }
func (this *ZendPropertyInfoList) SetPtr(value []*PropertyInfo) { this.ptr = value }

/**
 * ZendPropertyInfoSourceList
 */
type ZendPropertyInfoSourceList struct /* union */ {
	ptr  *PropertyInfo
	list uintptr
}

func (this *ZendPropertyInfoSourceList) GetPtr() *PropertyInfo      { return this.ptr }
func (this *ZendPropertyInfoSourceList) SetPtr(value *PropertyInfo) { this.ptr = value }
func (this *ZendPropertyInfoSourceList) GetList() uintptr           { return this.list }
func (this *ZendPropertyInfoSourceList) SetList(value uintptr)      { this.list = value }

/**
 * ZendReference
 */
type ZendReference struct {
	val     Zval
	sources ZendPropertyInfoSourceList
}

func NewZendReference(val *Zval) *ZendReference {
	var ref = &ZendReference{}

	ZVAL_COPY_VALUE(ref.GetVal(), val)
	ref.sources.SetPtr(nil)

	runtime.SetFinalizer(ref, func(ref *ZendReference) {
		b.Assert(ref.GetSources().GetPtr() != nil)
		zend.EfreeSize(ref, b.SizeOf("zend_reference"))
	})

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
	ast *zend.ZendAst
}

func NewAstRef(ast *zend.ZendAst) *ZendAstRef {
	b.Assert(ast != nil)

	// init
	var ref *ZendAstRef = &ZendAstRef{}
	zend.ZendAstTreeCopy(ast, ref.ast)

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
