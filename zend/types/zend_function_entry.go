package types

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
)

/**
 * FunctionEntry
 */
type FunctionEntry struct {
	funcName        string
	handler         zend.ZifHandler
	requiredNumArgs uint32
	argInfos        []zend.ArgInfo
	returnArgInfo   *zend.ArgInfo
	flags           uint32
}

// 只可用于 def.DefFunc 使用，后续会做不兼容修改
func DefFunctionEntry(funcName string, handler zend.ZifHandler, requiredNumArgs uint32, argInfos []zend.ArgInfo, returnArgInfo *zend.ArgInfo, flags uint32) FunctionEntry {
	return FunctionEntry{funcName: funcName, handler: handler, requiredNumArgs: requiredNumArgs, argInfos: argInfos, returnArgInfo: returnArgInfo, flags: flags}
}

func MakeZendFunctionEntryEx(funcName string, flags uint32, handler zend.ZifHandler, inputArgInfos []zend.ArgInfo) FunctionEntry {
	var requiredNumArgs int
	var argInfos []zend.ArgInfo
	var returnArgInfo *zend.ArgInfo

	if len(inputArgInfos) > 0 {
		requiredNumArgs = inputArgInfos[0].RequiredNumArgs()
		if requiredNumArgs < 0 { // 为 -1 时表示所有参数都必填
			requiredNumArgs = len(inputArgInfos) - 1
		}
	}
	if len(inputArgInfos) > 1 {
		argInfos = inputArgInfos[1:]
	}

	return FunctionEntry{
		funcName:        funcName,
		handler:         handler,
		requiredNumArgs: uint32(requiredNumArgs),
		argInfos:        argInfos,
		returnArgInfo:   returnArgInfo,
		flags:           flags,
	}
}

func (this *FunctionEntry) FuncName() string             { return this.funcName }
func (this *FunctionEntry) Handler() zend.ZifHandler     { return this.handler }
func (this *FunctionEntry) ArgInfos() []zend.ArgInfo     { return this.argInfos }
func (this *FunctionEntry) ReturnArgInfo() *zend.ArgInfo { return this.returnArgInfo }
func (this *FunctionEntry) RequiredNumArgs() uint32      { return this.requiredNumArgs }
func (this *FunctionEntry) NumArgs() uint32              { return uint32(len(this.argInfos)) }
func (this *FunctionEntry) Flags() uint32                { return this.flags }

func (this *FunctionEntry) GetFname() *byte { return b.CastStrPtr(this.funcName) }
func (this *FunctionEntry) GetArgInfo() *zend.ArgInfo {
	return b.Cast[zend.ArgInfo](&this.argInfos)
}
func (this *FunctionEntry) GetNumArgs() uint32 { return this.NumArgs() }
func (this *FunctionEntry) GetFlags() uint32   { return this.flags }

/* FunctionEntry.flags */
func (this *FunctionEntry) HasFlags(value uint32) bool { return this.flags&value != 0 }
func (this FunctionEntry) IsPppMask() bool             { return this.HasFlags(zend.AccPppMask) }
func (this FunctionEntry) IsAbstract() bool            { return this.HasFlags(zend.AccAbstract) }
func (this FunctionEntry) IsStatic() bool              { return this.HasFlags(zend.AccStatic) }

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

// func (this *ZendFcallInfo) SetFunctionName(value Zval) { this.functionName = value }
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
	calling_scope    *ClassEntry
	called_scope     *ClassEntry
	object           *ZendObject
}

func MakeZendFcallInfoCache(function_handler *ZendFunction, calling_scope *ClassEntry, called_scope *ClassEntry, object *ZendObject) ZendFcallInfoCache {
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
func (this *ZendFcallInfoCache) GetCallingScope() *ClassEntry      { return this.calling_scope }
func (this *ZendFcallInfoCache) SetCallingScope(value *ClassEntry) { this.calling_scope = value }
func (this *ZendFcallInfoCache) GetCalledScope() *ClassEntry       { return this.called_scope }
func (this *ZendFcallInfoCache) SetCalledScope(value *ClassEntry)  { this.called_scope = value }
func (this *ZendFcallInfoCache) GetObject() *ZendObject            { return this.object }
func (this *ZendFcallInfoCache) SetObject(value *ZendObject)       { this.object = value }
