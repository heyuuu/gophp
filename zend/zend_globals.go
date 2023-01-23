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

// #define ZEND_GLOBALS_H

// # include < setjmp . h >

// # include "zend_globals_macros.h"

// # include "zend_stack.h"

// # include "zend_ptr_stack.h"

// # include "zend_hash.h"

// # include "zend_llist.h"

// # include "zend_objects.h"

// # include "zend_objects_API.h"

// # include "zend_modules.h"

// # include "zend_float.h"

// # include "zend_multibyte.h"

// # include "zend_multiply.h"

// # include "zend_arena.h"

/* Define ZTS if you want a thread-safe Zend */

// #define SYMTABLE_CACHE_SIZE       32

// # include "zend_compile.h"

/* excpt.h on Digital Unix 4.0 defines function_table */

type ZendVmStack *_zendVmStack

// #define EG_FLAGS_INITIAL       ( 0 )

// #define EG_FLAGS_IN_SHUTDOWN       ( 1 << 0 )

// #define EG_FLAGS_OBJECT_STORE_NO_REUSE       ( 1 << 1 )

// #define EG_FLAGS_IN_RESOURCE_SHUTDOWN       ( 1 << 2 )

type ZendPhpScannerEvent = int

const (
	ON_TOKEN = iota
	ON_FEEDBACK
	ON_STOP
)
