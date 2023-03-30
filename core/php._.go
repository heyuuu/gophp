package core

import (
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
)

const PHP_API_VERSION = 20190902
const PHP_DEFAULT_CHARSET = "UTF-8"

/* Operating system family definition */

const PHP_OS_FAMILY = "Unknown"

/* PHP's DEBUG value must match Zend's ZEND_DEBUG value */

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

var ExplicitBzero = PhpExplicitBzero

/*
 * Then the ODBC support can use both iodbc and Solid,
 * uncomment this.
 * #define HAVE_ODBC __special__  (HAVE_IODBC|HAVE_SOLID)
 */

const INT_MAX = 2147483647
const INT_MIN = -INT_MAX - 1

/* double limits */
const EXEC_INPUT_BUF = 4096

/* macros */

const MAXPATHLEN = 256

/* global variables */

const PhpSleep = sleep

var Environ **byte

const PhpError = faults.Error

/* PHPAPI void php_error(int type, const char *format, ...); */

const PhpMemnstr = zend.ZendMemnstr

/* PHP-named Zend macro wrappers */

var PHP_MODULE_GLOBALS = ZEND_MODULE_GLOBALS

/* Output support */

/* Virtual current working directory support */

/* connection status states */

const PHP_CONNECTION_NORMAL = 0
const PHP_CONNECTION_ABORTED = 1
const PHP_CONNECTION_TIMEOUT = 2
