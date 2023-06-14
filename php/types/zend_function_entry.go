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
func (this FunctionEntry) IsPppMask() bool             { return this.HasFlags(AccPppMask) }
func (this FunctionEntry) IsAbstract() bool            { return this.HasFlags(AccAbstract) }
func (this FunctionEntry) IsStatic() bool              { return this.HasFlags(AccStatic) }
