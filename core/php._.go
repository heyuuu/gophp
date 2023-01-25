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

const PHP_API_VERSION = 20190902
const PHP_DEFAULT_CHARSET = "UTF-8"
const PhpSprintf = sprintf

/* Operating system family definition */

const PHP_OS_FAMILY = "Unknown"

/* PHP's DEBUG value must match Zend's ZEND_DEBUG value */

const PHP_DEBUG zend.ZendLong = ZEND_DEBUG
const PHP_DIR_SEPARATOR = '/'
const PHP_EOL = "\n"

/* Windows specific defines */

/*
 * This is a fast version of strlcpy which should be used, if you
 * know the size of the destination buffer and if you know
 * the length of the source string.
 *
 * size is the allocated number of bytes of dst
 * src_size is the number of bytes excluding the NUL of src
 */

const ExplicitBzero = PhpExplicitBzero

/*
 * Then the ODBC support can use both iodbc and Solid,
 * uncomment this.
 * #define HAVE_ODBC __special__  (HAVE_IODBC|HAVE_SOLID)
 */

const INT_MAX = 2147483647
const INT_MIN = -INT_MAX - 1

/* double limits */

const PHP_DOUBLE_MAX_LENGTH = 1080
const PHP_GCC_VERSION = zend.ZEND_GCC_VERSION
const EXEC_INPUT_BUF = 4096
const PHP_MIME_TYPE = "application/x-httpd-php"

/* macros */

const MAXPATHLEN = 256

/* global variables */

const PhpSleep = sleep

var Environ **byte
var Phperror func(error *byte)
var Debug func(format *byte, _ ...any) int
var Cfgparse func() int

const PhpError = zend.ZendError
const ErrorHandlingT = zend_error_handling_t

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

const PHP_MN = ZEND_MN
const PHP_ABSTRACT_ME = ZEND_ABSTRACT_ME
const PHP_ME_MAPPING = ZEND_ME_MAPPING
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

const PHP_GINIT = zend.ZEND_GINIT
const PHP_GSHUTDOWN = zend.ZEND_GSHUTDOWN
const PHP_GSHUTDOWN_FUNCTION = zend.ZEND_GSHUTDOWN_FUNCTION
const PHP_MODULE_GLOBALS = ZEND_MODULE_GLOBALS

/* Output support */

/* Virtual current working directory support */

/* connection status states */

const PHP_CONNECTION_NORMAL = 0
const PHP_CONNECTION_ABORTED = 1
const PHP_CONNECTION_TIMEOUT = 2
