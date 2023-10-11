package core

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/kits/slicekit"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
)

var OutputGlobals ZendOutputGlobals
var PhpOutputDirect = PhpOutputStderr

const PhpOutputDefaultHandlerName = "default output handler"

/**
 * functions
 */

func OG__() *ZendOutputGlobals { return &OutputGlobals }
func PHP_OUTPUT_HANDLER_INITBUF_SIZE(s int) int {
	if s > 1 {
		return s + PHP_OUTPUT_HANDLER_ALIGNTO_SIZE - s%PHP_OUTPUT_HANDLER_ALIGNTO_SIZE
	} else {
		return PHP_OUTPUT_HANDLER_DEFAULT_SIZE
	}
}
func PUTS(str string) int { return PhpOutputWrite(str) }
func PUTS_H(str string)   { PhpOutputWriteUnbuffered(str) }
func PUTC(c byte)         { PhpOutputWrite(string(c)) }

/**
 * types
 */

// PhpOutputHandler
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

func NewOutputHandlerUser(outputHandler *types.Zval, chunkSize int, flags int) *PhpOutputHandler {
	if outputHandler == nil || outputHandler.IsNull() {
		return NewOutputHandlerInternal(PhpOutputDefaultHandlerName, PhpOutputHandlerDefaultFunc, chunkSize, flags)
	}

	var handlerName *types.String = nil
	var err *byte = nil
	var handler *PhpOutputHandler = nil

	user := &PhpOutputHandlerUserFuncT{}
	if types.SUCCESS == zend.ZendFcallInfoInit(outputHandler, 0, user.GetFci(), user.GetFcc(), &handlerName, &error) {
		handler = newPhpOutputHandler(handlerName.GetStr(), chunkSize, flags & ^0xf | PHP_OUTPUT_HANDLER_USER)
		types.ZVAL_COPY(user.GetZoh(), outputHandler)
		handler.SetUser(user)
	}
	if err != nil {
		PhpErrorDocref("ref.outcontrol", faults.E_WARNING, "%s", err)
	}
	return handler
}

func NewOutputHandlerInternal(name string, outputHandler PhpOutputHandlerContextFuncT, chunkSize int, flags int) *PhpOutputHandler {
	handler := newPhpOutputHandler(name, chunkSize, flags & ^0xf | PHP_OUTPUT_HANDLER_INTERNAL)
	handler.func_.internal = outputHandler
	return handler
}

func newPhpOutputHandler(name string, chunkSize int, flags int) *PhpOutputHandler {
	handler := &PhpOutputHandler{
		name:  name,
		size:  chunkSize,
		flags: flags,
	}
	handler.GetBuffer().SetSize(PHP_OUTPUT_HANDLER_INITBUF_SIZE(chunkSize))
	handler.GetBuffer().SetData(zend.Emalloc(handler.GetBuffer().GetSize()))
	return handler
}

func (h *PhpOutputHandler) GetName() string                           { return h.name }
func (h *PhpOutputHandler) GetFlags() int                             { return h.flags }
func (h *PhpOutputHandler) GetLevel() int                             { return h.level }
func (h *PhpOutputHandler) GetSize() int                              { return h.size }
func (h *PhpOutputHandler) GetBuffer() *PhpOutputBuffer               { return &h.buffer }
func (h *PhpOutputHandler) GetOpaq() any                              { return h.opaq }
func (h *PhpOutputHandler) GetUser() *PhpOutputHandlerUserFuncT       { return h.func_.user }
func (h *PhpOutputHandler) GetInternal() PhpOutputHandlerContextFuncT { return h.func_.internal }

func (h *PhpOutputHandler) SetLevel(value int)                       { h.level = value }
func (h *PhpOutputHandler) SetOpaq(value any)                        { h.opaq = value }
func (h *PhpOutputHandler) SetUser(value *PhpOutputHandlerUserFuncT) { h.func_.user = value }
func (h *PhpOutputHandler) SetInternal(value PhpOutputHandlerContextFuncT) {
	h.func_.internal = value
}

