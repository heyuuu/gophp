// <<generate>>

package zend

// Source: <Zend/zend_globals.h>

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

/* Define ZTS if you want a thread-safe Zend */

const SYMTABLE_CACHE_SIZE = 32

/* excpt.h on Digital Unix 4.0 defines function_table */

type ZendVmStack *_zendVmStack

const EG_FLAGS_INITIAL = 0
const EG_FLAGS_IN_SHUTDOWN = 1 << 0
const EG_FLAGS_OBJECT_STORE_NO_REUSE = 1 << 1
const EG_FLAGS_IN_RESOURCE_SHUTDOWN = 1 << 2

type ZendPhpScannerEvent = int

const (
	ON_TOKEN = iota
	ON_FEEDBACK
	ON_STOP
)
