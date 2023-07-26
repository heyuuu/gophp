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
			if zend.ZendIsCompiling() != 0 {
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
	memset(&OutputGlobals, 0, b.SizeOf("zend_output_globals"))
	OG__().Handlers().Init()
	OG__().SetActivated(true)
	return types.SUCCESS
}
func PhpOutputDeactivate() {
	var handler **PhpOutputHandler = nil
	if OG__().IsActivated() {
		PhpOutputHeader()
		OG__().SetActivated(false)
		OG__().active = nil
		OG__().running = nil

		/* release all output handlers */

		if OG__().Handlers().elements {
			for lang.Assign(&handler, OG__().Handlers().Top()) {
				PhpOutputHandlerFree(handler)
				OG__().Handlers().DelTop()
			}
		}
		OG__().Handlers().Destroy()
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
	var context PhpOutputContext
	if OG__().active && (OG__().active.flags&PHP_OUTPUT_HANDLER_FLUSHABLE) != 0 {
		context.Init(PHP_OUTPUT_HANDLER_FLUSH)
		PhpOutputHandlerOp(OG__().active, &context)
		if context.GetOut().GetData() != nil && context.GetOut().GetUsed() != 0 {
			OG__().Handlers().DelTop()
			PhpOutputWrite(context.GetOut().GetData(), context.GetOut().GetUsed())
			OG__().Handlers().Push(&(OG__().active))
		}
		return types.SUCCESS
	}
	return types.FAILURE
}
func PhpOutputClean() int {
	var context PhpOutputContext
	if OG__().active && (OG__().active.flags&PHP_OUTPUT_HANDLER_CLEANABLE) != 0 {
		context.Init(PHP_OUTPUT_HANDLER_CLEAN)
		PhpOutputHandlerOp(OG__().active, &context)
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
	for OG__().active && PhpOutputStackPop(PHP_OUTPUT_POP_FORCE) != 0 {

	}
}
func PhpOutputDiscard() int {
	if PhpOutputStackPop(PHP_OUTPUT_POP_DISCARD|PHP_OUTPUT_POP_TRY) != 0 {
		return types.SUCCESS
	}
	return types.FAILURE
}
func PhpOutputDiscardAll() {
	for OG__().active {
		PhpOutputStackPop(PHP_OUTPUT_POP_DISCARD | PHP_OUTPUT_POP_FORCE)
	}
}
func PhpOutputGetLevel() int {
	if OG__().active {
		return OG__().Handlers().GetTop()
	} else {
		return 0
	}
}
func PhpOutputGetContents(p *types.Zval) int {
	if OG__().active {
		p.SetStringVal(b.CastStr(OG__().active.buffer.data, OG__().active.buffer.used))
		return types.SUCCESS
	} else {
		p.SetNull()
		return types.FAILURE
	}
}
func PhpOutputGetLength(p *types.Zval) int {
	if OG__().active {
		p.SetLong(OG__().active.buffer.used)
		return types.SUCCESS
	} else {
		p.SetNull()
		return types.FAILURE
	}
}
func PhpOutputStartDefault() int {
	var handler *PhpOutputHandler
	handler = PhpOutputHandlerCreateInternal(PhpOutputDefaultHandlerName, PhpOutputHandlerDefaultFunc, 0, PHP_OUTPUT_HANDLER_STDFLAGS)
	if types.SUCCESS == PhpOutputHandlerStart(handler) {
		return types.SUCCESS
	}
	PhpOutputHandlerFree(&handler)
	return types.FAILURE
}
func PhpOutputStartUser(output_handler *types.Zval, chunk_size int, flags int) int {
	var handler *PhpOutputHandler
	if output_handler != nil {
		handler = PhpOutputHandlerCreateUser(output_handler, chunk_size, flags)
	} else {
		handler = PhpOutputHandlerCreateInternal(PhpOutputDefaultHandlerName, PhpOutputHandlerDefaultFunc, chunk_size, flags)
	}
	if types.SUCCESS == PhpOutputHandlerStart(handler) {
		return types.SUCCESS
	}
	PhpOutputHandlerFree(&handler)
	return types.FAILURE
}
func PhpOutputStartInternal(name string, output_handler PhpOutputHandlerFuncT, chunk_size int, flags int) int {
	var handler *PhpOutputHandler
	handler = PhpOutputHandlerCreateInternal(name, PhpOutputHandlerCompatFunc, chunk_size, flags)
	PhpOutputHandlerSetContext(handler, output_handler)
	if types.SUCCESS == PhpOutputHandlerStart(handler) {
		return types.SUCCESS
	}
	PhpOutputHandlerFree(&handler)
	return types.FAILURE
}
func PhpOutputHandlerCreateUser(output_handler *types.Zval, chunk_size int, flags int) *PhpOutputHandler {
	var handler_name *types.String = nil
	var err *byte = nil
	var handler *PhpOutputHandler = nil
	switch output_handler.GetType() {
	case types.IS_NULL:
		handler = PhpOutputHandlerCreateInternal(PhpOutputDefaultHandlerName, PhpOutputHandlerDefaultFunc, chunk_size, flags)
	case types.IS_STRING:
		fallthrough
	default:
		user := &PhpOutputHandlerUserFuncT{}
		if types.SUCCESS == zend.ZendFcallInfoInit(output_handler, 0, user.GetFci(), user.GetFcc(), &handler_name, &error) {
			handler = NewPhpOutputHandler(handler_name.GetStr(), chunk_size, flags & ^0xf | PHP_OUTPUT_HANDLER_USER)
			types.ZVAL_COPY(user.GetZoh(), output_handler)
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
func PhpOutputHandlerCreateInternal(name string, output_handler PhpOutputHandlerContextFuncT, chunk_size int, flags int) *PhpOutputHandler {
	var handler *PhpOutputHandler
	var str = types.NewString(name)
	handler = PhpOutputHandlerInit(str, chunk_size, flags & ^0xf | PHP_OUTPUT_HANDLER_INTERNAL)
	handler.SetInternal(output_handler)
	return handler
}
func PhpOutputHandlerSetContext(handler *PhpOutputHandler, opaq any) {
	handler.SetOpaq(opaq)
}
func PhpOutputHandlerStart(handler *PhpOutputHandler) int {
	if PhpOutputLockError(PHP_OUTPUT_HANDLER_START) != 0 || handler == nil {
		return types.FAILURE
	}

	/* zend_stack_push returns stack level */

	handler.SetLevel(OG__().Handlers().Push(&handler))
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

	if op != 0 && OG__().active && OG__().running {

		/* fatal error */

		PhpOutputDeactivate()
		PhpErrorDocref("ref.outcontrol", faults.E_ERROR, "Cannot use output buffering in output buffering display handlers")
		return 1
	}
	return 0
}
func PhpOutputContextFeed(context *PhpOutputContext, data *byte, size int, used int, free bool) {
	if context.GetIn().IsFree() && context.GetIn().GetData() != nil {
		zend.Efree(context.GetIn().GetData())
	}
	context.SetIn(NewOutputBuffer(data, size, used, free))
}
func PhpOutputContextSwap(context *PhpOutputContext) {
	if context.GetIn().IsFree() && context.GetIn().GetData() != nil {
		zend.Efree(context.GetIn().GetData())
	}
	context.SetIn(context.GetOut())
	context.SetOut(EmptyOutputBuffer())
}
func PhpOutputContextPass(context *PhpOutputContext) {
	context.SetOut(context.GetIn())
	context.SetIn(EmptyOutputBuffer())
}
func PhpOutputHandlerInit(name *types.String, chunkSize int, flags int) *PhpOutputHandler {
	return NewPhpOutputHandler(name.GetStr(), flags, chunkSize)
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
	var original_op int = context.GetOp()
	if PhpOutputLockError(context.GetOp()) != 0 {

		/* fatal error */

		return PHP_OUTPUT_HANDLER_FAILURE

		/* fatal error */

	}

	/* storable? */

	if PhpOutputHandlerAppend(handler, context.GetIn()) != 0 && context.GetOp() == 0 {
		context.SetOp(original_op)
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
			ob_data.SetStringVal(b.CastStr(handler.GetBuffer().GetData(), handler.GetBuffer().GetUsed()))
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
					if retval.String().GetLen() != 0 {
						context.GetOut().SetFreeDataByStr(retval.StringVal())
						status = PHP_OUTPUT_HANDLER_SUCCESS
					}
				}
			} else {
				/* call failed, pass internal buffer along */
				status = PHP_OUTPUT_HANDLER_FAILURE
			}
			zend.ZendFcallInfoArgn(handler.GetUser().GetFci(), 0)
		} else {
			PhpOutputContextFeed(context, handler.GetBuffer().GetData(), handler.GetBuffer().GetSize(), handler.GetBuffer().GetUsed(), 0)
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

		/* discard any output */

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
	context.SetOp(original_op)
	return status
}
func PhpOutputOp(op int, str *byte, len_ int) {
	var context PhpOutputContext
	var active **PhpOutputHandler
	var obh_cnt int
	if PhpOutputLockError(op) != 0 {
		return
	}
	context.Init(op)

	/*
	 * broken up for better performance:
	 *  - apply op to the one active handler; note that OG__().active might be popped off the stack on a flush
	 *  - or apply op to the handler stack
	 */

	if OG__().active && lang.Assign(&obh_cnt, OG__().Handlers().GetTop()) {
		context.GetIn().SetData((*byte)(str))
		context.GetIn().SetUsed(len_)
		if obh_cnt > 1 {
			zend.ZendStackApplyWithArgument(&(OG__().handlers), zend.ZEND_STACK_APPLY_TOPDOWN, PhpOutputStackApplyOp, &context)
		} else if lang.Assign(&active, OG__().Handlers().Top()) && !active.IsDisabled() {
			PhpOutputHandlerOp(*active, &context)
		} else {
			PhpOutputContextPass(&context)
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
			PhpOutputContextSwap(context)
		}
		return 0
	case PHP_OUTPUT_HANDLER_FAILURE:
		fallthrough
	default:
		if was_disabled != 0 {

			/* pass input along, if it's the last handler in the stack */

			if handler.GetLevel() == 0 {
				PhpOutputContextPass(context)
			}

			/* pass input along, if it's the last handler in the stack */

		} else {

			/* swap buffers, unless this is the last handler */

			if handler.GetLevel() != 0 {
				PhpOutputContextSwap(context)
			}

			/* swap buffers, unless this is the last handler */

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
func PhpOutputStackApplyStatus(h any, z any) int {
	var handler *PhpOutputHandler = *((**PhpOutputHandler)(h))
	var arr types.Zval
	var array *types.Zval = (*types.Zval)(z)
	zend.AddNextIndexZval(array, PhpOutputHandlerStatus(handler, &arr))
	return 0
}
func PhpOutputHandlerStatus(handler *PhpOutputHandler, entry *types.Zval) *types.Zval {
	b.Assert(entry != nil)
	zend.ArrayInit(entry)
	entry.Array().KeyUpdate("name", types.NewZvalString(handler.GetName()))
	entry.Array().KeyUpdate("name", types.NewZvalLong(handler.GetFlags()&0xf))
	entry.Array().KeyUpdate("flags", types.NewZvalLong(handler.GetFlags()))
	entry.Array().KeyUpdate("level", types.NewZvalLong(handler.GetLevel()))
	entry.Array().KeyUpdate("chunk_size", types.NewZvalLong(handler.GetSize()))
	entry.Array().KeyUpdate("buffer_size", types.NewZvalLong(handler.GetBuffer().GetSize()))
	entry.Array().KeyUpdate("buffer_used", types.NewZvalLong(handler.GetBuffer().GetUsed()))
	return entry
}
func PhpOutputStackPop(flags int) int {
	var context PhpOutputContext
	var current **PhpOutputHandler
	var orphan **PhpOutputHandler = OG__().active
	if orphan == nil {
		if (flags & PHP_OUTPUT_POP_SILENT) == 0 {
			PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to %s buffer. No buffer to %s", lang.Cond((flags&PHP_OUTPUT_POP_DISCARD) != 0, "discard", "send"), lang.Cond((flags&PHP_OUTPUT_POP_DISCARD) != 0, "discard", "send"))
		}
		return 0
	} else if (flags&PHP_OUTPUT_POP_FORCE) == 0 && !orphan.HasFlags(PHP_OUTPUT_HANDLER_REMOVABLE) {
		if (flags & PHP_OUTPUT_POP_SILENT) == 0 {
			PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to %s buffer of %s (%d)", lang.Cond((flags&PHP_OUTPUT_POP_DISCARD) != 0, "discard", "send"), orphan.Name(), orphan.GetLevel())
		}
		return 0
	} else {
		context.Init(PHP_OUTPUT_HANDLER_FINAL)

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
			PhpOutputHandlerOp(orphan, &context)
		}

		/* pop it off the stack */

		OG__().Handlers().DelTop()
		if lang.Assign(&current, OG__().Handlers().Top()) {
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
	var func_ PhpOutputHandlerFuncT = *((*PhpOutputHandlerFuncT)(handler_context))
	if func_ != nil {
		var out_str *byte = nil
		var out_len int = 0
		func_(output_context.GetIn().GetData(), output_context.GetIn().GetUsed(), &out_str, &out_len, output_context.GetOp())
		if out_str != nil {
			output_context.GetOut().SetData(out_str)
			output_context.GetOut().SetUsed(out_len)
			output_context.GetOut().SetFree(true)
		} else {
			PhpOutputContextPass(output_context)
		}
		return types.SUCCESS
	}
	return types.FAILURE
}
func PhpOutputHandlerDefaultFunc(handler_context *any, output_context *PhpOutputContext) int {
	PhpOutputContextPass(output_context)
	return types.SUCCESS
}
func ZifObStart(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, userFunction *types.Zval, chunkSize *types.Zval, flags *types.Zval) {
	var output_handler *types.Zval = nil
	var chunk_size zend.ZendLong = 0
	var flags zend.ZendLong = PHP_OUTPUT_HANDLER_STDFLAGS
	if zend.ZendParseParameters(executeData.NumArgs(), "|zll", &output_handler, &chunk_size, &flags) == types.FAILURE {
		return
	}
	if chunk_size < 0 {
		chunk_size = 0
	}
	if PhpOutputStartUser(output_handler, chunk_size, flags) == types.FAILURE {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to create buffer")
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
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
func ZifObGetStatus(return_value zpp.Ret, _ zpp.Opt, fullStatus bool) {
	if OG__().active == nil {
		zend.ArrayInit(return_value)
		return
	}
	if fullStatus {
		zend.ArrayInit(return_value)
		zend.ZendStackApplyWithArgument(&(OG__().handlers), zend.ZEND_STACK_APPLY_BOTTOMUP, PhpOutputStackApplyStatus, return_value)
	} else {
		PhpOutputHandlerStatus(OG__().active, return_value)
	}
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
