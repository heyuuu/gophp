package core

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/shim/slices"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

var PhpOutputDefaultHandlerName = "default output handler"

func PhpOutputHandlerDefaultFunc(context *PhpOutputContext) bool {
	context.Pass()
	return true
}

func wrapOutputHandlerUser(outputHandler *types.Zval) (name string, handler OutputHandlerFunc, ok bool) {
	// 未定义或定义值为 nil 时
	if outputHandler == nil || outputHandler.IsNull() {
		return "", nil, true
	}

	// 自定义 callable 时
	var handlerName *types.String = nil
	var err *byte = nil
	user := &PhpOutputHandlerUserFuncT{}
	if types.SUCCESS == zend.ZendFcallInfoInit(outputHandler, 0, user.GetFci(), user.GetFcc(), &handlerName, &err) {
		types.ZVAL_COPY(user.GetZoh(), outputHandler)

		ok = true
		name = handlerName.GetStr()
		handler = func(context *PhpOutputContext) (result bool) {
			zend.ZendFcallInfoArgn(user.GetFci(), 2, types.NewZvalString(context.GetIn().String()), types.NewZvalLong(context.GetOp()))

			var retval types.Zval
			if zend.ZendFcallInfoCall(user.GetFci(), user.GetFcc(), &retval, nil) && !retval.IsUndef() && !retval.IsFalse() {
				result = true
				if !retval.IsTrue() {
					operators.ConvertToStringEx(&retval)
					context.GetOut().SetDataStr(retval.String())
				}
			}

			zend.ZendFcallInfoArgn(user.GetFci(), 0)
			return result
		}
	}
	if err != nil {
		PhpErrorDocref("ref.outcontrol", faults.E_WARNING, "%s", err)
	}
	return
}

func wrapHandlerFuncT(handler PhpOutputHandlerFuncT) OutputHandlerFunc {
	if handler == nil {
		return func(context *PhpOutputContext) bool {
			return false
		}
	}

	h := func(output string, mode int) (handledOutput string) {
		var outputPtr *byte
		var outputLen int
		handler(b.CastStrPtr(output), len(output), &outputPtr, &outputLen, mode)
		return b.CastStr(outputPtr, outputLen)
	}

	return func(context *PhpOutputContext) bool {
		var handledOutput string
		if data := context.GetOut().String(); data != "" {
			handledOutput = h(data, context.GetOp())
		}
		if len(handledOutput) > 0 {
			context.GetOut().SetDataStr(handledOutput)
		} else {
			context.Pass()
		}
		return true
	}
}

func wrapHandlerFuncContextT(handler PhpOutputHandlerContextFuncT) OutputHandlerFunc {
	return func(context *PhpOutputContext) bool {
		return handler(nil, context) == types.SUCCESS
	}
}

/**
 * types
 */
type OutputHandlerFunc func(context *PhpOutputContext) bool

// PhpOutputHandler
type PhpOutputHandler struct {
	name    string
	flags   int
	level   int
	size    int
	buffer  PhpOutputBuffer
	handler OutputHandlerFunc
}

func NewOutputHandlerUser(outputHandler *types.Zval, chunkSize int, flags int) *PhpOutputHandler {
	if name, userHandler, ok := wrapOutputHandlerUser(outputHandler); ok {
		return NewOutputHandler(name, userHandler, chunkSize, flags & ^0xf)
	} else {
		return nil
	}
}

func NewOutputHandler(name string, handlerFunc OutputHandlerFunc, chunkSize int, flags int) *PhpOutputHandler {
	if name == "" {
		name = PhpOutputDefaultHandlerName
	}
	handler := &PhpOutputHandler{
		name:    name,
		size:    chunkSize,
		flags:   flags &^ 0xf,
		handler: handlerFunc,
	}
	handler.buffer.Init(chunkSize)
	return handler
}

func (h *PhpOutputHandler) HandleContext(context *PhpOutputContext) bool {
	if h.handler != nil {
		return h.handler(context)
	}
	return PhpOutputHandlerDefaultFunc(context)
}

func (h *PhpOutputHandler) GetName() string             { return h.name }
func (h *PhpOutputHandler) GetFlags() int               { return h.flags }
func (h *PhpOutputHandler) GetLevel() int               { return h.level }
func (h *PhpOutputHandler) GetSize() int                { return h.size }
func (h *PhpOutputHandler) GetBuffer() *PhpOutputBuffer { return &h.buffer }

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
func (h PhpOutputHandler) IsStarted() bool   { return h.HasFlags(PHP_OUTPUT_HANDLER_STARTED) }
func (h PhpOutputHandler) IsDisabled() bool  { return h.HasFlags(PHP_OUTPUT_HANDLER_DISABLED) }
func (h PhpOutputHandler) IsProcessed() bool { return h.HasFlags(PHP_OUTPUT_HANDLER_PROCESSED) }
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

func (this *PhpOutputHandlerUserFuncT) GetFci() *types.ZendFcallInfo      { return &this.fci }
func (this *PhpOutputHandlerUserFuncT) GetFcc() *types.ZendFcallInfoCache { return &this.fcc }
func (this *PhpOutputHandlerUserFuncT) GetZoh() *types.Zval               { return &this.zoh }
