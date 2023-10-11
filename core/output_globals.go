package core

import (
	"github.com/heyuuu/gophp/kits/slicekit"
	"github.com/heyuuu/gophp/zend/faults"
	"io"
	"os"
)

var OutputGlobals ZendOutputGlobals

func OG__() *ZendOutputGlobals { return &OutputGlobals }
func PUTS(str string) int      { return OG__().WriteString(str) }
func PUTS_H(str string)        { OG__().WriteStringUnbuffered(str) }
func PUTC(c byte)              { OG__().Write([]byte{c}) }

// ZendOutputGlobals
type ZendOutputGlobals struct {
	activated           bool
	flags               uint8
	handlers            []*PhpOutputHandler
	active              *PhpOutputHandler
	running             *PhpOutputHandler
	outputStartFilename string
	outputStartLineno   int
	directOutput        io.Writer
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
	g.directOutput = os.Stdout
}
func (g *ZendOutputGlobals) Shutdown() {
	g.directOutput = os.Stderr
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

// common
func (g *ZendOutputGlobals) Write(str []byte) int {
	if g.activated {
		g.operate(PHP_OUTPUT_HANDLER_WRITE, str)
		return len(str)
	}
	if g.IsDisabled() {
		return 0
	}
	g.directOutput.Write(str)
	return len(str)
}

func (g *ZendOutputGlobals) WriteString(str string) int { return g.Write([]byte(str)) }
func (g *ZendOutputGlobals) WriteStringUnbuffered(str string) int {
	if g.activated {
		n, _ := SM__().UbWrite(str)
		return n
	}
	_, _ = io.WriteString(g.directOutput, str)
	return len(str)
}

func (g *ZendOutputGlobals) Flush() bool {
	if g.active != nil && g.active.IsFlushable() {
		context := InitOutputContext(PHP_OUTPUT_HANDLER_FLUSH)
		PhpOutputHandlerOp(g.active, context)
		if context.GetOut().Used() > 0 {
			OG__().PopHandler()
			OG__().Write(context.GetOut().Bytes())
			OG__().PushHandler(g.active)
		}
		return true
	}
	return false
}

func (g *ZendOutputGlobals) Clean() bool {
	if g.active != nil && g.active.IsCleanable() {
		context := InitOutputContext(PHP_OUTPUT_HANDLER_CLEAN)
		PhpOutputHandlerOp(g.active, context)
		return true
	}
	return false
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
	if g.lockError(PHP_OUTPUT_HANDLER_START) || h == nil {
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
	return g.CountHandlers()
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

//
func (g *ZendOutputGlobals) lockError(op int) bool {
	/* if there's no ob active, ob has been stopped */
	if op != 0 && g.active != nil && g.running != nil {
		/* fatal error */
		OG__().Deactivate()
		PhpErrorDocref("ref.outcontrol", faults.E_ERROR, "Cannot use output buffering in output buffering display handlers")
		return true
	}
	return false
}

func (g *ZendOutputGlobals) operate(op int, data []byte) {
	if g.lockError(op) {
		return
	}

	context := InitOutputContext(op)

	/*
	 * broken up for better performance:
	 *  - apply op to the one active handler; note that OG__().active might be popped off the stack on a flush
	 *  - or apply op to the handler stack
	 */
	if g.active != nil && g.CountHandlers() != 0 {
		obhCnt := OG__().CountHandlers()
		context.GetIn().SetData(data)
		if obhCnt > 1 {
			OG__().EachHandlerEx(false, func(h *PhpOutputHandler) bool {
				return PhpOutputStackApplyOp(h, context) == 0
			})
		} else if active := OG__().TopHandler(); active != nil && !active.IsDisabled() {
			PhpOutputHandlerOp(active, context)
		} else {
			context.Pass()
		}
	} else {
		context.GetOut().SetData(data)
	}
	if outData := context.GetOut().String(); outData != "" {
		PhpOutputHeader()

		if !OG__().IsDisabled() {
			SM__().UbWrite(outData)
			if OG__().IsImplicitFlush() {
				SapiFlush()
			}
			OG__().MarkSent()
		}
	}
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
