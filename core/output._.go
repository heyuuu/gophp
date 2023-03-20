// <<generate>>

package core

import (
	"sik/zend/types"
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

const PHP_OUTPUT_DEBUG = 0
const PHP_OUTPUT_NOINLINE = 0

var OutputGlobals ZendOutputGlobals
var PhpOutputDefaultHandlerName []byte = "default output handler"
var PhpOutputDevnullHandlerName []byte = "null output handler"

/* {{{ aliases, conflict and reverse conflict hash tables */

var PhpOutputHandlerAliases types.Array
var PhpOutputHandlerConflicts types.Array
var PhpOutputHandlerReverseConflicts types.Array
var PhpOutputDirect func(str *byte, str_len int) int = PhpOutputStderr

/* {{{ void php_output_startup(void)
 * Set up module globals and initialize the conflict and reverse conflict hash tables */

/* {{{ SUCCESS|FAILURE php_output_end(void)
 * Finalizes the most recent output handler at pops it off the stack if the handler is removable */

/* {{{ SUCCESS|FAILURE php_output_get_length(zval *z)
 * Get the length of the active output handlers buffer */

/* {{{ static zval *php_output_handler_status(php_output_handler *handler, zval *entry)
 * Returns an array with the status of the output handler */
