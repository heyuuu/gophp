// <<generate>>

package core

import (
	"sik/ext/standard"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <main/output.c>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Authors: Zeev Suraski <zeev@php.net>                                 |
   |          Thies C. Arntzen <thies@thieso.net>                         |
   |          Marcus Boerger <helly@php.net>                              |
   | New API: Michael Wallner <mike@php.net>                              |
   +----------------------------------------------------------------------+
*/

// #define PHP_OUTPUT_DEBUG       0

// #define PHP_OUTPUT_NOINLINE       0

// # include "php.h"

// # include "ext/standard/head.h"

// # include "ext/standard/url_scanner_ex.h"

// # include "SAPI.h"

// # include "zend_stack.h"

// # include "php_output.h"

var OutputGlobals ZendOutputGlobals
var PhpOutputDefaultHandlerName []byte = "default output handler"
var PhpOutputDevnullHandlerName []byte = "null output handler"

/* {{{ aliases, conflict and reverse conflict hash tables */

var PhpOutputHandlerAliases zend.HashTable
var PhpOutputHandlerConflicts zend.HashTable
var PhpOutputHandlerReverseConflicts zend.HashTable

/* }}} */

/* }}} */

func PhpOutputInitGlobals(G *ZendOutputGlobals) { memset(G, 0, g.SizeOf("* G")) }

/* }}} */

func PhpOutputStdout(str *byte, str_len int) int {
	fwrite(str, 1, str_len, stdout)
	return str_len
}
func PhpOutputStderr(str *byte, str_len int) int {
	fwrite(str, 1, str_len, stderr)

	/* See http://support.microsoft.com/kb/190351 */

	return str_len

	/* See http://support.microsoft.com/kb/190351 */
}

var PhpOutputDirect func(str *byte, str_len int) int = PhpOutputStderr

/* }}} */

func PhpOutputHeader() {
	if sapi_globals.GetHeadersSent() == 0 {
		if OutputGlobals.GetOutputStartFilename() == nil {
			if zend.ZendIsCompiling() != 0 {
				OutputGlobals.SetOutputStartFilename(zend.ZendGetCompiledFilename().val)
				OutputGlobals.SetOutputStartLineno(zend.ZendGetCompiledLineno())
			} else if zend.ZendIsExecuting() != 0 {
				OutputGlobals.SetOutputStartFilename(zend.ZendGetExecutedFilename())
				OutputGlobals.SetOutputStartLineno(zend.ZendGetExecutedLineno())
			}
		}
		if standard.PhpHeader() == 0 {
			OutputGlobals.SetFlags(OutputGlobals.GetFlags() | 0x2)
		}
	}
}

/* }}} */

func ReverseConflictDtor(zv *zend.Zval) {
	var ht *zend.HashTable = zv.value.ptr
	zend.ZendHashDestroy(ht)
}

/* {{{ void php_output_startup(void)
 * Set up module globals and initialize the conflict and reverse conflict hash tables */

func PhpOutputStartup() {
	PhpOutputInitGlobals(&OutputGlobals)
	zend._zendHashInit(&PhpOutputHandlerAliases, 8, nil, 1)
	zend._zendHashInit(&PhpOutputHandlerConflicts, 8, nil, 1)
	zend._zendHashInit(&PhpOutputHandlerReverseConflicts, 8, ReverseConflictDtor, 1)
	PhpOutputDirect = PhpOutputStdout
}

/* }}} */

func PhpOutputShutdown() {
	PhpOutputDirect = PhpOutputStderr
	zend.ZendHashDestroy(&PhpOutputHandlerAliases)
	zend.ZendHashDestroy(&PhpOutputHandlerConflicts)
	zend.ZendHashDestroy(&PhpOutputHandlerReverseConflicts)
}

/* }}} */

func PhpOutputActivate() int {
	memset(&OutputGlobals, 0, g.SizeOf("zend_output_globals"))
	zend.ZendStackInit(&(OutputGlobals.GetHandlers()), g.SizeOf("php_output_handler *"))
	OutputGlobals.SetFlags(OutputGlobals.GetFlags() | 0x100000)
	return zend.SUCCESS
}

/* }}} */

func PhpOutputDeactivate() {
	var handler **PhpOutputHandler = nil
	if (OutputGlobals.GetFlags() & 0x100000) != 0 {
		PhpOutputHeader()
		OutputGlobals.SetFlags(OutputGlobals.GetFlags() ^ 0x100000)
		OutputGlobals.SetActive(nil)
		OutputGlobals.SetRunning(nil)

		/* release all output handlers */

		if OutputGlobals.handlers.elements {
			for g.Assign(&handler, zend.ZendStackTop(&(OutputGlobals.GetHandlers()))) {
				PhpOutputHandlerFree(handler)
				zend.ZendStackDelTop(&(OutputGlobals.GetHandlers()))
			}
		}
		zend.ZendStackDestroy(&(OutputGlobals.GetHandlers()))
	}
}

/* }}} */

func PhpOutputRegisterConstants() {
	zend.ZendRegisterLongConstant("PHP_OUTPUT_HANDLER_START", g.SizeOf("\"PHP_OUTPUT_HANDLER_START\"")-1, 0x1, 1<<0|1<<1, 0)
	zend.ZendRegisterLongConstant("PHP_OUTPUT_HANDLER_WRITE", g.SizeOf("\"PHP_OUTPUT_HANDLER_WRITE\"")-1, 0x0, 1<<0|1<<1, 0)
	zend.ZendRegisterLongConstant("PHP_OUTPUT_HANDLER_FLUSH", g.SizeOf("\"PHP_OUTPUT_HANDLER_FLUSH\"")-1, 0x4, 1<<0|1<<1, 0)
	zend.ZendRegisterLongConstant("PHP_OUTPUT_HANDLER_CLEAN", g.SizeOf("\"PHP_OUTPUT_HANDLER_CLEAN\"")-1, 0x2, 1<<0|1<<1, 0)
	zend.ZendRegisterLongConstant("PHP_OUTPUT_HANDLER_FINAL", g.SizeOf("\"PHP_OUTPUT_HANDLER_FINAL\"")-1, 0x8, 1<<0|1<<1, 0)
	zend.ZendRegisterLongConstant("PHP_OUTPUT_HANDLER_CONT", g.SizeOf("\"PHP_OUTPUT_HANDLER_CONT\"")-1, 0x0, 1<<0|1<<1, 0)
	zend.ZendRegisterLongConstant("PHP_OUTPUT_HANDLER_END", g.SizeOf("\"PHP_OUTPUT_HANDLER_END\"")-1, 0x8, 1<<0|1<<1, 0)
	zend.ZendRegisterLongConstant("PHP_OUTPUT_HANDLER_CLEANABLE", g.SizeOf("\"PHP_OUTPUT_HANDLER_CLEANABLE\"")-1, 0x10, 1<<0|1<<1, 0)
	zend.ZendRegisterLongConstant("PHP_OUTPUT_HANDLER_FLUSHABLE", g.SizeOf("\"PHP_OUTPUT_HANDLER_FLUSHABLE\"")-1, 0x20, 1<<0|1<<1, 0)
	zend.ZendRegisterLongConstant("PHP_OUTPUT_HANDLER_REMOVABLE", g.SizeOf("\"PHP_OUTPUT_HANDLER_REMOVABLE\"")-1, 0x40, 1<<0|1<<1, 0)
	zend.ZendRegisterLongConstant("PHP_OUTPUT_HANDLER_STDFLAGS", g.SizeOf("\"PHP_OUTPUT_HANDLER_STDFLAGS\"")-1, 0x70, 1<<0|1<<1, 0)
	zend.ZendRegisterLongConstant("PHP_OUTPUT_HANDLER_STARTED", g.SizeOf("\"PHP_OUTPUT_HANDLER_STARTED\"")-1, 0x1000, 1<<0|1<<1, 0)
	zend.ZendRegisterLongConstant("PHP_OUTPUT_HANDLER_DISABLED", g.SizeOf("\"PHP_OUTPUT_HANDLER_DISABLED\"")-1, 0x2000, 1<<0|1<<1, 0)
}

/* }}} */

func PhpOutputSetStatus(status int) {
	OutputGlobals.SetFlags(OutputGlobals.GetFlags() & ^0xf | status&0xf)
}

/* }}} */

func PhpOutputGetStatus() int {
	return (OutputGlobals.GetFlags() | g.Cond(OutputGlobals.GetActive() != nil, 0x10, 0) | g.Cond(OutputGlobals.GetRunning() != nil, 0x20, 0)) & 0xff
}

/* }}} */

func PhpOutputWriteUnbuffered(str *byte, len_ int) int {
	if (OutputGlobals.GetFlags() & 0x100000) != 0 {
		return sapi_module.GetUbWrite()(str, len_)
	}
	return PhpOutputDirect(str, len_)
}

/* }}} */

func PhpOutputWrite(str *byte, len_ int) int {
	if (OutputGlobals.GetFlags() & 0x100000) != 0 {
		PhpOutputOp(0x0, str, len_)
		return len_
	}
	if (OutputGlobals.GetFlags() & 0x2) != 0 {
		return 0
	}
	return PhpOutputDirect(str, len_)
}

/* }}} */

func PhpOutputFlush() int {
	var context PhpOutputContext
	if OutputGlobals.GetActive() != nil && (OutputGlobals.GetActive().GetFlags()&0x20) != 0 {
		PhpOutputContextInit(&context, 0x4)
		PhpOutputHandlerOp(OutputGlobals.GetActive(), &context)
		if context.GetOut().GetData() != nil && context.GetOut().GetUsed() != 0 {
			zend.ZendStackDelTop(&(OutputGlobals.GetHandlers()))
			PhpOutputWrite(context.GetOut().GetData(), context.GetOut().GetUsed())
			zend.ZendStackPush(&(OutputGlobals.GetHandlers()), &(OutputGlobals.GetActive()))
		}
		PhpOutputContextDtor(&context)
		return zend.SUCCESS
	}
	return zend.FAILURE
}

/* }}} */

func PhpOutputFlushAll() {
	if OutputGlobals.GetActive() != nil {
		PhpOutputOp(0x4, nil, 0)
	}
}

/* }}} */

func PhpOutputClean() int {
	var context PhpOutputContext
	if OutputGlobals.GetActive() != nil && (OutputGlobals.GetActive().GetFlags()&0x10) != 0 {
		PhpOutputContextInit(&context, 0x2)
		PhpOutputHandlerOp(OutputGlobals.GetActive(), &context)
		PhpOutputContextDtor(&context)
		return zend.SUCCESS
	}
	return zend.FAILURE
}

/* }}} */

func PhpOutputCleanAll() {
	var context PhpOutputContext
	if OutputGlobals.GetActive() != nil {
		PhpOutputContextInit(&context, 0x2)
		zend.ZendStackApplyWithArgument(&(OutputGlobals.GetHandlers()), 1, PhpOutputStackApplyClean, &context)
	}
}

/* {{{ SUCCESS|FAILURE php_output_end(void)
 * Finalizes the most recent output handler at pops it off the stack if the handler is removable */

func PhpOutputEnd() int {
	if PhpOutputStackPop(0x0) != 0 {
		return zend.SUCCESS
	}
	return zend.FAILURE
}

/* }}} */

func PhpOutputEndAll() {
	for OutputGlobals.GetActive() != nil && PhpOutputStackPop(0x1) != 0 {

	}
}

/* }}} */

func PhpOutputDiscard() int {
	if PhpOutputStackPop(0x10|0x0) != 0 {
		return zend.SUCCESS
	}
	return zend.FAILURE
}

/* }}} */

func PhpOutputDiscardAll() {
	for OutputGlobals.GetActive() != nil {
		PhpOutputStackPop(0x10 | 0x1)
	}
}

/* }}} */

func PhpOutputGetLevel() int {
	if OutputGlobals.GetActive() != nil {
		return zend.ZendStackCount(&(OutputGlobals.GetHandlers()))
	} else {
		return 0
	}
}

/* }}} */

func PhpOutputGetContents(p *zend.Zval) int {
	if OutputGlobals.GetActive() != nil {
		var __z *zend.Zval = p
		var __s *zend.ZendString = zend.ZendStringInit(OutputGlobals.GetActive().GetBuffer().GetData(), OutputGlobals.GetActive().GetBuffer().GetUsed(), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return zend.SUCCESS
	} else {
		p.u1.type_info = 1
		return zend.FAILURE
	}
}

/* {{{ SUCCESS|FAILURE php_output_get_length(zval *z)
 * Get the length of the active output handlers buffer */

func PhpOutputGetLength(p *zend.Zval) int {
	if OutputGlobals.GetActive() != nil {
		var __z *zend.Zval = p
		__z.value.lval = OutputGlobals.GetActive().GetBuffer().GetUsed()
		__z.u1.type_info = 4
		return zend.SUCCESS
	} else {
		p.u1.type_info = 1
		return zend.FAILURE
	}
}

/* }}} */

func PhpOutputGetActiveHandler() *PhpOutputHandler { return OutputGlobals.GetActive() }

/* }}} */

func PhpOutputStartDefault() int {
	var handler *PhpOutputHandler
	handler = PhpOutputHandlerCreateInternal(PhpOutputDefaultHandlerName, g.SizeOf("php_output_default_handler_name")-1, PhpOutputHandlerDefaultFunc, 0, 0x70)
	if zend.SUCCESS == PhpOutputHandlerStart(handler) {
		return zend.SUCCESS
	}
	PhpOutputHandlerFree(&handler)
	return zend.FAILURE
}

/* }}} */

func PhpOutputStartDevnull() int {
	var handler *PhpOutputHandler
	handler = PhpOutputHandlerCreateInternal(PhpOutputDevnullHandlerName, g.SizeOf("php_output_devnull_handler_name")-1, PhpOutputHandlerDevnullFunc, 0x4000, 0)
	if zend.SUCCESS == PhpOutputHandlerStart(handler) {
		return zend.SUCCESS
	}
	PhpOutputHandlerFree(&handler)
	return zend.FAILURE
}

/* }}} */

func PhpOutputStartUser(output_handler *zend.Zval, chunk_size int, flags int) int {
	var handler *PhpOutputHandler
	if output_handler != nil {
		handler = PhpOutputHandlerCreateUser(output_handler, chunk_size, flags)
	} else {
		handler = PhpOutputHandlerCreateInternal(PhpOutputDefaultHandlerName, g.SizeOf("php_output_default_handler_name")-1, PhpOutputHandlerDefaultFunc, chunk_size, flags)
	}
	if zend.SUCCESS == PhpOutputHandlerStart(handler) {
		return zend.SUCCESS
	}
	PhpOutputHandlerFree(&handler)
	return zend.FAILURE
}

/* }}} */

func PhpOutputStartInternal(name string, name_len int, output_handler PhpOutputHandlerFuncT, chunk_size int, flags int) int {
	var handler *PhpOutputHandler
	handler = PhpOutputHandlerCreateInternal(name, name_len, PhpOutputHandlerCompatFunc, chunk_size, flags)
	PhpOutputHandlerSetContext(handler, output_handler, nil)
	if zend.SUCCESS == PhpOutputHandlerStart(handler) {
		return zend.SUCCESS
	}
	PhpOutputHandlerFree(&handler)
	return zend.FAILURE
}

/* }}} */

func PhpOutputHandlerCreateUser(output_handler *zend.Zval, chunk_size int, flags int) *PhpOutputHandler {
	var handler_name *zend.ZendString = nil
	var error *byte = nil
	var handler *PhpOutputHandler = nil
	var alias PhpOutputHandlerAliasCtorT = nil
	var user *PhpOutputHandlerUserFuncT = nil
	switch output_handler.u1.v.type_ {
	case 1:
		handler = PhpOutputHandlerCreateInternal(PhpOutputDefaultHandlerName, g.SizeOf("php_output_default_handler_name")-1, PhpOutputHandlerDefaultFunc, chunk_size, flags)
		break
	case 6:
		if output_handler.value.str.len_ != 0 && g.Assign(&alias, PhpOutputHandlerAlias(output_handler.value.str.val, output_handler.value.str.len_)) {
			handler = alias(output_handler.value.str.val, output_handler.value.str.len_, chunk_size, flags)
			break
		}
	default:
		user = zend._ecalloc(1, g.SizeOf("php_output_handler_user_func_t"))
		if zend.SUCCESS == zend.ZendFcallInfoInit(output_handler, 0, &user.fci, &user.fcc, &handler_name, &error) {
			handler = PhpOutputHandlerInit(handler_name, chunk_size, flags & ^0xf | 0x1)
			var _z1 *zend.Zval = &user.zoh
			var _z2 *zend.Zval = output_handler
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			if (_t & 0xff00) != 0 {
				zend.ZendGcAddref(&_gc.gc)
			}
			handler.SetUser(user)
		} else {
			zend._efree(user)
		}
		if error != nil {
			PhpErrorDocref("ref.outcontrol", 1<<1, "%s", error)
			zend._efree(error)
		}
		if handler_name != nil {
			zend.ZendStringReleaseEx(handler_name, 0)
		}
	}
	return handler
}

/* }}} */

func PhpOutputHandlerCreateInternal(name *byte, name_len int, output_handler PhpOutputHandlerContextFuncT, chunk_size int, flags int) *PhpOutputHandler {
	var handler *PhpOutputHandler
	var str *zend.ZendString = zend.ZendStringInit(name, name_len, 0)
	handler = PhpOutputHandlerInit(str, chunk_size, flags & ^0xf | 0x0)
	handler.SetInternal(output_handler)
	zend.ZendStringReleaseEx(str, 0)
	return handler
}

/* }}} */

func PhpOutputHandlerSetContext(handler *PhpOutputHandler, opaq any, dtor func(any)) {
	if handler.GetDtor() != nil && handler.GetOpaq() {
		handler.GetDtor()(handler.GetOpaq())
	}
	handler.SetDtor(dtor)
	handler.SetOpaq(opaq)
}

/* }}} */

func PhpOutputHandlerStart(handler *PhpOutputHandler) int {
	var rconflicts *zend.HashTable
	var conflict PhpOutputHandlerConflictCheckT
	if PhpOutputLockError(0x1) != 0 || handler == nil {
		return zend.FAILURE
	}
	if nil != g.Assign(&conflict, zend.ZendHashFindPtr(&PhpOutputHandlerConflicts, handler.GetName())) {
		if zend.SUCCESS != conflict(handler.GetName().val, handler.GetName().len_) {
			return zend.FAILURE
		}
	}
	if nil != g.Assign(&rconflicts, zend.ZendHashFindPtr(&PhpOutputHandlerReverseConflicts, handler.GetName())) {
		for {
			var __ht *zend.HashTable = rconflicts
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if _z.u1.v.type_ == 0 {
					continue
				}
				conflict = _z.value.ptr
				if zend.SUCCESS != conflict(handler.GetName().val, handler.GetName().len_) {
					return zend.FAILURE
				}
			}
			break
		}
	}

	/* zend_stack_push returns stack level */

	handler.SetLevel(zend.ZendStackPush(&(OutputGlobals.GetHandlers()), &handler))
	OutputGlobals.SetActive(handler)
	return zend.SUCCESS
}

/* }}} */

func PhpOutputHandlerStarted(name *byte, name_len int) int {
	var handlers **PhpOutputHandler
	var i int
	var count int = PhpOutputGetLevel()
	if count != 0 {
		handlers = (**PhpOutputHandler)(zend.ZendStackBase(&(OutputGlobals.GetHandlers())))
		for i = 0; i < count; i++ {
			if name_len == handlers[i].GetName().len_ && !(memcmp(handlers[i].GetName().val, name, name_len)) {
				return 1
			}
		}
	}
	return 0
}

/* }}} */

func PhpOutputHandlerConflict(handler_new *byte, handler_new_len int, handler_set *byte, handler_set_len int) int {
	if PhpOutputHandlerStarted(handler_set, handler_set_len) != 0 {
		if handler_new_len != handler_set_len || memcmp(handler_new, handler_set, handler_set_len) {
			PhpErrorDocref("ref.outcontrol", 1<<1, "output handler '%s' conflicts with '%s'", handler_new, handler_set)
		} else {
			PhpErrorDocref("ref.outcontrol", 1<<1, "output handler '%s' cannot be used twice", handler_new)
		}
		return 1
	}
	return 0
}

/* }}} */

func PhpOutputHandlerConflictRegister(name *byte, name_len int, check_func PhpOutputHandlerConflictCheckT) int {
	var str *zend.ZendString
	if zend.EG.current_module == nil {
		zend.ZendError(1<<0, "Cannot register an output handler conflict outside of MINIT")
		return zend.FAILURE
	}
	str = zend.ZendStringInitInterned(name, name_len, 1)
	zend.ZendHashUpdatePtr(&PhpOutputHandlerConflicts, str, check_func)
	zend.ZendStringReleaseEx(str, 1)
	return zend.SUCCESS
}

/* }}} */

func PhpOutputHandlerReverseConflictRegister(name *byte, name_len int, check_func PhpOutputHandlerConflictCheckT) int {
	var rev zend.HashTable
	var rev_ptr *zend.HashTable = nil
	if zend.EG.current_module == nil {
		zend.ZendError(1<<0, "Cannot register a reverse output handler conflict outside of MINIT")
		return zend.FAILURE
	}
	if nil != g.Assign(&rev_ptr, zend.ZendHashStrFindPtr(&PhpOutputHandlerReverseConflicts, name, name_len)) {
		if zend.ZendHashNextIndexInsertPtr(rev_ptr, check_func) {
			return zend.SUCCESS
		} else {
			return zend.FAILURE
		}
	} else {
		var str *zend.ZendString
		zend._zendHashInit(&rev, 8, nil, 1)
		if nil == zend.ZendHashNextIndexInsertPtr(&rev, check_func) {
			zend.ZendHashDestroy(&rev)
			return zend.FAILURE
		}
		str = zend.ZendStringInitInterned(name, name_len, 1)
		zend.ZendHashUpdateMem(&PhpOutputHandlerReverseConflicts, str, &rev, g.SizeOf("HashTable"))
		zend.ZendStringReleaseEx(str, 1)
		return zend.SUCCESS
	}
}

/* }}} */

func PhpOutputHandlerAlias(name *byte, name_len int) PhpOutputHandlerAliasCtorT {
	return zend.ZendHashStrFindPtr(&PhpOutputHandlerAliases, name, name_len)
}

/* }}} */

func PhpOutputHandlerAliasRegister(name *byte, name_len int, func_ PhpOutputHandlerAliasCtorT) int {
	var str *zend.ZendString
	if zend.EG.current_module == nil {
		zend.ZendError(1<<0, "Cannot register an output handler alias outside of MINIT")
		return zend.FAILURE
	}
	str = zend.ZendStringInitInterned(name, name_len, 1)
	zend.ZendHashUpdatePtr(&PhpOutputHandlerAliases, str, func_)
	zend.ZendStringReleaseEx(str, 1)
	return zend.SUCCESS
}

/* }}} */

func PhpOutputHandlerHook(type_ PhpOutputHandlerHookT, arg any) int {
	if OutputGlobals.GetRunning() != nil {
		switch type_ {
		case PHP_OUTPUT_HANDLER_HOOK_GET_OPAQ:
			*((**any)(arg)) = &(OutputGlobals.GetRunning()).opaq
			return zend.SUCCESS
		case PHP_OUTPUT_HANDLER_HOOK_GET_FLAGS:
			*((*int)(arg)) = OutputGlobals.GetRunning().GetFlags()
			return zend.SUCCESS
		case PHP_OUTPUT_HANDLER_HOOK_GET_LEVEL:
			*((*int)(arg)) = OutputGlobals.GetRunning().GetLevel()
			return zend.SUCCESS
		case PHP_OUTPUT_HANDLER_HOOK_IMMUTABLE:
			OutputGlobals.GetRunning().SetFlags(OutputGlobals.GetRunning().GetFlags() &^ (0x40 | 0x10))
			return zend.SUCCESS
		case PHP_OUTPUT_HANDLER_HOOK_DISABLE:
			OutputGlobals.GetRunning().SetFlags(OutputGlobals.GetRunning().GetFlags() | 0x2000)
			return zend.SUCCESS
		default:
			break
		}
	}
	return zend.FAILURE
}

/* }}} */

func PhpOutputHandlerDtor(handler *PhpOutputHandler) {
	if handler.GetName() != nil {
		zend.ZendStringReleaseEx(handler.GetName(), 0)
	}
	if handler.GetBuffer().GetData() != nil {
		zend._efree(handler.GetBuffer().GetData())
	}
	if (handler.GetFlags() & 0x1) != 0 {
		zend.ZvalPtrDtor(&handler.func_.user.GetZoh())
		zend._efree(handler.GetUser())
	}
	if handler.GetDtor() != nil && handler.GetOpaq() {
		handler.GetDtor()(handler.GetOpaq())
	}
	memset(handler, 0, g.SizeOf("* handler"))
}

/* }}} */

func PhpOutputHandlerFree(h **PhpOutputHandler) {
	if (*h) != nil {
		PhpOutputHandlerDtor(*h)
		zend._efree(*h)
		*h = nil
	}
}

/* }}} */

func PhpOutputSetImplicitFlush(flush int) {
	if flush != 0 {
		OutputGlobals.SetFlags(OutputGlobals.GetFlags() | 0x1)
	} else {
		OutputGlobals.SetFlags(OutputGlobals.GetFlags() &^ 0x1)
	}
}

/* }}} */

func PhpOutputGetStartFilename() *byte {
	return OutputGlobals.GetOutputStartFilename()
}

/* }}} */

func PhpOutputGetStartLineno() int {
	return OutputGlobals.GetOutputStartLineno()
}

/* }}} */

func PhpOutputLockError(op int) int {
	/* if there's no ob active, ob has been stopped */

	if op != 0 && OutputGlobals.GetActive() != nil && OutputGlobals.GetRunning() != nil {

		/* fatal error */

		PhpOutputDeactivate()
		PhpErrorDocref("ref.outcontrol", 1<<0, "Cannot use output buffering in output buffering display handlers")
		return 1
	}
	return 0
}

/* }}} */

func PhpOutputContextInit(context *PhpOutputContext, op int) {
	memset(context, 0, g.SizeOf("php_output_context"))
	context.SetOp(op)
}

/* }}} */

func PhpOutputContextReset(context *PhpOutputContext) {
	var op int = context.GetOp()
	PhpOutputContextDtor(context)
	memset(context, 0, g.SizeOf("php_output_context"))
	context.SetOp(op)
}

/* }}} */

func PhpOutputContextFeed(context *PhpOutputContext, data *byte, size int, used int, free zend.ZendBool) {
	if context.GetIn().GetFree() != 0 && context.GetIn().GetData() != nil {
		zend._efree(context.GetIn().GetData())
	}
	context.GetIn().SetData(data)
	context.GetIn().SetUsed(used)
	context.GetIn().SetFree(free)
	context.GetIn().SetSize(size)
}

/* }}} */

func PhpOutputContextSwap(context *PhpOutputContext) {
	if context.GetIn().GetFree() != 0 && context.GetIn().GetData() != nil {
		zend._efree(context.GetIn().GetData())
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

/* }}} */

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

/* }}} */

func PhpOutputContextDtor(context *PhpOutputContext) {
	if context.GetIn().GetFree() != 0 && context.GetIn().GetData() != nil {
		zend._efree(context.GetIn().GetData())
		context.GetIn().SetData(nil)
	}
	if context.GetOut().GetFree() != 0 && context.GetOut().GetData() != nil {
		zend._efree(context.GetOut().GetData())
		context.GetOut().SetData(nil)
	}
}

/* }}} */

func PhpOutputHandlerInit(name *zend.ZendString, chunk_size int, flags int) *PhpOutputHandler {
	var handler *PhpOutputHandler
	handler = zend._ecalloc(1, g.SizeOf("php_output_handler"))
	handler.SetName(zend.ZendStringCopy(name))
	handler.SetSize(chunk_size)
	handler.SetFlags(flags)
	if chunk_size > 1 {
		handler.GetBuffer().SetSize(chunk_size + 0x1000 - chunk_size%0x1000)
	} else {
		handler.GetBuffer().SetSize(0x4000)
	}
	handler.GetBuffer().SetData(zend._emalloc(handler.GetBuffer().GetSize()))
	return handler
}

/* }}} */

func PhpOutputHandlerAppend(handler *PhpOutputHandler, buf *PhpOutputBuffer) int {
	if buf.GetUsed() != 0 {
		OutputGlobals.SetFlags(OutputGlobals.GetFlags() | 0x4)

		/* store it away */

		if handler.GetBuffer().GetSize()-handler.GetBuffer().GetUsed() <= buf.GetUsed() {
			var grow_int int = g.CondF1(handler.GetSize() > 1, func() int { return handler.GetSize() + 0x1000 - handler.GetSize()%0x1000 }, 0x4000)
			var grow_buf int = g.CondF1(buf.GetUsed()-(handler.GetBuffer().GetSize()-handler.GetBuffer().GetUsed()) > 1, func() int {
				return buf.GetUsed() - (handler.GetBuffer().GetSize() - handler.GetBuffer().GetUsed()) + 0x1000 - (buf.GetUsed()-(handler.GetBuffer().GetSize()-handler.GetBuffer().GetUsed()))%0x1000
			}, 0x4000)
			var grow_max int = g.Cond(grow_int > grow_buf, grow_int, grow_buf)
			handler.GetBuffer().SetData(zend._safeErealloc(handler.GetBuffer().GetData(), 1, handler.GetBuffer().GetSize(), grow_max))
			handler.GetBuffer().SetSize(handler.GetBuffer().GetSize() + grow_max)
		}
		memcpy(handler.GetBuffer().GetData()+handler.GetBuffer().GetUsed(), buf.GetData(), buf.GetUsed())
		handler.GetBuffer().SetUsed(handler.GetBuffer().GetUsed() + buf.GetUsed())

		/* chunked buffering */

		if handler.GetSize() != 0 && handler.GetBuffer().GetUsed() >= handler.GetSize() {

			/* store away errors and/or any intermediate output */

			if OutputGlobals.GetRunning() != nil {
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

/* }}} */

func PhpOutputHandlerOp(handler *PhpOutputHandler, context *PhpOutputContext) PhpOutputHandlerStatusT {
	var status PhpOutputHandlerStatusT
	var original_op int = context.GetOp()
	if PhpOutputLockError(context.GetOp()) != 0 {

		/* fatal error */

		return PHP_OUTPUT_HANDLER_FAILURE

		/* fatal error */

	}

	/* storable? */

	if PhpOutputHandlerAppend(handler, &context.in) != 0 && context.GetOp() == 0 {
		context.SetOp(original_op)
		return PHP_OUTPUT_HANDLER_NO_DATA
	} else {

		/* need to start? */

		if (handler.GetFlags() & 0x1000) == 0 {
			context.SetOp(context.GetOp() | 0x1)
		}
		OutputGlobals.SetRunning(handler)
		if (handler.GetFlags() & 0x1) != 0 {
			var retval zend.Zval
			var ob_data zend.Zval
			var ob_mode zend.Zval
			var __z *zend.Zval = &ob_data
			var __s *zend.ZendString = zend.ZendStringInit(handler.GetBuffer().GetData(), handler.GetBuffer().GetUsed(), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			var __z *zend.Zval = &ob_mode
			__z.value.lval = zend.ZendLong(context.GetOp())
			__z.u1.type_info = 4
			zend.ZendFcallInfoArgn(&handler.func_.user.GetFci(), 2, &ob_data, &ob_mode)
			zend.ZvalPtrDtor(&ob_data)

			// #define PHP_OUTPUT_USER_SUCCESS(retval) ( ( Z_TYPE ( retval ) != IS_UNDEF ) && ! ( Z_TYPE ( retval ) == IS_FALSE ) )

			if zend.SUCCESS == zend.ZendFcallInfoCall(&handler.func_.user.GetFci(), &handler.func_.user.GetFcc(), &retval, nil) && (retval.u1.v.type_ != 0 && retval.u1.v.type_ != 2) {

				/* user handler may have returned TRUE */

				status = PHP_OUTPUT_HANDLER_NO_DATA
				if retval.u1.v.type_ != 2 && retval.u1.v.type_ != 3 {
					if &retval.u1.v.type_ != 6 {
						if &retval.u1.v.type_ != 6 {
							zend._convertToString(&retval)
						}
					}
					if retval.value.str.len_ != 0 {
						context.GetOut().SetData(zend._estrndup(retval.value.str.val, retval.value.str.len_))
						context.GetOut().SetUsed(retval.value.str.len_)
						context.GetOut().SetFree(1)
						status = PHP_OUTPUT_HANDLER_SUCCESS
					}
				}
			} else {

				/* call failed, pass internal buffer along */

				status = PHP_OUTPUT_HANDLER_FAILURE

				/* call failed, pass internal buffer along */

			}
			zend.ZendFcallInfoArgn(&handler.func_.user.GetFci(), 0)
			zend.ZvalPtrDtor(&retval)
		} else {
			PhpOutputContextFeed(context, handler.GetBuffer().GetData(), handler.GetBuffer().GetSize(), handler.GetBuffer().GetUsed(), 0)
			if zend.SUCCESS == handler.GetInternal()(&handler.opaq, context) {
				if context.GetOut().GetUsed() != 0 {
					status = PHP_OUTPUT_HANDLER_SUCCESS
				} else {
					status = PHP_OUTPUT_HANDLER_NO_DATA
				}
			} else {
				status = PHP_OUTPUT_HANDLER_FAILURE
			}
		}
		handler.SetFlags(handler.GetFlags() | 0x1000)
		OutputGlobals.SetRunning(nil)
	}
	switch status {
	case PHP_OUTPUT_HANDLER_FAILURE:

		/* disable this handler */

		handler.SetFlags(handler.GetFlags() | 0x2000)

		/* discard any output */

		if context.GetOut().GetData() != nil && context.GetOut().GetFree() != 0 {
			zend._efree(context.GetOut().GetData())
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
		handler.SetFlags(handler.GetFlags() | 0x4000)
		break
	}
	context.SetOp(original_op)
	return status
}

/* }}} */

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

	if OutputGlobals.GetActive() != nil && g.Assign(&obh_cnt, zend.ZendStackCount(&(OutputGlobals.GetHandlers()))) {
		context.GetIn().SetData((*byte)(str))
		context.GetIn().SetUsed(len_)
		if obh_cnt > 1 {
			zend.ZendStackApplyWithArgument(&(OutputGlobals.GetHandlers()), 1, PhpOutputStackApplyOp, &context)
		} else if g.Assign(&active, zend.ZendStackTop(&(OutputGlobals.GetHandlers()))) && ((*active).GetFlags()&0x2000) == 0 {
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
		if (OutputGlobals.GetFlags() & 0x2) == 0 {
			sapi_module.GetUbWrite()(context.GetOut().GetData(), context.GetOut().GetUsed())
			if (OutputGlobals.GetFlags() & 0x1) != 0 {
				SapiFlush()
			}
			OutputGlobals.SetFlags(OutputGlobals.GetFlags() | 0x8)
		}
	}
	PhpOutputContextDtor(&context)
}

/* }}} */

func PhpOutputStackApplyOp(h any, c any) int {
	var was_disabled int
	var status PhpOutputHandlerStatusT
	var handler *PhpOutputHandler = *((**PhpOutputHandler)(h))
	var context *PhpOutputContext = (*PhpOutputContext)(c)
	if g.Assign(&was_disabled, handler.GetFlags()&0x2000) {
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

/* }}} */

func PhpOutputStackApplyClean(h any, c any) int {
	var handler *PhpOutputHandler = *((**PhpOutputHandler)(h))
	var context *PhpOutputContext = (*PhpOutputContext)(c)
	handler.GetBuffer().SetUsed(0)
	PhpOutputHandlerOp(handler, context)
	PhpOutputContextReset(context)
	return 0
}

/* }}} */

func PhpOutputStackApplyList(h any, z any) int {
	var handler *PhpOutputHandler = *((**PhpOutputHandler)(h))
	var array *zend.Zval = (*zend.Zval)(z)
	zend.AddNextIndexStr(array, zend.ZendStringCopy(handler.GetName()))
	return 0
}

/* }}} */

func PhpOutputStackApplyStatus(h any, z any) int {
	var handler *PhpOutputHandler = *((**PhpOutputHandler)(h))
	var arr zend.Zval
	var array *zend.Zval = (*zend.Zval)(z)
	zend.AddNextIndexZval(array, PhpOutputHandlerStatus(handler, &arr))
	return 0
}

/* {{{ static zval *php_output_handler_status(php_output_handler *handler, zval *entry)
 * Returns an array with the status of the output handler */

func PhpOutputHandlerStatus(handler *PhpOutputHandler, entry *zend.Zval) *zend.Zval {
	assert(entry != nil)
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = entry
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	zend.AddAssocStrEx(entry, "name", strlen("name"), zend.ZendStringCopy(handler.GetName()))
	zend.AddAssocLongEx(entry, "type", strlen("type"), zend_long(handler.GetFlags()&0xf))
	zend.AddAssocLongEx(entry, "flags", strlen("flags"), zend.ZendLong(handler.GetFlags()))
	zend.AddAssocLongEx(entry, "level", strlen("level"), zend.ZendLong(handler.GetLevel()))
	zend.AddAssocLongEx(entry, "chunk_size", strlen("chunk_size"), zend.ZendLong(handler.GetSize()))
	zend.AddAssocLongEx(entry, "buffer_size", strlen("buffer_size"), zend.ZendLong(handler.GetBuffer().GetSize()))
	zend.AddAssocLongEx(entry, "buffer_used", strlen("buffer_used"), zend.ZendLong(handler.GetBuffer().GetUsed()))
	return entry
}

/* }}} */

func PhpOutputStackPop(flags int) int {
	var context PhpOutputContext
	var current **PhpOutputHandler
	var orphan **PhpOutputHandler = OutputGlobals.GetActive()
	if orphan == nil {
		if (flags & 0x100) == 0 {
			PhpErrorDocref("ref.outcontrol", 1<<3, "failed to %s buffer. No buffer to %s", g.Cond((flags&0x10) != 0, "discard", "send"), g.Cond((flags&0x10) != 0, "discard", "send"))
		}
		return 0
	} else if (flags&0x1) == 0 && (orphan.flags&0x40) == 0 {
		if (flags & 0x100) == 0 {
			PhpErrorDocref("ref.outcontrol", 1<<3, "failed to %s buffer of %s (%d)", g.Cond((flags&0x10) != 0, "discard", "send"), orphan.name.val, orphan.level)
		}
		return 0
	} else {
		PhpOutputContextInit(&context, 0x8)

		/* don't run the output handler if it's disabled */

		if (orphan.flags & 0x2000) == 0 {

			/* didn't it start yet? */

			if (orphan.flags & 0x1000) == 0 {
				context.SetOp(context.GetOp() | 0x1)
			}

			/* signal that we're cleaning up */

			if (flags & 0x10) != 0 {
				context.SetOp(context.GetOp() | 0x2)
			}
			PhpOutputHandlerOp(orphan, &context)
		}

		/* pop it off the stack */

		zend.ZendStackDelTop(&(OutputGlobals.GetHandlers()))
		if g.Assign(&current, zend.ZendStackTop(&(OutputGlobals.GetHandlers()))) {
			OutputGlobals.SetActive(*current)
		} else {
			OutputGlobals.SetActive(nil)
		}

		/* pass output along */

		if context.GetOut().GetData() != nil && context.GetOut().GetUsed() != 0 && (flags&0x10) == 0 {
			PhpOutputWrite(context.GetOut().GetData(), context.GetOut().GetUsed())
		}

		/* destroy the handler (after write!) */

		PhpOutputHandlerFree(&orphan)
		PhpOutputContextDtor(&context)
		return 1
	}
}

/* }}} */

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

/* }}} */

func PhpOutputHandlerDefaultFunc(handler_context *any, output_context *PhpOutputContext) int {
	PhpOutputContextPass(output_context)
	return zend.SUCCESS
}

/* }}} */

func PhpOutputHandlerDevnullFunc(handler_context *any, output_context *PhpOutputContext) int {
	return zend.SUCCESS
}

/* }}} */

func ZifObStart(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var output_handler *zend.Zval = nil
	var chunk_size zend.ZendLong = 0
	var flags zend.ZendLong = 0x70
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "|zll", &output_handler, &chunk_size, &flags) == zend.FAILURE {
		return
	}
	if chunk_size < 0 {
		chunk_size = 0
	}
	if PhpOutputStartUser(output_handler, chunk_size, flags) == zend.FAILURE {
		PhpErrorDocref("ref.outcontrol", 1<<3, "failed to create buffer")
		return_value.u1.type_info = 2
		return
	}
	return_value.u1.type_info = 3
	return
}

/* }}} */

func ZifObFlush(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if OutputGlobals.GetActive() == nil {
		PhpErrorDocref("ref.outcontrol", 1<<3, "failed to flush buffer. No buffer to flush")
		return_value.u1.type_info = 2
		return
	}
	if zend.SUCCESS != PhpOutputFlush() {
		PhpErrorDocref("ref.outcontrol", 1<<3, "failed to flush buffer of %s (%d)", OutputGlobals.GetActive().GetName().val, OutputGlobals.GetActive().GetLevel())
		return_value.u1.type_info = 2
		return
	}
	return_value.u1.type_info = 3
	return
}

/* }}} */

func ZifObClean(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if OutputGlobals.GetActive() == nil {
		PhpErrorDocref("ref.outcontrol", 1<<3, "failed to delete buffer. No buffer to delete")
		return_value.u1.type_info = 2
		return
	}
	if zend.SUCCESS != PhpOutputClean() {
		PhpErrorDocref("ref.outcontrol", 1<<3, "failed to delete buffer of %s (%d)", OutputGlobals.GetActive().GetName().val, OutputGlobals.GetActive().GetLevel())
		return_value.u1.type_info = 2
		return
	}
	return_value.u1.type_info = 3
	return
}

/* }}} */

func ZifObEndFlush(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if OutputGlobals.GetActive() == nil {
		PhpErrorDocref("ref.outcontrol", 1<<3, "failed to delete and flush buffer. No buffer to delete or flush")
		return_value.u1.type_info = 2
		return
	}
	if zend.SUCCESS == PhpOutputEnd() {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func ZifObEndClean(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if OutputGlobals.GetActive() == nil {
		PhpErrorDocref("ref.outcontrol", 1<<3, "failed to delete buffer. No buffer to delete")
		return_value.u1.type_info = 2
		return
	}
	if zend.SUCCESS == PhpOutputDiscard() {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func ZifObGetFlush(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if PhpOutputGetContents(return_value) == zend.FAILURE {
		PhpErrorDocref("ref.outcontrol", 1<<3, "failed to delete and flush buffer. No buffer to delete or flush")
		return_value.u1.type_info = 2
		return
	}
	if zend.SUCCESS != PhpOutputEnd() {
		PhpErrorDocref("ref.outcontrol", 1<<3, "failed to delete buffer of %s (%d)", OutputGlobals.GetActive().GetName().val, OutputGlobals.GetActive().GetLevel())
	}
}

/* }}} */

func ZifObGetClean(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if OutputGlobals.GetActive() == nil {
		return_value.u1.type_info = 2
		return
	}
	if PhpOutputGetContents(return_value) == zend.FAILURE {
		PhpErrorDocref("ref.outcontrol", 1<<3, "failed to delete buffer. No buffer to delete")
		return_value.u1.type_info = 2
		return
	}
	if zend.SUCCESS != PhpOutputDiscard() {
		PhpErrorDocref("ref.outcontrol", 1<<3, "failed to delete buffer of %s (%d)", OutputGlobals.GetActive().GetName().val, OutputGlobals.GetActive().GetLevel())
	}
}

/* }}} */

func ZifObGetContents(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if PhpOutputGetContents(return_value) == zend.FAILURE {
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

func ZifObGetLevel(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = PhpOutputGetLevel()
	__z.u1.type_info = 4
	return
}

/* }}} */

func ZifObGetLength(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if PhpOutputGetLength(return_value) == zend.FAILURE {
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

func ZifObListHandlers(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	if OutputGlobals.GetActive() == nil {
		return
	}
	zend.ZendStackApplyWithArgument(&(OutputGlobals.GetHandlers()), 2, PhpOutputStackApplyList, return_value)
}

/* }}} */

func ZifObGetStatus(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var full_status zend.ZendBool = 0
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "|b", &full_status) == zend.FAILURE {
		return
	}
	if OutputGlobals.GetActive() == nil {
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		return
	}
	if full_status != 0 {
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		zend.ZendStackApplyWithArgument(&(OutputGlobals.GetHandlers()), 2, PhpOutputStackApplyStatus, return_value)
	} else {
		PhpOutputHandlerStatus(OutputGlobals.GetActive(), return_value)
	}
}

/* }}} */

func ZifObImplicitFlush(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var flag zend.ZendLong = 1
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "|l", &flag) == zend.FAILURE {
		return
	}
	PhpOutputSetImplicitFlush(flag)
}

/* }}} */

func ZifOutputResetRewriteVars(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if standard.PhpUrlScannerResetVars() == zend.SUCCESS {
		return_value.u1.type_info = 3
		return
	} else {
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

func ZifOutputAddRewriteVar(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var name *byte
	var value *byte
	var name_len int
	var value_len int
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "ss", &name, &name_len, &value, &value_len) == zend.FAILURE {
		return
	}
	if standard.PhpUrlScannerAddVar(name, name_len, value, value_len, 1) == zend.SUCCESS {
		return_value.u1.type_info = 3
		return
	} else {
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */
