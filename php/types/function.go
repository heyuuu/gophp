package types

import "github.com/heyuuu/gophp/zend"

type IFunction interface {
	// type cast
	GetOpArray() *ZendOpArray
	GetInternalFunction() *InternalFunction

	// common fields
	GetType() uint8
	SetType(typ uint8)
	GetFnFlags() uint32
	SetFnFlags(flags uint32)
	GetFunctionName() *String
	SetFunctionName(value *String)
	GetScope() *ClassEntry
	SetScope(value *ClassEntry)
	GetPrototype() IFunction
	SetPrototype(value IFunction)
	GetNumArgs() uint32
	SetNumArgs(value uint32)
	GetRequiredNumArgs() uint32
	SetRequiredNumArgs(value uint32)
	GetArgInfo() []zend.ArgInfo
	SetArgInfo(value []zend.ArgInfo)

	/* fnFlags */
	AddFnFlags(value uint32)
	SubFnFlags(value uint32)
	HasFnFlags(value uint32) bool
	SwitchFnFlags(value uint32, cond bool)

	IsPublic() bool
	IsProtected() bool
	IsPrivate() bool
	IsChanged() bool
	IsStatic() bool
	IsAbstract() bool
	IsImmutable() bool
	IsHasTypeHints() bool
	IsTopLevel() bool
	IsPreloaded() bool
	IsAllowStatic() bool
	IsVariadic() bool
	IsHasReturnType() bool
	IsCallViaTrampoline() bool
	IsFakeClosure() bool
	IsUsesThis() bool
	IsGenerator() bool
	IsHeapRtCache() bool
	IsUserArgInfo() bool
	IsClosure() bool
	IsStrictTypes() bool
	IsDonePassTwo() bool
	IsDeprecated() bool
	IsFinal() bool
	IsCtor() bool
	IsReturnReference() bool
	IsArenaAllocated() bool
	IsDtor() bool
	IsEarlyBinding() bool
	IsHasFinallyBlock() bool
	IsTraitClone() bool

	SetIsPublic(b bool)
	SetIsProtected(b bool)
	SetIsPrivate(b bool)
	SetIsChanged(b bool)
	SetIsStatic(b bool)
	SetIsAbstract(b bool)
	SetIsVariadic(b bool)
	SetIsHasReturnType(b bool)
	SetIsCallViaTrampoline(b bool)
	SetIsAllowStatic(b bool)
	SetIsFakeClosure(b bool)
	SetIsUsesThis(b bool)
	SetIsGenerator(b bool)
	SetIsHeapRtCache(b bool)
	SetIsUserArgInfo(b bool)
	SetIsClosure(b bool)
	SetIsImmutable(b bool)
	SetIsStrictTypes(b bool)
	SetIsPreloaded(b bool)
	SetIsDonePassTwo(b bool)
	SetIsDeprecated(b bool)
	SetIsFinal(b bool)
	SetIsCtor(b bool)
	SetIsReturnReference(b bool)
	SetIsArenaAllocated(b bool)
	SetIsHasTypeHints(b bool)
	SetIsDtor(b bool)
	SetIsEarlyBinding(b bool)
	SetIsHasFinallyBlock(b bool)
	SetIsTopLevel(b bool)
	SetIsTraitClone(b bool)

	//
	CheckArgSendType(argNum1 uint32, mask uint8) bool
}

type functionHeader struct {
	// type(8) + argFlags(24) 的组合数据
	typ             uint8
	fnFlags         uint32
	functionName    *String
	scope           *ClassEntry
	prototype       IFunction
	numArgs         uint32
	requiredNumArgs uint32
	argInfos        []zend.ArgInfo
}

var _ IFunction = (*functionHeader)(nil)

func (f *functionHeader) GetOpArray() *ZendOpArray {
	panic("*functionHeader is not *ZendOpArray")
}
func (f *functionHeader) GetInternalFunction() *InternalFunction {
	panic("*functionHeader is not *InternalFunction")
}

