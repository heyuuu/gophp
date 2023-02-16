// <<generate>>

package zend

import b "sik/builtin"

/**
 * ZendFunctionEntry
 */
type ZendFunctionEntry struct {
	fname    *byte
	handler  ZifHandler
	arg_info *ArgInfo
	num_args uint32
	flags    uint32
}

func (this *ZendFunctionEntry) FuncName() string       { return b.CastStrAuto(this.fname) }
func (this *ZendFunctionEntry) Handler() ZifHandler    { return this.handler }
func (this *ZendFunctionEntry) ArgInfos() []ArgInfo    { return b.CastSlice(this.arg_info) }
func (this *ZendFunctionEntry) ReturnArgInfo() ArgInfo { return this.arg_info[-1] }
func (this *ZendFunctionEntry) Flags() uint32          { return this.flags }

func MakeZendFunctionEntryEx(fname string, flags uint32, handler ZifHandler, arg_info []ArgInfo) ZendFunctionEntry {
	return ZendFunctionEntry{
		fname:    b.CastStrPtr(fname),
		handler:  handler,
		arg_info: arg_info,
		num_args: uint32(len(arg_info)),
		flags:    flags,
	}
}

func MakeZendFunctionEntry(fname string, handler ZifHandler, arg_info *ArgInfo, num_args uint32, flags uint32) ZendFunctionEntry {
	return ZendFunctionEntry{
		fname:    b.CastStrPtr(fname),
		handler:  handler,
		arg_info: arg_info,
		num_args: num_args,
		flags:    flags,
	}
}
func (this *ZendFunctionEntry) GetFname() *byte        { return this.fname }
func (this *ZendFunctionEntry) GetHandler() ZifHandler { return this.handler }
func (this *ZendFunctionEntry) GetArgInfo() *ArgInfo   { return this.arg_info }
func (this *ZendFunctionEntry) GetNumArgs() uint32     { return this.num_args }
func (this *ZendFunctionEntry) GetFlags() uint32       { return this.flags }

/* ZendFunctionEntry.flags */
func (this *ZendFunctionEntry) AddFlags(value uint32)      { this.flags |= value }
func (this *ZendFunctionEntry) SubFlags(value uint32)      { this.flags &^= value }
func (this *ZendFunctionEntry) HasFlags(value uint32) bool { return this.flags&value != 0 }
func (this *ZendFunctionEntry) SwitchFlags(value uint32, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}
func (this ZendFunctionEntry) IsPppMask() bool          { return this.HasFlags(ZEND_ACC_PPP_MASK) }
func (this ZendFunctionEntry) IsAbstract() bool         { return this.HasFlags(ZEND_ACC_ABSTRACT) }
func (this ZendFunctionEntry) IsStatic() bool           { return this.HasFlags(ZEND_ACC_STATIC) }
func (this *ZendFunctionEntry) SetIsPppMask(cond bool)  { this.SwitchFlags(ZEND_ACC_PPP_MASK, cond) }
func (this *ZendFunctionEntry) SetIsAbstract(cond bool) { this.SwitchFlags(ZEND_ACC_ABSTRACT, cond) }
func (this *ZendFunctionEntry) SetIsStatic(cond bool)   { this.SwitchFlags(ZEND_ACC_STATIC, cond) }

/**
 * ZendFcallInfo
 */
type ZendFcallInfo struct {
	size          int
	function_name Zval
	retval        *Zval
	params        *Zval
	object        *ZendObject
	no_separation ZendBool
	param_count   uint32
}

func MakeZendFcallInfo(
	size int,
	function_name Zval,
	retval *Zval,
	params *Zval,
	object *ZendObject,
	no_separation ZendBool,
	param_count uint32,
) ZendFcallInfo {
	return ZendFcallInfo{
		size:          size,
		function_name: function_name,
		retval:        retval,
		params:        params,
		object:        object,
		no_separation: no_separation,
		param_count:   param_count,
	}
}
func (this *ZendFcallInfo) GetSize() int          { return this.size }
func (this *ZendFcallInfo) SetSize(value int)     { this.size = value }
func (this *ZendFcallInfo) GetFunctionName() Zval { return this.function_name }

// func (this *ZendFcallInfo) SetFunctionName(value Zval) { this.function_name = value }
func (this *ZendFcallInfo) GetRetval() *Zval               { return this.retval }
func (this *ZendFcallInfo) SetRetval(value *Zval)          { this.retval = value }
func (this *ZendFcallInfo) GetParams() *Zval               { return this.params }
func (this *ZendFcallInfo) SetParams(value *Zval)          { this.params = value }
func (this *ZendFcallInfo) GetObject() *ZendObject         { return this.object }
func (this *ZendFcallInfo) SetObject(value *ZendObject)    { this.object = value }
func (this *ZendFcallInfo) GetNoSeparation() ZendBool      { return this.no_separation }
func (this *ZendFcallInfo) SetNoSeparation(value ZendBool) { this.no_separation = value }
func (this *ZendFcallInfo) GetParamCount() uint32          { return this.param_count }
func (this *ZendFcallInfo) SetParamCount(value uint32)     { this.param_count = value }

/**
 * ZendFcallInfoCache
 */
type ZendFcallInfoCache struct {
	function_handler *ZendFunction
	calling_scope    *ZendClassEntry
	called_scope     *ZendClassEntry
	object           *ZendObject
}

func MakeZendFcallInfoCache(function_handler *ZendFunction, calling_scope *ZendClassEntry, called_scope *ZendClassEntry, object *ZendObject) ZendFcallInfoCache {
	return ZendFcallInfoCache{
		function_handler: function_handler,
		calling_scope:    calling_scope,
		called_scope:     called_scope,
		object:           object,
	}
}
func (this *ZendFcallInfoCache) GetFunctionHandler() *ZendFunction { return this.function_handler }
func (this *ZendFcallInfoCache) SetFunctionHandler(value *ZendFunction) {
	this.function_handler = value
}
func (this *ZendFcallInfoCache) GetCallingScope() *ZendClassEntry      { return this.calling_scope }
func (this *ZendFcallInfoCache) SetCallingScope(value *ZendClassEntry) { this.calling_scope = value }
func (this *ZendFcallInfoCache) GetCalledScope() *ZendClassEntry       { return this.called_scope }
func (this *ZendFcallInfoCache) SetCalledScope(value *ZendClassEntry)  { this.called_scope = value }
func (this *ZendFcallInfoCache) GetObject() *ZendObject                { return this.object }
func (this *ZendFcallInfoCache) SetObject(value *ZendObject)           { this.object = value }
