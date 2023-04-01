package types

import "github.com/heyuuu/gophp/zend"

type IFunction interface {
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

func (f *functionHeader) IsPublic() bool            { return f.HasFnFlags(zend.AccPublic) }
func (f *functionHeader) IsProtected() bool         { return f.HasFnFlags(zend.AccProtected) }
func (f *functionHeader) IsPrivate() bool           { return f.HasFnFlags(zend.AccPrivate) }
func (f *functionHeader) IsChanged() bool           { return f.HasFnFlags(zend.AccChanged) }
func (f *functionHeader) IsStatic() bool            { return f.HasFnFlags(zend.AccStatic) }
func (f *functionHeader) IsAbstract() bool          { return f.HasFnFlags(zend.AccAbstract) }
func (f *functionHeader) IsImmutable() bool         { return f.HasFnFlags(zend.AccImmutable) }
func (f *functionHeader) IsHasTypeHints() bool      { return f.HasFnFlags(zend.AccHasTypeHints) }
func (f *functionHeader) IsTopLevel() bool          { return f.HasFnFlags(zend.AccTopLevel) }
func (f *functionHeader) IsPreloaded() bool         { return f.HasFnFlags(zend.AccPreloaded) }
func (f *functionHeader) IsAllowStatic() bool       { return f.HasFnFlags(zend.AccAllowStatic) }
func (f *functionHeader) IsVariadic() bool          { return f.HasFnFlags(zend.AccVariadic) }
func (f *functionHeader) IsHasReturnType() bool     { return f.HasFnFlags(zend.AccHasReturnType) }
func (f *functionHeader) IsCallViaTrampoline() bool { return f.HasFnFlags(zend.AccCallViaTrampoline) }
func (f *functionHeader) IsFakeClosure() bool       { return f.HasFnFlags(zend.AccFakeClosure) }
func (f *functionHeader) IsUsesThis() bool          { return f.HasFnFlags(zend.AccUsesThis) }
func (f *functionHeader) IsGenerator() bool         { return f.HasFnFlags(zend.AccGenerator) }
func (f *functionHeader) IsHeapRtCache() bool       { return f.HasFnFlags(zend.AccHeapRtCache) }
func (f *functionHeader) IsUserArgInfo() bool       { return f.HasFnFlags(zend.AccUserArgInfo) }
func (f *functionHeader) IsClosure() bool           { return f.HasFnFlags(zend.AccClosure) }
func (f *functionHeader) IsStrictTypes() bool       { return f.HasFnFlags(zend.AccStrictTypes) }
func (f *functionHeader) IsDonePassTwo() bool       { return f.HasFnFlags(zend.AccDonePassTwo) }
func (f *functionHeader) IsDeprecated() bool        { return f.HasFnFlags(zend.AccDeprecated) }
func (f *functionHeader) IsFinal() bool             { return f.HasFnFlags(zend.AccFinal) }
func (f *functionHeader) IsCtor() bool              { return f.HasFnFlags(zend.AccCtor) }
func (f *functionHeader) IsReturnReference() bool   { return f.HasFnFlags(zend.AccReturnReference) }
func (f *functionHeader) IsArenaAllocated() bool    { return f.HasFnFlags(zend.AccArenaAllocated) }
func (f *functionHeader) IsDtor() bool              { return f.HasFnFlags(zend.AccDtor) }
func (f *functionHeader) IsEarlyBinding() bool      { return f.HasFnFlags(zend.AccEarlyBinding) }
func (f *functionHeader) IsHasFinallyBlock() bool   { return f.HasFnFlags(zend.AccHasFinallyBlock) }
func (f *functionHeader) IsTraitClone() bool        { return f.HasFnFlags(zend.AccTraitClone) }

func (f *functionHeader) SetIsPublic(b bool)        { f.SwitchFnFlags(zend.AccPublic, b) }
func (f *functionHeader) SetIsProtected(b bool)     { f.SwitchFnFlags(zend.AccProtected, b) }
func (f *functionHeader) SetIsPrivate(b bool)       { f.SwitchFnFlags(zend.AccPrivate, b) }
func (f *functionHeader) SetIsStatic(b bool)        { f.SwitchFnFlags(zend.AccStatic, b) }
func (f *functionHeader) SetIsAbstract(b bool)      { f.SwitchFnFlags(zend.AccAbstract, b) }
func (f *functionHeader) SetIsVariadic(b bool)      { f.SwitchFnFlags(zend.AccVariadic, b) }
func (f *functionHeader) SetIsHasReturnType(b bool) { f.SwitchFnFlags(zend.AccHasReturnType, b) }
func (f *functionHeader) SetIsCallViaTrampoline(b bool) {
	f.SwitchFnFlags(zend.AccCallViaTrampoline, b)
}
func (f *functionHeader) SetIsAllowStatic(b bool)     { f.SwitchFnFlags(zend.AccAllowStatic, b) }
func (f *functionHeader) SetIsFakeClosure(b bool)     { f.SwitchFnFlags(zend.AccFakeClosure, b) }
func (f *functionHeader) SetIsUsesThis(b bool)        { f.SwitchFnFlags(zend.AccUsesThis, b) }
func (f *functionHeader) SetIsGenerator(b bool)       { f.SwitchFnFlags(zend.AccGenerator, b) }
func (f *functionHeader) SetIsHeapRtCache(b bool)     { f.SwitchFnFlags(zend.AccHeapRtCache, b) }
func (f *functionHeader) SetIsUserArgInfo(b bool)     { f.SwitchFnFlags(zend.AccUserArgInfo, b) }
func (f *functionHeader) SetIsClosure(b bool)         { f.SwitchFnFlags(zend.AccClosure, b) }
func (f *functionHeader) SetIsImmutable(b bool)       { f.SwitchFnFlags(zend.AccImmutable, b) }
func (f *functionHeader) SetIsStrictTypes(b bool)     { f.SwitchFnFlags(zend.AccStrictTypes, b) }
func (f *functionHeader) SetIsPreloaded(b bool)       { f.SwitchFnFlags(zend.AccPreloaded, b) }
func (f *functionHeader) SetIsDonePassTwo(b bool)     { f.SwitchFnFlags(zend.AccDonePassTwo, b) }
func (f *functionHeader) SetIsDeprecated(b bool)      { f.SwitchFnFlags(zend.AccDeprecated, b) }
func (f *functionHeader) SetIsFinal(b bool)           { f.SwitchFnFlags(zend.AccFinal, b) }
func (f *functionHeader) SetIsCtor(b bool)            { f.SwitchFnFlags(zend.AccCtor, b) }
func (f *functionHeader) SetIsReturnReference(b bool) { f.SwitchFnFlags(zend.AccReturnReference, b) }
func (f *functionHeader) SetIsArenaAllocated(b bool)  { f.SwitchFnFlags(zend.AccArenaAllocated, b) }
func (f *functionHeader) SetIsHasTypeHints(b bool)    { f.SwitchFnFlags(zend.AccHasTypeHints, b) }
func (f *functionHeader) SetIsDtor(b bool)            { f.SwitchFnFlags(zend.AccDtor, b) }
func (f *functionHeader) SetIsEarlyBinding(b bool)    { f.SwitchFnFlags(zend.AccEarlyBinding, b) }
func (f *functionHeader) SetIsHasFinallyBlock(b bool) { f.SwitchFnFlags(zend.AccHasFinallyBlock, b) }
func (f *functionHeader) SetIsTopLevel(b bool)        { f.SwitchFnFlags(zend.AccTopLevel, b) }
func (f *functionHeader) SetIsTraitClone(b bool)      { f.SwitchFnFlags(zend.AccTraitClone, b) }

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

/**
 * ZendFunction
 */
type ZendFunction struct /* union */ {
	functionHeader
	op_array          ZendOpArray
	internal_function InternalFunction
}

func NewZendFunctionInternal(intern *InternalFunction) IFunction  { return intern }
func MakeZendFunctionInternal(intern *InternalFunction) IFunction { return intern }
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

func (this *ZendFunction) GetOpArray() *ZendOpArray               { return &this.op_array }
func (this *ZendFunction) GetInternalFunction() *InternalFunction { return &this.internal_function }
