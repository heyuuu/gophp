package zend

import "github.com/heyuuu/gophp/zend/types"

/**
 * ZendInternalFunction
 */
const ZEND_MAX_RESERVED_RESOURCES = 6

var _ IFunction = (*ZendInternalFunction)(nil)

type ZendInternalFunction struct {
	type_             types.ZendUchar
	arg_flags         []types.ZendUchar
	fn_flags          uint32
	function_name     *types.String
	scope             *types.ClassEntry
	prototype         *ZendFunction
	num_args          uint32
	required_num_args uint32
	arg_info          *ArgInfo
	handler           ZifHandler
	module            *ZendModuleEntry
	reserved          [ZEND_MAX_RESERVED_RESOURCES]any
}

func NewInternalFunction() *ZendInternalFunction {
	return &ZendInternalFunction{}
}

func NewInternalFunctionEx(funcName string, handler ZifHandler) *ZendInternalFunction {
	return &ZendInternalFunction{
		function_name: types.NewString(funcName),
		handler:       handler,
	}
}

func MakeInternalFunctionSimplify(handler ZifHandler) ZendInternalFunction {
	return ZendInternalFunction{handler: handler}
}

func (this *ZendInternalFunction) InitByEntry(entry *types.ZendFunctionEntry) {
	this.handler = entry.Handler()
	this.function_name = types.NewString(entry.FuncName())
	this.prototype = nil
}

func (this *ZendInternalFunction) GetType() uint8 { return ZEND_INTERNAL_FUNCTION }

func (this *ZendInternalFunction) SetType(value types.ZendUchar) { this.type_ = value }

// func (this *ZendInternalFunction)  GetArgFlags() []ZendUchar      { return this.arg_flags }
// func (this *ZendInternalFunction) SetArgFlags(value []ZendUchar) { this.arg_flags = value }
// func (this *ZendInternalFunction)  GetFnFlags() uint32      { return this.fn_flags }
func (this *ZendInternalFunction) SetFnFlags(value uint32)        { this.fn_flags = value }
func (this *ZendInternalFunction) GetFunctionName() *types.String { return this.function_name }
func (this *ZendInternalFunction) SetFunctionName(value *types.String) {
	this.function_name = value
}

// func (this *ZendInternalFunction)  GetScope() *ClassEntry      { return this.scope }
func (this *ZendInternalFunction) SetScope(value *types.ClassEntry) { this.scope = value }

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
