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
		if OG__().OutputStartFilename() == "" {
			if zend.ZendIsCompiling() {
				OG__().SetOutputStartFilename(zend.ZendGetCompiledFilename())
				OG__().SetOutputStartLineno(zend.ZendGetCompiledLineno())
			} else if zend.ZendIsExecuting() {
				OG__().SetOutputStartFilename(zend.ZendGetExecutedFilename())
				OG__().SetOutputStartLineno(zend.ZendGetExecutedLineno())
			}
		}
		if standard.PhpHeader() == 0 {
			OG__().MarkDisabled()
		}
	}
}
func PhpOutputStartup() {
	OutputGlobals.Init()
	PhpOutputDirect = PhpOutputStdout
}
func PhpOutputShutdown() {
	PhpOutputDirect = PhpOutputStderr
}
func PhpOutputActivate() int {
	OG__().Activate()
	return types.SUCCESS
}
func PhpOutputDeactivate() {
	if OG__().IsActivated() {
		PhpOutputHeader()
		OG__().Deactivate()
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
func PhpOutputFlush() int {
	if active := OG__().Active(); active != nil && active.IsFlushable() {
		context := InitOutputContext(PHP_OUTPUT_HANDLER_FLUSH)
		PhpOutputHandlerOp(active, context)
		if data, ok := context.GetOutData(); ok {
			OG__().PopHandler()
			PhpOutputWrite(data)
			OG__().PushHandler(&active)
		}
		return types.SUCCESS
	}
	return types.FAILURE
}
func PhpOutputClean() int {
	if active := OG__().Active(); active != nil && active.IsCleanable() {
		context := InitOutputContext(PHP_OUTPUT_HANDLER_CLEAN)
		PhpOutputHandlerOp(OG__().active, context)
		return types.SUCCESS
	}
	return types.FAILURE
}
func PhpOutputEnd() int {
	if PhpOutputStackPop(PHP_OUTPUT_POP_TRY) != 0 {
		return types.SUCCESS
	}
	return types.FAILURE
}
func PhpOutputEndAll() {
	for OG__().Active() != nil && PhpOutputStackPop(PHP_OUTPUT_POP_FORCE) != 0 {
	}
}
func PhpOutputDiscard() int {
	if PhpOutputStackPop(PHP_OUTPUT_POP_DISCARD|PHP_OUTPUT_POP_TRY) != 0 {
		return types.SUCCESS
	}
	return types.FAILURE
}
func PhpOutputDiscardAll() {
	for OG__().Active() != nil {
		PhpOutputStackPop(PHP_OUTPUT_POP_DISCARD | PHP_OUTPUT_POP_FORCE)
	}
}
func PhpOutputGetLevel() int {
	if OG__().Active() != nil {
		return OG__().CountHandlers()
	} else {
		return 0
	}
}
func PhpOutputGetContents(p *types.Zval) int {
	if OG__().Active() != nil {
		p.SetString(b.CastStr(OG__().active.buffer.data, OG__().active.buffer.used))
		return types.SUCCESS
	} else {
		p.SetNull()
		return types.FAILURE
	}
}
func PhpOutputGetLength(p *types.Zval) int {
	if OG__().Active() != nil {
		p.SetLong(OG__().active.buffer.used)
		return types.SUCCESS
	} else {
		p.SetNull()
		return types.FAILURE
	}
}
func PhpOutputStartDefault() int {
	var handler *PhpOutputHandler
	handler = NewOutputHandlerInternal(PhpOutputDefaultHandlerName, PhpOutputHandlerDefaultFunc, 0, PHP_OUTPUT_HANDLER_STDFLAGS)
	if types.SUCCESS == PhpOutputHandlerStart(handler) {
		return types.SUCCESS
	}
	PhpOutputHandlerFree(&handler)
	return types.FAILURE
}
func PhpOutputStartUser(outputHandler *types.Zval, chunkSize int, flags int) int {
	handler := NewOutputHandlerUser(outputHandler, chunkSize, flags)
	if types.SUCCESS == PhpOutputHandlerStart(handler) {
		return types.SUCCESS
	}
	PhpOutputHandlerFree(&handler)
	return types.FAILURE
}
func PhpOutputStartInternal(name string, output_handler PhpOutputHandlerFuncT, chunk_size int, flags int) int {
	var handler *PhpOutputHandler
	handler = NewOutputHandlerInternal(name, PhpOutputHandlerCompatFunc, chunk_size, flags)
	PhpOutputHandlerSetContext(handler, output_handler)
	if types.SUCCESS == PhpOutputHandlerStart(handler) {
		return types.SUCCESS
	}
	PhpOutputHandlerFree(&handler)
	return types.FAILURE
}
func PhpOutputHandlerSetContext(handler *PhpOutputHandler, opaq any) {
	handler.SetOpaq(opaq)
}
func PhpOutputHandlerStart(handler *PhpOutputHandler) int {
	if PhpOutputLockError(PHP_OUTPUT_HANDLER_START) != 0 || handler == nil {
		return types.FAILURE
	}

	/* zend_stack_push returns stack level */

	handler.SetLevel(OG__().PushHandler(&handler))
	OG__().active = handler
	return types.SUCCESS
}
func PhpOutputHandlerFree(h **PhpOutputHandler) {
	if (*h) != nil {
		*h = nil
	}
}
func PhpOutputSetImplicitFlush(flush int) {
	OG__().MarkImplicitFlush(flush != 0)
}
func PhpOutputGetStartFilename() string { return OG__().OutputStartFilename() }
func PhpOutputGetStartLineno() int      { return OG__().OutputStartLineno() }
func PhpOutputLockError(op int) int {
	/* if there's no ob active, ob has been stopped */
	if op != 0 && OG__().active != nil && OG__().running != nil {
		/* fatal error */
		PhpOutputDeactivate()
		PhpErrorDocref("ref.outcontrol", faults.E_ERROR, "Cannot use output buffering in output buffering display handlers")
		return 1
	}
	return 0
}
func PhpOutputHandlerAppend(handler *PhpOutputHandler, buf *PhpOutputBuffer) int {
	if buf.GetUsed() != 0 {
		OG__().MarkWritten()

		/* store it away */
		if handler.GetBuffer().GetSize()-handler.GetBuffer().GetUsed() <= buf.GetUsed() {
			var grow_int int = PHP_OUTPUT_HANDLER_INITBUF_SIZE(handler.GetSize())
			var grow_buf int = PHP_OUTPUT_HANDLER_INITBUF_SIZE(buf.GetUsed() - (handler.GetBuffer().GetSize() - handler.GetBuffer().GetUsed()))
			var grow_max int = b.Max(grow_int, grow_buf)
			handler.GetBuffer().SetData(zend.SafeErealloc(handler.GetBuffer().GetData(), 1, handler.GetBuffer().GetSize(), grow_max))
			handler.GetBuffer().SetSize(handler.GetBuffer().GetSize() + grow_max)
		}
		memcpy(handler.GetBuffer().GetData()+handler.GetBuffer().GetUsed(), buf.GetData(), buf.GetUsed())
		handler.GetBuffer().SetUsed(handler.GetBuffer().GetUsed() + buf.GetUsed())

		/* chunked buffering */

		if handler.GetSize() != 0 && handler.GetBuffer().GetUsed() >= handler.GetSize() {

			/* store away errors and/or any intermediate output */

			if OG__().running {
				return 1
			} else {
				return 0
			}

			/* store away errors and/or any intermediate output */

		}

		/* chunked buffering */

	}
	return 1
}
func PhpOutputHandlerOp(handler *PhpOutputHandler, context *PhpOutputContext) PhpOutputHandlerStatusT {
	var status PhpOutputHandlerStatusT
	var originalOp int = context.GetOp()
	if PhpOutputLockError(context.GetOp()) != 0 {
		/* fatal error */
		return PHP_OUTPUT_HANDLER_FAILURE
	}

	/* storable? */
	if PhpOutputHandlerAppend(handler, context.GetIn()) != 0 && context.GetOp() == 0 {
		context.SetOp(originalOp)
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
			ob_data.SetString(b.CastStr(handler.GetBuffer().GetData(), handler.GetBuffer().GetUsed()))
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
						context.GetOut().SetFreeDataByStr(retval.String())
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
				if context.GetOut().GetUsed() != 0 {
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
		if context.GetOut().GetData() != nil && context.GetOut().IsFree() {
			zend.Efree(context.GetOut().GetData())
		}

		/* returns handlers buffer */
		context.GetOut().SetFreeData(handler.GetBuffer())
		handler.GetBuffer().SetData(nil)
		handler.GetBuffer().SetUsed(0)
		handler.GetBuffer().SetSize(0)
	case PHP_OUTPUT_HANDLER_NO_DATA:
		/* handler ate all */
		context.Reset()
		fallthrough
	case PHP_OUTPUT_HANDLER_SUCCESS:
		/* no more buffered data */
		handler.GetBuffer().SetUsed(0)
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
		context.GetIn().SetData((*byte)(str))
		context.GetIn().SetUsed(len_)
		if obh_cnt > 1 {
			zend.ZendStackApplyWithArgument(&(OG__().handlers), zend.ZEND_STACK_APPLY_TOPDOWN, PhpOutputStackApplyOp, context)
		} else if active := OG__().TopHandler(); active != nil && !(*active).IsDisabled() {
			PhpOutputHandlerOp(*active, context)
		} else {
			context.Pass()
		}
	} else {
		context.GetOut().SetData((*byte)(str))
		context.GetOut().SetUsed(len_)
	}
	if context.GetOut().GetData() != nil && context.GetOut().GetUsed() != 0 {
		PhpOutputHeader()

		if !OG__().IsDisabled() {
			SM__().UbWrite(
				b.CastStr(context.GetOut().GetData(), context.GetOut().GetUsed()),
			)
			if OG__().IsImplicitFlush() {
				SapiFlush()
			}
			OG__().MarkSent()
		}
	}
}
func PhpOutputStackApplyOp(h any, c any) int {
	var was_disabled int
	var status PhpOutputHandlerStatusT
	var handler *PhpOutputHandler = *((**PhpOutputHandler)(h))
	var context *PhpOutputContext = (*PhpOutputContext)(c)
	if lang.Assign(&was_disabled, handler.GetFlags()&PHP_OUTPUT_HANDLER_DISABLED) {
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
		if was_disabled != 0 {
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

	/*
	 * handler ate all => break
	 * handler returned data or failed resp. is disabled => continue
	 */
}
func PhpOutputStackApplyList(h any, z any) int {
	var handler *PhpOutputHandler = *((**PhpOutputHandler)(h))
	var array *types.Zval = (*types.Zval)(z)
	zend.AddNextIndexStr(array, handler.GetName().Copy())
	array.Array().Append(handler.GetName())
	return 0
}
func PhpOutputStackPop(flags int) int {
	var current **PhpOutputHandler
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
		OG__().PopHandler()
		if current = OG__().TopHandler(); current != nil {
			OG__().SetActive(*current)
		} else {
			OG__().SetActive(nil)
		}

		/* pass output along */
		if context.GetOut().GetData() != nil && context.GetOut().GetUsed() != 0 && (flags&PHP_OUTPUT_POP_DISCARD) == 0 {
			PhpOutputWrite(context.GetOut().GetData(), context.GetOut().GetUsed())
		}

		/* destroy the handler (after write!) */

		PhpOutputHandlerFree(&orphan)
		return 1
	}
}
func PhpOutputHandlerCompatFunc(handler_context *any, output_context *PhpOutputContext) int {
	var func_ = (*handler_context).(PhpOutputHandlerFuncT)
	var handler = wrapOutputHandler(func_)

	if handler != nil {
		var handledOutput string
		if data, ok := output_context.GetOutData(); ok {
			handledOutput = handler(data, output_context.GetOp())
		}
		if len(handledOutput) > 0 {
			output_context.GetOut().SetFreeDataByStr(handledOutput)
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

	if PhpOutputStartUser(userFunction, chunkSize, flags) == types.FAILURE {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to create buffer")
		return false
	}
	return true
}
func ZifObFlush(executeData zpp.Ex, return_value zpp.Ret) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if !(OG__().active) {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to flush buffer. No buffer to flush")
		return_value.SetFalse()
		return
	}
	if types.SUCCESS != PhpOutputFlush() {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to flush buffer of %s (%d)", OG__().active.name.GetVal(), OG__().active.level)
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
}
func ZifObClean(executeData zpp.Ex, return_value zpp.Ret) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if !(OG__().active) {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to delete buffer. No buffer to delete")
		return_value.SetFalse()
		return
	}
	if types.SUCCESS != PhpOutputClean() {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to delete buffer of %s (%d)", OG__().active.name.GetVal(), OG__().active.level)
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
}
func ZifObEndFlush(executeData zpp.Ex, return_value zpp.Ret) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if !(OG__().active) {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to delete and flush buffer. No buffer to delete or flush")
		return_value.SetFalse()
		return
	}
	return_value.SetBool(types.SUCCESS == PhpOutputEnd())
	return
}
func ZifObEndClean(executeData zpp.Ex, return_value zpp.Ret) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if !(OG__().active) {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to delete buffer. No buffer to delete")
		return_value.SetFalse()
		return
	}
	return_value.SetBool(types.SUCCESS == PhpOutputDiscard())
	return
}
func ZifObGetFlush(executeData zpp.Ex, return_value zpp.Ret) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if PhpOutputGetContents(return_value) == types.FAILURE {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to delete and flush buffer. No buffer to delete or flush")
		return_value.SetFalse()
		return
	}
	if types.SUCCESS != PhpOutputEnd() {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to delete buffer of %s (%d)", OG__().active.name.GetVal(), OG__().active.level)
	}
}
func ZifObGetClean(executeData zpp.Ex, return_value zpp.Ret) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if !(OG__().active) {
		return_value.SetFalse()
		return
	}
	if PhpOutputGetContents(return_value) == types.FAILURE {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to delete buffer. No buffer to delete")
		return_value.SetFalse()
		return
	}
	if types.SUCCESS != PhpOutputDiscard() {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to delete buffer of %s (%d)", OG__().active.name.GetVal(), OG__().active.level)
	}
}
func ZifObGetContents(executeData zpp.Ex, return_value zpp.Ret) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if PhpOutputGetContents(return_value) == types.FAILURE {
		return_value.SetFalse()
		return
	}
}
func ZifObGetLevel(executeData zpp.Ex, return_value zpp.Ret) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	return_value.SetLong(PhpOutputGetLevel())
	return
}
func ZifObGetLength(executeData zpp.Ex, return_value zpp.Ret) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if PhpOutputGetLength(return_value) == types.FAILURE {
		return_value.SetFalse()
		return
	}
}
func ZifObListHandlers(executeData zpp.Ex, return_value zpp.Ret) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	zend.ArrayInit(return_value)
	if !(OG__().active) {
		return
	}
	zend.ZendStackApplyWithArgument(&(OG__().handlers), zend.ZEND_STACK_APPLY_BOTTOMUP, PhpOutputStackApplyList, return_value)
}

