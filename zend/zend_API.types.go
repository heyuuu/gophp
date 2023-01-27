// <<generate>>

package zend

/**
 * ZendFunctionEntry
 */
type ZendFunctionEntry struct {
	fname    *byte
	handler  ZifHandler
	arg_info *ZendInternalArgInfo
	num_args uint32
	flags    uint32
}

func (this *ZendFunctionEntry) GetFname() *byte                       { return this.fname }
func (this *ZendFunctionEntry) SetFname(value *byte)                  { this.fname = value }
func (this *ZendFunctionEntry) GetHandler() ZifHandler                { return this.handler }
func (this *ZendFunctionEntry) SetHandler(value ZifHandler)           { this.handler = value }
func (this *ZendFunctionEntry) GetArgInfo() *ZendInternalArgInfo      { return this.arg_info }
func (this *ZendFunctionEntry) SetArgInfo(value *ZendInternalArgInfo) { this.arg_info = value }
func (this *ZendFunctionEntry) GetNumArgs() uint32                    { return this.num_args }
func (this *ZendFunctionEntry) SetNumArgs(value uint32)               { this.num_args = value }
func (this *ZendFunctionEntry) GetFlags() uint32                      { return this.flags }
func (this *ZendFunctionEntry) SetFlags(value uint32)                 { this.flags = value }

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

func (this *ZendFcallInfo) GetSize() int                   { return this.size }
func (this *ZendFcallInfo) SetSize(value int)              { this.size = value }
func (this *ZendFcallInfo) GetFunctionName() Zval          { return this.function_name }
func (this *ZendFcallInfo) SetFunctionName(value Zval)     { this.function_name = value }
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
