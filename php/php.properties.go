package php

import (
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
)

// properties for Context
func (c *Context) Executor() *Executor {
	return c.executor
}

// properties for Engine
func (engine *Engine) Host() string {
	return engine.host
}
func (engine *Engine) Port() int {
	return engine.port
}
func (engine *Engine) BaseCtx() *Context {
	return engine.baseCtx
}

// properties for ExecutorGlobals
func (eg *ExecutorGlobals) SymbolTable() ISymtable {
	return eg.symbolTable
}
func (eg *ExecutorGlobals) ErrorReporting() perr.ErrorType {
	return eg.errorReporting
}
func (eg *ExecutorGlobals) SetErrorReporting(v perr.ErrorType) {
	eg.errorReporting = v
}
func (eg *ExecutorGlobals) ExitStatus() int {
	return eg.exitStatus
}
func (eg *ExecutorGlobals) SetExitStatus(v int) {
	eg.exitStatus = v
}
func (eg *ExecutorGlobals) Precision() int {
	return eg.precision
}
func (eg *ExecutorGlobals) SetPrecision(v int) {
	eg.precision = v
}
func (eg *ExecutorGlobals) ConstantTable() ConstantTable {
	return eg.constantTable
}
func (eg *ExecutorGlobals) FunctionTable() FunctionTable {
	return eg.functionTable
}
func (eg *ExecutorGlobals) ClassTable() ClassTable {
	return eg.classTable
}
func (eg *ExecutorGlobals) CurrentExecuteData() *ExecuteData {
	return eg.currentExecuteData
}
func (eg *ExecutorGlobals) SetCurrentExecuteData(v *ExecuteData) {
	eg.currentExecuteData = v
}

// properties for IniEntry
func (ini *IniEntry) ModuleNumber() int {
	return ini.moduleNumber
}
func (ini *IniEntry) Name() string {
	return ini.name
}
func (ini *IniEntry) Value() string {
	return ini.value
}
func (ini *IniEntry) HasValue() bool {
	return ini.hasValue
}
func (ini *IniEntry) Modifiable() IniModifiable {
	return ini.modifiable
}

// properties for IniEntryDef
func (d *IniEntryDef) Name() string {
	return d.name
}

// properties for IniGlobals
func (ig *IniGlobals) IniEntries() string {
	return ig.iniEntries
}
func (ig *IniGlobals) SetIniEntries(v string) {
	ig.iniEntries = v
}
func (ig *IniGlobals) IniIgnore() bool {
	return ig.iniIgnore
}
func (ig *IniGlobals) SetIniIgnore(v bool) {
	ig.iniIgnore = v
}
func (ig *IniGlobals) IniIgnoreCwd() bool {
	return ig.iniIgnoreCwd
}
func (ig *IniGlobals) SetIniIgnoreCwd(v bool) {
	ig.iniIgnoreCwd = v
}
func (ig *IniGlobals) IniPathOverride() string {
	return ig.iniPathOverride
}
func (ig *IniGlobals) SetIniPathOverride(v string) {
	ig.iniPathOverride = v
}
func (ig *IniGlobals) IniDefaultsFunc() func(*types.Array) {
	return ig.iniDefaultsFunc
}
func (ig *IniGlobals) SetIniDefaultsFunc(v func(*types.Array)) {
	ig.iniDefaultsFunc = v
}
func (ig *IniGlobals) HasPerDirConfig() bool {
	return ig.hasPerDirConfig
}
func (ig *IniGlobals) HasPerHostConfig() bool {
	return ig.hasPerHostConfig
}
func (ig *IniGlobals) ZendExtensions() []string {
	return ig.zendExtensions
}
func (ig *IniGlobals) PhpExtensions() []string {
	return ig.phpExtensions
}
func (ig *IniGlobals) IniOpenedPath() string {
	return ig.iniOpenedPath
}
func (ig *IniGlobals) SetIniOpenedPath(v string) {
	ig.iniOpenedPath = v
}
func (ig *IniGlobals) IniScannedPath() string {
	return ig.iniScannedPath
}
func (ig *IniGlobals) SetIniScannedPath(v string) {
	ig.iniScannedPath = v
}
func (ig *IniGlobals) IniScannedFiles() string {
	return ig.iniScannedFiles
}

// properties for StdInternObject
func (o *StdInternObject) Ctx() *Context {
	return o.ctx
}
func (o *StdInternObject) Obj() *types.Object {
	return o.obj
}
func (o *StdInternObject) Class() *types.Class {
	return o.class
}
