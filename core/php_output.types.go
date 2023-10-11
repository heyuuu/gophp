package core

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/kits/slicekit"
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

func (buf *PhpOutputBuffer) GetData() *byte      { return buf.data }
func (buf *PhpOutputBuffer) SetData(value *byte) { buf.data = value }
func (buf *PhpOutputBuffer) GetSize() int        { return buf.size }
func (buf *PhpOutputBuffer) SetSize(value int)   { buf.size = value }
func (buf *PhpOutputBuffer) GetUsed() int        { return buf.used }
func (buf *PhpOutputBuffer) SetUsed(value int)   { buf.used = value }
func (buf *PhpOutputBuffer) IsFree() bool        { return buf.free }
func (buf *PhpOutputBuffer) SetFree(value bool)  { buf.free = value }

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

/**
 * PhpOutputContext
 */
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
func (g *ZendOutputGlobals) TopHandler() **PhpOutputHandler {
	if len(g.handlersEx) == 0 {
		return nil
	}

	return g.handlersEx[len(g.handlersEx)-1]
}
func (g *ZendOutputGlobals) PopHandler() **PhpOutputHandler {
	if len(g.handlersEx) == 0 {
		return nil
	}

	h := g.handlersEx[len(g.handlersEx)-1]
	g.handlersEx = g.handlersEx[:len(g.handlersEx)-1]
	return h
}
func (g *ZendOutputGlobals) EachHandler(bottomUp bool, handler func(h **PhpOutputHandler)) {
	if bottomUp {
		slicekit.Each(g.handlersEx, handler)
	} else {
		slicekit.EachReserve(g.handlersEx, handler)
	}
}
func (g *ZendOutputGlobals) EachHandlerEx(bottomUp bool, handler func(h **PhpOutputHandler) bool) {
	if bottomUp {
		slicekit.EachEx(g.handlersEx, handler)
	} else {
		slicekit.EachReserveEx(g.handlersEx, handler)
	}
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
