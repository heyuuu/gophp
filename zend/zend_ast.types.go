package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/utils/slices"
)

/**
 * ZendAst
 */
type ZendAstList = ZendAst
type ZendAst struct {
	kind   ZendAstKind
	attr   ZendAstAttr
	lineno uint32
	child  []*ZendAst
}

func NewAst(kind ZendAstKind, attr ZendAstAttr, lineno uint32, children []*ZendAst) *ZendAst {
	b.Assert(int(kind)>>ZEND_AST_NUM_CHILDREN_SHIFT == len(children))
	return &ZendAst{
		kind:   kind,
		attr:   attr,
		lineno: lineno,
		child:  children,
	}
}
func CopyAst(old *ZendAst, childCopy func(child *ZendAst) *ZendAst) *ZendAst {
	child := slices.Map(old.child, childCopy)
	return NewAst(old.kind, old.attr, 0, child)
}

func (ast *ZendAst) IsSpecial() bool { return ast.kind>>ZEND_AST_SPECIAL_SHIFT&1 != 0 }
func (ast *ZendAst) IsList() bool    { return ast.kind>>ZEND_AST_IS_LIST_SHIFT&1 != 0 }
func (ast *ZendAst) IsZval() bool    { return ast.kind == ZEND_AST_ZVAL }

func (ast *ZendAst) AsAstList() *ZendAstList {
	b.Assert(ast.IsList())
	return ast
}
func (ast *ZendAst) AsAstZval() *ZendAstZval {
	b.Assert(ast.kind == ZEND_AST_ZVAL)
	return (*ZendAstZval)(ast)
}

// methods for List type
func (ast *ZendAst) AddChild(child *ZendAst) {
	b.Assert(ast.IsList())
	ast.child = append(ast.child, child)
}

// fields
func (ast *ZendAst) GetKind() ZendAstKind      { return ast.kind }
func (ast *ZendAst) SetKind(value ZendAstKind) { ast.kind = value }
func (ast *ZendAst) GetAttr() ZendAstAttr      { return ast.attr }
func (ast *ZendAst) SetAttr(value ZendAstAttr) { ast.attr = value }
func (ast *ZendAst) GetLineno() uint32         { return ast.lineno }
func (ast *ZendAst) SetLineno(value uint32)    { ast.lineno = value }
func (ast *ZendAst) GetChild() []*ZendAst      { return ast.child }
func (ast *ZendAst) GetChildren() uint32       { return uint32(len(ast.child)) }

/**
 * ZendAstZval
 */
type ZendAstZval struct {
	kind ZendAstKind
	attr ZendAstAttr
	// 新增 lineno，从 val.u2 中提出来
	lineno uint32
	val    types.Zval
}

func NewAstZval(kind ZendAstKind, attr ZendAstAttr, zv *types.Zval, lineno uint32) *ZendAstZval {
	ast := &ZendAstZval{kind: kind, attr: attr, lineno: lineno}
	ast.val.CopyValueFrom(zv)
	return ast
}

func CopyAstZval(old *ZendAstZval) *ZendAstZval {
	ast := &ZendAstZval{
		kind: ZEND_AST_ZVAL,
		attr: old.attr,
	}
	ast.val.CopyFrom(&old.val)
	return ast
}

func (this *ZendAstZval) GetKind() ZendAstKind      { return this.kind }
func (this *ZendAstZval) SetKind(value ZendAstKind) { this.kind = value }
func (this *ZendAstZval) GetAttr() ZendAstAttr      { return this.attr }
func (this *ZendAstZval) SetAttr(value ZendAstAttr) { this.attr = value }
func (this *ZendAstZval) GetVal() *types.Zval       { return &this.val }
func (this *ZendAstZval) GetLineno() uint32         { return this.lineno }

/**
 * ZendAstDecl
 */
type ZendAstDecl struct {
	kind        ZendAstKind
	attr        ZendAstAttr
	startLineno uint32
	endLineno   uint32
	flags       uint32
	lexPos      *uint8
	docComment  *types.String
	name        *types.String
	child       []*ZendAst
}

