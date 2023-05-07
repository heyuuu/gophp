package core

import (
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
)

const PHP_API_VERSION = 20190902
const PHP_DEFAULT_CHARSET = "UTF-8"

/* Operating system family definition */

const PHP_OS_FAMILY = "Unknown"

const PHP_DIR_SEPARATOR = '/'
const PHP_EOL = "\n"

const INT_MAX = 2147483647
const INT_MIN = -INT_MAX - 1

/* double limits */
const EXEC_INPUT_BUF = 4096

/* macros */

const MAXPATHLEN = 256

/* global variables */

const PhpSleep = sleep

var Environ **byte

func PhpError(typ int, format string, args ...any) {
	faults.Error(typ, format, args...)
}

/* PHPAPI void php_error(int type, const char *format, ...); */

var PhpMemnstr = zend.ZendMemnstr

const PHP_CONNECTION_NORMAL = 0
const PHP_CONNECTION_ABORTED = 1
const PHP_CONNECTION_TIMEOUT = 2
