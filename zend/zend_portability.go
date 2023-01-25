// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/core"
)

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

// #define ZEND_PORTABILITY_H

// #define BEGIN_EXTERN_C()

// #define END_EXTERN_C()

/*
 * general definitions
 */

// # include < zend_config . h >

const ZEND_PATHS_SEPARATOR = ':'

// failed # include "../TSRM/TSRM.h"

// # include < stdio . h >

// # include < assert . h >

// # include < math . h >

// # include < stdarg . h >

// # include < stddef . h >

// # include < dlfcn . h >

// # include < limits . h >

// # include < alloca . h >

// # include "zend_range_check.h"

/* GCC x.y.z supplies __GNUC__ = x and __GNUC_MINOR__ = y */

const ZEND_GCC_VERSION = 0

/* Compatibility with non-clang compilers */

func __hasAttribute(x __auto__) int { return 0 }
func __hasBuiltin(x __auto__) int   { return 0 }
func __hasFeature(x __auto__) int   { return 0 }

// #define ZEND_ASSUME(c)

func ZEND_ASSERT(c bool) {}

/* Only use this macro if you know for sure that all of the switches values
   are covered by its case statements */

// #define EMPTY_SWITCH_DEFAULT_CASE() default : ZEND_ASSUME ( 0 ) ; break ;

func ZEND_IGNORE_VALUE(x __auto__) { void(x) }
func ZendQuietWrite()              { ZEND_IGNORE_VALUE(write(__VA_ARGS__)) }

/* all HAVE_XXX test have to be after the include of zend_config above */

const RTLD_LAZY = 1
const RTLD_GLOBAL = 0
const PHP_RTLD_MODE = RTLD_LAZY

func DL_LOAD(libname *byte) __auto__ {
	return dlopen(libname, PHP_RTLD_MODE|RTLD_GLOBAL)
}

const DL_UNLOAD = dlclose
const DL_FETCH_SYMBOL = dlsym
const DL_ERROR = dlerror

// #define DL_HANDLE       void *

const ZEND_EXTENSIONS_SUPPORT = 1

/* AIX requires this to be the first thing in the file.  */

// #define ZEND_ATTRIBUTE_MALLOC

// #define ZEND_ATTRIBUTE_ALLOC_SIZE(X)

// #define ZEND_ATTRIBUTE_ALLOC_SIZE2(X,Y)

// #define ZEND_ATTRIBUTE_FORMAT(type,idx,first)

// #define ZEND_ATTRIBUTE_PTR_FORMAT(type,idx,first)

// #define ZEND_ATTRIBUTE_DEPRECATED

// #define ZEND_ATTRIBUTE_UNUSED

// #define ZEND_COLD

// #define ZEND_HOT

// #define ZEND_OPT_SIZE

// #define ZEND_OPT_SPEED

// #define ZEND_ATTRIBUTE_UNUSED_LABEL

// #define ZEND_ATTRIBUTE_COLD_LABEL

// #define ZEND_ATTRIBUTE_HOT_LABEL

// #define ZEND_FASTCALL

// #define ZEND_NORETURN

func ZEND_CONST_COND(_condition __auto__, _default __auto__) __auto__ { return _default }

// #define zend_always_inline       inline

// #define zend_never_inline

func EXPECTED(condition bool) __auto__   { return __builtin_expect(!!condition, 1) }
func UNEXPECTED(condition bool) __auto__ { return __builtin_expect(!!condition, 0) }

// #define XtOffset(p_type,field) ( ( zend_long ) ( ( ( char * ) ( & ( ( ( p_type ) NULL ) -> field ) ) ) - ( ( char * ) NULL ) ) )

// #define XtOffsetOf(s_type,field) XtOffset ( s_type * , field )

// #define ALLOCA_FLAG(name)

// #define SET_ALLOCA_FLAG(name)

func DoAlloca(p int, use_heap __auto__) any          { return Emalloc(p) }
func FreeAlloca(p any, use_heap __auto__)            { Efree(p) }
func SETJMP(a JMP_BUF) __auto__                      { return sigsetjmp(a, 0) }
func LONGJMP(a JMP_BUF, b ZEND_RESULT_CODE) __auto__ { return siglongjmp(a, b) }

const JMP_BUF = sigjmp_buf

