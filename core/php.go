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

const PHP_API_VERSION = 20190902

// #define PHP_HAVE_STREAMS

// #define YYDEBUG       0

const PHP_DEFAULT_CHARSET = "UTF-8"

// # include "php_version.h"

// # include "zend.h"

// # include "zend_sort.h"

// # include "php_compat.h"

// # include "zend_API.h"

const PhpSprintf = sprintf

/* Operating system family definition */

const PHP_OS_FAMILY = "Unknown"

/* PHP's DEBUG value must match Zend's ZEND_DEBUG value */

const PHP_DEBUG zend.ZendLong = ZEND_DEBUG

// #define PHPAPI

// #define THREAD_LS

const PHP_DIR_SEPARATOR = '/'
const PHP_EOL = "\n"

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

func PHP_STRLCPY(dst []char, src __auto__, size int, src_size int) {
	var php_str_len int
	if src_size >= size {
		php_str_len = size - 1
	} else {
		php_str_len = src_size
	}
	memcpy(dst, src, php_str_len)
	dst[php_str_len] = '0'
}

const ExplicitBzero = PhpExplicitBzero

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

const INT_MAX = 2147483647
const INT_MIN = -INT_MAX - 1

/* double limits */

// # include < float . h >

const PHP_DOUBLE_MAX_LENGTH = 1080
const PHP_GCC_VERSION = zend.ZEND_GCC_VERSION

// #define PHP_ATTRIBUTE_MALLOC       ZEND_ATTRIBUTE_MALLOC

// #define PHP_ATTRIBUTE_FORMAT       ZEND_ATTRIBUTE_FORMAT

// # include "snprintf.h"

// # include "spprintf.h"

const EXEC_INPUT_BUF = 4096
const PHP_MIME_TYPE = "application/x-httpd-php"

/* macros */

func STR_PRINT(str *byte) string {
	if str != nil {
		return str
	} else {
		return ""
	}
}

const MAXPATHLEN = 256

func PhpIgnoreValue(x __auto__) { zend.ZEND_IGNORE_VALUE(x) }

/* global variables */

// #define PHP_SLEEP_NON_VOID

const PhpSleep = sleep

var Environ **byte
var Phperror func(error *byte)

// # include "php_syslog.h"

func PhpLogErr(msg *byte) { PhpLogErrWithSeverity(msg, LOG_NOTICE) }

var Debug func(format *byte, _ ...any) int
var Cfgparse func() int

const PhpError = zend.ZendError
const ErrorHandlingT = zend_error_handling_t

func PhpSetErrorHandling(error_handling error_handling_t, exception_class *zend.ZendClassEntry) {
	zend.ZendReplaceErrorHandling(error_handling, exception_class, nil)
}
func PhpStdErrorHandling() {}

/* PHPAPI void php_error(int type, const char *format, ...); */

const Zenderror = Phperror
const Zendlex = phplex
const Phpparse = zend.Zendparse
const Phprestart = zendrestart
const Phpin = zendin
const PhpMemnstr = zend.ZendMemnstr

/* functions */

var PhpRegisterPreRequestShutdown func(func_ func(any), userdata any)

/* PHP-named Zend macro wrappers */

// #define PHP_FN       ZEND_FN

const PHP_MN = ZEND_MN

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

const PHP_ABSTRACT_ME = ZEND_ABSTRACT_ME
const PHP_ME_MAPPING = ZEND_ME_MAPPING

// #define PHP_FE_END       ZEND_FE_END

const PHP_MODULE_STARTUP_N = ZEND_MODULE_STARTUP_N
const PHP_MODULE_SHUTDOWN_N = ZEND_MODULE_SHUTDOWN_N
const PHP_MODULE_ACTIVATE_N = ZEND_MODULE_ACTIVATE_N
const PHP_MODULE_DEACTIVATE_N = ZEND_MODULE_DEACTIVATE_N
const PHP_MODULE_INFO_N = ZEND_MODULE_INFO_N
const PHP_MODULE_STARTUP_D = ZEND_MODULE_STARTUP_D
const PHP_MODULE_SHUTDOWN_D = ZEND_MODULE_SHUTDOWN_D
const PHP_MODULE_ACTIVATE_D = ZEND_MODULE_ACTIVATE_D
const PHP_MODULE_DEACTIVATE_D = ZEND_MODULE_DEACTIVATE_D
const PHP_MODULE_INFO_D = ZEND_MODULE_INFO_D

/* Compatibility macros */

// #define PHP_MINIT       ZEND_MODULE_STARTUP_N

// #define PHP_MSHUTDOWN       ZEND_MODULE_SHUTDOWN_N

// #define PHP_RINIT       ZEND_MODULE_ACTIVATE_N

// #define PHP_RSHUTDOWN       ZEND_MODULE_DEACTIVATE_N

// #define PHP_MINFO       ZEND_MODULE_INFO_N

const PHP_GINIT = zend.ZEND_GINIT
const PHP_GSHUTDOWN = zend.ZEND_GSHUTDOWN

// #define PHP_MINIT_FUNCTION       ZEND_MODULE_STARTUP_D

// #define PHP_MSHUTDOWN_FUNCTION       ZEND_MODULE_SHUTDOWN_D

// #define PHP_RINIT_FUNCTION       ZEND_MODULE_ACTIVATE_D

// #define PHP_RSHUTDOWN_FUNCTION       ZEND_MODULE_DEACTIVATE_D

// #define PHP_MINFO_FUNCTION       ZEND_MODULE_INFO_D

// #define PHP_GINIT_FUNCTION       ZEND_GINIT_FUNCTION

const PHP_GSHUTDOWN_FUNCTION = zend.ZEND_GSHUTDOWN_FUNCTION
const PHP_MODULE_GLOBALS = ZEND_MODULE_GLOBALS

/* Output support */

// # include "main/php_output.h"

// # include "php_streams.h"

// # include "php_memory_streams.h"

// # include "fopen_wrappers.h"

/* Virtual current working directory support */

// # include "zend_virtual_cwd.h"

// # include "zend_constants.h"

/* connection status states */

const PHP_CONNECTION_NORMAL = 0
const PHP_CONNECTION_ABORTED = 1
const PHP_CONNECTION_TIMEOUT = 2

// # include "php_reentrancy.h"