func (f *functionHeader) GetType() uint8                  { return f.typ }
func (f *functionHeader) SetType(typ uint8)               { f.typ = typ }
func (f *functionHeader) GetFnFlags() uint32              { return f.fnFlags }
func (f *functionHeader) SetFnFlags(flags uint32)         { f.fnFlags = flags }
func (f *functionHeader) GetFunctionName() *String        { return f.functionName }
func (f *functionHeader) SetFunctionName(value *String)   { f.functionName = value }
func (f *functionHeader) GetScope() *ClassEntry           { return f.scope }
func (f *functionHeader) SetScope(value *ClassEntry)      { f.scope = value }
func (f *functionHeader) GetPrototype() IFunction         { return f.prototype }
func (f *functionHeader) SetPrototype(value IFunction)    { f.prototype = value }
func (f *functionHeader) GetNumArgs() uint32              { return uint32(len(f.argInfos)) }
func (f *functionHeader) SetNumArgs(value uint32)         { f.numArgs = value }
func (f *functionHeader) GetRequiredNumArgs() uint32      { return f.requiredNumArgs }
func (f *functionHeader) SetRequiredNumArgs(value uint32) { f.requiredNumArgs = value }
func (f *functionHeader) GetArgInfo() []zend.ArgInfo      { return f.argInfos }
func (f *functionHeader) SetArgInfo(value []zend.ArgInfo) { f.argInfos = value }

/* fnFlags */
func (f *functionHeader) AddFnFlags(value uint32)      { f.fnFlags |= value }
func (f *functionHeader) SubFnFlags(value uint32)      { f.fnFlags &^= value }
func (f *functionHeader) HasFnFlags(value uint32) bool { return f.fnFlags&value != 0 }
func (f *functionHeader) SwitchFnFlags(value uint32, cond bool) {
	if cond {
		f.AddFnFlags(value)
	} else {
		f.SubFnFlags(value)
	}
}

func (f *functionHeader) IsPublic() bool            { return f.HasFnFlags(AccPublic) }
func (f *functionHeader) IsProtected() bool         { return f.HasFnFlags(AccProtected) }
func (f *functionHeader) IsPrivate() bool           { return f.HasFnFlags(AccPrivate) }
func (f *functionHeader) IsChanged() bool           { return f.HasFnFlags(AccChanged) }
func (f *functionHeader) IsStatic() bool            { return f.HasFnFlags(AccStatic) }
func (f *functionHeader) IsAbstract() bool          { return f.HasFnFlags(AccAbstract) }
func (f *functionHeader) IsImmutable() bool         { return f.HasFnFlags(AccImmutable) }
func (f *functionHeader) IsHasTypeHints() bool      { return f.HasFnFlags(AccHasTypeHints) }
func (f *functionHeader) IsTopLevel() bool          { return f.HasFnFlags(AccTopLevel) }
func (f *functionHeader) IsPreloaded() bool         { return f.HasFnFlags(AccPreloaded) }
func (f *functionHeader) IsAllowStatic() bool       { return f.HasFnFlags(AccAllowStatic) }
func (f *functionHeader) IsVariadic() bool          { return f.HasFnFlags(AccVariadic) }
func (f *functionHeader) IsHasReturnType() bool     { return f.HasFnFlags(AccHasReturnType) }
func (f *functionHeader) IsCallViaTrampoline() bool { return f.HasFnFlags(AccCallViaTrampoline) }
func (f *functionHeader) IsFakeClosure() bool       { return f.HasFnFlags(AccFakeClosure) }
func (f *functionHeader) IsUsesThis() bool          { return f.HasFnFlags(AccUsesThis) }
func (f *functionHeader) IsGenerator() bool         { return f.HasFnFlags(AccGenerator) }
func (f *functionHeader) IsHeapRtCache() bool       { return f.HasFnFlags(AccHeapRtCache) }
func (f *functionHeader) IsUserArgInfo() bool       { return f.HasFnFlags(AccUserArgInfo) }
func (f *functionHeader) IsClosure() bool           { return f.HasFnFlags(AccClosure) }
func (f *functionHeader) IsStrictTypes() bool       { return f.HasFnFlags(AccStrictTypes) }
func (f *functionHeader) IsDonePassTwo() bool       { return f.HasFnFlags(AccDonePassTwo) }
func (f *functionHeader) IsDeprecated() bool        { return f.HasFnFlags(AccDeprecated) }
func (f *functionHeader) IsFinal() bool             { return f.HasFnFlags(AccFinal) }
func (f *functionHeader) IsCtor() bool              { return f.HasFnFlags(AccCtor) }
func (f *functionHeader) IsReturnReference() bool   { return f.HasFnFlags(AccReturnReference) }
func (f *functionHeader) IsArenaAllocated() bool    { return f.HasFnFlags(AccArenaAllocated) }
func (f *functionHeader) IsDtor() bool              { return f.HasFnFlags(AccDtor) }
func (f *functionHeader) IsEarlyBinding() bool      { return f.HasFnFlags(AccEarlyBinding) }
func (f *functionHeader) IsHasFinallyBlock() bool   { return f.HasFnFlags(AccHasFinallyBlock) }
func (f *functionHeader) IsTraitClone() bool        { return f.HasFnFlags(AccTraitClone) }

