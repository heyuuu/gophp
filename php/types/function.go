package types

import (
	"github.com/heyuuu/gophp/compile/ast"
)

type FunctionType uint8

const (
	TypeInternalFunction FunctionType = 1
	TypeUserFunction     FunctionType = 2
	TypeEvalCode         FunctionType = 4
)

// Function
type Function struct {
	typ      FunctionType `get:"Type"`
	flags    uint32       `prop:""`
	name     string       `prop:""`
	scope    *Class       `prop:""`
	argInfos []ArgInfo    `get:""`

	// fields for internal function
	moduleNumber int
	handler      any `get:""`

	// fields for user function
	stmts []ast.Stmt `get:""`

	blockInfo
}

func NewInternalFunction(name string, handler any, moduleNumber int) *Function {
	return &Function{
		typ:          TypeInternalFunction,
		name:         name,
		handler:      handler,
		moduleNumber: moduleNumber,
	}
}

func NewInternalFunctionByEntry(moduleNumber int, entry FunctionDecl) *Function {
	return &Function{
		typ:          TypeInternalFunction,
		name:         entry.name,
		handler:      entry.handler,
		argInfos:     entry.argInfos,
		moduleNumber: moduleNumber,
	}
}

func NewAstFunction(name string, argInfos []ArgInfo, stmts []ast.Stmt) *Function {
	return &Function{
		typ:      TypeUserFunction,
		name:     name,
		argInfos: argInfos,
		stmts:    stmts,
	}
}

func (f *Function) IsInternalFunction() bool { return f.typ == TypeInternalFunction }
func (f *Function) IsUserFunction() bool     { return f.typ == TypeUserFunction }
func (f *Function) IsEvalCode() bool         { return f.typ == TypeEvalCode }
func (f *Function) IsUserCode() bool         { return f.typ == TypeUserFunction || f.typ == TypeEvalCode }

/* flags */
func (f *Function) AddFnFlags(value uint32)      { f.flags |= value }
func (f *Function) SubFnFlags(value uint32)      { f.flags &^= value }
func (f *Function) HasFnFlags(value uint32) bool { return f.flags&value != 0 }
func (f *Function) SwitchFnFlags(value uint32, cond bool) {
	if cond {
		f.AddFnFlags(value)
	} else {
		f.SubFnFlags(value)
	}
}

func (f *Function) IsPublic() bool            { return f.HasFnFlags(AccPublic) }
func (f *Function) IsProtected() bool         { return f.HasFnFlags(AccProtected) }
func (f *Function) IsPrivate() bool           { return f.HasFnFlags(AccPrivate) }
func (f *Function) IsChanged() bool           { return f.HasFnFlags(AccChanged) }
func (f *Function) IsStatic() bool            { return f.HasFnFlags(AccStatic) }
func (f *Function) IsAbstract() bool          { return f.HasFnFlags(AccAbstract) }
func (f *Function) IsImmutable() bool         { return f.HasFnFlags(AccImmutable) }
func (f *Function) IsHasTypeHints() bool      { return f.HasFnFlags(AccHasTypeHints) }
func (f *Function) IsTopLevel() bool          { return f.HasFnFlags(AccTopLevel) }
func (f *Function) IsPreloaded() bool         { return f.HasFnFlags(AccPreloaded) }
func (f *Function) IsAllowStatic() bool       { return f.HasFnFlags(AccAllowStatic) }
func (f *Function) IsVariadic() bool          { return f.HasFnFlags(AccVariadic) }
func (f *Function) IsCallViaTrampoline() bool { return f.HasFnFlags(AccCallViaTrampoline) }
func (f *Function) IsFakeClosure() bool       { return f.HasFnFlags(AccFakeClosure) }
func (f *Function) IsUsesThis() bool          { return f.HasFnFlags(AccUsesThis) }
func (f *Function) IsGenerator() bool         { return f.HasFnFlags(AccGenerator) }
func (f *Function) IsHeapRtCache() bool       { return f.HasFnFlags(AccHeapRtCache) }
func (f *Function) IsUserArgInfo() bool       { return f.HasFnFlags(AccUserArgInfo) }
func (f *Function) IsClosure() bool           { return f.HasFnFlags(AccClosure) }
func (f *Function) IsStrictTypes() bool       { return f.HasFnFlags(AccStrictTypes) }
func (f *Function) IsDonePassTwo() bool       { return f.HasFnFlags(AccDonePassTwo) }
func (f *Function) IsDeprecated() bool        { return f.HasFnFlags(AccDeprecated) }
func (f *Function) IsFinal() bool             { return f.HasFnFlags(AccFinal) }
func (f *Function) IsCtor() bool              { return f.HasFnFlags(AccCtor) }
func (f *Function) IsReturnReference() bool   { return f.HasFnFlags(AccReturnReference) }
func (f *Function) IsDtor() bool              { return f.HasFnFlags(AccDtor) }
func (f *Function) IsEarlyBinding() bool      { return f.HasFnFlags(AccEarlyBinding) }
func (f *Function) IsHasFinallyBlock() bool   { return f.HasFnFlags(AccHasFinallyBlock) }
func (f *Function) IsTraitClone() bool        { return f.HasFnFlags(AccTraitClone) }

