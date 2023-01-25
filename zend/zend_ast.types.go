// <<generate>>

package zend

/**
 * ZendAst
 */
type ZendAst struct {
	kind   ZendAstKind
	attr   ZendAstAttr
	lineno uint32
	child  []*ZendAst
}

func (this ZendAst) GetKind() ZendAstKind       { return this.kind }
func (this *ZendAst) SetKind(value ZendAstKind) { this.kind = value }
func (this ZendAst) GetAttr() ZendAstAttr       { return this.attr }
func (this *ZendAst) SetAttr(value ZendAstAttr) { this.attr = value }
func (this ZendAst) GetLineno() uint32          { return this.lineno }
func (this *ZendAst) SetLineno(value uint32)    { this.lineno = value }
func (this ZendAst) GetChild() []*ZendAst       { return this.child }
func (this *ZendAst) SetChild(value []*ZendAst) { this.child = value }

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

func (this ZendAstList) GetKind() ZendAstKind       { return this.kind }
func (this *ZendAstList) SetKind(value ZendAstKind) { this.kind = value }
func (this ZendAstList) GetAttr() ZendAstAttr       { return this.attr }
func (this *ZendAstList) SetAttr(value ZendAstAttr) { this.attr = value }
func (this ZendAstList) GetLineno() uint32          { return this.lineno }
func (this *ZendAstList) SetLineno(value uint32)    { this.lineno = value }
func (this ZendAstList) GetChildren() uint32        { return this.children }
func (this *ZendAstList) SetChildren(value uint32)  { this.children = value }
func (this ZendAstList) GetChild() []*ZendAst       { return this.child }
func (this *ZendAstList) SetChild(value []*ZendAst) { this.child = value }

/**
 * ZendAstZval
 */
type ZendAstZval struct {
	kind ZendAstKind
	attr ZendAstAttr
	val  Zval
}

func (this ZendAstZval) GetKind() ZendAstKind       { return this.kind }
func (this *ZendAstZval) SetKind(value ZendAstKind) { this.kind = value }
func (this ZendAstZval) GetAttr() ZendAstAttr       { return this.attr }
func (this *ZendAstZval) SetAttr(value ZendAstAttr) { this.attr = value }
func (this ZendAstZval) GetVal() Zval               { return this.val }
func (this *ZendAstZval) SetVal(value Zval)         { this.val = value }

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
	doc_comment  *ZendString
	name         *ZendString
	child        []*ZendAst
}

func (this ZendAstDecl) GetKind() ZendAstKind             { return this.kind }
func (this *ZendAstDecl) SetKind(value ZendAstKind)       { this.kind = value }
func (this ZendAstDecl) GetAttr() ZendAstAttr             { return this.attr }
func (this *ZendAstDecl) SetAttr(value ZendAstAttr)       { this.attr = value }
func (this ZendAstDecl) GetStartLineno() uint32           { return this.start_lineno }
func (this *ZendAstDecl) SetStartLineno(value uint32)     { this.start_lineno = value }
func (this ZendAstDecl) GetEndLineno() uint32             { return this.end_lineno }
func (this *ZendAstDecl) SetEndLineno(value uint32)       { this.end_lineno = value }
func (this ZendAstDecl) GetFlags() uint32                 { return this.flags }
func (this *ZendAstDecl) SetFlags(value uint32)           { this.flags = value }
func (this ZendAstDecl) GetLexPos() *uint8                { return this.lex_pos }
func (this *ZendAstDecl) SetLexPos(value *uint8)          { this.lex_pos = value }
func (this ZendAstDecl) GetDocComment() *ZendString       { return this.doc_comment }
func (this *ZendAstDecl) SetDocComment(value *ZendString) { this.doc_comment = value }
func (this ZendAstDecl) GetName() *ZendString             { return this.name }
func (this *ZendAstDecl) SetName(value *ZendString)       { this.name = value }
func (this ZendAstDecl) GetChild() []*ZendAst             { return this.child }
func (this *ZendAstDecl) SetChild(value []*ZendAst)       { this.child = value }

/* ZendAstDecl.flags */
func (this *ZendAstDecl) AddFlags(value uint32)     { this.flags |= value }
func (this *ZendAstDecl) SubFlags(value uint32)     { this.flags &^= value }
func (this ZendAstDecl) HasFlags(value uint32) bool { return this.flags&value != 0 }
func (this *ZendAstDecl) SwitchFlags(value uint32, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}
func (this ZendAstDecl) isPublic() bool          { return this.HasFlags(ZEND_ACC_PUBLIC) }
func (this ZendAstDecl) isProtected() bool       { return this.HasFlags(ZEND_ACC_PROTECTED) }
func (this ZendAstDecl) isPrivate() bool         { return this.HasFlags(ZEND_ACC_PRIVATE) }
func (this ZendAstDecl) isStatic() bool          { return this.HasFlags(ZEND_ACC_STATIC) }
func (this ZendAstDecl) isAbstract() bool        { return this.HasFlags(ZEND_ACC_ABSTRACT) }
func (this ZendAstDecl) isFinal() bool           { return this.HasFlags(ZEND_ACC_FINAL) }
func (this ZendAstDecl) isReturnReference() bool { return this.HasFlags(ZEND_ACC_RETURN_REFERENCE) }
func (this ZendAstDecl) isInterface() bool       { return this.HasFlags(ZEND_ACC_INTERFACE) }
func (this ZendAstDecl) isTrait() bool           { return this.HasFlags(ZEND_ACC_TRAIT) }
func (this ZendAstDecl) isExplicitAbstractClass() bool {
	return this.HasFlags(ZEND_ACC_EXPLICIT_ABSTRACT_CLASS)
}
func (this ZendAstDecl) isAnonClass() bool         { return this.HasFlags(ZEND_ACC_ANON_CLASS) }
func (this *ZendAstDecl) setIsPublic(cond bool)    { this.SwitchFlags(ZEND_ACC_PUBLIC, cond) }
func (this *ZendAstDecl) setIsProtected(cond bool) { this.SwitchFlags(ZEND_ACC_PROTECTED, cond) }
func (this *ZendAstDecl) setIsPrivate(cond bool)   { this.SwitchFlags(ZEND_ACC_PRIVATE, cond) }
func (this *ZendAstDecl) setIsStatic(cond bool)    { this.SwitchFlags(ZEND_ACC_STATIC, cond) }
func (this *ZendAstDecl) setIsAbstract(cond bool)  { this.SwitchFlags(ZEND_ACC_ABSTRACT, cond) }
func (this *ZendAstDecl) setIsFinal(cond bool)     { this.SwitchFlags(ZEND_ACC_FINAL, cond) }
func (this *ZendAstDecl) setIsReturnReference(cond bool) {
	this.SwitchFlags(ZEND_ACC_RETURN_REFERENCE, cond)
}
func (this *ZendAstDecl) setIsInterface(cond bool) { this.SwitchFlags(ZEND_ACC_INTERFACE, cond) }
func (this *ZendAstDecl) setIsTrait(cond bool)     { this.SwitchFlags(ZEND_ACC_TRAIT, cond) }
func (this *ZendAstDecl) setIsExplicitAbstractClass(cond bool) {
	this.SwitchFlags(ZEND_ACC_EXPLICIT_ABSTRACT_CLASS, cond)
}
func (this *ZendAstDecl) setIsAnonClass(cond bool) { this.SwitchFlags(ZEND_ACC_ANON_CLASS, cond) }
