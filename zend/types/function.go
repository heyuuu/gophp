package types

import "github.com/heyuuu/gophp/zend"

type IFunction interface {
	GetType() uint8
}

type functionHeader struct {
	// type(8) + argFlags(24) 的组合数据
	typ             uint8
	fnFlags         uint32
	functionName    *String
	scope           *ClassEntry
	prototype       *ZendFunction
	numArgs         uint32
	requiredNumArgs uint32
	argInfos        []zend.ArgInfo
}

func (this *functionHeader) GetType() uint8                   { return this.typ }
func (this *functionHeader) SetType(typ uint8)                { this.typ = typ }
func (this *functionHeader) GetFnFlags() uint32               { return this.fnFlags }
func (this *functionHeader) SetFnFlags(flags uint32)          { this.fnFlags = flags }
func (this *functionHeader) GetFunctionName() *String         { return this.functionName }
func (this *functionHeader) SetFunctionName(value *String)    { this.functionName = value }
func (this *functionHeader) GetScope() *ClassEntry            { return this.scope }
func (this *functionHeader) SetScope(value *ClassEntry)       { this.scope = value }
func (this *functionHeader) GetPrototype() *ZendFunction      { return this.prototype }
func (this *functionHeader) SetPrototype(value *ZendFunction) { this.prototype = value }
func (this *functionHeader) GetNumArgs() uint32               { return uint32(len(this.argInfos)) }
func (this *functionHeader) SetNumArgs(value uint32)          { this.numArgs = value }
func (this *functionHeader) GetRequiredNumArgs() uint32       { return this.requiredNumArgs }
func (this *functionHeader) SetRequiredNumArgs(value uint32)  { this.requiredNumArgs = value }
func (this *functionHeader) GetArgInfo() []zend.ArgInfo       { return this.argInfos }
func (this *functionHeader) SetArgInfo(value []zend.ArgInfo)  { this.argInfos = value }

/* functionHeader.fnFlags */
func (this *functionHeader) AddFnFlags(value uint32)      { this.fnFlags |= value }
func (this *functionHeader) SubFnFlags(value uint32)      { this.fnFlags &^= value }
func (this *functionHeader) HasFnFlags(value uint32) bool { return this.fnFlags&value != 0 }
func (this *functionHeader) SwitchFnFlags(value uint32, cond bool) {
	if cond {
		this.AddFnFlags(value)
	} else {
		this.SubFnFlags(value)
	}
}
func (this *functionHeader) IsPublic() bool    { return this.HasFnFlags(zend.AccPublic) }
func (this *functionHeader) IsProtected() bool { return this.HasFnFlags(zend.AccProtected) }
func (this *functionHeader) IsPrivate() bool   { return this.HasFnFlags(zend.AccPrivate) }
func (this *functionHeader) IsChanged() bool   { return this.HasFnFlags(zend.AccChanged) }
func (this *functionHeader) IsStatic() bool    { return this.HasFnFlags(zend.AccStatic) }
func (this *functionHeader) IsAbstract() bool  { return this.HasFnFlags(zend.AccAbstract) }
func (this *functionHeader) IsImmutable() bool { return this.HasFnFlags(zend.AccImmutable) }
func (this *functionHeader) IsHasTypeHints() bool {
	return this.HasFnFlags(zend.AccHasTypeHints)
}
func (this *functionHeader) IsTopLevel() bool    { return this.HasFnFlags(zend.AccTopLevel) }
func (this *functionHeader) IsPreloaded() bool   { return this.HasFnFlags(zend.AccPreloaded) }
func (this *functionHeader) IsAllowStatic() bool { return this.HasFnFlags(zend.AccAllowStatic) }

func (this *functionHeader) IsVariadic() bool { return this.HasFnFlags(zend.AccVariadic) }
func (this *functionHeader) IsHasReturnType() bool {
	return this.HasFnFlags(zend.AccHasReturnType)
}
func (this *functionHeader) IsCallViaTrampoline() bool {
	return this.HasFnFlags(zend.AccCallViaTrampoline)
}
func (this *functionHeader) IsFakeClosure() bool { return this.HasFnFlags(zend.AccFakeClosure) }
func (this *functionHeader) IsUsesThis() bool    { return this.HasFnFlags(zend.AccUsesThis) }
func (this *functionHeader) IsGenerator() bool   { return this.HasFnFlags(zend.AccGenerator) }
func (this *functionHeader) IsHeapRtCache() bool { return this.HasFnFlags(zend.AccHeapRtCache) }
func (this *functionHeader) IsUserArgInfo() bool { return this.HasFnFlags(zend.AccUserArgInfo) }
func (this *functionHeader) IsClosure() bool     { return this.HasFnFlags(zend.AccClosure) }
func (this *functionHeader) IsStrictTypes() bool { return this.HasFnFlags(zend.AccStrictTypes) }
func (this *functionHeader) IsDonePassTwo() bool { return this.HasFnFlags(zend.AccDonePassTwo) }
func (this *functionHeader) IsDeprecated() bool  { return this.HasFnFlags(zend.AccDeprecated) }
func (this *functionHeader) IsFinal() bool       { return this.HasFnFlags(zend.AccFinal) }
func (this *functionHeader) IsCtor() bool        { return this.HasFnFlags(zend.AccCtor) }
func (this *functionHeader) IsReturnReference() bool {
	return this.HasFnFlags(zend.AccReturnReference)
}
func (this *functionHeader) IsArenaAllocated() bool {
	return this.HasFnFlags(zend.AccArenaAllocated)
}
func (this *functionHeader) IsDtor() bool { return this.HasFnFlags(zend.AccDtor) }