/* PhpOutputHandler.flags */
func (h *PhpOutputHandler) AddFlags(value int)      { h.flags |= value }
func (h *PhpOutputHandler) SubFlags(value int)      { h.flags &^= value }
func (h *PhpOutputHandler) HasFlags(value int) bool { return h.flags&value != 0 }
func (h *PhpOutputHandler) SwitchFlags(value int, cond bool) {
	if cond {
		h.AddFlags(value)
	} else {
		h.SubFlags(value)
	}
}
func (h PhpOutputHandler) IsUser() bool         { return h.HasFlags(PHP_OUTPUT_HANDLER_USER) }
func (h PhpOutputHandler) IsStarted() bool      { return h.HasFlags(PHP_OUTPUT_HANDLER_STARTED) }
func (h PhpOutputHandler) IsDisabled() bool     { return h.HasFlags(PHP_OUTPUT_HANDLER_DISABLED) }
func (h PhpOutputHandler) IsProcessed() bool    { return h.HasFlags(PHP_OUTPUT_HANDLER_PROCESSED) }
func (h *PhpOutputHandler) SetIsUser(cond bool) { h.SwitchFlags(PHP_OUTPUT_HANDLER_USER, cond) }
func (h *PhpOutputHandler) SetIsStarted(cond bool) {
	h.SwitchFlags(PHP_OUTPUT_HANDLER_STARTED, cond)
}
func (h *PhpOutputHandler) SetIsDisabled(cond bool) {
	h.SwitchFlags(PHP_OUTPUT_HANDLER_DISABLED, cond)
}
func (h *PhpOutputHandler) SetIsProcessed(cond bool) {
	h.SwitchFlags(PHP_OUTPUT_HANDLER_PROCESSED, cond)
}

func (h *PhpOutputHandler) IsCleanable() bool { return h.HasFlags(PHP_OUTPUT_HANDLER_CLEANABLE) }
func (h *PhpOutputHandler) IsFlushable() bool { return h.HasFlags(PHP_OUTPUT_HANDLER_FLUSHABLE) }
func (h *PhpOutputHandler) IsRemovable() bool { return h.HasFlags(PHP_OUTPUT_HANDLER_REMOVABLE) }

// PhpOutputBuffer
type PhpOutputBuffer struct {
	data *byte
	size int
	used int
	free bool
}

func (buf *PhpOutputBuffer) GetData() *byte      { return buf.data }
func (buf *PhpOutputBuffer) SetData(value *byte) { buf.data = value }
func (buf *PhpOutputBuffer) GetSize() int        { return buf.size }
func (buf *PhpOutputBuffer) SetSize(value int)   { buf.size = value }
func (buf *PhpOutputBuffer) GetUsed() int        { return buf.used }
func (buf *PhpOutputBuffer) SetUsed(value int)   { buf.used = value }
func (buf *PhpOutputBuffer) IsFree() bool        { return buf.free }
func (buf *PhpOutputBuffer) SetFree(value bool)  { buf.free = value }

func (buf *PhpOutputBuffer) GetString() string { return b.CastStr(buf.data, buf.used) }

func (buf *PhpOutputBuffer) SetFreeData(handler *PhpOutputBuffer) {
	buf.data = handler.data
	buf.used = handler.used
	buf.free = true
}
func (buf *PhpOutputBuffer) SetFreeDataByStr(data string) {
	buf.data = b.CastStrPtr(data)
	buf.used = len(data)
	buf.free = true
}

// PhpOutputContext
type PhpOutputContext struct {
	op  int
	in  PhpOutputBuffer
	out PhpOutputBuffer
}

func InitOutputContext(op int) *PhpOutputContext {
	return &PhpOutputContext{op: op}
}

func (c *PhpOutputContext) Reset() {
	c.in = PhpOutputBuffer{}
	c.out = PhpOutputBuffer{}
}

func (c *PhpOutputContext) GetOutData() (string, bool) {
	if c.out.data == nil || c.out.used == 0 {
		return "", false
	}

	return b.CastStr(c.out.data, c.out.used), true
}

func (c *PhpOutputContext) Feed(buffer PhpOutputBuffer) {
	if c.in.free && c.in.data != nil {
		zend.Efree(c.in.data)
	}
	c.in = buffer
}

func (c *PhpOutputContext) Swap() {
	if c.in.free && c.in.data != nil {
		zend.Efree(c.in.data)
	}
	c.in = c.out
	c.out = PhpOutputBuffer{}
}

