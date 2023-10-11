package core

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"github.com/heyuuu/gophp/zend/zpp"
	"os"
)

func PhpOutputStdout(str string) int {
	os.Stdout.WriteString(str)
	return len(str)
}
func PhpOutputStderr(str string) int {
	os.Stderr.WriteString(str)
	return len(str)
}
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
func PhpOutputWriteUnbuffered(str string) int {
	if OG__().IsActivated() {
		return SM__().UbWrite(str)
	}
	return PhpOutputDirect(str)
}
func PhpOutputWrite(str string) int {
	if OG__().IsActivated() {
		ptr := b.CastStrPtr(str)
		len_ := len(str)
		PhpOutputOp(PHP_OUTPUT_HANDLER_WRITE, ptr, len_)
		return len(str)
	}
	if OG__().IsDisabled() {
		return 0
	}
	return PhpOutputDirect(str)
}
func PhpOutputFlush() bool {
	if active := OG__().Active(); active != nil && active.IsFlushable() {
		context := InitOutputContext(PHP_OUTPUT_HANDLER_FLUSH)
		PhpOutputHandlerOp(active, context)
		if data := context.GetOut().String(); data != "" {
			OG__().PopHandler()
			PhpOutputWrite(data)
			OG__().PushHandler(active)
		}
		return true
	}
	return false
}
func PhpOutputClean() bool {
	if active := OG__().Active(); active != nil && active.IsCleanable() {
		context := InitOutputContext(PHP_OUTPUT_HANDLER_CLEAN)
		PhpOutputHandlerOp(OG__().active, context)
		return true
	}
	return false
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
func PhpOutputGetContents(p *types.Zval) bool {
	if contents, ok := OG__().GetContents(); ok {
		p.SetString(contents)
		return true
	} else {
		p.SetNull()
		return false
	}
}
func PhpOutputStartDefault() bool {
	handler := NewOutputHandlerInternal(PhpOutputDefaultHandlerName, PhpOutputHandlerDefaultFunc, 0, PHP_OUTPUT_HANDLER_STDFLAGS)
	return PhpOutputHandlerStart(handler)
}
func PhpOutputStartUser(outputHandler *types.Zval, chunkSize int, flags int) bool {
	handler := NewOutputHandlerUser(outputHandler, chunkSize, flags)
	return PhpOutputHandlerStart(handler)
}
func PhpOutputStartInternal(name string, outputHandler PhpOutputHandlerFuncT, chunkSize int, flags int) bool {
	handler := NewOutputHandlerInternal(name, PhpOutputHandlerCompatFunc, chunkSize, flags)
	PhpOutputHandlerSetContext(handler, outputHandler)
	return PhpOutputHandlerStart(handler)
}
func PhpOutputHandlerSetContext(handler *PhpOutputHandler, opaq any) {
	handler.SetOpaq(opaq)
}
func PhpOutputHandlerStart(handler *PhpOutputHandler) bool {
	if PhpOutputLockError(PHP_OUTPUT_HANDLER_START) != 0 {
		return false
	}
	return OG__().StartHandler(handler)
}
func PhpOutputGetStartFilename() string { return OG__().StartFilename() }
func PhpOutputGetStartLineno() int      { return OG__().StartLineno() }
func PhpOutputLockError(op int) int {
	/* if there's no ob active, ob has been stopped */
	if op != 0 && OG__().active != nil && OG__().running != nil {
		/* fatal error */
		OG__().Deactivate()
		PhpErrorDocref("ref.outcontrol", faults.E_ERROR, "Cannot use output buffering in output buffering display handlers")
		return 1
	}
	return 0
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
	if PhpOutputLockError(context.GetOp()) != 0 {
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
		if handler.IsUser() {
			var retval types.Zval
			var ob_data types.Zval
			var ob_mode types.Zval
			ob_data.SetString(handler.GetBuffer().String())
			ob_mode.SetLong(zend.ZendLong(context.GetOp()))
			zend.ZendFcallInfoArgn(handler.GetUser().GetFci(), 2, &ob_data, &ob_mode)
			// zend.ZvalPtrDtor(&ob_data)
			var PHP_OUTPUT_USER_SUCCESS func(retval types.Zval) bool = func(retval types.Zval) bool {
				return retval.IsNotUndef() && !retval.IsFalse()
			}
			if types.SUCCESS == zend.ZendFcallInfoCall(handler.GetUser().GetFci(), handler.GetUser().GetFcc(), &retval, nil) && PHP_OUTPUT_USER_SUCCESS(retval) {

				/* user handler may have returned TRUE */

				status = PHP_OUTPUT_HANDLER_NO_DATA
				if !retval.IsFalse() && !retval.IsTrue() {
					operators.ConvertToStringEx(&retval)
					if retval.StringEx().GetLen() != 0 {
						context.GetOut().SetDataStr(retval.String())
						status = PHP_OUTPUT_HANDLER_SUCCESS
					}
				}
			} else {
				/* call failed, pass internal buffer along */
				status = PHP_OUTPUT_HANDLER_FAILURE
			}
			zend.ZendFcallInfoArgn(handler.GetUser().GetFci(), 0)
		} else {
			context.Feed(*handler.GetBuffer())
			if types.SUCCESS == handler.GetInternal()(handler.GetOpaq(), context) {
				if context.GetOut().Used() != 0 {
					status = PHP_OUTPUT_HANDLER_SUCCESS
				} else {
					status = PHP_OUTPUT_HANDLER_NO_DATA
				}
			} else {
				status = PHP_OUTPUT_HANDLER_FAILURE
			}
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
func PhpOutputOp(op int, str *byte, len_ int) {
	if PhpOutputLockError(op) != 0 {
		return
	}

	context := InitOutputContext(op)

	/*
	 * broken up for better performance:
	 *  - apply op to the one active handler; note that OG__().active might be popped off the stack on a flush
	 *  - or apply op to the handler stack
	 */
	if OG__().Active() != nil && OG__().CountHandlers() != 0 {
		obh_cnt := OG__().CountHandlers()
		context.GetIn().SetData(b.CastBytes(str, len_))
		if obh_cnt > 1 {
			OG__().EachHandlerEx(false, func(h *PhpOutputHandler) bool {
				return PhpOutputStackApplyOp(h, context) == 0
			})
		} else if active := OG__().TopHandler(); active != nil && !active.IsDisabled() {
			PhpOutputHandlerOp(active, context)
		} else {
			context.Pass()
		}
	} else {
		context.GetOut().SetData(b.CastBytes(str, len_))
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
			PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to %s buffer of %s (%d)", lang.Cond((flags&PHP_OUTPUT_POP_DISCARD) != 0, "discard", "send"), orphan.Name(), orphan.GetLevel())
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
			PhpOutputWrite(context.GetOut().String())
		}

		/* destroy the handler (after write!) */
		return 1
	}
}
func PhpOutputHandlerCompatFunc(handler_context *any, output_context *PhpOutputContext) int {
	var func_ = (*handler_context).(PhpOutputHandlerFuncT)
	var handler = wrapOutputHandler(func_)

	if handler != nil {
		var handledOutput string
		if data := output_context.GetOut().String(); data != "" {
			handledOutput = handler(data, output_context.GetOp())
		}
		if len(handledOutput) > 0 {
			output_context.GetOut().SetDataStr(handledOutput)
		} else {
			output_context.Pass()
		}
		return types.SUCCESS
	}
	return types.FAILURE
}
func PhpOutputHandlerDefaultFunc(handler_context *any, output_context *PhpOutputContext) int {
	output_context.Pass()
	return types.SUCCESS
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
	if !PhpOutputFlush() {
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
	if !PhpOutputClean() {
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
	arr.KeyUpdate("type", types.NewZvalLong(handler.GetFlags()&0xf))
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
