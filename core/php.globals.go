// <<generate>>

package core

import (
	"sik/zend"
)

const PHP_API_VERSION = 20190902
const PHP_DEFAULT_CHARSET = "UTF-8"
const PhpSprintf = sprintf
const PHP_OS_FAMILY = "Unknown"
const PHP_DEBUG zend.ZendLong = ZEND_DEBUG
const PHP_DIR_SEPARATOR = '/'
const PHP_EOL = "\n"
const ExplicitBzero = PhpExplicitBzero
const INT_MAX = 2147483647
const INT_MIN = -INT_MAX - 1
const PHP_DOUBLE_MAX_LENGTH = 1080
const PHP_GCC_VERSION = zend.ZEND_GCC_VERSION
const EXEC_INPUT_BUF = 4096
const PHP_MIME_TYPE = "application/x-httpd-php"
const MAXPATHLEN = 256
const PhpSleep = sleep

var Environ **byte
var Phperror func(error *byte)
var Debug func(format *byte, _ ...any) int
var Cfgparse func() int

const PhpError = zend.ZendError
const ErrorHandlingT = zend_error_handling_t
const Zenderror = Phperror
const Zendlex = phplex
const Phpparse = zend.Zendparse
const Phprestart = zendrestart
const Phpin = zendin
const PhpMemnstr = zend.ZendMemnstr

var PhpRegisterPreRequestShutdown func(func_ func(any), userdata any)

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
const PHP_GINIT = zend.ZEND_GINIT
const PHP_GSHUTDOWN = zend.ZEND_GSHUTDOWN
const PHP_GSHUTDOWN_FUNCTION = zend.ZEND_GSHUTDOWN_FUNCTION
const PHP_MODULE_GLOBALS = ZEND_MODULE_GLOBALS
const PHP_CONNECTION_NORMAL = 0
const PHP_CONNECTION_ABORTED = 1
const PHP_CONNECTION_TIMEOUT = 2
