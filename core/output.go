package core

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
)

var OutputGlobals ZendOutputGlobals
var PhpOutputDefaultHandlerName = "default output handler"

var PhpOutputDirect func(str string) int = PhpOutputStderr

/**
 * types
 */

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

func PhpOutputHandlerCreateUser(outputHandler *types.Zval, chunkSize int, flags int) *PhpOutputHandler {
	var handlerName *types.String = nil
	var err *byte = nil
	var handler *PhpOutputHandler = nil
	switch outputHandler.Type() {
	case types.IsNull:
		handler = PhpOutputHandlerCreateInternal(PhpOutputDefaultHandlerName, PhpOutputHandlerDefaultFunc, chunkSize, flags)
	case types.IsString:
		fallthrough
	default:
		user := &PhpOutputHandlerUserFuncT{}
		if types.SUCCESS == zend.ZendFcallInfoInit(outputHandler, 0, user.GetFci(), user.GetFcc(), &handlerName, &error) {
			handler = NewPhpOutputHandler(handlerName.GetStr(), chunkSize, flags & ^0xf | PHP_OUTPUT_HANDLER_USER)
			types.ZVAL_COPY(user.GetZoh(), outputHandler)
			handler.SetUser(user)
		} else {
			zend.Efree(user)
		}
		if err != nil {
			PhpErrorDocref("ref.outcontrol", faults.E_WARNING, "%s", err)
		}
	}
	return handler
}

func PhpOutputHandlerCreateInternal(name string, outputHandler PhpOutputHandlerContextFuncT, chunkSize int, flags int) *PhpOutputHandler {
	var handler *PhpOutputHandler
	var str = types.NewString(name)
	handler = PhpOutputHandlerInit(str, chunkSize, flags & ^0xf | PHP_OUTPUT_HANDLER_INTERNAL)
	handler.SetInternal(outputHandler)
	return handler
}
func PhpOutputHandlerInit(name *types.String, chunkSize int, flags int) *PhpOutputHandler {
	return NewPhpOutputHandler(name.GetStr(), flags, chunkSize)
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