//zif -old "|b"
func ZifObGetStatus(_ zpp.Opt, fullStatus bool) *types.Array {
	if OG__().active == nil {
		return types.NewArray()
	}
	if fullStatus {
		retArr := types.NewArrayCap(OG__().CountHandlers())
		OG__().EachHandler(true, func(h **PhpOutputHandler) {
			status := outputHandlerStatus(*h)
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
	arr.KeyUpdate("buffer_size", types.NewZvalLong(handler.GetBuffer().GetSize()))
	arr.KeyUpdate("buffer_used", types.NewZvalLong(handler.GetBuffer().GetUsed()))
	return arr
}

func ZifObImplicitFlush(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, flag *types.Zval) {
	var flag zend.ZendLong = 1
	if zend.ZendParseParameters(executeData.NumArgs(), "|l", &flag) == types.FAILURE {
		return
	}
	PhpOutputSetImplicitFlush(flag)
}
func ZifOutputResetRewriteVars(executeData zpp.Ex, return_value zpp.Ret) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if standard.PhpUrlScannerResetVars() == types.SUCCESS {
		return_value.SetTrue()
		return
	} else {
		return_value.SetFalse()
		return
	}
}
func ZifOutputAddRewriteVar(executeData zpp.Ex, return_value zpp.Ret, name *types.Zval, value *types.Zval) {
	var name *byte
	var value *byte
	var name_len int
	var value_len int
	if zend.ZendParseParameters(executeData.NumArgs(), "ss", &name, &name_len, &value, &value_len) == types.FAILURE {
		return
	}
	if standard.PhpUrlScannerAddVar(name, name_len, value, value_len, 1) == types.SUCCESS {
		return_value.SetTrue()
		return
	} else {
		return_value.SetFalse()
		return
	}
}