func (f *functionHeader) SetIsPublic(b bool)        { f.SwitchFnFlags(AccPublic, b) }
func (f *functionHeader) SetIsProtected(b bool)     { f.SwitchFnFlags(AccProtected, b) }
func (f *functionHeader) SetIsPrivate(b bool)       { f.SwitchFnFlags(AccPrivate, b) }
func (f *functionHeader) SetIsChanged(b bool)       { f.SwitchFnFlags(AccChanged, b) }
func (f *functionHeader) SetIsStatic(b bool)        { f.SwitchFnFlags(AccStatic, b) }
func (f *functionHeader) SetIsAbstract(b bool)      { f.SwitchFnFlags(AccAbstract, b) }
func (f *functionHeader) SetIsVariadic(b bool)      { f.SwitchFnFlags(AccVariadic, b) }
func (f *functionHeader) SetIsHasReturnType(b bool) { f.SwitchFnFlags(AccHasReturnType, b) }
func (f *functionHeader) SetIsCallViaTrampoline(b bool) {
	f.SwitchFnFlags(AccCallViaTrampoline, b)
}
func (f *functionHeader) SetIsAllowStatic(b bool)     { f.SwitchFnFlags(AccAllowStatic, b) }
func (f *functionHeader) SetIsFakeClosure(b bool)     { f.SwitchFnFlags(AccFakeClosure, b) }
func (f *functionHeader) SetIsUsesThis(b bool)        { f.SwitchFnFlags(AccUsesThis, b) }
func (f *functionHeader) SetIsGenerator(b bool)       { f.SwitchFnFlags(AccGenerator, b) }
func (f *functionHeader) SetIsHeapRtCache(b bool)     { f.SwitchFnFlags(AccHeapRtCache, b) }
func (f *functionHeader) SetIsUserArgInfo(b bool)     { f.SwitchFnFlags(AccUserArgInfo, b) }
func (f *functionHeader) SetIsClosure(b bool)         { f.SwitchFnFlags(AccClosure, b) }
func (f *functionHeader) SetIsImmutable(b bool)       { f.SwitchFnFlags(AccImmutable, b) }
func (f *functionHeader) SetIsStrictTypes(b bool)     { f.SwitchFnFlags(AccStrictTypes, b) }
func (f *functionHeader) SetIsPreloaded(b bool)       { f.SwitchFnFlags(AccPreloaded, b) }
func (f *functionHeader) SetIsDonePassTwo(b bool)     { f.SwitchFnFlags(AccDonePassTwo, b) }
func (f *functionHeader) SetIsDeprecated(b bool)      { f.SwitchFnFlags(AccDeprecated, b) }
func (f *functionHeader) SetIsFinal(b bool)           { f.SwitchFnFlags(AccFinal, b) }
func (f *functionHeader) SetIsCtor(b bool)            { f.SwitchFnFlags(AccCtor, b) }
func (f *functionHeader) SetIsReturnReference(b bool) { f.SwitchFnFlags(AccReturnReference, b) }
func (f *functionHeader) SetIsArenaAllocated(b bool)  { f.SwitchFnFlags(AccArenaAllocated, b) }
func (f *functionHeader) SetIsHasTypeHints(b bool)    { f.SwitchFnFlags(AccHasTypeHints, b) }
func (f *functionHeader) SetIsDtor(b bool)            { f.SwitchFnFlags(AccDtor, b) }
func (f *functionHeader) SetIsEarlyBinding(b bool)    { f.SwitchFnFlags(AccEarlyBinding, b) }
func (f *functionHeader) SetIsHasFinallyBlock(b bool) { f.SwitchFnFlags(AccHasFinallyBlock, b) }
func (f *functionHeader) SetIsTopLevel(b bool)        { f.SwitchFnFlags(AccTopLevel, b) }
func (f *functionHeader) SetIsTraitClone(b bool)      { f.SwitchFnFlags(AccTraitClone, b) }

func (f *functionHeader) CheckArgSendType(argNum1 uint32, mask uint8) bool {
	argNum := int(argNum1 - 1) // 传入参数从 1 开始计数
	if argNum >= len(f.argInfos) {
		if !f.IsVariadic() {
			return false
		}
		argNum = len(f.argInfos) - 1
	}
	return f.argInfos[argNum].ByReference()&mask != 0
}

func CopyFunction(function IFunction) IFunction {
	switch f := function.(type) {
	case *InternalFunction:
		var tmp InternalFunction = *f
		return &tmp
	case *UserFunction:
		var tmp UserFunction = *f
		return &tmp
	default:
		return nil
	}
}