func (f *Function) SetIsPublic(b bool)            { f.SwitchFnFlags(AccPublic, b) }
func (f *Function) SetIsProtected(b bool)         { f.SwitchFnFlags(AccProtected, b) }
func (f *Function) SetIsPrivate(b bool)           { f.SwitchFnFlags(AccPrivate, b) }
func (f *Function) SetIsChanged(b bool)           { f.SwitchFnFlags(AccChanged, b) }
func (f *Function) SetIsStatic(b bool)            { f.SwitchFnFlags(AccStatic, b) }
func (f *Function) SetIsAbstract(b bool)          { f.SwitchFnFlags(AccAbstract, b) }
func (f *Function) SetIsVariadic(b bool)          { f.SwitchFnFlags(AccVariadic, b) }
func (f *Function) SetIsHasReturnType(b bool)     { f.SwitchFnFlags(AccHasReturnType, b) } // todo use SetReturnType()
func (f *Function) SetIsCallViaTrampoline(b bool) { f.SwitchFnFlags(AccCallViaTrampoline, b) }
func (f *Function) SetIsAllowStatic(b bool)       { f.SwitchFnFlags(AccAllowStatic, b) }
func (f *Function) SetIsFakeClosure(b bool)       { f.SwitchFnFlags(AccFakeClosure, b) }
func (f *Function) SetIsUsesThis(b bool)          { f.SwitchFnFlags(AccUsesThis, b) }
func (f *Function) SetIsGenerator(b bool)         { f.SwitchFnFlags(AccGenerator, b) }
func (f *Function) SetIsHeapRtCache(b bool)       { f.SwitchFnFlags(AccHeapRtCache, b) }
func (f *Function) SetIsUserArgInfo(b bool)       { f.SwitchFnFlags(AccUserArgInfo, b) }
func (f *Function) SetIsClosure(b bool)           { f.SwitchFnFlags(AccClosure, b) }
func (f *Function) SetIsImmutable(b bool)         { f.SwitchFnFlags(AccImmutable, b) }
func (f *Function) SetIsStrictTypes(b bool)       { f.SwitchFnFlags(AccStrictTypes, b) }
func (f *Function) SetIsPreloaded(b bool)         { f.SwitchFnFlags(AccPreloaded, b) }
func (f *Function) SetIsDonePassTwo(b bool)       { f.SwitchFnFlags(AccDonePassTwo, b) }
func (f *Function) SetIsDeprecated(b bool)        { f.SwitchFnFlags(AccDeprecated, b) }
func (f *Function) SetIsFinal(b bool)             { f.SwitchFnFlags(AccFinal, b) }
func (f *Function) SetIsCtor(b bool)              { f.SwitchFnFlags(AccCtor, b) }
func (f *Function) SetIsReturnReference(b bool)   { f.SwitchFnFlags(AccReturnReference, b) }
func (f *Function) SetIsHasTypeHints(b bool)      { f.SwitchFnFlags(AccHasTypeHints, b) }
func (f *Function) SetIsDtor(b bool)              { f.SwitchFnFlags(AccDtor, b) }
func (f *Function) SetIsEarlyBinding(b bool)      { f.SwitchFnFlags(AccEarlyBinding, b) }
func (f *Function) SetIsHasFinallyBlock(b bool)   { f.SwitchFnFlags(AccHasFinallyBlock, b) }
func (f *Function) SetIsTopLevel(b bool)          { f.SwitchFnFlags(AccTopLevel, b) }
func (f *Function) SetIsTraitClone(b bool)        { f.SwitchFnFlags(AccTraitClone, b) }

func (f *Function) CalleeName() string {
	if f.name == "" {
		return "main"
	} else if f.scope == nil {
		return f.name
	} else {
		return f.scope.Name() + "::" + f.name
	}
}
