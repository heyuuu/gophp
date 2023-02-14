// <<generate>>

package zend

type functionHeader struct {
	// type(8) + argFlags(24) 的组合数据
	typeArgFlags uint32
}

func (this *functionHeader) GetType() uint8 { return uint8(this.typeArgFlags | 0xff) }
func (this *functionHeader) SetType(typ uint8) {
	this.typeArgFlags = (this.typeArgFlags & 0x00ffffff) | uint32(typ)
}

/**
 * ZendFunction
 */
type ZendFunction struct /* union */ {
	functionHeader
	type_           ZendUchar
	quick_arg_flags uint32
	common          struct {
		type_             ZendUchar
		arg_flags         [3]ZendUchar
		fn_flags          uint32
		function_name     *ZendString
		scope             *ZendClassEntry
		prototype         *ZendFunction
		num_args          uint32
		required_num_args uint32
		arg_info          *ZendArgInfo
	}
	op_array          ZendOpArray
	internal_function ZendInternalFunction
}

func (this *ZendFunction) ArgInfos() []ZendArgInfo {} // todo

func (this *ZendFunction) GetQuickArgFlags() uint32                  { return this.quick_arg_flags }
func (this *ZendFunction) GetCommonType() ZendUchar                  { return this.common.type_ }
func (this *ZendFunction) GetArgFlags() []ZendUchar                  { return this.common.arg_flags }
func (this *ZendFunction) GetFnFlags() uint32                        { return this.common.fn_flags }
func (this *ZendFunction) SetFnFlags(value uint32)                   { this.common.fn_flags = value }
func (this *ZendFunction) GetFunctionName() *ZendString              { return this.common.function_name }
func (this *ZendFunction) SetFunctionName(value *ZendString)         { this.common.function_name = value }
func (this *ZendFunction) GetScope() *ZendClassEntry                 { return this.common.scope }
func (this *ZendFunction) SetScope(value *ZendClassEntry)            { this.common.scope = value }
func (this *ZendFunction) GetPrototype() *ZendFunction               { return this.common.prototype }
func (this *ZendFunction) SetPrototype(value *ZendFunction)          { this.common.prototype = value }
func (this *ZendFunction) GetNumArgs() uint32                        { return this.common.num_args }
func (this *ZendFunction) GetRequiredNumArgs() uint32                { return this.common.required_num_args }
func (this *ZendFunction) GetArgInfo() *ZendArgInfo                  { return this.common.arg_info }
func (this *ZendFunction) SetArgInfo(value *ZendArgInfo)             { this.common.arg_info = value }
func (this *ZendFunction) GetOpArray() ZendOpArray                   { return this.op_array }
func (this *ZendFunction) GetInternalFunction() ZendInternalFunction { return this.internal_function }

/* ZendFunction.quick_arg_flags */
func (this *ZendFunction) AddQuickArgFlags(value uint32) { this.quick_arg_flags |= value }

/* ZendFunction.common.fn_flags */
func (this *ZendFunction) AddFnFlags(value uint32)      { this.common.fn_flags |= value }
func (this *ZendFunction) SubFnFlags(value uint32)      { this.common.fn_flags &^= value }
func (this *ZendFunction) HasFnFlags(value uint32) bool { return this.common.fn_flags&value != 0 }
func (this *ZendFunction) SwitchFnFlags(value uint32, cond bool) {
	if cond {
		this.AddFnFlags(value)
	} else {
		this.SubFnFlags(value)
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
