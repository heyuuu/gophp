// <<generate>>

package core

import (
	"sik/zend"
)

/**
 * PhpOutputBuffer
 */
type PhpOutputBuffer struct {
	data      *byte
	size      int
	used      int
	free      uint32
	_reserved uint32
}

func (this *PhpOutputBuffer) GetData() *byte           { return this.data }
func (this *PhpOutputBuffer) SetData(value *byte)      { this.data = value }
func (this *PhpOutputBuffer) GetSize() int             { return this.size }
func (this *PhpOutputBuffer) SetSize(value int)        { this.size = value }
func (this *PhpOutputBuffer) GetUsed() int             { return this.used }
func (this *PhpOutputBuffer) SetUsed(value int)        { this.used = value }
func (this *PhpOutputBuffer) GetFree() uint32          { return this.free }
func (this *PhpOutputBuffer) SetFree(value uint32)     { this.free = value }
func (this *PhpOutputBuffer) GetReserved() uint32      { return this._reserved }
func (this *PhpOutputBuffer) SetReserved(value uint32) { this._reserved = value }

/**
 * PhpOutputContext
 */
type PhpOutputContext struct {
	op  int
	in  PhpOutputBuffer
	out PhpOutputBuffer
}

func (this *PhpOutputContext) GetOp() int                   { return this.op }
func (this *PhpOutputContext) SetOp(value int)              { this.op = value }
func (this *PhpOutputContext) GetIn() PhpOutputBuffer       { return this.in }
func (this *PhpOutputContext) SetIn(value PhpOutputBuffer)  { this.in = value }
func (this *PhpOutputContext) GetOut() PhpOutputBuffer      { return this.out }
func (this *PhpOutputContext) SetOut(value PhpOutputBuffer) { this.out = value }

/**
 * PhpOutputHandlerUserFuncT
 */
type PhpOutputHandlerUserFuncT struct {
	fci zend.ZendFcallInfo
	fcc zend.ZendFcallInfoCache
	zoh zend.Zval
}

func (this *PhpOutputHandlerUserFuncT) GetFci() zend.ZendFcallInfo           { return this.fci }
func (this *PhpOutputHandlerUserFuncT) SetFci(value zend.ZendFcallInfo)      { this.fci = value }
func (this *PhpOutputHandlerUserFuncT) GetFcc() zend.ZendFcallInfoCache      { return this.fcc }
func (this *PhpOutputHandlerUserFuncT) SetFcc(value zend.ZendFcallInfoCache) { this.fcc = value }
func (this *PhpOutputHandlerUserFuncT) GetZoh() zend.Zval                    { return this.zoh }
func (this *PhpOutputHandlerUserFuncT) SetZoh(value zend.Zval)               { this.zoh = value }

/**
 * PhpOutputHandler
 */
type PhpOutputHandler struct {
	name   *zend.ZendString
	flags  int
	level  int
	size   int
	buffer PhpOutputBuffer
	opaq   any
	dtor   func(opaq any)
	func_  struct /* union */ {
		user     *PhpOutputHandlerUserFuncT
		internal PhpOutputHandlerContextFuncT
	}
}

func (this *PhpOutputHandler) GetName() *zend.ZendString                 { return this.name }
func (this *PhpOutputHandler) SetName(value *zend.ZendString)            { this.name = value }
func (this *PhpOutputHandler) GetFlags() int                             { return this.flags }
func (this *PhpOutputHandler) SetFlags(value int)                        { this.flags = value }
func (this *PhpOutputHandler) GetLevel() int                             { return this.level }
func (this *PhpOutputHandler) SetLevel(value int)                        { this.level = value }
func (this *PhpOutputHandler) GetSize() int                              { return this.size }
func (this *PhpOutputHandler) SetSize(value int)                         { this.size = value }
func (this *PhpOutputHandler) GetBuffer() PhpOutputBuffer                { return this.buffer }
func (this *PhpOutputHandler) SetBuffer(value PhpOutputBuffer)           { this.buffer = value }
func (this *PhpOutputHandler) GetOpaq() any                              { return this.opaq }
func (this *PhpOutputHandler) SetOpaq(value any)                         { this.opaq = value }
func (this *PhpOutputHandler) GetDtor() func(opaq any)                   { return this.dtor }
func (this *PhpOutputHandler) SetDtor(value func(opaq any))              { this.dtor = value }
func (this *PhpOutputHandler) GetUser() *PhpOutputHandlerUserFuncT       { return this.func_.user }
func (this *PhpOutputHandler) SetUser(value *PhpOutputHandlerUserFuncT)  { this.func_.user = value }
func (this *PhpOutputHandler) GetInternal() PhpOutputHandlerContextFuncT { return this.func_.internal }
func (this *PhpOutputHandler) SetInternal(value PhpOutputHandlerContextFuncT) {
	this.func_.internal = value
}

