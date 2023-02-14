// <<generate>>

package zend

/**
 * ZendInternalFunction
 */
type ZendInternalFunction struct {
	type_             ZendUchar
	arg_flags         []ZendUchar
	fn_flags          uint32
	function_name     *ZendString
	scope             *ZendClassEntry
	prototype         *ZendFunction
	num_args          uint32
	required_num_args uint32
	arg_info          *ArgInfo
	handler           ZifHandler
	module            *ZendModuleEntry
	reserved          []any
}

func MakeZendInternalFunction(
	type_ ZendUchar,
	arg_flags []ZendUchar,
	fn_flags uint32,
	function_name *ZendString,
	scope *ZendClassEntry,
	prototype *ZendFunction,
	num_args uint32,
	required_num_args uint32,
	arg_info *ArgInfo,
	handler ZifHandler,
	module *ZendModuleEntry,
	reserved []any,
) ZendInternalFunction {
	return ZendInternalFunction{
		type_:             type_,
		arg_flags:         arg_flags,
		fn_flags:          fn_flags,
		function_name:     function_name,
		scope:             scope,
		prototype:         prototype,
		num_args:          num_args,
		required_num_args: required_num_args,
		arg_info:          arg_info,
		handler:           handler,
		module:            module,
		reserved:          reserved,
	}
}

// func (this *ZendInternalFunction)  GetType() ZendUchar      { return this.type_ }
func (this *ZendInternalFunction) SetType(value ZendUchar) { this.type_ = value }

// func (this *ZendInternalFunction)  GetArgFlags() []ZendUchar      { return this.arg_flags }
// func (this *ZendInternalFunction) SetArgFlags(value []ZendUchar) { this.arg_flags = value }
// func (this *ZendInternalFunction)  GetFnFlags() uint32      { return this.fn_flags }
func (this *ZendInternalFunction) SetFnFlags(value uint32)           { this.fn_flags = value }
func (this *ZendInternalFunction) GetFunctionName() *ZendString      { return this.function_name }
func (this *ZendInternalFunction) SetFunctionName(value *ZendString) { this.function_name = value }

// func (this *ZendInternalFunction)  GetScope() *ZendClassEntry      { return this.scope }
func (this *ZendInternalFunction) SetScope(value *ZendClassEntry) { this.scope = value }

// func (this *ZendInternalFunction)  GetPrototype() *ZendFunction      { return this.prototype }
func (this *ZendInternalFunction) SetPrototype(value *ZendFunction) { this.prototype = value }
func (this *ZendInternalFunction) GetNumArgs() uint32               { return this.num_args }
func (this *ZendInternalFunction) SetNumArgs(value uint32)          { this.num_args = value }

// func (this *ZendInternalFunction)  GetRequiredNumArgs() uint32      { return this.required_num_args }
func (this *ZendInternalFunction) SetRequiredNumArgs(value uint32)  { this.required_num_args = value }
func (this *ZendInternalFunction) GetArgInfo() *ArgInfo             { return this.arg_info }
func (this *ZendInternalFunction) SetArgInfo(value *ArgInfo)        { this.arg_info = value }
func (this *ZendInternalFunction) GetHandler() ZifHandler           { return this.handler }
func (this *ZendInternalFunction) SetHandler(value ZifHandler)      { this.handler = value }
func (this *ZendInternalFunction) GetModule() *ZendModuleEntry      { return this.module }
func (this *ZendInternalFunction) SetModule(value *ZendModuleEntry) { this.module = value }
func (this *ZendInternalFunction) GetReserved() []any               { return this.reserved }

// func (this *ZendInternalFunction) SetReserved(value []any) { this.reserved = value }

/* ZendInternalFunction.fn_flags */
func (this *ZendInternalFunction) AddFnFlags(value uint32)      { this.fn_flags |= value }
func (this *ZendInternalFunction) SubFnFlags(value uint32)      { this.fn_flags &^= value }
func (this *ZendInternalFunction) HasFnFlags(value uint32) bool { return this.fn_flags&value != 0 }
func (this *ZendInternalFunction) SwitchFnFlags(value uint32, cond bool) {
	if cond {
		this.AddFnFlags(value)
	} else {
		this.SubFnFlags(value)
	}
}
func (this ZendInternalFunction) IsVariadic() bool { return this.HasFnFlags(ZEND_ACC_VARIADIC) }
func (this ZendInternalFunction) IsReturnReference() bool {
	return this.HasFnFlags(ZEND_ACC_RETURN_REFERENCE)
}
func (this ZendInternalFunction) IsHasReturnType() bool {
	return this.HasFnFlags(ZEND_ACC_HAS_RETURN_TYPE)
}
func (this ZendInternalFunction) IsUserArgInfo() bool { return this.HasFnFlags(ZEND_ACC_USER_ARG_INFO) }
func (this *ZendInternalFunction) SetIsVariadic(cond bool) {
	this.SwitchFnFlags(ZEND_ACC_VARIADIC, cond)
}
func (this *ZendInternalFunction) SetIsReturnReference(cond bool) {
	this.SwitchFnFlags(ZEND_ACC_RETURN_REFERENCE, cond)
}
func (this *ZendInternalFunction) SetIsHasReturnType(cond bool) {
	this.SwitchFnFlags(ZEND_ACC_HAS_RETURN_TYPE, cond)
}
func (this *ZendInternalFunction) SetIsUserArgInfo(cond bool) {
	this.SwitchFnFlags(ZEND_ACC_USER_ARG_INFO, cond)
}
