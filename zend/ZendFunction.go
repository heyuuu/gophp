package zend

import "github.com/heyuuu/gophp/zend/types"

type IFunction interface {
	GetType() uint8
}

type functionHeader struct {
	// type(8) + argFlags(24) 的组合数据
	typ             uint8
	fnFlags         uint32
	functionName    string
	scope           *types.ClassEntry
	prototype       *ZendFunction
	requiredNumArgs uint32
	argInfos        []ArgInfo
}

func (this *functionHeader) GetType() uint8             { return this.typ }
func (this *functionHeader) SetType(typ uint8)          { this.typ = typ }
func (this *functionHeader) GetFnFlags() uint32         { return this.fnFlags }
func (this *functionHeader) SetFnFlags(flags uint32)    { this.fnFlags = flags }
func (this *functionHeader) GetNumArgs() uint32         { return uint32(len(this.argInfos)) }
func (this *functionHeader) GetRequiredNumArgs() uint32 { return this.requiredNumArgs }

/**
 * ZendFunction
 */
type ZendFunction struct /* union */ {
	functionHeader
	common struct {
		function_name *types.String
		scope         *types.ClassEntry
		prototype     *ZendFunction
		arg_info      *ZendArgInfo
	}
	op_array          ZendOpArray
	internal_function ZendInternalFunction
}

func NewZendFunctionInternal(intern *ZendInternalFunction) *ZendFunction {
	// todo
	return &ZendFunction{}
}

