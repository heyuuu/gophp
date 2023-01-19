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

// #define ZEND_PORTABILITY_H

// #define BEGIN_EXTERN_C()

// #define END_EXTERN_C()

/*
 * general definitions
 */

// # include < zend_config . h >

// #define ZEND_PATHS_SEPARATOR       ':'

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

// #define ZEND_GCC_VERSION       0

/* Compatibility with non-clang compilers */

// #define __has_attribute(x) 0

// #define __has_builtin(x) 0

// #define __has_feature(x) 0

// #define ZEND_ASSUME(c)

// #define ZEND_ASSERT(c) ZEND_ASSUME ( c )

/* Only use this macro if you know for sure that all of the switches values
   are covered by its case statements */

// #define EMPTY_SWITCH_DEFAULT_CASE() default : ZEND_ASSUME ( 0 ) ; break ;

// #define ZEND_IGNORE_VALUE(x) ( ( void ) ( x ) )

// #define zend_quiet_write() ZEND_IGNORE_VALUE ( write ( __VA_ARGS__ ) )

/* all HAVE_XXX test have to be after the include of zend_config above */

// #define RTLD_LAZY       1

// #define RTLD_GLOBAL       0

// #define PHP_RTLD_MODE       RTLD_LAZY

// #define DL_LOAD(libname) dlopen ( libname , PHP_RTLD_MODE | RTLD_GLOBAL )

// #define DL_UNLOAD       dlclose

// #define DL_FETCH_SYMBOL       dlsym

// #define DL_ERROR       dlerror

// #define DL_HANDLE       void *

// #define ZEND_EXTENSIONS_SUPPORT       1

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

// #define ZEND_CONST_COND(_condition,_default) ( _default )

// #define zend_always_inline       inline

// #define zend_never_inline

// #define EXPECTED(condition) __builtin_expect ( ! ! ( condition ) , 1 )

// #define UNEXPECTED(condition) __builtin_expect ( ! ! ( condition ) , 0 )

// #define XtOffset(p_type,field) ( ( zend_long ) ( ( ( char * ) ( & ( ( ( p_type ) NULL ) -> field ) ) ) - ( ( char * ) NULL ) ) )

// #define XtOffsetOf(s_type,field) XtOffset ( s_type * , field )

// #define ALLOCA_FLAG(name)

// #define SET_ALLOCA_FLAG(name)

// #define do_alloca(p,use_heap) emalloc ( p )

// #define free_alloca(p,use_heap) efree ( p )

// #define SETJMP(a) sigsetjmp ( a , 0 )

// #define LONGJMP(a,b) siglongjmp ( a , b )

// #define JMP_BUF       sigjmp_buf

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

// #define ZTS_V       0

// #define LONG_MAX       2147483647L

// #define LONG_MIN       ( - LONG_MAX - 1 )

// #define MAX_LENGTH_OF_DOUBLE       32

// #define MAX(a,b) ( ( ( a ) > ( b ) ) ? ( a ) : ( b ) )

// #define MIN(a,b) ( ( ( a ) < ( b ) ) ? ( a ) : ( b ) )

// #define ZEND_BIT_TEST(bits,bit) ( ( ( bits ) [ ( bit ) / ( sizeof ( ( bits ) [ 0 ] ) * 8 ) ] >> ( ( bit ) & ( sizeof ( ( bits ) [ 0 ] ) * 8 - 1 ) ) ) & 1 )

/* We always define a __special__  function, even if there's a macro or expression we could
 * alias, so that using it in contexts where we can't make function calls
 * won't fail to compile on some machines and not others.
 */

func _zendGetInf() float64 { return HUGE_VAL }

// #define ZEND_INFINITY       ( _zend_get_inf ( ) )

func _zendGetNan() float64 { return 0.0 / 0.0 }

// #define ZEND_NAN       ( _zend_get_nan ( ) )

// #define ZEND_STRL(str) ( str ) , ( sizeof ( str ) - 1 )

// #define ZEND_STRS(str) ( str ) , ( sizeof ( str ) )

// #define ZEND_NORMALIZE_BOOL(n) ( ( n ) ? ( ( ( n ) < 0 ) ? - 1 : 1 ) : 0 )

// #define ZEND_TRUTH(x) ( ( x ) ? 1 : 0 )

// #define ZEND_LOG_XOR(a,b) ( ZEND_TRUTH ( a ) ^ ZEND_TRUTH ( b ) )

// #define ZEND_MAX_RESERVED_RESOURCES       6

/* excpt.h on Digital Unix 4.0 defines function_table */

// #define ZEND_SECURE_ZERO(var,size) explicit_bzero ( ( var ) , ( size ) )

/* This check should only be used on network socket, not file descriptors */

// #define ZEND_VALID_SOCKET(sock) ( ( sock ) >= 0 )

/* va_copy() is __va_copy() in old gcc versions.
 * According to the autoconf manual, using
 * memcpy(&dst, &src, sizeof(va_list))
 * gives maximum portability. */

// #define va_copy(dest,src) memcpy ( & ( dest ) , & ( src ) , sizeof ( va_list ) )

/* Intrinsics macros start. */

/* Do not use for conditional declaration of API functions! */

// #define ZEND_INTRIN_SSSE3_FUNC_DECL(func)

/* Do not use for conditional declaration of API functions! */

// #define ZEND_INTRIN_SSE4_2_FUNC_DECL(func)

/* Do not use for conditional declaration of API functions! */

// #define ZEND_INTRIN_AVX2_FUNC_DECL(func)

/* Intrinsics macros end. */

// #define ZEND_SET_ALIGNED(alignment,decl) decl

// #define ZEND_SLIDE_TO_ALIGNED(alignment,ptr) ( ( ( zend_uintptr_t ) ( ptr ) + ( ( alignment ) - 1 ) ) & ~ ( ( alignment ) - 1 ) )

// #define ZEND_SLIDE_TO_ALIGNED16(ptr) ZEND_SLIDE_TO_ALIGNED ( Z_UL ( 16 ) , ptr )

// #define ZEND_EXPAND_VA(code) code

/* On CPU with few registers, it's cheaper to reload value then use spill slot */

// #define ZEND_IGNORE_LEAKS_BEGIN()

// #define ZEND_IGNORE_LEAKS_END()
