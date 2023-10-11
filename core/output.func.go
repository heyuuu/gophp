package core

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/zpp"
)

func PhpOutputHeader() {
	if !SG__().headersSent {
		if OG__().StartFilename() == "" {
			if zend.ZendIsCompiling() {
				OG__().SetStartFilename(zend.ZendGetCompiledFilename())
				OG__().SetStartLineno(zend.ZendGetCompiledLineno())
			} else if zend.ZendIsExecuting() {
				OG__().SetStartFilename(zend.ZendGetExecutedFilename())
				OG__().SetStartLineno(zend.ZendGetExecutedLineno())
			}
		}
		if standard.PhpHeader() == 0 {
			OG__().MarkDisabled()
		}
	}
}
func PhpOutputRegisterConstants() {
	zend.RegisterMainLongConstant("PHP_OUTPUT_HANDLER_START", PHP_OUTPUT_HANDLER_START, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.RegisterMainLongConstant("PHP_OUTPUT_HANDLER_WRITE", PHP_OUTPUT_HANDLER_WRITE, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.RegisterMainLongConstant("PHP_OUTPUT_HANDLER_FLUSH", PHP_OUTPUT_HANDLER_FLUSH, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.RegisterMainLongConstant("PHP_OUTPUT_HANDLER_CLEAN", PHP_OUTPUT_HANDLER_CLEAN, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.RegisterMainLongConstant("PHP_OUTPUT_HANDLER_FINAL", PHP_OUTPUT_HANDLER_FINAL, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.RegisterMainLongConstant("PHP_OUTPUT_HANDLER_CONT", PHP_OUTPUT_HANDLER_WRITE, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.RegisterMainLongConstant("PHP_OUTPUT_HANDLER_END", PHP_OUTPUT_HANDLER_FINAL, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.RegisterMainLongConstant("PHP_OUTPUT_HANDLER_CLEANABLE", PHP_OUTPUT_HANDLER_CLEANABLE, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.RegisterMainLongConstant("PHP_OUTPUT_HANDLER_FLUSHABLE", PHP_OUTPUT_HANDLER_FLUSHABLE, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.RegisterMainLongConstant("PHP_OUTPUT_HANDLER_REMOVABLE", PHP_OUTPUT_HANDLER_REMOVABLE, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.RegisterMainLongConstant("PHP_OUTPUT_HANDLER_STDFLAGS", PHP_OUTPUT_HANDLER_STDFLAGS, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.RegisterMainLongConstant("PHP_OUTPUT_HANDLER_STARTED", PHP_OUTPUT_HANDLER_STARTED, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.RegisterMainLongConstant("PHP_OUTPUT_HANDLER_DISABLED", PHP_OUTPUT_HANDLER_DISABLED, zend.CONST_CS|zend.CONST_PERSISTENT)
}
func PhpOutputEnd() bool {
	return PhpOutputStackPop(PHP_OUTPUT_POP_TRY) != 0
}
func PhpOutputEndAll() {
	for OG__().Active() != nil && PhpOutputStackPop(PHP_OUTPUT_POP_FORCE) != 0 {
	}
}
func PhpOutputDiscard() bool {
	return PhpOutputStackPop(PHP_OUTPUT_POP_DISCARD|PHP_OUTPUT_POP_TRY) != 0
}
func PhpOutputDiscardAll() {
	for OG__().Active() != nil {
		PhpOutputStackPop(PHP_OUTPUT_POP_DISCARD | PHP_OUTPUT_POP_FORCE)
	}
}
func PhpOutputStartDefault() bool {
	handler := NewOutputHandler(PhpOutputDefaultHandlerName, nil, 0, PHP_OUTPUT_HANDLER_STDFLAGS)
	return OG__().StartHandler(handler)
}
func PhpOutputStartUser(outputHandler *types.Zval, chunkSize int, flags int) bool {
	handler := NewOutputHandlerUser(outputHandler, chunkSize, flags)
	return OG__().StartHandler(handler)
}
func PhpOutputStartInternal(name string, outputHandler PhpOutputHandlerFuncT, chunkSize int, flags int) bool {
	handler := NewOutputHandler(name, wrapHandlerFuncT(outputHandler), chunkSize, flags)
	return OG__().StartHandler(handler)
}
func PhpOutputHandlerAppend(handler *PhpOutputHandler, buf *PhpOutputBuffer) bool {
	if buf.Used() != 0 {
		OG__().MarkWritten()

		/* store it away */
		handler.GetBuffer().Append(buf.Bytes())

		/* chunked buffering */
		if handler.GetSize() != 0 && handler.GetBuffer().Used() >= handler.GetSize() {
			/* store away errors and/or any intermediate output */
			if OG__().Running() != nil {
				return true
			} else {
				return false
			}
		}
	}
	return true
}
func PhpOutputHandlerOp(handler *PhpOutputHandler, context *PhpOutputContext) PhpOutputHandlerStatusT {
	if OG__().lockError(context.GetOp()) {
		/* fatal error */
		return PHP_OUTPUT_HANDLER_FAILURE
	}

	var status PhpOutputHandlerStatusT
	var originalOp int = context.GetOp()

	/* storable? */
	if PhpOutputHandlerAppend(handler, context.GetIn()) && context.GetOp() == 0 {
		return PHP_OUTPUT_HANDLER_NO_DATA
	} else {
		/* need to start? */
		if !handler.IsStarted() {
			context.SetOp(context.GetOp() | PHP_OUTPUT_HANDLER_START)
		}
		OG__().running = handler

		context.Feed(*handler.GetBuffer())
		if handler.HandleContext(context) {
			status = PHP_OUTPUT_HANDLER_NO_DATA
			if context.GetOut().Used() != 0 {
				status = PHP_OUTPUT_HANDLER_SUCCESS
			}
		} else {
			status = PHP_OUTPUT_HANDLER_FAILURE
		}
		handler.SetIsStarted(true)
		OG__().running = nil
	}
	switch status {
	case PHP_OUTPUT_HANDLER_FAILURE:
		/* disable this handler */
		handler.SetIsDisabled(true)

		/* returns handlers buffer */
		context.GetOut().SetDataNoClone(handler.GetBuffer().Bytes())
		handler.GetBuffer().Clean()
	case PHP_OUTPUT_HANDLER_NO_DATA:
		/* handler ate all */
		context.Reset()
		fallthrough
	case PHP_OUTPUT_HANDLER_SUCCESS:
		/* no more buffered data */
		handler.GetBuffer().Reset()
		handler.SetIsProcessed(true)
	}
	context.SetOp(originalOp)
	return status
}
func PhpOutputStackApplyOp(handler *PhpOutputHandler, context *PhpOutputContext) int {
	var status PhpOutputHandlerStatusT

	wasDisabled := handler.IsDisabled()
	if wasDisabled {
		status = PHP_OUTPUT_HANDLER_FAILURE
	} else {
		status = PhpOutputHandlerOp(handler, context)
	}

	/*
	 * handler ate all => break
	 * handler returned data or failed resp. is disabled => continue
	 */
	switch status {
	case PHP_OUTPUT_HANDLER_NO_DATA:
		return 1
	case PHP_OUTPUT_HANDLER_SUCCESS:
		/* swap contexts buffers, unless this is the last handler in the stack */
		if handler.GetLevel() != 0 {
			context.Swap()
		}
		return 0
	case PHP_OUTPUT_HANDLER_FAILURE:
		fallthrough
	default:
		if wasDisabled {
			/* pass input along, if it's the last handler in the stack */
			if handler.GetLevel() == 0 {
				context.Pass()
			}
		} else {
			/* swap buffers, unless this is the last handler */
			if handler.GetLevel() != 0 {
				context.Swap()
			}
		}
		return 0
	}
}
func PhpOutputStackPop(flags int) int {
	var orphan *PhpOutputHandler = OG__().Active()
	if orphan == nil {
		if (flags & PHP_OUTPUT_POP_SILENT) == 0 {
			PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to %s buffer. No buffer to %s", lang.Cond((flags&PHP_OUTPUT_POP_DISCARD) != 0, "discard", "send"), lang.Cond((flags&PHP_OUTPUT_POP_DISCARD) != 0, "discard", "send"))
		}
		return 0
	} else if (flags&PHP_OUTPUT_POP_FORCE) == 0 && !orphan.IsRemovable() {
		if (flags & PHP_OUTPUT_POP_SILENT) == 0 {
			PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to %s buffer of %s (%d)", lang.Cond((flags&PHP_OUTPUT_POP_DISCARD) != 0, "discard", "send"), orphan.GetName(), orphan.GetLevel())
		}
		return 0
	} else {
		context := InitOutputContext(PHP_OUTPUT_HANDLER_FINAL)

		/* don't run the output handler if it's disabled */

		if !orphan.IsDisabled() {

			/* didn't it start yet? */

			if !orphan.IsStarted() {
				context.SetOp(context.GetOp() | PHP_OUTPUT_HANDLER_START)
			}

			/* signal that we're cleaning up */

			if (flags & PHP_OUTPUT_POP_DISCARD) != 0 {
				context.SetOp(context.GetOp() | PHP_OUTPUT_HANDLER_CLEAN)
			}
			PhpOutputHandlerOp(orphan, context)
		}

		/* pop it off the stack */
		OG__().EndHandler()

		/* pass output along */
		if context.GetOut().Used() != 0 && (flags&PHP_OUTPUT_POP_DISCARD) == 0 {
			OG__().WriteString(context.GetOut().String())
		}

		/* destroy the handler (after write!) */
		return 1
	}
}

//zif -old "|zll"
func ZifObStart(_ zpp.Opt, userFunction *types.Zval, chunkSize int, flags_ *int) bool {
	var flags = b.Option(flags_, PHP_OUTPUT_HANDLER_STDFLAGS)
	if chunkSize < 0 {
		chunkSize = 0
	}

	if !PhpOutputStartUser(userFunction, chunkSize, flags) {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to create buffer")
		return false
	}
	return true
}
func ZifObFlush() bool {
	if OG__().Active() == nil {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to flush buffer. No buffer to flush")
		return false
	}
	if !OG__().Flush() {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to flush buffer of %s (%d)", OG__().Active().GetName(), OG__().Active().GetLevel())
		return false
	}
	return true
}
func ZifObClean() bool {
	if OG__().Active() == nil {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to delete buffer. No buffer to delete")
		return false
	}
	if !OG__().Clean() {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to delete buffer of %s (%d)", OG__().active.name.GetVal(), OG__().active.level)
		return false
	}
	return true
}
func ZifObEndFlush() bool {
	if OG__().Active() == nil {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to delete and flush buffer. No buffer to delete or flush")
		return false
	}
	return PhpOutputEnd()
}
func ZifObEndClean() bool {
	if OG__().Active() == nil {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to delete buffer. No buffer to delete")
		return false
	}
	return PhpOutputDiscard()
}
func ZifObGetFlush() (string, bool) {
	contents, ok := OG__().GetContents()
	if !ok {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to delete and flush buffer. No buffer to delete or flush")
		return "", false
	}
	if !PhpOutputEnd() {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to delete buffer of %s (%d)", OG__().Active().GetName(), OG__().Active().GetLevel())
	}
	return contents, true
}
func ZifObGetClean() (string, bool) {
	if OG__().Active() == nil {
		return "", false
	}

	contents, ok := OG__().GetContents()
	if !ok {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to delete buffer. No buffer to delete")
		return "", false
	}
	if !PhpOutputDiscard() {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to delete buffer of %s (%d)", OG__().Active().GetName(), OG__().Active().GetLevel())
	}
	return contents, true
}
func ZifObGetContents() (string, bool) {
	return OG__().GetContents()
}
func ZifObGetLevel() int {
	return OG__().GetLevel()
}
func ZifObGetLength() (int, bool) {
	return OG__().GetLength()
}
func ZifObListHandlers() *types.Array {
	retArr := types.NewArray()
	if OG__().Active() != nil {
		OG__().EachHandler(true, func(h *PhpOutputHandler) {
			retArr.Append(types.NewZvalString(h.name))
		})
	}
	return retArr
}

//zif -old "|b"
func ZifObGetStatus(_ zpp.Opt, fullStatus bool) *types.Array {
	if OG__().active == nil {
		return types.NewArray()
	}
	if fullStatus {
		retArr := types.NewArrayCap(OG__().CountHandlers())
		OG__().EachHandler(true, func(h *PhpOutputHandler) {
			status := outputHandlerStatus(h)
			retArr.Append(types.NewZvalArray(status))
		})
		return retArr
	} else {
		return outputHandlerStatus(OG__().active)
	}
}
func outputHandlerStatus(handler *PhpOutputHandler) *types.Array {
	arr := types.NewArrayCap(7)
	arr.KeyUpdate("name", types.NewZvalString(handler.GetName()))
	arr.KeyUpdate("type", types.NewZvalLong(handler.GetFlags()))
	arr.KeyUpdate("flags", types.NewZvalLong(handler.GetFlags()))
	arr.KeyUpdate("level", types.NewZvalLong(handler.GetLevel()))
	arr.KeyUpdate("chunk_size", types.NewZvalLong(handler.GetSize()))
	arr.KeyUpdate("buffer_size", types.NewZvalLong(handler.GetBuffer().Size()))
	arr.KeyUpdate("buffer_used", types.NewZvalLong(handler.GetBuffer().Used()))
	return arr
}

func ZifObImplicitFlush(_ zpp.Opt, flag_ *int) {
	flag := b.Option(flag_, 1)
	OG__().MarkImplicitFlush(flag != 0)
}
func ZifOutputResetRewriteVars() bool {
	standard.PhpUrlScannerResetVars()
	return true
}
func ZifOutputAddRewriteVar(name string, value string) bool {
	return standard.PhpUrlScannerAddVar(name, value, 1)
}
