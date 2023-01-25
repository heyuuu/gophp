// <<generate>>

package core

// Source: <main/php_output.h>

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
   | Author: Michael Wallner <mike@php.net>                               |
   +----------------------------------------------------------------------+
*/

// #define PHP_OUTPUT_H

const PHP_OUTPUT_NEWAPI = 1

/* handler ops */

const PHP_OUTPUT_HANDLER_WRITE = 0x0
const PHP_OUTPUT_HANDLER_START = 0x1
const PHP_OUTPUT_HANDLER_CLEAN = 0x2
const PHP_OUTPUT_HANDLER_FLUSH = 0x4
const PHP_OUTPUT_HANDLER_FINAL = 0x8
const PHP_OUTPUT_HANDLER_CONT = PHP_OUTPUT_HANDLER_WRITE
const PHP_OUTPUT_HANDLER_END = PHP_OUTPUT_HANDLER_FINAL

/* handler types */

const PHP_OUTPUT_HANDLER_INTERNAL = 0x0
const PHP_OUTPUT_HANDLER_USER = 0x1

/* handler ability flags */

const PHP_OUTPUT_HANDLER_CLEANABLE = 0x10
const PHP_OUTPUT_HANDLER_FLUSHABLE = 0x20
const PHP_OUTPUT_HANDLER_REMOVABLE = 0x40
const PHP_OUTPUT_HANDLER_STDFLAGS = 0x70

/* handler status flags */

const PHP_OUTPUT_HANDLER_STARTED = 0x1000
const PHP_OUTPUT_HANDLER_DISABLED = 0x2000
const PHP_OUTPUT_HANDLER_PROCESSED = 0x4000

/* handler op return values */

type PhpOutputHandlerStatusT = int

const (
	PHP_OUTPUT_HANDLER_FAILURE = iota
	PHP_OUTPUT_HANDLER_SUCCESS
	PHP_OUTPUT_HANDLER_NO_DATA
)

/* php_output_stack_pop() flags */

const PHP_OUTPUT_POP_TRY = 0x0
const PHP_OUTPUT_POP_FORCE = 0x1
const PHP_OUTPUT_POP_DISCARD = 0x10
const PHP_OUTPUT_POP_SILENT = 0x100

/* real global flags */

const PHP_OUTPUT_IMPLICITFLUSH = 0x1
const PHP_OUTPUT_DISABLED = 0x2
const PHP_OUTPUT_WRITTEN = 0x4
const PHP_OUTPUT_SENT = 0x8

/* supplementary flags for php_output_get_status() */

const PHP_OUTPUT_ACTIVE = 0x10
const PHP_OUTPUT_LOCKED = 0x20

/* output layer is ready to use */

const PHP_OUTPUT_ACTIVATED = 0x100000

/* handler hooks */

type PhpOutputHandlerHookT = int

const (
	PHP_OUTPUT_HANDLER_HOOK_GET_OPAQ = iota
	PHP_OUTPUT_HANDLER_HOOK_GET_FLAGS
	PHP_OUTPUT_HANDLER_HOOK_GET_LEVEL
	PHP_OUTPUT_HANDLER_HOOK_IMMUTABLE
	PHP_OUTPUT_HANDLER_HOOK_DISABLE
	PHP_OUTPUT_HANDLER_HOOK_LAST
)

func PHP_OUTPUT_HANDLER_INITBUF_SIZE(s int) __auto__ {
	if s > 1 {
		return s + PHP_OUTPUT_HANDLER_ALIGNTO_SIZE - s%PHP_OUTPUT_HANDLER_ALIGNTO_SIZE
	} else {
		return PHP_OUTPUT_HANDLER_DEFAULT_SIZE
	}
}

const PHP_OUTPUT_HANDLER_ALIGNTO_SIZE = 0x1000
const PHP_OUTPUT_HANDLER_DEFAULT_SIZE = 0x4000

/* old-style, stateless callback */

type PhpOutputHandlerFuncT func(output *byte, output_len int, handled_output **byte, handled_output_len *int, mode int)

/* new-style, opaque context callback */

type PhpOutputHandlerContextFuncT func(handler_context *any, output_context *PhpOutputContext) int

/* output handler context dtor */

type PhpOutputHandlerContextDtorT func(opaq any)

/* conflict check callback */

type PhpOutputHandlerConflictCheckT func(handler_name *byte, handler_name_len int) int

/* ctor for aliases */

type PhpOutputHandlerAliasCtorT func(handler_name *byte, handler_name_len int, chunk_size int, flags int) *PhpOutputHandler

/* there should not be a need to use OG() from outside of output.c */

func OG(v **PhpOutputHandler) __auto__ { return OutputGlobals.v }

/* convenience macros */

func PHPWRITE(str *byte, str_len int) int { return PhpOutputWrite(str, str_len) }
func PHPWRITE_H(str *byte, str_len int) int {
	return PhpOutputWriteUnbuffered(str, str_len)
}
func PUTC(c __auto__) int { return PhpOutputWrite((*byte)(&c), 1) }
func PUTC_H(c __auto__) int {
	return PhpOutputWriteUnbuffered((*byte)(&c), 1)
}
func PUTS(str string) {
	var __str *byte = str
	PhpOutputWrite(__str, strlen(__str))
}
func PUTS_H(str *byte) {
	var __str *byte = str
	PhpOutputWriteUnbuffered(__str, strlen(__str))
}
func PhpOutputTearup() {
	PhpOutputStartup()
	PhpOutputActivate()
}
func PhpOutputTeardown() {
	PhpOutputEndAll()
	PhpOutputDeactivate()
	PhpOutputShutdown()
}

/* MINIT */

/* MSHUTDOWN */

/* RINIT */

/* RSHUTDOWN */
