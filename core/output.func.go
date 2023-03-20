// <<generate>>

package core

import (
	b "sik/builtin"
	r "sik/builtin/file"
	"sik/ext/standard"
	"sik/zend"
	"sik/zend/faults"
	"sik/zend/types"
)

func PhpOutputInitGlobals(G *ZendOutputGlobals) { memset(G, 0, b.SizeOf("* G")) }
func PhpOutputStdout(str *byte, str_len int) int {
	r.Fwrite(str, 1, str_len, stdout)
	return str_len
}
func PhpOutputStderr(str *byte, str_len int) int {
	r.Fwrite(str, 1, str_len, stderr)

	/* See http://support.microsoft.com/kb/190351 */

	return str_len

	/* See http://support.microsoft.com/kb/190351 */
}
func PhpOutputHeader() {
	if !(SG__().headers_sent) {
		if !(OG(output_start_filename)) {
			if zend.ZendIsCompiling() != 0 {
				OG(output_start_filename) = zend.ZendGetCompiledFilename().GetVal()
				OG(output_start_lineno) = zend.ZendGetCompiledLineno()
			} else if zend.ZendIsExecuting() != 0 {
				OG(output_start_filename) = zend.ZendGetExecutedFilename()
				OG(output_start_lineno) = zend.ZendGetExecutedLineno()
			}
		}
		if standard.PhpHeader() == 0 {
			OG(flags) |= PHP_OUTPUT_DISABLED
		}
	}
}
func ReverseConflictDtor(zv *types.Zval) {
	var ht *types.HashTable = zv.GetPtr()
	ht.Destroy()
}
func PhpOutputStartup() {
	PhpOutputInitGlobals(&OutputGlobals)
	types.ZendHashInit(&PhpOutputHandlerAliases, 8, nil, nil, 1)
	types.ZendHashInit(&PhpOutputHandlerConflicts, 8, nil, nil, 1)
	types.ZendHashInit(&PhpOutputHandlerReverseConflicts, 8, nil, ReverseConflictDtor, 1)
	PhpOutputDirect = PhpOutputStdout
}
func PhpOutputShutdown() {
	PhpOutputDirect = PhpOutputStderr
	PhpOutputHandlerAliases.Destroy()
	PhpOutputHandlerConflicts.Destroy()
	PhpOutputHandlerReverseConflicts.Destroy()
}
func PhpOutputActivate() int {
	memset(&OutputGlobals, 0, b.SizeOf("zend_output_globals"))
	OG(handlers).Init()
	OG(flags) |= PHP_OUTPUT_ACTIVATED
	return types.SUCCESS
}
func PhpOutputDeactivate() {
	var handler **PhpOutputHandler = nil
	if (OG(flags) & PHP_OUTPUT_ACTIVATED) != 0 {
		PhpOutputHeader()
		OG(flags) ^= PHP_OUTPUT_ACTIVATED
		OG(active) = nil
		OG(running) = nil

		/* release all output handlers */

		if OG(handlers).elements {
			for b.Assign(&handler, OG(handlers).Top()) {
				PhpOutputHandlerFree(handler)
				OG(handlers).DelTop()
			}
		}
		OG(handlers).Destroy()
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
func PhpOutputSetStatus(status int) {
	OG(flags) = OG(flags) & ^0xf | status&0xf
}
func PhpOutputGetStatus() int {
	return (OG(flags) | b.Cond(OG(active), PHP_OUTPUT_ACTIVE, 0) | b.Cond(OG(running), PHP_OUTPUT_LOCKED, 0)) & 0xff
}
func PhpOutputWriteUnbuffered(str *byte, len_ int) int {
	if (OG(flags) & PHP_OUTPUT_ACTIVATED) != 0 {
		s := b.CastStr(str, len_)
		return SM__().UbWrite(s)
	}
	return PhpOutputDirect(str, len_)
}
func PhpOutputWrite(str string) int {
	ptr := b.CastStrPtr(str)
	len_ := len(str)

	if (OG(flags) & PHP_OUTPUT_ACTIVATED) != 0 {
		PhpOutputOp(PHP_OUTPUT_HANDLER_WRITE, ptr, len_)
		return len_
	}
	if (OG(flags) & PHP_OUTPUT_DISABLED) != 0 {
		return 0
	}
	return PhpOutputDirect(str, len_)
}
func PhpOutputFlush() int {
	var context PhpOutputContext
	if OG(active) && (OG(active).flags&PHP_OUTPUT_HANDLER_FLUSHABLE) != 0 {
		PhpOutputContextInit(&context, PHP_OUTPUT_HANDLER_FLUSH)
		PhpOutputHandlerOp(OG(active), &context)
		if context.GetOut().GetData() != nil && context.GetOut().GetUsed() != 0 {
			OG(handlers).DelTop()
			PhpOutputWrite(context.GetOut().GetData(), context.GetOut().GetUsed())
			OG(handlers).Push(&(OG(active)))
		}
		PhpOutputContextDtor(&context)
		return types.SUCCESS
	}
	return types.FAILURE
}
func PhpOutputFlushAll() {
	if OG(active) {
		PhpOutputOp(PHP_OUTPUT_HANDLER_FLUSH, nil, 0)
	}
}
func PhpOutputClean() int {
	var context PhpOutputContext
	if OG(active) && (OG(active).flags&PHP_OUTPUT_HANDLER_CLEANABLE) != 0 {
		PhpOutputContextInit(&context, PHP_OUTPUT_HANDLER_CLEAN)
		PhpOutputHandlerOp(OG(active), &context)
		PhpOutputContextDtor(&context)
		return types.SUCCESS
	}
	return types.FAILURE
}
func PhpOutputCleanAll() {
	var context PhpOutputContext
	if OG(active) {
		PhpOutputContextInit(&context, PHP_OUTPUT_HANDLER_CLEAN)
		zend.ZendStackApplyWithArgument(&(OG(handlers)), zend.ZEND_STACK_APPLY_TOPDOWN, PhpOutputStackApplyClean, &context)
	}
}
func PhpOutputEnd() int {
	if PhpOutputStackPop(PHP_OUTPUT_POP_TRY) != 0 {
		return types.SUCCESS
	}
	return types.FAILURE
}
func PhpOutputEndAll() {
	for OG(active) && PhpOutputStackPop(PHP_OUTPUT_POP_FORCE) != 0 {

	}
}
func PhpOutputDiscard() int {
	if PhpOutputStackPop(PHP_OUTPUT_POP_DISCARD|PHP_OUTPUT_POP_TRY) != 0 {
		return types.SUCCESS
	}
	return types.FAILURE
}
func PhpOutputDiscardAll() {
	for OG(active) {
		PhpOutputStackPop(PHP_OUTPUT_POP_DISCARD | PHP_OUTPUT_POP_FORCE)
	}
}
func PhpOutputGetLevel() int {
	if OG(active) {
		return OG(handlers).GetTop()
	} else {
		return 0
	}
}
func PhpOutputGetContents(p *types.Zval) int {
	if OG(active) {
		p.SetRawString(b.CastStr(OG(active).buffer.data, OG(active).buffer.used))
		return types.SUCCESS
	} else {
		p.SetNull()
		return types.FAILURE
	}
}
func PhpOutputGetLength(p *types.Zval) int {
	if OG(active) {
		p.SetLong(OG(active).buffer.used)
		return types.SUCCESS
	} else {
		p.SetNull()
		return types.FAILURE
	}
}
func PhpOutputGetActiveHandler() *PhpOutputHandler { return OG(active) }
func PhpOutputStartDefault() int {
	var handler *PhpOutputHandler
	handler = PhpOutputHandlerCreateInternal(zend.ZEND_STRL(PhpOutputDefaultHandlerName), PhpOutputHandlerDefaultFunc, 0, PHP_OUTPUT_HANDLER_STDFLAGS)
	if types.SUCCESS == PhpOutputHandlerStart(handler) {
		return types.SUCCESS
	}
	PhpOutputHandlerFree(&handler)
	return types.FAILURE
}
func PhpOutputStartDevnull() int {
	var handler *PhpOutputHandler
	handler = PhpOutputHandlerCreateInternal(zend.ZEND_STRL(PhpOutputDevnullHandlerName), PhpOutputHandlerDevnullFunc, PHP_OUTPUT_HANDLER_DEFAULT_SIZE, 0)
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
		handler = PhpOutputHandlerCreateInternal(zend.ZEND_STRL(PhpOutputDefaultHandlerName), PhpOutputHandlerDefaultFunc, chunk_size, flags)
	}
	if types.SUCCESS == PhpOutputHandlerStart(handler) {
		return types.SUCCESS
	}
	PhpOutputHandlerFree(&handler)
	return types.FAILURE
}
func PhpOutputStartInternal(name *byte, name_len int, output_handler PhpOutputHandlerFuncT, chunk_size int, flags int) int {
	var handler *PhpOutputHandler
	handler = PhpOutputHandlerCreateInternal(name, name_len, PhpOutputHandlerCompatFunc, chunk_size, flags)
	PhpOutputHandlerSetContext(handler, output_handler, nil)
	if types.SUCCESS == PhpOutputHandlerStart(handler) {
		return types.SUCCESS
	}
	PhpOutputHandlerFree(&handler)
	return types.FAILURE
}
func PhpOutputHandlerCreateUser(output_handler *types.Zval, chunk_size int, flags int) *PhpOutputHandler {
	var handler_name *types.String = nil
	var error *byte = nil
	var handler *PhpOutputHandler = nil
	var alias PhpOutputHandlerAliasCtorT = nil
	var user *PhpOutputHandlerUserFuncT = nil
	switch output_handler.GetType() {
	case types.IS_NULL:
		handler = PhpOutputHandlerCreateInternal(zend.ZEND_STRL(PhpOutputDefaultHandlerName), PhpOutputHandlerDefaultFunc, chunk_size, flags)
	case types.IS_STRING:
		if output_handler.GetStr().GetLen() != 0 && b.Assign(&alias, PhpOutputHandlerAlias(output_handler.GetStr().GetVal(), output_handler.GetStr().GetLen())) {
			handler = alias(output_handler.GetStr().GetVal(), output_handler.GetStr().GetLen(), chunk_size, flags)
			break
		}
		fallthrough
	default:
		user = zend.Ecalloc(1, b.SizeOf("php_output_handler_user_func_t"))
		if types.SUCCESS == zend.ZendFcallInfoInit(output_handler, 0, user.GetFci(), user.GetFcc(), &handler_name, &error) {
			handler = PhpOutputHandlerInit(handler_name, chunk_size, flags & ^0xf | PHP_OUTPUT_HANDLER_USER)
			types.ZVAL_COPY(user.GetZoh(), output_handler)
			handler.SetUser(user)
		} else {
			zend.Efree(user)
		}
		if error != nil {
			PhpErrorDocref("ref.outcontrol", faults.E_WARNING, "%s", error)
			zend.Efree(error)
		}
		if handler_name != nil {
			types.ZendStringReleaseEx(handler_name, 0)
		}
	}
	return handler
}
func PhpOutputHandlerCreateInternal(name *byte, name_len int, output_handler PhpOutputHandlerContextFuncT, chunk_size int, flags int) *PhpOutputHandler {
	var handler *PhpOutputHandler
	var str *types.String = types.NewString(b.CastStr(name, name_len))
	handler = PhpOutputHandlerInit(str, chunk_size, flags & ^0xf | PHP_OUTPUT_HANDLER_INTERNAL)
	handler.SetInternal(output_handler)
	types.ZendStringReleaseEx(str, 0)
	return handler
}
func PhpOutputHandlerSetContext(handler *PhpOutputHandler, opaq any, dtor func(any)) {
	if handler.GetDtor() != nil && handler.GetOpaq() {
		handler.GetDtor()(handler.GetOpaq())
	}
	handler.SetDtor(dtor)
	handler.SetOpaq(opaq)
}
func PhpOutputHandlerStart(handler *PhpOutputHandler) int {
	var rconflicts *types.HashTable
	var conflict PhpOutputHandlerConflictCheckT
	if PhpOutputLockError(PHP_OUTPUT_HANDLER_START) != 0 || handler == nil {
		return types.FAILURE
	}
	if nil != b.Assign(&conflict, types.ZendHashFindPtr(&PhpOutputHandlerConflicts, handler.GetName())) {
		if types.SUCCESS != conflict(handler.GetName().GetVal(), handler.GetName().GetLen()) {
			return types.FAILURE
		}
	}
	if nil != b.Assign(&rconflicts, types.ZendHashFindPtr(&PhpOutputHandlerReverseConflicts, handler.GetName())) {
		var __ht *types.HashTable = rconflicts
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()

			conflict = _z.GetPtr()
			if types.SUCCESS != conflict(handler.GetName().GetVal(), handler.GetName().GetLen()) {
				return types.FAILURE
			}
		}
	}

	/* zend_stack_push returns stack level */

	handler.SetLevel(OG(handlers).Push(&handler))
	OG(active) = handler
	return types.SUCCESS
}
func PhpOutputHandlerStarted(name *byte, name_len int) int {
	var handlers **PhpOutputHandler
	var i int
	var count int = PhpOutputGetLevel()
	if count != 0 {
		handlers = (**PhpOutputHandler)(OG(handlers).GetElements())
		for i = 0; i < count; i++ {
			if name_len == handlers[i].GetName().GetLen() && !(memcmp(handlers[i].GetName().GetVal(), name, name_len)) {
				return 1
			}
		}
	}
	return 0
}
func PhpOutputHandlerConflict(handler_new *byte, handler_new_len int, handler_set *byte, handler_set_len int) int {
	if PhpOutputHandlerStarted(handler_set, handler_set_len) != 0 {
		if handler_new_len != handler_set_len || memcmp(handler_new, handler_set, handler_set_len) {
			PhpErrorDocref("ref.outcontrol", faults.E_WARNING, "output handler '%s' conflicts with '%s'", handler_new, handler_set)
		} else {
			PhpErrorDocref("ref.outcontrol", faults.E_WARNING, "output handler '%s' cannot be used twice", handler_new)
		}
		return 1
	}
	return 0
}
func PhpOutputHandlerConflictRegister(name *byte, name_len int, check_func PhpOutputHandlerConflictCheckT) int {
	var str *types.String
	if zend.EG__().GetCurrentModule() == nil {
		faults.Error(faults.E_ERROR, "Cannot register an output handler conflict outside of MINIT")
		return types.FAILURE
	}
	str = types.ZendStringInitInterned(name, name_len, 1)
	types.ZendHashUpdatePtr(&PhpOutputHandlerConflicts, str, check_func)
	types.ZendStringReleaseEx(str, 1)
	return types.SUCCESS
}
func PhpOutputHandlerReverseConflictRegister(name *byte, name_len int, check_func PhpOutputHandlerConflictCheckT) int {
	var rev types.HashTable
	var rev_ptr *types.HashTable = nil
	if zend.EG__().GetCurrentModule() == nil {
		faults.Error(faults.E_ERROR, "Cannot register a reverse output handler conflict outside of MINIT")
		return types.FAILURE
	}
	if nil != b.Assign(&rev_ptr, types.ZendHashStrFindPtr(&PhpOutputHandlerReverseConflicts, name, name_len)) {
		if types.ZendHashNextIndexInsertPtr(rev_ptr, check_func) {
			return types.SUCCESS
		} else {
			return types.FAILURE
		}
	} else {
		var str *types.String
		types.ZendHashInit(&rev, 8, nil, nil, 1)
		if nil == types.ZendHashNextIndexInsertPtr(&rev, check_func) {
			rev.Destroy()
			return types.FAILURE
		}
		str = types.ZendStringInitInterned(name, name_len, 1)
		types.ZendHashUpdateMem(&PhpOutputHandlerReverseConflicts, str, &rev, b.SizeOf("HashTable"))
		types.ZendStringReleaseEx(str, 1)
		return types.SUCCESS
	}
}
func PhpOutputHandlerAlias(name *byte, name_len int) PhpOutputHandlerAliasCtorT {
	return types.ZendHashStrFindPtr(&PhpOutputHandlerAliases, name, name_len)
}
func PhpOutputHandlerAliasRegister(name *byte, name_len int, func_ PhpOutputHandlerAliasCtorT) int {
	var str *types.String
	if zend.EG__().GetCurrentModule() == nil {
		faults.Error(faults.E_ERROR, "Cannot register an output handler alias outside of MINIT")
		return types.FAILURE
	}
	str = types.ZendStringInitInterned(name, name_len, 1)
	types.ZendHashUpdatePtr(&PhpOutputHandlerAliases, str, func_)
	types.ZendStringReleaseEx(str, 1)
	return types.SUCCESS
}
func PhpOutputHandlerHook(type_ PhpOutputHandlerHookT, arg any) int {
	if OG(running) {
		switch type_ {
		case PHP_OUTPUT_HANDLER_HOOK_GET_OPAQ:
			*((**any)(arg)) = OG(running).opaq
			return types.SUCCESS
		case PHP_OUTPUT_HANDLER_HOOK_GET_FLAGS:
			*((*int)(arg)) = OG(running).flags
			return types.SUCCESS
		case PHP_OUTPUT_HANDLER_HOOK_GET_LEVEL:
			*((*int)(arg)) = OG(running).level
			return types.SUCCESS
		case PHP_OUTPUT_HANDLER_HOOK_IMMUTABLE:
			OG(running).flags &= ^(PHP_OUTPUT_HANDLER_REMOVABLE | PHP_OUTPUT_HANDLER_CLEANABLE)
			return types.SUCCESS
		case PHP_OUTPUT_HANDLER_HOOK_DISABLE:
			OG(running).flags |= PHP_OUTPUT_HANDLER_DISABLED
			return types.SUCCESS
		default:

		}
	}
	return types.FAILURE
}
func PhpOutputHandlerDtor(handler *PhpOutputHandler) {
	if handler.GetName() != nil {
		types.ZendStringReleaseEx(handler.GetName(), 0)
	}
	if handler.GetBuffer().GetData() != nil {
		zend.Efree(handler.GetBuffer().GetData())
	}
	if handler.IsUser() {
		zend.ZvalPtrDtor(handler.GetUser().GetZoh())
		zend.Efree(handler.GetUser())
	}
	if handler.GetDtor() != nil && handler.GetOpaq() {
		handler.GetDtor()(handler.GetOpaq())
	}
	memset(handler, 0, b.SizeOf("* handler"))
}
func PhpOutputHandlerFree(h **PhpOutputHandler) {
	if (*h) != nil {
		PhpOutputHandlerDtor(*h)
		zend.Efree(*h)
		*h = nil
	}
}
func PhpOutputSetImplicitFlush(flush int) {
	if flush != 0 {
		OG(flags) |= PHP_OUTPUT_IMPLICITFLUSH
	} else {
		OG(flags) &= ^PHP_OUTPUT_IMPLICITFLUSH
	}
}
func PhpOutputGetStartFilename() *byte { return OG(output_start_filename) }
func PhpOutputGetStartLineno() int     { return OG(output_start_lineno) }
func PhpOutputLockError(op int) int {
	/* if there's no ob active, ob has been stopped */

	if op != 0 && OG(active) && OG(running) {

		/* fatal error */

		PhpOutputDeactivate()
		PhpErrorDocref("ref.outcontrol", faults.E_ERROR, "Cannot use output buffering in output buffering display handlers")
		return 1
	}
	return 0
}
func PhpOutputContextInit(context *PhpOutputContext, op int) {
	memset(context, 0, b.SizeOf("php_output_context"))
	context.SetOp(op)
}
func PhpOutputContextReset(context *PhpOutputContext) {
	var op int = context.GetOp()
	PhpOutputContextDtor(context)
	memset(context, 0, b.SizeOf("php_output_context"))
	context.SetOp(op)
}
func PhpOutputContextFeed(context *PhpOutputContext, data *byte, size int, used int, free types.ZendBool) {
	if context.GetIn().GetFree() != 0 && context.GetIn().GetData() != nil {
		zend.Efree(context.GetIn().GetData())
	}
	context.GetIn().SetData(data)
	context.GetIn().SetUsed(used)
	context.GetIn().SetFree(free)
	context.GetIn().SetSize(size)
}
func PhpOutputContextSwap(context *PhpOutputContext) {
	if context.GetIn().GetFree() != 0 && context.GetIn().GetData() != nil {
		zend.Efree(context.GetIn().GetData())
	}
	context.GetIn().SetData(context.GetOut().GetData())
	context.GetIn().SetUsed(context.GetOut().GetUsed())
	context.GetIn().SetFree(context.GetOut().GetFree())
	context.GetIn().SetSize(context.GetOut().GetSize())
	context.GetOut().SetData(nil)
	context.GetOut().SetUsed(0)
	context.GetOut().SetFree(0)
	context.GetOut().SetSize(0)
}
func PhpOutputContextPass(context *PhpOutputContext) {
	context.GetOut().SetData(context.GetIn().GetData())
	context.GetOut().SetUsed(context.GetIn().GetUsed())
	context.GetOut().SetSize(context.GetIn().GetSize())
	context.GetOut().SetFree(context.GetIn().GetFree())
	context.GetIn().SetData(nil)
	context.GetIn().SetUsed(0)
	context.GetIn().SetFree(0)
	context.GetIn().SetSize(0)
}
func PhpOutputContextDtor(context *PhpOutputContext) {
	if context.GetIn().GetFree() != 0 && context.GetIn().GetData() != nil {
		zend.Efree(context.GetIn().GetData())
		context.GetIn().SetData(nil)
	}
	if context.GetOut().GetFree() != 0 && context.GetOut().GetData() != nil {
		zend.Efree(context.GetOut().GetData())
		context.GetOut().SetData(nil)
	}
}
func PhpOutputHandlerInit(name *types.String, chunk_size int, flags int) *PhpOutputHandler {
	var handler *PhpOutputHandler
	handler = zend.Ecalloc(1, b.SizeOf("php_output_handler"))
	handler.SetName(name.Copy())
	handler.SetSize(chunk_size)
	handler.SetFlags(flags)
	handler.GetBuffer().SetSize(PHP_OUTPUT_HANDLER_INITBUF_SIZE(chunk_size))
	handler.GetBuffer().SetData(zend.Emalloc(handler.GetBuffer().GetSize()))
	return handler
}
func PhpOutputHandlerAppend(handler *PhpOutputHandler, buf *PhpOutputBuffer) int {
	if buf.GetUsed() != 0 {
		OG(flags) |= PHP_OUTPUT_WRITTEN

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

			if OG(running) {
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
		OG(running) = handler
		if handler.IsUser() {
			var retval types.Zval
			var ob_data types.Zval
			var ob_mode types.Zval
			ob_data.SetRawString(b.CastStr(handler.GetBuffer().GetData(), handler.GetBuffer().GetUsed()))
			ob_mode.SetLong(zend.ZendLong(context.GetOp()))
			zend.ZendFcallInfoArgn(handler.GetUser().GetFci(), 2, &ob_data, &ob_mode)
			zend.ZvalPtrDtor(&ob_data)
			var PHP_OUTPUT_USER_SUCCESS func(retval types.Zval) bool = func(retval types.Zval) bool {
				return retval.GetType() != types.IS_UNDEF && retval.GetType() != types.IS_FALSE
			}
			if types.SUCCESS == zend.ZendFcallInfoCall(handler.GetUser().GetFci(), handler.GetUser().GetFcc(), &retval, nil) && PHP_OUTPUT_USER_SUCCESS(retval) {

				/* user handler may have returned TRUE */

				status = PHP_OUTPUT_HANDLER_NO_DATA
				if retval.GetType() != types.IS_FALSE && retval.GetType() != types.IS_TRUE {
					zend.ConvertToStringEx(&retval)
					if retval.GetStr().GetLen() != 0 {
						context.GetOut().SetData(zend.Estrndup(retval.GetStr().GetVal(), retval.GetStr().GetLen()))
						context.GetOut().SetUsed(retval.GetStr().GetLen())
						context.GetOut().SetFree(1)
						status = PHP_OUTPUT_HANDLER_SUCCESS
					}
				}
			} else {

				/* call failed, pass internal buffer along */

				status = PHP_OUTPUT_HANDLER_FAILURE

				/* call failed, pass internal buffer along */

			}
			zend.ZendFcallInfoArgn(handler.GetUser().GetFci(), 0)
			zend.ZvalPtrDtor(&retval)
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
		OG(running) = nil
	}
	switch status {
	case PHP_OUTPUT_HANDLER_FAILURE:

		/* disable this handler */

		handler.SetIsDisabled(true)

		/* discard any output */

		if context.GetOut().GetData() != nil && context.GetOut().GetFree() != 0 {
			zend.Efree(context.GetOut().GetData())
		}

		/* returns handlers buffer */

		context.GetOut().SetData(handler.GetBuffer().GetData())
		context.GetOut().SetUsed(handler.GetBuffer().GetUsed())
		context.GetOut().SetFree(1)
		handler.GetBuffer().SetData(nil)
		handler.GetBuffer().SetUsed(0)
		handler.GetBuffer().SetSize(0)
	case PHP_OUTPUT_HANDLER_NO_DATA:

		/* handler ate all */

		PhpOutputContextReset(context)
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
	PhpOutputContextInit(&context, op)

	/*
	 * broken up for better performance:
	 *  - apply op to the one active handler; note that OG(active) might be popped off the stack on a flush
	 *  - or apply op to the handler stack
	 */

	if OG(active) && b.Assign(&obh_cnt, OG(handlers).GetTop()) {
		context.GetIn().SetData((*byte)(str))
		context.GetIn().SetUsed(len_)
		if obh_cnt > 1 {
			zend.ZendStackApplyWithArgument(&(OG(handlers)), zend.ZEND_STACK_APPLY_TOPDOWN, PhpOutputStackApplyOp, &context)
		} else if b.Assign(&active, OG(handlers).Top()) && !active.IsDisabled() {
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
		if (OG(flags) & PHP_OUTPUT_DISABLED) == 0 {
			SM__().UbWrite(
				b.CastStr(context.GetOut().GetData(), context.GetOut().GetUsed()),
			)
			if (OG(flags) & PHP_OUTPUT_IMPLICITFLUSH) != 0 {
				SapiFlush()
			}
			OG(flags) |= PHP_OUTPUT_SENT
		}
	}
	PhpOutputContextDtor(&context)
}
func PhpOutputStackApplyOp(h any, c any) int {
	var was_disabled int
	var status PhpOutputHandlerStatusT
	var handler *PhpOutputHandler = *((**PhpOutputHandler)(h))
	var context *PhpOutputContext = (*PhpOutputContext)(c)
	if b.Assign(&was_disabled, handler.GetFlags()&PHP_OUTPUT_HANDLER_DISABLED) {
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
func PhpOutputStackApplyClean(h any, c any) int {
	var handler *PhpOutputHandler = *((**PhpOutputHandler)(h))
	var context *PhpOutputContext = (*PhpOutputContext)(c)
	handler.GetBuffer().SetUsed(0)
	PhpOutputHandlerOp(handler, context)
	PhpOutputContextReset(context)
	return 0
}
func PhpOutputStackApplyList(h any, z any) int {
	var handler *PhpOutputHandler = *((**PhpOutputHandler)(h))
	var array *types.Zval = (*types.Zval)(z)
	zend.AddNextIndexStr(array, handler.GetName().Copy())
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
	zend.AddAssocStr(entry, "name", handler.GetName().GetStr())
	zend.AddAssocLong(entry, "type", zend_long(handler.GetFlags()&0xf))
	zend.AddAssocLong(entry, "flags", zend.ZendLong(handler.GetFlags()))
	zend.AddAssocLong(entry, "level", zend.ZendLong(handler.GetLevel()))
	zend.AddAssocLong(entry, "chunk_size", zend.ZendLong(handler.GetSize()))
	zend.AddAssocLong(entry, "buffer_size", zend.ZendLong(handler.GetBuffer().GetSize()))
	zend.AddAssocLong(entry, "buffer_used", zend.ZendLong(handler.GetBuffer().GetUsed()))
	return entry
}
func PhpOutputStackPop(flags int) int {
	var context PhpOutputContext
	var current **PhpOutputHandler
	var orphan **PhpOutputHandler = OG(active)
	if orphan == nil {
		if (flags & PHP_OUTPUT_POP_SILENT) == 0 {
			PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to %s buffer. No buffer to %s", b.Cond((flags&PHP_OUTPUT_POP_DISCARD) != 0, "discard", "send"), b.Cond((flags&PHP_OUTPUT_POP_DISCARD) != 0, "discard", "send"))
		}
		return 0
	} else if (flags&PHP_OUTPUT_POP_FORCE) == 0 && !orphan.HasFlags(PHP_OUTPUT_HANDLER_REMOVABLE) {
		if (flags & PHP_OUTPUT_POP_SILENT) == 0 {
			PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to %s buffer of %s (%d)", b.Cond((flags&PHP_OUTPUT_POP_DISCARD) != 0, "discard", "send"), orphan.GetName().GetVal(), orphan.GetLevel())
		}
		return 0
	} else {
		PhpOutputContextInit(&context, PHP_OUTPUT_HANDLER_FINAL)

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

		OG(handlers).DelTop()
		if b.Assign(&current, OG(handlers).Top()) {
			OG(active) = *current
		} else {
			OG(active) = nil
		}

		/* pass output along */

		if context.GetOut().GetData() != nil && context.GetOut().GetUsed() != 0 && (flags&PHP_OUTPUT_POP_DISCARD) == 0 {
			PhpOutputWrite(context.GetOut().GetData(), context.GetOut().GetUsed())
		}

		/* destroy the handler (after write!) */

		PhpOutputHandlerFree(&orphan)
		PhpOutputContextDtor(&context)
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
			output_context.GetOut().SetFree(1)
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
func PhpOutputHandlerDevnullFunc(handler_context *any, output_context *PhpOutputContext) int {
	return types.SUCCESS
}
func ZifObStart(executeData *zend.ZendExecuteData, return_value *types.Zval) {
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
func ZifObFlush(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if !(OG(active)) {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to flush buffer. No buffer to flush")
		return_value.SetFalse()
		return
	}
	if types.SUCCESS != PhpOutputFlush() {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to flush buffer of %s (%d)", OG(active).name.GetVal(), OG(active).level)
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
}
func ZifObClean(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if !(OG(active)) {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to delete buffer. No buffer to delete")
		return_value.SetFalse()
		return
	}
	if types.SUCCESS != PhpOutputClean() {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to delete buffer of %s (%d)", OG(active).name.GetVal(), OG(active).level)
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
}
func ZifObEndFlush(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if !(OG(active)) {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to delete and flush buffer. No buffer to delete or flush")
		return_value.SetFalse()
		return
	}
	types.ZVAL_BOOL(return_value, types.SUCCESS == PhpOutputEnd())
	return
}
func ZifObEndClean(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if !(OG(active)) {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to delete buffer. No buffer to delete")
		return_value.SetFalse()
		return
	}
	types.ZVAL_BOOL(return_value, types.SUCCESS == PhpOutputDiscard())
	return
}
func ZifObGetFlush(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if PhpOutputGetContents(return_value) == types.FAILURE {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to delete and flush buffer. No buffer to delete or flush")
		return_value.SetFalse()
		return
	}
	if types.SUCCESS != PhpOutputEnd() {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to delete buffer of %s (%d)", OG(active).name.GetVal(), OG(active).level)
	}
}
func ZifObGetClean(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if !(OG(active)) {
		return_value.SetFalse()
		return
	}
	if PhpOutputGetContents(return_value) == types.FAILURE {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to delete buffer. No buffer to delete")
		return_value.SetFalse()
		return
	}
	if types.SUCCESS != PhpOutputDiscard() {
		PhpErrorDocref("ref.outcontrol", faults.E_NOTICE, "failed to delete buffer of %s (%d)", OG(active).name.GetVal(), OG(active).level)
	}
}
func ZifObGetContents(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if PhpOutputGetContents(return_value) == types.FAILURE {
		return_value.SetFalse()
		return
	}
}
func ZifObGetLevel(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	return_value.SetLong(PhpOutputGetLevel())
	return
}
func ZifObGetLength(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if PhpOutputGetLength(return_value) == types.FAILURE {
		return_value.SetFalse()
		return
	}
}
func ZifObListHandlers(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	zend.ArrayInit(return_value)
	if !(OG(active)) {
		return
	}
	zend.ZendStackApplyWithArgument(&(OG(handlers)), zend.ZEND_STACK_APPLY_BOTTOMUP, PhpOutputStackApplyList, return_value)
}
func ZifObGetStatus(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var full_status types.ZendBool = 0
	if zend.ZendParseParameters(executeData.NumArgs(), "|b", &full_status) == types.FAILURE {
		return
	}
	if !(OG(active)) {
		zend.ArrayInit(return_value)
		return
	}
	if full_status != 0 {
		zend.ArrayInit(return_value)
		zend.ZendStackApplyWithArgument(&(OG(handlers)), zend.ZEND_STACK_APPLY_BOTTOMUP, PhpOutputStackApplyStatus, return_value)
	} else {
		PhpOutputHandlerStatus(OG(active), return_value)
	}
}
func ZifObImplicitFlush(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var flag zend.ZendLong = 1
	if zend.ZendParseParameters(executeData.NumArgs(), "|l", &flag) == types.FAILURE {
		return
	}
	PhpOutputSetImplicitFlush(flag)
}
func ZifOutputResetRewriteVars(executeData *zend.ZendExecuteData, return_value *types.Zval) {
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
func ZifOutputAddRewriteVar(executeData *zend.ZendExecuteData, return_value *types.Zval) {
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
