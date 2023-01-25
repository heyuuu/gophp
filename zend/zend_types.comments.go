// <<generate>>

package zend

// Source: <Zend/zend_types.h>

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
   |          Dmitry Stogov <dmitry@php.net>                              |
   |          Xinchen Hui <xinchen.h@zend.com>                            |
   +----------------------------------------------------------------------+
*/

/*
 * zend_type - is an abstraction layer to represent information about type hint.
 * It shouldn't be used directly. Only through ZEND_TYPE_* macros.
 *
 * ZEND_TYPE_IS_SET()     - checks if type-hint exists
 * ZEND_TYPE_IS_CODE()    - checks if type-hint refer to standard type
 * ZEND_TYPE_IS_CLASS()   - checks if type-hint refer to some class
 * ZEND_TYPE_IS_CE()      - checks if type-hint refer to some class by zend_class_entry *
 * ZEND_TYPE_IS_NAME()    - checks if type-hint refer to some class by zend_string *
 *
 * ZEND_TYPE_NAME()       - returns referenced class name
 * ZEND_TYPE_CE()         - returns referenced class entry
 * ZEND_TYPE_CODE()       - returns standard type code (e.g. IS_LONG, _IS_BOOL)
 *
 * ZEND_TYPE_ALLOW_NULL() - checks if NULL is allowed
 *
 * ZEND_TYPE_ENCODE() and ZEND_TYPE_ENCODE_CLASS() should be used for
 * construction.
 */

/*
 * HashTable Data Layout
 * =====================
 *
 *                 +=============================+
 *                 | HT_HASH(ht, ht->nTableMask) |
 *                 | ...                         |
 *                 | HT_HASH(ht, -1)             |
 *                 +-----------------------------+
 * ht->arData ---> | Bucket[0]                   |
 *                 | ...                         |
 *                 | Bucket[ht->nTableSize-1]    |
 *                 +=============================+
 */

/* regular data types */

/* constant expressions */

/* internal types */

/* fake types used only for type hinting (Z_TYPE(zv) can not use them) */

/* we should never set just Z_TYPE, we should set Z_TYPE_INFO */

/* zval_gc_flags(zval.value->gc.u.type_info) (common flags) */

/* zval.u1.v.type_flags */

/* This optimized version assumes that we have a single "type_flag" */

/* extended types */

/* string flags (zval.value->gc.u.flags) */

/* array flags */

/* object flags (zval.value->gc.u.flags) */

/* Recursion protection macros must be used only for arrays and objects */

/* All data types < IS_STRING have their constructor/destructors skipped */

/* This optimized version assumes that we have a single "type_flag" */

/* deprecated: (COPYABLE is the same as IS_ARRAY) */

/* deprecated: (IMMUTABLE is the same as IS_ARRAY && !REFCOUNTED) */

/* the following Z_OPT_* macros make better code when Z_TYPE_INFO accessed before */

/* deprecated: (COPYABLE is the same as IS_ARRAY) */

/* ZVAL_COPY_OR_DUP() should be used instead of ZVAL_COPY() and ZVAL_DUP()
 * in all places where the source may be a persistent zval.
 */

/* Properties store a flag distinguishing unset and unintialized properties
 * (both use IS_UNDEF type) in the Z_EXTRA space. As such we also need to copy
 * the Z_EXTRA space when copying property default values etc. We define separate __special__
 * macros for this purpose, so this workaround is easier to remove in the future. */
