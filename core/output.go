package core

import (
	"github.com/heyuuu/gophp/kits/slicekit"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/shim/slices"
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
func PUTS(str string) int      { return PhpOutputWrite(str) }
func PUTS_H(str string)        { PhpOutputWriteUnbuffered(str) }
func PUTC(c byte)              { PhpOutputWrite(string(c)) }

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
	handler.buffer.Init(chunkSize)
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
type PhpOutputBuffer struct{ data []byte }

func (buf *PhpOutputBuffer) Init(chunkSize int) {
	buf.data = make([]byte, buf.initSize(chunkSize))
}
func (buf *PhpOutputBuffer) initSize(s int) int {
	if s > 1 {
		return s + PHP_OUTPUT_HANDLER_ALIGNTO_SIZE - s%PHP_OUTPUT_HANDLER_ALIGNTO_SIZE
	}
	return PHP_OUTPUT_HANDLER_DEFAULT_SIZE
}

func (buf *PhpOutputBuffer) Append(data []byte)         { buf.data = append(buf.data, data...) }
func (buf *PhpOutputBuffer) SetData(data []byte)        { buf.data = slices.Clone(data) }
func (buf *PhpOutputBuffer) SetDataNoClone(data []byte) { buf.data = data }
func (buf *PhpOutputBuffer) SetDataStr(data string)     { buf.data = []byte(data) }
func (buf *PhpOutputBuffer) Reset()                     { buf.data = buf.data[:0] }
func (buf *PhpOutputBuffer) Clean()                     { buf.data = nil }

func (buf *PhpOutputBuffer) Bytes() []byte  { return buf.data }
func (buf *PhpOutputBuffer) String() string { return string(buf.data) }
func (buf *PhpOutputBuffer) Size() int      { return cap(buf.data) }
func (buf *PhpOutputBuffer) Used() int      { return len(buf.data) }

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

func (c *PhpOutputContext) Feed(buffer PhpOutputBuffer) {
	c.in = buffer
}

func (c *PhpOutputContext) Swap() {
	c.in = c.out
	c.out = PhpOutputBuffer{}
}

func (c *PhpOutputContext) Pass() {
	c.out = c.in
	c.in = PhpOutputBuffer{}
}

func (c *PhpOutputContext) GetOp() int               { return c.op }
func (c *PhpOutputContext) SetOp(value int)          { c.op = value }
func (c *PhpOutputContext) GetIn() *PhpOutputBuffer  { return &c.in }
func (c *PhpOutputContext) GetOut() *PhpOutputBuffer { return &c.out }

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
	activated           bool
	flags               uint8
	handlers            []*PhpOutputHandler
	active              *PhpOutputHandler
	running             *PhpOutputHandler
	outputStartFilename string
	outputStartLineno   int
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
func (g *ZendOutputGlobals) EndHandler() {
	g.PopHandler()
	g.active = g.TopHandler()
}
func (g *ZendOutputGlobals) GetLevel() int {
	if g.active != nil {
		return g.CountHandlers()
	}
	return 0
}
func (g *ZendOutputGlobals) GetContents() (string, bool) {
	if g.active != nil {
		return g.active.buffer.String(), true
	}
	return "", false
}
func (g *ZendOutputGlobals) GetLength() (int, bool) {
	if g.active != nil {
		return g.active.buffer.Used(), true
	}
	return 0, false
}

// fields
func (g *ZendOutputGlobals) IsActivated() bool                { return g.activated }
func (g *ZendOutputGlobals) Active() *PhpOutputHandler        { return g.active }
func (g *ZendOutputGlobals) Running() *PhpOutputHandler       { return g.running }
func (g *ZendOutputGlobals) StartFilename() string            { return g.outputStartFilename }
func (g *ZendOutputGlobals) SetStartFilename(filename string) { g.outputStartFilename = filename }
func (g *ZendOutputGlobals) StartLineno() int                 { return g.outputStartLineno }
func (g *ZendOutputGlobals) SetStartLineno(lineno int)        { g.outputStartLineno = lineno }

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
