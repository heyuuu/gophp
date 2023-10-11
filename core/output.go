package core

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/shim/slices"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

const PhpOutputDefaultHandlerName = "default output handler"

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

type OutputHandler func(context *PhpOutputContext) bool

func (h *PhpOutputHandler) Handler() OutputHandler {
	if h.IsUser() {
		user := h.func_.user
		return func(context *PhpOutputContext) (result bool) {
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
	} else {
		internal := h.func_.internal
		return func(context *PhpOutputContext) bool {
			return internal(h.GetOpaq(), context) == types.SUCCESS
		}
	}
}

func (h *PhpOutputHandler) GetName() string                           { return h.name }
func (h *PhpOutputHandler) GetFlags() int                             { return h.flags }
func (h *PhpOutputHandler) GetLevel() int                             { return h.level }
func (h *PhpOutputHandler) GetSize() int                              { return h.size }
func (h *PhpOutputHandler) GetBuffer() *PhpOutputBuffer               { return &h.buffer }
func (h *PhpOutputHandler) GetOpaq() any                              { return h.opaq }
func (h *PhpOutputHandler) GetUser() *PhpOutputHandlerUserFuncT       { return h.func_.user }
func (h *PhpOutputHandler) GetInternal() PhpOutputHandlerContextFuncT { return h.func_.internal }
func (h *PhpOutputHandler) SetLevel(value int)                        { h.level = value }
func (h *PhpOutputHandler) SetOpaq(value any)                         { h.opaq = value }
func (h *PhpOutputHandler) SetUser(value *PhpOutputHandlerUserFuncT)  { h.func_.user = value }
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

func (this *PhpOutputHandlerUserFuncT) GetFci() *types.ZendFcallInfo      { return &this.fci }
func (this *PhpOutputHandlerUserFuncT) GetFcc() *types.ZendFcallInfoCache { return &this.fcc }
func (this *PhpOutputHandlerUserFuncT) GetZoh() *types.Zval               { return &this.zoh }
