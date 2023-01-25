// <<generate>>

package zend

// Source: <Zend/zend_compile.h>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

/* On 64-bit systems less optimal, but more compact VM code leads to better
 * performance. So on 32-bit systems we use absolute addresses for jump
 * targets and constants, but on 64-bit systems realtive 32-bit offsets */

/* Temporarily defined here, to avoid header ordering issues */

/* Compilation context that is different for each file, but shared between op arrays. */

/* Compilation context that is different for each op array. */

/* Class, property and method flags                  class|meth.|prop.|const*/

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/*                                                        |     |     |     */

/* call through internal function handler. e.g. Closure::invoke() */

/* arg_info for internal functions */

/* arg_info for user functions */

/* the following structure repeats the layout of zend_internal_arg_info,
 * but its fields have different meaning. It's used as the first element of
 * arg_info array to define properties __special__  of internal functions.
 * It's also used for the return type.
 */

/* zend_internal_function_handler */

/* Top 16 bits of Z_TYPE_INFO(EX(This)) are used as call_info flags */

/* run-time jump target */

/* convert jump target from compile-time to run-time */

/* convert jump target back from run-time to compile-time */

/* constant-time constant */

/* At run-time, constants are allocated together with op_array->opcodes
 * and addressed relatively to current opline.
 */

/* convert constant from compile-time to run-time */

/* convert constant back from run-time to compile-time */

/* Used during AST construction */

/* parser-driven code generators */

/* helper functions in zend_language_scanner.l */

/* BEGIN: OPCODES */

/* END: OPCODES */

/* var status for backpatching */

/* A quick check (type == ZEND_USER_FUNCTION || type == ZEND_EVAL_CODE) */

/* global/local fetches */

/* Only one of these can ever be in use */

/* Quick API to check first 12 arguments */

/* Attribute for ternary inside parentheses */

/* For "use" AST nodes and the seen symbol table */

/* All increment opcodes are even (decrement are odd) */

/* Pseudo-opcodes that are used only temporarily during compilation */

/* The following constants may be combined in CG(compiler_options)
 * to change the default compiler behavior */

/* call op_array handler of extendions */

/* generate ZEND_INIT_FCALL_BY_NAME for internal functions instead of ZEND_INIT_FCALL */

/* don't perform early binding for classes inherited form internal ones;
 * in namespaces assume that internal class that doesn't exist at compile-time
 * may apper in run-time */

/* generate ZEND_DECLARE_CLASS_DELAYED opcode to delay early binding */

/* disable constant substitution at compile-time */

/* disable usage of builtin instruction for strlen() */

/* disable substitution of persistent constants at compile-time */

/* generate ZEND_INIT_FCALL_BY_NAME for userland functions instead of ZEND_INIT_FCALL */

/* force ZEND_ACC_USE_GUARDS for all classes */

/* disable builtin special case function calls */

/* result of compilation may be stored in file cache */

/* ignore functions and classes declared in other files */

/* this flag is set when compiler invoked by opcache_compile_file() */

/* this flag is set when compiler invoked during preloading */

/* disable jumptable optimization for switch statements */

/* this flag is set when compiler invoked during preloading in separate process */

/* The default value for CG(compiler_options) */

/* The default value for CG(compiler_options) during eval() */

// Source: <Zend/zend_compile.c>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   |          Nikita Popov <nikic@php.net>                                |
   +----------------------------------------------------------------------+
*/

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* Common part of zend_add_literal and zend_append_individual_literal */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* Propagate refs used on leaf elements to the surrounding list() structures. */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}}*/

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */
