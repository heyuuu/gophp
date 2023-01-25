// <<generate>>

package zend

// Source: <Zend/zend_portability.h>

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
   |          Dmitry Stogov <zeev@php.net>                                |
   +----------------------------------------------------------------------+
*/

/*
 * general definitions
 */

// failed # include "../TSRM/TSRM.h"

/* GCC x.y.z supplies __GNUC__ = x and __GNUC_MINOR__ = y */

/* Compatibility with non-clang compilers */

/* Only use this macro if you know for sure that all of the switches values
   are covered by its case statements */

/* all HAVE_XXX test have to be after the include of zend_config above */

/* AIX requires this to be the first thing in the file.  */

/* We always define a __special__  function, even if there's a macro or expression we could
 * alias, so that using it in contexts where we can't make function calls
 * won't fail to compile on some machines and not others.
 */

/* excpt.h on Digital Unix 4.0 defines function_table */

/* This check should only be used on network socket, not file descriptors */

/* va_copy() is __va_copy() in old gcc versions.
 * According to the autoconf manual, using
 * memcpy(&dst, &src, sizeof(va_list))
 * gives maximum portability. */

/* Intrinsics macros start. */

/* Do not use for conditional declaration of API functions! */

/* Do not use for conditional declaration of API functions! */

/* Do not use for conditional declaration of API functions! */

/* Intrinsics macros end. */

/* On CPU with few registers, it's cheaper to reload value then use spill slot */