func (c *PhpOutputContext) Pass() {
	c.out = c.in
	c.in = PhpOutputBuffer{}
}

func (c *PhpOutputContext) GetOp() int              { return c.op }
func (c *PhpOutputContext) SetOp(value int)         { c.op = value }
func (c *PhpOutputContext) GetIn() *PhpOutputBuffer { return &c.in }
func (c *PhpOutputContext) SetIn(buffer PhpOutputBuffer) {
	c.in = buffer
}
func (c *PhpOutputContext) GetOut() *PhpOutputBuffer { return &c.out }
func (c *PhpOutputContext) SetOut(buffer PhpOutputBuffer) {
	c.in = buffer
}

// PhpOutputHandlerUserFuncT
type PhpOutputHandlerUserFuncT struct {
	fci types.ZendFcallInfo
	fcc types.ZendFcallInfoCache
	zoh types.Zval
}

func (this *PhpOutputHandlerUserFuncT) GetFci() types.ZendFcallInfo      { return this.fci }
func (this *PhpOutputHandlerUserFuncT) GetFcc() types.ZendFcallInfoCache { return this.fcc }
func (this *PhpOutputHandlerUserFuncT) GetZoh() types.Zval               { return this.zoh }

// ZendOutputGlobals
type ZendOutputGlobals struct {
	handlers            []*PhpOutputHandler
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

// life cycle
func (g *ZendOutputGlobals) StartUp() {
	g.reset()
	PhpOutputDirect = PhpOutputStdout
}
func (g *ZendOutputGlobals) Shutdown() {
	PhpOutputDirect = PhpOutputStderr
}
func (g *ZendOutputGlobals) Activate() {
	g.reset()
	g.activated = true
}

func (g *ZendOutputGlobals) Deactivate() {
	if g.IsActivated() {
		PhpOutputHeader()

		g.activated = false
		g.active = nil
		g.running = nil
		g.handlers = nil
	}
}

func (g *ZendOutputGlobals) Teardown() {
	PhpOutputEndAll()
	g.Deactivate()
	g.Shutdown()
}

func (g *ZendOutputGlobals) reset() {
	*g = ZendOutputGlobals{}
}

// handlers
func (g *ZendOutputGlobals) CountHandlers() int { return len(g.handlers) }
func (g *ZendOutputGlobals) PushHandler(h *PhpOutputHandler) int {
	g.handlers = append(g.handlers, h)
	return len(g.handlers)
}
func (g *ZendOutputGlobals) TopHandler() *PhpOutputHandler {
	if len(g.handlers) == 0 {
		return nil
	}

	return g.handlers[len(g.handlers)-1]
}
func (g *ZendOutputGlobals) PopHandler() *PhpOutputHandler {
	if len(g.handlers) == 0 {
		return nil
	}

	h := g.handlers[len(g.handlers)-1]
	g.handlers = g.handlers[:len(g.handlers)-1]
	return h
}
func (g *ZendOutputGlobals) EachHandler(bottomUp bool, handler func(h *PhpOutputHandler)) {
	if bottomUp {
		slicekit.Each(g.handlers, handler)
	} else {
		slicekit.EachReserve(g.handlers, handler)
	}
}
func (g *ZendOutputGlobals) EachHandlerEx(bottomUp bool, handler func(h *PhpOutputHandler) bool) {
	if bottomUp {
		slicekit.EachEx(g.handlers, handler)
	} else {
		slicekit.EachReserveEx(g.handlers, handler)
	}
}
func (g *ZendOutputGlobals) StartHandler(h *PhpOutputHandler) bool {
	if h == nil {
		return false
	}

	h.level = len(g.handlers)
	g.handlers = append(g.handlers, h)
	g.active = h
	return true
}

func (g *ZendOutputGlobals) GetLevel() int {
	if g.active != nil {
		return g.CountHandlers()
	}
	return 0
}
func (g *ZendOutputGlobals) GetContents() (string, bool) {
	if g.active != nil {
		return g.active.buffer.GetString(), true
	}
	return "", false
}
func (g *ZendOutputGlobals) GetLength() (int, bool) {
	if g.active != nil {
		return g.active.buffer.GetUsed(), true
	}
	return 0, false
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
