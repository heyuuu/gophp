// <<generate>>

package core

import (
	b "sik/builtin"
	"sik/ext/standard"
	r "sik/runtime"
	"sik/zend"
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
	if !(SG(headers_sent)) {
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
func ReverseConflictDtor(zv *zend.Zval) {
	var ht *zend.HashTable = zv.GetPtr()
	zend.ZendHashDestroy(ht)
}
func PhpOutputStartup() {
	PhpOutputInitGlobals(&OutputGlobals)
	zend.ZendHashInit(&PhpOutputHandlerAliases, 8, nil, nil, 1)
	zend.ZendHashInit(&PhpOutputHandlerConflicts, 8, nil, nil, 1)
	zend.ZendHashInit(&PhpOutputHandlerReverseConflicts, 8, nil, ReverseConflictDtor, 1)
	PhpOutputDirect = PhpOutputStdout
}
func PhpOutputShutdown() {
	PhpOutputDirect = PhpOutputStderr
	zend.ZendHashDestroy(&PhpOutputHandlerAliases)
	zend.ZendHashDestroy(&PhpOutputHandlerConflicts)
	zend.ZendHashDestroy(&PhpOutputHandlerReverseConflicts)
}
func PhpOutputActivate() int {
	memset(&OutputGlobals, 0, b.SizeOf("zend_output_globals"))
	zend.ZendStackInit(&(OG(handlers)), b.SizeOf("php_output_handler *"))
	OG(flags) |= PHP_OUTPUT_ACTIVATED
	return zend.SUCCESS
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
			for b.Assign(&handler, zend.ZendStackTop(&(OG(handlers)))) {
				PhpOutputHandlerFree(handler)
				zend.ZendStackDelTop(&(OG(handlers)))
			}
		}
		zend.ZendStackDestroy(&(OG(handlers)))
	}
}
func PhpOutputRegisterConstants() {
	zend.REGISTER_MAIN_LONG_CONSTANT("PHP_OUTPUT_HANDLER_START", PHP_OUTPUT_HANDLER_START, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_MAIN_LONG_CONSTANT("PHP_OUTPUT_HANDLER_WRITE", PHP_OUTPUT_HANDLER_WRITE, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_MAIN_LONG_CONSTANT("PHP_OUTPUT_HANDLER_FLUSH", PHP_OUTPUT_HANDLER_FLUSH, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_MAIN_LONG_CONSTANT("PHP_OUTPUT_HANDLER_CLEAN", PHP_OUTPUT_HANDLER_CLEAN, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_MAIN_LONG_CONSTANT("PHP_OUTPUT_HANDLER_FINAL", PHP_OUTPUT_HANDLER_FINAL, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_MAIN_LONG_CONSTANT("PHP_OUTPUT_HANDLER_CONT", PHP_OUTPUT_HANDLER_WRITE, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_MAIN_LONG_CONSTANT("PHP_OUTPUT_HANDLER_END", PHP_OUTPUT_HANDLER_FINAL, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_MAIN_LONG_CONSTANT("PHP_OUTPUT_HANDLER_CLEANABLE", PHP_OUTPUT_HANDLER_CLEANABLE, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_MAIN_LONG_CONSTANT("PHP_OUTPUT_HANDLER_FLUSHABLE", PHP_OUTPUT_HANDLER_FLUSHABLE, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_MAIN_LONG_CONSTANT("PHP_OUTPUT_HANDLER_REMOVABLE", PHP_OUTPUT_HANDLER_REMOVABLE, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_MAIN_LONG_CONSTANT("PHP_OUTPUT_HANDLER_STDFLAGS", PHP_OUTPUT_HANDLER_STDFLAGS, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_MAIN_LONG_CONSTANT("PHP_OUTPUT_HANDLER_STARTED", PHP_OUTPUT_HANDLER_STARTED, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_MAIN_LONG_CONSTANT("PHP_OUTPUT_HANDLER_DISABLED", PHP_OUTPUT_HANDLER_DISABLED, zend.CONST_CS|zend.CONST_PERSISTENT)
}
func PhpOutputSetStatus(status int) {
	OG(flags) = OG(flags) & ^0xf | status&0xf
}
func PhpOutputGetStatus() int {
	return (OG(flags) | b.Cond(OG(active), PHP_OUTPUT_ACTIVE, 0) | b.Cond(OG(running), PHP_OUTPUT_LOCKED, 0)) & 0xff
}
func PhpOutputWriteUnbuffered(str *byte, len_ int) int {
	if (OG(flags) & PHP_OUTPUT_ACTIVATED) != 0 {
		return sapi_module.GetUbWrite()(str, len_)
	}
	return PhpOutputDirect(str, len_)
}
func PhpOutputWrite(str *byte, len_ int) int {
	if (OG(flags) & PHP_OUTPUT_ACTIVATED) != 0 {
		PhpOutputOp(PHP_OUTPUT_HANDLER_WRITE, str, len_)
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
			zend.ZendStackDelTop(&(OG(handlers)))
			PhpOutputWrite(context.GetOut().GetData(), context.GetOut().GetUsed())
			zend.ZendStackPush(&(OG(handlers)), &(OG(active)))
		}
		PhpOutputContextDtor(&context)
		return zend.SUCCESS
	}
	return zend.FAILURE
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
		return zend.SUCCESS
	}
	return zend.FAILURE
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
		return zend.SUCCESS
	}
	return zend.FAILURE
}
func PhpOutputEndAll() {
	for OG(active) && PhpOutputStackPop(PHP_OUTPUT_POP_FORCE) != 0 {

	}
}
func PhpOutputDiscard() int {
	if PhpOutputStackPop(PHP_OUTPUT_POP_DISCARD|PHP_OUTPUT_POP_TRY) != 0 {
		return zend.SUCCESS
	}
	return zend.FAILURE
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
func PhpOutputGetContents(p *zend.Zval) int {
	if OG(active) {
		zend.ZVAL_STRINGL(p, OG(active).buffer.data, OG(active).buffer.used)
		return zend.SUCCESS
	} else {
		zend.ZVAL_NULL(p)
		return zend.FAILURE
	}
}
func PhpOutputGetLength(p *zend.Zval) int {
	if OG(active) {
		zend.ZVAL_LONG(p, OG(active).buffer.used)
		return zend.SUCCESS
	} else {
		zend.ZVAL_NULL(p)
		return zend.FAILURE
	}
}
func PhpOutputGetActiveHandler() *PhpOutputHandler { return OG(active) }
func PhpOutputStartDefault() int {
	var handler *PhpOutputHandler
	handler = PhpOutputHandlerCreateInternal(zend.ZEND_STRL(PhpOutputDefaultHandlerName), PhpOutputHandlerDefaultFunc, 0, PHP_OUTPUT_HANDLER_STDFLAGS)
	if zend.SUCCESS == PhpOutputHandlerStart(handler) {
		return zend.SUCCESS
	}
	PhpOutputHandlerFree(&handler)
	return zend.FAILURE
}
func PhpOutputStartDevnull() int {
	var handler *PhpOutputHandler
	handler = PhpOutputHandlerCreateInternal(zend.ZEND_STRL(PhpOutputDevnullHandlerName), PhpOutputHandlerDevnullFunc, PHP_OUTPUT_HANDLER_DEFAULT_SIZE, 0)
	if zend.SUCCESS == PhpOutputHandlerStart(handler) {
		return zend.SUCCESS
	}
	PhpOutputHandlerFree(&handler)
	return zend.FAILURE
}
func PhpOutputStartUser(output_handler *zend.Zval, chunk_size int, flags int) int {
	var handler *PhpOutputHandler
	if output_handler != nil {
		handler = PhpOutputHandlerCreateUser(output_handler, chunk_size, flags)
	} else {
		handler = PhpOutputHandlerCreateInternal(zend.ZEND_STRL(PhpOutputDefaultHandlerName), PhpOutputHandlerDefaultFunc, chunk_size, flags)
	}
	if zend.SUCCESS == PhpOutputHandlerStart(handler) {
		return zend.SUCCESS
	}
	PhpOutputHandlerFree(&handler)
	return zend.FAILURE
}
func PhpOutputStartInternal(name *byte, name_len int, output_handler PhpOutputHandlerFuncT, chunk_size int, flags int) int {
	var handler *PhpOutputHandler
	handler = PhpOutputHandlerCreateInternal(name, name_len, PhpOutputHandlerCompatFunc, chunk_size, flags)
	PhpOutputHandlerSetContext(handler, output_handler, nil)
	if zend.SUCCESS == PhpOutputHandlerStart(handler) {
		return zend.SUCCESS
	}
	PhpOutputHandlerFree(&handler)
	return zend.FAILURE
}
func PhpOutputHandlerCreateUser(output_handler *zend.Zval, chunk_size int, flags int) *PhpOutputHandler {
	var handler_name *zend.ZendString = nil
	var error *byte = nil
	var handler *PhpOutputHandler = nil
	var alias PhpOutputHandlerAliasCtorT = nil
	var user *PhpOutputHandlerUserFuncT = nil
	switch output_handler.GetType() {
	case zend.IS_NULL:
		handler = PhpOutputHandlerCreateInternal(zend.ZEND_STRL(PhpOutputDefaultHandlerName), PhpOutputHandlerDefaultFunc, chunk_size, flags)
		break
	case zend.IS_STRING:
		if zend.Z_STRLEN_P(output_handler) != 0 && b.Assign(&alias, PhpOutputHandlerAlias(zend.Z_STRVAL_P(output_handler), zend.Z_STRLEN_P(output_handler))) {
			handler = alias(zend.Z_STRVAL_P(output_handler), zend.Z_STRLEN_P(output_handler), chunk_size, flags)
			break
		}
	default:
		user = zend.Ecalloc(1, b.SizeOf("php_output_handler_user_func_t"))
		if zend.SUCCESS == zend.ZendFcallInfoInit(output_handler, 0, user.GetFci(), user.GetFcc(), &handler_name, &error) {
			handler = PhpOutputHandlerInit(handler_name, chunk_size, flags & ^0xf | PHP_OUTPUT_HANDLER_USER)
			zend.ZVAL_COPY(user.GetZoh(), output_handler)
			handler.SetUser(user)
		} else {
			zend.Efree(user)
		}
		if error != nil {
			PhpErrorDocref("ref.outcontrol", zend.E_WARNING, "%s", error)
			zend.Efree(error)
		}
		if handler_name != nil {
			zend.ZendStringReleaseEx(handler_name, 0)
		}
	}
	return handler
}
func PhpOutputHandlerCreateInternal(name *byte, name_len int, output_handler PhpOutputHandlerContextFuncT, chunk_size int, flags int) *PhpOutputHandler {
	var handler *PhpOutputHandler
	var str *zend.ZendString = zend.ZendStringInit(name, name_len, 0)
	handler = PhpOutputHandlerInit(str, chunk_size, flags & ^0xf | PHP_OUTPUT_HANDLER_INTERNAL)
	handler.SetInternal(output_handler)
	zend.ZendStringReleaseEx(str, 0)
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
	var rconflicts *zend.HashTable
	var conflict PhpOutputHandlerConflictCheckT
	if PhpOutputLockError(PHP_OUTPUT_HANDLER_START) != 0 || handler == nil {
		return zend.FAILURE
	}
	if nil != b.Assign(&conflict, zend.ZendHashFindPtr(&PhpOutputHandlerConflicts, handler.GetName())) {
		if zend.SUCCESS != conflict(handler.GetName().GetVal(), handler.GetName().GetLen()) {
			return zend.FAILURE
		}
	}
	if nil != b.Assign(&rconflicts, zend.ZendHashFindPtr(&PhpOutputHandlerReverseConflicts, handler.GetName())) {
		for {
			var __ht *zend.HashTable = rconflicts
			var _p *zend.Bucket = __ht.GetArData()
			var _end *zend.Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *zend.Zval = _p.GetVal()

				if _z.IsType(zend.IS_UNDEF) {
					continue
				}
				conflict = _z.GetPtr()
				if zend.SUCCESS != conflict(handler.GetName().GetVal(), handler.GetName().GetLen()) {
					return zend.FAILURE
				}
			}
			break
		}
	}

	/* zend_stack_push returns stack level */

	handler.SetLevel(zend.ZendStackPush(&(OG(handlers)), &handler))
	OG(active) = handler
	return zend.SUCCESS
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
			PhpErrorDocref("ref.outcontrol", zend.E_WARNING, "output handler '%s' conflicts with '%s'", handler_new, handler_set)
		} else {
			PhpErrorDocref("ref.outcontrol", zend.E_WARNING, "output handler '%s' cannot be used twice", handler_new)
		}
		return 1
	}
	return 0
}
func PhpOutputHandlerConflictRegister(name *byte, name_len int, check_func PhpOutputHandlerConflictCheckT) int {
	var str *zend.ZendString
	if zend.__EG().GetCurrentModule() == nil {
		zend.ZendError(zend.E_ERROR, "Cannot register an output handler conflict outside of MINIT")
		return zend.FAILURE
	}
	str = zend.ZendStringInitInterned(name, name_len, 1)
	zend.ZendHashUpdatePtr(&PhpOutputHandlerConflicts, str, check_func)
	zend.ZendStringReleaseEx(str, 1)
	return zend.SUCCESS
}
func PhpOutputHandlerReverseConflictRegister(name *byte, name_len int, check_func PhpOutputHandlerConflictCheckT) int {
	var rev zend.HashTable
	var rev_ptr *zend.HashTable = nil
	if zend.__EG().GetCurrentModule() == nil {
		zend.ZendError(zend.E_ERROR, "Cannot register a reverse output handler conflict outside of MINIT")
		return zend.FAILURE
	}
	if nil != b.Assign(&rev_ptr, zend.ZendHashStrFindPtr(&PhpOutputHandlerReverseConflicts, name, name_len)) {
		if zend.ZendHashNextIndexInsertPtr(rev_ptr, check_func) {
			return zend.SUCCESS
		} else {
			return zend.FAILURE
		}
	} else {
		var str *zend.ZendString
		zend.ZendHashInit(&rev, 8, nil, nil, 1)
		if nil == zend.ZendHashNextIndexInsertPtr(&rev, check_func) {
			zend.ZendHashDestroy(&rev)
			return zend.FAILURE
		}
		str = zend.ZendStringInitInterned(name, name_len, 1)
		zend.ZendHashUpdateMem(&PhpOutputHandlerReverseConflicts, str, &rev, b.SizeOf("HashTable"))
		zend.ZendStringReleaseEx(str, 1)
		return zend.SUCCESS
	}
}
func PhpOutputHandlerAlias(name *byte, name_len int) PhpOutputHandlerAliasCtorT {
	return zend.ZendHashStrFindPtr(&PhpOutputHandlerAliases, name, name_len)
}
func PhpOutputHandlerAliasRegister(name *byte, name_len int, func_ PhpOutputHandlerAliasCtorT) int {
	var str *zend.ZendString
	if zend.__EG().GetCurrentModule() == nil {
		zend.ZendError(zend.E_ERROR, "Cannot register an output handler alias outside of MINIT")
		return zend.FAILURE
	}
	str = zend.ZendStringInitInterned(name, name_len, 1)
	zend.ZendHashUpdatePtr(&PhpOutputHandlerAliases, str, func_)
	zend.ZendStringReleaseEx(str, 1)
	return zend.SUCCESS
}
func PhpOutputHandlerHook(type_ PhpOutputHandlerHookT, arg any) int {
	if OG(running) {
		switch type_ {
		case PHP_OUTPUT_HANDLER_HOOK_GET_OPAQ:
			*((**any)(arg)) = OG(running).opaq
			return zend.SUCCESS
		case PHP_OUTPUT_HANDLER_HOOK_GET_FLAGS:
			*((*int)(arg)) = OG(running).flags
			return zend.SUCCESS
		case PHP_OUTPUT_HANDLER_HOOK_GET_LEVEL:
			*((*int)(arg)) = OG(running).level
			return zend.SUCCESS
		case PHP_OUTPUT_HANDLER_HOOK_IMMUTABLE:
			OG(running).flags &= ^(PHP_OUTPUT_HANDLER_REMOVABLE | PHP_OUTPUT_HANDLER_CLEANABLE)
			return zend.SUCCESS
		case PHP_OUTPUT_HANDLER_HOOK_DISABLE:
			OG(running).flags |= PHP_OUTPUT_HANDLER_DISABLED
			return zend.SUCCESS
		default:
			break
		}
	}
	return zend.FAILURE
}
func PhpOutputHandlerDtor(handler *PhpOutputHandler) {
	if handler.GetName() != nil {
		zend.ZendStringReleaseEx(handler.GetName(), 0)
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
		PhpErrorDocref("ref.outcontrol", zend.E_ERROR, "Cannot use output buffering in output buffering display handlers")
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
func PhpOutputContextFeed(context *PhpOutputContext, data *byte, size int, used int, free zend.ZendBool) {
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
func PhpOutputHandlerInit(name *zend.ZendString, chunk_size int, flags int) *PhpOutputHandler {
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
			var grow_max int = zend.MAX(grow_int, grow_buf)
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
			var retval zend.Zval
			var ob_data zend.Zval
			var ob_mode zend.Zval
			zend.ZVAL_STRINGL(&ob_data, handler.GetBuffer().GetData(), handler.GetBuffer().GetUsed())
			zend.ZVAL_LONG(&ob_mode, zend.ZendLong(context.GetOp()))
			zend.ZendFcallInfoArgn(handler.GetUser().GetFci(), 2, &ob_data, &ob_mode)
			zend.ZvalPtrDtor(&ob_data)
			var PHP_OUTPUT_USER_SUCCESS func(retval zend.Zval) bool = func(retval zend.Zval) bool {
				return retval.GetType() != zend.IS_UNDEF && retval.GetType() != zend.IS_FALSE
			}
			if zend.SUCCESS == zend.ZendFcallInfoCall(handler.GetUser().GetFci(), handler.GetUser().GetFcc(), &retval, nil) && PHP_OUTPUT_USER_SUCCESS(retval) {

				/* user handler may have returned TRUE */

				status = PHP_OUTPUT_HANDLER_NO_DATA
				if retval.GetType() != zend.IS_FALSE && retval.GetType() != zend.IS_TRUE {
					zend.ConvertToStringEx(&retval)
					if zend.Z_STRLEN(retval) != 0 {
						context.GetOut().SetData(zend.Estrndup(zend.Z_STRVAL(retval), zend.Z_STRLEN(retval)))
						context.GetOut().SetUsed(zend.Z_STRLEN(retval))
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
			if zend.SUCCESS == handler.GetInternal()(handler.GetOpaq(), context) {
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
		break
	case PHP_OUTPUT_HANDLER_NO_DATA:

		/* handler ate all */

		PhpOutputContextReset(context)
	case PHP_OUTPUT_HANDLER_SUCCESS:

		/* no more buffered data */

		handler.GetBuffer().SetUsed(0)
		handler.SetIsProcessed(true)
		break
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
		} else if b.Assign(&active, zend.ZendStackTop(&(OG(handlers)))) && !active.IsDisabled() {
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
			sapi_module.GetUbWrite()(context.GetOut().GetData(), context.GetOut().GetUsed())
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
	var array *zend.Zval = (*zend.Zval)(z)
	zend.AddNextIndexStr(array, handler.GetName().Copy())
	return 0
}
func PhpOutputStackApplyStatus(h any, z any) int {
	var handler *PhpOutputHandler = *((**PhpOutputHandler)(h))
	var arr zend.Zval
	var array *zend.Zval = (*zend.Zval)(z)
	zend.AddNextIndexZval(array, PhpOutputHandlerStatus(handler, &arr))
	return 0
}
func PhpOutputHandlerStatus(handler *PhpOutputHandler, entry *zend.Zval) *zend.Zval {
	zend.ZEND_ASSERT(entry != nil)
	zend.ArrayInit(entry)
	zend.AddAssocStr(entry, "name", handler.GetName().Copy())
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
			PhpErrorDocref("ref.outcontrol", zend.E_NOTICE, "failed to %s buffer. No buffer to %s", b.Cond((flags&PHP_OUTPUT_POP_DISCARD) != 0, "discard", "send"), b.Cond((flags&PHP_OUTPUT_POP_DISCARD) != 0, "discard", "send"))
		}
		return 0
	} else if (flags&PHP_OUTPUT_POP_FORCE) == 0 && !orphan.HasFlags(PHP_OUTPUT_HANDLER_REMOVABLE) {
		if (flags & PHP_OUTPUT_POP_SILENT) == 0 {
			PhpErrorDocref("ref.outcontrol", zend.E_NOTICE, "failed to %s buffer of %s (%d)", b.Cond((flags&PHP_OUTPUT_POP_DISCARD) != 0, "discard", "send"), orphan.GetName().GetVal(), orphan.GetLevel())
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

		zend.ZendStackDelTop(&(OG(handlers)))
		if b.Assign(&current, zend.ZendStackTop(&(OG(handlers)))) {
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
		return zend.SUCCESS
	}
	return zend.FAILURE
}
func PhpOutputHandlerDefaultFunc(handler_context *any, output_context *PhpOutputContext) int {
	PhpOutputContextPass(output_context)
	return zend.SUCCESS
}
func PhpOutputHandlerDevnullFunc(handler_context *any, output_context *PhpOutputContext) int {
	return zend.SUCCESS
}
func ZifObStart(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var output_handler *zend.Zval = nil
	var chunk_size zend.ZendLong = 0
	var flags zend.ZendLong = PHP_OUTPUT_HANDLER_STDFLAGS
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "|zll", &output_handler, &chunk_size, &flags) == zend.FAILURE {
		return
	}
	if chunk_size < 0 {
		chunk_size = 0
	}
	if PhpOutputStartUser(output_handler, chunk_size, flags) == zend.FAILURE {
		PhpErrorDocref("ref.outcontrol", zend.E_NOTICE, "failed to create buffer")
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_TRUE
	return
}
func ZifObFlush(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if !(OG(active)) {
		PhpErrorDocref("ref.outcontrol", zend.E_NOTICE, "failed to flush buffer. No buffer to flush")
		zend.RETVAL_FALSE
		return
	}
	if zend.SUCCESS != PhpOutputFlush() {
		PhpErrorDocref("ref.outcontrol", zend.E_NOTICE, "failed to flush buffer of %s (%d)", OG(active).name.GetVal(), OG(active).level)
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_TRUE
	return
}
func ZifObClean(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if !(OG(active)) {
		PhpErrorDocref("ref.outcontrol", zend.E_NOTICE, "failed to delete buffer. No buffer to delete")
		zend.RETVAL_FALSE
		return
	}
	if zend.SUCCESS != PhpOutputClean() {
		PhpErrorDocref("ref.outcontrol", zend.E_NOTICE, "failed to delete buffer of %s (%d)", OG(active).name.GetVal(), OG(active).level)
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_TRUE
	return
}
func ZifObEndFlush(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if !(OG(active)) {
		PhpErrorDocref("ref.outcontrol", zend.E_NOTICE, "failed to delete and flush buffer. No buffer to delete or flush")
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_BOOL(zend.SUCCESS == PhpOutputEnd())
	return
}
func ZifObEndClean(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if !(OG(active)) {
		PhpErrorDocref("ref.outcontrol", zend.E_NOTICE, "failed to delete buffer. No buffer to delete")
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_BOOL(zend.SUCCESS == PhpOutputDiscard())
	return
}
func ZifObGetFlush(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if PhpOutputGetContents(return_value) == zend.FAILURE {
		PhpErrorDocref("ref.outcontrol", zend.E_NOTICE, "failed to delete and flush buffer. No buffer to delete or flush")
		zend.RETVAL_FALSE
		return
	}
	if zend.SUCCESS != PhpOutputEnd() {
		PhpErrorDocref("ref.outcontrol", zend.E_NOTICE, "failed to delete buffer of %s (%d)", OG(active).name.GetVal(), OG(active).level)
	}
}
func ZifObGetClean(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if !(OG(active)) {
		zend.RETVAL_FALSE
		return
	}
	if PhpOutputGetContents(return_value) == zend.FAILURE {
		PhpErrorDocref("ref.outcontrol", zend.E_NOTICE, "failed to delete buffer. No buffer to delete")
		zend.RETVAL_FALSE
		return
	}
	if zend.SUCCESS != PhpOutputDiscard() {
		PhpErrorDocref("ref.outcontrol", zend.E_NOTICE, "failed to delete buffer of %s (%d)", OG(active).name.GetVal(), OG(active).level)
	}
}
func ZifObGetContents(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if PhpOutputGetContents(return_value) == zend.FAILURE {
		zend.RETVAL_FALSE
		return
	}
}
func ZifObGetLevel(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_LONG(PhpOutputGetLevel())
	return
}
func ZifObGetLength(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if PhpOutputGetLength(return_value) == zend.FAILURE {
		zend.RETVAL_FALSE
		return
	}
}
func ZifObListHandlers(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ArrayInit(return_value)
	if !(OG(active)) {
		return
	}
	zend.ZendStackApplyWithArgument(&(OG(handlers)), zend.ZEND_STACK_APPLY_BOTTOMUP, PhpOutputStackApplyList, return_value)
}
func ZifObGetStatus(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var full_status zend.ZendBool = 0
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "|b", &full_status) == zend.FAILURE {
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
func ZifObImplicitFlush(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var flag zend.ZendLong = 1
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "|l", &flag) == zend.FAILURE {
		return
	}
	PhpOutputSetImplicitFlush(flag)
}
func ZifOutputResetRewriteVars(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if standard.PhpUrlScannerResetVars() == zend.SUCCESS {
		zend.RETVAL_TRUE
		return
	} else {
		zend.RETVAL_FALSE
		return
	}
}
func ZifOutputAddRewriteVar(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var name *byte
	var value *byte
	var name_len int
	var value_len int
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "ss", &name, &name_len, &value, &value_len) == zend.FAILURE {
		return
	}
	if standard.PhpUrlScannerAddVar(name, name_len, value, value_len, 1) == zend.SUCCESS {
		zend.RETVAL_TRUE
		return
	} else {
		zend.RETVAL_FALSE
		return
	}
}