/* PhpOutputHandler.flags */
func (this *PhpOutputHandler) AddFlags(value int)      { this.flags |= value }
func (this *PhpOutputHandler) SubFlags(value int)      { this.flags &^= value }
func (this *PhpOutputHandler) HasFlags(value int) bool { return this.flags&value != 0 }
func (this *PhpOutputHandler) SwitchFlags(value int, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}
func (this PhpOutputHandler) IsUser() bool         { return this.HasFlags(PHP_OUTPUT_HANDLER_USER) }
func (this PhpOutputHandler) IsStarted() bool      { return this.HasFlags(PHP_OUTPUT_HANDLER_STARTED) }
func (this PhpOutputHandler) IsDisabled() bool     { return this.HasFlags(PHP_OUTPUT_HANDLER_DISABLED) }
func (this PhpOutputHandler) IsProcessed() bool    { return this.HasFlags(PHP_OUTPUT_HANDLER_PROCESSED) }
func (this *PhpOutputHandler) SetIsUser(cond bool) { this.SwitchFlags(PHP_OUTPUT_HANDLER_USER, cond) }
func (this *PhpOutputHandler) SetIsStarted(cond bool) {
	this.SwitchFlags(PHP_OUTPUT_HANDLER_STARTED, cond)
}
func (this *PhpOutputHandler) SetIsDisabled(cond bool) {
	this.SwitchFlags(PHP_OUTPUT_HANDLER_DISABLED, cond)
}
func (this *PhpOutputHandler) SetIsProcessed(cond bool) {
	this.SwitchFlags(PHP_OUTPUT_HANDLER_PROCESSED, cond)
}

/**
 * ZendOutputGlobals
 */
type ZendOutputGlobals struct {
	handlers              zend.ZendStack
	active                *PhpOutputHandler
	running               *PhpOutputHandler
	output_start_filename *byte
	output_start_lineno   int
	flags                 int
}

func (this *ZendOutputGlobals) GetHandlers() zend.ZendStack        { return this.handlers }
func (this *ZendOutputGlobals) SetHandlers(value zend.ZendStack)   { this.handlers = value }
func (this *ZendOutputGlobals) GetActive() *PhpOutputHandler       { return this.active }
func (this *ZendOutputGlobals) SetActive(value *PhpOutputHandler)  { this.active = value }
func (this *ZendOutputGlobals) GetRunning() *PhpOutputHandler      { return this.running }
func (this *ZendOutputGlobals) SetRunning(value *PhpOutputHandler) { this.running = value }
func (this *ZendOutputGlobals) GetOutputStartFilename() *byte      { return this.output_start_filename }
func (this *ZendOutputGlobals) SetOutputStartFilename(value *byte) {
	this.output_start_filename = value
}
func (this *ZendOutputGlobals) GetOutputStartLineno() int      { return this.output_start_lineno }
func (this *ZendOutputGlobals) SetOutputStartLineno(value int) { this.output_start_lineno = value }
func (this *ZendOutputGlobals) GetFlags() int                  { return this.flags }
func (this *ZendOutputGlobals) SetFlags(value int)             { this.flags = value }

/* ZendOutputGlobals.flags */
func (this *ZendOutputGlobals) AddFlags(value int)      { this.flags |= value }
func (this *ZendOutputGlobals) SubFlags(value int)      { this.flags &^= value }
func (this *ZendOutputGlobals) HasFlags(value int) bool { return this.flags&value != 0 }
func (this *ZendOutputGlobals) SwitchFlags(value int, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}