func (this *functionHeader) IsEarlyBinding() bool {
	return this.HasFnFlags(zend.AccEarlyBinding)
}
func (this *functionHeader) IsHasFinallyBlock() bool {
	return this.HasFnFlags(zend.AccHasFinallyBlock)
}
func (this *functionHeader) IsTraitClone() bool { return this.HasFnFlags(zend.AccTraitClone) }

func (this *functionHeader) SetIsVariadic(cond bool) {
	this.SwitchFnFlags(zend.AccVariadic, cond)
}
func (this *functionHeader) SetIsStatic(cond bool) { this.SwitchFnFlags(zend.AccStatic, cond) }
func (this *functionHeader) SetIsHasReturnType(cond bool) {
	this.SwitchFnFlags(zend.AccHasReturnType, cond)
}
func (this *functionHeader) SetIsCallViaTrampoline(cond bool) {
	this.SwitchFnFlags(zend.AccCallViaTrampoline, cond)
}
func (this *functionHeader) SetIsPrivate(cond bool) { this.SwitchFnFlags(zend.AccPrivate, cond) }
func (this *functionHeader) SetIsPublic(cond bool)  { this.SwitchFnFlags(zend.AccPublic, cond) }
func (this *functionHeader) SetIsAbstract(cond bool) {
	this.SwitchFnFlags(zend.AccAbstract, cond)
}
func (this *functionHeader) SetIsAllowStatic(cond bool) {
	this.SwitchFnFlags(zend.AccAllowStatic, cond)
}
func (this *functionHeader) SetIsProtected(cond bool) {
	this.SwitchFnFlags(zend.AccProtected, cond)
}
func (this *functionHeader) SetIsFakeClosure(cond bool) {
	this.SwitchFnFlags(zend.AccFakeClosure, cond)
}
func (this *functionHeader) SetIsUsesThis(cond bool) {
	this.SwitchFnFlags(zend.AccUsesThis, cond)
}
func (this *functionHeader) SetIsGenerator(cond bool) {
	this.SwitchFnFlags(zend.AccGenerator, cond)
}
func (this *functionHeader) SetIsHeapRtCache(cond bool) {
	this.SwitchFnFlags(zend.AccHeapRtCache, cond)
}
func (this *functionHeader) SetIsUserArgInfo(cond bool) {
	this.SwitchFnFlags(zend.AccUserArgInfo, cond)
}
func (this *functionHeader) SetIsClosure(cond bool) { this.SwitchFnFlags(zend.AccClosure, cond) }
func (this *functionHeader) SetIsImmutable(cond bool) {
	this.SwitchFnFlags(zend.AccImmutable, cond)
}
func (this *functionHeader) SetIsStrictTypes(cond bool) {
	this.SwitchFnFlags(zend.AccStrictTypes, cond)
}
func (this *functionHeader) SetIsPreloaded(cond bool) {
	this.SwitchFnFlags(zend.AccPreloaded, cond)
}
func (this *functionHeader) SetIsDonePassTwo(cond bool) {
	this.SwitchFnFlags(zend.AccDonePassTwo, cond)
}
func (this *functionHeader) SetIsDeprecated(cond bool) {
	this.SwitchFnFlags(zend.AccDeprecated, cond)
}
func (this *functionHeader) SetIsFinal(cond bool) { this.SwitchFnFlags(zend.AccFinal, cond) }
func (this *functionHeader) SetIsCtor(cond bool)  { this.SwitchFnFlags(zend.AccCtor, cond) }
func (this *functionHeader) SetIsReturnReference(cond bool) {
	this.SwitchFnFlags(zend.AccReturnReference, cond)
}
func (this *functionHeader) SetIsArenaAllocated(cond bool) {
	this.SwitchFnFlags(zend.AccArenaAllocated, cond)
}
func (this *functionHeader) SetIsHasTypeHints(cond bool) {
	this.SwitchFnFlags(zend.AccHasTypeHints, cond)
}
func (this *functionHeader) SetIsDtor(cond bool) { this.SwitchFnFlags(zend.AccDtor, cond) }
func (this *functionHeader) SetIsEarlyBinding(cond bool) {
	this.SwitchFnFlags(zend.AccEarlyBinding, cond)
}
func (this *functionHeader) SetIsHasFinallyBlock(cond bool) {
	this.SwitchFnFlags(zend.AccHasFinallyBlock, cond)
}
func (this *functionHeader) SetIsTopLevel(cond bool) {
	this.SwitchFnFlags(zend.AccTopLevel, cond)
}
func (this *functionHeader) SetIsTraitClone(cond bool) {
	this.SwitchFnFlags(zend.AccTraitClone, cond)
}

/**
 * ZendFunction
 */
type ZendFunction struct /* union */ {
	functionHeader
	op_array          ZendOpArray
	internal_function InternalFunction
}

func NewZendFunctionInternal(intern *InternalFunction) *ZendFunction {
	// todo
	return &ZendFunction{}
}

func MakeZendFunctionInternal(intern *InternalFunction) ZendFunction {
	// todo
	return ZendFunction{}
}

func (this *ZendFunction) ZEND_CHECK_ARG_FLAG(argNum1 uint32, mask uint8) bool {
	return this.CheckArgSendType(argNum1, mask)
}

func (this *ZendFunction) CheckArgSendType(argNum1 uint32, mask uint8) bool {
	argNum := int(argNum1 - 1) // 传入参数从 1 开始计数
	if argNum >= len(this.argInfos) {
		if !this.IsVariadic() {
			return false
		}
		argNum = len(this.argInfos) - 1
	}
	return this.argInfos[argNum].ByReference()&mask != 0
}

func (this *ZendFunction) GetOpArray() *ZendOpArray               { return &this.op_array }
func (this *ZendFunction) GetInternalFunction() *InternalFunction { return &this.internal_function }
