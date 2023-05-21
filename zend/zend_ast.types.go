package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

/**
 * ZendAst
 */
type ZendAst struct {
	kind   ZendAstKind
	attr   ZendAstAttr
	lineno uint32
	child  []*ZendAst
}

func (this *ZendAst) GetKind() ZendAstKind      { return this.kind }
func (this *ZendAst) SetKind(value ZendAstKind) { this.kind = value }
func (this *ZendAst) GetAttr() ZendAstAttr      { return this.attr }
func (this *ZendAst) SetAttr(value ZendAstAttr) { this.attr = value }
func (this *ZendAst) GetLineno() uint32         { return this.lineno }
func (this *ZendAst) SetLineno(value uint32)    { this.lineno = value }
func (this *ZendAst) GetChild() []*ZendAst      { return this.child }

// func (this *ZendAst) SetChild(value []*ZendAst) { this.child = value }

/**
 * ZendAstList
 */
type ZendAstList struct {
	kind     ZendAstKind
	attr     ZendAstAttr
	lineno   uint32
	children uint32
	child    []*ZendAst
}

// func MakeZendAstList(kind ZendAstKind, attr ZendAstAttr, lineno uint32, children uint32, child []*ZendAst) ZendAstList {
//     return ZendAstList{
//         kind:kind,
//         attr:attr,
//         lineno:lineno,
//         children:children,
//         child:child,
//     }
// }
func (this *ZendAstList) GetKind() ZendAstKind      { return this.kind }
func (this *ZendAstList) SetKind(value ZendAstKind) { this.kind = value }
func (this *ZendAstList) GetAttr() ZendAstAttr      { return this.attr }
func (this *ZendAstList) SetAttr(value ZendAstAttr) { this.attr = value }

// func (this *ZendAstList)  GetLineno() uint32      { return this.lineno }
func (this *ZendAstList) SetLineno(value uint32)   { this.lineno = value }
func (this *ZendAstList) GetChildren() uint32      { return this.children }
func (this *ZendAstList) SetChildren(value uint32) { this.children = value }
func (this *ZendAstList) GetChild() []*ZendAst     { return this.child }

// func (this *ZendAstList) SetChild(value []*ZendAst) { this.child = value }

/**
 * ZendAstZval
 */
type ZendAstZval struct {
	kind ZendAstKind
	attr ZendAstAttr
	val  types.Zval
	// 新增 lineno，从 val.u2 中提出来
	lineno uint32
}

func NewZendAstZval(kind ZendAstKind, attr ZendAstAttr, zv *types.Zval, lineno uint32) *ZendAstZval {
	ast := &ZendAstZval{kind: kind, attr: attr, lineno: lineno}
	ast.val.CopyValueFrom(zv)
	return ast
}

func (this *ZendAstZval) GetKind() ZendAstKind      { return this.kind }
func (this *ZendAstZval) SetKind(value ZendAstKind) { this.kind = value }
func (this *ZendAstZval) GetAttr() ZendAstAttr      { return this.attr }
func (this *ZendAstZval) SetAttr(value ZendAstAttr) { this.attr = value }
func (this *ZendAstZval) GetVal() *types.Zval       { return &this.val }
func (this *ZendAstZval) GetLineno() uint32         { return this.lineno }

// func (this *ZendAstZval) SetVal(value Zval) { this.val = value }

/**
 * ZendAstDecl
 */
type ZendAstDecl struct {
	kind         ZendAstKind
	attr         ZendAstAttr
	start_lineno uint32
	end_lineno   uint32
	flags        uint32
	lex_pos      *uint8
	doc_comment  *types.String
	name         *types.String
	child        []*ZendAst
}

//             func MakeZendAstDecl(
// kind ZendAstKind,
// attr ZendAstAttr,
// start_lineno uint32,
// end_lineno uint32,
// flags uint32,
// lex_pos *uint8,
// doc_comment *String,
// name *String,
// child []*ZendAst,
// ) ZendAstDecl {
//                 return ZendAstDecl{
//                     kind:kind,
//                     attr:attr,
//                     start_lineno:start_lineno,
//                     end_lineno:end_lineno,
//                     flags:flags,
//                     lex_pos:lex_pos,
//                     doc_comment:doc_comment,
//                     name:name,
//                     child:child,
//                 }
//             }
func (this *ZendAstDecl) GetKind() ZendAstKind      { return this.kind }
func (this *ZendAstDecl) SetKind(value ZendAstKind) { this.kind = value }

// func (this *ZendAstDecl)  GetAttr() ZendAstAttr      { return this.attr }
func (this *ZendAstDecl) SetAttr(value ZendAstAttr)   { this.attr = value }
func (this *ZendAstDecl) GetStartLineno() uint32      { return this.start_lineno }
func (this *ZendAstDecl) SetStartLineno(value uint32) { this.start_lineno = value }
func (this *ZendAstDecl) GetEndLineno() uint32        { return this.end_lineno }
func (this *ZendAstDecl) SetEndLineno(value uint32)   { this.end_lineno = value }
func (this *ZendAstDecl) GetFlags() uint32            { return this.flags }
func (this *ZendAstDecl) SetFlags(value uint32)       { this.flags = value }

// func (this *ZendAstDecl)  GetLexPos() *uint8      { return this.lex_pos }
func (this *ZendAstDecl) SetLexPos(value *uint8)            { this.lex_pos = value }
func (this *ZendAstDecl) GetDocComment() *types.String      { return this.doc_comment }
func (this *ZendAstDecl) SetDocComment(value *types.String) { this.doc_comment = value }
func (this *ZendAstDecl) GetName() *types.String            { return this.name }
func (this *ZendAstDecl) SetName(value *types.String)       { this.name = value }
func (this *ZendAstDecl) GetChild() []*ZendAst              { return this.child }

// func (this *ZendAstDecl) SetChild(value []*ZendAst) { this.child = value }

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