func MakeZendFunctionInternal(intern *ZendInternalFunction) ZendFunction {
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

func (this *ZendFunction) GetCommonType() types.ZendUchar            { return this.GetType() }
func (this *ZendFunction) GetFunctionName() *types.String            { return this.common.function_name }
func (this *ZendFunction) SetFunctionName(value *types.String)       { this.common.function_name = value }
func (this *ZendFunction) GetScope() *types.ClassEntry               { return this.common.scope }
func (this *ZendFunction) SetScope(value *types.ClassEntry)          { this.common.scope = value }
func (this *ZendFunction) GetPrototype() *ZendFunction               { return this.common.prototype }
func (this *ZendFunction) SetPrototype(value *ZendFunction)          { this.common.prototype = value }
func (this *ZendFunction) GetArgInfo() *ZendArgInfo                  { return this.common.arg_info }
func (this *ZendFunction) SetArgInfo(value *ZendArgInfo)             { this.common.arg_info = value }
func (this *ZendFunction) GetOpArray() *ZendOpArray                  { return &this.op_array }
func (this *ZendFunction) GetInternalFunction() ZendInternalFunction { return this.internal_function }

/* ZendFunction.common.fn_flags */
func (this *ZendFunction) HasFnFlags(value uint32) bool { return this.GetFnFlags()&value != 0 }
func (this *ZendFunction) SwitchFnFlags(value uint32, cond bool) {
	if cond {
		this.SetFnFlags(this.GetFnFlags() | value)
	} else {
		this.SetFnFlags(this.GetFnFlags() &^ value)
	}
}
func (this ZendFunction) IsVariadic() bool      { return this.HasFnFlags(ZEND_ACC_VARIADIC) }
func (this ZendFunction) IsStatic() bool        { return this.HasFnFlags(ZEND_ACC_STATIC) }
func (this ZendFunction) IsHasReturnType() bool { return this.HasFnFlags(ZEND_ACC_HAS_RETURN_TYPE) }
func (this ZendFunction) IsCallViaTrampoline() bool {
	return this.HasFnFlags(ZEND_ACC_CALL_VIA_TRAMPOLINE)
}
func (this ZendFunction) IsPrivate() bool          { return this.HasFnFlags(ZEND_ACC_PRIVATE) }
func (this ZendFunction) IsPublic() bool           { return this.HasFnFlags(ZEND_ACC_PUBLIC) }
func (this ZendFunction) IsAbstract() bool         { return this.HasFnFlags(ZEND_ACC_ABSTRACT) }
func (this ZendFunction) IsAllowStatic() bool      { return this.HasFnFlags(ZEND_ACC_ALLOW_STATIC) }
func (this ZendFunction) IsProtected() bool        { return this.HasFnFlags(ZEND_ACC_PROTECTED) }
func (this ZendFunction) IsFakeClosure() bool      { return this.HasFnFlags(ZEND_ACC_FAKE_CLOSURE) }
func (this ZendFunction) IsUsesThis() bool         { return this.HasFnFlags(ZEND_ACC_USES_THIS) }
func (this ZendFunction) IsGenerator() bool        { return this.HasFnFlags(ZEND_ACC_GENERATOR) }
func (this ZendFunction) IsHeapRtCache() bool      { return this.HasFnFlags(ZEND_ACC_HEAP_RT_CACHE) }
func (this ZendFunction) IsUserArgInfo() bool      { return this.HasFnFlags(ZEND_ACC_USER_ARG_INFO) }
func (this ZendFunction) IsClosure() bool          { return this.HasFnFlags(ZEND_ACC_CLOSURE) }
func (this ZendFunction) IsImmutable() bool        { return this.HasFnFlags(ZEND_ACC_IMMUTABLE) }
func (this ZendFunction) IsStrictTypes() bool      { return this.HasFnFlags(ZEND_ACC_STRICT_TYPES) }
func (this ZendFunction) IsPreloaded() bool        { return this.HasFnFlags(ZEND_ACC_PRELOADED) }
func (this ZendFunction) IsDonePassTwo() bool      { return this.HasFnFlags(ZEND_ACC_DONE_PASS_TWO) }
func (this ZendFunction) IsDeprecated() bool       { return this.HasFnFlags(ZEND_ACC_DEPRECATED) }
func (this ZendFunction) IsFinal() bool            { return this.HasFnFlags(ZEND_ACC_FINAL) }
func (this ZendFunction) IsCtor() bool             { return this.HasFnFlags(ZEND_ACC_CTOR) }
func (this ZendFunction) IsReturnReference() bool  { return this.HasFnFlags(ZEND_ACC_RETURN_REFERENCE) }
func (this ZendFunction) IsArenaAllocated() bool   { return this.HasFnFlags(ZEND_ACC_ARENA_ALLOCATED) }
func (this ZendFunction) IsHasTypeHints() bool     { return this.HasFnFlags(ZEND_ACC_HAS_TYPE_HINTS) }
func (this ZendFunction) IsDtor() bool             { return this.HasFnFlags(ZEND_ACC_DTOR) }
func (this ZendFunction) IsChanged() bool          { return this.HasFnFlags(ZEND_ACC_CHANGED) }
func (this *ZendFunction) SetIsVariadic(cond bool) { this.SwitchFnFlags(ZEND_ACC_VARIADIC, cond) }
func (this *ZendFunction) SetIsStatic(cond bool)   { this.SwitchFnFlags(ZEND_ACC_STATIC, cond) }
func (this *ZendFunction) SetIsHasReturnType(cond bool) {
	this.SwitchFnFlags(ZEND_ACC_HAS_RETURN_TYPE, cond)
}
func (this *ZendFunction) SetIsCallViaTrampoline(cond bool) {
	this.SwitchFnFlags(ZEND_ACC_CALL_VIA_TRAMPOLINE, cond)
}
func (this *ZendFunction) SetIsPrivate(cond bool)  { this.SwitchFnFlags(ZEND_ACC_PRIVATE, cond) }
func (this *ZendFunction) SetIsPublic(cond bool)   { this.SwitchFnFlags(ZEND_ACC_PUBLIC, cond) }
func (this *ZendFunction) SetIsAbstract(cond bool) { this.SwitchFnFlags(ZEND_ACC_ABSTRACT, cond) }
func (this *ZendFunction) SetIsAllowStatic(cond bool) {
	this.SwitchFnFlags(ZEND_ACC_ALLOW_STATIC, cond)
}
func (this *ZendFunction) SetIsProtected(cond bool) { this.SwitchFnFlags(ZEND_ACC_PROTECTED, cond) }
func (this *ZendFunction) SetIsFakeClosure(cond bool) {
	this.SwitchFnFlags(ZEND_ACC_FAKE_CLOSURE, cond)
}
func (this *ZendFunction) SetIsUsesThis(cond bool)  { this.SwitchFnFlags(ZEND_ACC_USES_THIS, cond) }
func (this *ZendFunction) SetIsGenerator(cond bool) { this.SwitchFnFlags(ZEND_ACC_GENERATOR, cond) }
func (this *ZendFunction) SetIsHeapRtCache(cond bool) {
	this.SwitchFnFlags(ZEND_ACC_HEAP_RT_CACHE, cond)
}
func (this *ZendFunction) SetIsUserArgInfo(cond bool) {
	this.SwitchFnFlags(ZEND_ACC_USER_ARG_INFO, cond)
}
func (this *ZendFunction) SetIsClosure(cond bool)   { this.SwitchFnFlags(ZEND_ACC_CLOSURE, cond) }
func (this *ZendFunction) SetIsImmutable(cond bool) { this.SwitchFnFlags(ZEND_ACC_IMMUTABLE, cond) }
func (this *ZendFunction) SetIsStrictTypes(cond bool) {
	this.SwitchFnFlags(ZEND_ACC_STRICT_TYPES, cond)
}
func (this *ZendFunction) SetIsPreloaded(cond bool) { this.SwitchFnFlags(ZEND_ACC_PRELOADED, cond) }
func (this *ZendFunction) SetIsDonePassTwo(cond bool) {
	this.SwitchFnFlags(ZEND_ACC_DONE_PASS_TWO, cond)
}
func (this *ZendFunction) SetIsDeprecated(cond bool) { this.SwitchFnFlags(ZEND_ACC_DEPRECATED, cond) }
func (this *ZendFunction) SetIsFinal(cond bool)      { this.SwitchFnFlags(ZEND_ACC_FINAL, cond) }
func (this *ZendFunction) SetIsCtor(cond bool)       { this.SwitchFnFlags(ZEND_ACC_CTOR, cond) }
func (this *ZendFunction) SetIsReturnReference(cond bool) {
	this.SwitchFnFlags(ZEND_ACC_RETURN_REFERENCE, cond)
}
func (this *ZendFunction) SetIsArenaAllocated(cond bool) {
	this.SwitchFnFlags(ZEND_ACC_ARENA_ALLOCATED, cond)
}
func (this *ZendFunction) SetIsHasTypeHints(cond bool) {
	this.SwitchFnFlags(ZEND_ACC_HAS_TYPE_HINTS, cond)
}
func (this *ZendFunction) SetIsDtor(cond bool)    { this.SwitchFnFlags(ZEND_ACC_DTOR, cond) }
func (this *ZendFunction) SetIsChanged(cond bool) { this.SwitchFnFlags(ZEND_ACC_CHANGED, cond) }
