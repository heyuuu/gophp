package types

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
)

/**
 * ZendResource
 */
type Resource struct {
	handle int
	type_  int
	ptr    any
}

func NewZendResource(handle int, ptr any, type_ int) *Resource {
	var res = &Resource{
		handle: handle,
		type_:  type_,
		ptr:    ptr,
	}

	return res
}

func (this *Resource) GetHandle() int      { return this.handle }
func (this *Resource) SetHandle(value int) { this.handle = value }
func (this *Resource) GetType() int        { return this.type_ }
func (this *Resource) SetType(value int)   { this.type_ = value }
func (this *Resource) GetPtr() any         { return this.ptr }
func (this *Resource) SetPtr(value any)    { this.ptr = value }

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
 * Reference
 */
type Reference struct {
	val     Zval
	sources ZendPropertyInfoSourceList
}

func NewZendReference(val *Zval) *Reference {
	var ref = &Reference{}

	ZVAL_COPY_VALUE(ref.GetVal(), val)
	ref.sources.SetPtr(nil)

	return ref
}

func (this *Reference) GetVal() *Zval                               { return &this.val }
func (this *Reference) SetVal(value Zval)                           { this.val = value }
func (this *Reference) GetSources() *ZendPropertyInfoSourceList     { return &this.sources }
func (this *Reference) SetSources(value ZendPropertyInfoSourceList) { this.sources = value }

/**
 * ZendAstRef
 */
type ZendAstRef struct {
	ast *zend.ZendAst
}

func NewAstRef(ast *zend.ZendAst) *ZendAstRef {
	b.Assert(ast != nil)

	return &ZendAstRef{
		ast: ast.TreeCopy(),
	}
}

func (this ZendAstRef) GcAst() *zend.ZendAst {
	//func GC_AST(p *ZendAstRef) *ZendAst {
	//	return (*ZendAst)((*byte)(p) + b.SizeOf("zend_ast_ref"))
	//}
	return this.ast
}
