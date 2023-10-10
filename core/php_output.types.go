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