func (this *ZendAstDecl) GetKind() ZendAstKind              { return this.kind }
func (this *ZendAstDecl) SetKind(value ZendAstKind)         { this.kind = value }
func (this *ZendAstDecl) SetAttr(value ZendAstAttr)         { this.attr = value }
func (this *ZendAstDecl) GetStartLineno() uint32            { return this.startLineno }
func (this *ZendAstDecl) SetStartLineno(value uint32)       { this.startLineno = value }
func (this *ZendAstDecl) GetEndLineno() uint32              { return this.endLineno }
func (this *ZendAstDecl) SetEndLineno(value uint32)         { this.endLineno = value }
func (this *ZendAstDecl) GetFlags() uint32                  { return this.flags }
func (this *ZendAstDecl) SetFlags(value uint32)             { this.flags = value }
func (this *ZendAstDecl) SetLexPos(value *uint8)            { this.lexPos = value }
func (this *ZendAstDecl) GetDocComment() *types.String      { return this.docComment }
func (this *ZendAstDecl) SetDocComment(value *types.String) { this.docComment = value }
func (this *ZendAstDecl) GetName() *types.String            { return this.name }
func (this *ZendAstDecl) SetName(value *types.String)       { this.name = value }
func (this *ZendAstDecl) GetChild() []*ZendAst              { return this.child }

/* ZendAstDecl.flags */
func (this *ZendAstDecl) AddFlags(value uint32)      { this.flags |= value }
func (this *ZendAstDecl) SubFlags(value uint32)      { this.flags &^= value }
func (this *ZendAstDecl) HasFlags(value uint32) bool { return this.flags&value != 0 }
func (this *ZendAstDecl) SwitchFlags(value uint32, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}
func (this ZendAstDecl) IsPublic() bool          { return this.HasFlags(types.AccPublic) }
func (this ZendAstDecl) IsProtected() bool       { return this.HasFlags(types.AccProtected) }
func (this ZendAstDecl) IsPrivate() bool         { return this.HasFlags(types.AccPrivate) }
func (this ZendAstDecl) IsStatic() bool          { return this.HasFlags(types.AccStatic) }
func (this ZendAstDecl) IsAbstract() bool        { return this.HasFlags(types.AccAbstract) }
func (this ZendAstDecl) IsFinal() bool           { return this.HasFlags(types.AccFinal) }
func (this ZendAstDecl) IsReturnReference() bool { return this.HasFlags(types.AccReturnReference) }
func (this ZendAstDecl) IsInterface() bool       { return this.HasFlags(types.AccInterface) }
func (this ZendAstDecl) IsTrait() bool           { return this.HasFlags(types.AccTrait) }
func (this ZendAstDecl) IsExplicitAbstractClass() bool {
	return this.HasFlags(types.AccExplicitAbstractClass)
}
func (this ZendAstDecl) IsAnonClass() bool         { return this.HasFlags(types.AccAnonClass) }
func (this *ZendAstDecl) SetIsPublic(cond bool)    { this.SwitchFlags(types.AccPublic, cond) }
func (this *ZendAstDecl) SetIsProtected(cond bool) { this.SwitchFlags(types.AccProtected, cond) }
func (this *ZendAstDecl) SetIsPrivate(cond bool)   { this.SwitchFlags(types.AccPrivate, cond) }
func (this *ZendAstDecl) SetIsStatic(cond bool)    { this.SwitchFlags(types.AccStatic, cond) }
func (this *ZendAstDecl) SetIsAbstract(cond bool)  { this.SwitchFlags(types.AccAbstract, cond) }
func (this *ZendAstDecl) SetIsFinal(cond bool)     { this.SwitchFlags(types.AccFinal, cond) }
func (this *ZendAstDecl) SetIsReturnReference(cond bool) {
	this.SwitchFlags(types.AccReturnReference, cond)
}
func (this *ZendAstDecl) SetIsInterface(cond bool) { this.SwitchFlags(types.AccInterface, cond) }
func (this *ZendAstDecl) SetIsTrait(cond bool)     { this.SwitchFlags(types.AccTrait, cond) }
func (this *ZendAstDecl) SetIsExplicitAbstractClass(cond bool) {
	this.SwitchFlags(types.AccExplicitAbstractClass, cond)
}
func (this *ZendAstDecl) SetIsAnonClass(cond bool) { this.SwitchFlags(types.AccAnonClass, cond) }
