// <<generate>>

package zend

import (
	b "sik/builtin"
)

// Source: <Zend/zend_gc.h>

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
   | Authors: David Wang <planetbeing@gmail.com>                          |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

var GcCollectCycles func() int

/* enable/disable automatic start of GC collection */

/* enable/disable possible root additions */

/* The default implementation of the gc_collect_cycles callback. */

// Source: <Zend/zend_gc.c>

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
   | Authors: David Wang <planetbeing@gmail.com>                          |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

const GC_BENCH = 0
const ZEND_GC_DEBUG = 0

/* GC_INFO layout */

const GC_ADDRESS = 0xfffff
const GC_COLOR = 0x300000
const GC_BLACK = 0x0
const GC_WHITE = 0x100000
const GC_GREY = 0x200000
const GC_PURPLE = 0x300000

/* Debug tracing */

/* GC_INFO access */

/* bit stealing tags for gc_root_buffer.ref */

const GC_BITS = 0x3
const GC_ROOT = 0x0
const GC_UNUSED = 0x1
const GC_GARBAGE = 0x2
const GC_DTOR_GARBAGE = 0x3

/* GC address conversion */

/* GC buffers */

const GC_INVALID = 0
const GC_FIRST_ROOT = 1
const GC_DEFAULT_BUF_SIZE uint32 = 16 * 1024
const GC_BUF_GROW_STEP uint32 = 128 * 1024
const GC_MAX_UNCOMPRESSED = 512 * 1024
const GC_MAX_BUF_SIZE = 0x40000000
const GC_THRESHOLD_DEFAULT = 10000
const GC_THRESHOLD_STEP = 10000
const GC_THRESHOLD_MAX = 1000000000
const GC_THRESHOLD_TRIGGER = 100

/* GC flags */

const GC_HAS_DESTRUCTORS = 1 << 0

/* unused buffers */

var GcGlobals ZendGcGlobals

const GC_STACK_SEGMENT_SIZE = (4096-ZEND_MM_OVERHEAD)/b.SizeOf("void *") - 2

/* Two-Finger compaction algorithm */
