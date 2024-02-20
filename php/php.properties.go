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

// properties for PhpGlobals
func (pg *PhpGlobals) ImplicitFlush() bool {
	return pg.implicitFlush
}
func (pg *PhpGlobals) SetImplicitFlush(v bool) {
	pg.implicitFlush = v
}
func (pg *PhpGlobals) OutputBuffering() int {
	return pg.outputBuffering
}
func (pg *PhpGlobals) SetOutputBuffering(v int) {
	pg.outputBuffering = v
}
func (pg *PhpGlobals) OutputHandler() string {
	return pg.outputHandler
}
func (pg *PhpGlobals) SetOutputHandler(v string) {
	pg.outputHandler = v
}
func (pg *PhpGlobals) UnserializeCallbackFunc() string {
	return pg.unserializeCallbackFunc
}
func (pg *PhpGlobals) SetUnserializeCallbackFunc(v string) {
	pg.unserializeCallbackFunc = v
}
func (pg *PhpGlobals) SerializePrecision() int {
	return pg.serializePrecision
}
func (pg *PhpGlobals) SetSerializePrecision(v int) {
	pg.serializePrecision = v
}
func (pg *PhpGlobals) MemoryLimit() int {
	return pg.memoryLimit
}
func (pg *PhpGlobals) SetMemoryLimit(v int) {
	pg.memoryLimit = v
}
func (pg *PhpGlobals) MaxInputTime() int {
	return pg.maxInputTime
}
func (pg *PhpGlobals) SetMaxInputTime(v int) {
	pg.maxInputTime = v
}
func (pg *PhpGlobals) TrackErrors() bool {
	return pg.trackErrors
}
func (pg *PhpGlobals) SetTrackErrors(v bool) {
	pg.trackErrors = v
}
func (pg *PhpGlobals) DisplayErrors() int {
	return pg.displayErrors
}
func (pg *PhpGlobals) SetDisplayErrors(v int) {
	pg.displayErrors = v
}
func (pg *PhpGlobals) DisplayStartupErrors() bool {
	return pg.displayStartupErrors
}
func (pg *PhpGlobals) SetDisplayStartupErrors(v bool) {
	pg.displayStartupErrors = v
}
func (pg *PhpGlobals) LogErrors() bool {
	return pg.logErrors
}
func (pg *PhpGlobals) SetLogErrors(v bool) {
	pg.logErrors = v
}
func (pg *PhpGlobals) LogErrorsMaxLen() int {
	return pg.logErrorsMaxLen
}
func (pg *PhpGlobals) SetLogErrorsMaxLen(v int) {
	pg.logErrorsMaxLen = v
}
func (pg *PhpGlobals) IgnoreRepeatedErrors() bool {
	return pg.ignoreRepeatedErrors
}
func (pg *PhpGlobals) SetIgnoreRepeatedErrors(v bool) {
	pg.ignoreRepeatedErrors = v
}
func (pg *PhpGlobals) IgnoreRepeatedSource() bool {
	return pg.ignoreRepeatedSource
}
func (pg *PhpGlobals) SetIgnoreRepeatedSource(v bool) {
	pg.ignoreRepeatedSource = v
}
func (pg *PhpGlobals) ReportMemleaks() bool {
	return pg.reportMemleaks
}
func (pg *PhpGlobals) SetReportMemleaks(v bool) {
	pg.reportMemleaks = v
}
func (pg *PhpGlobals) ErrorLog() string {
	return pg.errorLog
}
func (pg *PhpGlobals) SetErrorLog(v string) {
	pg.errorLog = v
}
func (pg *PhpGlobals) DocRoot() string {
	return pg.docRoot
}
func (pg *PhpGlobals) SetDocRoot(v string) {
	pg.docRoot = v
}
func (pg *PhpGlobals) UserDir() string {
	return pg.userDir
}
func (pg *PhpGlobals) SetUserDir(v string) {
	pg.userDir = v
}
func (pg *PhpGlobals) IncludePath() string {
	return pg.includePath
}
func (pg *PhpGlobals) SetIncludePath(v string) {
	pg.includePath = v
}
func (pg *PhpGlobals) OpenBasedir() string {
	return pg.openBasedir
}
func (pg *PhpGlobals) SetOpenBasedir(v string) {
	pg.openBasedir = v
}
func (pg *PhpGlobals) ExtensionDir() string {
	return pg.extensionDir
}
func (pg *PhpGlobals) SetExtensionDir(v string) {
	pg.extensionDir = v
}
func (pg *PhpGlobals) PhpBinary() string {
	return pg.phpBinary
}
func (pg *PhpGlobals) SetPhpBinary(v string) {
	pg.phpBinary = v
}
func (pg *PhpGlobals) SysTempDir() string {
	return pg.sysTempDir
}
func (pg *PhpGlobals) SetSysTempDir(v string) {
	pg.sysTempDir = v
}
func (pg *PhpGlobals) UploadTmpDir() string {
	return pg.uploadTmpDir
}
func (pg *PhpGlobals) SetUploadTmpDir(v string) {
	pg.uploadTmpDir = v
}
func (pg *PhpGlobals) UploadMaxFilesize() int {
	return pg.uploadMaxFilesize
}
func (pg *PhpGlobals) SetUploadMaxFilesize(v int) {
	pg.uploadMaxFilesize = v
}
func (pg *PhpGlobals) ErrorAppendString() string {
	return pg.errorAppendString
}
func (pg *PhpGlobals) SetErrorAppendString(v string) {
	pg.errorAppendString = v
}
func (pg *PhpGlobals) ErrorPrependString() string {
	return pg.errorPrependString
}
func (pg *PhpGlobals) SetErrorPrependString(v string) {
	pg.errorPrependString = v
}
func (pg *PhpGlobals) AutoPrependFile() string {
	return pg.autoPrependFile
}
func (pg *PhpGlobals) SetAutoPrependFile(v string) {
	pg.autoPrependFile = v
}
func (pg *PhpGlobals) AutoAppendFile() string {
	return pg.autoAppendFile
}
func (pg *PhpGlobals) SetAutoAppendFile(v string) {
	pg.autoAppendFile = v
}
func (pg *PhpGlobals) InputEncoding() string {
	return pg.inputEncoding
}
func (pg *PhpGlobals) SetInputEncoding(v string) {
	pg.inputEncoding = v
}
func (pg *PhpGlobals) InternalEncoding() string {
	return pg.internalEncoding
}
func (pg *PhpGlobals) SetInternalEncoding(v string) {
	pg.internalEncoding = v
}
func (pg *PhpGlobals) OutputEncoding() string {
	return pg.outputEncoding
}
func (pg *PhpGlobals) SetOutputEncoding(v string) {
	pg.outputEncoding = v
}
func (pg *PhpGlobals) VariablesOrder() string {
	return pg.variablesOrder
}
func (pg *PhpGlobals) SetVariablesOrder(v string) {
	pg.variablesOrder = v
}
func (pg *PhpGlobals) ConnectionStatus() int16 {
	return pg.connectionStatus
}
func (pg *PhpGlobals) SetConnectionStatus(v int16) {
	pg.connectionStatus = v
}
func (pg *PhpGlobals) IgnoreUserAbort() bool {
	return pg.ignoreUserAbort
}
func (pg *PhpGlobals) SetIgnoreUserAbort(v bool) {
	pg.ignoreUserAbort = v
}
func (pg *PhpGlobals) HeaderIsBeingSent() uint8 {
	return pg.headerIsBeingSent
}
func (pg *PhpGlobals) SetHeaderIsBeingSent(v uint8) {
	pg.headerIsBeingSent = v
}
func (pg *PhpGlobals) ExposePhp() bool {
	return pg.exposePhp
}
func (pg *PhpGlobals) SetExposePhp(v bool) {
	pg.exposePhp = v
}
func (pg *PhpGlobals) RegisterArgcArgv() bool {
	return pg.registerArgcArgv
}
func (pg *PhpGlobals) SetRegisterArgcArgv(v bool) {
	pg.registerArgcArgv = v
}
func (pg *PhpGlobals) AutoGlobalsJit() bool {
	return pg.autoGlobalsJit
}
func (pg *PhpGlobals) SetAutoGlobalsJit(v bool) {
	pg.autoGlobalsJit = v
}
func (pg *PhpGlobals) DocrefRoot() string {
	return pg.docrefRoot
}
func (pg *PhpGlobals) SetDocrefRoot(v string) {
	pg.docrefRoot = v
}
func (pg *PhpGlobals) DocrefExt() string {
	return pg.docrefExt
}
func (pg *PhpGlobals) SetDocrefExt(v string) {
	pg.docrefExt = v
}
func (pg *PhpGlobals) HtmlErrors() bool {
	return pg.htmlErrors
}
func (pg *PhpGlobals) SetHtmlErrors(v bool) {
	pg.htmlErrors = v
}
func (pg *PhpGlobals) XmlrpcErrors() bool {
	return pg.xmlrpcErrors
}
func (pg *PhpGlobals) SetXmlrpcErrors(v bool) {
	pg.xmlrpcErrors = v
}
func (pg *PhpGlobals) XmlrpcErrorNumber() int {
	return pg.xmlrpcErrorNumber
}
func (pg *PhpGlobals) SetXmlrpcErrorNumber(v int) {
	pg.xmlrpcErrorNumber = v
}
func (pg *PhpGlobals) ModulesActivated() bool {
	return pg.modulesActivated
}
func (pg *PhpGlobals) SetModulesActivated(v bool) {
	pg.modulesActivated = v
}
func (pg *PhpGlobals) FileUploads() bool {
	return pg.fileUploads
}
func (pg *PhpGlobals) SetFileUploads(v bool) {
	pg.fileUploads = v
}
func (pg *PhpGlobals) DuringRequestStartup() bool {
	return pg.duringRequestStartup
}
func (pg *PhpGlobals) SetDuringRequestStartup(v bool) {
	pg.duringRequestStartup = v
}
func (pg *PhpGlobals) AllowUrlFopen() bool {
	return pg.allowUrlFopen
}
func (pg *PhpGlobals) SetAllowUrlFopen(v bool) {
	pg.allowUrlFopen = v
}
func (pg *PhpGlobals) EnablePostDataReading() bool {
	return pg.enablePostDataReading
}
func (pg *PhpGlobals) SetEnablePostDataReading(v bool) {
	pg.enablePostDataReading = v
}
func (pg *PhpGlobals) ReportZendDebug() bool {
	return pg.reportZendDebug
}
func (pg *PhpGlobals) SetReportZendDebug(v bool) {
	pg.reportZendDebug = v
}
func (pg *PhpGlobals) LastError() *LastError {
	return pg.lastError
}
func (pg *PhpGlobals) DisableFunctions() string {
	return pg.disableFunctions
}
func (pg *PhpGlobals) SetDisableFunctions(v string) {
	pg.disableFunctions = v
}
func (pg *PhpGlobals) DisableClasses() string {
	return pg.disableClasses
}
func (pg *PhpGlobals) SetDisableClasses(v string) {
	pg.disableClasses = v
}
func (pg *PhpGlobals) AllowUrlInclude() bool {
	return pg.allowUrlInclude
}
func (pg *PhpGlobals) SetAllowUrlInclude(v bool) {
	pg.allowUrlInclude = v
}
func (pg *PhpGlobals) MaxInputNestingLevel() int {
	return pg.maxInputNestingLevel
}
func (pg *PhpGlobals) SetMaxInputNestingLevel(v int) {
	pg.maxInputNestingLevel = v
}
func (pg *PhpGlobals) MaxInputVars() int {
	return pg.maxInputVars
}
func (pg *PhpGlobals) SetMaxInputVars(v int) {
	pg.maxInputVars = v
}
func (pg *PhpGlobals) InUserInclude() bool {
	return pg.inUserInclude
}
func (pg *PhpGlobals) SetInUserInclude(v bool) {
	pg.inUserInclude = v
}
func (pg *PhpGlobals) UserIniFilename() string {
	return pg.userIniFilename
}
func (pg *PhpGlobals) SetUserIniFilename(v string) {
	pg.userIniFilename = v
}
func (pg *PhpGlobals) UserIniCacheTtl() int {
	return pg.userIniCacheTtl
}
func (pg *PhpGlobals) SetUserIniCacheTtl(v int) {
	pg.userIniCacheTtl = v
}
func (pg *PhpGlobals) RequestOrder() string {
	return pg.requestOrder
}
func (pg *PhpGlobals) SetRequestOrder(v string) {
	pg.requestOrder = v
}
func (pg *PhpGlobals) MailXHeader() bool {
	return pg.mailXHeader
}
func (pg *PhpGlobals) SetMailXHeader(v bool) {
	pg.mailXHeader = v
}
func (pg *PhpGlobals) MailLog() string {
	return pg.mailLog
}
func (pg *PhpGlobals) SetMailLog(v string) {
	pg.mailLog = v
}
func (pg *PhpGlobals) InErrorLog() bool {
	return pg.inErrorLog
}
func (pg *PhpGlobals) SetInErrorLog(v bool) {
	pg.inErrorLog = v
}
func (pg *PhpGlobals) SyslogFacility() int {
	return pg.syslogFacility
}
func (pg *PhpGlobals) SetSyslogFacility(v int) {
	pg.syslogFacility = v
}
func (pg *PhpGlobals) SyslogIdent() string {
	return pg.syslogIdent
}
func (pg *PhpGlobals) SetSyslogIdent(v string) {
	pg.syslogIdent = v
}
func (pg *PhpGlobals) HaveCalledOpenlog() bool {
	return pg.haveCalledOpenlog
}
func (pg *PhpGlobals) SetHaveCalledOpenlog(v bool) {
	pg.haveCalledOpenlog = v
}
func (pg *PhpGlobals) SyslogFilter() int {
	return pg.syslogFilter
}
func (pg *PhpGlobals) SetSyslogFilter(v int) {
	pg.syslogFilter = v
}
func (pg *PhpGlobals) DefaultMimetype() string {
	return pg.defaultMimetype
}
func (pg *PhpGlobals) SetDefaultMimetype(v string) {
	pg.defaultMimetype = v
}
func (pg *PhpGlobals) DefaultCharset() string {
	return pg.defaultCharset
}
func (pg *PhpGlobals) SetDefaultCharset(v string) {
	pg.defaultCharset = v
}
func (pg *PhpGlobals) PostMaxSize() int {
	return pg.postMaxSize
}
func (pg *PhpGlobals) SetPostMaxSize(v int) {
	pg.postMaxSize = v
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
