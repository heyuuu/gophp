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

func (this *PhpOutputHandler) IsCleanable() bool { return this.HasFlags(PHP_OUTPUT_HANDLER_CLEANABLE) }
func (this *PhpOutputHandler) IsFlushable() bool { return this.HasFlags(PHP_OUTPUT_HANDLER_FLUSHABLE) }
func (this *PhpOutputHandler) IsRemovable() bool { return this.HasFlags(PHP_OUTPUT_HANDLER_REMOVABLE) }

/**
 * ZendOutputGlobals
 */
type ZendOutputGlobals struct {
	handlers            zend.ZendStack
	handlersEx          []**PhpOutputHandler
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

func (g *ZendOutputGlobals) Init() {
	*g = ZendOutputGlobals{}
}

func (g *ZendOutputGlobals) Activate() {
	g.Init()
	g.SetActivated(true)
}

func (g *ZendOutputGlobals) Deactivate() {
	if g.IsActivated() {
		g.SetActivated(false)
		g.active = nil
		g.running = nil
		g.handlersEx = nil
	}
}

// handlers
func (g *ZendOutputGlobals) Handlers() *zend.ZendStack { return g.handlers }
func (g *ZendOutputGlobals) CountHandlers() int        { return len(g.handlersEx) }
func (g *ZendOutputGlobals) PushHandler(h **PhpOutputHandler) int {
	g.handlersEx = append(g.handlersEx, h)
	return len(g.handlersEx)
}
func (g *ZendOutputGlobals) PopHandler() **PhpOutputHandler {
	var c **PhpOutputHandler
	if len(g.handlersEx) > 0 {
		c = g.handlersEx[len(g.handlersEx)-1]
		g.handlersEx = g.handlersEx[:len(g.handlersEx)-1]
	}
	return c
}

// fields
func (g *ZendOutputGlobals) Active() *PhpOutputHandler          { return g.active }
func (g *ZendOutputGlobals) SetActive(active *PhpOutputHandler) { g.active = active }

func (g *ZendOutputGlobals) Running() *PhpOutputHandler           { return g.running }
func (g *ZendOutputGlobals) SetRunning(running *PhpOutputHandler) { g.running = running }

func (g *ZendOutputGlobals) OutputStartFilename() string { return g.outputStartFilename }
func (g *ZendOutputGlobals) SetOutputStartFilename(outputStartFilename string) {
	g.outputStartFilename = outputStartFilename
}

func (g *ZendOutputGlobals) OutputStartLineno() int { return g.outputStartLineno }
func (g *ZendOutputGlobals) SetOutputStartLineno(outputStartLineno int) {
	g.outputStartLineno = outputStartLineno
}

// activated
func (g *ZendOutputGlobals) IsActivated() bool   { return g.activated }
func (g *ZendOutputGlobals) SetActivated(v bool) { g.activated = v }

// flags
func (g *ZendOutputGlobals) IsImplicitFlush() bool { return g.flags&outputImplicitFlush != 0 }
func (g *ZendOutputGlobals) MarkImplicitFlush(v bool) {
	if v {
		g.flags |= outputImplicitFlush
	} else {
		g.flags &^= outputImplicitFlush
	}
}

func (g *ZendOutputGlobals) IsDisabled() bool   { return g.flags&OutputDisabled != 0 }
func (g *ZendOutputGlobals) MarkDisabled()      { g.flags |= OutputDisabled }
func (g *ZendOutputGlobals) SetStatusDisabled() { g.flags = OutputDisabled }

func (g *ZendOutputGlobals) IsWritten() bool { return g.flags&outputWritten != 0 }
func (g *ZendOutputGlobals) MarkWritten()    { g.flags |= outputWritten }

func (g *ZendOutputGlobals) IsSend() bool { return g.flags&outputSent != 0 }
func (g *ZendOutputGlobals) MarkSent()    { g.flags |= outputSent }
