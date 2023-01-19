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

// #define PHP_OUTPUT_NEWAPI       1

/* handler ops */

// #define PHP_OUTPUT_HANDLER_WRITE       0x00

// #define PHP_OUTPUT_HANDLER_START       0x01

// #define PHP_OUTPUT_HANDLER_CLEAN       0x02

// #define PHP_OUTPUT_HANDLER_FLUSH       0x04

// #define PHP_OUTPUT_HANDLER_FINAL       0x08

// #define PHP_OUTPUT_HANDLER_CONT       PHP_OUTPUT_HANDLER_WRITE

// #define PHP_OUTPUT_HANDLER_END       PHP_OUTPUT_HANDLER_FINAL

/* handler types */

// #define PHP_OUTPUT_HANDLER_INTERNAL       0x0000

// #define PHP_OUTPUT_HANDLER_USER       0x0001

/* handler ability flags */

// #define PHP_OUTPUT_HANDLER_CLEANABLE       0x0010

// #define PHP_OUTPUT_HANDLER_FLUSHABLE       0x0020

// #define PHP_OUTPUT_HANDLER_REMOVABLE       0x0040

// #define PHP_OUTPUT_HANDLER_STDFLAGS       0x0070

/* handler status flags */

// #define PHP_OUTPUT_HANDLER_STARTED       0x1000

// #define PHP_OUTPUT_HANDLER_DISABLED       0x2000

// #define PHP_OUTPUT_HANDLER_PROCESSED       0x4000

/* handler op return values */

type PhpOutputHandlerStatusT = int

const (
	PHP_OUTPUT_HANDLER_FAILURE = iota
	PHP_OUTPUT_HANDLER_SUCCESS
	PHP_OUTPUT_HANDLER_NO_DATA
)

/* php_output_stack_pop() flags */

// #define PHP_OUTPUT_POP_TRY       0x000

// #define PHP_OUTPUT_POP_FORCE       0x001

// #define PHP_OUTPUT_POP_DISCARD       0x010

// #define PHP_OUTPUT_POP_SILENT       0x100

/* real global flags */

// #define PHP_OUTPUT_IMPLICITFLUSH       0x01

// #define PHP_OUTPUT_DISABLED       0x02

// #define PHP_OUTPUT_WRITTEN       0x04

// #define PHP_OUTPUT_SENT       0x08

/* supplementary flags for php_output_get_status() */

// #define PHP_OUTPUT_ACTIVE       0x10

// #define PHP_OUTPUT_LOCKED       0x20

/* output layer is ready to use */

// #define PHP_OUTPUT_ACTIVATED       0x100000

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

// #define PHP_OUTPUT_HANDLER_INITBUF_SIZE(s) ( ( ( s ) > 1 ) ? ( s ) + PHP_OUTPUT_HANDLER_ALIGNTO_SIZE - ( ( s ) % ( PHP_OUTPUT_HANDLER_ALIGNTO_SIZE ) ) : PHP_OUTPUT_HANDLER_DEFAULT_SIZE )

// #define PHP_OUTPUT_HANDLER_ALIGNTO_SIZE       0x1000

// #define PHP_OUTPUT_HANDLER_DEFAULT_SIZE       0x4000

// @type PhpOutputBuffer struct

// @type PhpOutputContext struct

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

// @type PhpOutputHandlerUserFuncT struct

// @type PhpOutputHandler struct

// @type ZendOutputGlobals struct

/* there should not be a need to use OG() from outside of output.c */

// #define OG(v) ( output_globals . v )

/* convenience macros */

// #define PHPWRITE(str,str_len) php_output_write ( ( str ) , ( str_len ) )

// #define PHPWRITE_H(str,str_len) php_output_write_unbuffered ( ( str ) , ( str_len ) )

// #define PUTC(c) php_output_write ( ( const char * ) & ( c ) , 1 )

// #define PUTC_H(c) php_output_write_unbuffered ( ( const char * ) & ( c ) , 1 )

// #define PUTS(str) do { const char * __str = ( str ) ; php_output_write ( __str , strlen ( __str ) ) ; } while ( 0 )

// #define PUTS_H(str) do { const char * __str = ( str ) ; php_output_write_unbuffered ( __str , strlen ( __str ) ) ; } while ( 0 )

// #define php_output_tearup() php_output_startup ( ) ; php_output_activate ( )

// #define php_output_teardown() php_output_end_all ( ) ; php_output_deactivate ( ) ; php_output_shutdown ( )

/* MINIT */

/* MSHUTDOWN */

/* RINIT */

/* RSHUTDOWN */
