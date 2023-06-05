package core

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

/**
 * PhpOutputBuffer
 */
type PhpOutputBuffer struct {
	data *byte
	size int
	used int
	free bool
}

func NewOutputBuffer(data *byte, size int, used int, free bool) *PhpOutputBuffer {
	return &PhpOutputBuffer{data: data, size: size, used: used, free: free}
}
func EmptyOutputBuffer() *PhpOutputBuffer {
	return NewOutputBuffer(nil, 0, 0, false)
}

func (this *PhpOutputBuffer) GetData() *byte      { return this.data }
func (this *PhpOutputBuffer) SetData(value *byte) { this.data = value }
func (this *PhpOutputBuffer) GetSize() int        { return this.size }
func (this *PhpOutputBuffer) SetSize(value int)   { this.size = value }
func (this *PhpOutputBuffer) GetUsed() int        { return this.used }
func (this *PhpOutputBuffer) SetUsed(value int)   { this.used = value }
func (this *PhpOutputBuffer) IsFree() bool        { return this.free }
func (this *PhpOutputBuffer) SetFree(value bool)  { this.free = value }

func (this *PhpOutputBuffer) SetFreeData(handler *PhpOutputBuffer) {
	this.data = handler.data
	this.used = handler.used
	this.free = true
}
func (this *PhpOutputBuffer) SetFreeDataByStr(data string) {
	this.data = b.CastStrPtr(data)
	this.used = len(data)
	this.free = true
}

/**
 * PhpOutputContext
 */
type PhpOutputContext struct {
	op  int
	in  *PhpOutputBuffer
	out *PhpOutputBuffer
}

func (this *PhpOutputContext) Init(op int) {
	this.op = op
	this.in = nil
	this.out = nil
}

func (this *PhpOutputContext) Reset() {
	op := this.op
	this.Init(op)
}

func (this *PhpOutputContext) GetOp() int              { return this.op }
func (this *PhpOutputContext) SetOp(value int)         { this.op = value }
func (this *PhpOutputContext) GetIn() *PhpOutputBuffer { return this.in }
func (this *PhpOutputContext) SetIn(buffer *PhpOutputBuffer) {
	this.in = buffer
}
func (this *PhpOutputContext) GetOut() *PhpOutputBuffer { return this.out }
func (this *PhpOutputContext) SetOut(buffer *PhpOutputBuffer) {
	this.in = buffer
}

/**
 * PhpOutputHandlerUserFuncT
 */
type PhpOutputHandlerUserFuncT struct {
	fci types.ZendFcallInfo
	fcc types.ZendFcallInfoCache
	zoh types.Zval
}

func (this *PhpOutputHandlerUserFuncT) GetFci() types.ZendFcallInfo      { return this.fci }
func (this *PhpOutputHandlerUserFuncT) GetFcc() types.ZendFcallInfoCache { return this.fcc }
func (this *PhpOutputHandlerUserFuncT) GetZoh() types.Zval               { return this.zoh }

/**
 * PhpOutputHandler
 */
type PhpOutputHandler struct {
	name   string
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

func NewPhpOutputHandler(name string, flags int, chunkSize int) *PhpOutputHandler {
	handler := &PhpOutputHandler{
		name:  name,
		size:  chunkSize,
		flags: flags,
	}
	handler.GetBuffer().SetSize(PHP_OUTPUT_HANDLER_INITBUF_SIZE(chunkSize))
	handler.GetBuffer().SetData(zend.Emalloc(handler.GetBuffer().GetSize()))
	return handler
}

func (this *PhpOutputHandler) GetName() string                           { return this.name }
func (this *PhpOutputHandler) GetFlags() int                             { return this.flags }
func (this *PhpOutputHandler) GetLevel() int                             { return this.level }
func (this *PhpOutputHandler) GetSize() int                              { return this.size }
func (this *PhpOutputHandler) GetBuffer() *PhpOutputBuffer               { return &this.buffer }
func (this *PhpOutputHandler) GetOpaq() any                              { return this.opaq }
func (this *PhpOutputHandler) GetUser() *PhpOutputHandlerUserFuncT       { return this.func_.user }
func (this *PhpOutputHandler) GetInternal() PhpOutputHandlerContextFuncT { return this.func_.internal }

func (this *PhpOutputHandler) SetLevel(value int)                       { this.level = value }
func (this *PhpOutputHandler) SetOpaq(value any)                        { this.opaq = value }
func (this *PhpOutputHandler) SetUser(value *PhpOutputHandlerUserFuncT) { this.func_.user = value }
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
	handlers            zend.ZendStack
	active              *PhpOutputHandler
	running             *PhpOutputHandler
	outputStartFilename string
	outputStartLineno   int
	activated           bool
	flags               uint8
}

const (
	/* output global flags */
	outputImplicitFlush = 0x1
	OutputDisabled      = 0x2
	outputWritten       = 0x4
	outputSent          = 0x8
)

// fields
func (this *ZendOutputGlobals) Handlers() zend.ZendStack            { return this.handlers }
func (this *ZendOutputGlobals) SetHandlers(handlers zend.ZendStack) { this.handlers = handlers }

func (this *ZendOutputGlobals) Active() *PhpOutputHandler          { return this.active }
func (this *ZendOutputGlobals) SetActive(active *PhpOutputHandler) { this.active = active }

func (this *ZendOutputGlobals) Running() *PhpOutputHandler           { return this.running }
func (this *ZendOutputGlobals) SetRunning(running *PhpOutputHandler) { this.running = running }

func (this *ZendOutputGlobals) OutputStartFilename() string { return this.outputStartFilename }
func (this *ZendOutputGlobals) SetOutputStartFilename(outputStartFilename string) {
	this.outputStartFilename = outputStartFilename
}

func (this *ZendOutputGlobals) OutputStartLineno() int { return this.outputStartLineno }
func (this *ZendOutputGlobals) SetOutputStartLineno(outputStartLineno int) {
	this.outputStartLineno = outputStartLineno
}

// activated
func (this *ZendOutputGlobals) IsActivated() bool   { return this.activated }
func (this *ZendOutputGlobals) SetActivated(v bool) { this.activated = v }

// flags
func (this *ZendOutputGlobals) IsImplicitFlush() bool { return this.flags&outputImplicitFlush != 0 }
func (this *ZendOutputGlobals) MarkImplicitFlush(v bool) {
	if v {
		this.flags |= outputImplicitFlush
	} else {
		this.flags &^= outputImplicitFlush
	}
}

func (this *ZendOutputGlobals) IsDisabled() bool   { return this.flags&OutputDisabled != 0 }
func (this *ZendOutputGlobals) MarkDisabled()      { this.flags |= OutputDisabled }
func (this *ZendOutputGlobals) SetStatusDisabled() { this.flags = OutputDisabled }

func (this *ZendOutputGlobals) IsWritten() bool { return this.flags&outputWritten != 0 }
func (this *ZendOutputGlobals) MarkWritten()    { this.flags |= outputWritten }

func (this *ZendOutputGlobals) IsSend() bool { return this.flags&outputSent != 0 }
func (this *ZendOutputGlobals) MarkSent()    { this.flags |= outputSent }