// #define ZEND_FILE_LINE_D       void

// #define ZEND_FILE_LINE_DC

// #define ZEND_FILE_LINE_ORIG_D       void

// #define ZEND_FILE_LINE_ORIG_DC

// #define ZEND_FILE_LINE_RELAY_C

// #define ZEND_FILE_LINE_RELAY_CC

// #define ZEND_FILE_LINE_C

// #define ZEND_FILE_LINE_CC

// #define ZEND_FILE_LINE_EMPTY_C

// #define ZEND_FILE_LINE_EMPTY_CC

// #define ZEND_FILE_LINE_ORIG_RELAY_C

// #define ZEND_FILE_LINE_ORIG_RELAY_CC

// #define Z_DBG(expr)

const ZTS_V = 0
const LONG_MAX = 2147483647
const LONG_MIN = -LONG_MAX - 1
const MAX_LENGTH_OF_DOUBLE = 32

func MAX(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func MIN(a __auto__, b __auto__) __auto__ {
	if a < b {
		return a
	} else {
		return b
	}
}
func ZEND_BIT_TEST(bits []uint32, bit uint32) int {
	return bits[bit/(b.SizeOf("( bits ) [ 0 ]")*8)] >> (bit&b.SizeOf("( bits ) [ 0 ]")*8 - 1) & 1
}

/* We always define a __special__  function, even if there's a macro or expression we could
 * alias, so that using it in contexts where we can't make function calls
 * won't fail to compile on some machines and not others.
 */

func _zendGetInf() float64 { return HUGE_VAL }

const ZEND_INFINITY float64 = _zendGetInf()

func _zendGetNan() float64 { return 0.0 / 0.0 }

const ZEND_NAN float64 = _zendGetNan()

func ZEND_STRL(str string) {
	str
	b.SizeOf("str") - 1
}
func ZEND_STRS(str __auto__) {
	str
	b.SizeOf("str")
}
func ZEND_NORMALIZE_BOOL(n ZendLong) int {
	if n != 0 {
		if n < 0 {
			return -1
		} else {
			return 1
		}
	} else {
		return 0
	}
}
func ZEND_TRUTH(x __auto__) int {
	if x {
		return 1
	} else {
		return 0
	}
}
func ZEND_LOG_XOR(a __auto__, b __auto__) int { return ZEND_TRUTH(a) ^ ZEND_TRUTH(b) }

const ZEND_MAX_RESERVED_RESOURCES = 6

/* excpt.h on Digital Unix 4.0 defines function_table */

func ZEND_SECURE_ZERO(var_ __auto__, size __auto__) __auto__ { return core.ExplicitBzero(var_, size) }

/* This check should only be used on network socket, not file descriptors */

func ZEND_VALID_SOCKET(sock core.PhpSocketT) bool { return sock >= 0 }

/* va_copy() is __va_copy() in old gcc versions.
 * According to the autoconf manual, using
 * memcpy(&dst, &src, sizeof(va_list))
 * gives maximum portability. */

func VaCopy(dest ...any, src ...any) __auto__ {
	return memcpy(&dest, &src, b.SizeOf("va_list"))
}

/* Intrinsics macros start. */

/* Do not use for conditional declaration of API functions! */

// #define ZEND_INTRIN_SSSE3_FUNC_DECL(func)

/* Do not use for conditional declaration of API functions! */

// #define ZEND_INTRIN_SSE4_2_FUNC_DECL(func)

/* Do not use for conditional declaration of API functions! */

// #define ZEND_INTRIN_AVX2_FUNC_DECL(func)

/* Intrinsics macros end. */

// #define ZEND_SET_ALIGNED(alignment,decl) decl

func ZEND_SLIDE_TO_ALIGNED(alignment int, ptr __auto__) int {
	return zend_uintptr_t(ptr) + (alignment-1) & ^(alignment-1)
}
func ZEND_SLIDE_TO_ALIGNED16(ptr __auto__) int {
	return ZEND_SLIDE_TO_ALIGNED(Z_UL(16), ptr)
}
func ZEND_EXPAND_VA(code __auto__) __auto__ { return code }

/* On CPU with few registers, it's cheaper to reload value then use spill slot */

// #define ZEND_IGNORE_LEAKS_BEGIN()

// #define ZEND_IGNORE_LEAKS_END()
