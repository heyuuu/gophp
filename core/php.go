// <<generate>>

package core

import (
	"sik/zend"
)

// Source: <main/php.h>

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
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

// #define PHP_H

// #define PHP_API_VERSION       20190902

// #define PHP_HAVE_STREAMS

// #define YYDEBUG       0

// #define PHP_DEFAULT_CHARSET       "UTF-8"

// # include "php_version.h"

// # include "zend.h"

// # include "zend_sort.h"

// # include "php_compat.h"

// # include "zend_API.h"

// #define php_sprintf       sprintf

/* Operating system family definition */

// #define PHP_OS_FAMILY       "Unknown"

/* PHP's DEBUG value must match Zend's ZEND_DEBUG value */

// #define PHP_DEBUG       ZEND_DEBUG

// #define PHPAPI

// #define THREAD_LS

// #define PHP_DIR_SEPARATOR       '/'

// #define PHP_EOL       "\n"

/* Windows specific defines */

// #define NDEBUG

// # include < assert . h >

// # include < alloca . h >

// # include < build - defs . h >

/*
 * This is a fast version of strlcpy which should be used, if you
 * know the size of the destination buffer and if you know
 * the length of the source string.
 *
 * size is the allocated number of bytes of dst
 * src_size is the number of bytes excluding the NUL of src
 */

// #define PHP_STRLCPY(dst,src,size,src_size) { size_t php_str_len ; if ( src_size >= size ) php_str_len = size - 1 ; else php_str_len = src_size ; memcpy ( dst , src , php_str_len ) ; dst [ php_str_len ] = '\0' ; }

// #define explicit_bzero       php_explicit_bzero

// #define CREATE_MUTEX(a,b)

// #define SET_MUTEX(a)

// #define FREE_MUTEX(a)

/*
 * Then the ODBC support can use both iodbc and Solid,
 * uncomment this.
 * #define HAVE_ODBC __special__  (HAVE_IODBC|HAVE_SOLID)
 */

// # include < stdlib . h >

// # include < ctype . h >

// # include < unistd . h >

// # include < stdarg . h >

// # include "php_stdint.h"

// # include "zend_hash.h"

// # include "zend_alloc.h"

// # include "zend_stack.h"

// # include < string . h >

// # include < pwd . h >

// # include < sys / param . h >

// # include < limits . h >

// #define INT_MAX       2147483647

// #define INT_MIN       ( - INT_MAX - 1 )

/* double limits */

// # include < float . h >

// #define PHP_DOUBLE_MAX_LENGTH       1080

// #define PHP_GCC_VERSION       ZEND_GCC_VERSION

// #define PHP_ATTRIBUTE_MALLOC       ZEND_ATTRIBUTE_MALLOC

// #define PHP_ATTRIBUTE_FORMAT       ZEND_ATTRIBUTE_FORMAT

// # include "snprintf.h"

// # include "spprintf.h"

// #define EXEC_INPUT_BUF       4096

// #define PHP_MIME_TYPE       "application/x-httpd-php"

/* macros */

// #define STR_PRINT(str) ( ( str ) ? ( str ) : "" )

// #define MAXPATHLEN       256

// #define php_ignore_value(x) ZEND_IGNORE_VALUE ( x )

/* global variables */

// #define PHP_SLEEP_NON_VOID

// #define php_sleep       sleep

var Environ **byte
var Phperror func(error *byte)

// # include "php_syslog.h"

// #define php_log_err(msg) php_log_err_with_severity ( msg , LOG_NOTICE )

var Debug func(format *byte, _ ...any) int
var Cfgparse func() int

// #define php_error       zend_error

// #define error_handling_t       zend_error_handling_t

func PhpSetErrorHandling(error_handling zend.ZendErrorHandlingT, exception_class *zend.ZendClassEntry) {
	zend.ZendReplaceErrorHandling(error_handling, exception_class, nil)
}
func PhpStdErrorHandling() {}

/* PHPAPI void php_error(int type, const char *format, ...); */

// #define zenderror       phperror

// #define zendlex       phplex

// #define phpparse       zendparse

// #define phprestart       zendrestart

// #define phpin       zendin

// #define php_memnstr       zend_memnstr

/* functions */

var PhpRegisterPreRequestShutdown func(func_ func(any), userdata any)

/* PHP-named Zend macro wrappers */

// #define PHP_FN       ZEND_FN

// #define PHP_MN       ZEND_MN

// #define PHP_NAMED_FUNCTION       ZEND_NAMED_FUNCTION

// #define PHP_FUNCTION       ZEND_FUNCTION

// #define PHP_METHOD       ZEND_METHOD

// #define PHP_RAW_NAMED_FE       ZEND_RAW_NAMED_FE

// #define PHP_NAMED_FE       ZEND_NAMED_FE

// #define PHP_FE       ZEND_FE

// #define PHP_DEP_FE       ZEND_DEP_FE

// #define PHP_FALIAS       ZEND_FALIAS

// #define PHP_DEP_FALIAS       ZEND_DEP_FALIAS

// #define PHP_ME       ZEND_ME

// #define PHP_MALIAS       ZEND_MALIAS

// #define PHP_ABSTRACT_ME       ZEND_ABSTRACT_ME

// #define PHP_ME_MAPPING       ZEND_ME_MAPPING

// #define PHP_FE_END       ZEND_FE_END

// #define PHP_MODULE_STARTUP_N       ZEND_MODULE_STARTUP_N

// #define PHP_MODULE_SHUTDOWN_N       ZEND_MODULE_SHUTDOWN_N

// #define PHP_MODULE_ACTIVATE_N       ZEND_MODULE_ACTIVATE_N

// #define PHP_MODULE_DEACTIVATE_N       ZEND_MODULE_DEACTIVATE_N

// #define PHP_MODULE_INFO_N       ZEND_MODULE_INFO_N

// #define PHP_MODULE_STARTUP_D       ZEND_MODULE_STARTUP_D

// #define PHP_MODULE_SHUTDOWN_D       ZEND_MODULE_SHUTDOWN_D

// #define PHP_MODULE_ACTIVATE_D       ZEND_MODULE_ACTIVATE_D

// #define PHP_MODULE_DEACTIVATE_D       ZEND_MODULE_DEACTIVATE_D

// #define PHP_MODULE_INFO_D       ZEND_MODULE_INFO_D

/* Compatibility macros */

// #define PHP_MINIT       ZEND_MODULE_STARTUP_N

// #define PHP_MSHUTDOWN       ZEND_MODULE_SHUTDOWN_N

// #define PHP_RINIT       ZEND_MODULE_ACTIVATE_N

// #define PHP_RSHUTDOWN       ZEND_MODULE_DEACTIVATE_N

// #define PHP_MINFO       ZEND_MODULE_INFO_N

// #define PHP_GINIT       ZEND_GINIT

// #define PHP_GSHUTDOWN       ZEND_GSHUTDOWN

// #define PHP_MINIT_FUNCTION       ZEND_MODULE_STARTUP_D

// #define PHP_MSHUTDOWN_FUNCTION       ZEND_MODULE_SHUTDOWN_D

// #define PHP_RINIT_FUNCTION       ZEND_MODULE_ACTIVATE_D

// #define PHP_RSHUTDOWN_FUNCTION       ZEND_MODULE_DEACTIVATE_D

// #define PHP_MINFO_FUNCTION       ZEND_MODULE_INFO_D

// #define PHP_GINIT_FUNCTION       ZEND_GINIT_FUNCTION

// #define PHP_GSHUTDOWN_FUNCTION       ZEND_GSHUTDOWN_FUNCTION

// #define PHP_MODULE_GLOBALS       ZEND_MODULE_GLOBALS

/* Output support */

// # include "main/php_output.h"

// # include "php_streams.h"

// # include "php_memory_streams.h"

// # include "fopen_wrappers.h"

/* Virtual current working directory support */

// # include "zend_virtual_cwd.h"

// # include "zend_constants.h"

/* connection status states */

// #define PHP_CONNECTION_NORMAL       0

// #define PHP_CONNECTION_ABORTED       1

// #define PHP_CONNECTION_TIMEOUT       2

// # include "php_reentrancy.h"
